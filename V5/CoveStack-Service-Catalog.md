# CoveStack V5 — Product Services, Capabilities and Boundaries

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Core product services and capabilities

This detailed catalogue mirrors the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md). Core is product architecture, not shipped support, a deadline or separate microservice count. A Partially Implemented supporting component does not establish a complete service gate. Actual resource/service runtime evidence remains pending.

| ID | Service / capability | Category | Required outcome / dependencies | Current evidence status |
| --- | --- | --- | --- | --- |
| 1 | Compute Instances (VMs) | Compute | Image, vCPU/RAM/disk/NIC, secure initialization; create/inspect/start/stop/restart/resize/snapshot/backup/restore/delete; depends on IAM, images, admission, volumes, NIC/network/firewall, agent/provider | Planned |
| 2 | App Containers | Compute | OCI image, ports, env/secret refs, CPU/RAM, volumes/network, health/logs, start/stop/restart/update/delete; depends on IAM, image verification, capacity, network/storage/secrets, OCI driver | Planned |
| 3 | Kubernetes Clusters | Compute | Create/delete/inspect/lifecycle, nodes/limits, reserved API endpoint, protected kubeconfig, health/network and standard-tool workload deployment; depends on Tested k3d/k3s/Docker host, ports/capacity/network, secret export | Planned |
| 4 | VM Images and Templates | Compute support | Approved OS images/custom-image intake, versioned VM templates, validated cloud-init or equivalent, provenance; depends on Trusted image origin/signatures, image/disk capacity, supported guest/backend | Planned |
| 5 | VPC / Private Networks | Networking | Owned network create/inspect/membership/attachment/dependency-safe delete; depends on Account/project policy, coordinated IPAM, NIC/firewall and backend network capability | Planned |
| 6 | Subnets | Networking | Operator-reviewed ranges, allocate subnet from permitted pool, reservations and no overlap; depends on Private network, route/pool validation, one authoritative allocator | Planned |
| 7 | IPAM | Networking | Allocate/reserve/inspect/release private addresses; interrupted-allocation recovery; depends on Owned subnet/host/NIC, transaction locks and backend allocator coordination | Planned |
| 8 | Network Interfaces | Networking | Owned interface create/inspect and validated attach/detach; address tracking; depends on VM/container capability, VPC/subnet/IPAM, host/owner binding | Planned |
| 9 | Firewalls / Security Groups | Networking/security | Default deny, ingress/egress protocol/port/source/destination rules, associations and audit; depends on NIC/network mapping, host policy backend, IAM, validated rule model | Planned |
| 10 | Block Storage / Volumes | Storage | Create/inspect/attach/detach, supported resize, ownership/usage and safe delete; depends on Pool capability/capacity, VM/container bindings, secret/encryption policy | Planned |
| 11 | Object Storage | Storage | S3 bytes/buckets, scoped keys/policies, quota, persistent objects, backup/restore; capability-declared versioning/lifecycle; depends on Maintained provider ADR, private service/admin path, capacity, key custody, backup adapter | Planned |
| 12 | Managed PostgreSQL | Database | Private database/user/credential lifecycle, real readiness, persistent storage, start/stop/health/backup/restore/delete/rotate; depends on Compute placement, volume/network, secrets, backup consistency | Planned |
| 13 | Application Environments | Application infrastructure | Owned resources and bounded dependencies, provision order/readiness/status/progress, safe retry/recovery/cleanup; depends on Constituent service gates, project ownership, durable child/attempt records, capacity | Planned |
| 14 | IAM and RBAC | Identity/security | Users, organization/account membership, roles/permissions, default-deny resource/project authorization; depends on Reviewed identity/membership schema and concrete auth stores | Partially Implemented |
| 15 | API Keys and Service Accounts | Identity/security | Scoped programmatic identity, issuance/revocation/rotation/expiration policy and trusted agent assignment; depends on IAM, protected key provider, hash/encryption distinction, replay protection | Partially Implemented |
| 16 | Secrets Management | Identity/security | Sensitive credential references, scoped retrieval/redaction/audit/rotation through evaluated mature backend; depends on Secret provider/custody ADR, bootstrap/recovery, IAM and TLS | Planned |
| 17 | Audit Logs and Events | Operations/security | API/admin/agent/resource/security/credential events, correlation, append-only app access; depends on Event schema, IAM, retention and protected logging storage | Planned |
| 18 | Resource Quotas | Operations/admission | CPU/RAM/disk/VM/container/cluster/storage/network per-account limits; depends on Atomic reservations, canonical host identity, IAM and actual capacity | Planned |
| 19 | Monitoring and Resource Health | Observability | Host/VM/container/Kubernetes/agent/scheduler CPU/RAM/disk/health and observation freshness; depends on Authenticated collection, timestamp/sequence semantics, real runtime probes | Planned |
| 20 | Basic Logging | Observability | Platform/agent/provisioning and authorized workload logs, sanitized errors and retention; depends on Collection scope, IAM, secret redaction, bounded storage | Planned |
| 21 | Snapshots and Backup/Restore Foundations | Storage/recovery | Owned VM/volume recovery points and metadata, bounded schedules/retention, verify/restore/history; depends on Backend consistency support, off-host storage/key recovery, capacity, IAM | Planned |
| 22 | CLI and API | User interface | Unified scoped lifecycle API and thin CLI, JSON, idempotency, wait/cancel/progress and explicit errors; depends on Existing signed transport plus real server/user auth, reviewed contracts/module layout | Partially Implemented |


## 2. Platform and backend boundaries

VM Instance is real general-purpose compute; App Container is distinct trusted shared-kernel compute. Incus/QEMU/KVM is preferred VM evaluation; k3d is a Kubernetes candidate. RustFS is preferred Planned object storage with Java-owned bucket/policy/credential reconcile and operator-owned process/volumes. Backend IPAM/scheduler features must be mapped, not duplicated as competing authority. Terraform/Ansible/UI/Floci integrations do not replace durable state/security.

## 3. Future candidates

Basic snapshots/backup, IAM/secrets, quotas, monitoring/health/logs and network enforcement are core foundations even when advanced policies/services are future. No “extended” filename demotes those requirements.

| ID | Future candidate | Dependencies / reason for deferral |
| --- | --- | --- |
| 23 | Load Balancers | Owned targets/network/firewall and ingress/ports; Defer until target health is trustworthy; remove dead member and rebuild config; no management route exposure |
| 24 | Private DNS | Network identity/IPAM, resolver control, IAM; Defer managed DNS lifecycle; prove scoped resolve/update/delete and deny cross-account disclosure |
| 25 | DNS Management | Private/public DNS provider credentials, domain policy, audit; Defer external provider dependency; proof ownership and safe reassignment with no old-tenant serving |
| 26 | NAT / Egress Gateway | Routing/firewall, capacity and host network privilege; Defer beyond private-network subset; actual authorized egress and denied path tests, no invented internet reachability |
| 27 | VPN Gateway | IAM/secrets, routing, firewall and supported VPN backend; Defer remote-access attack surface; lost/revoked key denied and routes removed |
| 28 | Application Gateway / Ingress | Verified hostname, owned app endpoint, firewall, certificate provider; Defer public edge; unregistered host denied, deleted app route removed and control ports private |
| 29 | Public IP / Port Exposure | Operator-authorized routable address, NAT/firewall and admission; Defer topology-specific publication; fake IP never routes; verify remote path and collision-free cleanup |
| 30 | File Storage | Secure per-account storage/export identity, network policy and supported backend; Defer shared-filesystem attack surface; foreign mount denied and retained data survives restart |
| 31 | Advanced Snapshot Policies | Foundation snapshots, backend hooks, capacity/retention; Defer policy complexity until restore proven; enforce ownership and retention under failure |
| 32 | Advanced Backup Policies | Verified foundation backup, protected keys/targets and scheduler; Defer beyond bounded foundation; test missed runs, corrupt copy, authorized retention and known-data recovery |
| 33 | Managed Valkey | Approved runtime/network/secrets and persistence profile; Separate from platform nonce store; defer demand; validate scoped auth and declared expiry/restart behavior |
| 34 | Managed MySQL | DB provider, image/license, secrets/storage/network; Defer engine breadth; real SQL/data/restore and foreign access tests |
| 35 | Managed MariaDB | Reviewed engine artifact and DB provider conformance; Defer engine breadth; no automatic MySQL equivalence; actual SQL/persistence/restore |
| 36 | MongoDB-compatible Service | User demand, exact engine license/support and provider review; Defer demand/rights/compatibility uncertainty; test chosen client/auth/data/recovery subset |
| 37 | Private Container Registry | Harbor or lighter OCI registry evaluation, blob storage, IAM/TLS; Defer registry operation burden; authenticated push/pull by immutable digest and foreign denial |
| 38 | Artifact Registry | Owned blob storage/integrity/provider and IAM; Defer package-protocol breadth; exact artifact bytes/digest and retained-version access |
| 39 | Advanced Application Deployment | Environment state, immutable app versions and health; Defer beyond basic app image update; failed rollout restores prior declared healthy app without data loss |
| 40 | Infrastructure Templates | Environment resource/dependency schemas and ownership; Defer beyond one fixed initial definition; preview matches actual bounded mutations; no arbitrary DAG executor |
| 41 | Deployment Automation | Existing CI/CD/tool identity, app API and artifact trust; Defer pipeline scope; do not build full CI/CD; test authorized integration and failure propagation |
| 42 | TLS / Certificate Management | Mature certificate backend, DNS/ownership/secrets, trusted network; Defer external issuance complexity; validate actual TLS, expiry/renewal and foreign key denial |
| 43 | Advanced Key Management | Core secret references, custody/restore and KMS integration; Defer custom cryptography; revoked/rotated key and recovery policy tested |
| 44 | Security Posture Checks | Reliable inventory, versioned rules and IAM/audit; Defer unvalidated compliance claims; findings evidence-based, exceptions explicit |
| 45 | Image Vulnerability Scanning | Image digest/provenance, scanner/version/feed and policy; Defer user-facing scanning service; baseline release scanning remains required; findings not proof absence |
| 46 | Supply-Chain Security Integrations | Core trusted artifact verification, approved signing/attestation backends; Defer integrations, not baseline release trust; reject tampered artifact and preserve provenance |
| 47 | Kubernetes Node Pools | Tested cluster lifecycle, capacity and node ownership; Defer topology expansion until node ownership, measured capacity and lifecycle conformance are established; historical academic node limits are not product limits |
| 48 | Autoscaling Groups | Reliable compute/health/admission, cooldown and ownership model; Defer automated amplification; min/max/failure/outage tests without quota bypass |
| 49 | Multi-host Scheduling | Single-host admission plus network/storage locality, assignments and partition-safe task authority; Defer distributed failure domains; no implicit HA; reject stale hosts and prevent split ownership |
| 50 | Advanced Load Balancing | Basic LB, measured health/capacity and gateway policy; Defer breadth; failover proven under partial failure with no cross-account target binding |
| 51 | GPU Instances | VM device/backend support, hardware profile and driver/license security; Defer scarce privileged hardware; no unsafe broad passthrough; isolation/reset/capacity tests |
| 52 | Cross-host Resource Migration | Multi-host placement, backend migration, storage/network consistency and claims; Defer data-loss risk; interrupted move one authority, verified data and rollback/resume |
| 53 | Disaster Recovery | Verified backups/keys/off-host targets, recovery mapping, consistency and operational ownership; Defer orchestrated multi-host/site guarantee; foundation restore remains core; rehearse stated recovery objectives |


## 4. Unnumbered optional integrations

Optional external Floci is an application-testing integration, not mandatory core, a hypervisor or a rewrite target. React UI, Ansible inventory, metadata and richer monitoring/publication link to their appropriate domain/future requirements; do not inflate the resource catalogue with every interface. No full AWS/DO/provider parity or unique market advantage is established.

## 5. Common resource information

Persist immutable owned identity, account/project references, desired configuration/hash/generation, operation/attempt/provider task, capacity/dependencies and redacted observed freshness. ID kinds, roles and secret export policy need reviewed schema. Actual current wire kinds/strings remain 1.1.0; these product names are not new fields. [API](CoveStack-API-and-Compatibility.md), [schema](CoveStack-Database-Schema.md).

## 6. Acceptance and use

G1: PostgreSQL desired state and generation are authoritative. G2: crash-safe inspection/reconciliation rebuilds state/config without duplicates. G3: account authorization denies foreign API resources (404 without existence leak) and actual traffic/storage boundaries are tested. G4: meaningful failure acceptance runs in CI on an authorized capable runner. **Acceptance Gates Passed** requires this evidence; **Supported** additionally needs a declared subset, qualified profile and maintenance policy.

Each service needs its own protocol/guest/data/traffic/readiness/restore test in addition to common lifecycle. Read [resource behavior](CoveStack-Resource-Services.md), [service gates](CoveStack-Service-Gates.md), [tests](CoveStack-Testing-and-Acceptance.md) and [future dependencies](CoveStack-Extended-Roadmap.md). No service count substitutes for an actual usable environment.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-service-catalog) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
