# CoveStack V5 — Test Matrix and Gate Evidence

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

## 3. CORE-5 acceptance

C owns ALL acceptance tests from day 1. The [canonical exact-assertion table](CoveStack-AntiWrapper-Acceptance.md#2-core-5-acceptance--srs-tests) is authoritative; do not maintain a contradictory second numbering scheme.

| Service | Tests | Required evidence |
|---|---|---|
| Compute/Droplets | 1–6 | Crash, fencing, tenant 404, atomic idempotency, stale generation, server outage |
| Kubernetes | 7–10 | Crash-safe one cluster; scoped kubeconfig; owned cleanup; third-cluster quota rejection |
| Managed PostgreSQL | 11 | Known rows/checksum restored; private/scoped credentials; cross-account network denial |
| Object storage | 12–13 | Cross-account S3 denial; Java reconstructs bucket/credential policy after MinIO restart |
| VPC/IPAM | 14 | Ten IP allocations in one VPC; ten accounts creating non-overlapping VPCs; tenant denial |

Authentication/replay, invalid inputs, actual readiness, retained data and migration/security regressions supplement these tests. Never cut tests 1–15.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 4. Extended acceptance — relocated

The former extended/proxy/UI/metadata/domain/metrics scenarios are preserved in the [roadmap](CoveStack-Extended-Roadmap.md). Canonical E1–E8 assertions are in the [acceptance appendix](CoveStack-AntiWrapper-Acceptance.md#appendix-a-extended-track-tests--not-srs-acceptance). They are not SRS requirements. Interface tests reuse account/lifecycle evidence; bonus demonstrations require Gates passed.

## 5. Isolation checks

Test allowed same-VPC paths and forbidden account/control-plane paths by actual IP/port, gateway and published host endpoint. Inspect workload mounts/capabilities. Never equate a failed DNS lookup with blocked access. Use synthetic workloads and authorized lab targets.

## 6. Stack acceptance scenario — test 15

| Order | Exact target |
|---|---|
| One CLI command | Store one owned fixed-template stack |
| Dependencies | Private network; PostgreSQL + scoped MinIO; then app Droplet |
| Inspect | API-backed child IDs/generations and truthful progress |
| Partial failure/retry | Record failed step; resume without duplicate children |
| Delete with retention | Protect owned persistent data per reviewed policy |

No DAG engine, UI or public route is part of this integration acceptance. A stub does not establish readiness.

## 7. Fair comparison and mentor evidence

Use equal workload/hardware and comparable security assumptions. Measure account denial, atomic idempotency, restart convergence, quotas and known-data restore. Credit runtime tools for existing mechanisms; platform value must be shown by CORE-5 gate evidence, not by forms replacing YAML. The CORE-5 viva questions are in [Acceptance](CoveStack-AntiWrapper-Acceptance.md#5-core-5-viva-questions). No mentor acceptance is invented.

## 8. Release checklist

| Check | Required evidence |
|---|---|
| Gate register | G1–G4 per CORE-5 service; status honestly retained if missing |
| CI | Tests 1–15 and schema/auth/client regressions; C owns results |
| Installer/artifacts | Clean VM; pinned tested inputs/checksums; Ansible second run changed=0 |
| CLI | API-only dependency check; three cross-builds; two actual OS runs; stable JSON |
| Live demo | Tests 1, 3, 4 and get view; remaining tests as CI/video |
| Documentation | No fictional main or unapproved wire/DDL presented as runnable |
| Recovery | Control-plane backup and workload restore evidence; measured scenario |
| Network | Exact [Known constraints](CoveStack-Complete-Working-Setup.md#10-known-constraints), hotspot backup video |

Record untested/failed cases with owner; do not promote service status from mock results. No new CORE-5 features after 10 January 2027; deadline Monday, 25 January 2027.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
