# API, CLI and compatibility boundaries

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. The [current 1.1.0 contract](../../contracts/internal-api.yaml) is unchanged. This is documentation, not a new OpenAPI/DTO/SQL version.

## What exists

| Existing boundary | Source and behavior | Limitation |
| --- | --- | --- |
| Signed agent discovery | GET /internal/v1/agent/work; 200 with array, [] means no work | Read-only; server-owned groups/buckets excluded; not a VM/claim API |
| Agent configuration | KAN_SERVER_URL, KAN_ACCESS_KEY, KAN_SECRET_KEY, KAN_RULES_PATH | Environment loader does not source .env; one configured credential context |
| Signing | Current KAN1 HMAC headers/canonical bytes and nonce rules | Not encryption, user RBAC or backend work ownership by itself |
| Client/loop | Bounded response, timeout, redirect refusal, cancellable poll/backoff/auth stop | Stops on returned work; no provisioning driver |
| Public auth filters | Existing /v2 and /kan/v1 registration and service interfaces | No runnable scoped resource controllers/concrete credential store established |
| CLI | Go/cobra target | No current public user CLI/provider implementation found |

The default http://localhost:4700 is a contract example for same-host development, not a verified live listener. Remote agent origin requires verified HTTPS. No runnable VM commands are published until code and acceptance exist.

## Target public API and clients

All management requests authenticate user/service identity and authorize account/action before lookup/mutation. Ordinary CLI uses user credentials, not an agent service key. No direct backend or platform DB calls. A future UI and Terraform provider consume the same authoritative API. CLI config is endpoint/key identifier/profile; secrets protected separately; no local resource database.

Mutations carry a logical idempotency key and validated method/route/body fingerprint. Same key/body returns original result; changed fingerprint is rejected per agreed existing policy, not conflated with two identical intentional requests. Atomic transaction includes response, resource/operation/reservations/events. Create returns pending with durable identity; wait polls until actual convergence/failure/cancellation. Every list/get supports stable JSON. Redacted get separates desired/converged revision, observation freshness, operation/lease and error; secret export is separately authorized. Exact output fields/schema are future review, not invented here.

Optional Terraform stays bounded to approved compute/Kubernetes compatibility resources until VM API migration is settled. Do not claim current doctl/godo/AWS SDK/full Terraform compatibility. Internal UUID, CLI public ID and numeric /v2 provider/action IDs remain separate mappings requiring B+D approval.

## Mandatory future versioned evolution

| Contract gap | Required coordinated change |
| --- | --- |
| VM versus App Container | Explicit type/spec/backend capability distinction and guest-image/disk/NIC validation; no silent old container conversion |
| Claim | Separate authenticated idempotent mutation, account/agent assignment, exact revision, epoch/expiry and duplicate-result behavior |
| Renewal/report | Live unchanged claim correlation, progress versus convergence, stale/future/out-of-order rejection and lease release |
| Capacity | Fresh host evidence, durable reservations and graceful refusal/queue semantics |
| Networks/storage | Server-owned network and allocation exchange plus owned attach/cleanup evidence |
| Compatibility | Mixed-version agent/client matrix, rejected unknown schemas and recorded migration/rollback path |

D/B own schema/DTO and B SQL; A client/provider; C security/acceptance. Edit authoritative base sources, fixtures and generated outputs together in a separately authorized implementation change. Do not edit a generated contract alone or predeclare an approved v1.2 wire version.

## Error/retry requirements

Timeouts and transient errors must not masquerade as [] or success. Fresh signature per HTTP attempt; retry a mutation only under reviewed idempotency/outcome semantics. Account 401/403 stops that context; temporary backoff/429 policy and cancellation follow tested loop behavior. Invalid spec, unsupported host/backend or capacity refusal is explicit and sanitized; no fabricated status/error name introduced here. No response body, secret reference value, signing material or kubeconfig in ordinary logs/audit/errors.
