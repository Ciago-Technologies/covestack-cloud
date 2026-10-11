package com.ciagotech.kan.server.auth;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.*;
import java.util.stream.Collectors;

/**
 * Same algorithm as kan-agent's Go sigv1 package — see CANONICAL-REQUEST.md.
 * Used two ways:
 *   - sign(): by the kan CLI when it calls the Kan-native REST API
 *   - verify(): by KanSigV1AuthFilter, on every incoming agent request
 */
public final class KanSigV1 {

    private KanSigV1() {}

    public record CanonicalResult(String canonicalRequest, String signedHeaders) {}

    public static CanonicalResult buildCanonicalRequest(String method, String path, String query,
            Map<String, String> headers, byte[] body, SigV1Rules rules) {

        Set<String> signSet = new TreeSet<>();
        for (String h : rules.alwaysSignedHeaders()) signSet.add(h.toLowerCase());
        if (body != null && body.length > 0 && rules.signBodyContentTypeHeader()) {
            signSet.add("content-type");
        }

        Map<String, String> lower = new HashMap<>();
        for (var e : headers.entrySet()) {
            if (e.getValue() != null) lower.put(e.getKey().toLowerCase(), e.getValue().trim());
        }

        List<String> names = signSet.stream()
                .filter(lower::containsKey)
                .sorted()
                .collect(Collectors.toList());

        StringBuilder canonicalHeaders = new StringBuilder();
        for (String n : names) canonicalHeaders.append(n).append(":").append(lower.get(n)).append("\n");
        String signedHeaders = String.join(";", names);
        String hashedPayload = sha256Hex(body == null ? new byte[0] : body);

        String canonicalRequest = String.join("\n",
                method, path, query, canonicalHeaders.toString(), signedHeaders, hashedPayload);

        return new CanonicalResult(canonicalRequest, signedHeaders);
    }

    public static String stringToSign(SigV1Rules rules, String dateStr, String canonicalRequest) {
        return String.join("\n", rules.algorithm(), dateStr,
                sha256Hex(canonicalRequest.getBytes(StandardCharsets.UTF_8)));
    }

    public static String sign(String secretKey, String method, String path, String query,
            Map<String, String> headers, byte[] body, String dateStr, SigV1Rules rules) {
        CanonicalResult cr = buildCanonicalRequest(method, path, query, headers, body, rules);
        String sts = stringToSign(rules, dateStr, cr.canonicalRequest());
        return hmacSha256Hex(secretKey, sts);
    }

    /** Constant-time compare — do not replace with String.equals(), that leaks timing info. */
    public static boolean verify(String providedSignature, String secretKey, String method, String path,
            String query, Map<String, String> headers, byte[] body, String dateStr, SigV1Rules rules) {
        String expected = sign(secretKey, method, path, query, headers, body, dateStr, rules);
        return constantTimeEquals(expected, providedSignature);
    }

    private static boolean constantTimeEquals(String a, String b) {
        if (a == null || b == null) return false;
        return MessageDigest.isEqual(
                a.getBytes(StandardCharsets.UTF_8),
                b.getBytes(StandardCharsets.UTF_8));
    }

    private static String sha256Hex(byte[] data) {
        try {
            var md = MessageDigest.getInstance("SHA-256");
            return toHex(md.digest(data));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    private static String hmacSha256Hex(String secretKey, String data) {
        try {
            Mac mac = Mac.getInstance("HmacSHA256");
            mac.init(new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), "HmacSHA256"));
            return toHex(mac.doFinal(data.getBytes(StandardCharsets.UTF_8)));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    private static String toHex(byte[] bytes) {
        StringBuilder sb = new StringBuilder(bytes.length * 2);
        for (byte b : bytes) sb.append(String.format("%02x", b));
        return sb.toString();
    }
}
