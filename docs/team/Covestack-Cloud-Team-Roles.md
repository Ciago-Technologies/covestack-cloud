# Covestack Cloud — Team Roles and Acceptance Ownership

Updated 11 October 2026. Baseline 349a2f4 is published on fix/security-baseline. Five-person plan, submission **Tuesday, 15 December 2026**. Use [workflow](Covestack-Cloud-Project-Workflow.md) for dates and [design](Covestack-Cloud-Technical-Design.md) for technical rules. This file assigns work; it does not restate conflicting signing algorithms or introduce different states.

## 1. Team responsibilities

| Person | Primary responsibility | Actual folders |
|---|---|---|
| Anuj — A | Go agent, Docker drivers, infra, CI, security review | kan-agent/, deploy/, .github/, contracts/ |
| B | Java core, migrations, repositories, agent API, actions/ASG | kan-server/src/.../repository, service, agent; resources/db/migration |
| C | Auth, keys, authorization, audit/security tests | kan-server/src/.../auth/security/audit; signing code |
| D | DO endpoints/DTOs, golden files and SDK compatibility | kan-server/src/.../do; compat/ |
| E | Kan management API, CLI, doctor, Mailpit integration | kan-server/src/.../kan/mail; new kan-cli/ |

Anuj designs/reviews security with C; ownership does not imply one person can fix every issue alone. B owns backend correctness; E's CLI and D's provider adapters call B's scoped services. No business logic is duplicated in controllers/CLI.

## 2. Anuj: agent, infrastructure and review

Current priority: review the committed baseline with C, implement config/client/main/polling and keep integration infrastructure working. Toolchain pins, signer and vectors exist; verify actual CI and runtime acceptance.

Deliverables:

- Registered/signed private agent communication; account/agent lease/fencing and reserved labels.
- Sharded workers, in-flight deduplication, non-blocking dispatch and dirty-event retry on the same worker.
- Startup port rebuild from running/stopped Docker inspect, explicit bind addresses and safe allocation failure handling.
- Periodic inspection of converged resources, pull/readiness timeouts and crash vs intentional-stop handling.
- Key-only SSH-ready images with tested minimum capabilities; safe argument arrays/file writes.
- Persistent managed DB volumes; correct MinIO /data; workload/account isolation and measured IP-level checks.
- Valid LB includes/config directory/reload, bounded k3d integration and scoped snapshot images.
- Actual Java/Go/contract/integration CI, pinned/verified tools/images and release checksums.

Acceptance: real create/delete, agent restart with stopped port owners, failed Ensure not falsely acknowledged, second-agent fencing, backend unreachable by IP, saved data surviving recreation and valid negative signature tests.

Do not implement every optional driver before the first real end-to-end flow. Anuj and C pair on security; Anuj and B pair on reconciliation failure semantics.

## 3. B: core state, SQL and agent API

First priorities: Spring entry point/config, real migrations and scoped repositories. Boot/Flyway dependency changes are committed; verify the full build. Existing schema prose is not an applied schema.

Deliverables:

- Transactional resource/action creation; numeric external IDs; explicit account/agent predicates.
- Work/heartbeat/known-IDs/events API matching reviewed contract; groups never agent work.
- Observations with timestamp/generation/epoch/sequence checks and null health-only reports.
- Successful-generation acknowledgement only on confirmed convergence; errored action/retry on failure.
- Explicit reboot intent and safe completion; operation locking derived from open rows.
- Atomic scoped idempotency claim with body fingerprint and stored response.
- ASG counts pending reservations, handles scale-down/retirement, pauses during agent outage and binds parent/account.
- DB, Spaces metadata, LB, cluster and snapshot services behind provider/CLI interfaces.
- Deleted_at/GC ordering and data retention/restore contracts.

Acceptance: stale reports cannot corrupt newer intent; failed actions cannot become completed; concurrent idempotent creates yield one result; slow pulls do not inflate ASG; fresh migration and restored data both work.

FakeAgent is explicit test/dev profile only and uses the same service/action path. It reports kind-correct states and does not run alongside the real agent on the same resources. Never remove a protective profile annotation to make a release work.

## 4. C: authentication and authorization

Filter registration, body bounds, signed identity, verify-before-nonce and replay plumbing are committed. Next implement real credential/token stores, expiry/revocation and permissions, then prove these controls on the running server.

Deliverables:

- Route matrix: sigv1 service keys for internal API; opaque Bearer tokens for /v2 and /kan management; minimal public health exception.
- Simple HMAC protocol from contracts; no four-step derivation or AWS Authorization branch.
- Required/duplicate header validation, raw-byte body signing, constant-time comparison and negative vectors.
- Atomic authenticated replay protection, bounded authenticated/unauthenticated rate limits, restart durability policy.
- AES-GCM secrets with protected external master key, rotation and restricted credential files.
- Account/agent/function permission checks in every service/endpoint; hashed Bearer token lookup and revocation.
- Structured audit identity/outcome and append-only DB permissions; no decrypted secret/URI/kubeconfig logging.
- Scoped S3 policies, security scan triage, k3d privilege boundary and network reachability review.

Acceptance: modified nonce rejected, original replay rejected, bad MAC cannot poison nonce, missing/oversized chunked body handled, user key cannot call agent API, cross-account IDs denied and store outage fails closed. Use vetted constant-time primitives; no fragile 5% request-timing assertion.

## 5. D: DigitalOcean adapters and real clients

First priorities: contract states/examples/schema agreement and numeric ID/response inventory. Use actual DO reference/godo types rather than invented endpoint shapes.

Deliverables:

- Account, Droplet CRUD, actions/read polling, catalogs, tags and validated SSH key references.
- Kind-specific DTOs/status mapping, bounded pagination and error codes; no internal ULID leaks.
- Managed Postgres API with correct connection/status and deliberate credential visibility.
- HTTP LB and bounded DOKS subset after core gates.
- Snapshot through POST /v2/droplets/{id}/actions with type snapshot; public image ID mapping.
- Golden examples and schema checks plus godo/doctl smoke tests on the actual signed/Bearer integration stack.
- Explicit documented 501 for unsupported methods/features; no success-shaped stubs.

Acceptance: SDK errors checked, bounded eventual assertions prove active and action completion, failed operations have truthful error outcomes and every mutation validates resource ownership.

## 6. E: management API, CLI and mail

First priorities: local administrative bootstrap, restricted credential files, minimal public liveness, authenticated detailed doctor and standalone CLI entry point.

Deliverables:

- Account-scoped key/token management with server-side permissions; no public unauthenticated key issuer.
- CLI configure/env/instances/actions/asg/db/spaces/lb/k8s and supported-feature help.
- Bearer credentials for management/provider APIs; sigv1 service secrets reserved for agent requests.
- Correct Compose path/env-file handling; local up/down without deleting volumes.
- Remote doctor reads server-side status; no assumption the remote laptop owns the Docker socket or Postgres.
- Scoped S3 upload/download with physical bucket mapping; password-masked list output.
- Kubeconfig exported separately or safely merged with backup/permissions.
- Post-commit mail outbox/retries; Mailpit UI via restricted proxy or tunnel.
- Tested packaging/install instructions and release checksums; foreground agent remains supported.

Acceptance: fresh configuration works without token bypass; secret files/ACLs are restrictive; existing kubeconfig survives; SMTP outage does not break a committed create; remote doctor reports actual remote readiness.

## 7. Pairing and reviews

| Boundary | Owners | Required review/evidence |
|---|---|---|
| Contract/signing | A + B + C | Java/Go vectors, negative tests, schema/examples |
| First real Droplet | A + B + D | API → DB → Docker → observed → action |
| Bootstrap/CLI credentials | C + E | Local issuance, permissions, revocation |
| ASG/retries | A + B | Slow pull, outage, failure/retirement |
| DB/Spaces | A + B + C + E | Persistence, credentials, cross-account denial |
| Public DNS/TLS | A + E | Actual second-device and private-route tests |

Use reviewed branches, small changes and meaningful tests for behavior/security changes. After 90 minutes blocked, pair; after one working day, reassign help and revise gate status. All owners can explain the code they accept.

## 8. Schedule and completion evidence

The workflow is the single dated plan. G1–G3 finish secure compute by Oct 30; G4–G5 finish ASG/data by Nov 13; G6 targets bounded advanced paths Nov 20; half demo Nov 28; feature freeze Dec 1; release candidate Dec 8; submit Dec 15.

Map every accepted remedy to an audit ID in [setup fixes](Covestack-Cloud-Current-Setup-Fixes.md). Evidence records actual commit, command, result and limitations. Documentation corrections alone never close a code defect.
