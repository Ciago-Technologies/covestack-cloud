package com.ciagotech.kan.server.auth;

/** Person C implements hash-only token lookup, expiry/revocation and account membership. */
@FunctionalInterface
public interface BearerTokenService {
    /** Null for unknown, expired or revoked tokens; never store the plaintext token. */
    KeyService.Principal findActivePrincipal(String opaqueToken);
}
