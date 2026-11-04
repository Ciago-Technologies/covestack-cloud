# Research references and evidence limits

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Reviewed: 4 November 2026. URLs below were opened/searched in official documentation during this task. Web publication is evidence of documented behavior, not runtime validation. No version pin was selected from a latest/main URL.

| Source | What it informs | Scope limit |
| --- | --- | --- |
| [Incus VM/container distinction](https://linuxcontainers.org/incus/docs/main/explanation/containers_and_vms/) | Guest-kernel VM versus host-kernel container | Not an accepted CoveStack driver |
| [Incus security](https://linuxcontainers.org/incus/docs/main/explanation/security/) | Privileged daemon/socket and patch boundaries | No claim of protection from host/root compromise |
| [Incus projects](https://linuxcontainers.org/incus/docs/main/projects/) | Existing ownership/isolation features | Mapping to CoveStack accounts must be designed/tested |
| [Incus requirements](https://linuxcontainers.org/incus/docs/main/requirements/) | Host/platform evaluation prerequisites | No exact CoveStack OS/kernel/backend pair or hardware density certified |
| [KVM](https://www.linux-kvm.org/page/Main_Page) | Linux hardware virtualization | Hardware/nested capability must be measured |
| [QEMU introduction](https://www.qemu.org/docs/master/system/introduction.html) | System emulation, accelerators and management | TCG emulation is not claimed as the production KVM profile |
| [Floci repository](https://github.com/floci-io/floci) | AWS emulator and real selected runtimes | Service/mode-specific fidelity; no runtime audit |
| [LocalStack](https://docs.localstack.cloud/aws/) | Cloud API testing and integrations | Edition/configuration and tested API subset matter |
| [Coolify](https://coolify.io/docs/) | Application hosting workflows | No invented VM feature absence |
| [Dokku](https://dokku.com/docs/) | Self-hosted application deployment | Plugin and access model need bounded assessment |
| [Portainer](https://docs.portainer.io/) | Container/environment management | Distinguish CE/BE capabilities |
| [DevPod](https://devpod.sh/docs/what-is-devpod) | Devcontainer/provider workspaces | Provider machines may be VMs; not inherently container-host-only |
| [CloudStack](https://docs.cloudstack.apache.org/en/latest/conceptsandterminology/concepts.html) | Existing IaaS capabilities | Latest docs do not certify CoveStack comparison performance |
| [Nova](https://docs.openstack.org/nova/latest/) | Compute dependencies and inventory | Old architecture guide is not a current installation baseline |
| [Keystone](https://docs.openstack.org/keystone/latest/admin/identity-concepts.html) | Identity/tenancy concepts | No plan to embed or clone Keystone |
| [Neutron](https://docs.openstack.org/neutron/latest/admin/intro-os-networking.html) | Network roles | No OpenStack parity target |
| [OPA](https://www.openpolicyagent.org/docs) | Optional policy evaluation candidate | No new dependency or requirement chosen |
| [OpenSSF Best Practices](https://www.bestpractices.dev/en) | Release/security practice criteria | No self-certification or badge awarded here |

The requested bestpractices.coreinfrastructure.org URL redirects to bestpractices.dev. Documentation can move; resolve current official links before pinning release requirements. Coolify knowledge-base/teams and DevPod architecture/overview lookup failures were replaced by verified team/how-it-works URLs in competitive analysis. Unknown feature support, hardware requirements, licenses and vulnerability claims are recorded as requiring validation rather than guessed.

## Additional official sources for the catalogue refinement

| Source | Evidence and limit |
| --- | --- |
| [Proxmox VE overview](https://www.proxmox.com/en/products/proxmox-virtual-environment/overview) | KVM/LXC and integrated infrastructure; no benchmark comparison |
| [Proxmox official documentation mirror](https://github.com/proxmox/pve-docs) | Role/storage documentation fallback; primary pve-docs endpoint returned 403 during research |
| [OpenTofu introduction](https://opentofu.org/docs/intro/) and [state locking](https://opentofu.org/docs/language/state/locking/) | Provider/state orchestration; state locking depends on backend, not runtime fencing |
| [MinIO upstream](https://github.com/minio/minio) | Repository archive/AGPL marker requires artifact-specific maintenance/legal review, not a claim all commercial offerings ended |
| [Garage](https://garagehq.deuxfleurs.fr/documentation/quick-start/) | Compatibility caveats include unsupported ACL/policy features; do not assume AWS IAM parity |
| [SeaweedFS](https://github.com/seaweedfs/seaweedfs) | S3/file candidate with Apache-2.0 repository notice; scope of exact artifacts still reviewed |
| [Ceph RGW](https://docs.ceph.com/en/squid/radosgw/) | S3/Swift over RADOS; stable-version reference is not a selected support pin |
| [Harbor](https://goharbor.io/docs/) | Later registry candidate, not installed or selected |
| [OpenBao](https://openbao.org/docs/) | Secrets-provider candidate; custody integration and operational sizing unverified |

No provider is declared maintained, supported or secure from website availability alone. Record exact artifacts, license notices, current security/release evidence and accountable maintainers before selecting a default.

## RustFS and licensing correction sources

Reviewed 4 November 2026: [RustFS S3 matrix](https://docs.rustfs.com/en/reference/s3-compatibility), [IAM](https://docs.rustfs.com/en/security-compliance/iam), [access keys](https://docs.rustfs.com/en/security-compliance/iam/access-token), [installation/container](https://docs.rustfs.com/en/installation/container), [Linux persistent path](https://docs.rustfs.com/en/installation/linux/quick-start), [Kopia destination integration](https://docs.rustfs.com/en/developer/integration/backup/kopia), [RustFS license](https://github.com/rustfs/rustfs/blob/main/LICENSE), [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0.txt), [MIT](https://opensource.org/license/mit). These substantiate candidate/recommendation review, not CoveStack runtime or licensing approval. Older MinIO sources remain historical/provider-comparison evidence only.

Maintainer-approved licensing update: complete official [Apache-2.0 text](https://www.apache.org/licenses/LICENSE-2.0.txt) copied unchanged to [root LICENSE](../../LICENSE). Approval comes from the latest maintainer instruction, not from the research sources; third-party attribution remains separate.
