package com.ciagotech.kan.server.auth;

/**
 * Looks up the secret key for a signed request. Real implementation reads
 * access_keys.secret_enc from Postgres and decrypts with AES-256-GCM
 * (per the schema — see TDD key management section). NOT implemented here.
 */
public interface KeyService {
    /** Returns the decrypted secret, or null if the access key is unknown or inactive. */
    String getSecretKey(String accessKeyId);
}
