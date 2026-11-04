# CoveStack V5 — State Reconciliation Acceptance and Viva

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Scope and honest claims

Retained filename is a compatibility link; acceptance is about demonstrable ownership, state, isolation and recovery, not a slogan. A handler that directly runs a backend create or writes proxy config without durable intent, authorization, idempotency and safe retry does not meet this design. Existing signed no-execution client is useful source functionality, not VM/RustFS acceptance.

Tests 1–15 and E1–E8 retain historical traceability labels. The **definitions below are current target assertions**; old definitions/results are in the historical snapshot and do not prove the VM/product tests pass. Prefix E is not a universal future-scope label: snapshot foundations and freshness are core requirements, while advanced policy/alerts remain classified separately.

## 2. Product acceptance definitions

| Test | Case | Exact assertion | Owner / execution | Schema or authority relied on |
| --- | --- | --- | --- | --- |
| 1 | VM interrupted create | Kill agent after backend accepts creation, before acknowledgement; restart inspects owned task and converges to exactly one real guest-kernel VM. Separate App Container variant remains required. | A+B; C tests; meaningful automated CI acceptance required, not run here | Owned resource/provider task binding, current generation and uniqueness; target schema review |
| 2 | Fencing | A claims, lease expires, B claims; A convergence report is rejected by owner/epoch/exact generation. Unknown old runtime task blocks conflicting takeover until settled. | B+A; C tests; meaningful automated CI acceptance required, not run here | Atomic claim/epoch and fenced update predicate; DB fence is not backend fence |
| 3 | Tenant denial | A cannot GET/DELETE/export B resource; foreign API returns 404 without existence leak. Add actual network/disk/storage denial. | C+B+D/A; meaningful automated CI acceptance required, not run here | Owner-scoped lookup and composite account references; project policy review |
| 4 | Idempotency | Ten parallel same-key/same-body creates yield one resource/response; same key with different fingerprint returns reviewed mismatch (existing design 422), without a second resource. | B; C tests; meaningful automated CI acceptance required, not run here | Unique account/key and resource/operation/response in one transaction |
| 5 | Stale generation | Report generation 3 after desired 4 cannot acknowledge 4 or replace newer observations; progress never advances acknowledged generation. | B+A; C tests; meaningful automated CI acceptance required, not run here | Exact-generation convergence plus independent observation ordering |
| 6 | Control-plane outage | Agent cannot authorize new mutations/renewals when server unavailable; creates/deletes nothing new. Empty/error/partial discovery is never cleanup authority. | A+B; C tests; meaningful automated CI acceptance required, not run here | Stored assignment/claim and authorized mutation policy; no insecure fallback |
| 7 | Kubernetes interrupted create | Kill agent during backend cluster create; restart finds owned partial task/cluster and converges to one cluster with usable API/workload. | A+B; C tests; meaningful automated CI acceptance required, not run here | Owned cluster/provider task and reserved API endpoint identity |
| 8 | Kubernetes secret scoping | B cannot export A kubeconfig (404); kubeconfig absent from ordinary list/get/events/observations. | C+B+D/A; meaningful automated CI acceptance required, not run here | Owner-authorized secret export and protected reference |
| 9 | Kubernetes cleanup | Delete removes only owned cluster containers/volumes/networks under retention; no owned execution leftovers or unrelated deletions. | A; C tests; meaningful automated CI acceptance required, not run here | Scoped complete backend inventory and confirmed cleanup before release |
| 10 | Kubernetes capacity/quota | A request over the configured measured profile/account cluster/node capacity is explicitly rejected; admitted clusters unaffected and parallel requests cannot bypass reservation. | B+A; C tests; meaningful automated CI acceptance required, not run here | Host/account locked capacity and unique live ports; no fixed academic count |
| 11 | PostgreSQL restore | Backup known rows, remove only test data, restore and compare expected rows/checksum; persistence survives restart and foreign account cannot reach/authenticate to A DB. | A+B; C tests; meaningful automated CI acceptance required, not run here | Owned volume/DB/secret/backup/restore records with verified recovery evidence |
| 12 | RustFS tenant access | A key cannot list/read/write/delete B bucket or objects/admin; valid own bytes round-trip and key revoke/rotate behavior enforced. | B Java; C tests; A service; meaningful automated CI acceptance required, not run here | Account bucket/policy/credential binding and protected admin custody |
| 13 | RustFS restart/reconcile | Restart provider and Java; existing bytes persist, Java restores missing intended bucket/policy/credential configuration from DB safely, without duplicates or resurrecting deleted intent. | B Java+A process; C tests; meaningful automated CI acceptance required, not run here | Durable operator volume and account bucket/key/policy intent; config rebuild is not byte restore |
| 14 | Network/IPAM isolation | Ten IP allocations in one VPC have no duplicate; ten accounts create networks in parallel with no CIDR overlap. Foreign inspection/attach denied and forbidden packets actually fail. | C allocator+B SQL+A backend; meaningful automated CI acceptance required, not run here | Live allocation uniqueness/pool-overlap protection and owned attachments |
| 15 | Environment partial retry | One authenticated request realizes usable app, SQL and RustFS bucket; failure is recorded; restart/retry reuses owned children; shared/retained dependencies survive deletion. | B+A+D/E; C tests; meaningful automated CI acceptance required, not run here | Stable environment/child/dependency/idempotency/reservation bindings |


## 3. Additional retained and cross-domain cases

| Test | Case | Exact assertion | Owner | Schema / authority |
| --- | --- | --- | --- | --- |
| E1 | Autoscaling | Missing member replaced only under current min/max/cooldown and capacity; server unavailable causes no unauthorized scaling. | B coordinator+A runtime; C tests | Owned membership and authoritative policy/reservations |
| E2 | Load balancing | Backend failure changes validated health/routing; missing derived proxy config is rebuilt from DB without exposing unrelated accounts. | D/B APIs+A runtime; C tests | Owned backends/health/route intent; derived config not authority |
| E3 | Snapshot foundation | Create supported VM/volume recovery point, change test data, restore to safe new owned target and compare file hash; B cannot read A snapshot. | A+B; C tests | Owned recovery metadata, consistency/key/retention and target identity |
| E4 | Guest metadata | Workload A cannot read B self-metadata through spoof/NAT/reused identity; no agent/cloud/DB/S3 secrets. | C policy+A delivery; B/D API | Current owned workload/interface/assignment identity; future schema review |
| E5 | Publication | Unregistered hostname returns 404; delete removes owned route; private management port cannot be published through it; actual HTTPS path verified. | A+B+D; C tests | Approved durable publication/target/domain generation and allowed port |
| E6 | Domains | Unverified hostname serves no workload; reassigned domain cannot serve previous tenant and previous proof is invalid. | C+B/D; A route | Owned domain claim generation/proof and route binding |
| E7 | Observation freshness | Stop agent: UI/API becomes stale/unknown, never healthy from old desired state. If alerts are enabled, verify the policy event/delivery separately. | B+A/D/E; C tests | Ordered timestamped observations and freshness policy; alerts future where classified |
| E8 | Ansible inventory | Bob inventory includes none of Alice resources; output from authenticated API/CLI JSON, never backend inventory or DB. | API consumer owner reviewed; C tests | User ownership/permission and sanitized stable API output |


VM image provenance, guest readiness, unauthorized disk/NIC attachment, App Container shared-kernel privilege denial, actual allow/deny traffic, parallel physical admission, scoped membership/secret rotation, independent RustFS byte restore, control-plane DB/key restore and clean-install/upgrade/rollback require named domain tests as well. Do not squeeze every new product requirement into a legacy number and lose its assertion.

## 4. Reference concurrency design — proposal, B finalizes

| Flow | Required guard / transaction | Failure assertion |
| --- | --- | --- |
| Claim | Trusted credential/host/account assignment, locked due resource, generation/owner/epoch/expiry; bump epoch atomically | Concurrent claims are distinct/exclusive under reviewed policy; arbitrary caller account parameter is not scope |
| Convergence | Match owned resource, exact generation, owner and epoch; expiry check follows explicitly agreed policy | Old/future/stale report affects no convergence row; release lease_owner and lease_expires_at together, retain epoch |
| Progress/health | Separate ordered observation/attempt statement; no acked_generation update | A healthy/progress report cannot acknowledge desired state |
| Idempotent create | Same transaction: key/fingerprint, resource, operation/reservation/event and stored response | Same key ten callers one result; changed fingerprint rejected |
| Unknown runtime outcome | Serialize resource/child operations and inspect provider task before takeover | No competing Incus/Docker/k3d mutation just because DB lease expired |
| Delete/release | Complete owned inventory, confirmed teardown and retention policy | No unrelated cleanup; port/IP/disk remains reserved while outcome unknown |

These are not executable SQL, approved VM wire messages or proof the DB exists. [Schema/concurrency](CoveStack-Database-Schema.md), [API gap](CoveStack-API-and-Compatibility.md).

## 5. Live demonstration and retained evidence

Use real highlights only for behavior that exists: signed no-work/auth-stop and cancellation first; once APIs/runtime exist, demonstrate interrupted create, tenant denial, same-key retry and actual desired/observed freshness. Show remaining real runtime/failure/restore evidence from CI or a recorded exact-revision video. Never claim mocked or historical screenshots as new VM/RustFS tests. No dependency on phone hotspot alone; no fabricated CLI/Java startup command.

## 6. Twelve viva questions — truthful answers

| Question | Answer and evidence boundary |
| --- | --- |
| What if the agent crashes mid-create? | Target: inspect owned provider task and resume one resource. Current agent has no executor; real crash test remains pending. |
| What if two agents act? | DB owner/epoch/generation guards plus host/runtime exclusion and unknown-task settlement are required. A DB fence alone does not stop a stale backend command. |
| Why not directly use Docker? | Docker already supplies useful execution/network features. CoveStack must prove additional account/environment ownership, durable intent, admission and recovery workflow. |
| Why not directly use Incus? | Incus already manages VMs/projects/network/storage. CoveStack is only valuable if its account/app/SQL/S3/recovery experience demonstrably reduces user work. |
| Why not directly use k3d? | k3d provides cluster execution. Target platform adds owned API, safe operations/ports/capacity and scoped exports; tests must prove the integration. |
| Is a compute instance a VM? | Product VM Instance must have a guest kernel. App Container is explicitly a different shared-kernel type; no VM driver is claimed today. |
| Can A read B kubeconfig or storage? | Target: foreign export denied and scoped object/traffic tests fail access. Existing filters alone are not proof of full service isolation. |
| What if the server is unavailable? | No new unauthorized runtime mutations. Current poll-loop no-execution guard remains; general transport retry behavior still needs integration review. |
| How many VMs/clusters fit? | Measure the exact profile and reserve platform overhead; no old 8 GB/two-cluster assumption qualifies a production host. |
| Why does the CLI not use the DB/backend directly? | Authenticated API centralizes ownership, transactions, admission and operations; direct access would bypass these boundaries. |
| Which services have passed acceptance? | Use actual evidence ledger and CI/runtime artifacts. This document promotes none and establishes no supported production profile. |
| What remains unimplemented? | Server foundation/credentials/assignment/claims, VM/runtime/data services and full clients require code and real evidence; see current status rather than a diagram. |

G1: PostgreSQL desired state and generation are authoritative. G2: crash-safe inspection/reconciliation rebuilds state/config without duplicates. G3: account authorization denies foreign API resources (404 without existence leak) and actual traffic/storage boundaries are tested. G4: meaningful failure acceptance runs in CI on an authorized capable runner. **Acceptance Gates Passed** requires this evidence; **Supported** additionally needs a declared subset, qualified profile and maintenance policy.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-antiwrapper-acceptance) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
