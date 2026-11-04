# Provider architecture and backend evaluation

Reviewed: 4 November 2026. **No final object-storage/secrets/registry provider selected, no driver implemented by this task.** [Catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) owns scope; this document owns backend evidence and selection criteria.

## Provider domains

| Domain | Direction to evaluate | Alternatives / selection boundary |
| --- | --- | --- |
| VM Instance | Incus management API with QEMU/KVM on supported Linux | Other hypervisors later; no public VM spec tied to one daemon API |
| App Container | Docker/approved OCI runtime | Other approved runtimes; Incus system container is not automatically the same app-container contract |
| Kubernetes | k3d/k3s with Docker runtime | Other distributions later; no managed production Kubernetes claim |
| Block volumes | Incus/host storage capabilities | Additional pools/backends with enforced capacity, ownership and restore semantics |
| Networking | Incus/Linux integration and explicit container/k3d paths | Coordinated allocator/ACL/routing backend, one authoritative allocation model |
| Object storage | Evaluate Garage, SeaweedFS and Ceph RGW | MinIO legacy choice under maintenance/distribution review, not new default |
| Database | Approved PostgreSQL workload backend | MySQL/MariaDB/cache/document engines only after demand/licensing tests |
| Registry | Harbor versus lighter OCI registry | Future candidate; do not build registry protocols from scratch |
| Secrets | Mature provider integration, e.g. OpenBao evaluation | Secret service availability, bootstrap/key recovery/rotation and license reviewed; no homemade vault |
| AWS-like app testing | Optional external Floci | Not authoritative infrastructure provider, mandatory dependency or Java rewrite target |

Existing Java control plane owns account/environment intent/transactions. Go agent/provider adapters own bounded local execution. New product-owned Go scheduler/adapters are future boundary decisions. No raw backend IDs, privileged socket or arbitrary host command is tenant authority.

## Interface and conformance responsibilities

A provider declares tested type/spec/API versions, supported operations/limits, consistency/readiness, cancellation/timeout ambiguity, placement locality, privilege and secret needs. It returns owner-bound opaque task/object references and observations; the control plane remains authority for intent, reservations and permissions. This describes responsibilities, not function signatures, new fields or proto files.

Inspect before create; deterministic logical identity plus provider task correlation; async accepted is not ready; verify actual guest/service. Same conformance tests cover idempotent create, foreign-object rejection, lost response/crash, safe retry, scoped inventory, data retention, explicit unsupported options and confirmed cleanup. Add backend-specific packet/disk/guest/protocol tests. A standalone adapter is not a supported service until its consumer and API path pass gates.

Minimize runtime administration: no dynamic untrusted plugin or arbitrary shell hook. New built-in providers register capabilities only after conformance and ownership review. Multiple competing schedulers/IP allocators are forbidden; delegate to backend under explicit mapping and reconcile returned placement/allocation. [Incus clustering](https://linuxcontainers.org/incus/docs/main/clustering/) already supplies backend mechanisms; CoveStack's layer adds account/environment policy and durable reservations, not a reimplementation of Incus internals.

## Object-storage findings and tradeoffs

| Candidate | Official evidence reviewed | Tradeoff / validation still required | Decision |
| --- | --- | --- | --- |
| Legacy MinIO | Upstream repository is marked archived/read-only and identifies AGPL-3.0 | Community maintenance/security-fix path, exact source/binary distribution and license obligations must be reviewed; historical integration is not a supported default | Do not adopt as a new default without explicit ADR |
| Garage | Official site targets self-contained distributed S3 storage; quick-start documents scoped keys/buckets and a compatibility list | Some S3 ACL/policy operations are not implemented in that guide; test equivalent permission requirements, quota/versioning/lifecycle and backup; no assumption full AWS policy parity | Candidate; no supported profile or minimum RAM inferred |
| SeaweedFS | Upstream provides distributed object/file storage, S3 API docs and Apache-2.0 license; release history is published | Benchmark chosen topology and metadata/security/backup operations; feature catalogue does not prove our tenant policy or compatibility subset | Candidate; exact artifact/support pin pending |
| Ceph RGW | Official stable-version docs describe S3/Swift gateway over Ceph/RADOS with user-management | Requires Ceph storage operations; measure topology/resource/upgrade burden rather than assume small-host fit; test supported S3/policy subset and multi-tenant recovery | Candidate especially where Ceph already exists; no default selected |

Sources: [MinIO repository](https://github.com/minio/minio), [Garage](https://garagehq.deuxfleurs.fr/), [Garage quick-start/compatibility caveat](https://garagehq.deuxfleurs.fr/documentation/quick-start/), [SeaweedFS](https://github.com/seaweedfs/seaweedfs), [S3 API](https://github.com/seaweedfs/seaweedfs/wiki/Amazon-S3-API), [releases](https://github.com/seaweedfs/seaweedfs/releases), [Ceph RGW stable-version docs](https://docs.ceph.com/en/squid/radosgw/). Main/latest docs are not a chosen support pin.

The archive marker is a maintenance-risk signal, not proof every MinIO offering is abandoned. No alternative is declared actively security-maintained merely because a site loads or a release page exists. Before selection record dated release/security-policy/issue evidence and responsible maintainers for the exact artifact. Garage/Ceph artifact-specific license/distribution details need further review; no blanket compatibility legal conclusion.

## Required selection evidence

| Criterion | Gate before choosing default |
| --- | --- |
| Technical maturity/API stability | Documented versioned API and supported requested subset; unknown/unsupported operations rejected |
| Maintenance | Current upstream release/security advisory/support policy for selected artifact; upgrade path and ownership |
| Footprint | Measured RAM/CPU/disk/metadata/replication overhead under actual object sizes and concurrency; no invented minimum |
| Identity/policy | Scoped key issuance/revocation/rotation, cross-account denial and admin isolation; policies/quota enforced, not just accepted |
| Durability/recovery | Restart retained bytes, consistent backup, known-object restore, corruption rejection and key recovery |
| S3 capability | SDK/CLI exact-byte tests, declared versioning/lifecycle/multipart/policy/quota limitations; no blanket full S3 claim |
| Operations | Verified installation, TLS, audit/metrics, disk-full handling and data-safe upgrade/rollback |
| Rights/supply chain | Exact license/notices/source/binary acquisition, digest/signature/SBOM and vulnerabilities |

No benchmark, install or object operation ran. B+C/A select object-admin/custody/provider mapping; C owns denial/recovery tests. Preserve old MinIO-specific acceptance intent but map it to selected provider with team/mentor review rather than silently changing CORE-5.

## Registry and secrets candidates

[Harbor](https://goharbor.io/docs/) documents project/registry operations; compare its full operational footprint against a lighter OCI registry for the future scope. Validate authenticated push/pull/retention/scanning interface and distribution rights, not just successful container startup.

[OpenBao](https://openbao.org/docs/) documents secrets/auth/audit and CLI/API mechanisms; evaluate authorization, rotation, unseal/bootstrap, backup/key recovery and availability. Its presence would not by itself solve CoveStack secret isolation. No provider added/installed or final pin chosen. Keep library/config secrets out of request/observation/log responses regardless of provider.
