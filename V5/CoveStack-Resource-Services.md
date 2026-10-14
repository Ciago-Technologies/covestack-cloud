# CoveStack V5 — Resource Specifications and Service Behavior

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Common resource envelope

| Field/behavior | Requirement |
|---|---|
| ID, account and kind | Stable identity from trusted server context |
| Host/agent and lease | Authorized execution placement |
| Desired spec/generation | Validated versioned intent |
| Observed state/time | Actual evidence, not requested state echoed back |
| Operation | Durable create/update/delete attempt and result |
| Secret references | Separate protected retrieval |
| Ownership labels | kan.* reserved labels are proposed; approve contract before use; existing fixtures do not establish these labels |

Use authoritative current contract vocabulary, not new enums copied from this document. Unsupported fields/actions are rejected explicitly. Failed work never acknowledges successful generation. Contract 1.1.0 wins; unknown strings are TBD with B. Convergence ack requires exact owner/epoch/generation and clears both lease fields; progress never changes acked_generation.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. Droplets / compute instances

| Input | Bounded behavior |
|---|---|
| Image | Approved tested workload image |
| Size | Enforced CPU/memory/storage policy |
| Network | Owned VPC on selected execution host |
| SSH key | Validated owned public key; only if SSH mode supported |
| App port | Private application listener; public routing is a roadmap target |
| Persistent storage | Explicit volume identity and retention |

Inspect approved reserved labels and stable runtime identity before Ensure creates anything. Start/stop must respect user intent; an intentional stop is not a failure requiring replacement. Reboot requires an explicit operation or restart marker, not an active-to-active state change with no effect. Convergence includes agreed readiness, not only container existence.

Key-only SSH requires an SSH-ready image and tested minimal capability/transport configuration. A raw Ubuntu image does not automatically provide a usable Droplet session. Do not expose shared root passwords. Nested Docker is not included in the base specification, and no workload receives the platform daemon socket.

**Status: Planned.** Execution backend only (Docker). Platform value: tenancy, desired/observed generation, lease and fencing, crash-safe restart.

Droplets are container-backed on one host, never VMs. Tests 1–6 prove crash/fence/account/idempotency/generation/outage behavior. Reboot/restart intent needs coordinated schema/fixture review; no new wire marker is invented. Public routing is relocated to the roadmap.

## 3. Kubernetes clusters

| Area | Initial support |
|---|---|
| Version | One tested pinned k3d+k3s pair; verify every CLI flag |
| Shape | 1 server + max 2 agents; hard cap 2 clusters/host |
| Network | Explicit Docker, Pod and Service CIDRs without collisions |
| Administration | Authorized separately exported kubeconfig |
| Website | Private workload rollout; public routing details relocated to roadmap |
| Lifecycle | Create/inspect/delete; other actions enabled after tests |

Cluster provisioning is actual local Kubernetes, not a renamed app container. Track provisioning, readiness, failure and safe cleanup using the shared lifecycle. Protect cluster administration from ordinary public website routing. Preserve the caller's existing kubeconfig on export/merge.

Reserve only necessary ports, not the whole NodePort range. Define the advertised API address and certificate coverage for the actual client path; a loopback API address is unusable on a remote laptop unless approved tunnel/VPN forwarding exists. Do not claim a website tunnel automatically makes kubectl connectivity work.

**Status: Planned.** Execution backend only (k3d). Platform value: ownership, scoping, crash-safe lifecycle, quota, scoped kubeconfig.

Tests 7–10: exactly one cluster after crash; other-account kubeconfig 404 and absent from ordinary views; clean owned deletion; third cluster rejected. Reserve API port before create; release only after cleanup. Ensure finds half-create by approved labels and finishes or recreates under the same deterministic name. Node-pool changes go through reconcile within node limits, not independent manual commands. Driver calls k3d CLI, not library.

8 GB host; 0.5–1 GB/cluster is an **unverified estimate, confirm in K8s host spike**. One-day November spike covers AlmaLinux cgroup v2, SELinux, firewalld and RAM with one tested pair. [Main Kubernetes ownership table](DocsV5-DO_Team.md#52-kubernetes--planned-required) defines platform-owned logic.

## 4. Managed PostgreSQL

| Area | Required policy |
|---|---|
| Engine | One reviewed tested PostgreSQL line |
| Data | Named owned volume and stable attachment |
| Credentials | Scoped generated secret and protected connection details |
| Network | User VPC/private route, separate from control-plane DB |
| Readiness | Authenticated usable connection under timeout |
| Delete | Explicit retain/delete policy |
| Backup/restore | Measured known-row and checksum restore; scoped artifact |

The database used by the control plane is not a user-managed database. Restart/replacement must preserve data and account identity. Existing volumes have their initialized password; changing an environment variable is not a password rotation. Schema migrations for platform state differ from user database application migrations.

**Status: Planned.** Execution backend only (container). Platform value: volume ownership, backup/restore, credential scoping.

Generic container driver: image/env/volume/labels/health, not a third PostgreSQL driver. Test 11 requires known rows/checksum restored and cross-account private-network denial. Backup, restore, credential scoping and private networking are all required before the service may be Gates passed.

## 5. Object storage through MinIO

| Area | Required behavior |
|---|---|
| Bucket | Account-owned logical/physical mapping |
| S3 endpoint | Actual client-reachable supported protocol endpoint |
| Credential | Scoped policy; not global MinIO root keys |
| Operations | Selected put/get/list/delete tests |
| Persistence | Correct mounted data path and object verification |
| Network | Defined shared-service/private/public endpoint policy |

A bucket is a server-managed logical resource; it need not be one container or bridge. Prefixing its name avoids collisions but does not authorize access. Shared MinIO can be used if credentials/policies isolate accounts. Define approved app access to that shared endpoint rather than pretending each bucket is physically attached to a VPC.

MinIO administrative access stays private. Proxying S3 traffic must preserve the protocol's signed authority/path semantics. A generic website route can break client signatures. Test a real supported S3 client against the selected connection method.

**Status: Planned.** Execution backend only (MinIO). Platform value: per-account credentials, bucket scoping, cross-account denial.

MinIO process uses the generic container driver; bucket and credential reconciliation is Java server-side under storage.bucket, never agent bucket work. Tests 12–13 prove denied cross-account key and server recreation of bucket/policy after MinIO restart. Admin custody, rotation and server-to-MinIO private path are open B+C decisions.

## 6. HTTP load balancer — relocated

**Extended track.** All former design and acceptance text is preserved in the [roadmap](CoveStack-Extended-Roadmap.md#4-load-balancing). Not an SRS requirement or deliverable.

## 7. Auto scaling groups — relocated

**Extended track.** All former design and acceptance text is preserved in the [roadmap](CoveStack-Extended-Roadmap.md#3-auto-scaling-groups). Not an SRS requirement or deliverable.

## 8. Snapshots — relocated

**Extended track.** All former design and acceptance text is preserved in the [roadmap](CoveStack-Extended-Roadmap.md#8-snapshots). Not an SRS requirement or deliverable.

## 9. Capacity and fidelity

A full stack includes control-plane overhead, application workloads, DB/storage and several Kubernetes node containers. Record actual CPU/memory/disk use on the lab host; quotas must leave headroom for platform dependencies. The required shape stays within the 2-cluster host cap and 1-server/max-2-agent limits.

Do not promise cloud hardware isolation, SLA, full provider networking, provider billing or exact performance. Connection metadata and supported actions must describe what the local execution model can really do.

### 9.1 VPC/IPAM — Planned

Execution backend only (Docker bridge). Platform value: allocator, ownership, collision-free subnets and IPs. Server-owned vpcs/ip_allocations records drive C's allocator and A's container-driver network helper. One private network/account, no custom CIDR. Test 14 includes ten parallel allocations in one VPC and ten accounts creating VPCs without overlap. No VPC wire kind exists; attachment is coordinated B+A contract work.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
