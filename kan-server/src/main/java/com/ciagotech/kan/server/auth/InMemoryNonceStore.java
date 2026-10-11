package com.ciagotech.kan.server.auth;

import java.time.Clock;
import java.util.HashMap;
import java.util.Map;

/** Bounded, expiring DEV/TEST ONLY store. Never wire into production. */
public class InMemoryNonceStore implements NonceStore {
    private final Map<String, Long> seen = new HashMap<>();
    private final Clock clock;
    private final int capacity;
    public InMemoryNonceStore() { this(Clock.systemUTC(), 10000); }
    public InMemoryNonceStore(Clock clock, int capacity) {
        if (capacity < 1) throw new IllegalArgumentException("capacity must be positive");
        this.clock = clock;
        this.capacity = capacity;
    }
    @Override
    public synchronized boolean checkAndStore(String accessKeyId, String nonce, int ttlSeconds) {
        if (ttlSeconds < 1) throw new IllegalArgumentException("TTL must be positive");
        long now = clock.millis();
        seen.entrySet().removeIf(e -> e.getValue() <= now);
        String key = accessKeyId.length() + ":" + accessKeyId + nonce;
        if (seen.containsKey(key)) return false;
        if (seen.size() >= capacity) throw new IllegalStateException("nonce store capacity exhausted");
        seen.put(key, now + ttlSeconds * 1000L);
        return true;
    }
}
