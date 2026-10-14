# CoveStack V5 — Services, Capabilities and Release Boundaries

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Original resource services

**Planned offering is not implemented availability.** Each enabled service must pass its acceptance test and expose truthful status. First deliver Droplets and Kubernetes, then the dependencies of the complete stack.

| Service | Bounded release behavior | Limit |
|---|---|---|
| Droplet | Container-backed compute; create/inspect/start/stop/delete | Not an independent-kernel VM |
| Kubernetes | One reviewed version/shape; create/inspect/delete | k3d-backed local cluster |
| PostgreSQL | One supported engine line; persistence and connection details | No broad HA/database-engine promise |
| Object storage | MinIO S3 endpoint; owned buckets and scoped access | Selected S3 operations tested |
| HTTP load balancer | Registered backends, health and safe configuration reload | No generic TCP/UDP LB promise |
| Auto scaling group | Bounded replica controller, cooldown and limits | Initially approved compute members |
| Snapshots | Scoped immutable instance filesystem images | Mounted data/memory excluded |

## 2. Extensions agreed for V5

| Request | Feasible bounded version | Important limit |
|---|---|---|
| VPC | Owned private bridge per VPC on one host | Local network simulation |
| IPAM | Non-overlapping subnet leases; recorded actual IPs | One authoritative allocation policy |
| Internet website | Public hostname through tunnel and reverse proxy | Host, tunnel and app must run |
| Simulated public IP | Optional visibly simulated API/UI value | Not internet-routable |
| Direct IP browsing | Real gateway or routed private access | Fake address is insufficient |
| Custom domain | Verified owned name, DNS/TLS and route | Provider-specific setup required |
| Instance metadata | Read-only self-metadata | Identity/routing must be enforced |
| React SPA | Scoped resources, CLI help and real health | Backend still enforces policy |

| Additional capability | Required delivery |
|---|---|
| Stack templates | One-request dependency orchestration and resource references |
| Terraform | CoveStack provider for supported API resources |
| Ansible | Inventory/connectivity plus supported configuration workflows |
| Monitoring | Actual metrics/events and alert evidence |
| Recovery | Safe restart, bounded retry and verified restoration |
| UI/CLI parity | Same backend records, policy and operation results |

## 3. Publication is standard support, not automatic exposure

Every eligible Droplet/Kubernetes application can use the built-in publication workflow. A shared gateway is installed by the administrator. Users explicitly request publication for an approved service/port; no database, metadata or private management service becomes public merely because its resource exists.

The initial public endpoint is an HTTPS hostname. Arbitrary customer domains are a later gated extension. Kubernetes requests reach a Service/Ingress through the configured cluster ingress path, not a hardcoded Pod IP.

## 4. Resource information

| Resource | Normal metadata | Protected retrieval |
|---|---|---|
| Droplet | ID, image, state, VPC, private IP, URL | SSH access material under policy |
| Kubernetes | ID, version, readiness, nodes, endpoint | Kubeconfig/credentials |
| Database | ID, engine, host/port, storage, readiness | Password/connection secret |
| Bucket | ID, name, endpoint, policy summary | Scoped access credentials |
| VPC | ID, CIDR, gateway, attachments | No ordinary secret |
| Stack | ID, template version, child references, progress | Referenced secrets under authorization |

The control-plane metadata shown here is distinct from the endpoint inside a simulated instance. Avoid exposing a password just because it is a connection attribute.

## 5. Demonstrated release tiers

| Tier | Required outcome |
|---|---|
| Foundation | Real authenticated backend/agent and first compute flow |
| Initial services | Droplet and Kubernetes lifecycle with UI/CLI consistency |
| Network/publication | Owned VPC, address policy and external HTTPS app |
| Full template | App, PostgreSQL, S3 bucket, network and publication |
| Automation/operations | Provider, playbook, monitoring and restore evidence |
| Expanded services | LB, ASG and snapshots with explicit tested limits |

All agreed services remain targets. If a gate fails, report the affected feature unsupported/pending and revise workload with the mentor; do not publish a success-shaped placeholder to preserve a timetable.

## 6. Reuse and differentiation

Docker already offers IPAM/network primitives; Compose/Swarm support coordinated workloads; Terraform/Ansible can manage Docker. CoveStack adds its selected account/resource/stack contract and integrated workflow. It is not original merely because forms replace YAML. Evaluate whether the complete workflow helps the intended shared lab.

## 7. Optional future features

| Future feature | Reason to defer |
|---|---|
| Nested Docker | Privilege/storage/network complexity |
| VM-backed Droplets | New virtualization driver and operational model |
| Multi-host VPC | Routing, distributed allocation and fencing |
| Additional database engines | Engine-specific persistence/readiness tests |
| Full provider parity | Scope and semantics substantially exceed MVP |

Never mount the platform Docker socket into user workloads as a shortcut to nested Docker.
