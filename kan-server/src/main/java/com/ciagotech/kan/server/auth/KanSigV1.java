package com.ciagotech.kan.server.auth;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.*;

/**
 * Same algorithm as kan-agent's Go sigv1 package — see CANONICAL-REQUEST.md.
 * sign() supports shared-vector tests; verify() authenticates internal agent requests.
 * Public CLI APIs use separate Bearer authentication.
 */
public final class KanSigV1 {

    private KanSigV1() {}

    public record CanonicalResult(String canonicalRequest, String signedHeaders) {}

    public static CanonicalResult buildCanonicalRequest(String method, String path, String query,
            Map<String, String> headers, byte[] body, SigV1Rules rules) {

        Set<String> signSet = new TreeSet<>();
        for (String h : rules.alwaysSignedHeaders()) signSet.add(h.toLowerCase(Locale.ROOT));
        if (body != null && body.length > 0 && rules.signBodyContentTypeHeader()) {
            signSet.add("content-type");
        }

        Map<String, String> lower = new HashMap<>();
        for (var e : headers.entrySet()) {
            String name = e.getKey().toLowerCase(Locale.ROOT);
            if (lower.containsKey(name)) throw new IllegalArgumentException("duplicate header: " + name);
            lower.put(name, e.getValue() == null ? null : e.getValue().replaceAll("^[ \t]+|[ \t]+$", ""));
        }

        for (String name : signSet) {
            String value = lower.get(name);
            if (value == null || value.isEmpty() || value.contains("\r") || value.contains("\n"))
                throw new IllegalArgumentException("missing or invalid signed header: " + name);
        }
        List<String> names = new ArrayList<>(signSet);

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
        String headerDate = headers.entrySet().stream()
                .filter(e -> e.getKey().equalsIgnoreCase(rules.dateHeader()))
                .map(Map.Entry::getValue).filter(Objects::nonNull).findFirst().orElse("")
                .replaceAll("^[ \t]+|[ \t]+$", "");
        if (!headerDate.equals(dateStr) || secretKey == null || secretKey.isEmpty())
            throw new IllegalArgumentException("date mismatch or empty signing secret");
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
