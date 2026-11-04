# Private networks, subnets, IPAM, interfaces and firewall policy

Updated: 4 November 2026. **Planned core capabilities 5–9**, not packet-level support. [Catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) is scope authority; [security](security.md) supplies trust requirements.

## Control model

One account owner with project/environment references; network/subnet/interface/IP/rule identities are separate conceptual resources. Initial academic constraint remains one private network/account and no tenant custom CIDR. Product subnet/range configuration is operator-reviewed from allowed pools; tenant custom ranges, cross-project sharing and eventual richer networks need an explicit scope/contract decision. No current VPC kind or new network state is invented.

| Core capability | Desired/observed and ownership responsibility | Failure/cleanup requirement |
| --- | --- | --- |
| Private network | Owned backend network plus host/domain policy; inspect membership and allowed sharing | Reject teardown with live/reserved attachments; unknown inventory never authorizes deletion |
| Subnet | Range/pool containment, reserved addresses, overlap with host/VPN/runtime networks | Keep lease through uncertain backend creation; transactional overlap tests |
| IPAM | One authority coordinating backend allocation and PostgreSQL reservations | Do not run competing Incus/Docker/CoveStack allocators; release only after confirmed detach |
| NIC | Owner/resource/host/network/address binding; actual VM/container attachment state | Reject foreign/wrong-host/double attach; recover partial detach without address reuse |
| Security group | Validated ingress/egress protocol/port/source/destination and association intent | Apply actual backend enforcement, default deny before readiness, retain guarded retry/error state |

Incus VM vNIC networking differs from Docker/k3d networking; a Docker bridge helper is not a universal network driver. Evaluate Incus/Linux ACL/firewall facilities and correct forwarding/bridge behavior for each runtime pair; pin/test actual backend semantics. Do not select nftables/iptables/OVN topology solely by naming it in docs. Link enforcement capability/version to observed installed policy and actual packets.

## Policy transaction and enforcement

Authenticate/authorize rule actor and network/resource ownership → validate typed protocol/range/destination and management-route invariants → persist generation/operation/policy references with admission → assigned agent claims future reviewed policy work → inspect owned installed policy → apply narrow backend changes → verify traffic → guarded observation/convergence. No host-shell command or arbitrary firewall fragment is accepted from tenant input.

Default deny inbound/outbound across VM-to-VM, container, cross-account and control-plane paths; explicit reviewed exceptions for app/SQL/S3/cluster/API/bootstrap/DNS. Return traffic behavior, IPv4/IPv6 and DNS/metadata access must be specified/tested; disabling IPv6 implicitly is not policy proof. Restrict backend/agent/API administration separately. Tenant-assigned addresses/labels cannot spoof owner identity; rule updates/removal need replay/fencing and audit as much as create.

Stored policy success is not enforcement: installed-rule readback plus actual permitted/forbidden connection tests are required. A rule change is not safely converged while stale broader policy remains. Rollback must not turn deny into allow-all; backend failure keeps new workload unready or quarantined under reviewed policy rather than temporarily exposed.

## Tests and future expansion

C owns parallel subnet/IP/port allocation, foreign attach, same-account allowed app/DB path, cross-account/management denial, source spoofing, explicit ingress/egress revoke, restart/drift, partial delete and route overlap tests. A implements runtime mapping; B owns durable reservations/transactions; D schema alignment. Tests must send real packets on authorized capable hosts; mock rule dictionaries do not pass.

Later private/public DNS, NAT, VPN, LB, ingress and public exposure are separate catalogue candidates. Core connectivity requires defined usable address/resolver behavior, not a full DNS-management service. A simulated public IP is metadata only. Public publication needs real reachable host/NAT/firewall path, hostname/TLS policy and explicit tenant/operator authorization; no guessed public address or automatic exposure.
