package com.ciagotech.kan.server.auth;

import java.time.Instant;
import java.util.concurrent.ConcurrentHashMap;

/**
 * DEV/TEST ONLY — good enough to run the FakeAgent locally and to make
 * this filter compile and testable today. Do NOT deploy this: it never
 * expires entries (memory leak) and doesn't survive a restart, which
 * defeats replay protection across a deploy. Swap for real Valkey
 * (SET NX EX) before Week 3 per the roadmap — this is exactly the kind
 * of stub the "kill list" is meant to catch if it's still here at D21.
 */
public class InMemoryNonceStore implements NonceStore {
    private final ConcurrentHashMap<String, Instant> seen = new ConcurrentHashMap<>();

    @Override
    public boolean checkAndStore(String accessKeyId, String nonce, int ttlSeconds) {
        String key = accessKeyId + ":" + nonce;
        Instant prior = seen.putIfAbsent(key, Instant.now());
        return prior == null;
    }
}
