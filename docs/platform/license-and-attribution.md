# License, attribution and release-rights review

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **Apache License 2.0 adopted with maintainer approval** for CoveStack-owned code and documentation. SPDX: `Apache-2.0`. The complete official [root LICENSE](../../LICENSE) is present. Existing third-party notices and licenses remain in force; adoption does not clear every redistributed dependency/image or copied document. See the [adoption decision](licensing-decision.md).

| Item | Required maintainer action before distribution |
| --- | --- |
| CoveStack code/docs | Apache-2.0 approved and root LICENSE added; maintain contribution provenance and resolve rights for copied/third-party material before redistribution |
| Third-party dependencies | Inventory exact resolved versions/artifacts and licenses; include required notices and source-offer obligations where applicable |
| Incus/QEMU/KVM and guest images | Review package/image/firmware redistribution rights, notices, kernel/guest obligations and supported acquisition path |
| Docker/k3d and OCI images | Review each image/runtime/backend separately; an image tag is not a licensing decision |
| RustFS preferred candidate / historical MinIO | Verify selected server/client artifact/version license and distribution/use obligations; do not assume all storage integrations are permissively licensed |
| Optional Floci | Preserve applicable upstream license/notices; keep it an external integration, not misattributed original code |
| Historical copied handbook | HANDBOOK.md is LocalCloud reference material, not CoveStack implementation or an established license grant; verify origin/permission before redistribution |
| Generated models/vectors/assets | Record origin, generator and rights; generated output does not automatically remove upstream obligations |
| Release | SBOM and notices match actual binaries/images; signatures/provenance and vulnerability policy reviewed |

Floci's repository identifies an MIT license; verify the exact pinned artifact and retain required notices at integration time: [upstream](https://github.com/floci-io/floci). No license compatibility conclusion for the combined product is made here. Official software documentation links are research references, not permission to copy implementation or claim certification.

Apache-2.0 is the project license for CoveStack-owned material; do not infer production certification or blanket third-party redistribution clearance from it. Designating a real security contact, governance/maintainer responsibilities, support policy and dependency-update process is release work. No invented legal text, legal advice, agreement, contact address or mentor confirmation is supplied.
