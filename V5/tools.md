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
| Go | Agent, cobra CLI and later thin provider | Pin after testing; record actual version; race/vet/build |
| Java/Maven/Spring | Authenticated control plane | Compiler/runtime, dependencies, packaged boot |
| Python | Contract generation/fixtures and V5 print tooling | Script version and checks |
| Docker/Compose | Workload execution/private dependencies | Engine/plugin and firewall backend |
| k3d + k3s | Required bounded Kubernetes; driver calls CLI | Pin after testing; record actual tested pair and flags |
| kubectl | Scoped cluster verification | Pin after testing; record actual version |
| PostgreSQL | Durable desired state and owned workload DB | Image/digest and backup/restore |
| Valkey | Shared replay-related state | Auth, persistence and failure policy |
| MinIO/S3 client | Generic container process; Java bucket reconcile | Pin after testing; record actual version/digest and tenant/rebuild tests |
| Extended tools | Proxy, tunnel and browser build choices | TBD decision records in [roadmap](CoveStack-Extended-Roadmap.md); not selected here |
| cobra | E thin Go CLI | Pin after testing; record actual version |
| Terraform + terraform-plugin-framework | Optional thin API provider, Droplet/K8s only | Pin after testing; record actual versions and lifecycle fixtures |
| Ansible | Part 1 operator installation | Pin after testing; record actual version; second run changed=0 |
| ansible-lint | Installer checks | Pin after testing; record actual version and result |
| ansible-vault | Installer secret protection | Pin after testing; record actual version; custody B+C |
| golangci-lint | Reviewed Go static checks | Pin after testing; record actual version and rules |
| GOOS/GOARCH targets | linux/amd64, darwin/arm64, windows/amd64 | Pin compiler after testing; record actual version; two actual OS runs |

## 2. Development versus deployment

| Tool | Use | Restriction |
|---|---|---|
| Git/GitHub | Branches, reviewed PRs and source | Do not push secrets/main directly |
| API test client | Repeat positive/negative HTTP cases | Sanitized exports; real auth only |
| curl | Inspect supported HTTP endpoints | Do not bypass TLS verification in release |
| Editor Go/Java tooling | Navigate, format and debug | Editor output is not CI evidence |
| Kubernetes tools | Scoped workload validation | Protect kubeconfig; no public admin endpoint |
| Extended measurements | Actual scoped health/metrics | Roadmap decision; no invented tool selection |
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
| Extended tools | Future deployment-specific choices recorded in roadmap |
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

Run module checks from the proper working directory. Extra tools are tied to planned features and become release requirements only with their gate evidence. Lack of Node/Terraform does not prevent checking current signing/lifecycle code.

## 5. Automation integration policy

Terraform and Ansible already integrate with execution tools; no exclusivity claim is made. CoveStack adds authenticated resource/account semantics. CLI, thin Terraform provider and Ansible inventory talk only to the authenticated API; none import/call Docker/k3d. The operator Ansible Part 1 installer is distinct and intentionally installs host dependencies.

Optional Terraform is Droplet and Kubernetes resources only, after CLI. Part 1 requires reviewed roles, checksum-verified binaries, no_log/vault, systemd hardening and repeat-run changed=0. Part 2 inventory and unrestricted platform-run job ideas remain roadmap targets; arbitrary uploaded configurations must not run with host credentials. Preserve job identity, permissions, timeout, logs/state and exit-result requirements for that future scope.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).

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
| [Extended decisions](CoveStack-Extended-Roadmap.md) | Proxy/tunnel candidates and references; tool choice TBD |

These references are mechanisms, not endorsements of the complete CoveStack design. Match tested versions and document provider-specific constraints.
