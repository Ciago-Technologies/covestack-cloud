# CoveStack V5 — Anuj’s Agent, Infrastructure and ITISS Work

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. What you already have

You are Person A. The Go config loader, signer, signed HTTP work client, cancellable polling loop and cmd/kan-agent entry point already exist. They contact a configured server but do not establish trusted provisioning on their own. When resource work arrives the loop deliberately stops: no claim, execution, acknowledgement or deletion. Do not start by recreating client.go/main/loop.

## 2. Your immediate discussion with B+C (D reviews API)

| Question to settle | Why it matters | Actual output |
| --- | --- | --- |
| Which real origin/HTTPS path is running? | Contract localhost example is not proof of a listener | Nonsecret confirmed endpoint and startup evidence |
| How is a test service credential created/delivered/revoked? | Signing must use real stored identity, not default/fake key | Protected delivery and lifecycle agreement; no secret in issue/docs |
| How does server bind account/host/agent assignment? | Current principal has account/key/kind only; caller IDs/labels cannot grant authority | Reviewed stored assignment and negative cases |
| What is no-work/error behavior? | Empty success differs from auth/server/transport failure | Actual GET /internal/v1/agent/work 200 [] plus observed errors |
| Which missing protocol changes are prerequisites? | Current 1.1.0 is not full VM/type/claim/capacity support | B+D+A/C design/fixture issue, not invented wire fields |

## 3. First coding/integration slice

1. Explain config.Load → client.New → agent.Run under signal context in the existing main.
2. On an authorized dev setup run relevant current tests/static checks and attach exact results.
3. Once B+C real server/identity exists, prove signed no-work, invalid/revoked credentials, replay behavior with C, Ctrl+C and server/outage observations.
4. Review bounded transport-error policy: current loop retries timeout-class/selected HTTP, not every connection error. Change only the agreed missing behavior and matching tests.
5. Keep no-execution safe stop until assignment, claims, type compatibility, runtime ownership and acceptance are approved.

[Existing-code change plan](../EXISTING-CODE-CHANGE-PLAN.md) maps exact present files; it does not invent missing driver/controller names. [Starter](../Starter.md) and [team checkpoint](../TEAM-FIRST-CHECKPOINT.md) define first issues/reviews.

## 4. Your full dependency order

| Order | Your responsibility | Dependency / real evidence |
| --- | --- | --- |
| 1 | Integrate existing single-account signed client/loop | B+C real auth/server/assignment; C acceptance |
| 2 | Review future shared client/module and host/account contexts | A+B+C+E; keep CLI user auth separate; no signer move assumed |
| 3 | Establish trusted execution/type/claim boundary with server | B+D schemas/SQL, C policy, current-generation/epoch and unknown-task design |
| 4 | Authorized Incus/VM host spike and capability inspection | Real guest kernel/image/provenance, disks/NIC/default deny and measured overhead |
| 5 | Owned runtime inspect/reconcile and provider-task recovery | Durable desired intent/claims/reservations; crash/lost reply one VM, no unrelated cleanup |
| 6 | Separate trusted App Container and k3d execution | Explicit type/trust, real readiness, measured ports/quotas and cleanup |
| 7 | Network/storage/data integration | C IPAM/B records, owned attach/traffic; PostgreSQL real SQL/restore; operator RustFS process/volume while Java owns buckets/keys/policies |
| 8 | Bounded environment execution and independent verified recovery | Stable child bindings/readiness/retention/shared-dependency policy with B/D/C/E |
| 9 | CI/installer/hardening/operations evidence | Clean repeat install, artifact/profile/secret policy, actual failure/upgrade/restore and ITISS explanation |

Do not add a second account context before first integration is understood; later prove account failure/credential isolation under host-wide coordination. Do not wait until the end to add meaningful failure tests: C owns acceptance throughout.

## 5. Files and folder boundaries

| Existing path | Purpose | What you do now |
| --- | --- | --- |
| kan-agent/internal/config/ | Validated environment/rules and secret-safe formatting | Retain validation; reviewed future loading/custody changes only |
| kan-agent/internal/client/ | Signed bounded read-only work request | Retain safeguards; real-server integration, future reviewed APIs later |
| kan-agent/internal/agent/ | Cancellable poll/backoff/auth/no-execution behavior | Test actual error/cancel behavior; no runtime command inserted now |
| kan-agent/cmd/kan-agent/ | Startup, signals, config/client/loop wiring | Explain and integrate existing main |
| kan-agent/cmd/auth-smoke/ | Separate signing diagnostic | Keep mock evidence distinct from real Java/Valkey |
| kan-agent/sigv1/ | Canonical signing/request/rules | Preserve cross-language compatibility; no speculative crypto rewrite |
| kan-agent/internal/lifecycle/ | Generated data and pure validation | Do not hand-edit generated catalogue or treat it as tenant/lease authorization |
| pkg/kanclient, drivers and other proposed modules | Planned layout only | Decide names/interfaces with reviewers before creating/moving code |

## 6. What you do not take over

B owns all SQL/server durable coordination, C all acceptance/security/Java allocator, D Java DTO/API and no Go, E thin CLI/packaging/later provider. You review E's Go, not implement the CLI for them. Java reconciles RustFS buckets/policies/tenant keys; your process/volume/host work must not become a second bucket controller. Advanced proxy/publication/health/scaling/snapshot extensions follow actual dependencies and do not starve safe current work.

## 7. ITISS evidence and code explanation

Explain why each signed request is fresh, why account/host assignment is trusted, why DB fencing is not runtime cancellation, why a VM has a guest kernel, why packets are actually denied and why a file backup is not a verified restore. Keep host/socket/image/key/traffic/capacity/incident/restore records with exact revision and sanitized results. AI-generated Go is reference only: type, understand and explain contributed code. No mentor no-coding policy or approval is invented; ask if the course requires it.

## 8. Daily next action

One issue, one bounded change, real dependency owner and assertion. Read actual files/contract; retain current guards, add meaningful tests, attach actual scope/results, open PR and update checkpoint. No service support from a diagram, source file or mock. No code/host execution authorized by this document alone.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-anuj-next-steps) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
