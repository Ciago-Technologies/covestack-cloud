package com.ciagotech.kan.server.auth;

/** Person C implements active Postgres lookup and AES-GCM secret decryption. */
@FunctionalInterface
public interface KeyService {
    /** Null for unknown/revoked keys. Never return a default development credential. */
    Credential findActiveCredential(String accessKeyId);

    record Credential(String accountId, String accessKeyId, String kind, String secretKey) {}
    record Principal(String accountId, String accessKeyId, String kind) {}
}
