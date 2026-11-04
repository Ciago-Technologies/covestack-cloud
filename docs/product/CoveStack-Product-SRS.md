# CoveStack product SRS — infrastructure control-plane requirements

Updated: 4 November 2026. Status: **Planned target specification**, not implementation or acceptance evidence. Product scope is independent of academic CORE-5; old V5 SRS remains historical reference only. This SRS adopts the core/future classifications in the [canonical service catalogue](CoveStack-Product-Architecture-and-Service-Catalogue.md), without creating new wire types, fields or states.

## Scope and intended outcome

Account/project members request actual VM, container, Kubernetes, network, storage and database infrastructure through an authenticated API/CLI. The control plane persists owned intent and operations, admits safe capacity, coordinates bounded environments, reconciles interrupted work and verifies recovery. Core platform capabilities include IAM, scoped keys/secrets, quotas, audit/events, health/logging and backup foundations. Future candidates remain deferred as classified in the catalogue; old academic service counts, node caps, dates and template limits do not constrain the product.

## Compute and orchestration

Primary compute is a real guest-kernel VM; App Container is a distinct shared-kernel resource for trusted workloads. Incus/QEMU/KVM is preferred VM evaluation, not implemented support. k3d is the initial Kubernetes candidate, with capacity/profile limits measured before support. Owned images, disks, NICs, networks and firewall policy are explicit dependencies. A bounded environment records child ownership/readiness, retries safely and preserves shared or retained data; no arbitrary DAG/script engine is promised. See [HLD/LLD](../platform/architecture.md), [resource model](../platform/resource-model.md) and [environments](../platform/environments.md).

## Object storage — RustFS preferred, integration Planned

| Requirement | Required behavior | Acceptance evidence |
| --- | --- | --- |
| Operator-managed backend | Operator owns RustFS process, durable volumes, private administration and TLS | Tenants cannot reach administration; host/process restart preserves known bytes |
| Account-owned buckets | Java persists/reconciles bucket, policy and credential intent through a replaceable provider adapter | Repeated create/uncertain admin response yields one owned result; restart rebuilds configuration from DB without deleting data |
| Scoped access | Narrow per-account credentials and bucket policy; no tenant root credential | A cannot list/read/write/delete B's data; revocation and rotation enforced |
| Actual storage | Upload, read, inspect and delete supported objects through authenticated S3 subset | Known bytes/checksums and advertised listing/multipart behavior tested with selected Java SDK/artifact |
| Verified recovery | Independent backup of required bytes, metadata, ownership/policies and protected keys | Restore into isolated target and verify known data plus access denial; empty bucket recreation is not restoration |
| Honest compatibility | Expose only tested provider capabilities; reject unavailable options | No full S3 parity or unsupported ACL/access-logging/ownership-control promise |

RustFS is a preference, not a qualified deployment. Exact version/digest, admin client/permissions, key custody, persistent layout, consistent backup/restore, upgrade and security profile remain review items. SeaweedFS and other providers remain possible through conformance-tested adapters. Upstream findings and citations are centralized in [RustFS compatibility and acceptance](../platform/provider-evaluation.md#rustfs-compatibility-and-acceptance).

## Security, capacity and failure behavior

Authorize account/project/action before mutation; retain desired/observed generations and operation history. Persist idempotent intent/reservations atomically. Claims and runtime outcomes require coordinated versioned contracts; exact-generation/owner/epoch fencing is a target, not a current wire guarantee. Privileged agent/runtime sockets are tenant-inaccessible. Default-deny traffic and scoped secrets/storage require actual negative tests. Reject unsafe capacity or stale placement evidence; an 8 GB historical host is not a production size guarantee. Single-host operation is not HA.

G1 persisted authoritative intent; G2 crash-safe inspect/reconcile without duplicates; G3 enforced ownership and real network/storage denial; G4 meaningful automated failure acceptance. Supported additionally requires a qualified profile and maintenance policy. [Readiness](../platform/readiness.md), [capacity](../platform/capacity.md), [threat model](../platform/security.md) and [operations](../platform/installation-and-operations.md) define domain-specific evidence.

## Interfaces, current evidence and licensing

CLI/provider/inventory use the authenticated API only, never privileged backend sockets or direct database access. Mutations use idempotency; list/get offer stable JSON and truthful desired/observed freshness. Existing contract 1.1.0 is unchanged; B/D/A/C review future schema/fixtures/generated outputs before implementation.

Signed client, single-account polling and deliberate no-execution behavior exist in source. Full credential provisioning, trusted account/host assignment, claims and runtime execution remain pending. Source/test existence is not runtime qualification; [implementation ledger](../platform/implementation-status.md) is the evidence authority.

The maintainer approved Apache License 2.0 (SPDX `Apache-2.0`) for CoveStack-owned code/documentation; the complete [root LICENSE](../../LICENSE) is present. [Adoption and remaining obligations](../platform/licensing-decision.md) preserve third-party terms. No production support claim is made.
