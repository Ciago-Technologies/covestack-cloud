# CoveStack V5 — Application Publication, Domains and Tunnels

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Default supported capability

**Publication support is built in for eligible Droplets and Kubernetes applications. Exposure is explicit.** One shared tunnel/gateway can carry many hostname routes; each published app receives a validated owned mapping. This is not one new tunnel per container.

```text
Internet browser
   -> real HTTPS hostname / public edge
   -> shared tunnel or VPS gateway
   -> trusted local reverse proxy
      -> Droplet application port
      -> Kubernetes ingress -> Service -> ready Pods
```

| Layer | Responsibility |
|---|---|
| Public DNS/TLS | Real hostname and trusted certificate |
| Tunnel | Reach the private lab without inbound router forwarding |
| Reverse proxy | Dispatch approved hostnames to approved upstreams |
| Application | Listen on expected interface/port and satisfy health criteria |
| Control plane | Account permission, domain claim and route lifecycle |

## 2. Request and route lifecycle

| Stage | Required behavior |
|---|---|
| Select resource | Must belong to the caller |
| Select service/port | Allowed app endpoint, not arbitrary upstream URL |
| Reserve hostname | Unique account/resource binding |
| Configure | Generate validated config safely; atomic apply/reload |
| Verify | Test endpoint and record actual readiness |
| Return | Public URL separate from private/simulated addresses |
| Stop/failure | Truthful unavailable status |
| Delete/unpublish | Remove routes before address/resource reuse |

The user should not paste Docker IPs into proxy configuration. Resolve upstreams from trusted resource state. Reject configuration injection, arbitrary private URLs and attempts to publish the control-plane DB, replay store, metadata or daemon.

## 3. Proxy topology choices

| Topology | Consequence |
|---|---|
| Proxy on host | Can use explicitly loopback-published workload ports |
| Proxy in container | Its localhost is not host localhost; requires approved connectivity |
| Proxy joined to tenant bridges | Increases cross-network reach; review compromise boundary |
| Kubernetes path | Stable Service/Ingress with mapped cluster ingress port |

Record the chosen method and measure isolation. Kubernetes Pod IPs can change. Cluster admin connectivity is separate from website ingress; do not publish admin APIs just to serve an app.

Cloudflare is one candidate for public hostname routing; a VPS is another. Do not require a specific vendor unless the team's chosen deployment does. [Tunnel routing](https://developers.cloudflare.com/tunnel/concepts/routing/), [k3d service exposure](https://k3d.io/stable/usage/exposing_services/).

## 4. Domain types

| Type | Initial handling |
|---|---|
| Team-owned subdomain | Automatically assigned within reviewed zone |
| Customer subdomain | Ownership challenge, DNS target and certificate validation |
| Apex domain | Provider-specific DNS/certificate method |
| Local-only name | Lab DNS; not public internet access |
| Fake public IP | Not a substitute for any domain route |

One certificate wildcard covers one label depth. Validate the actual `app-id.apps.domain` certificate rather than assuming `*.domain` covers it. Domain registrar and DNS provider can differ.

Ordinary tunnel setup for an owned zone is not automatic arbitrary-domain support. Use a supported custom-hostname product or a tested VPS certificate flow. Validate hostname ownership separately from certificate issuance. [Custom hostname validation](https://developers.cloudflare.com/cloudflare-for-platforms/cloudflare-for-saas/domain-support/hostname-validation/).

## 5. Security and operations

| Condition | Rule |
|---|---|
| Unknown Host/SNI | Reject/404; never default to another tenant |
| Unverified customer domain | No serving tenant traffic |
| Stale route after delete | Remove before reassignment |
| Provider token | Restricted operator secret; no browser exposure |
| Forwarded headers | Trust only known proxies; no spoofed identity |
| Internal signed API | Private route; preserve original signed authority |
| Tunnel outage | Show endpoint unavailable; resources/data remain |
| TLS expiry | Monitor and test renewal/reload |

## 6. First acceptance demo

Publish a Droplet app and a Kubernetes app under separate real HTTPS hostnames. Access both from mobile data or another unrelated network without VPN. Deny an unknown hostname and another account's publish request. Stop the app, restart the tunnel, delete the resource and verify route removal. Only real HTTP/HTTPS app access is promised; SSH/database/UDP publication requires separate protocol design.
