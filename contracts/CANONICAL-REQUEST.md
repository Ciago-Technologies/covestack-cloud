# What is a "canonical request" and how kan sigv1 signs one

## Honesty first
Your own docs say kan sigv1 shares its HMAC-SHA-256 verifier with AWS SigV4,
and gets tested against AWS's own SigV4 worked example. They do NOT spell out
the exact canonical-request string format anywhere I've seen. So the design
below is **me proposing a format that fits what your docs already say** —
it's a real, working algorithm you can implement today, not a quote from
your TDD. Get Person C (sigcore owner) to sign off on it before it's frozen.

## What "canonical request" means, in easy words
Before you sign anything, both sides (the one signing, the one checking)
must build the EXACT same string of text from the HTTP request. If Java
builds the string slightly differently than Go — extra space, different
header order, different casing — the signatures will never match, even
with the same secret key.

"Canonical" just means: one agreed, boring, predictable way to turn a
request into text. No creativity allowed.

## The recipe (6 lines, joined with newlines)

```
CanonicalRequest =
  HTTP_METHOD          (e.g. "GET")
  + "\n" +
  PATH                 (e.g. "/internal/v1/agent/work")
  + "\n" +
  QUERY_STRING         (empty string if none)
  + "\n" +
  CANONICAL_HEADERS    (see below)
  + "\n" +
  SIGNED_HEADERS       (see below)
  + "\n" +
  SHA256_HEX(BODY)     (hash of the raw request body; hash of "" if no body)
```

**Canonical headers**: pick `host` and `x-kan-date` always, plus
`content-type` if the request has a body. Lowercase the header name,
trim the value, sort alphabetically by name, one per line as
`name:value\n`.

**Signed headers**: the same header names, lowercased, joined with `;`,
in the same sorted order. This tells the verifier which headers were
actually included in the signature.

## Then: string to sign

```
StringToSign =
  "KAN1-HMAC-SHA256"
  + "\n" +
  X_KAN_DATE            (same value as the header, e.g. "20260926T120000Z")
  + "\n" +
  SHA256_HEX(CanonicalRequest)
```

## Then: the actual signature

```
Signature = HMAC_SHA256_HEX( secretKey, StringToSign )
```

That hex string is what goes in the `Authorization` (or `X-Kan-Signature`)
header. The server re-derives the same canonical request from the incoming
HTTP request, re-runs this exact recipe with the account's known secret
key, and checks the two signatures match. If they don't, reject with 401.

## Worked example (real, computed — not made up)

Request: `GET /internal/v1/agent/work`, no query, no body.

Headers used for signing:
```
Host: localhost:4700
X-Kan-Date: 20260926T120000Z
```

Step 1 — Canonical Request:
```
GET
/internal/v1/agent/work

host:localhost:4700
x-kan-date:20260926T120000Z

host;x-kan-date
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```
(that last line is `SHA256("")` — the hash of an empty body, always the
same value, worth memorizing so you recognize it in tests)

Step 2 — String To Sign:
```
KAN1-HMAC-SHA256
20260926T120000Z
a20b0876060aa38062365c07d59eecde6d327dd9a2c71b30fa452721e58530c3
```
(that last line is `SHA256(CanonicalRequest above)`)

Step 3 — Signature (secretKey = `test-secret-do-not-use-in-prod`):
```
4f934e9d22c52f8a04b7158964dec9f952c77d7d081cbd6fc3a4a3c8db39c86a
```

A second worked example, this time with a body (`PUT /observed`), is in
`sigv1-vectors.json` — same recipe, just with `content-type` added to the
signed headers and the real body hash used.

## How to check your own implementation
Run `sign_example.py` in this folder (no dependencies, just Python 3).
It prints the same three steps above. If your Java or Go implementation
produces the identical Signature hex string for the identical inputs,
your implementation is correct — for THIS proposed algorithm. Confirming
this is the algorithm your team actually wants is still Person C's call.
