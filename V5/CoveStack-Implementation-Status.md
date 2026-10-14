# CoveStack V5 — Baseline Status, Missing Work and Contract Changes

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Snapshot interpretation

The inspected working tree contains local configuration changes and existing security/lifecycle libraries. This document does not verify a new build or live deployment. **Documentation V5 describes the target, not the state of implementation.** Verify current PR/CI status before onboarding another developer.

| Component | Inspected evidence | Next verification |
|---|---|---|
| Go config | Loader/tests exist locally | Review/commit and pinned toolchain |
| Signing | Rules/request/signer and shared tests | Real server/store acceptance |
| Lifecycle | Source/generator/helpers and fixtures | Wire into runtime services/drivers |
| Agent command | Auth-smoke exists | Production client/main/polling |
| Java auth | Filters/interfaces/stores scaffold | Real credentials/tokens and application boot |
| Java packaging | Boot plugin skip remains enabled | Entry point then package/boot |
| Compose | Private Postgres/Valkey/Mailpit | Actual configured readiness/restore |
| Public resource services | Contract/design records | Gate evidence for real workflows |

Earlier local tests passed with Go 1.27.1; that is not pinned 1.27.2 evidence. V5 generation itself does not run Java/Go integration tests or inspect real secret values.

## 2. Planned work register

| CORE-5 service | Status | Required evidence | Owner |
|---|---|---|---|
| Compute/Droplets | Planned | G1–G4; tests 1–6 | A+B; C tests; E CLI |
| Kubernetes | Planned | G1–G4; tests 7–10; hard cap 2 clusters/host | A+B+D; C tests; E CLI |
| Managed PostgreSQL | Planned | G1–G4; test 11; backup/restore/scoped credentials/private network | A+B; C tests |
| Object storage | Planned | G1–G4; tests 12–13; Java recreates buckets after MinIO restart | B+C server; A process |
| VPC/IPAM | Planned | G1–G4; test 14; ten IP allocations and ten account VPCs | C allocator; B records; A helper |

Cross-cutting status is Planned: authenticated CLI, Security, Recovery, optional Terraform Droplet/K8s, one stack integration (test 15), Ansible Part 1. None is a sixth service. Extended features remain [Extended track](CoveStack-Extended-Roadmap.md), not deliverables; optional future items are separate.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 3. Contract and schema decision checklist

| Open item | Required coordinated review |
|---|---|
| VPC/compute attachment | VPC server-owned, no contract kind; attachment schema/capability change B+A; state TBD with B |
| Multi-account polling | One credential-account poll; trusted agent lookup and multi-account credential model B+C+A |
| Stack | One fixed-sequence integration; missing states TBD with B; no invented DAG wire model |
| IDs/DTOs | Internal uuid, CLI public_id, numeric /v2 provider_id; B+D finalize |
| Credential custody | Hash-only api_tokens; encrypted service-kind signing_credentials; B+C |
| MinIO administration | Server reconcile needs private administration path, custody and rotation B+C |
| Unauthenticated audit | Nullable reference option versus separate security events; final B+C decision |
| Host/labels/ports | Canonical physical agent; proposed reserved kan.* labels; port release after cleanup |

Lifecycle contract stays **1.1.0**. It supplies existing kinds/states; missing enums are TBD with B. Any new message type needs coordinated schema, fixtures and generated outputs, not an isolated catalog edit. Flyway is declared but no migrations/entities/repositories were found; **no schema exists yet**. The markdown DDL is reference design, not executable or migration-ready. B finalizes staged migrations; this task creates none.

Extended route/hostname, metadata, health freshness, domain and broader automation decisions remain in the roadmap; their prior design requirements are preserved there, not promoted by this register.

## 4. Update and freeze procedure

Record revision/PR, environment, exact test results and constraints. Use only Planned, In progress, Gates passed or Extended track. A mock/source file is not service gate evidence. Promote a service only when G1–G4 all pass; cite acceptance test numbers and CI evidence. At freeze, relabel any CORE-5 service missing a gate honestly and list the missing evidence. Never write that all services are complete.

Keep pre-existing source changes separate from documentation review. Batch 3 topic guides now use CORE-5 or roadmap relocation pointers; historical text is preserved in the roadmap and does not override the SRS/gate register. Submission deadline: Monday, 25 January 2027.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
