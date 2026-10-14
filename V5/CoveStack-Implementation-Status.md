# CoveStack V5 — Baseline Status, Missing Work and Contract Changes

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Snapshot interpretation

The inspected working tree contains local configuration changes and existing security/lifecycle libraries. This document does not verify a new complete build or live deployment. **Documentation V5 describes the target, not the state of implementation.** Verify current PR/CI status before onboarding another developer.

| Component | Inspected evidence | Next verification |
|---|---|---|
| Go config | Loader/tests exist locally | Review/commit and pinned toolchain |
| Signing | Rules/request/signer and shared tests | Real server/store acceptance |
| Lifecycle | Source/generator/helpers and fixtures | Wire into runtime services/drivers |
| Agent command | Auth-smoke exists | Production client/main/polling |
| Java auth | Filters/interfaces/stores scaffold | Real credentials/tokens and application boot |
| Java packaging | Boot plugin skip remains enabled | Entry point then package/boot |
| Compose | Private Postgres/Valkey/Mailpit | Actual configured readiness/restore |
| Public resource services | Contract/design records | Complete real workflows |

Earlier local tests passed with Go 1.27.1; that is not pinned 1.27.2 evidence. V5 generation itself does not run Java/Go integration tests or inspect real secret values.

## 2. Planned work register

| Area | Initial state | Completion evidence |
|---|---|---|
| Droplets | Driver/controller pending | Real lifecycle and restart |
| Kubernetes | Integration pending | Cluster and app connectivity |
| DB/S3 | Runtime/policy pending | Persistence and scoped clients |
| VPC/IPAM | New extension | Lease/isolation/recovery |
| Gateway/domains | New managed workflow | External HTTPS and safe claims |
| Metadata | New self-identity service | Exact response and denied access |
| Stacks | New orchestration | Dependency/failure/retention |
| UI/CLI | Common client workflows pending | Real cross-interface parity |
| Terraform/Ansible | New integration | Provider lifecycle and playbook |
| Monitoring/recovery | Plan/evidence pending | Alert and measured restore |
| LB/ASG/snapshots | Designed services | Bounded complete acceptance |

## 3. Contract migration checklist

| Change | Required coordinated update |
|---|---|
| Network or attachment work | Kind/ownership/capability/state/schema decision |
| Stack model | Separate orchestration lifecycle and child references |
| Ingress mapping | Route/hostname/port/status/security schema |
| Metadata endpoint | Provider subset, identity and forbidden fields |
| IaC integration | Resource ownership/import/refresh/retention semantics |
| Health data | Scope, timestamps and stale/unknown representation |

Keep `lifecycle-rules.json` and its version authoritative for existing runtime behavior. Review additions with A/B/C, generate all outputs, validate examples and negative fixtures, then implement. Do not edit generated catalogs alone.

## 4. Update procedure

For each feature add commit/PR, accepted behavior, tests and actual constraints. Mark planned, in progress, verified or unsupported with evidence. Do not confuse source presence with operational completion. The supported-feature matrix is the release truth; archived documents are historical reference.
