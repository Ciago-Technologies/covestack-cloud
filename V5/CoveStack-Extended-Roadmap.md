# CoveStack V5 — Future Services and Dependency Register

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Scope and authority

This retained filename is a **future dependency register**, not a release-by-release calendar or an academic extended-service ceiling. Current [core catalogue](CoveStack-Service-Catalog.md#1-core-product-services-and-capabilities) includes image/network/security/storage/recovery/monitoring/CLI/environment foundations. Do not defer those baseline controls because a related richer feature appears below. Every future candidate is Deferred until a bounded owner/dependency/acceptance plan is approved.

## 2. Future product candidates

| ID | Future candidate | Target scope | Dependency | Reason for deferral / required evidence |
| --- | --- | --- | --- | --- |
| 23 | Load Balancers | TCP/HTTP pools, health and configuration reconcile | Owned targets/network/firewall and ingress/ports | Defer until target health is trustworthy; remove dead member and rebuild config; no management route exposure |
| 24 | Private DNS | Owned internal records/zones and discovery | Network identity/IPAM, resolver control, IAM | Defer managed DNS lifecycle; prove scoped resolve/update/delete and deny cross-account disclosure |
| 25 | DNS Management | Zones/records, verification and external-domain integration | Private/public DNS provider credentials, domain policy, audit | Defer external provider dependency; proof ownership and safe reassignment with no old-tenant serving |
| 26 | NAT / Egress Gateway | Address translation, controlled outbound policy and egress observations | Routing/firewall, capacity and host network privilege | Defer beyond private-network subset; actual authorized egress and denied path tests, no invented internet reachability |
| 27 | VPN Gateway | Secure private-network access, device/user auth, routes and key rotation | IAM/secrets, routing, firewall and supported VPN backend | Defer remote-access attack surface; lost/revoked key denied and routes removed |
| 28 | Application Gateway / Ingress | Explicit HTTP/HTTPS publication, reverse proxy, TLS and health | Verified hostname, owned app endpoint, firewall, certificate provider | Defer public edge; unregistered host denied, deleted app route removed and control ports private |
| 29 | Public IP / Port Exposure | Real reachable host/network path with explicit port/IP binding | Operator-authorized routable address, NAT/firewall and admission | Defer topology-specific publication; fake IP never routes; verify remote path and collision-free cleanup |
| 30 | File Storage | NFS/SMB/shared filesystem, ACL and quotas | Secure per-account storage/export identity, network policy and supported backend | Defer shared-filesystem attack surface; foreign mount denied and retained data survives restart |
| 31 | Advanced Snapshot Policies | Expanded consistency/retention and recovery-point automation | Foundation snapshots, backend hooks, capacity/retention | Defer policy complexity until restore proven; enforce ownership and retention under failure |
| 32 | Advanced Backup Policies | Richer schedules, remote copies/retention and recovery policy | Verified foundation backup, protected keys/targets and scheduler | Defer beyond bounded foundation; test missed runs, corrupt copy, authorized retention and known-data recovery |
| 33 | Managed Valkey | Owned cache auth/persistence/quotas/health | Approved runtime/network/secrets and persistence profile | Separate from platform nonce store; defer demand; validate scoped auth and declared expiry/restart behavior |
| 34 | Managed MySQL | Instance/users/private endpoint/storage/backup/restore | DB provider, image/license, secrets/storage/network | Defer engine breadth; real SQL/data/restore and foreign access tests |
| 35 | Managed MariaDB | Instance/user/storage/recovery lifecycle | Reviewed engine artifact and DB provider conformance | Defer engine breadth; no automatic MySQL equivalence; actual SQL/persistence/restore |
| 36 | MongoDB-compatible Service | Document engine lifecycle and scoped durable access | User demand, exact engine license/support and provider review | Defer demand/rights/compatibility uncertainty; test chosen client/auth/data/recovery subset |
| 37 | Private Container Registry | Repositories/credentials/tags/digests/pull permissions/retention | Harbor or lighter OCI registry evaluation, blob storage, IAM/TLS | Defer registry operation burden; authenticated push/pull by immutable digest and foreign denial |
| 38 | Artifact Registry | Versioned artifacts/packages and scoped access | Owned blob storage/integrity/provider and IAM | Defer package-protocol breadth; exact artifact bytes/digest and retained-version access |
| 39 | Advanced Application Deployment | Richer app update/rollout/readiness/rollback history | Environment state, immutable app versions and health | Defer beyond basic app image update; failed rollout restores prior declared healthy app without data loss |
| 40 | Infrastructure Templates | Versioned reusable definitions, validated variables and previews | Environment resource/dependency schemas and ownership | Defer beyond one fixed initial definition; preview matches actual bounded mutations; no arbitrary DAG executor |
| 41 | Deployment Automation | Integrate approved tools/jobs with deployment history/rollback | Existing CI/CD/tool identity, app API and artifact trust | Defer pipeline scope; do not build full CI/CD; test authorized integration and failure propagation |
| 42 | TLS / Certificate Management | Issuance/import/renewal and gateway binding | Mature certificate backend, DNS/ownership/secrets, trusted network | Defer external issuance complexity; validate actual TLS, expiry/renewal and foreign key denial |
| 43 | Advanced Key Management | Managed key lifecycle/rotation/access through mature backend | Core secret references, custody/restore and KMS integration | Defer custom cryptography; revoked/rotated key and recovery policy tested |
| 44 | Security Posture Checks | Misconfiguration/weak-policy checks and recommendations | Reliable inventory, versioned rules and IAM/audit | Defer unvalidated compliance claims; findings evidence-based, exceptions explicit |
| 45 | Image Vulnerability Scanning | Expose trusted scanner findings for selected images | Image digest/provenance, scanner/version/feed and policy | Defer user-facing scanning service; baseline release scanning remains required; findings not proof absence |
| 46 | Supply-Chain Security Integrations | Richer provenance/signature/SBOM policy integrations | Core trusted artifact verification, approved signing/attestation backends | Defer integrations, not baseline release trust; reject tampered artifact and preserve provenance |
| 47 | Kubernetes Node Pools | Worker group sizing/constraints/node lifecycle and future scaling | Tested cluster lifecycle, capacity and node ownership | Defer topology expansion until node ownership, measured capacity and lifecycle conformance are established; historical academic node limits are not product limits |
| 48 | Autoscaling Groups | Bounded desired membership and scaling policy | Reliable compute/health/admission, cooldown and ownership model | Defer automated amplification; min/max/failure/outage tests without quota bypass |
| 49 | Multi-host Scheduling | Place account/environment resources across owned registered machines | Single-host admission plus network/storage locality, assignments and partition-safe task authority | Defer distributed failure domains; no implicit HA; reject stale hosts and prevent split ownership |
| 50 | Advanced Load Balancing | Richer protocols/algorithms/topologies and failover | Basic LB, measured health/capacity and gateway policy | Defer breadth; failover proven under partial failure with no cross-account target binding |
| 51 | GPU Instances | Approved device placement/isolation and quota | VM device/backend support, hardware profile and driver/license security | Defer scarce privileged hardware; no unsafe broad passthrough; isolation/reset/capacity tests |
| 52 | Cross-host Resource Migration | Reviewed move of owned compute/data/network attachment | Multi-host placement, backend migration, storage/network consistency and claims | Defer data-loss risk; interrupted move one authority, verified data and rollback/resume |
| 53 | Disaster Recovery | Coordinated recovery after major host/site/control-plane loss | Verified backups/keys/off-host targets, recovery mapping, consistency and operational ownership | Defer orchestrated multi-host/site guarantee; foundation restore remains core; rehearse stated recovery objectives |


## 3. Retained integrations and core-versus-future distinctions

| Domain | Core foundation | Later candidate / gate |
| --- | --- | --- |
| Monitoring | Fresh health/metrics/basic logs/events and account scope | Rich dashboards, history/prediction/tracing/alerts after trustworthy observations; E7 freshness required now, alert delivery when enabled |
| Snapshot/backup | Owned consistent recovery point and verified VM/SQL/object/key restore | Advanced schedules/replication/retention/DR; E3 plus named data restore |
| Metadata | Owned management API inspection and resource identity | Future guest self-metadata with reliable caller identity, no secrets; E4 |
| React UI | API/CLI authoritative state and user authorization | Read-only list first, then validated forms/operations; UI never authorization/DB authority |
| Ansible | Reviewed operator installer and API boundary | Future inventory/guest configuration with owned API JSON, key injection and guest prerequisites; E8 |
| Floci | Optional external app-testing environment integration | Exact emulator artifact/API/licensing/cleanup subset; not mandatory VM/runtime foundation or rewrite target |
| Publication/domain | Private workloads do not become public automatically | Shared operator gateway, explicit app route, verified domain/TLS/management denial; E5/E6 |

## 4. Domain-specific future design

| Area | Durable reconstruction / ownership | Owners and decisions |
| --- | --- | --- |
| ASG | Owned desired membership and capacity; missing child replaced under bounds/cooldown, no scale on lost authority | B coordinator, A runtime, C E1; min/max/cooldown/claim safety reviewed |
| Load balancer | Backend/health/route intent authoritative; derived config rebuilt and safe apply, no foreign target | D/B API/state, A future execution, C E2; protocol/tool/health/certificate decisions |
| Publication | Registered allowed target/hostname/port/generation from DB, not an arbitrary reverse proxy file | A runtime/B state/D API/C E5; shared tunnel/provider/config/TLS/SSRF policy |
| Domains | Current owner claim/proof generation; reassignment clears old authority | C verification/B state/D API; proof/expiry/provider and E6 |
| Guest metadata | Current workload/host/NIC binding and approved self payload only | C policy/A delivery/B+D API; route/token/spoof/NAT design, E4 |
| Inventory | Bob sees only Bob API resources; no docker ps/Incus/DB | Reviewed consumer owner; C E8; sshd/key injection and guest tooling prerequisites |

Future Go proxy/health/snapshot execution belongs to A only through reviewed dependencies; do not starve core safety/reviews or assume a backend command is an accepted service. Separate VM/volume snapshot foundation from advanced orchestration.

## 5. Universal future provider gates

G1: PostgreSQL desired state and generation are authoritative. G2: crash-safe inspection/reconciliation rebuilds state/config without duplicates. G3: account authorization denies foreign API resources (404 without existence leak) and actual traffic/storage boundaries are tested. G4: meaningful failure acceptance runs in CI on an authorized capable runner. **Acceptance Gates Passed** requires this evidence; **Supported** additionally needs a declared subset, qualified profile and maintenance policy.

Every provider declares actual operations/capabilities/consistency, secure identity, image/artifact rights, unknown-outcome handling, locality and cleanup. Real API/guest/data/packet/export/restore evidence is needed; no service becomes Supported because a table, contract kind, UI form or proposal exists.

## 6. What this does not authorize

No arbitrary host workflow runner, Labs/assignment/instructor product, full cloud parity, invented DNS/tunnel choice or supported multi-host/HA claim. No release dates or workload density figures. GPU/migration/DR need actual hardware/backend/network/storage/custody/demand/security review. Prior extension text is preserved in the snapshot rather than silently dropped; new scope follows canonical classifications.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-extended-roadmap) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
