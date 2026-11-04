# Feature specifications and implementation boundaries

Updated: 4 November 2026. Target requirements, not executable schemas or wire field names. Each feature records the user outcome, account/data/runtime boundary, failure behavior and evidence. Read [status](implementation-status.md) before interpreting a scope designation as available. Every provisioning subset requires baseline security/idempotency/capacity. No release plan is generated or endorsed here.

Product requirements are independent of historical CORE-5; [historical reference](academic-transition.md) preserves the old records without limiting this specification. Product VM/App Container changes require engineering contract review, not academic scope approval. No dependency is added by this document.

## VM Instances

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need owned general-purpose guest OS |
| User story | As an account member, request a VM and manage it without hypervisor access |
| Expected outcome | One ready owned guest-kernel VM with persistent owned disk/NIC |
| API responsibilities | Validate guest spec/capability, authorize, admit quota/capacity and record operation |
| Data ownership | Account/resource/host/provider-task binding; proposed VM records and reservations need B review |
| Runtime/backend responsibility | Incus API preferred evaluation; QEMU/KVM target; inspect before create |
| Security boundary | Guest kernel stronger boundary; no host socket/mount/device/exec exposure |
| Failure behavior | Lost task response blocks conflicting mutation; readiness failure preserves data |
| Acceptance tests | Actual guest kernel, lifecycle, interrupted create one VM, foreign disk/NIC denial, restore |
| Current implementation status | Planned |
| Dependencies | Versioned VM/claim schema, host spike, keys, capacity, networks/storage |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## App Containers

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need lightweight trusted applications |
| User story | As a trusted user, run an approved image under explicit resource limits |
| Expected outcome | Owned ready shared-kernel app, never represented as VM |
| API responsibilities | Distinct type/spec, authorization, quotas and operation records |
| Data ownership | Owned image/runtime/volume bindings; separate from old compute interpretation |
| Runtime/backend responsibility | Restricted container provider; Docker candidate |
| Security boundary | Trusted-workload class; no privileged daemon/socket/root mounts |
| Failure behavior | Inspect same owned task/object before retry; no silent VM fallback |
| Acceptance tests | Type distinction, readiness, privilege denial, crash/idempotency and safe cleanup |
| Current implementation status | Planned |
| Dependencies | Versioned type/spec, policy, host capacity/network and driver |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Kubernetes clusters

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need usable local Kubernetes, not cluster metadata |
| User story | As an authorized account user, obtain a cluster and scoped kubeconfig |
| Expected outcome | Reachable API and workload result with correct ownership |
| API responsibilities | Quota/port/capacity reservation, operation and separately authorized kubeconfig export |
| Data ownership | Account cluster/ports/node/host identity; secrets stored by protected reference |
| Runtime/backend responsibility | k3d CLI provider; tested pinned pair; profile limits set by measured capacity |
| Security boundary | Protected kubeconfig and management API; no generic tenant privileged host access |
| Failure behavior | Recover half-create by recorded identity/task; deletion holds quota until confirmed |
| Acceptance tests | Existing tests 7–10 plus actual workload and measured capacity |
| Current implementation status | Planned |
| Dependencies | Claim/type/network/capacity, tested host, scoped secrets |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Managed PostgreSQL

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need actual private durable SQL and restore |
| User story | As an account user, write rows and recover deleted test data |
| Expected outcome | Ready private SQL endpoint with scoped credential and verified restore |
| API responsibilities | Own instance/volume/credential/backup intent; no fake readiness |
| Data ownership | Account-owned DB/volume/backup and protected secret references |
| Runtime/backend responsibility | Reviewed VM or trusted container hosting; actual SQL readiness |
| Security boundary | Separate from platform DB; no other-account network/key access |
| Failure behavior | Retain volume on stop/recreate; failed restore cannot destroy only recovery copy |
| Acceptance tests | Test 11, canonical rows/checksum and denied foreign reachability |
| Current implementation status | Planned |
| Dependencies | Compute/runtime, network/storage, key and backup policy |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Object storage

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need real byte storage with scoped access |
| User story | As an account user, upload/download an object without administrator keys |
| Expected outcome | Exact bytes and owned bucket/policy persist |
| API responsibilities | Java owns bucket/credential reconciliation and authorized retrieval |
| Data ownership | Tenant bucket records separate from operator-owned object-provider service/storage |
| Runtime/backend responsibility | Preferred RustFS candidate through replaceable provider adapter; exact artifact/profile pending; [provider evidence](provider-evaluation.md) |
| Security boundary | Per-account S3 policy; operator admin custody/console private |
| Failure behavior | Process restart rebuilds policies, not lost bytes; object restore separate |
| Acceptance tests | Tests 12–13 plus retained-object checksum and foreign-key denial |
| Current implementation status | Planned |
| Dependencies | Shared runtime bootstrap, secret custody, network path, data backup |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

RustFS is preferred; integration remains Planned. Operator-managed storage and Java-side bucket/policy/credential reconcile are distinct from agent process lifecycle. See [compatibility gaps and failure/restore requirements](provider-evaluation.md#rustfs-compatibility-and-acceptance).

## VPC/IPAM

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core product |
| User problem | Need predictable private connectivity without collisions |
| User story | As an account user, connect owned resources while others are denied |
| Expected outcome | One initial owned network/account with actual allowed/denied paths |
| API responsibilities | Allocate/reserve atomically; authorize attachment and deletion |
| Data ownership | Owned pool/subnet/IP/port/host/NIC bindings; no hidden JSON ownership |
| Runtime/backend responsibility | Incus VM networking plus explicit container/k3d capabilities; coordinated allocator |
| Security boundary | Default-deny ingress/egress and management separation; no competing IPAM |
| Failure behavior | Hold reservations through uncertainty and cleanup; never reuse on timeout alone |
| Acceptance tests | Test 14 races plus VM/container traffic, spoofing and management denial |
| Current implementation status | Planned |
| Dependencies | Versioned network exchange, backend capabilities, route/pool review |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Resource scheduler

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE safety; later Go extraction |
| User problem | Need safe admission on shared limited hardware |
| User story | As an operator, reject unsafe load without affecting existing resources |
| Expected outcome | Durable fair admission and understandable refusal/queue |
| API responsibilities | Validate fresh host capability and quota; reserve in intent transaction |
| Data ownership | Single authoritative host/account reservations; no duplicate usage authority |
| Runtime/backend responsibility | Agent reports measured capacity; coordinator initially Java-owned |
| Security boundary | No tenant priority escalation or stale capacity admission |
| Failure behavior | No reservation release for ambiguous task or unconfirmed teardown |
| Acceptance tests | Parallel admission, pressure/full disk, stale host and restart accounting |
| Current implementation status | Planned |
| Dependencies | Capacity ADR, reports, locking, profile measurements; future Go ownership review |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Accounts and memberships

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE security |
| User problem | Need people and resources with clear authorization |
| User story | As a member, operate only permitted account resources |
| Expected outcome | Scoped tokens/actions and safe account administration |
| API responsibilities | Authenticate every management API, check membership/RBAC before mutation |
| Data ownership | Users/global identity, memberships, owned records, hash-only tokens/encrypted service keys |
| Runtime/backend responsibility | Backend project mapping is defense in depth, not authorization replacement |
| Security boundary | No operator rights from tenant identity; revocation/expiry and secure custody |
| Failure behavior | Store/policy outage fails closed; old account context never falls back |
| Acceptance tests | Foreign-ID 404, role/revocation/assignment and lookup/mutation race tests |
| Current implementation status | Partially Implemented; filters/interfaces only |
| Dependencies | Concrete credential store, foundational schema, role/custody decisions |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## CLI

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Core platform capability |
| User problem | Need programmable lifecycle without runtime administration |
| User story | As a user, create/wait/get/list in stable JSON |
| Expected outcome | Same authoritative API outcome across clients |
| API responsibilities | User-authenticated requests, mutation idempotency and scoped secret export |
| Data ownership | Config only; no local resource DB or service-secret reuse |
| Runtime/backend responsibility | No direct Docker/Incus/k3d; only API transport |
| Security boundary | No credentials in arguments/logs; finite timeout/no redirect |
| Failure behavior | Explicit non-2xx/protocol errors; cancellation differs from convergence |
| Acceptance tests | API-only dependency review, JSON, retries/wait and real user workflow |
| Current implementation status | Planned |
| Dependencies | Reviewed module/client boundary and public API/credential flow |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Go agent

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE execution foundation |
| User problem | Need trusted execution without public host admin |
| User story | As an operator, run one host agent with assigned account contexts |
| Expected outcome | Signed bounded communication and future authorized actions |
| API responsibilities | Server supplies verified assignment and current work/claim semantics |
| Data ownership | One canonical host identity; immutable per-account clients and task correlation |
| Runtime/backend responsibility | Built-in provider adapters; no arbitrary tenant commands/plugins |
| Security boundary | Protect privileged sockets/keys; replay-resistant TLS API |
| Failure behavior | Stop auth context safely; uncertain backend task blocks mutation |
| Acceptance tests | Existing unit loop/client plus real Java/no-work/outage/signal tests; future driver tests |
| Current implementation status | Partially Implemented; signed no-execution loop in code |
| Dependencies | Real server/keys/assignment and future protocol before driver dispatch |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Reconciliation

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE safety |
| User problem | Need restart/drift recovery without duplicates |
| User story | As a user, retry an interrupted request and retain one resource |
| Expected outcome | Durable convergence/attempt visibility and owner-safe repair |
| API responsibilities | Atomic intent/idempotency/claims; mandatory exact generation/owner/epoch guards. Lease-expiry/renewal/acknowledgement acceptance policy remains an explicit B+C+A decision; D reviews any wire/DTO impact |
| Data ownership | DB intent plus observed freshness/provider-task identity; progress not acknowledgment |
| Runtime/backend responsibility | Inspect before mutation, serialize and settle backend outcomes |
| Security boundary | DB fence plus host exclusion; no cleanup from failed inventory |
| Failure behavior | Unknown outcome fails closed; periodic inspections use current intent |
| Acceptance tests | Crash/lost response/paused worker/stale observation and drift cases |
| Current implementation status | Planned |
| Dependencies | Claims, durable attempts, provider conformance and host exclusion |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

Expiry is not silently a second fence or a runtime cancellation signal. Before implementation, B+C+A must agree consistent DB-time predicates for renewal, convergence and observations, including delayed/current-epoch reports and takeover races. Progress never advances acknowledgement; accepted convergence clears lease owner and expiry together while preserving epoch. An expired lease alone never permits conflicting backend work: inspect/settle the actual provider task under resource exclusion. [Open decision](decisions.md#lease-expiry-renewal-and-acknowledgement--open), [current V5 concurrency reference](../../V5/CoveStack-Database-Schema.md#73-exact-generation-acknowledgement-and-separate-progress).

## Snapshots and Backups

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE recovery plus operator responsibility |
| User problem | Need actual recovery after data/host failure |
| User story | As an operator, restore known workload and control-plane data |
| Expected outcome | Verified usable data/VM, not just an archive |
| API responsibilities | Authorize backup/restore, retention and isolated recovery orchestration |
| Data ownership | Owned backup metadata, protected off-host copy and separate key recovery |
| Runtime/backend responsibility | Backend-specific consistency/export and DB/object restore adapters |
| Security boundary | Encrypted transport/storage where tested; foreign reads denied |
| Failure behavior | Reject corruption/incompatibility; preserve last good copy; reconcile restored intent |
| Acceptance tests | Known rows/files/object checksums, VM boot/snapshot consistency and control DB/key recovery |
| Current implementation status | Planned |
| Dependencies | Storage/secret format/retention ADR and authorized recovery host |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Audit

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | CORE security/operations |
| User problem | Need explainable administrative decisions |
| User story | As an operator, trace an action without secret disclosure |
| Expected outcome | Scoped append-only correlated events and redacted user diagnostics |
| API responsibilities | Capture identity/account/action/resource/operation/result and failed auth |
| Data ownership | App-role INSERT-only events; unauthenticated-event scope decision |
| Runtime/backend responsibility | Backend task/request correlation, never arbitrary raw logs as authority |
| Security boundary | Protected reads/retention; administrator rewrite threat explicit |
| Failure behavior | Audit policy under storage outage reviewed; no false tamper-proof claim |
| Acceptance tests | Privilege denial, missing/denied actions, redaction and operation correlation |
| Current implementation status | Planned |
| Dependencies | DB privileges/event shape/retention and monitored operations |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Optional Floci integration

| Field | Requirement / evidence boundary |
| --- | --- |
| Scope | Proposed; not core emulator |
| User problem | Need AWS-like app integration tests inside owned infrastructure |
| User story | As a developer, opt into a local emulator for a chosen app test |
| Expected outcome | Reachable pinned emulator with documented service/mode limits |
| API responsibilities | Authorize environment placement and lifecycle; do not proxy arbitrary cloud requests |
| Data ownership | Own environment/runtime/volume reference; not Floci internal tables |
| Runtime/backend responsibility | External Floci artifact in suitable trusted VM/container; no Java rewrite |
| Security boundary | Protected endpoint, explicit enforcement and no real AWS credentials/socket exposure to tenants |
| Failure behavior | Bounded startup/health/cleanup; external emulator persistence has its own semantics |
| Acceptance tests | SDK workflow, permission-mode verification, restart and owner-safe cleanup |
| Current implementation status | Proposed |
| Dependencies | Core runtime/network/capacity, license/artifact pin and use-case validation |
| Scope priority | Dependency-bound product target; [canonical core/future classification](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); no release assignment |

## Catalogue coverage and inherited specification fields

The [canonical catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) owns all 22 core and 31 later entries. Detailed specifications above describe common execution components; they do not reduce the full proposed core to those sections. For every catalogue entry: its requested capabilities define the user request/outcome, dependency column defines prerequisites, acceptance column defines required domain evidence, and status column is the current evidence disposition. Common ownership/lifecycle rules apply; unsupported operations are rejected, never guessed into a wire kind. No release target is assigned.

| Additional core capability | User problem / story and outcome | API/data/runtime/security and failure boundary | Evidence and dependencies |
| --- | --- | --- | --- |
| VM Images and Templates | Account user needs a known guest image and repeatable secure initialization | Authorize approved immutable image/reference; trusted provider fetch verifies provenance; failed verification blocks boot, no arbitrary host import | Boot exact image; reject modified/unauthorized image; depends on VM driver, artifact trust and storage |
| Subnets | Account user needs a non-overlapping owned range | Control plane owns reservations; backend applies approved subnet; ambiguous cleanup keeps allocation, no second allocator | Parallel accounts cannot overlap; depends on VPC/IPAM and actual routes |
| Network Interfaces | Account user needs an owned VM/container address attachment | Authorize both endpoint owners; provider realizes NIC/routes; interrupted attach inspected before retry | Foreign attachment denied and real connectivity verified; depends on network/storage ownership and VM driver |
| Firewalls / Security Groups | Account user needs explicitly permitted connectivity | Validate rules/owner and persist generation; backend enforces default deny; stale/partial apply never reported safe | Allowed packets pass and forbidden packets fail; depends on enforceable networking/provider |
| Block Storage / Volumes | Account user needs data surviving compute restart | Record owner/attachment/capacity and retention; provider realizes disk; failure preserves disk and prevents double attach | Actual data/restart/attach/resize/restore and foreign denial; depends on storage backend and capacity |
| Secrets Management | Authorized workload needs credentials without embedding plaintext | API authorizes protected references; mature custody backend/agent delivers narrowly; unavailable secret fails closed, logs redacted | Rotate/revoke, foreign retrieval denied, protected recovery; depends on reviewed provider/key custody |
| Resource Quotas | Account user needs predictable fair capacity | Durable limits/reservations checked with host capacity in transaction; stale observations block admission | Parallel requests cannot overspend; depends on schema/admission and authenticated host inventory |
| Monitoring and Resource Health / Basic Logging | Operator needs useful fresh status and sanitized failure diagnosis | Owned timestamped observations and correlated events/logs; telemetry cannot rewrite intent or expose other accounts/secrets | Lost agent becomes unknown/stale; redaction and cross-account queries tested; depends on signed reports, IAM and retention policy |

Current status for these entries is Planned, as recorded in the canonical table. Project/organization hierarchy, network policy fields and secret references are conceptual requirements, not new current contract fields. All require B/D schema/API review, C security/acceptance and A provider review; E exposes only authenticated API behavior.
