# CoveStack V5 — Tools, Versions and Why They Are Used

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Tool selection policy

Prefer the smallest tested set. Each release has one recorded matrix of compiler, runtime, images, CLI tools and plugin versions. Existing pins are observed repository inputs; verify actual builds and current support before release. Do not blindly use `latest`, a course's old version, or an unverified download.

| Tool/group | Role | Evidence to record |
|---|---|---|
| Go | Agent and proposed Terraform provider | Compiler selection, tests/race/vet |
| Java/Maven/Spring | Authenticated control plane | Compiler/runtime, dependencies, packaged boot |
| Python | Contract generation/fixtures and V5 print tooling | Script version and checks |
| Docker/Compose | Workload execution/private dependencies | Engine/plugin and firewall backend |
| k3d/k3s/kubectl | Bounded local Kubernetes | Compatible version pair, client and cluster |
| PostgreSQL | Durable desired state and owned workload DB | Image/digest and backup/restore |
| Valkey | Shared replay-related state | Auth, persistence and failure policy |
| MinIO/S3 client | Object storage | Version/digest, operations and policy tests |
| Nginx or equivalent proxy | HTTP routing and LB execution | Configuration test/reload and TLS |
| Tunnel connector | Real public app connectivity | Provider/config/version and outage behavior |
| React/Node build tools | SPA | Lockfile/build and browser/API tests |
| Terraform/framework | API provisioning integration | Provider version/state acceptance tests |
| Ansible | Supported configuration/inventory | Core/collection versions and repeat runs |

## 2. Development versus deployment

| Tool | Use | Restriction |
|---|---|---|
| Git/GitHub | Branches, reviewed PRs and source | Do not push secrets/main directly |
| API test client | Repeat positive/negative HTTP cases | Sanitized exports; real auth only |
| curl | Inspect supported HTTP endpoints | Do not bypass TLS verification in release |
| Editor Go/Java tooling | Navigate, format and debug | Editor output is not CI evidence |
| Kubernetes tools | Scoped workload validation | Protect kubeconfig; no public admin endpoint |
| Monitoring stack/CLI | Actual health/metrics | Avoid exposing host/admin data to all accounts |
| Image/config scanners | Findings and remediation | Scanner success is not proof of safety |

## 3. Version manifest contents

The team should add a release manifest under the approved deployment/evidence location. V5 does not alter runtime pins.

| Field | Meaning |
|---|---|
| Commit and date | Source snapshot tested |
| Host and VM OS/kernel | Network/privilege context |
| Go and Java actual version | Not only configured preference |
| Image references/digests | Exact control-plane/workload inputs |
| k3d/k3s pair | Supported cluster shape/version |
| Proxy/tunnel tooling | Deployment-specific routing behavior |
| Terraform/Ansible plugins | Accepted integrations and dependencies |
| Verification commands/results | Build, integration, security and restore |

## 4. Useful checks

```bash
go version
java -version
javac -version
mvn -version
docker version
docker compose version
```

Run module checks from the proper working directory. Extra tools become requirements only when the relevant feature is implemented. Lack of Node/Terraform does not prevent checking current signing/lifecycle code.

## 5. Automation integration policy

Terraform and Ansible already integrate with Docker. CoveStack's integration must add its resource/account/stack semantics rather than claim exclusive support. A provider calls the API; an Ansible module/inventory or approved playbook uses authenticated operations and appropriate private connectivity.

Do not upload arbitrary configurations and execute them with host credentials. Begin with reviewed templates and validated variables. Track job identity, permissions, timeout, logs, state and exit result.

## 6. Networking tools

Record routes, DNS answers and actual listener addresses using platform-appropriate tools. Confirm Docker's iptables/nftables mode and test forwarded traffic, not only host INPUT policy. Use the host's supported firewall mechanism; do not copy unrelated Linux distro instructions or disable SELinux to hide a failure.

## 7. Primary references

| Reference | Use |
|---|---|
| [Go documentation](https://go.dev/doc/) | Language, tooling and standard library |
| [Docker networking](https://docs.docker.com/engine/network/) | Execution/network mechanisms |
| [k3d exposure](https://k3d.io/stable/usage/exposing_services/) | Cluster ingress/port integration |
| [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) | Provider implementation |
| [Ansible Docker collection](https://docs.ansible.com/projects/ansible/latest/collections/community/docker/index.html) | Existing automation mechanisms |
| [Cloudflare Tunnel routing](https://developers.cloudflare.com/tunnel/concepts/routing/) | Candidate public gateway |

These references are mechanisms, not endorsements of the complete CoveStack design. Match tested versions and document provider-specific constraints.
