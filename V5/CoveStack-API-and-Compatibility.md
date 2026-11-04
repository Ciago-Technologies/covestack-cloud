# CoveStack V5 — API, Identity and Compatibility Boundaries

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Existing wire behavior

| Existing contract / code | Meaning | Limit |
| --- | --- | --- |
| contracts/internal-api.yaml, generated from base/catalogue | Wire version 1.1.0 and current exact kinds/states | Not VM-first type/claim/capacity support |
| GET /internal/v1/agent/work | Read-only assigned agent/account discovery; successful [] means no work | Actual server/foundation/assignment path still missing or unverified |
| Heartbeat schema | Existing agentVersion/dockerVersion/supportedKinds | Existing Go work client does not implement full host registration/capability exchange; Docker-only version field needs future reviewed evolution |
| PUT observed / events / known IDs schemas | Observation/event/scoped inventory definitions | No current agent execution/report/cleanup implementation is established |
| KAN1 signed internal requests | Shared canonical bytes, nonce/time validation | Authentication is not authorization, trusted host assignment or runtime fencing |
| Public auth filter | User Bearer identity on /v2 and /kan/v1 | Concrete token store/membership and full scoped APIs remain pending |

Current resource kinds include compute.instance, db.instance, k8s.cluster, compute.group, compute.snapshot, network.load-balancer and storage.bucket. Contract membership is not service support. Server-owned groups/buckets/operations/metadata are never ordinary agent work. Labels/caller IDs are not trusted ownership.

## 2. Public product resource model

VM Instance, App Container, project/environment, NIC/firewall/volume and provider task are product concepts requiring reviewed schemas; no new wire kind/name/version is selected here. Preserve current kind/state/header/module names until coordinated evolution. Do not reinterpret a legacy container image/spec as a VM guest image or silently convert running containers. Public “Droplet” compatibility mapping needs an explicit migration/adapter decision.

Internal UUID, prefixed CLI public ID and numeric bounded /v2 provider ID are separate identities; B+D finalize mappings/uniqueness. Backend IDs are opaque owned references, not client-supplied authority. Kubeconfig, DB passwords and S3/signing secrets never appear in ordinary list/get/spec/observed/events.

## 3. Target request semantics

Authenticate → authorize membership/account/project/action → validate type/capability/dependencies → physical/quota admission → atomically persist intent/operation/reservations/idempotent response/event. Create returns a durable pending operation; actual convergence may occur later. Same idempotency key/fingerprint returns one result; changed body returns reviewed mismatch behavior. No backend operation inside the resource transaction.

List/get use stable JSON and owner predicates; get separates desired/converged revision, observed freshness, operation/attempt and sanitized failure. Foreign resources return 404 without existence leak. Secret/export endpoints are separately authorized/audited; never export kubeconfig through ordinary cluster list/get or observations.

## 4. Agent assignment and future authority

One process per host; separate immutable credential/client context per account under host-wide resource coordination. Server derives scope from stored credential/account/host assignment. Current principal record carries account/key/kind only; it does not implement trusted host binding. B+C+A/D must agree enrollment, provisioning, revocation, capability binding and dispatch. Do not add unsigned arbitrary agentId/account selectors as a workaround.

| Future contract area | Required agreement before code |
| --- | --- |
| VM versus App Container | Distinct schema/type/backend capability and old-container migration; no implicit fallback |
| Claim / renew | Authenticated idempotent current resource/generation/owner/epoch/expiry and host assignment |
| Progress / convergence | Separate progress/health from successful acknowledgement; exact generation/owner/epoch; release lease together |
| Capacity / capabilities | Fresh host evidence, backend compatibility, durable reservations and safe refusal |
| Network/storage attachment | Ownership, locality, allocation generation and confirmed detach/cleanup |
| Async provider task | Accepted-task identity, lost response, cancellation, exclusion and unknown-outcome resolution |
| Upgrade / rollback | Compatible old/new clients/agents, schema/fixture negotiation and rejection of unknown work |

No v1.2.0 or future state string is presumed approved. Authoritative base/catalogue, fixtures, generated wire/language output, Java DTO/controller and Go client change together in an authorized task. [Existing-file plan](../EXISTING-CODE-CHANGE-PLAN.md).

## 5. Server-owned service APIs

VPC/IPAM have no current wire kind; reviewed Java records/allocator own intent, agent capability applies reviewed host network changes only after a contract decision. RustFS service process/volumes are operator-managed; Java adapter reconciles account-owned buckets, scoped credentials and policy. Bucket reconciliation must not be invented as a generic agent resource. Workload PostgreSQL is distinct from control-plane PostgreSQL.

## 6. Client, provider and UI rules

CLI/Terraform/OpenTofu provider/Ansible inventory/UI use only authenticated API; no direct DB or privileged runtime calls. CLI is Go/cobra, config-only, secret-safe, stable JSON, idempotent mutation and finite wait/cancel/error behavior. Terraform state may contain sensitive exports; minimize and protect it. Do not advertise full AWS/DO/provider parity. Runtime-driver backend APIs remain private implementation, not public control-plane schema.

## 7. Error, retry and evidence

Errors are never translated into []. Fresh signature/nonce per retry; TLS outside loopback and no credential-bearing redirects. Auth rejection stops that account context. Current loop retries selected HTTP statuses and timeout-class transport errors; connection-refusal/general outage policy needs review, not a claim of universal retry. Unknown mutation outcome requires idempotency/task inspection, not blind retry. Preserve useful sanitized correlation, never raw secret responses. Tests must include real assignment, auth/replay/outage and cross-account behavior, not only fixtures.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-api-and-compatibility) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
