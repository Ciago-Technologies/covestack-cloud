# Covestack Cloud — Team Roles and Acceptance Ownership

| Item | Details |
|---|---|
| 1 | Updated 11 October 2026. Baseline 349a2f4 is published on fix/security-baseline. Five-person plan, submission **Tuesday, 15 December 2026**. Use [workflow](Covestack-Cloud-Project-Workflow.md) for dates and [design](Covestack-Cloud-Technical-Design.md) for technical rules. This file assigns work; it does not restate conflicting signing algorithms or introduce different states. |

## 1. Team responsibilities

| Person | Primary responsibility | Actual folders |
|---|---|---|
| Anuj — A | Go agent, Docker drivers, infra, CI, security review | kan-agent/, deploy/, .github/, contracts/ |
| B | Java core, migrations, repositories, agent API, actions/ASG | kan-server/src/.../repository, service, agent; resources/db/migration |
| C | Auth, keys, authorization, audit/security tests | kan-server/src/.../auth/security/audit; signing code |
| D | DO endpoints/DTOs, golden files and SDK compatibility | kan-server/src/.../do; compat/ |
| E | Kan management API, CLI, doctor, Mailpit integration | kan-server/src/.../kan/mail; new kan-cli/ |

| Item | Details |
|---|---|
| 1 | Anuj designs/reviews security with C; ownership does not imply one person can fix every issue alone. B owns backend correctness; E's CLI and D's provider adapters call B's scoped services. No business logic is duplicated in controllers/CLI. |

## 2. Anuj: agent, infrastructure and review

| Item | Details |
|---|---|
| 1 | Current priority: review the committed baseline with C, implement config/client/main/polling and keep integration infrastructure working. Toolchain pins, signer and vectors exist; verify actual CI and runtime acceptance. |
| 2 | Deliverables: |
| 3 | Registered/signed private agent communication; account/agent lease/fencing and reserved labels. |
| 4 | Sharded workers, in-flight deduplication, non-blocking dispatch and dirty-event retry on the same worker. |
| 5 | Startup port rebuild from running/stopped Docker inspect, explicit bind addresses and safe allocation failure handling. |
| 6 | Periodic inspection of converged resources, pull/readiness timeouts and crash vs intentional-stop handling. |
| 7 | Key-only SSH-ready images with tested minimum capabilities; safe argument arrays/file writes. |
| 8 | Persistent managed DB volumes; correct MinIO /data; workload/account isolation and measured IP-level checks. |
| 9 | Valid LB includes/config directory/reload, bounded k3d integration and scoped snapshot images. |
| 10 | Actual Java/Go/contract/integration CI, pinned/verified tools/images and release checksums. |
| 11 | Acceptance: real create/delete, agent restart with stopped port owners, failed Ensure not falsely acknowledged, second-agent fencing, backend unreachable by IP, saved data surviving recreation and valid negative signature tests. |
| 12 | Do not implement every optional driver before the first real end-to-end flow. Anuj and C pair on security; Anuj and B pair on reconciliation failure semantics. |

## 3. B: core state, SQL and agent API

| Item | Details |
|---|---|
| 1 | First priorities: Spring entry point/config, real migrations and scoped repositories. Boot/Flyway dependency changes are committed; verify the full build. Existing schema prose is not an applied schema. |
| 2 | Deliverables: |
| 3 | Transactional resource/action creation; numeric external IDs; explicit account/agent predicates. |
| 4 | Work/heartbeat/known-IDs/events API matching reviewed contract; groups never agent work. |
| 5 | Observations with timestamp/generation/epoch/sequence checks and null health-only reports. |
| 6 | Successful-generation acknowledgement only on confirmed convergence; failed operation attempt/new linked retry on failure. |
| 7 | Explicit reboot intent and safe completion; operation locking derived from open rows. |
| 8 | Atomic scoped idempotency claim with body fingerprint and stored response. |
| 9 | ASG counts pending reservations, handles scale-down/retirement, pauses during agent outage and binds parent/account. |
| 10 | DB, Spaces metadata, LB, cluster and snapshot services behind provider/CLI interfaces. |
| 11 | Deleted_at/GC ordering and data retention/restore contracts. |
| 12 | Acceptance: stale reports cannot corrupt newer intent; failed actions cannot become completed; concurrent idempotent creates yield one result; slow pulls do not inflate ASG; fresh migration and restored data both work. |
| 13 | FakeAgent is explicit test/dev profile only and uses the same service/action path. It reports kind-correct states and does not run alongside the real agent on the same resources. Never remove a protective profile annotation to make a release work. |

## 4. C: authentication and authorization

| Item | Details |
|---|---|
| 1 | Filter registration, body bounds, signed identity, verify-before-nonce and replay plumbing are committed. Next implement real credential/token stores, expiry/revocation and permissions, then prove these controls on the running server. |
| 2 | Deliverables: |
| 3 | Route matrix: sigv1 service keys for internal API; opaque Bearer tokens for /v2 and /kan management; minimal public health exception. |
| 4 | Simple HMAC protocol from contracts; no four-step derivation or AWS Authorization branch. |
| 5 | Required/duplicate header validation, raw-byte body signing, constant-time comparison and negative vectors. |
| 6 | Atomic authenticated replay protection, bounded authenticated/unauthenticated rate limits, restart durability policy. |
| 7 | AES-GCM secrets with protected external master key, rotation and restricted credential files. |
| 8 | Account/agent/function permission checks in every service/endpoint; hashed Bearer token lookup and revocation. |
| 9 | Structured audit identity/outcome and append-only DB permissions; no decrypted secret/URI/kubeconfig logging. |
| 10 | Scoped S3 policies, security scan triage, k3d privilege boundary and network reachability review. |
| 11 | Acceptance: modified nonce rejected, original replay rejected, bad MAC cannot poison nonce, missing/oversized chunked body handled, user key cannot call agent API, cross-account IDs denied and store outage fails closed. Use vetted constant-time primitives; no fragile 5% request-timing assertion. |

## 5. D: DigitalOcean adapters and real clients

| Item | Details |
|---|---|
| 1 | First priorities: contract states/examples/schema agreement and numeric ID/response inventory. Use actual DO reference/godo types rather than invented endpoint shapes. |
| 2 | Deliverables: |
| 3 | Account, Droplet CRUD, actions/read polling, catalogs, tags and validated SSH key references. |
| 4 | Kind-specific DTOs/status mapping, bounded pagination and error codes; no internal ULID leaks. |
| 5 | Managed Postgres API with correct connection/status and deliberate credential visibility. |
| 6 | HTTP LB and bounded DOKS subset after core gates. |
| 7 | Snapshot through POST /v2/droplets/{id}/actions with type snapshot; public image ID mapping. |
| 8 | Golden examples and schema checks plus godo/doctl smoke tests on the actual signed/Bearer integration stack. |
| 9 | Explicit documented 501 for unsupported methods/features; no success-shaped stubs. |
| 10 | Acceptance: SDK errors checked, bounded eventual assertions prove active and action completion, failed operations have truthful error outcomes and every mutation validates resource ownership. |

## 6. E: management API, CLI and mail

| Item | Details |
|---|---|
| 1 | First priorities: local administrative bootstrap, restricted credential files, minimal public liveness, authenticated detailed doctor and standalone CLI entry point. |
| 2 | Deliverables: |
| 3 | Account-scoped key/token management with server-side permissions; no public unauthenticated key issuer. |
| 4 | CLI configure/env/instances/actions/asg/db/spaces/lb/k8s and supported-feature help. |
| 5 | Bearer credentials for management/provider APIs; sigv1 service secrets reserved for agent requests. |
| 6 | Correct Compose path/env-file handling; local up/down without deleting volumes. |
| 7 | Remote doctor reads server-side status; no assumption the remote laptop owns the Docker socket or Postgres. |
| 8 | Scoped S3 upload/download with physical bucket mapping; password-masked list output. |
| 9 | Kubeconfig exported separately or safely merged with backup/permissions. |
| 10 | Post-commit mail outbox/retries; Mailpit UI via restricted proxy or tunnel. |
| 11 | Tested packaging/install instructions and release checksums; foreground agent remains supported. |
| 12 | Acceptance: fresh configuration works without token bypass; secret files/ACLs are restrictive; existing kubeconfig survives; SMTP outage does not break a committed create; remote doctor reports actual remote readiness. |

## 7. Pairing and reviews

| Boundary | Owners | Required review/evidence |
|---|---|---|
| Contract/signing | A + B + C | Java/Go vectors, negative tests, schema/examples |
| First real Droplet | A + B + D | API → DB → Docker → observed → action |
| Bootstrap/CLI credentials | C + E | Local issuance, permissions, revocation |
| ASG/retries | A + B | Slow pull, outage, failure/retirement |
| DB/Spaces | A + B + C + E | Persistence, credentials, cross-account denial |
| Public DNS/TLS | A + E | Actual second-device and private-route tests |

| Item | Details |
|---|---|
| 1 | Use reviewed branches, small changes and meaningful tests for behavior/security changes. After 90 minutes blocked, pair; after one working day, reassign help and revise gate status. All owners can explain the code they accept. |

## 8. Schedule and completion evidence

| Item | Details |
|---|---|
| 1 | The workflow is the single dated plan. G1–G3 finish secure compute by Oct 30; G4–G5 finish ASG/data by Nov 13; G6 targets bounded advanced paths Nov 20; half demo Nov 28; feature freeze Dec 1; release candidate Dec 8; submit Dec 15. |
| 2 | Map every accepted remedy to an audit ID in [setup fixes](Covestack-Cloud-Current-Setup-Fixes.md). Evidence records actual commit, command, result and limitations. Documentation corrections alone never close a code defect. |
