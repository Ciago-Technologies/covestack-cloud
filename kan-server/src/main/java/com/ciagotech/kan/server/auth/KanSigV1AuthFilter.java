package com.ciagotech.kan.server.auth;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.filter.OncePerRequestFilter;
import java.io.IOException;
import java.time.Instant;
import java.time.ZoneOffset;
import java.time.format.DateTimeFormatter;
import java.time.format.ResolverStyle;
import java.util.*;

/** Agent authentication only. Public APIs require a separate Bearer filter. */
public class KanSigV1AuthFilter extends OncePerRequestFilter {
    private static final DateTimeFormatter DATE_FMT = DateTimeFormatter
            .ofPattern("uuuuMMdd'T'HHmmss'Z'").withResolverStyle(ResolverStyle.STRICT)
            .withZone(ZoneOffset.UTC);
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
        String route = request.getServletPath();
        if (route.isEmpty()) route = request.getRequestURI().substring(request.getContextPath().length());
        if (!(route.equals("/internal/v1/agent") || route.startsWith("/internal/v1/agent/"))) {
            chain.doFilter(request, response);
            return;
        }
        String accessKeyId, dateStr, nonce, signature;
        Map<String, String> headers = new HashMap<>();
        try {
            accessKeyId = singleHeader(request, rules.accessKeyHeader());
            dateStr = singleHeader(request, rules.dateHeader());
            nonce = singleHeader(request, rules.nonceHeader());
            signature = singleHeader(request, rules.signatureHeader());
            if (!accessKeyId.matches("[A-Za-z0-9_-]{1,128}") || !nonce.matches("[A-Za-z0-9_-]{16,128}")
                    || !signature.matches("[0-9a-f]{64}") || !dateStr.matches("[0-9]{8}T[0-9]{6}Z"))
                throw new IllegalArgumentException();
            for (String name : rules.alwaysSignedHeaders()) headers.put(name, singleHeader(request, name));
            if (request.getHeader("Content-Type") != null)
                headers.put("content-type", singleHeader(request, "Content-Type"));
            Instant timestamp = Instant.from(DATE_FMT.parse(dateStr));
            long skew = Math.abs(Instant.now().getEpochSecond() - timestamp.getEpochSecond());
            if (skew > rules.maxClockSkewSeconds()) throw new IllegalArgumentException();
        } catch (RuntimeException e) {
            reject(response, 401, "invalid signing headers");
            return;
        }
        KeyService.Credential credential = keyService.findActiveCredential(accessKeyId);
        if (credential == null || !accessKeyId.equals(credential.accessKeyId())
                || !"service".equals(credential.kind()) || credential.accountId() == null
                || credential.accountId().isBlank() || credential.secretKey() == null
                || credential.secretKey().isEmpty()) {
            reject(response, 401, "invalid agent credential");
            return;
        }
        CachedBodyHttpServletRequest cached;
        try {
            cached = new CachedBodyHttpServletRequest(request);
        } catch (CachedBodyHttpServletRequest.BodyTooLargeException e) {
            reject(response, 413, "body too large");
            return;
        }
        try {
            if (!KanSigV1.verify(signature, credential.secretKey(), request.getMethod(),
                    request.getRequestURI(), Optional.ofNullable(request.getQueryString()).orElse(""),
                    headers, cached.getCachedBody(), dateStr, rules)) {
                reject(response, 401, "invalid signature");
                return;
            }
        } catch (IllegalArgumentException e) {
            reject(response, 401, "invalid signed headers");
            return;
        }
        // Reserve AFTER MAC validation, BEFORE controllers or other side effects.
        try {
            if (!nonceStore.checkAndStore(accessKeyId, nonce, rules.nonceReplayWindowSeconds())) {
                reject(response, 401, "replayed nonce");
                return;
            }
        } catch (RuntimeException e) {
            reject(response, 503, "replay protection unavailable");
            return;
        }
        cached.setAttribute("kan.principal", new KeyService.Principal(
                credential.accountId(), accessKeyId, credential.kind()));
        chain.doFilter(cached, response);
    }

    private static String singleHeader(HttpServletRequest request, String name) {
        Enumeration<String> values = request.getHeaders(name);
        if (values == null || !values.hasMoreElements()) throw new IllegalArgumentException();
        String value = values.nextElement();
        if (values.hasMoreElements() || value.isBlank() || value.length() > 1024
                || value.contains("\r") || value.contains("\n")) throw new IllegalArgumentException();
        return value.replaceAll("^[ \t]+|[ \t]+$", "");
    }

    private static void reject(HttpServletResponse response, int status, String reason) throws IOException {
        response.setStatus(status);
        response.setContentType("application/json");
        response.getWriter().write("{\"error\":\"" + reason + "\"}");
    }
}
