# CoveStack V5 — Service Gates and Product Promotion

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Promotion rule and statuses

G1: PostgreSQL desired state and generation are authoritative. G2: crash-safe inspection/reconciliation rebuilds state/config without duplicates. G3: account authorization denies foreign API resources (404 without existence leak) and actual traffic/storage boundaries are tested. G4: meaningful failure acceptance runs in CI on an authorized capable runner. **Acceptance Gates Passed** requires this evidence; **Supported** additionally needs a declared subset, qualified profile and maintenance policy.

Status vocabulary: Proposed, Planned, In Progress, Partially Implemented, Implemented, Runtime Verified, Acceptance Gates Passed, Supported, Deferred, Blocked. Source-confirmed libraries can be Implemented while service integration is Planned/Blocked. The catalogue's Partially Implemented interface/IAM entries describe component evidence, not full service support. No row below has passed runtime gates by documentation.

## 2. Core product gate register

| ID / capability | G1 authoritative records | G2 rebuild / recovery | G3 ownership / isolation | G4 required assertion | Owner / reviewers | Current category |
| --- | --- | --- | --- | --- | --- | --- |
| 1 Compute Instances (VMs) | Owned resource/VM/image/task generation and reservation | Inspect VM/task; reattach only owned disk/NIC; resume one guest | Foreign ID/disk/NIC/export denied | 1–6 + guest/readiness/privilege | A+B; C; D API | Planned |
| 2 App Containers | Distinct App Container type/spec and volume/network binding | Inspect owned container/task, image/health and retained data | Foreign network/volume denied; no privileged sockets | 1–6 container variants + privilege/update | A+B; C | Planned |
| 3 Kubernetes Clusters | Cluster desired revision/task/API port/capacity | Find partial owned cluster, settle task and verify actual API/workload | Foreign kubeconfig 404; no ordinary-view secret | 7–10 | A+B; C; D | Planned |
| 4 VM Images and Templates | Approved immutable image/version/provenance and initialization policy | Inspect exact digest and owned image/task cache; reject tampering | Foreign/private image/bootstrap access denied | Named image/provenance/boot failure | A+B; C | Planned |
| 5 VPC / Private Networks | Owned network/pool/generation and references | Rebuild owned network mapping/lease policy, preserve unrelated routes | Foreign network inspect/attach and packets denied | 14 + traffic/restart | C+B+A | Planned |
| 6 Subnets | Owned range/pool containment/nonoverlap reservation | Inspect actual subnet/backend mapping before apply/release | Foreign range attach/lookup denied | 14 parallel ranges | C+B+A | Planned |
| 7 IPAM | Unique live address/resource/interface leases | Rebuild intended allocations and inspect actual attachments | Foreign lease/resource denied; spoofing tested | 14 parallel IP/reuse | C+B+A | Planned |
| 8 Network Interfaces | Owned NIC/resource/host/network/current revision | Inspect attachment/task; resume one mapping and preserve reservations | Foreign attach/detach denied | Named interrupted NIC/traffic test | A+B; C | Planned |
| 9 Firewalls / Security Groups | Owned policy/rules/association/generation | Rebuild actual default-deny enforcement from DB intent | Forbidden ingress/egress/management packets fail | Named packet/restart/apply-failure | A+C+B | Planned |
| 10 Block Storage / Volumes | Owned disk/attachment/capacity/retention and provider ref | Inspect disk/task and reattach safely; no double attach | Foreign disk/backup/restore denied | Named volume persistence/restore/failure | A+B; C | Planned |
| 11 Object Storage | Account bucket/policy/key intent and generation | Java recreates intended config on RustFS; persistent bytes recovered independently | A key cannot access B bucket/admin | 12–13 + isolated byte restore | B Java; A operator; C | Planned |
| 12 Managed PostgreSQL | Owned DB/volume/private secret/backup references | Inspect workload/volume and restore known rows under current policy | Foreign SQL reachability/login denied | 11 | A+B; C | Planned |
| 13 Application Environments | Owned definition/version/child/dependency/reservations | Rebuild child progress; resume stable IDs after partial failure | Foreign/shared-resource deletion denied | 15 | B+A; C; D/E API | Planned |
| 14 IAM and RBAC | Accounts/projects/memberships/roles/revocation authority | Rebuild permissions/derived caches from authoritative records | Foreign/project/role downgrade denied | 3 + named membership race | C+B; D | Partially Implemented |
| 15 API Keys and Service Accounts | Hash tokens/protected signing material/assignment and revocation | Recover protected identity/custody without fallback key | User/service confusion, wrong-host/key denied | Named real auth/replay/rotate/outage | C+B; A | Partially Implemented |
| 16 Secrets Management | Protected secret references/policy/key version/rotation | Restore authorized keys/secret metadata through reviewed custody | Foreign retrieval/log disclosure denied | Named rotate/revoke/key recovery | C+B; A | Planned |
| 17 Audit Logs and Events | Append-only correlated audit/events with tenant policy | Rebuild views from authoritative events; no fabricated trusted account | Foreign read; app UPDATE/DELETE blocked | Named audit/unauth event test | C+B; A/D inputs | Planned |
| 18 Resource Quotas | Quotas plus locked host/resource reservations | Rebuild accounting from authoritative allocations + scoped inventory | Cross-account bypass or over-allocation denied | 10 + named parallel admission | B+A; C | Planned |
| 19 Monitoring and Resource Health | Timestamped ordered observation/probe policy | Rebuild actual freshness, never desired state as health | Foreign metrics/probe data denied | E7 + named stale/ordering | A+B; C; D/E | Planned |
| 20 Basic Logging | Scoped correlated sanitized logs/retention policy | Rebuild diagnostic views; errors never erase intent | Foreign logs and secrets denied | Named redaction/log outage | A+B; C | Planned |
| 21 Snapshots and Backup/Restore Foundations | Owned recovery point/artifact/checksum/key/retention | Restore known VM/SQL/object bytes, not just metadata | Foreign backup/restore target denied | E3 + 11 + named S3/control DB restore | A+B; C | Planned |
| 22 CLI and API | API-owned state/operations; no client resource DB | Client reconnect/read through actual API; finite wait/error/idempotency | User credentials and scoped views/export | 3–4 + named JSON/wait/cancel | E CLI; B/D API; A review; C | Partially Implemented |


## 3. Future candidate gate register

Each future item needs G1 durable owned intent/reservations, G2 safe provider/config/task reconstruction, G3 actual domain access denial, G4 meaningful CI failure test. The table adds domain-specific target; no unassigned test number is invented.

| Candidate | Rebuild / failure focus | Tenant boundary / test focus | Status |
| --- | --- | --- | --- |
| 23 Load Balancers | Owned targets/network/firewall and ingress/ports | Defer until target health is trustworthy; remove dead member and rebuild config; no management route exposure | Deferred |
| 24 Private DNS | Network identity/IPAM, resolver control, IAM | Defer managed DNS lifecycle; prove scoped resolve/update/delete and deny cross-account disclosure | Deferred |
| 25 DNS Management | Private/public DNS provider credentials, domain policy, audit | Defer external provider dependency; proof ownership and safe reassignment with no old-tenant serving | Deferred |
| 26 NAT / Egress Gateway | Routing/firewall, capacity and host network privilege | Defer beyond private-network subset; actual authorized egress and denied path tests, no invented internet reachability | Deferred |
| 27 VPN Gateway | IAM/secrets, routing, firewall and supported VPN backend | Defer remote-access attack surface; lost/revoked key denied and routes removed | Deferred |
| 28 Application Gateway / Ingress | Verified hostname, owned app endpoint, firewall, certificate provider | Defer public edge; unregistered host denied, deleted app route removed and control ports private | Deferred |
| 29 Public IP / Port Exposure | Operator-authorized routable address, NAT/firewall and admission | Defer topology-specific publication; fake IP never routes; verify remote path and collision-free cleanup | Deferred |
| 30 File Storage | Secure per-account storage/export identity, network policy and supported backend | Defer shared-filesystem attack surface; foreign mount denied and retained data survives restart | Deferred |
| 31 Advanced Snapshot Policies | Foundation snapshots, backend hooks, capacity/retention | Defer policy complexity until restore proven; enforce ownership and retention under failure | Deferred |
| 32 Advanced Backup Policies | Verified foundation backup, protected keys/targets and scheduler | Defer beyond bounded foundation; test missed runs, corrupt copy, authorized retention and known-data recovery | Deferred |
| 33 Managed Valkey | Approved runtime/network/secrets and persistence profile | Separate from platform nonce store; defer demand; validate scoped auth and declared expiry/restart behavior | Deferred |
| 34 Managed MySQL | DB provider, image/license, secrets/storage/network | Defer engine breadth; real SQL/data/restore and foreign access tests | Deferred |
| 35 Managed MariaDB | Reviewed engine artifact and DB provider conformance | Defer engine breadth; no automatic MySQL equivalence; actual SQL/persistence/restore | Deferred |
| 36 MongoDB-compatible Service | User demand, exact engine license/support and provider review | Defer demand/rights/compatibility uncertainty; test chosen client/auth/data/recovery subset | Deferred |
| 37 Private Container Registry | Harbor or lighter OCI registry evaluation, blob storage, IAM/TLS | Defer registry operation burden; authenticated push/pull by immutable digest and foreign denial | Deferred |
| 38 Artifact Registry | Owned blob storage/integrity/provider and IAM | Defer package-protocol breadth; exact artifact bytes/digest and retained-version access | Deferred |
| 39 Advanced Application Deployment | Environment state, immutable app versions and health | Defer beyond basic app image update; failed rollout restores prior declared healthy app without data loss | Deferred |
| 40 Infrastructure Templates | Environment resource/dependency schemas and ownership | Defer beyond one fixed initial definition; preview matches actual bounded mutations; no arbitrary DAG executor | Deferred |
| 41 Deployment Automation | Existing CI/CD/tool identity, app API and artifact trust | Defer pipeline scope; do not build full CI/CD; test authorized integration and failure propagation | Deferred |
| 42 TLS / Certificate Management | Mature certificate backend, DNS/ownership/secrets, trusted network | Defer external issuance complexity; validate actual TLS, expiry/renewal and foreign key denial | Deferred |
| 43 Advanced Key Management | Core secret references, custody/restore and KMS integration | Defer custom cryptography; revoked/rotated key and recovery policy tested | Deferred |
| 44 Security Posture Checks | Reliable inventory, versioned rules and IAM/audit | Defer unvalidated compliance claims; findings evidence-based, exceptions explicit | Deferred |
| 45 Image Vulnerability Scanning | Image digest/provenance, scanner/version/feed and policy | Defer user-facing scanning service; baseline release scanning remains required; findings not proof absence | Deferred |
| 46 Supply-Chain Security Integrations | Core trusted artifact verification, approved signing/attestation backends | Defer integrations, not baseline release trust; reject tampered artifact and preserve provenance | Deferred |
| 47 Kubernetes Node Pools | Tested cluster lifecycle, capacity and node ownership | Defer topology expansion until node ownership, measured capacity and lifecycle conformance are established; historical academic node limits are not product limits | Deferred |
| 48 Autoscaling Groups | Reliable compute/health/admission, cooldown and ownership model | Defer automated amplification; min/max/failure/outage tests without quota bypass | Deferred |
| 49 Multi-host Scheduling | Single-host admission plus network/storage locality, assignments and partition-safe task authority | Defer distributed failure domains; no implicit HA; reject stale hosts and prevent split ownership | Deferred |
| 50 Advanced Load Balancing | Basic LB, measured health/capacity and gateway policy | Defer breadth; failover proven under partial failure with no cross-account target binding | Deferred |
| 51 GPU Instances | VM device/backend support, hardware profile and driver/license security | Defer scarce privileged hardware; no unsafe broad passthrough; isolation/reset/capacity tests | Deferred |
| 52 Cross-host Resource Migration | Multi-host placement, backend migration, storage/network consistency and claims | Defer data-loss risk; interrupted move one authority, verified data and rollback/resume | Deferred |
| 53 Disaster Recovery | Verified backups/keys/off-host targets, recovery mapping, consistency and operational ownership | Defer orchestrated multi-host/site guarantee; foundation restore remains core; rehearse stated recovery objectives | Deferred |


Metadata, React UI and Ansible inventory remain separate future/interface integration discussions; they do not demote core monitoring or snapshot foundations. E1–E8 labels are preserved in [acceptance](CoveStack-AntiWrapper-Acceptance.md), but scope follows the current catalogue rather than the letter E.

## 4. Interface / automation gates

CLI/provider/UI/inventory do not each own a second G1 database. Prove API-only ownership, idempotent lifecycle, stable JSON, secret-safe output, finite wait/cancel/errors, no direct backend/DB and real user permission. Installer proves artifact/profile/secret safety, clean install, idempotent repeat and recoverable upgrade. A forms screen or terraform apply screenshot is not a service gate.

## 5. Review and honest relabelling

Before any readiness/release/profile review, record exact revision, environment/versions, test assertions/results, negative cases, reviewer and limitations. Missing runtime evidence keeps the accurate current category and a blocker, not Supported. Regression reopens the gate and prevents unsupported claims. Historical build/unit evidence remains dated; no fixed academic freeze date applies. Never write that all services are complete.

## 6. Evidence links

[Acceptance assertions](CoveStack-AntiWrapper-Acceptance.md), [test process](CoveStack-Testing-and-Acceptance.md), [implementation status](CoveStack-Implementation-Status.md), [schema](CoveStack-Database-Schema.md), [readiness profiles](../docs/platform/readiness.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-service-gates) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
