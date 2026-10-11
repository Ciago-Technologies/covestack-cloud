package com.ciagotech.kan.server.auth;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.ServletRequest;
import jakarta.servlet.ServletResponse;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.time.Instant;
import java.time.ZoneOffset;
import java.time.format.DateTimeFormatter;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

/**
 * Runs once per request, before any controller. Only guards the
 * /internal/v1/agent/* routes here — the AWS SigV4 side reuses the same
 * KanSigV1 crypto but a different header set, in its own filter.
 *
 * Order of checks (cheapest / hardest-to-fake first, so a request gets
 * rejected before spending effort on the expensive parts):
 *   1. required headers present
 *   2. clock skew acceptable
 *   3. nonce not replayed
 *   4. access key known
 *   5. signature matches
 */
public class KanSigV1AuthFilter extends OncePerRequestFilter {

    private static final DateTimeFormatter DATE_FMT =
            DateTimeFormatter.ofPattern("yyyyMMdd'T'HHmmss'Z'").withZone(ZoneOffset.UTC);

    private final SigV1Rules rules;
    private final KeyService keyService;
    private final NonceStore nonceStore;

    public KanSigV1AuthFilter(SigV1Rules rules, KeyService keyService, NonceStore nonceStore) {
        this.rules = rules;
        this.keyService = keyService;
        this.nonceStore = nonceStore;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response,
            FilterChain chain) throws ServletException, IOException {

        if (!request.getRequestURI().startsWith("/internal/v1/agent/")) {
            chain.doFilter(request, response);
            return;
        }

        String accessKeyId = request.getHeader(rules.accessKeyHeader());
        String dateStr = request.getHeader(rules.dateHeader());
        String nonce = request.getHeader(rules.nonceHeader());
        String providedSignature = request.getHeader(rules.signatureHeader());

        if (accessKeyId == null || dateStr == null || nonce == null || providedSignature == null) {
            reject(response, "missing required signing headers");
            return;
        }

        Instant reqTime;
        try {
            reqTime = Instant.from(DATE_FMT.parse(dateStr));
        } catch (Exception e) {
            reject(response, "bad date format");
            return;
        }
        long skewSeconds = Math.abs(Instant.now().getEpochSecond() - reqTime.getEpochSecond());
        if (skewSeconds > rules.maxClockSkewSeconds()) {
            reject(response, "clock skew too large");
            return;
        }

        if (!nonceStore.checkAndStore(accessKeyId, nonce, rules.nonceReplayWindowSeconds())) {
            reject(response, "replayed nonce");
            return;
        }

        String secretKey = keyService.getSecretKey(accessKeyId);
        if (secretKey == null) {
            reject(response, "unknown or inactive access key");
            return;
        }

        // Wrap so the body can be read here AND again downstream by Jackson.
        CachedBodyHttpServletRequest cachedRequest = new CachedBodyHttpServletRequest(request);
        byte[] body = cachedRequest.getCachedBody();

        Map<String, String> headers = new HashMap<>();
        headers.put("host", request.getHeader("Host"));
        headers.put(rules.dateHeader().toLowerCase(), dateStr);
        String contentType = request.getContentType();
        if (contentType != null) headers.put("content-type", contentType);

        boolean ok = KanSigV1.verify(providedSignature, secretKey, request.getMethod(),
                request.getRequestURI(), Optional.ofNullable(request.getQueryString()).orElse(""),
                headers, body, dateStr, rules);

        if (!ok) {
            reject(response, "signature mismatch");
            return;
        }

        chain.doFilter(cachedRequest, response);
    }

    private void reject(HttpServletResponse response, String reason) throws IOException {
        response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
        response.setContentType("application/json");
        response.getWriter().write("{\"error\":\"" + reason + "\"}");
    }
}
