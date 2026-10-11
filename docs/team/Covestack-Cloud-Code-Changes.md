# Covestack Cloud: implemented code changes — 11 October 2026

## Current location and commit status

The security/build baseline is committed as `349a2f4` (Fixed security issues)
on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`.
The working tree was clean when inspected. Remote main/PR merge status has not been
verified; review and merge through the shared PR. Three signature vectors, negative
tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates
and auth-smoke exist. The config.go file currently contains only `package config`;
it is a placeholder, not a working loader. Real credential/token lookup, application
main, migrations/controllers, full agent and live deployment tests remain pending.

The original security patch was applied and committed by Anuj. Do not reapply
`Kan-Cloud-Code-Fixes.patch`. The table below records the original security patch's
38 file changes/line blocks, not every subsequent edit. Commit 349a2f4 also adds
`kan-agent/internal/config/config.go`, currently just `package config`.

The new branding work is prepared in `Claude/covestack-cloud` because this session
cannot edit the original checkout. See `Covestack-Cloud-Branding-Changes.md` for its
separate checked patch and updated document files.

## Changed files and exact lines

Line numbers refer to the resulting fixed copy. Each cell lists changed/new blocks;
removed-only blocks point to the resulting insertion location. The patch contains
exact before/after contents, including original line endings.

| File | Changed lines | Problem | Solution |
|---|---|---|---|
| `.github/workflows/build.yml` | 1–32 | No Go/Java/contracts build CI. | Add pinned setup actions, Go 1.27.2 race/vet, Java 25 Maven verify and Python vectors. |
| `.github/workflows/secrets-scan.yml` | 4–5, 10–10, 15–24 | Unchecked download and broad workflow permissions. | Fail on downloads, verify release checksum, pin checkout SHA, contents read-only. |
| `.gitignore` | 1–6 | Local credentials could be committed. | Ignore deploy secrets/.env and build caches. |
| `README.md` | 3–4, 6–6, 8–14, 17–17, 19–22, 24–24, 25–26, 29–33, 35–35, 37–43, 45–49 | Quickstart promised nonexistent runnable executables and wrong Compose path. | Provide truthful verification/setup commands and remaining app requirements. |
| `contracts/CANONICAL-REQUEST.md` | 1–1, 3–4, 6–6, 8–12, 14–19, 21–23, 25–25, 27–33, 35–38, 40–40, 42–45 | Documented unsigned nonce and outdated sample hashes. | Specify new signed fields, bytes/path/query rules, replay order and coordinated upgrade. |
| `contracts/FROZEN-STATES.md` | 10–10, 16–16, 18–19 | Lifecycle example contradicted its table and DB lacked stopped. | Use provisioning/running; include DB stopped. |
| `contracts/README.md` | 11–11, 22–23, 31–33 | README described obsolete vectors and confused desired/observed K8s states. | Describe current computed vectors and coordinated Anuj/C review. |
| `contracts/examples/put-observed-request-k8s.json` | 3–3, 9–10 | Example omitted timestamp and used the wrong observed state. | Add observedAt and running. |
| `contracts/internal-api.yaml` | 18–23, 90–90, 97–97, 125–125, 134–138, 155–156, 184–184, 188–191, 193–196, 201–202 | No declared auth and incomplete observed report semantics. | Declare all four required auth schemes, int64 generations, observedAt, nullable health generation, applied response. |
| `contracts/sign_example.py` | 2–2, 5–6, 8–9, 12–29, 31–38, 41–46 | Python selected different headers than Go/Java. | Load shared rules, reject missing/duplicate headers, check all vectors. |
| `contracts/sigv1-rules.json` | 7–12 | Nonce/key ID could be changed without invalidating the MAC. | Sign both fields mandatorily. |
| `contracts/sigv1-vectors.json` | 2–2, 12–13, 17–17, 28–29, 31–31, 33–50 | Old vectors encode the insecure signing format and wrong K8s state. | Regenerate all signatures; add UTF-8/raw-query vector. |
| `deploy/.env.example` | 1–4 | No way to record selected image versions. | Provide required local image reference variables; no invented untested digests. |
| `deploy/docker-compose.yml` | 3–3, 7–8, 11–11, 14–16, 18–19, 22–22, 26–26, 28–34, 36–37 | Inline password, volatile unauthenticated Valkey and floating image refs. | File secrets, persistent /data, noeviction, backend network and mandatory operator-selected image refs. |
| `deploy/valkey.conf.example` | 1–7 | No authenticated/durable Valkey configuration template. | AOF with always fsync, bounded noeviction and replaceable password. |
| `docs/ANUJ-PERSON-C-HANDOFF.md` | 1–107 | No actionable owner/interface/integration handoff. | Record breaking APIs, ownership, smoke procedure, remaining work and verification limits. |
| `kan-agent/cmd/auth-smoke/main.go` | 1–96 | No runnable handoff check against Java. | Add read-only five-case smoke command, timeout, HTTPS policy and no redirects. |
| `kan-agent/cmd/auth-smoke/main_test.go` | 1–62 | Smoke command needed repeatable acceptance coverage without live services. | Test five expected responses through an in-process transport. |
| `kan-agent/go.mod` | 4–5 | No pinned patched toolchain. | Pin go1.27.2; retain existing Go 1.22 language minimum. |
| `kan-agent/sigv1/negative_test.go` | 1–96 | No regression coverage for signed identity, missing fields or request byte mutation. | Add meaningful negative and immutable-body/fresh-nonce tests. |
| `kan-agent/sigv1/request.go` | 1–65 | No safe request constructor for the Go agent. | Add immutable body signing, random nonce, raw URL signing and body bound. |
| `kan-agent/sigv1/rules.go` | 5–5, 33–46 | Unsafe rule configuration could be loaded. | Reject missing security fields, wrong algorithm/names and too-short replay retention. |
| `kan-agent/sigv1/sign.go` | 12–12, 26–26, 38–42, 47–49, 51–51, 68–68, 79–79, 81–93, 100–100 | Missing required fields silently disappeared and duplicate names overwrote each other. | Return errors; reject missing/duplicate/CRLF fields and date mismatch. |
| `kan-agent/sigv1/sign_test.go` | 49–52 | Tests called the old signer API. | Handle errors while verifying all shared vectors. |
| `kan-server/pom.xml` | 10–10, 41–44, 68–69 | Boot 3.3.5 baseline mismatched Java 25 and PostgreSQL Flyway module was absent. | Select Java 25-compatible Boot 3.5.16, add Flyway PostgreSQL module; skip app repackage until main exists. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/AgentAuthConfiguration.java` | 1–51 | Filters were not registered and production nonce store not wired. | Register agent/public filters, load validated rules, wire Valkey and require credential services. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/BearerAuthFilter.java` | 1–35 | Public APIs had no authentication filter. | Require a single opaque Bearer token and active user account principal. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/BearerTokenService.java` | 1–8 | Public token lookup contract was absent. | Define active user principal lookup SPI; C implements hashing/expiry/revocation. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/CachedBodyHttpServletRequest.java` | 24–27, 32–34, 47–49 | Unbounded readAllBytes allowed memory exhaustion. | Check declared length and read at most MAX+1, reject oversized unknown-length bodies. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/InMemoryNonceStore.java` | 3–5, 7–7, 9–17, 19–27 | Entries never expired and capacity was unbounded. | Bound/expire test-only entries; reject saturation instead of evicting active nonces. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/KanSigV1.java` | 8–8, 11–12, 24–24, 31–33, 36–41, 61–66 | Verifier canonicalizer omitted required fields and used default-locale casing. | Strict canonical fields, Locale.ROOT, ASCII trim and date consistency. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/KanSigV1AuthFilter.java` | 5–5, 8–8, 12–13, 15–15, 17–19, 33–35, 39–56, 59–64, 67–71, 74–82, 85–92, 95–97, 100–110 | Invalid MACs consumed nonces; identity/header validation incomplete. | Check active service credential/MAC before atomic nonce reservation; attach principal; 413 bounds and fail-closed 503. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/KeyService.java` | 3–4, 6–10 | Secret-only interface could not carry account/key kind. | Define active Credential and non-secret Principal records; persistence remains C task. |
| `kan-server/src/main/java/com/ciagotech/kan/server/auth/ValkeyNonceStore.java` | 1–23 | No shared production replay store existed. | Atomic setIfAbsent with duration; key-safe encoding; no fallback on failure. |
| `kan-server/src/test/java/com/ciagotech/kan/server/auth/AgentAuthSecurityTest.java` | 1–112 | No filter failure-case tests. | Cover valid principal, replay/tamper/poisoning, body limits, wrong kinds, expiry and outage. |
| `kan-server/src/test/java/com/ciagotech/kan/server/auth/BearerAuthFilterTest.java` | 1–28 | No public auth rejection checks. | Test missing/unknown/duplicate tokens and user/service separation. |
| `kan-server/src/test/java/com/ciagotech/kan/server/auth/KanSigV1Test.java` | 36–36 | Body test encoding relied on platform default. | Use explicit UTF-8 for cross-language vectors. |
| `kan-server/src/test/java/com/ciagotech/kan/server/auth/ValkeyNonceStoreTest.java` | 1–37 | No check of atomic reservation/TTL API use. | Test conditional set with TTL, safe key encoding and null reply fail closed; real Valkey pending. |

38 files changed/added.

## Verification

- Python: all three shared vectors PASS.
- Go: race tests for signer/request helper and smoke command PASS; vet PASS.
  Ran with installed Go 1.27.1 and GOTOOLCHAIN=local; pinned Go 1.27.2 download blocked.
- Java: all main/test sources compiled with cached dependencies; 14 test methods
  PASS in a temporary direct runner. Full Maven verify blocked by inability to
  download Boot 3.5.16 from Maven Central (DNS/network restrictions).
- YAML: contracts, Compose and workflows parse; Compose config validates with
  temporary placeholder images/secrets. No infrastructure was deployed.
- Live Go->Java, real Valkey restart/multi-server/outage tests and GitHub CI execution
  remain pending. Action commit pins/version support verified against official sources.

## Remaining implementation — do not treat this as a completed emulator

C must implement active Postgres credential lookup/AES-GCM and hash-only Bearer token
lookup/revocation/expiry, plus permissions. Java core owner must implement the app
main/configuration/migrations and tenant-scoped controllers; enforce generation and
observation ordering. Anuj must build the full agent, polling/retries and Docker
resource reconciliation. Auth filters do not implement resource authorization.
Image refs remain deliberately operator-selected; validate and record exact VM
images. MinIO and workload isolation are planned separately. Keep the 38-finding
setup checklist for remaining application/infra issues.

## Official references checked

- [Boot3.5 Java 25 compatibility](https://docs.spring.io/spring-boot/3.5/system-requirements.html)
- [Go 1.27.2 security release](https://go.dev/doc/devel/release)
- [Setup Java pinned commit](https://github.com/actions/setup-java/commit/c5195efecf7bdfc987ee8bae7a71cb8b11521c00)
- [Setup Go pinned commit](https://github.com/actions/setup-go/commit/d35c59abb061a4a6fb18e82ac0862c26744d6ab5)
