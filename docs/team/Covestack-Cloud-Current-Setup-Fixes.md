# Covestack Cloud — Current Setup Fixes and Implementation Checklist

Updated 11 October 2026. This checklist covers **all 38 findings** from the historical [audit](Covestack-Cloud-Audit-2026-10-11.md). Security/build changes are applied and committed. Rows describe remediation/acceptance scope, not an assertion that every fix remains missing or is fully validated.

The security/build baseline is committed as `349a2f4` (Fixed security issues)
on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`.
The working tree was clean when inspected. Remote main/PR merge status has not been
verified; review and merge through the shared PR. Three signature vectors, negative
tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates
and auth-smoke exist. The config.go file currently contains only `package config`;
it is a placeholder, not a working loader. Real credential/token lookup, application
main, migrations/controllers, full agent and live deployment tests remain pending.

Final submission: **Tuesday, 15 December 2026**. Use [the dated workflow](Covestack-Cloud-Project-Workflow.md) and [corrected setup guide](Covestack-Cloud-Complete-Working-Setup.md).

## 1. Continue from the committed baseline

Boot 3.5.16, Java 25, Flyway PostgreSQL support and Go 1.27.2 toolchain/CI are configured.
Run normal builds in CI/on Fedora with dependency access; earlier local Go checks
used 1.27.1 and Maven downloads were blocked. Filter registration, signed nonce/key,
1 MiB body bound, principal/service checks and Valkey reservation are implemented;
real credential lookup, permissions and live integration acceptance remain pending.

Compose now has file secrets, a persistent Valkey volume and a backend network.
Finish local ignored image references/secrets, validate actual service readiness,
then add MinIO/workload network isolation when implementing those features. Preserve
existing volumes and coordinate initialized Postgres password rotation.

Anuj starts config/client/polling; B adds main/migrations/controllers; C implements
credential/token stores and authorization. Do not expose private backend ports.

## 2. Safe baseline commands

From Claude, enter the actual repository:

```bash
cd ../Project-CDAC/kan-cloud
git status --short
java -version
javac -version
mvn -version
go version
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
```

The last command validates the **current** infrastructure syntax, not its security or app readiness. Preserve current local edits. No data-reset command is necessary.

Current signing baseline:

```bash
(cd kan-agent && go test ./...)
(cd kan-server && mvn -B test)
python3 contracts/sign_example.py
```

Run these against the committed baseline, plus live negative/integration acceptance. Maven dependency/cache errors are recorded as failures/blocked verification, never converted into success. The three vector signatures have already been regenerated; later protocol changes must update them together.

Once deploy/.env and local secrets plus corrected Compose exist, use:

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
```

The server/agent commands become usable only after their main/config files exist. Confirm actual readiness, not just containers being listed.

## 3. Source-file fixes: authentication and builds

All paths below are relative to kan-cloud. The numbered audit IDs match the unchanged audit report.

| Audit ID | File/location to change | Problem → required solution | Owner / target | Acceptance |
|---|---|---|---|---|
| 1 | contracts/sigv1-rules.json; Java KanSigV1AuthFilter/KanSigV1; Go sigv1/sign.go; vectors/Python | Unsigned nonce permits modified replay → require/sign nonce and key ID, collect them in Java and regenerate signatures | A+C / Oct 16 | Change only nonce/key ID without resigning: rejected |
| 2 | KanSigV1AuthFilter.doFilterInternal; production NonceStore | Nonces stored before authentication → verify first, then atomic authenticated SET NX EX; invalid ingress separately limited | C / Oct 16 | Bad signature does not consume nonce; concurrent identical valid requests execute once |
| 3 | CachedBodyHttpServletRequest constructor; ingress settings | readAllBytes unbounded → enforce 1 MiB declared and actual-byte cap, including chunked input, with 413 | C / Oct 16 | Oversized known-key bad-MAC/chunked input rejected without full allocation |
| 4 | New Spring app/filter registration, Bearer verifier and route policy | Public/Kan routes skipped and no wiring → register filters and explicitly cover all routes except minimal liveness | B+C / Oct 16 | Real HTTP routing rejects unsigned/invalid requests on protected paths |
| 5 | KeyService credential model, filter principal and repositories | Secret-only lookup lacks identity/kind → active account/key/kind/permissions record, trusted principal and scoped reads/updates | B+C / Oct 16–23 | User credential denied on agent route; cross-account IDs denied; principal exists |
| 6 | New ValkeyNonceStore; InMemoryNonceStore; profiles | Memory store ignores TTL/restarts → shared atomic store; bounded expiring test-only fallback | C / Oct 16 | Replay rejected across instances/restart policy; store outage fails closed |
| 7 | Java/Go canonicalization and rules validation | Missing required signed headers silently omitted → reject missing/duplicate/ambiguous input and define normalization | A+C / Oct 16 | Missing host/content type/nonce fails; raw-path/query edge vectors agree |
| 8 | contracts/sign_example.py; Java tests; vectors | Python signs all supplied headers → shared selection policy; load real vectors and UTF-8 bytes | A+C / Oct 16 | Python/Java/Go compare same complete header inputs |
| 9 | README, Java application/config/migrations, kan-agent/cmd/kan-agent/main.go | Startup commands target nonexistent apps → implement entry points; fix Compose path and runnable Go command | A+B+E / Oct 16–23 | Clean checkout runs documented app and real agent |
| 10 | kan-server/pom.xml; wrapper/CI/runtime images/versions.env | Boot 3.3.5 outside Java 25 support → reviewed maintained compatible parent + one tested matrix | B+C / Oct 16 | Java 25 build/package/boot passes; dependency scan triaged |
| 11 | kan-server/pom.xml and migrations | Postgres Flyway plugin absent → add aligned flyway-database-postgresql; mail starter before email | B+E / Oct 16 | Actual Postgres migration succeeds from empty database |
| 12 | .github/workflows/ | Only Gitleaks exists → Java verify, Go test/vet, contract validation and integration checks | A+C / Oct 16–23 | Real required CI jobs execute and fail on regressions |
| 13 | secrets-scan workflow; versions/install/image manifests | Unverified archive/moving tags → pinned action commits/tools/images and checksum verification | A / Oct 16, release Dec 4 | Tampered download fails; release artifacts/digests reproducible |
| 14 | deploy/docker-compose.yml; local secrets/networks; VM firewall | Default secrets/implicit networking → separate secrets/config, explicit backend/storage and workload-account policies | A+C / Oct 16–23 | Backend inaccessible from client/workload IP; no default passwords |
| 15 | Valkey Compose volume/config and readiness/recovery | AOF without managed persistence policy → named /data volume; authenticated durable store and fail-closed loss handling | A+C / Oct 16 | Store replacement/restart preserves valid-window protection or blocks until recovery |

## 4. Source-file fixes: contracts

| Audit ID | File/location to change | Problem → required solution | Owner / target | Acceptance |
|---|---|---|---|---|
| 16 | contracts/internal-api.yaml ObservedReport; examples; Java/Go DTOs | Missing timestamp/nonnullable generation → required observedAt, nullable int64 generation, report ordering and applied response | A+B / Oct 16 | Null health reports/schema examples pass; stale report ignored |
| 17 | contracts/FROZEN-STATES.md; YAML/examples; all fakes/drivers | Conflicting kinds/states → corrected per-kind vocabulary from design section 5; add DB stopped state | A+B+D / Oct 16 | Every accepted state/example valid for its kind |
| 18 | YAML kind/spec enum and additional agent paths | Planned features outside contract → reviewed compute.snapshot/network.load-balancer kinds and ssh-keys/lb-members APIs; exclude groups | A+B+D / Oct 16 core, Nov 13 extensions | Unsupported kinds reject; extensions validated before drivers merge |
| 19 | OpenAPI security/headers/specs/errors and known-ID protocol | Auth only prose, unconstrained values → required signing header/security declarations, bounded typed specs, complete scoped snapshots | A+B+C / Oct 16–23 | Examples/negative payloads validated; partial known-ID response never deletes objects |

Do not label corrected contracts frozen until Java/Go/schema tests and the required reviewers accept them. Protocol additions for snapshots/LB/k8s must be versioned, not ad hoc strings in weekly tasks.

## 5. Source/config fixes: behavior, networking and data

| Audit ID | File/location to create/change | Problem → required solution | Owner / target | Acceptance |
|---|---|---|---|---|
| 20 | Agent reconcile.handle; Java observed/action/retry service | Failure acknowledged as success → preserve successful generation, record attempt/error and schedule bounded retry; action errors truthfully | A+B / Oct 23–30 | Failed create/stop/delete never completes successfully |
| 21 | Observed service/repository; schema/DTO epoch/sequence | Old generation/timestamp can corrupt state → scoped locked validation, monotonic success, generation/sequence/epoch and timestamp bounds | B+A / Oct 23–30 | Newer intent survives old-generation/newer-time reports |
| 22 | Work spec, OperationService, agent restart driver | active→active cannot express reboot → explicit restart intent and executed marker | A+B / Oct 30 | Process/container start actually changes; retry does not blindly repeat |
| 23 | AsgScheduler/group repo; periodic agent inspection/lease | Only active/recent count creates excess replacements → pending reservations, parent binding, deadlines, serialized scaling, scale-down and outage pause | A+B / Nov 6 | Slow pull/agent outage stays within desired/max; failed children retired |
| 24 | Managed DB driver/workload network policy | User DB joins control-plane network → account workload network, narrow allowed client routes | A+C / Nov 13 | User DB cannot reach backend IP/ports |
| 25 | SSH image build, key injection and host command runners | Shared root password/shell interpolation → key-only SSH, safe file write/argument arrays, bind/firewall policy | A+C+D / Oct 30 | Real SSH works; malicious key/name cannot execute shell data |
| 26 | Image catalog, Dockerfile and HostConfig capability profile | Plain Ubuntu/no tested sshd capability setup → pinned SSH-ready image with reviewed minimum and real session tests | A / Oct 30 | Key login and session work under exact release capabilities |
| 27 | Managed DB volume lifecycle; MinIO Compose | Recreated DB loses attached-data identity; MinIO mounts wrong path → labeled named DB volumes and miniodata:/data, restore/deletion policy | A+B / Nov 13 | Tables/objects survive replacement; backup restore passes |
| 28 | SpacesService, token/policy issuer, S3 client mapping | Prefixes/root creds do not isolate S3 → per-account credential policies and scoped physical-bucket metadata | C+E+B / Nov 13 | A credential cannot read/write/list B bucket |
| 29 | VM NAT/DNS/firewall/proxy configs and client metadata | Forwarding reaches guest interface, loopback-only services inaccessible → correct guest listeners/proxy, DNS TCP+UDP, laptop-address answer | A+E / Nov 20 | Two-device DNS/TLS and user service endpoints pass |
| 30 | Nginx/host policy and private agent URL | Signed Host rewritten; internal publicly proxied; Mailpit HTTP treated as HTTPS → authority preservation, internal denial and explicit restricted TLS/tunnel | A+C / Nov 20 | Public internal request blocked; private signed Host test succeeds |
| 31 | LB config generator/persistent directory/membership loop; web workload image | Invalid top-level fragment/stale member updates → valid conf.d include, nginx -t, safe directory replacement and health/IP refresh | A+B+D / Nov 20 | LB boot/reload and member removal/empty pool behavior pass |
| 32 | Droplet action API, SnapshotService/image mapping/docs | Invented snapshot endpoint/full-disk claim → type snapshot action, scoped public image IDs and documented volume exclusion | A+B+D / Nov 20 | godo snapshot action works; volumes not claimed as captured |
| 33 | Idempotency/operation/credential/event migrations and services | Concurrent token race/undefined columns/weak audit → atomic fingerprinted claim, consistent lock/revocation model and append-only identity/outcome audit | B+C / Oct 30 | Concurrent create yields one resource; altered-token payload rejects; audit immutable to app |
| 34 | CLI bootstrap/config/kubeconfig export; auth matrix | Unrestricted credentials/overwritten kubeconfig/auth mismatch → local bootstrap, restrictive ACL/mode, atomic writes, Bearer CLI and separate/merged kubeconfig | C+E / Oct 23, kubeconfig Nov 20 | Credential mode passes; original kubeconfig preserved |
| 35 | SDK tests, hardening/health scripts, isolation evidence | Ignored errors/timeouts/always-zero script → checked errors and bounded assertions, nonzero failed checks, accurate limited-compliance claims | A+C+D / Oct 23–Dec 4 | Deliberate bad state/failed control makes test/script fail |
| 36 | Work predicate, periodic inspect, agent lease, orphan cleanup and GC | Converged crashes missed/terminal rows churn/partial list dangerous → retry/health protocol, fencing, complete scoped snapshot, deletion acknowledgement/deleted_at | A+B / Nov 6 | Crash heals, stopped resources stay stopped, partial responses never destroy live data |
| 37 | All Claude guides plus repo README/contracts/runtime config | Conflicting algorithm/scope/layout/versions → one corrected design/workflow and actual executable contract | All / docs done; repo Oct 16–23 | No active contradictory recipes; docs match implemented entry points/features |
| 38 | Workflow/calendar | Incorrect weekdays → corrected 2026 dates and explicit submission buffer | All / documentation corrected | Calendar dates verified; Dec 15 is Tuesday |

### Report/failure acceptance sketch

This is an algorithm requirement, not drop-in production code:

```text
lock account+agent-scoped resource
validate assignment/lease epoch, report sequence, timestamp and generation
if obsolete report: return applied=false
if health-only: update current inspection, do not acknowledge target generation
if Ensure failed: record attempt/error, error matching action, schedule allowed retry
if target actually reached: advance successful observed generation monotonically
complete action only if its expected result and target generation are confirmed
commit; emit notifications through post-commit outbox
```

The agent retains failed dirty inspections until successfully reported. Work ordering and full periodic scans must be tested with Docker events disconnected. No design assumes a healthy process is observed forever from its first successful report.

## 6. OS and network settings to change on the VM

These are proposed administrative changes; this task did not alter the VM.

| Setting | Current/design issue | Required target |
|---|---|---|
| Guest IP | Forwarding documentation assumed auto/blank IP | Reserved actual NAT guest IP and recorded interface |
| Public API | Forwarding raw app port exposes private paths | Nginx HTTPS 443; deny internal path; private signed app route |
| DNS | Loopback-only guest listener; UDP-only forwarding | Guest-interface DNS listener; narrow TCP/UDP 53 forwarding/firewall |
| Backend DB/cache | tools suggested direct Postgres forward | No hotspot forwarding; loopback admin or SSH tunnel |
| Mailpit UI | HTTP port described as HTTPS | Private tunnel or explicit restricted TLS proxy |
| S3 | Root credentials/direct HTTP across devices | Scoped credentials + suitable TLS endpoint; private admin console |
| Workload ports | Dynamic VM ports not reachable from other devices | Small tested per-port NAT/firewall budget and client-reachable metadata |
| SSH | Shared image root password | Uploaded validated public keys, password auth disabled |
| Docker authority | Host execution described as inherently safer | Trusted agent daemon access; maintain VM/host boundaries |
| SELinux/firewall | Broad trusted-zone workaround | Leave enabled, diagnose/allow narrowly required paths |
| Certificates | Permanent trust/renewal assumptions | DNS-01 provision before rehearsal, renewal/reload and tested fallback |

Before changing storage or passwords, inventory existing data/volumes and back it up. Do not use destructive reset/prune to avoid solving a configuration bug. Store master-key backup separately and test restore; losing it makes encrypted signing/DB credentials unusable.

## 7. Workload and compatibility budget

For the demonstration, cap total workload counts against measured VM capacity. Suggested starting budget: two ASG Droplets, one independent Droplet, one managed Postgres, one LB and one small k3d cluster. Run Kubernetes separately if simultaneous load exceeds available memory. Budget ports before configuring per-port forwards. Admission quotas prevent requests from exhausting memory/PIDs/disk.

Use one tested image/size catalog, one PostgreSQL version, one HTTP LB rule and one Kubernetes version/node shape. Add snapshots through the correct action API. Advanced features require their own acceptance; no success-shaped placeholder is allowed.

## 8. Fix batches and dates

| Batch | Deadline | Audit IDs | Merge evidence |
|---|---|---|---|
| Security/build/contracts | Oct 16 | 1–8, 10–17, 19, core 37 | Real registered filter + negative vectors + migration/build CI |
| End-to-end core | Oct 23 | 9, 18 core, 20–21, 34 core, 35 tests | Real created container and truthful observed/action result |
| Full compute | Oct 30 | 22, 25–26, 33 | Actual reboot/key SSH, retry/idempotency tests |
| Health/scaling | Nov 6 | 23, 36 | Slow-pull/outage/lease/GC tests |
| Persistent data | Nov 13 | 24, 27–28 | Recreated DB/MinIO retains data; scoped S3 denied across accounts |
| Advanced/demo connectivity | Nov 20 | 18 extensions, 29–32, 34 kubeconfig | Actual second device plus bounded LB/k8s/snapshot |
| Release verification | Dec 4–8 | All applicable, final 13/35/37 | Scans, clean install, full supported demo and honest limitations |
| Calendar | Already corrected | 38 | Workflow dates verified |

If a batch fails, the next work slot fixes it before expansion. The deadline cannot make an incomplete security control safe; reduce feature depth and retain required controls. Explicitly track blockers rather than marking docs as implementation.

## 9. Evidence to collect before submission

For each audit ID store owner, PR/commit, acceptance command, sanitized result and status. Use Not started / In progress / Implemented awaiting validation / Accepted. Documentation corrections cover problem descriptions, not proof that code behavior changed.

Collect CI artifacts, negative auth results, cross-account denial, real SDK lifecycle, port rebuild/agent outage, persistence/restore, network reachability by IP, second-device TLS and release checksums. Keep screenshots free of secrets. A limited container-hardening script is labeled by its checked controls, not full CIS compliance.

[Documentation change log](Covestack-Cloud-Documentation-Change-Log.md) gives before/after lines and problems/solutions for every changed guide section. [Workflow](Covestack-Cloud-Project-Workflow.md) puts submission assembly on Dec 14 so Dec 15 is reserved for delivery.


## Implementation progress — 11 October 2026

The security/build baseline is committed as `349a2f4` (Fixed security issues)
on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`.
The working tree was clean when inspected. Remote main/PR merge status has not been
verified; review and merge through the shared PR. Three signature vectors, negative
tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates
and auth-smoke exist. The config.go file currently contains only `package config`;
it is a placeholder, not a working loader. Real credential/token lookup, application
main, migrations/controllers, full agent and live deployment tests remain pending.

See [code changes](Covestack-Cloud-Code-Changes.md) and the
[handoff](../ANUJ-PERSON-C-HANDOFF.md). The original security patch
has already been applied; do not apply it again. The separate Covestack branding
patch updates the published baseline and includes these team guides.
