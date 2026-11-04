# Future infrastructure candidates and dependencies

Updated: 4 November 2026. **No release-by-release plan, dates or Labs product.** The [canonical catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md#5-later-candidates-dependencies-and-deferral) owns entry IDs/classification; this index explains the dependency review. Historical numbered roadmap is superseded, not a new release commitment.

Core foundations include snapshots/backup, IAM/secrets, monitoring/health/logs, quotas, network enforcement and environments even while no service is supported. Richer future services must not become excuses to defer those protections. Product core is independent of academic CORE-5; dependencies and evidence, not historical academic limits, determine implementation priority.

| Future ID | Candidate | Prerequisites | Deferral reason / required proof |
| --- | --- | --- | --- |
| 23 | Load Balancers | Owned targets/network/firewall and ingress/ports | Defer until target health is trustworthy; remove dead member and rebuild config; no management route exposure |
| 24 | Private DNS | Network identity/IPAM, resolver control, IAM | Defer managed DNS lifecycle; prove scoped resolve/update/delete and deny cross-account disclosure |
| 25 | DNS Management | Private/public DNS provider credentials, domain policy, audit | Defer external provider dependency; proof ownership and safe reassignment with no old-tenant serving |
| 26 | NAT / Egress Gateway | Routing/firewall, capacity and host network privilege | Defer beyond private-network subset; actual authorized egress and denied path tests, no invented internet reachability |
| 27 | VPN Gateway | IAM/secrets, routing, firewall and supported VPN backend | Defer remote-access attack surface; lost/revoked key denied and routes removed |
| 28 | Application Gateway / Ingress | Verified hostname, owned app endpoint, firewall, certificate provider | Defer public edge; unregistered host denied, deleted app route removed and control ports private |
| 29 | Public IP / Port Exposure | Operator-authorized routable address, NAT/firewall and admission | Defer topology-specific publication; fake IP never routes; verify remote path and collision-free cleanup |
| 30 | File Storage | Secure per-account storage/export identity, network policy and supported backend | Defer shared-filesystem attack surface; foreign mount denied and retained data survives restart |
| 31 | Advanced Snapshot Policies | Foundation snapshots, backend hooks, capacity/retention | Defer policy complexity until restore proven; enforce ownership and retention under failure |
| 32 | Advanced Backup Policies | Verified foundation backup, protected keys/targets and scheduler | Defer beyond bounded foundation; test missed runs, corrupt copy, authorized retention and known-data recovery |
| 33 | Managed Valkey | Approved runtime/network/secrets and persistence profile | Separate from platform nonce store; defer demand; validate scoped auth and declared expiry/restart behavior |
| 34 | Managed MySQL | DB provider, image/license, secrets/storage/network | Defer engine breadth; real SQL/data/restore and foreign access tests |
| 35 | Managed MariaDB | Reviewed engine artifact and DB provider conformance | Defer engine breadth; no automatic MySQL equivalence; actual SQL/persistence/restore |
| 36 | MongoDB-compatible Service | User demand, exact engine license/support and provider review | Defer demand/rights/compatibility uncertainty; test chosen client/auth/data/recovery subset |
| 37 | Private Container Registry | Harbor or lighter OCI registry evaluation, blob storage, IAM/TLS | Defer registry operation burden; authenticated push/pull by immutable digest and foreign denial |
| 38 | Artifact Registry | Owned blob storage/integrity/provider and IAM | Defer package-protocol breadth; exact artifact bytes/digest and retained-version access |
| 39 | Advanced Application Deployment | Environment state, immutable app versions and health | Defer beyond basic app image update; failed rollout restores prior declared healthy app without data loss |
| 40 | Infrastructure Templates | Environment resource/dependency schemas and ownership | Defer beyond one fixed initial definition; preview matches actual bounded mutations; no arbitrary DAG executor |
| 41 | Deployment Automation | Existing CI/CD/tool identity, app API and artifact trust | Defer pipeline scope; do not build full CI/CD; test authorized integration and failure propagation |
| 42 | TLS / Certificate Management | Mature certificate backend, DNS/ownership/secrets, trusted network | Defer external issuance complexity; validate actual TLS, expiry/renewal and foreign key denial |
| 43 | Advanced Key Management | Core secret references, custody/restore and KMS integration | Defer custom cryptography; revoked/rotated key and recovery policy tested |
| 44 | Security Posture Checks | Reliable inventory, versioned rules and IAM/audit | Defer unvalidated compliance claims; findings evidence-based, exceptions explicit |
| 45 | Image Vulnerability Scanning | Image digest/provenance, scanner/version/feed and policy | Defer user-facing scanning service; baseline release scanning remains required; findings not proof absence |
| 46 | Supply-Chain Security Integrations | Core trusted artifact verification, approved signing/attestation backends | Defer integrations, not baseline release trust; reject tampered artifact and preserve provenance |
| 47 | Kubernetes Node Pools | Tested cluster lifecycle, capacity and node ownership | Defer topology expansion until node ownership, measured capacity and lifecycle conformance are established; historical academic node limits are not product limits |
| 48 | Autoscaling Groups | Reliable compute/health/admission, cooldown and ownership model | Defer automated amplification; min/max/failure/outage tests without quota bypass |
| 49 | Multi-host Scheduling | Single-host admission plus network/storage locality, assignments and partition-safe task authority | Defer distributed failure domains; no implicit HA; reject stale hosts and prevent split ownership |
| 50 | Advanced Load Balancing | Basic LB, measured health/capacity and gateway policy | Defer breadth; failover proven under partial failure with no cross-account target binding |
| 51 | GPU Instances | VM device/backend support, hardware profile and driver/license security | Defer scarce privileged hardware; no unsafe broad passthrough; isolation/reset/capacity tests |
| 52 | Cross-host Resource Migration | Multi-host placement, backend migration, storage/network consistency and claims | Defer data-loss risk; interrupted move one authority, verified data and rollback/resume |
| 53 | Disaster Recovery | Verified backups/keys/off-host targets, recovery mapping, consistency and operational ownership | Defer orchestrated multi-host/site guarantee; foundation restore remains core; rehearse stated recovery objectives |

Advanced observability/alerting and optional Floci are cross-cutting extensions associated with these foundations, not additional claimed core services. Optional emulator integration keeps its own artifact/license/enforcement/data/cleanup model; CoveStack remains infrastructure authority. No built-in CI/CD, learning/assessment workflow, release manager or general arbitrary DAG engine is introduced.

Review one concrete user need, dependency readiness, capacity, security boundary, ownership, failure/restore and consumer API tests before selecting a candidate for implementation. No demand or priority is established by this catalogue. Backend alternatives require [provider evaluation](provider-evaluation.md); no service is Supported before its [domain acceptance gates](readiness.md) pass.
