# Covestack Cloud — kan sigv1 canonical request

Shared contract for Anuj's Go signer and Person C's Java verifier. This revision
changes signatures: deploy Go, Java, rules and vectors together. Old signatures fail.

## Canonical request

Join method, raw escaped path, raw query (without `?`), canonical headers,
signed-header names and lowercase SHA256 hex of the raw body with `\n`.
Canonical headers themselves end with `\n`, so there is a blank line before names.
Do not decode/re-encode the path, sort the query, or serialize the body after signing.
Java uses request URI/query; Go callers use URL.EscapedPath()/URL.RawQuery.

Always sign `host`, `x-kan-access-key`, `x-kan-date`, `x-kan-nonce`.
For a nonempty body also sign `content-type`. Select only these names from the
shared rules, lowercase names, trim only outer ASCII space/tab, sort names, then
write `name:value\n`. Join the same names with `;`. Required values must be
present and nonempty; duplicate case-insensitive names or CR/LF values are invalid.
Other headers do not affect the signature. Use UTF-8 for textual hash/HMAC input.

String to sign: `KAN1-HMAC-SHA256\n<date>\n<SHA256_HEX(canonical request)>`.
Signature: lowercase HMAC-SHA256 hex, using the raw secret bytes, sent only in
`X-Kan-Signature`. Date: strict UTC `yyyyMMdd'T'HHmmss'Z'`.

## Transport and replay protection

Use a service credential for internal agent calls. Nonces contain 16–128 URL-safe
ASCII characters; the Go request helper generates 32 cryptographically random
bytes as hex. Each request/retry gets a fresh date/nonce/signature.
The server verifies headers, time, active credential, body bound and MAC first,
then atomically reserves the nonce in Valkey before any controller runs.
Clock skew is at most 300 seconds; nonce retention is 600 seconds.
Maximum signed body size is 1 MiB, including chunked bodies.

Authenticated requests carry `kan.principal` (accountId/accessKeyId/kind).
Controllers MUST scope reads/writes to this account. Public `/v2/` and `/kan/v1/`
APIs are registered with the separate BearerAuthFilter; Person C must implement
BearerTokenService lookup. Neither filter replaces controller authorization.

## Compatibility gate

From repository root: `python3 contracts/sign_example.py`,
`cd kan-agent && go test ./...`, `cd kan-server && mvn verify`.
Shared vectors include empty/body/UTF-8/raw-query cases. Negative filter tests cover
nonce tampering, replay, poisoning, oversized bodies and invalid credential kinds.
