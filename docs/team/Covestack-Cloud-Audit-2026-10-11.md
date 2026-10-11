# Covestack Cloud audit — 11 October 2026

**Historical pre-fix audit.** Evidence, identifiers and original paths below describe
the earlier inspected state. Baseline 349a2f4 now includes security/build fixes; use
the current code/status guide to assess remaining work. Product branding is now
Covestack Cloud. Findings are not automatically closed by a passing unit test.

Reviewed all 6 existing Markdown files in `Claude/` (7,330 lines), then all 8 Markdown files in `../Project-CDAC/kan-cloud/` (4 are empty placeholders), and every implementation, test, contract, example, dependency manifest, Compose file and CI workflow present in that repository. No project source files were changed. Existing uncommitted changes were preserved.

This is an early skeleton, not a running emulator. Severity describes the consequence if the existing classes or documented starters are used as intended. Design findings below are not claims of exploits in code that has not been implemented. No deployed server, Docker runtime, tenant services or image vulnerability scanner was exercised.

## Verification results

- Go: `go test ./...` passed in a temporary copy with an isolated build cache and network downloads disabled.
- Java: all existing auth source files compiled with Java 25 against locally cached dependencies. Both shared signature vectors passed in a standalone probe. This is not a successful Maven build.
- Maven: `mvn -o -B test` was blocked by missing cached Maven plugin dependencies. No conclusion about the full Maven test/build result follows from that failure.
- Mock-servlet probes of the actual Java filter reproduced: original signature + nonce A → 200; identical signature + nonce B → 200; bad signature + nonce C → 401; subsequent valid signature + nonce C → 401; unsigned `/kan/v1/keys` forwarded; valid internal request forwarded without `kan.principal`.
- Actual cached-body wrapper accepted 2,097,152 bytes, exceeding the documented 1 MiB limit.
- Python reference signer with all headers from shared vector 1 did not match that vector; manually selecting its signed headers did match.
- These probes exercised classes directly. There is no application entry point or registered filter configuration to test over a live HTTP server.

## Existing implementation and deployment findings

### 1. High — Replay nonce is not signed

Evidence: `contracts/sigv1-rules.json:7` signs only `host` and `x-kan-date`. `KanSigV1AuthFilter.java:95` forwards only host, date and optional content type to verification. Both canonicalizers omit the nonce. A captured request can be replayed during its timestamp-validity window by changing `X-Kan-Nonce`, without knowing the secret. This was reproduced using the actual filter.

Fix: require and sign `x-kan-nonce` and `x-kan-access-key` in Java and Go; include them in the verifier's canonical header map; regenerate vectors. Validate nonce format and length. Changing only the JSON rules is insufficient because the Java filter does not currently collect these headers.

Acceptance: identical replay fails; changing nonce without resigning fails; a newly signed request with a fresh nonce succeeds.

### 2. High — Unauthenticated requests consume nonce storage

Evidence: `KanSigV1AuthFilter.java:79` calls `checkAndStore` before key lookup and signature verification. An attacker can submit current dates, arbitrary key IDs and unique nonces to grow the store. A bad signature also poisons a nonce before a legitimate request uses it; reproduced.

Fix: validate bounded headers and body, look up the active key, verify the MAC, then atomically claim the nonce with `SET NX EX` before invoking downstream code. Apply a separate ingress/IP limit to invalid requests. Fail closed when the replay store is unavailable.

### 3. High — Cached body has no size limit

Evidence: `CachedBodyHttpServletRequest.java:28` uses `readAllBytes()`. A request with a known access-key ID and plausible headers reaches this allocation before its signature is checked. This does not require a valid signature. A 2 MiB payload was accepted.

Fix: check declared length and read at most `MAX_BYTES + 1`, reject over-limit input with 413, including chunked input and missing/false length. Add read timeouts and a consistent ingress limit. Do not rely exclusively on `Content-Length`.

### 4. High — Authentication coverage is incomplete and not wired

Evidence: `KanSigV1AuthFilter.java:51` skips all routes outside `/internal/v1/agent/`. No Bearer verifier, filter bean/registration or Spring application exists. Unsigned `/kan/v1/keys` is forwarded by this class; `/v2/*` is also outside its scope.

Fix: add explicit filter registration and route policy covering internal, DO and Kan APIs. Implement Bearer verification; explicitly allow only intended public health endpoints. Integration-test unauthenticated access and path matching with the actual servlet routing configuration.

Do not interpret this as evidence that an already deployed public API exists: no controllers are present.

### 5. High — Authenticated identity and service-key authorization are missing

Evidence: `KeyService.java:10` returns only a secret. `KanSigV1AuthFilter.java:109` forwards without attaching account identity or checking key kind. The probe confirmed absence of `kan.principal`.

Fix: resolve an active credential record containing account ID, key ID, kind and permissions; require service credentials for agent endpoints; attach the trusted principal. Derive tenant/account scope from it, never from user-controlled labels. Require account/agent ownership in every read and update, including events, known IDs, SSH-key fetches and LB members.

### 6. High when used beyond isolated tests — In-memory replay store never expires

Evidence: `InMemoryNonceStore.java:15` retains entries indefinitely and ignores `ttlSeconds`; restart clears replay history, and separate server instances have separate maps. Its own comment correctly warns against production use, but the workflow calls it fine for demos.

Fix: use an atomic shared Valkey implementation. Keep a bounded, expiring in-memory implementation exclusively in a test/dev profile. Define persistence/restart behavior so still-valid signed requests cannot be replayed after cache loss.

### 7. Medium — Required signing headers are silently omitted

Evidence: Java `KanSigV1.java` filters signing names with `lower::containsKey`; Go `sigv1/sign.go` includes them only if found. Non-empty bodies can be signed without content type; missing host is not rejected by the filter's required-header check.

Fix: canonicalization must fail on a missing required header rather than weaken the signature. Specify raw path/query handling, percent encoding, duplicate headers/parameters, whitespace, Unicode and UTF-8. Reject ambiguous duplicate signing headers and test Java/Go parity.

### 8. Medium — Python referee uses different header selection

Evidence: `contracts/sign_example.py:21` signs every supplied header; Java and Go use the selected rules. Passing vector 1's full header dictionary to Python produces a different signature because it includes the access-key header, which Java/Go currently omit.

Fix: use one shared header-selection policy in all three implementations. Run the Python implementation against the actual vector file rather than hand-selected examples. Use raw UTF-8 bytes consistently; Java's test currently uses default-charset `getBytes()`.

### 9. High integration blocker — Application and agent cannot run as documented

Evidence: no `@SpringBootApplication`, controllers, configuration, migrations, concrete KeyService, production NonceStore, Go `main`, Docker driver, poller or CLI exists. `README.md:21` invokes Compose from the root, but the file lives in `deploy/`. Its `go run .` has no runnable package.

Fix: correct README to `docker compose -f deploy/docker-compose.yml up -d`; implement a minimal authenticated vertical slice with entry points before claiming an operational stack. Clearly mark planned features and placeholders. The current Compose starts only Postgres, Valkey and Mailpit, not the server or MinIO.

### 10. Medium — Spring Boot and Java support mismatch

Evidence: `kan-server/pom.xml:10` selects Boot 3.3.5 and line 20 targets Java 25. Official Boot 3.3 documentation supports Java through 23. Local Java compilation succeeding does not establish supported runtime/tooling compatibility.

Fix: use a maintained Boot release that explicitly supports Java 25, or choose a supported Java target across build/runtime/CI. A Boot 3.5 release explicitly supports Java 25, but select a maintained patch version at implementation time and verify migration compatibility. The old dependency stack needs a resolved-dependency scan; this audit does not assert a reachable CVE merely from its age.

Sources: [Boot 3.3 requirements](https://docs.spring.io/spring-boot/3.3/system-requirements.html), [Boot 3.5 requirements](https://docs.spring.io/spring-boot/3.5/system-requirements.html).

### 11. High future startup blocker — Flyway PostgreSQL module is missing

Evidence: `kan-server/pom.xml:36` includes only `flyway-core`, not the PostgreSQL database plugin.

Fix: add `org.flywaydb:flyway-database-postgresql` with a version aligned with the selected Boot dependency management, then test migrations against the actual Postgres version. Add the mail starter when implementing `JavaMailSender`.

Source: [Flyway PostgreSQL dependency documentation](https://github.com/flyway/flyway/blob/main/documentation/Reference/Database%20Driver%20Reference/PostgreSQL%20Database.md).

### 12. Medium — CI does not test the project

Evidence: the only workflow, `.github/workflows/secrets-scan.yml`, runs Gitleaks. There is no Java/Go build, unit test, contract validation or integration job, despite docs claiming these checks run.

Fix: add Java verify and Go test/vet jobs with correct module working directories, followed by contract-example validation and an authenticated vertical-slice integration test. Required branch checks should correspond to actual jobs.

### 13. Medium — Supply-chain artifacts are not verified or immutable

Evidence: CI downloads and executes a Gitleaks archive without checksum verification; checkout uses a mutable version tag. Compose uses mutable image tags and `mailpit:latest`. Several installation snippets execute scripts from moving branches or select `latest` binaries.

Fix: pin action commits and tool versions, verify published checksums/signatures, use `curl --fail`, set least-privilege workflow permissions and pin image digests with a planned update mechanism. Digest pinning does not replace vulnerability scanning.

### 14. Medium — Default credentials and backend isolation remain development-only

Evidence: Compose sets `POSTGRES_PASSWORD: kan-dev-only`, Valkey has no authentication, and services share the default Compose network. All published infrastructure ports bind to loopback, which is a useful restriction; the file itself does not expose them to the LAN.

Fix: distinguish dev and demo/deployment configurations, supply non-default credentials, and define backend/workload networks explicitly. Do not forward the control-plane Postgres port to other demo devices as `Covestack-Cloud-Tools.md:293` suggests. Enforce the intended workload-to-host rules and test reachability by IP, not only DNS name.

### 15. Medium — Replay-store persistence is not explicitly managed

Evidence: Valkey enables AOF but Compose declares no named Valkey volume. Do not assume AOF survives replacement of its storage; anonymous-image-volume behavior is not a reproducible persistence policy.

Fix: define a named volume for the actual persistence directory and test container replacement. Retain nonces through the maximum timestamp validity window, use `noeviction`, monitor capacity, and reject signed requests when storage fails.

## Contract defects

### 16. High — ObservedReport cannot express the documented ordering protocol

Evidence: `contracts/internal-api.yaml:170` requires integer `observedGeneration`, omits `observedAt`, and does not mark generation nullable. The TDD requires a timestamp and permits null generation for crash/state-only reports. `examples/put-observed-request-k8s.json` also has no timestamp. Responses specify `{ok:true}`, whereas the design needs an explicit applied/ignored result.

Fix: require `observedAt` with date-time format, make `observedGeneration` optional/nullable with int64 and nonnegative bounds, document stale-report behavior, and align examples and response DTOs. Validate examples against the schema in CI.

### 17. High — Frozen state names contradict their own examples

Evidence: `FROZEN-STATES.md:11` defines Kubernetes observed states as `provisioning/running/...`, but lines 16–19 and YAML examples use `creating/active`. DB desired state includes `stopped`, but the observed-state list lacks `stopped`. Workflow Week 6 uses kind `database` and desired `active/archive`, whereas the contract uses `db.instance` and `online/stopped/deleted`.

Fix: define one kind-specific state machine and update YAML, prose, examples, fakes and adapters together. Include a database stopped state if stopping is supported. Keep DO status mapping distinct from internal state vocabulary.

### 18. High — Expanded resource kinds and endpoints are outside the contract

Evidence: YAML kind enum at line 139 supports only 4 kinds. Workflow adds `snapshot`, `load_balancer`, `/ssh-keys` and `/lb-members`, none of which is specified. `compute.group` is included in the work-item enum even though the agent is told to ignore it.

Fix: settle release scope, extend/version the contract before implementation, define per-kind spec schemas and transitions, and exclude server-only groups from agent work.

### 19. Medium — OpenAPI omits authentication and meaningful validation

Evidence: YAML prose says calls are signed, but it has no security scheme/requirement or required signing-header parameters. `spec`, `detail`, state and desiredState are mostly unconstrained. Generation has no int64 format; work and known-IDs have no completeness/pagination contract. Heartbeat is 200 in YAML but 204 in the workflow.

Fix: describe required signature headers and rejection responses; use kind-discriminated schemas and numeric/string limits; unify heartbeat status. Make known-IDs explicitly complete or paginated with a completion marker. Never delete orphans on a failed, truncated or partially read response.

## Bugs and vulnerabilities in documented starters and plans

### 20. High — Failed reconciliations are acknowledged as successful generations

Evidence: `Covestack-Cloud-Technical-Design.md:1090` creates an error report when Ensure fails, then unconditionally sets `ObservedGeneration = &gen`. The completion SQL marks operations complete solely from generation. This can remove failed work from the queue and report a failed action as completed.

Fix: distinguish attempted and successfully converged generations. On failure preserve the successful acknowledgement, record retry/backoff and mark the corresponding action errored; complete only when the expected state and generation have actually been reached. Test failed create, stop and delete.

### 21. High — Generation can regress or old reports can overwrite state

Evidence: workflow Week 3 and team-role observed SQL use only `LEAST(:obsGen,generation)`, losing the TDD's monotonic/null protection. TDD SQL clamps generation but still overwrites state/detail from an older-generation report with a newer timestamp. An extreme future observedAt can suppress later updates.

Fix: preserve monotonic acknowledged generation, handle null explicitly, reject impossible future/negative generations and timestamp skew, and check report generation before updating reconciliation state/detail. Define equal-time ordering or an agent epoch/sequence. Include account/agent ownership in the update predicate.

### 22. High — Reboot intent is not represented

Evidence: action types include reboot/power_cycle, but work carries only desired state/spec/generation. Changing active→active and bumping generation does not tell Ensure to restart an already running container rather than acknowledge it as active.

Fix: include an operation/restart intent or explicit restart generation in the work spec and persist its completed marker. Test that a reboot action actually changes the container's start time/process instance and is applied once under retries.

### 23. High — Auto-scaling can create replacements indefinitely

Evidence: workflow `Covestack-Cloud-Project-Workflow.md:554` counts only active, recently observed instances. Work polling excludes converged resources, and periodic observations of them are not specified. Healthy resources age out after 60s. Pending creates also do not count, so slow startup triggers extra creates each scheduler pass. There is no specified scale-down or failed-child retirement. Some starter creates do not establish parent/group membership.

Fix: periodically inspect converged resources; distinguish agent outage from a confirmed workload failure; count pending reservations; lock/serialize decisions per group; enforce min/max and scale-down; associate new children atomically and retire failed children before replacement. Test image pulls longer than multiple scheduler periods and agent outages.

### 24. High — Managed workload databases share the control-plane backend

Evidence: workflow `Covestack-Cloud-Project-Workflow.md:634` starts user DB containers on `kan-backend`. This violates the intended separation from Postgres/Valkey/control-plane services.

Fix: use workload/account networks for user databases, allow only the intended client paths and keep backend services unreachable. A separate shared droplets network alone also does not isolate one tenant's Droplets from another tenant's Droplets.

### 25. High — Static root SSH password and shell interpolation

Evidence: TDD `Covestack-Cloud-Technical-Design.md:1459` sets the same root password in every image and enables root password login. Published SSH bindings omit HostIP. Workflow Week 6 writes a public key using an interpolated `sh -c` string. No safe parsing/escaping policy is specified.

Fix: use validated public keys and disable password authentication; choose explicit bind addresses and firewall rules for demo access. Write authorized_keys as file bytes through a safe Docker API rather than shell interpolation. Invoke k3d and other host commands with separate validated arguments, not shell strings; reserve ownership labels and prohibit user-selected host mounts/privilege flags.

### 26. Medium — SSH image and capability policy are incompatible without testing

Evidence: design flow pulls plain `ubuntu:22.04`, which has no running sshd; a different section builds an SSH image under a different name. `secureHostConfig` drops all caps except NET_BIND_SERVICE, whereas conventional root sshd/PAM sessions can require identity/permission-related capabilities.

Fix: map image slugs to prebuilt immutable SSH-enabled images. Test actual key login and shell startup under the exact capability policy; use a carefully justified minimum or a compatible non-root SSH setup rather than indiscriminately restoring all capabilities.

### 27. High data-loss risk — Database and MinIO volume plans are incomplete

Evidence: workflow DB creation specifies no managed named data volume or ownership/retention rules, so recreated databases cannot reliably reattach existing data. MinIO command stores data at `/data` while examples mount `miniodata:/var/lib/minio/data` (`Covestack-Cloud-Complete-Working-Setup.md:189`, workflow line 782).

Fix: attach named volumes to each database's actual data directory, label/scope them and define deletion/backup behavior. Mount MinIO storage at `/data`. Test delete/recreate of service containers while retaining volumes and verify inserted data/objects survive.

### 28. High — Direct S3 access has no tenant IAM design

Evidence: planned MinIO uses one root credential and direct S3 access. Account-prefixed bucket names prevent collisions but do not restrict what a client with root credentials can access. Metadata filtering cannot secure direct S3 requests.

Fix: issue account-scoped S3 credentials/policies, keep root credentials server-side, verify bucket ownership on mutations and test account A attempting account B's bucket by exact name.

### 29. High demo blocker — Loopback services cannot receive NAT-forwarded guest traffic

Evidence: TDD dnsmasq listens only at `127.0.0.1` (`Covestack-Cloud-Technical-Design.md:1277`), but VMware forwards to the VM's interface IP. Compose publishes services on guest loopback, while docs expect raw guest-IP forwarding for API/Mailpit/MinIO. DNS forwarding is UDP-only.

Fix: expose the intended guest interface through a controlled proxy/listener, keep backend ports private, and add firewall rules restricted to the demo subnet. Listen for DNS on the needed VM address and allow both TCP and UDP 53. Return the laptop's reachable hotspot address to external devices, not the hidden NAT guest address.

### 30. Medium — Nginx rewrites a signed Host and publicly proxies internal routes

Evidence: `Covestack-Cloud-Technical-Design.md:1399` uses `proxy_set_header Host $host`, which can strip a nondefault port from the authority the client signed. The same location proxies `/internal/*` publicly despite docs calling it isolated. HTTPS on port 8025 is described without a corresponding TLS listener; the existing Mailpit port is HTTP.

Fix: preserve/validate the exact agreed authority (`$http_host` where appropriate) and define proxy trust semantics. Deny internal paths on the public proxy and give the host agent a private authenticated route. Put Mailpit behind an explicit restricted TLS proxy location/listener.

### 31. Medium — New load-balancer design will not start Nginx as written

Evidence: workflow line 1024 generates only `upstream` and `server` blocks but mounts them as `/etc/nginx/nginx.conf`, which needs top-level `events` and `http` context. Updating an atomically replaced single bind-mounted file can leave the container reading the old inode. Membership triggers cover tag edits, not all create/delete/failure/IP changes. SSH images do not provide an HTTP server on port 80.

Fix: mount a valid include under `conf.d/` or generate a complete config, persist/mount a config directory, run `nginx -t` before reload, and reconcile actual healthy membership changes. Define a workload image/user-data setup with a working web service for the LB demo. Validate all generated identifiers and backend ports/addresses.

### 32. Medium — Snapshot compatibility and semantics are overstated

Evidence: workflow uses `POST /v2/droplets/{id}/snapshots`, but DigitalOcean snapshots are initiated through `/actions` with `type: snapshot`. It also calls docker commit a true disk-image equivalent.

Fix: implement snapshot as a DO action and preserve numeric/public image ID mapping. Document that container commit excludes mounted-volume data, does not capture VM memory, and requires separate consistency/backup handling. Avoid leaking source credentials and SSH host keys into cloned images; regenerate instance-specific identity.

Sources: [DO snapshot action](https://docs.digitalocean.com/products/snapshots/how-to/snapshot-droplets/), [Docker commit semantics](https://docs.docker.com/reference/cli/docker/container/commit/).

### 33. Medium — Idempotency and audit integrity are underspecified

Evidence: token lookup is described before creation without an atomic claim/body fingerprint; concurrent requests can create duplicates. Events schema lacks account/key/outcome and has no append-only permission enforcement. Workflow refers to nonexistent `resources.locked`, `resources.client_token` and `access_keys.revoked_at` columns.

Fix: atomically reserve `(account,scope,token)` in the resource/action transaction and store request fingerprint/result; reject token reuse with different payload. Derive locked state from open operations or migrate it consistently. Choose one revocation model. Add audit identity/outcome and DB permissions preventing application update/delete of audit rows.

### 34. Medium — Credentials and kubeconfig handling lack safe persistence rules

Evidence: CLI credentials are saved to `~/.kan/credentials` without a permissions/ACL requirement. Kubeconfig is redirected directly over `~/.kube/config`. Bootstrap credentials and `/kan/v1/*` authentication conflict: CLI is described as sigv1 while starter routing treats non-internal calls as Bearer.

Fix: define a one-time local bootstrap flow, separate Bearer tokens from public access-key identifiers, use restrictive file mode/Windows ACL and atomic writes, never print secrets in env/list output, and export to a separate kubeconfig or safely merge with backup. Specify auth for every CLI endpoint.

### 35. Medium — Test snippets can report false success

Evidence: godo loops in roles ignore GET errors and do not assert active was reached. Timing tests demand a 5% whole-request difference, which is unreliable under JIT/GC/scheduling. CIS script prints failures but ends with a successful echo, returning 0; checking 4 settings does not establish full CIS compliance. DNS failure by hostname alone can reflect missing DNS rather than blocked connectivity.

Fix: fail on SDK errors, use bounded eventual assertions and verify completed actions. Test cryptographic behavior, use the platform's constant-time primitive and avoid noisy CI timing thresholds. Track failed controls and return nonzero; label the script a limited hardening check. Test isolation by backend IP, host gateway and published ports as well as names.

### 36. Medium — Work/health/deletion protocol can stall, churn or delete live objects

Evidence: roles work query includes all `off/archive` rows even when converged; another query includes only unequal generations and misses a converged standalone crash. No periodic full reconcile or agent lease/epoch is specified. known-IDs cleanup lacks safety for partial results. TDD archive garbage collection conflicts with self-referencing parent rows and relies on updated_at despite repeated observation writes.

Fix: define a canonical work predicate with retry scheduling, full health scans, intentional-stop handling and terminal deletion acknowledgement. Add single-agent ownership/lease or explicitly restrict deployment to one agent. Clean orphans only from a complete successful scoped snapshot after a grace period. Track deleted_at independently and explicitly handle parent/child garbage collection.

### 37. Low–Medium — Documentation conflicts will produce incompatible implementations

Evidence: current source/TDD use simple HMAC, but setup Part 6, team-role Crypto and early workflow require a four-step derivation and AWS-style Authorization. Later workflow locks simple HMAC. Scope alternates between 501 for DOKS/DB and full DOKS/DB/LB/Spaces. README targets Java 21 while POM targets 25; tool docs mix Postgres 16/17 and Valkey 7/8. Compose build context `./control-plane/kan-server` is wrong relative to `deploy/` and the current layout. Claims that security classes are COMPLETE/do-not-modify obscure real defects.

Fix: designate versioned contracts as the executable source of truth, archive obsolete designs, document one actual layout/scope/auth policy and align setup commands with it. Existing simple HMAC is not inherently a vulnerability; contradictory recipes are the bug. Existing two vectors establish agreement only for their exact inputs, not comprehensive compatibility/security.

### 38. Low — Timeline weekdays are wrong

Evidence: workflow calls Oct 6, 2026 Monday, Oct 10 Friday, Nov 28 Friday and Dec 15 Monday. They are Tuesday, Saturday, Saturday and Tuesday respectively (verified with Python dates).

Fix: correct dates/weekday labels and milestones before relying on Friday gates. This is a planning defect, not a security issue.

## Recommended implementation order

1. Fix nonce signing, verification order, body limits, filter registration and principal/service-key checks together; add negative servlet tests.
2. Align ObservedReport, per-kind states, response codes, required signing headers and test vectors; validate every contract example.
3. Resolve Java/Boot/Flyway compatibility and build one authenticated create→agent→observed vertical slice with real entry points.
4. Correct generation/error/action/reboot and ASG logic before expanding feature scope.
5. Separate workloads from backend/other accounts; implement key-only SSH, persistent data volumes and scoped S3 credentials.
6. Add repeatable CI, checksummed tools/images, real cross-device networking tests and truthful security evidence.

Do not announce replay protection, a working emulator, complete DO compatibility or CIS compliance based solely on the current two happy-path signing vectors and design snippets.
