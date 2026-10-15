# CoveStack V5 — API Boundaries and Compatibility Matrix

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. One backend model

All user-facing interfaces call core authorized services. Provider-style and management controllers adapt those services; they must not implement duplicate resource state or bypass account checks.

| API area              | Consumer                                 | Policy                                                           |
| --------------------- | ---------------------------------------- | ---------------------------------------------------------------- |
| Internal signed agent | Go worker                                | Service identity and private route                               |
| Provider-style subset | Tested /v2 consumers                     | Bearer account scope and exact DTO mapping; no full parity claim |
| CoveStack management  | CLI and optional provider                | CORE-5 resources and one stack integration; stable JSON          |
| S3 endpoint           | Application/S3 client                    | Scoped storage credential and signed protocol                    |
| Kubernetes API        | Authorized kubectl/workload operator     | Protected kubeconfig and route                                   |
| Extended interfaces   | [Roadmap](CoveStack-Extended-Roadmap.md) | Prior interface design relocated; not SRS scope                  |

## 2. Existing internal endpoint inventory

| Method | Path                                         | Purpose                             |
| ------ | -------------------------------------------- | ----------------------------------- |
| POST   | `/internal/v1/agent/heartbeat`               | Presence and installed capabilities |
| GET    | `/internal/v1/agent/work`                    | Scoped desired work                 |
| PUT    | `/internal/v1/agent/resources/{id}/observed` | State/generation observation        |
| GET    | `/internal/v1/agent/known-ids`               | Defined complete scoped inventory   |
| POST   | `/internal/v1/agent/events`                  | Bounded report/event input          |

GET work/known-ids scopes are authenticated agent/account plus advertised supportedKinds, excluding server-owned groups/buckets. No caller-selected account parameter exists. Current service-key verification derives accountId/accessKeyId/kind; assignment to an agent and multi-account provisioning/poll selection remain B+C+A decisions. Ownership labels do not grant account scope. Known-ids cleanup requires a complete scoped inventory; a timeout/partial response authorizes no deletion.

Actual request/response shapes remain in `contracts/internal-api.yaml`. V5 does not replace them or certify running controllers. New fields/routes require reviewed schemas, examples and tests.

## 3. CORE-5 management models

| Model              | Required semantics                                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------------------ |
| Resource/operation | Account scope, generation, lease/epoch; operation states queued, running, completed, failed, cancelled |
| VPC/IPAM           | Server-owned intent; one private network/account; unknown states TBD with B; attachment schema B+A     |
| Stack/template     | One immutable template version and fixed sequence; child IDs; no DAG engine                            |
| IDs                | Internal uuid; CLI prefixed public_id; /v2 numeric provider_id; B+D finalize                           |
| Automation JSON    | Stable list/get output for every service; no credential/kubeconfig content                             |

CLI, Terraform provider and future inventory use only authenticated API, never Docker/k3d or direct DB. Every mutating CLI command carries an idempotency key; pending is user-facing behavior, not a newly invented operation state. Create returns immediately; --wait polls. Endpoint additions/fields/error strings need B+D review and coordinated schemas/fixtures/generated outputs; absent values are TBD with B.

## 4. Compatibility register

| Field              | Record per tested feature              |
| ------------------ | -------------------------------------- |
| Endpoint/method    | Exact supported operation              |
| Request fields     | Required/optional/unsupported          |
| Response and error | Actual status/body and pagination      |
| Client/tool        | Version tested and endpoint override   |
| Async behavior     | Polling/timeout/action result          |
| Limitation         | Container/cluster/provider differences |

DigitalOcean-style IDs/actions need explicit mappings rather than leaking internal identifiers by accident. A custom Terraform provider does not automatically make the standard DigitalOcean provider compatible. A real S3 endpoint is distinct from a management bucket record. Website reachability is distinct from Kubernetes admin connectivity.

## 5. Cross-interface conflicts

Store a durable operation once, then let CLI wait or inspect it. Protect concurrent writes with optimistic/version checks and idempotency. Define optional provider ownership of desired fields. Broader UI/IaC conflict rules are preserved in the roadmap, not an early UI task.

## 6. Secrets and error policy

Ordinary resource responses contain safe metadata. Passwords, tokens and kubeconfigs use separately authorized retrieval with restricted logging. Errors name the failed operation without returning secret configuration or upstream response dumps. Other-account GET/DELETE/kubeconfig requests return 404 with no existence leak; tests verify side effects as well as status.

## 7. Acceptance

C owns tests 1–15 and every acceptance fixture from day 1. A reviews Go/client and wire integration; B schema/core; D Java /v2 DTOs/fixtures; E thin CLI and later provider. D writes no Go. Test real clients, negative cases and the mandatory bounded Kubernetes subset; do not return success-shaped stubs.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
