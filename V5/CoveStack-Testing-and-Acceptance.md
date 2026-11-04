# CoveStack V5 — Testing, Failure Evidence and Acceptance Matrix

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Evidence levels

| Evidence | Demonstrates | Does not demonstrate |
| --- | --- | --- |
| Source inspection | Code/interface exists and apparent behavior | Running server, provisioning or isolation |
| Unit/mock/vector | Local logic/fixtures under test | Real DB/runtime/network/credential behavior |
| Build/static/contract | Compilation/analysis/generated/schema consistency | Resource readiness or restore |
| Real integration | Authenticated server/DB/Valkey/adapter protocol behavior | Every backend/profile or full service support |
| Runtime/failure acceptance | Actual guest/data/traffic/restore behavior under stated failures | Arbitrary production capacity/HA/security |
| Qualified deployment profile | Exact supported matrix/operating evidence | Other hardware/runtime versions or single-host HA |

C owns acceptance; A/B/D/E implement and supply real results. No application/infrastructure test ran for this documentation alignment.

## 2. Existing baseline commands

From kan-agent/ on an authorized dev setup:

```bash
go test -race ./...
go vet ./...
```

From kan-server/:

```bash
mvn --batch-mode --no-transfer-progress verify
```

From repository root for contract changes:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
python3 contracts/generate_lifecycle.py --check
```

Current CI runs Go race/vet, Java verify and signing/lifecycle checks. Generator freshness/integration/profile jobs are future reviewed changes if not already enforced. Manifest toolchain pins, prerequisites and historical version mismatch must be checked; no auto-download/install or test run is implied here. [Existing-file change plan](../EXISTING-CODE-CHANGE-PLAN.md).

## 3. Traceable failure matrix

| ID | Assertion family | Implementation/test ownership |
| --- | --- | --- |
| 1 | VM interrupted create | A+B; C tests |
| 2 | Fencing | B+A; C tests |
| 3 | Tenant denial | C+B+D/A |
| 4 | Idempotency | B; C tests |
| 5 | Stale generation | B+A; C tests |
| 6 | Control-plane outage | A+B; C tests |
| 7 | Kubernetes interrupted create | A+B; C tests |
| 8 | Kubernetes secret scoping | C+B+D/A |
| 9 | Kubernetes cleanup | A; C tests |
| 10 | Kubernetes capacity/quota | B+A; C tests |
| 11 | PostgreSQL restore | A+B; C tests |
| 12 | RustFS tenant access | B Java; C tests; A service |
| 13 | RustFS restart/reconcile | B Java+A process; C tests |
| 14 | Network/IPAM isolation | C allocator+B SQL+A backend |
| 15 | Environment partial retry | B+A+D/E; C tests |
| E1 | Autoscaling | B coordinator+A runtime; C tests |
| E2 | Load balancing | D/B APIs+A runtime; C tests |
| E3 | Snapshot foundation | A+B; C tests |
| E4 | Guest metadata | C policy+A delivery; B/D API |
| E5 | Publication | A+B+D; C tests |
| E6 | Domains | C+B/D; A route |
| E7 | Observation freshness | B+A/D/E; C tests |
| E8 | Ansible inventory | API consumer owner reviewed; C tests |


Exact assertions are centralized in [State Reconciliation Acceptance](CoveStack-AntiWrapper-Acceptance.md); do not maintain incompatible abbreviated definitions here. The historical 1–15/E1–E8 mapping remains for traceability; current VM/quota/RustFS definitions require fresh tests. Snapshot foundation/freshness core requirements are not demoted because their IDs start E.

## 4. Additional product acceptance

| Domain | Required negative/failure proof |
| --- | --- |
| VM images/guest | Exact guest kernel/readiness; tampered image/bootstrap rejected; interrupted resize/start/delete and foreign disk/NIC denied |
| App Container | Explicit shared-kernel type, no privileged host/socket/device mounts, image update failure and retained data |
| Network/firewall | Real allowed/forbidden traffic across VM/container/K8s/management, address spoof/reuse and parallel ranges/IPs |
| Capacity | Physical/account reservations cannot overspend; stale host/disk pressure safely refuse; unknown task retains capacity |
| Auth/secrets | Real assignment, revoked/expired/user-service/host confusion, replay outage, rotation/recovery and redaction |
| Recovery | Known VM/file/SQL/object content and policy/keys restored on isolated target; corrupt/foreign backup denied |
| Environments | Usable app data flow, partial failure/restart and safe compensation/shared dependency preservation |
| Observability | Stale/unknown versus desired/converged state, out-of-order probe, sanitized audit/logs and access denial |
| Operations | Independent clean install, mixed-version upgrade/rollback, secret/DB restore and exact artifact rights/provenance |

## 5. Test environment and CI policy

Real backend tests need a separately authorized disposable host/data profile, protected secrets and bounded cleanup; never run destructive restore/delete against production or unrelated workloads. Current sandbox lacks authorized runtime acceptance; use named mocks honestly for unit learning. Real Valkey tests are distinct from mocked setIfAbsent tests; clean migration tests cannot be inferred from V5/V6 SQL presence.

CI records exact revision/host/artifact versions, setup, assertions, timeout, result, cleanup and limitations. Failure injection includes lost provider response, paused stale worker, parent/child death, backend/server outage and unknown outcome—not only kill-and-restart happy paths. Never suppress failure or swap a fake for the real test to claim a gate.

## 6. User experience and demo

Immediate team checkpoint proves real signed no-work/auth/cancel behavior, no resource execution. Future CLI/API prove idempotency, owner denial, truthful generation/freshness and actual ready resources. Show runtime restore/traffic data as exact-revision CI/recorded evidence if not live; have backup video/network-independent plan. Do not show old MinIO/container results as new RustFS/VM acceptance.

## 7. Gate review and report template

Record service/subset, status, G1–G4 evidence links, environment/profile, failures/unverified options, security/retention/restore limits and owner/reviewer. Board Done is a merged task, not a service gate. [Gates](CoveStack-Service-Gates.md), [readiness](../docs/platform/readiness.md), [operations](CoveStack-Operations-and-Recovery.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-testing-and-acceptance) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
