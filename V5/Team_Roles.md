# CoveStack V5 — Ownership, Reviews and ITISS Responsibilities

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Retained five-person ownership

| Person | Current responsibility | First reviewable slice | Later work / boundary |
| --- | --- | --- | --- |
| A — Anuj | All agent/shared client/signing/driver/reconcile Go, CI, installer and ITISS evidence | Explain existing config/client/loop, agree credentials/assignment with B+C, prove real no-work polling | VM/container/k3d/network/storage execution; never assume a new driver exists; reviews E's Go without taking over CLI |
| B | Java core and ALL SQL/migrations, operations, leases/fences/idempotency, quotas/reservations and environments | Reviewed runnable server/foundation/migration prerequisites with real auth wiring | Java RustFS orchestration, persistence/recovery, future ASG coordinator; no JPA |
| C | Authentication/security, ALL acceptance tests, Java IPAM allocator | Concrete credential/assignment/replay design and real-server negative tests | Scoped storage policies, packet isolation, domain/metadata policy and verified recovery |
| D | Java APIs/DTOs, bounded /v2 mappings, schema/API review and provider fixtures | Inventory actual endpoint gaps; review current discovery and future VM/claim fixtures | Further provider mappings and compatibility; writes no Go |
| E | Thin Go CLI/cobra, JSON/errors/wait, packaging; later thin provider | Reviewed module/client plan and mock HTTP learning slice | Real user API integration, optional Terraform/OpenTofu adapter and eventual read-only UI after ownership/capacity needs are reviewed; no agent/signing/drivers/reconcile |

These responsibilities do not impose academic service limits or deadlines. A future Go scheduler/extra platform service needs explicit ownership, not an assumed transfer from B's current coordination.

## 2. A's immediate sequence

1. Read existing kan-agent/cmd/kan-agent/main.go, internal/config, internal/client, internal/agent/loop.go and sigv1; explain current no-execution behavior.
2. Agree protected test credential provisioning, account/host assignment and real server/no-work response with B+C and D compatibility review.
3. Run existing relevant unit/static checks in an authorized dev environment and retain actual evidence.
4. Integrate the existing signed loop against the real Java endpoint: [] response, invalid/revoked credentials, cancellation and observed server-outage/retry behavior.
5. Coordinate future VM/type/claim/network/capacity schemas before execution code; separately approve Incus host evaluation.
6. Build owned inspection/reconciliation only once protocol, capacity and runtime safety are reviewable; prove actual guest/service behavior.

Do not recreate the existing loop or advertise driver capabilities before implementation. AI-generated agent code is reference material: A types, understands and explains contributed Go. See [Anuj sequence](CoveStack-Anuj-Next-Steps.md).

## 3. B's Java / SQL boundary

B establishes actual startup and applied migration history before changing dependencies or migration order. V5/V6 publication/domain SQL in source reference missing foundation tables; their presence is not a clean migration chain. B owns relational ownership, allocation/port uniqueness, quotas/reservations, exact-generation fenced reports, idempotency transactions and retained deletion. Never edit an applied migration; current reference schema is not an executable VM migration.

## 4. C's acceptance and security boundary

C starts with real auth/assignment/replay tests, not only late demonstrations. API tokens use hash-only lookup; signing credentials use protected recoverable secrets, kind service and reviewed custody. Caller account strings/labels are not trusted ownership. C owns actual foreign packet/disk/bucket/key/restore denial and freshness/unknown-outcome tests; implementers pair and supply real evidence.

## 5. D's contract boundary

D checks base schemas/DTOs/fixtures against actual 1.1.0 behavior. New VM versus App Container, capabilities, claims and network attachment are coordinated changes; no invented current fields/version. Generated outputs follow authoritative inputs. Public API compatibility is a tested subset, not full cloud-provider parity.

## 6. E's Go boundary

CLI authenticates as the user and talks only to the API, holds config rather than local resource DB, uses idempotency on mutations and JSON on list/get. E does not implement privileged execution or signing logic. Shared pkg/kanclient/module layout is Planned and reviewed by A+B+C/E; do not move modules independently. [Go handover](CoveStack-Go-Handover.md).

## 7. Shared responsibility by domain

| Domain | Design / implementation | Acceptance |
| --- | --- | --- |
| VM/App Container / k3d | A execution; B persistence/admission; D spec/API; C policy | C with A real backend evidence |
| Network/IPAM | C allocator; B transactions; A backend policy/attachments; D API | Actual allow/deny and parallel allocation |
| PostgreSQL / volumes | B coordinator/records; A workload/disks; C secrets; D API | Real SQL persistence/known-row restore and foreign denial |
| RustFS | B Java buckets/policies/keys; A operator process/volume; C policy/custody; D API | Actual bytes, scoped access, restart and isolated verified restore |
| Environments / quotas | B durable child/reservation/coordination; A executor; D DTO; C security | Partial retry no duplicates, no overspend/shared deletion |
| UI / inventory / optional providers | API-only consumer with reviewed owner; E bounded client work | No authorization bypass or direct backend access |

## 8. Pair reviews and bottleneck control

| Change | Review |
| --- | --- |
| Any Go | A |
| Shared client/signing / VM / k3d execution boundary | A+B+C; D when wire/DTO changes |
| SQL/schema/reservations | B+C+A; D DTO alignment |
| Auth/security/acceptance | C plus affected owner |
| Java API mapping | B/C/D as affected; A agent-facing |

A is a bottleneck across runtime, CI, infrastructure and ITISS. B owns SQL, C tests/allocator, D mappings and E CLI to keep this bounded. Do not start speculative extensions at the expense of current reviews/security/recovery. Track dependencies on issues, not guessed percentages.

## 9. ITISS and communication

ITISS evidence is host/virtualization/network policy, identity/replay, secret custody, capacity admission, hardening, incident diagnosis and verified recovery—not just number of container commands. Ask the mentor in writing about course-specific coding/ownership expectations if needed; no confirmation is assumed and it does not constrain product architecture. [Security/ITISS](CoveStack-Security-and-ITISS.md).

Use Starter.md for onboarding, TEAM-FIRST-CHECKPOINT.md for shared work/board, GitHub issues for live status and root checkpoint.md for AI continuity. No new deadline or role transfer is declared here.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-team-roles) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
