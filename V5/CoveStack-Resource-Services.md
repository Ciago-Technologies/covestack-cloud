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
| Ownership labels | Reserved and immutable to user inputs |

Use authoritative current contract vocabulary, not new enums copied from this document. Unsupported fields/actions are rejected explicitly. Failed work never acknowledges successful generation.

## 2. Droplets / compute instances

| Input | Bounded behavior |
|---|---|
| Image | Approved tested workload image |
| Size | Enforced CPU/memory/storage policy |
| Network | Owned VPC on selected execution host |
| SSH key | Validated owned public key; only if SSH mode supported |
| App port | Approved listener for optional publication |
| Persistent storage | Explicit volume identity and retention |

Inspect reserved labels and stable runtime identity before Ensure creates anything. Start/stop must respect user intent; an intentional stop is not a failure requiring replacement. Reboot requires an explicit operation or restart marker, not an active-to-active state change with no effect. Convergence includes agreed readiness, not only container existence.

Key-only SSH requires an SSH-ready image and tested minimal capability/transport configuration. A raw Ubuntu image does not automatically provide a usable Droplet session. Do not expose shared root passwords. Nested Docker is not included in the base specification, and no workload receives the platform daemon socket.

Acceptance: real create/inspect/stop/start/reboot/delete, quota enforcement, approved publication, denied other-account action, malicious input rejection and restart without duplicate containers.

## 3. Kubernetes clusters

| Area | Initial support |
|---|---|
| Version | One recorded tested k3d/k3s pair |
| Shape | Bounded server/agent node counts |
| Network | Explicit Docker, Pod and Service CIDRs without collisions |
| Administration | Authorized separately exported kubeconfig |
| Website | Service/Ingress through mapped cluster ingress |
| Lifecycle | Create/inspect/delete; other actions enabled after tests |

Cluster provisioning is actual local Kubernetes, not a renamed app container. Track provisioning, readiness, failure and safe cleanup using the shared lifecycle. Protect cluster administration from ordinary public website routing. Preserve the caller's existing kubeconfig on export/merge.

Reserve only necessary ports, not the whole NodePort range. Define the advertised API address and certificate coverage for the actual client path; a loopback API address is unusable on a remote laptop unless approved tunnel/VPN forwarding exists. Do not claim a website tunnel automatically makes kubectl connectivity work.

Acceptance: cluster ready, workload rollout, scoped kubeconfig retrieval, ingress reachable externally, wrong-account access denied, deliberate failure reported and delete removes owned nodes/ports without unrelated cleanup.

## 4. Managed PostgreSQL

| Area | Required policy |
|---|---|
| Engine | One reviewed tested PostgreSQL line |
| Data | Named owned volume and stable attachment |
| Credentials | Scoped generated secret and protected connection details |
| Network | User VPC/private route, separate from control-plane DB |
| Readiness | Authenticated usable connection under timeout |
| Delete | Explicit retain/delete policy |

The database used by the control plane is not a user-managed database. Restart/replacement must preserve data and account identity. Existing volumes have their initialized password; changing an environment variable is not a password rotation. Schema migrations for platform state differ from user database application migrations.

Acceptance: connect from approved app, deny forbidden/private-control-plane path, retain known table through replacement, rotate credentials by procedure and restore a database-consistent backup into another environment.

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

Acceptance: A can operate on A bucket; A cannot list/read/write/delete B bucket; saved objects survive recreation; endpoint and secret retrieval remain truthful.

## 6. HTTP load balancer

| Input | Validation |
|---|---|
| Backend resources | Owned, allowed app endpoints |
| Listener/rule | Initial bounded HTTP routing behavior |
| Health check | Approved path/protocol and timeouts |
| Public endpoint | Optional registered gateway publication |

CoveStack owns membership/health decisions; the proxy executes traffic forwarding. Generate valid complete include structures, test configuration before atomically reloading, and keep the last good config on failure. Reconcile changed IP/port assignments and removed members. Define empty-pool/unhealthy behavior rather than routing to a default unrelated app.

Acceptance: actual requests reach healthy members, failed/removed backend stops receiving new traffic, invalid reload preserves service, and cross-account/private-service targets are rejected.

## 7. Auto scaling groups

| Policy | Requirement |
|---|---|
| Minimum/desired/maximum | Ordered bounded values under account quota |
| Member template | Approved image/size/VPC and parent ownership |
| Scale signal | Actual selected metric or explicit desired-size request |
| Cooldown | Bounded rate of decisions and stable window |
| Pending capacity | Count reserved/provisioning members |
| Outage | Pause speculative replacement when agent information is stale |

The controller—not Docker's mere ability to run replicas—decides scaling. Serialize group changes; avoid excess launches during slow image pulls. Retire failed attempts and scale down under defined ownership/data policy. Do not delete persistent data as a hidden consequence of membership replacement.

Acceptance: manual/metric-driven supported change, min/max/quota enforcement, pending counts during slow provision, cooldown, lease outage and owned member cleanup. Record the tested metrics; do not present random chart values as scaling inputs.

## 8. Snapshots

| Snapshot contains | Snapshot does not contain |
|---|---|
| Selected container filesystem/image state | Mounted database/object volume data |
| Owned immutable image reference | VM memory/kernel state |
| Source/resource/account mapping | Automatically safe reusable credentials |

Create through the documented mapped action, record immutable identity, and restrict image access. Restored instances need their own identity/SSH host-key/credential policy; copying secrets across instances is not harmless. Snapshot retention and garbage collection must respect references and owners.

Acceptance: owned snapshot action completes honestly, restored instance is usable, wrong-account image use is denied, mounted data exclusion is explicit, and unrelated images are not pruned.

## 9. Capacity and fidelity

A full stack includes control-plane overhead, application workloads, DB/storage and several Kubernetes node containers. Record actual CPU/memory/disk use on the lab host; quotas must leave headroom for platform dependencies. A smaller tested supported cluster is better than an unmeasured large shape.

Do not promise cloud hardware isolation, SLA, full provider networking, provider billing or exact performance. Connection metadata and supported actions must describe what the local execution model can really do.
