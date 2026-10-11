package com.ciagotech.kan.server.auth;

import org.springframework.data.redis.core.StringRedisTemplate;
import java.time.Duration;
import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.Objects;

/** Atomic SET NX EX through Spring Data Redis; errors propagate and fail closed. */
public final class ValkeyNonceStore implements NonceStore {
    private final StringRedisTemplate redis;
    public ValkeyNonceStore(StringRedisTemplate redis) { this.redis = Objects.requireNonNull(redis); }
    @Override
    public boolean checkAndStore(String accessKeyId, String nonce, int ttlSeconds) {
        if (ttlSeconds < 1) throw new IllegalArgumentException("TTL must be positive");
        var encoder = Base64.getUrlEncoder().withoutPadding();
        String key = "kan:nonce:" + encoder.encodeToString(accessKeyId.getBytes(StandardCharsets.UTF_8))
                + ":" + encoder.encodeToString(nonce.getBytes(StandardCharsets.UTF_8));
        Boolean stored = redis.opsForValue().setIfAbsent(key, "1", Duration.ofSeconds(ttlSeconds));
        if (stored == null) throw new IllegalStateException("Missing nonce-store response");
        return stored;
    }
}
