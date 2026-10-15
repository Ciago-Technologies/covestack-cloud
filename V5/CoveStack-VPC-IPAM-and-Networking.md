# CoveStack V5 — VPC, IPAM and Connectivity Policy

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. CORE-5 private networking boundary

**Status: Planned.** Execution backend only (Docker bridge). Platform value: allocator, ownership, collision-free subnets and IPs.

One private network per account on one host; no custom CIDR. VPC is server-owned, like storage.bucket, but has no kind/state in lifecycle contract 1.1.0. State is TBD with B; attaching compute.instance to this network requires coordinated schema/fixture/generated-output review B+A. No wire field is invented.

Private east-west traffic, reviewed outbound egress and public ingress are different concerns. Public ingress design is preserved in the [roadmap](CoveStack-Extended-Roadmap.md#5-application-publication); it is not implied by VPC creation. Containers share the host kernel.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. Address allocation

| Allocation                | Authority                                                           | Durable record                                           |
| ------------------------- | ------------------------------------------------------------------- | -------------------------------------------------------- |
| Account subnet            | C Java simple allocator; B transaction/constraints                  | subnet_pool and one vpcs row/account                     |
| Workload IP               | C allocator reserves DB lease; A helper applies reviewed assignment | ip_allocations; partial live-IP uniqueness               |
| Bridge/backend IPAM       | Docker enforces selected network/IP assignment                      | Inspect actual runtime and reconcile with DB reservation |
| Gateway/system exclusions | Reviewed pool policy B+A/C                                          | Explicit exclusions before workload allocation           |

Pool/subnet sizes are **TBD with B+A** after host/VM/VPN/Docker/Pod/Service-route overlap checks; no numeric pool is selected here. The old illustrative pool is preserved as unapproved historical text in the roadmap. Use one authoritative allocator: backend allocation cannot conflict with Java reservations. Explicit-IP assignment/route policy needs a tested decision before claiming the gate.

## 3. Lease lifecycle

| Step      | Rule                                                                                                |
| --------- | --------------------------------------------------------------------------------------------------- |
| Reserve   | Atomic uniqueness/locking, not check-then-insert race                                               |
| Provision | Create only the intended owned bridge                                                               |
| Record    | Save runtime network ID and actual addressing                                                       |
| Attach    | Validate resource/VPC account and host                                                              |
| Recover   | Inspect/reconcile pending leases after failure                                                      |
| Release   | Confirm owned attachments/runtime removed; future route/identity cleanup also required before reuse |
| Reuse     | Only after owned objects have actually gone                                                         |

Track allocation exhaustion and explicit conflicting-CIDR errors. VPC failure-state strings remain TBD with B. A partially created bridge/DB lease needs recorded retry/error intent; deleting unrelated networks is never a valid recovery shortcut.

## 4. Isolation and firewall policy

| Path                                    | Initial policy                                                         |
| --------------------------------------- | ---------------------------------------------------------------------- |
| Same VPC approved app → DB              | Allowed under selected service policy                                  |
| Different account/VPC                   | Denied unless explicitly modeled                                       |
| Workload → control-plane DB/cache       | Denied                                                                 |
| Workload → Docker management API/socket | Denied                                                                 |
| Public/private admin access             | No CORE-5 public route; private management stays denied from workloads |

Bridge defaults and published host ports can create alternate paths. Test backend IPs, gateways, published ports and both directions. DNS lookup failure alone is not proof of isolation. Inspect the actual Docker iptables/nftables behavior; do not assume ordinary host INPUT rules govern forwarded container packets.

Kubernetes adds its own service/Pod CIDRs and network policy mechanisms. Reserve those ranges and verify cluster-to-VPC routes. A cluster network does not automatically become an isolated VPC attachment; the driver must define its boundary.

## 5. Address presentation

| Value                          | Meaning                                                                                                                             |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------- |
| Private IP                     | Actual inspected assignment through permitted private routes                                                                        |
| API endpoint/kubeconfig        | Separately authorized private administrative connection                                                                             |
| Future public/simulated fields | Relocated to [publication design](CoveStack-Extended-Roadmap.md#5-application-publication); no fake address routes internet traffic |

Safe account views exclude credential content. Existing example addresses are not an approved allocation pool.

## 6. Future constraints — relocated

Overlapping tenant ranges, multi-host routing, real per-resource public addresses, private remote administration and stronger hostile-tenant boundaries remain Optional future work or extended publication decisions in the [roadmap](CoveStack-Extended-Roadmap.md). No extra CORE-5 deliverable is introduced.

## 7. Acceptance evidence

C owns test 14: ten parallel IP allocations in one VPC have no duplicate; ten accounts create VPCs concurrently with no overlapping CIDR; other-account network read is 404. Exercise allocator overlap/exhaustion internally, not a custom-CIDR user form. Verify allowed app/DB paths and denied control-plane/daemon/cross-account paths, restart bridge/lease reconciliation and safe release/reuse.

B owns SQL/constraints, C allocator/tests, A network execution helper inside the container driver. Record actual routes/rules/versions; DNS failure alone does not prove isolation.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
