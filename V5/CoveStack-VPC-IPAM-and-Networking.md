# CoveStack V5 — VPC, IPAM and Connectivity Policy

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Three separate networking concerns

| Concern | Purpose |
|---|---|
| Private east-west | Permitted app/database/workload communication |
| Outbound egress | Workloads reaching approved outside destinations |
| Public ingress | Explicit published applications reached by visitors |

A VPC is not automatically an internet website. The initial implementation maps an owned VPC onto a user-defined bridge on one host. Shared-kernel containers remain a limited isolation model.

## 2. Address allocation

Example pools are illustrative and must be checked against actual routes. Choose a reviewed private pool, such as `10.80.0.0/16`, only when it does not overlap host/VM/LAN/VPN or existing Docker networks. Centrally allocate bounded subnets (for example /24) in the initial routing domain.

| Allocation | Authority | Storage |
|---|---|---|
| VPC subnet | CoveStack transactional allocator | Account/host/CIDR lease |
| Container IP — MVP | Docker IPAM | Inspect and record actual assignment |
| Explicit stable IP — later | One coordinated allocator | Reserved IP plus Docker acceptance |
| Gateway/system addresses | Deployment policy | Reserved before allocation |
| Public hostname | CoveStack domain/route registry | Unique verified mapping |

Do not run two independent workload-IP allocators. Docker already provides IPAM; using it under CoveStack's owned lease policy is legitimate integration. [Docker network options](https://docs.docker.com/reference/cli/docker/network/create/).

## 3. Lease lifecycle

| Step | Rule |
|---|---|
| Reserve | Atomic uniqueness/locking, not check-then-insert race |
| Provision | Create only the intended owned bridge |
| Record | Save runtime network ID and actual addressing |
| Attach | Validate resource/VPC account and host |
| Recover | Inspect/reconcile pending leases after failure |
| Release | Confirm attachments, ingress and metadata cleanup |
| Reuse | Only after owned objects have actually gone |

Track allocation exhaustion and explicit conflicting-CIDR errors. A partially created bridge/DB lease needs a reconciled failure state; deleting unrelated networks is never a valid recovery shortcut.

## 4. Isolation and firewall policy

| Path | Initial policy |
|---|---|
| Same VPC approved app → DB | Allowed under selected service policy |
| Different account/VPC | Denied unless explicitly modeled |
| Workload → control-plane DB/cache | Denied |
| Workload → Docker management API/socket | Denied |
| Visitor → published app | Approved gateway only |
| Visitor → metadata/private admin | Denied |

Bridge defaults and published host ports can create alternate paths. Test backend IPs, gateways, published ports and both directions. DNS lookup failure alone is not proof of isolation. Inspect the actual Docker iptables/nftables behavior; do not assume ordinary host INPUT rules govern forwarded container packets.

Kubernetes adds its own service/Pod CIDRs and network policy mechanisms. Reserve those ranges and verify cluster-to-VPC routes. A cluster network does not automatically become an isolated VPC attachment; the driver must define its boundary.

## 5. Address presentation

| Value | User meaning |
|---|---|
| Private IP | Reachable only through permitted private routes |
| Simulated public address | Explicit emulation value; not real internet reachability |
| Public HTTPS URL | Actually routed visitor endpoint |
| API endpoint/kubeconfig | Separately authorized administrative connectivity |

Documentation-only `203.0.113.42` is a fictitious example, not an address to allocate to internet services. DNS cannot make a fake address reachable. Use a real tunnel/gateway hostname for external visitors.

## 6. Future constraints

| Request | Additional requirement |
|---|---|
| Overlapping tenant ranges | Stronger namespace/routing isolation |
| Multi-host VPC | Overlay or explicitly managed routing and distributed leases |
| Per-resource public IP | Real address supply and routing; not API fields alone |
| Private remote admin | VPN/client routes and authentication |
| Hostile tenant isolation | Evaluate VM/sandbox boundaries, not bridge names |

## 7. Acceptance evidence

Run concurrent subnet allocation, exhausted-pool behavior, denied cross-account attach, allowed app/DB traffic, forbidden backend access, daemon-access denial, restart reconciliation and deletion/reuse. Record routes, runtime inspections, rule sets and versions with sanitized results.
