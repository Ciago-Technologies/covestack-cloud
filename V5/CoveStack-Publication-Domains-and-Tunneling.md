# CoveStack V5 — Application Publication, Domains and Gateways

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Future opt-in capability

A user may later publish an owned app running in a VM, trusted container or Kubernetes workload through an operator-approved gateway/reverse proxy and a real reachable HTTPS hostname. This is future scope, not already available from a create request. Support for publishing is distinct from exposing every resource by default. Management/DB/admin ports never become public because a workload exists.

## 2. Durable publication and domain lifecycle

Authenticate ownership → validate allowed workload port and actual readiness → verify hostname/domain control → persist route/domain binding and generation → render only approved DB-owned routes → validate/apply gateway config safely → verify actual HTTPS request → report scoped URL/freshness. Delete disables route and revokes domain binding before resource cleanup. Unknown gateway outcome inspects config/state, does not overwrite unrelated routes.

| Record / subsystem | Target responsibility |
| --- | --- |
| Publication | Account/resource/hostname/port policy, generation and intended removal |
| Domain verification | Proof bound to current claim generation and owner; previous proof invalid after reassignment |
| Gateway config | Derived cache from approved DB records, never a second authoritative route registry |
| Tunnel | Shared operator-managed reachable path; no mandatory separate tunnel per VM/pod |
| TLS | Controlled certificate lifecycle, secure keys/renewal and verified hostname |

Existing V5 publication and V6 domain SQL are partial files with foundation prerequisites; no API/reconciler/gateway is proven by them. Tool/provider choice and DNS challenge interface require reviewed artifact/policy decisions; this doc does not select a tunnel vendor or guarantee Nginx execution exists.

## 3. Routing and address facts

| Path | Requirement / limit |
| --- | --- |
| Internet hostname | Real DNS → reachable tunnel/VPS/gateway → registered reverse proxy → owned reachable workload |
| Private address | Requires appropriate private routing or approved VPN; not universally reachable |
| Simulated public IP | Optional demonstration field only; does not route internet traffic |
| Direct public IP | Real routable gateway/IP with host/NAT/firewall/ISP permission, not an arbitrary fake address |
| Kubernetes app | Correct service/ingress/workload port and network path; a pod starting is not an external URL |
| VM app | Service listens on permitted interface/port and survives policy/TLS path verification |

Host, workload, gateway and tunnel must all run. Phone hotspot may restrict inbound reachability; no router DNS/DHCP ownership is assumed. Do not promise creating arbitrary globally resolvable domains without domain control.

## 4. Gateway and identity safety

Reject unregistered hostnames; no default tenant backend or public management route. Validate target owner, host/address resolution, port and policy to prevent SSRF/management/metadata reachability. Preserve operator config and per-account boundaries; sanitized logs and quotas protect traffic/route exhaustion. Certificate/private keys are protected references, never list/get output.

## 5. Domains and verification

Platform subdomain allocation and verified custom domains have different authority. Only an owned/approved zone can issue a platform name. Custom domain proof/expiry/revalidation and reassignment must bind to the new owner/generation; stale proof cannot serve the old tenant. Define DNS/TLS control and cleanup before calling a domain verified. A documented SQL status is not proof a DNS check ran.

## 6. Owners, dependencies and acceptance

B durable publication/operations; D API mappings; C domain verification/security tests; A future gateway/tunnel config execution after safe runtime foundation; E API consumer. E5: unknown hostname 404, route disappears after delete and management port unreachable. E6: unverified domain serves nothing, reassignment invalidates old serving/proof. Add reload failure/config recreation from DB, actual HTTPS, cross-account target and tunnel outage tests. Public URL claims need real external reachability evidence, not screenshots alone.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-publication-domains-and-tunneling) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
