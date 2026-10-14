# CoveStack V5 — API Boundaries and Compatibility Matrix

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. One backend model

All user-facing interfaces call core authorized services. Provider-style and management controllers adapt those services; they must not implement duplicate resource state or bypass account checks.

| API area | Consumer | Policy |
|---|---|---|
| Internal signed agent | Go worker | Service identity and private route |
| Provider-style subset | Supported SDK/CLI | Bearer account scope and exact DTO mapping |
| CoveStack management | UI/CLI/provider/inventory | Stack/network/publication operations |
| S3 endpoint | Application/S3 client | Scoped storage credential and signed protocol |
| Kubernetes API | Authorized kubectl/workload operator | Protected kubeconfig and route |
| Self-metadata | Bound workload | Instance identity, no control-plane secret |

## 2. Existing internal endpoint inventory

| Method | Path | Purpose |
|---|---|---|
| POST | `/internal/v1/agent/heartbeat` | Presence and installed capabilities |
| GET | `/internal/v1/agent/work` | Scoped desired work |
| PUT | `/internal/v1/agent/resources/{id}/observed` | State/generation observation |
| GET | `/internal/v1/agent/known-ids` | Defined complete scoped inventory |
| POST | `/internal/v1/agent/events` | Bounded report/event input |

Actual request/response shapes remain in `contracts/internal-api.yaml`. V5 does not replace them or certify running controllers. New fields/routes require reviewed schemas, examples and tests.

## 3. Proposed management models

| Model | Required semantics |
|---|---|
| Resource/action | Stable IDs, pending/completed/failed outcome |
| VPC/attachment | Account/host/CIDR ownership and uniqueness |
| Stack/template | Immutable template version, DAG and child IDs |
| Publication/domain | Verified hostname, approved upstream and lifecycle |
| Metrics/doctor | Scope, timestamps and unknown/stale status |

Endpoint paths for these extensions must be agreed before coding; illustrative names are not an implemented API. Document pagination, timeouts, errors and idempotency for each operation.

## 4. Compatibility register

| Field | Record per tested feature |
|---|---|
| Endpoint/method | Exact supported operation |
| Request fields | Required/optional/unsupported |
| Response and error | Actual status/body and pagination |
| Client/tool | Version tested and endpoint override |
| Async behavior | Polling/timeout/action result |
| Limitation | Container/cluster/provider differences |

DigitalOcean-style IDs/actions need explicit mappings rather than leaking internal identifiers by accident. A custom Terraform provider does not automatically make the standard DigitalOcean provider compatible. A real S3 endpoint is distinct from a management bucket record. Website reachability is distinct from Kubernetes admin connectivity.

## 5. Cross-interface conflicts

Store a durable operation once, then let UI/CLI wait or inspect it. Protect concurrent writes with optimistic/version checks and appropriate idempotency. Define external IaC ownership; do not let UI and Terraform silently fight over the same desired fields.

## 6. Secrets and error policy

Ordinary resource responses contain safe metadata. Passwords, tokens and kubeconfigs use separately authorized retrieval with restricted logging. Errors name the failed operation without returning secret configuration or upstream response dumps. Missing/inaccessible resources may use a reviewed non-disclosing status policy; tests verify side effects as well as status.

## 7. Acceptance

Test real clients and failures, not only schema conformance. Record explicit unsupported operations and reject them rather than returning fake success. Coordinate A/B/C for wire changes, D for provider mapping and E for user interface behavior.
