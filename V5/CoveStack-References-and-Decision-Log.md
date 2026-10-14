# CoveStack V5 — References, Assumptions and Decisions

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Document authority

| Material | Authority |
|---|---|
| V5 docs | New target architecture/schedule and agreed scope |
| Current runtime contracts | Implemented wire/generator source |
| Actual commits/tests/deployment | Feature completion evidence |
| Old team guides/archives | Historical design and findings |
| Mentor rubric | External acceptance requirement; obtain current version |

No official C-DAC marks allocation or no-coding rule is assumed. Final deadline is interpreted as 25 January **2027**, the January following the current October 2026 planning period.

## 2. Primary source links

| Reference | Purpose |
|---|---|
| [Docker bridge](https://docs.docker.com/engine/network/drivers/bridge/) | Local network primitives and limits |
| [Docker IPAM options](https://docs.docker.com/reference/cli/docker/network/create/) | Existing address mechanisms |
| [Docker security](https://docs.docker.com/engine/security/) | Daemon/workload authority |
| [Cloudflare routing](https://developers.cloudflare.com/tunnel/concepts/routing/) | Candidate public hostname transport |
| [Domain validation](https://developers.cloudflare.com/cloudflare-for-platforms/cloudflare-for-saas/domain-support/hostname-validation/) | Ownership/certificate distinction |
| [DigitalOcean metadata](https://docs.digitalocean.com/products/droplets/how-to/access-metadata/) | Selected provider endpoint semantics |
| [AWS metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html) | Separate IMDSv2 behavior |
| [k3d exposure](https://k3d.io/stable/usage/exposing_services/) | Ingress integration |
| [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) | Provider development |
| [Ansible Docker](https://docs.ansible.com/projects/ansible/latest/collections/community/docker/index.html) | Existing automation capabilities |
| [OWASP object permissions](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/) | Authorization test rationale |

These mechanisms were researched during the preceding design work; k3d, Terraform and Ansible references were checked again while preparing V5. Exact tool/provider compatibility must be retested against release versions. Do not claim Docker lacks IPAM or Terraform/Ansible integrations.

## 3. Decisions to record in G0

| Decision | Required information |
|---|---|
| First stack | Supported app interface and dependencies |
| Network pool | Host/VM/VPN overlap evidence |
| Proxy/tunnel | Namespace placement, DNS/TLS and provider terms |
| Kubernetes | Tested shape/version, ingress and admin access |
| Metadata | Explicit endpoint first versus link-local milestone |
| Resource ownership | UI/CLI/IaC conflict and adoption policy |
| Persistent-data default | Retain/delete/restore semantics |
| Team capacity | Named frontend/provider/Ansible owners |

## 4. Completion claims

Do not claim full DigitalOcean parity, cloud VM security, public fake-IP routing, tamper-proof logs or an operational service merely because V5 describes it. The defensible contribution is the implemented/tested workflow and its stated boundaries.
