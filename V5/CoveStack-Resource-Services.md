# CoveStack V5 — Resource Specifications and Runtime Behavior

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Common owned resource envelope

Every target resource has an immutable owned logical identity, account/project policy, desired configuration/hash/generation, operation/attempt history, provider/host/task binding, dependency/reservation references and redacted observed state/freshness. These are schema responsibilities, not new current wire fields. Mutations persist intent before backend calls; readiness and cleanup require actual evidence. [API](CoveStack-API-and-Compatibility.md), [schema](CoveStack-Database-Schema.md).

| Action | Required behavior |
| --- | --- |
| Create / update | Validate image/type/capability/owner/dependencies, admit capacity, persist idempotent intent; no hidden reassignment |
| Inspect | Return scoped redacted desired/observed generation, actual readiness/freshness and useful operation errors |
| Start / stop / restart | Only applicable types and authorized retained data; graceful/force policy explicit; re-admit capacity where needed |
| Resize | Backend/image-supported delta, safe capacity and explicit offline/online limits; no unapproved disk shrink |
| Snapshot / backup / restore | Declared consistency and ownership; independent protected data/key recovery and known-value verification |
| Delete / recover | Inspect owned provider/task, preserve shared/retained data, release reservations only after confirmed teardown |

## 2. VM Instances — Planned

| Requirement | Target behavior / evidence |
| --- | --- |
| User outcome | Real guest OS/kernel with allocated CPU/RAM, owned disk/vNIC/network and protected initial access |
| Backend | Incus management API preferred evaluation over QEMU/KVM; verify host/kernel/architecture/nested virtualization and image compatibility |
| Provision | Approve image provenance, reserve capacity/disks/addresses, persist task binding and inspect before create; deterministic logical identity |
| Readiness | Hypervisor accepted/running is not enough; verify guest/kernel plus required network/service readiness |
| Lifecycle | Start/stop/restart/resize/snapshot/backup/restore/delete only declared tested operations; preserve disk/retention on failure |
| Security | Trusted agent socket access only; no tenant host paths, unrestricted devices, privileged daemon or arbitrary host command |
| Failure | Lost provider response/parent death blocks conflicting action until owned task/object is inspected; one VM after retry |
| Tests | Test 1 plus actual guest/volume/NIC/admission/privilege/restore cases; foreign attach and cleanup denied |

Execution backend only (Incus/QEMU/KVM). Platform value to prove: account/environment ownership, durable generation/operation, capacity, safe recovery, scoped access and verified restoration. Do not call a container a VM or claim this already runs.

## 3. App Containers — Planned

Reviewed OCI image with explicit trust class, CPU/RAM, narrow capabilities, ports, network, secret references, volumes and health/logs. Docker is a candidate backend, not guest-kernel isolation. No daemon/Incus socket, privileged mode, broad devices or host-root mounts. Image update/readiness failure preserves owned persistent data and a usable rollback policy; unknown task inspects before retry. Test shared-kernel type distinction, privilege/foreign network denial, actual app health, restart and retained-volume cleanup.

## 4. Kubernetes — Planned

k3d CLI/k3s over Docker is initial evaluation. Pin a tested pair after host spike and verify every documented flag; no library/flag/version is invented. Reserve API endpoint/port before creation, track owned cluster/task/node bindings and export kubeconfig only on authorized request. Node topology/cluster quota come from measured profile and account share, not old academic limits. An actual API plus sample workload/network test is required; this is not a managed production Kubernetes SLA.

| CoveStack logic to implement | Backend mechanism / test boundary |
| --- | --- |
| Account/host/resource assignment and reserved labels | Proposed kan.account / kan.resource / kan.agent labels and deterministic names are backend correlation, not authorization; reserved-label contract still needs review |
| Desired revision/image/node spec and provider task | Persist before dispatch; inspect half-created cluster and settle/recreate only owned deterministic object |
| Port/capacity reservation and quota | Concurrent admission cannot allocate a duplicate port or overspend physical host |
| Scoped kubeconfig export | Foreign account 404; never list/get/observed secret material |
| Reconciled node changes / health | Backend commands follow current ownership/spec; stale worker cannot authoritatively report convergence |
| Deletion | Verify all owned containers/volumes/networks removed under retention; no unrelated cleanup |

Tests 7–10 cover crash, denial, cleanup and measured quota; add actual workload/readiness and node-change evidence.

## 5. Managed PostgreSQL — Planned

This name requires all of backup, restore, scoped credentials and private networking in acceptance. A reviewed workload backend uses persistent owned volume, resource limits and readiness through a real SQL connection. Distinguish platform DB from tenant DB. Generate narrow users/password references, redact ordinary DTOs, rotate/revoke safely, retain policy-governed volumes and test data after restart. Test 11 restores known rows/checksum and denies a foreign account's access. Snapshot alone is not a consistent database backup.

Execution backend only (reviewed PostgreSQL workload). Platform value: volume/credential/network ownership, data lifecycle and measured restore. A VM-backed profile may be needed for untrusted tenant workloads; a container-hosted database does not inherit VM isolation.

## 6. Object storage — RustFS preferred, integration Planned

Operator owns service process, persistent volumes, management path and protected administration. Java reconciles account-owned buckets, policy, credential references and intended deletion through a replaceable adapter. S3 requests authenticate narrow tenant credentials; root/admin never becomes a tenant key. Future SeaweedFS/other adapters must pass the same conformance.

Official RustFS docs describe tested common bucket CRUD/policy and object/multipart operations, IAM users/derived access keys. ACL authorization is unsupported; bucket logging/ownership controls are planned upstream; selected versioning/encryption behavior is not complete S3 parity. Full evidence/citations: [provider evaluation](../docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance).

Tests 12–13 require actual upload/download/checksums, denied foreign access, key rotation/revocation, restart-persistent bytes and Java DB-driven configuration rebuild. Independently back up/restore required object data/metadata/policy/key state and recheck denial. Recreating an empty bucket does not restore lost data. Exact artifact/admin API, custody/rotation/private network and consistent service-wide backup need B+C+A review.

## 7. Networks, IPAM, NICs, firewall and volumes — Planned

These are first-class product ownership/capacity capabilities, not merely a Docker bridge. Subnet pools and range policy must prevent overlap with host/VPN/Incus/Docker routes; one coordinated allocator authority. Interface attachment validates both owners/host/locality/current generation. Default-deny ingress/egress must actually block packets; stored rules alone are not enforcement. Volumes have disk/capacity/retention and attach/restore policy. Tests include 10 parallel IPs in one network, 10 accounts allocating networks with no overlap, denied foreign attach/traffic and retained data after restart.

## 8. Images, secrets, quotas and observations

Images are approved immutable provenance-bound artifacts with versioned initialization; no arbitrary host imports. Secret references are scoped and recoverable with a mature custody provider under review. Quotas combine authoritative account limits and host reservations; do not double-store derived usage. Health is observed and freshness-aware, not inferred from desired state or old acknowledgement. Core logging/audit/backup are part of lifecycle, not optional polish.

## 9. Broader services and environments

The [catalogue](CoveStack-Service-Catalog.md) lists all core/future entries. [Environment orchestration](CoveStack-Stack-Templates.md) combines owned typed children and readiness without an arbitrary DAG engine. [Future dependencies](CoveStack-Extended-Roadmap.md) cover LB/ASG/publication/domains and other candidates; they are not automatically available from contract/schema examples. No service is Supported before domain evidence and G1–G4/profile qualification.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-resource-services) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
