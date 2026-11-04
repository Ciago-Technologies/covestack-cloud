# CoveStack V5 — Tools, Versions and Responsibilities

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Tool selection policy

Use current source pins for existing code; select runtime/provider artifacts only after compatibility, security, rights and capacity testing. A version string or available website is not support. No software was installed/upgraded or runtime started by this documentation alignment. Keep Java control plane and Go agent; no wholesale rewrite or mandatory cloud account.

## 2. Existing development and dependency inputs

| Tool / surface | Current input | Purpose / limit |
| --- | --- | --- |
| Git/editor | Actual developer environment | Feature branches/PR; never discard user work or commit secrets |
| Go | go.mod/CI toolchain 1.27.2, language baseline 1.22 | Existing agent/signing/client tests; historical tool output differs and must be labelled |
| Java / Spring Boot | pom Java 25 / Boot 3.5.16 | Auth/JDBC control-plane source; actual app foundation pending |
| Maven | No wrapper assumed; inspect local version | Existing verify command; build is not startup/schema proof |
| Python 3 | Current generator/checkers use stdlib | Contract fixture/output validation; no application provisioning |
| PostgreSQL | Current platform 17 direction; Compose image must be exact tested tag/digest | Durable intent/operations, not JPA; foundation/real migration acceptance missing |
| Valkey | Compose exact reviewed image input, example family 8 | Authenticated atomic replay store; not resource DB/lease authority |
| Mailpit | Exact tested image input | Local development email capture, no production SMTP guarantee |
| Existing CI scanner | Workflow-pinned gitleaks with verified checksums | Source secret scan; not a full dependency/security audit |

## 3. Planned runtime and interface toolchain

| Tool / backend | Responsibility | Selection / acceptance |
| --- | --- | --- |
| Incus / QEMU / KVM | Preferred real VM management stack | Pin after authorized Linux/image/guest/disk/NIC/security/task-recovery spike; no minimum RAM/density invented |
| Docker / approved OCI | Distinct trusted App Containers and internal service runtime | Privileged socket restricted; actual image/resources/health/volume/policy tests |
| k3d + k3s | Local Kubernetes cluster lifecycle via CLI | One tested pair per profile; verify actual flags/version before documentation; no production Kubernetes SLA |
| RustFS | Preferred Planned S3 provider; operator process/volume, Java bucket/key/policy adapter | Exact artifact/admin SDK/custody/policies/restart/verified restore; not full S3 parity |
| SeaweedFS / Garage / Ceph RGW | Alternative object adapters if selected | Capability/security/maintenance/licensing/resource/recovery review; no binary swap guarantee |
| cobra / Go HTTP client | E's thin API CLI | Module/user-auth boundary with A; no runtime/DB import; JSON/wait/idempotency tests |
| Terraform plugin framework / Terraform / OpenTofu | Later thin approved compute/Kubernetes provider | Pin actual tested framework/tool/provider subset; refresh/import only if implemented and verified |
| Ansible / lint / Vault or approved custody | Operator installer; later API inventory | Exact tested versions, protected secrets, artifact verify, lint, repeat install/no unexpected changes |
| Go lint / static/dependency tools | Reviewed source quality/supply-chain checks | Existing versus proposed checks labelled; no guessed pin or installation |
| Mature secret backend | Evaluate custody/rotation/recovery, e.g. OpenBao candidate | Not selected/installed; no homemade vault or assumed KMS compatibility |
| Gateway/proxy/tunnel/certificate tools | Future explicit publication/domain service | TBD decision record; real routing/TLS/SSRF/management denial required |

## 4. Build targets and checks

CLI packaging target linux/amd64, darwin/arm64 and windows/amd64 after module exists; actual runtime platform support requires OS-run evidence, not cross-compile alone. Existing Go/Java/contract checks are in [setup](CoveStack-Complete-Working-Setup.md). Never run fictional CLI/server/driver commands or change pins to hide unavailable tooling. B/C/D need real Java foundation, A/E correct Go source/module boundary, C authoritative tests.

## 5. Deployment profiles and resource constraints

Old 8 GB AlmaLinux context is development/lab only. Supported self-hosted production needs exact measured compatible OS/kernel/backend/image/storage/network/workload plus hardening/restore/upgrade evidence. Single-host is not HA, multi-host future. Incus placement/IPAM mechanisms need explicit mapping to platform account/environment policy, not competing schedulers/allocators.

## 6. Networking, storage and rights

No router DNS/DHCP or inbound internet path assumed on hotspot. Fake public IP is demonstration-only; actual route/tunnel/domain and authorization needed. Privileged sockets/host mounts never tenant access. Apache-2.0 covers CoveStack-owned code/docs; runtime/guest/container/SDK dependencies retain their licenses/notices. Record artifact source/digest/signature, SBOM and required redistribution obligations.

## 7. References

[Canonical source references](../docs/platform/references.md), [provider evaluation](../docs/platform/provider-evaluation.md), [license/attribution](../docs/platform/license-and-attribution.md), [setup/profile](CoveStack-Complete-Working-Setup.md), [operations](CoveStack-Operations-and-Recovery.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-tools) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
