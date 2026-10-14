# CoveStack V5 — Instance Metadata and Identity Boundaries

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Two meanings of metadata

| Meaning | Reader | Example |
|---|---|---|
| Control-plane resource information | Authorized UI/CLI/API user | ID, state, network and endpoint |
| Instance metadata service | Workload querying itself | Own ID, hostname, private IP and tags |

These are different services. The current catalog of accounts/tags/regions does not mean an IMDS implementation exists.

## 2. Provider contract

| Contract | Endpoint model | V5 handling |
|---|---|---|
| DigitalOcean | `/metadata/v1/` at link-local address | Document/test selected Droplet subset |
| AWS EC2 | `/latest/meta-data/` and IMDSv2 token flow | Separate future compatibility module |
| CoveStack custom fields | Explicit namespace | Real publication URL and simulator attributes |

Do not call generic JSON automatically AWS-compatible; many leaf metadata responses are text. DigitalOcean and AWS use different paths and semantics. [Droplet metadata](https://docs.digitalocean.com/products/droplets/how-to/access-metadata/), [AWS IMDSv2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html).

## 3. Bounded first response

| Field | Policy |
|---|---|
| Instance ID/hostname | Bound to this workload |
| Private address/VPC | Actual inspected assignment |
| Region/tags | Reviewed self-information subset |
| Public app URL | Separate custom field, if published |
| Simulated address | Clearly marked; not a working URL |
| User data | Excluded initially or separately reviewed for sensitivity |
| Signing/provider/database secrets | Never included |

## 4. Delivery stages

| Stage | Design | Acceptance |
|---|---|---|
| First | Per-instance read-only sidecar; explicit internal endpoint | Only its bound instance data |
| Compatibility | Namespace-local `169.254.169.254` listener | Correct routes/binding and provider response tests |
| Shared gateway alternative | Enforced identity/anti-spoofing | Cannot select another instance |

An HTTP listener does not automatically assign a virtual IP or intercept packets. Link-local support needs namespace/address/route setup. A controlled helper can configure the interface; the user workload should not receive host network administration privileges.

A per-instance sidecar sharing the instance network namespace can serve data bound to that instance. Assigning a link-local /32 inside separate namespaces avoids a single ambiguous host listener, but must be prototyped against actual Docker mode and host routes. Check that no genuine host/cloud metadata endpoint leaks into the workload.

## 5. Caller identity

| Untrusted input | Why insufficient |
|---|---|
| Requested instance/account ID | User can change it |
| `X-Forwarded-For` | Can be spoofed without a trusted proxy boundary |
| Source IP alone | NAT/spoofing/reuse can break identity |
| Token alone | Must be bound, scoped and validated under the chosen model |

The simplest MVP has no caller-selected target: the sidecar has only its assigned instance data. A token may mitigate some SSRF patterns but does not make an already authorized compromised workload unable to read its own metadata. Document that boundary.

## 6. Tests and operational policy

Deny public ingress to metadata, cross-instance lookup, forged identifiers, stale data after delete/reuse and access to secret fields. Bound request rates/response sizes. Metadata can use last approved local state during a control-plane outage, with freshness explicitly represented. Delete/recreate must not deliver the old instance's account data to a new owner.

No link-local compatibility claim is accepted until a real workload query succeeds with the exact expected response and negative identity tests pass.
