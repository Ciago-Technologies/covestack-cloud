package com.ciagotech.kan.server.auth;

/**
 * Replay protection. Production implementation must be Valkey:
 *   SET kan:nonce:{accessKeyId}:{nonce} 1 NX EX {ttlSeconds}
 * (see FROZEN-STATES.md / your architecture doc — this is already
 * specified, not a new idea). See InMemoryNonceStore for why that
 * class must never reach production.
 */
public interface NonceStore {
    /** Returns true if this nonce is fresh (and records it). False = replay, reject the request. */
    boolean checkAndStore(String accessKeyId, String nonce, int ttlSeconds);
}
