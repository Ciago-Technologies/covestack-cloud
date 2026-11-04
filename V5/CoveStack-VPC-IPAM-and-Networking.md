# CoveStack V5 — Private Networks, IPAM and Traffic Enforcement

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Product networking scope

Core includes account-owned private networks, subnets, IPAM, NICs and firewall/security-group policy. No academic one-network/no-custom-CIDR rule constrains the product. The selected initial subset still needs reviewed pool/range/sharing policy and backend capabilities; unsupported requests fail explicitly. Incus/Linux/host networking and container/k3d paths differ; a Docker-only helper does not automatically attach VM guests.

## 2. Ownership and allocation model

| Record / capability | Authoritative responsibility | Acceptance |
| --- | --- | --- |
| Network/subnet | Account/project ownership, pool containment, overlapping-route rejection, immutable allocation identity | Parallel accounts allocate disjoint ranges; foreign inspection/deletion denied |
| Address lease | Owned subnet/interface/resource binding, live uniqueness and controlled release | 10 parallel IP allocations in one network have no duplicates |
| NIC/attachment | Validate both owners, current revision, host/backend/locality and selected policy | Foreign attach denied; interruption resumes with one binding |
| Firewall/policy | Desired rule/version, default deny and actual backend apply observation | Allowed traffic works and forbidden traffic fails before/after restart |
| Routes/DNS/egress | Scoped naming, resolver/route restrictions and private management separation | No cross-account disclosure, spoof bypass or management reachability |
| Capacity/recovery | Durable reservations, owned bridge/NIC references and complete inventory | Lost response does not free an in-use address or destroy an unrelated route |

Backend IPAM can be delegated under an explicit recorded mapping; do not let server and backend allocate the same pool independently. Pool values, reserved ranges/ports, tenant custom ranges and network-sharing rules require A+B+C decision, not invented CIDRs.

## 3. Allocation and attachment lifecycle

Validate request/ownership/pool → ordered DB locks and reserve unique range/address/port → persist intent/operation/attachment generation → apply reviewed owned network/interface/firewall capability → verify actual traffic and address → record observation. Detach/delete first disables unsafe access and checks dependencies; release only after confirmed teardown. Unknown task or partial inventory retains allocation and requires reconciliation.

Network leases are not worker execution leases. Worker expiry does not prove address cleanup. Tombstones/retention prevent accidental public-ID/resource reuse; released-address reuse requires actual old attachment absence. Lock ordering and serialization retries must be defined and tested by B+C.

## 4. Trust boundaries and enforcement

Separate management plane (API, control DB, Valkey, agent, Incus/Docker administration, RustFS admin) from tenant data paths. No tenant socket/root/device access; scoped user APIs cannot bridge to arbitrary host routes. Default-deny both ingress and egress with explicit required DNS/app/SQL/S3 destinations; do not assume private addressing alone isolates tenants. VM-to-container and cluster-to-workload traffic need their own tests.

NIC identity must resist address/source spoofing. Name/IP alone is not trusted metadata or tenancy identity. Policy changes are versioned; partial apply is not healthy/converged. Evaluate host firewall/Incus ACL mechanisms on the selected OS/profile and preserve operator-managed unrelated rules.

## 5. Current contract gap

No current VPC kind or complete network-attachment protocol exists in 1.1.0. Java owns network/IPAM intent; A's future agent capability executes only reviewed owned host operations. VM/NIC/subnet policy changes need B+D+A/C-approved schema/fixtures and generated outputs. Do not invent wire fields, new state values or label compatibility.

## 6. Internet access and address presentation

| Address/path | What it means |
| --- | --- |
| Private address | Reachable only through configured local/private routes or approved VPN |
| Simulated public address | Optional UI/demo annotation associated with future publication; never internet routing |
| Real gateway/IP | May be externally reachable only with actual routing/NAT/firewall/ISP permissions |
| Tunnel and reverse proxy | Future explicit registered app route through an operator gateway; workload and tunnel must run |
| Custom domain | Controlled verified DNS/TLS/route; no arbitrary global names without ownership |

Publication, DNS/NAT/VPN/real exposure are future services, not automatically enabled by network creation. No router DHCP/DNS control is assumed on a phone hotspot; plan reproducible local evidence and backup recordings. [Publication](CoveStack-Publication-Domains-and-Tunneling.md).

## 7. Acceptance and operations

Test 14: 10 IP allocations in one network, 10 accounts creating networks concurrently with no CIDR overlap, foreign network API denial. Add actual allow/deny traffic, source spoof/management protection, crashed attachment recovery, DB/provider unavailable behavior, release/reuse and shared dependency deletion. Restoring DB intent must inspect actual network state before recreation. Capacity and multi-host networking require separate measured/locality/partition design; no current multi-host support.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-vpc-ipam-and-networking) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
