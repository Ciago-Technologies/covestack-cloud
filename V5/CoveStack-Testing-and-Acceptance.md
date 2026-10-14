# CoveStack V5 — Test Matrix and Definition of Done

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Evidence levels

| Evidence | Establishes | Does not establish |
|---|---|---|
| Unit/vector | Local behavior and contract fixtures | Deployed system correctness |
| Fake transport/agent | Client/service orchestration under controlled inputs | Java/Docker integration |
| Real integration | Defined services work together | All production/security scenarios |
| Negative/security | Tested attack/control cases | Absolute security guarantee |
| Clean install/restore | Defined reproducibility/recovery | Untested hardware/provider conditions |

Every result identifies commit, date, actual tool/runtime versions, test setup, expected behavior, actual behavior and limitations. Store sanitized output. Do not fabricate success to align with a milestone.

## 2. Baseline commands

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
(cd kan-agent && go test -race ./...)
(cd kan-agent && go vet ./...)
(cd kan-server && mvn --batch-mode --no-transfer-progress verify)
```

Run from the repository root. A dependency-download or missing entry-point problem is recorded honestly. A local older toolchain pass does not prove the pinned version.

## 3. Service acceptance

| Area | Required cases |
|---|---|
| Authentication | Active/invalid/revoked/expired credentials; wrong credential kind |
| Signing | Tampered nonce/key/body/path/query; replay and nonce-poisoning |
| Droplet | Create/read/start/stop/delete; intentional stop vs crash |
| Recovery | Restart mid-work; inspect IDs; no duplicates |
| Observations | Stale/future generation, old lease and readiness acknowledgment |
| Kubernetes | Real ready cluster, kubeconfig and deployed test app |
| PostgreSQL | Connection readiness, private access and retained/restored data |
| S3 | Put/get/list/delete on scoped bucket; cross-account deny |
| VPC/IPAM | Concurrent/overlap/exhaustion allocation and denied attachment |
| HTTP LB | Healthy member routing, removal, empty pool and reload failure |
| ASG | Pending reservations, min/max/cooldown and outage pause |
| Snapshot | Owned image restore; volume data explicitly excluded |

## 4. Extension acceptance

| Area | Required cases |
|---|---|
| Public apps | Droplet and Kubernetes site from external mobile network |
| Domain policy | Unknown/unverified/stale hostname and wrong account |
| Metadata | Self-only identity; no secrets/public ingress; IP reuse cleanup |
| Stack | Dependencies/readiness, retry, partial failure and data retention |
| Terraform | Create/read/update/delete/import, drift and interrupted apply |
| Ansible | Supported configuration twice; denied/unreachable target |
| UI/CLI | Create in either, inspect/mutate in other; identical policy |
| Monitoring | Real timestamps, stale/unknown states and triggered alert |
| Restore | Separate environment with known records/objects verified |

## 5. Isolation checks

Test allowed same-VPC paths and forbidden account/control-plane paths by actual IP/port, gateway and published host endpoint. Inspect workload mounts/capabilities. Never equate a failed DNS lookup with blocked access. Use synthetic workloads and authorized lab targets.

## 6. Stack acceptance scenario

| Order | Expected result |
|---|---|
| Submit approved template | One durable stack/action |
| Network allocation | Owned non-overlapping bridge |
| DB/bucket readiness | Private DB and scoped storage |
| App deploy/config | References correct service endpoints |
| Publication | Trusted real HTTPS URL |
| CLI inspect UI-created stack | Same child IDs and outcomes |
| Dependency failure/retry | Truthful state and no duplicate children |
| Delete with retain | Data remains protected/recoverable |

A ready stack cannot be backed by success-shaped stubs. Test a failure at each dependency boundary; retain user data even if the app step fails.

## 7. Fair comparison and mentor evidence

Use the same workload/hardware and comparable security assumptions for Docker/Compose versus CoveStack. Measure workflow effort, account controls, publication setup, restart behavior, overhead and restoration. Credit Docker/Compose for mechanisms already supplied; show the added cloud-style account/template policy.

## 8. Release checklist

| Release check | Required evidence |
|---|---|
| Supported-feature matrix | Exact endpoints/actions/limits |
| Tests | Required CI and recorded live cases |
| Security | Cross-account denial and private-boundary tests |
| Artifacts | Pinned inputs/checksums and clean install |
| Docs | No fictional commands presented as runnable |
| Data | Retention and restore proof |
| Interfaces | Same API/state and non-sensitive listings |
| Rehearsal | Second-device app and complete stack demonstration |

All unresolved tests are recorded as pending/failed with an owner; “not tested” is not “passed.”
