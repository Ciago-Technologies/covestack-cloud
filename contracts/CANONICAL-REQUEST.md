# CoveStack Cloud — Canonical Request and Signing Guide

## What this file means

| Term              | Plain meaning                                                                                                                         |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Canonical request | An exact text representation of the request. Go and Java must produce identical bytes.                                                |
| Body hash         | SHA-256 of the exact bytes being sent; formatting JSON differently changes it.                                                        |
| String to sign    | Algorithm name, request date and canonical-request hash, separated by newline bytes.                                                  |
| Signature         | HMAC-SHA256 of the string to sign using the secret key; sent as lowercase hexadecimal.                                                |
| Owner             | Anuj builds the Go signer; C builds the Java verifier; B enforces scoped controller/service behavior.                                 |
| Branding          | CoveStack is the product name. Existing `KAN1`, `KAN_*` and `X-Kan-*` identifiers remain unchanged.                                   |
| Compatibility     | The signed-nonce/access-key security revision already exists. This documentation rewrite does not change the algorithm or signatures. |

## Required request headers

| Header             | Example                         | Signed?             | Rule                                                                               |
| ------------------ | ------------------------------- | ------------------- | ---------------------------------------------------------------------------------- |
| `Host`             | `localhost:4700`                | Always              | Match the authority sent to the verifier, including port when present.             |
| `X-Kan-Access-Key` | `KANKAGENT0000000001`           | Always              | Public credential ID; internal calls require an active service credential.         |
| `X-Kan-Date`       | `20260926T120000Z`              | Always              | Strict UTC `yyyyMMdd'T'HHmmss'Z'`; maximum skew 300 seconds.                       |
| `X-Kan-Nonce`      | `vector-nonce-0000000000000001` | Always              | 16–128 URL-safe ASCII characters; use a fresh random nonce for each request/retry. |
| `Content-Type`     | `application/json`              | For a nonempty body | Required and signed when body bytes are nonempty.                                  |
| `X-Kan-Signature`  | Calculated HMAC hex             | No                  | Carries the result; do not include it in canonical headers.                        |
| Other headers      | `User-Agent`                    | No                  | Do not add them to the signed-header set.                                          |

## Build the canonical request in this order

| Position | Component           | Exact rule                                                                                                                                          |
| -------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1        | HTTP method         | Actual method, such as `GET` or `PUT`.                                                                                                              |
| 2        | Raw escaped path    | Preserve escaping; Go uses `URL.EscapedPath()`, Java uses request URI.                                                                              |
| 3        | Raw query           | Preserve ordering and encoding, without `?`; empty query is an empty line.                                                                          |
| 4        | Canonical headers   | Lowercase selected names, sort alphabetically, trim only outer ASCII space/tab from values; write `name:value` followed by newline for each header. |
| 5        | Signed-header names | Same sorted names joined with `;`.                                                                                                                  |
| 6        | Body hash           | Lowercase SHA-256 hexadecimal of the exact raw body bytes.                                                                                          |

| Formatting detail       | Required behavior                                                                                 |
| ----------------------- | ------------------------------------------------------------------------------------------------- |
| Separator               | Join the six components with actual newline bytes (`\n`), not the two characters backslash and n. |
| Blank line before names | Canonical headers already end in newline; joining the next component adds another newline.        |
| No final newline        | Do not append a newline after the body-hash component.                                            |
| Validation              | Reject missing/empty required headers, case-insensitive duplicate names and CR/LF values.         |
| Text encoding           | UTF-8 for the canonical request and string to sign.                                               |
| Body                    | Serialize once, sign those bytes, send those same bytes.                                          |
| Query/path              | Do not sort query parameters, decode/re-encode the path or normalize escapes.                     |

## Calculate the signature

| Step | Formula                                                                                                         |
| ---- | --------------------------------------------------------------------------------------------------------------- |
| 1    | `canonicalHash = SHA256_HEX(UTF8(canonicalRequest))`                                                            |
| 2    | `stringToSign = "KAN1-HMAC-SHA256" + "\n" + date + "\n" + canonicalHash`                                        |
| 3    | `signature = HMAC_SHA256_HEX(rawSecretBytes, UTF8(stringToSign))`                                               |
| 4    | Put the lowercase hex result in `X-Kan-Signature`. No AWS-style key derivation or Authorization branch is used. |

## Worked example: GET agent work

| Input            | Value                                                                                         |
| ---------------- | --------------------------------------------------------------------------------------------- |
| Method / path    | `GET /internal/v1/agent/work`                                                                 |
| Query / body     | Both empty                                                                                    |
| Test-only secret | `test-secret-do-not-use-in-prod`                                                              |
| Source           | `get-work-basic` in [sigv1-vectors.json](sigv1-vectors.json)                                  |
| Scope            | Offline deterministic example; the historical timestamp will fail live clock-skew validation. |

The exact canonical text is below. Keep the blank query line and the blank line before signed-header names.

```text
GET
/internal/v1/agent/work

host:localhost:4700
x-kan-access-key:KANKAGENT0000000001
x-kan-date:20260926T120000Z
x-kan-nonce:vector-nonce-0000000000000001

host;x-kan-access-key;x-kan-date;x-kan-nonce
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```

| Stage              | Result                                                             |
| ------------------ | ------------------------------------------------------------------ |
| Signed headers     | `host;x-kan-access-key;x-kan-date;x-kan-nonce`                     |
| Canonical hash     | `179f1b6fcc31e5476c6ca4e5e3ec1843eb030e9f1e675f5528d73cb572ec3e9e` |
| Expected signature | `1e34d04c4dd572f7b8f21596ecafce1dbea7cb8c7a464b3d5c776ffbfcd69a81` |

Exact string to sign:

```text
KAN1-HMAC-SHA256
20260926T120000Z
179f1b6fcc31e5476c6ca4e5e3ec1843eb030e9f1e675f5528d73cb572ec3e9e
```

## Reproduce the example locally

Run from the repository root. This uses the existing reference signer and prints no real credentials.

```bash
python3 - <<'PYTHON'
import json
from contracts.sign_example import sign
with open("contracts/sigv1-vectors.json") as stream:
    vector = json.load(stream)["vectors"][0]
canonical, string_to_sign, signature, names = sign(
    vector["secretKey"], vector["method"], vector["path"],
    vector["query"], vector["headers"], vector["body"],
    vector["headers"]["X-Kan-Date"],
)
assert signature == vector["expectedSignature"]
print("CANONICAL REQUEST:\n" + canonical)
print("STRING TO SIGN:\n" + string_to_sign)
print("SIGNATURE:", signature)
PYTHON
```

## Example with a JSON body

| Item                   | Rule / example                                                                                          |
| ---------------------- | ------------------------------------------------------------------------------------------------------- |
| Shared fixture         | `put-observed-k8s-cluster` in [sigv1-vectors.json](sigv1-vectors.json)                                  |
| Method                 | `PUT`                                                                                                   |
| Path                   | `/internal/v1/agent/resources/01J8ZK9XCLUSTER1/observed`                                                |
| Content type           | `application/json`; included before `host` in sorted canonical headers.                                 |
| Required report fields | `kind`, `observedAt`, `state`; successful convergence may acknowledge the current `observedGeneration`. |
| Expected signature     | `24fc2c36e16ec4df9a15b45fa23f7f9e67418e4f501884fc193cdebfb5570a81`                                      |
| Body bytes             | Use the exact `body` string from the fixture; pretty-printing it changes the signature.                 |
| Heartbeat              | Lifecycle 1.1.0 requires `supportedKinds`; advertise only tested installed drivers.                     |

## Server verification and replay protection

| Order | Server action                                                         | Failure behavior                                                            |
| ----- | --------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| 1     | Validate headers/date, active service credential and body bound.      | Reject malformed, expired/revoked or oversized requests.                    |
| 2     | Recompute canonical request and compare MAC in constant time.         | Reject tampering; do not reserve the nonce for an invalid MAC.              |
| 3     | Atomically reserve the authenticated nonce in Valkey for 600 seconds. | Reject replay; nonce-store outage must fail closed.                         |
| 4     | Attach `kan.principal` containing accountId/accessKeyId/kind.         | Controllers/services still enforce account, agent and function permissions. |
| 5     | Execute the authorized handler.                                       | A valid signature alone never grants cross-account access.                  |

| Additional rule     | Value                                                                                         |
| ------------------- | --------------------------------------------------------------------------------------------- |
| Maximum signed body | 1 MiB, including chunked bodies.                                                              |
| Go nonce generation | Request helper uses 32 cryptographically random bytes encoded as hex.                         |
| Retry               | New date, nonce and signature every attempt.                                                  |
| Public APIs         | `/v2/` and `/kan/v1/` use separate Bearer authentication; C must implement real token lookup. |

## Checks before merging

| Check              | Command / expected result                                                                                                                     |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------- |
| Shared signatures  | `python3 contracts/sign_example.py` — all three vectors PASS.                                                                                 |
| Lifecycle contract | `python3 contracts/check_lifecycle.py` — generated files and fixtures PASS.                                                                   |
| Go                 | From `kan-agent/`: `go test ./...`.                                                                                                           |
| Java               | From `kan-server/`: `mvn verify`.                                                                                                             |
| Live integration   | Valid request succeeds; replay/tampering fail; bad MAC cannot poison a legitimate nonce; enforce permissions and fail closed on store outage. |
| Remaining work     | Passing offline vectors does not prove real credential lookup, authorization or deployment is complete.                                       |
