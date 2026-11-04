# License, attribution and release-rights review

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **Distribution blocker:** no root LICENSE/COPYING file was found in this checkout. Open-source is the intended product direction, not a verified repository license grant. This documentation task does not choose or add a license.

| Item | Required maintainer action before distribution |
| --- | --- |
| CoveStack code/docs | Confirm ownership/contributor rights and approve an actual license/grant; publish root license and contribution terms in a separately authorized change |
| Third-party dependencies | Inventory exact resolved versions/artifacts and licenses; include required notices and source-offer obligations where applicable |
| Incus/QEMU/KVM and guest images | Review package/image/firmware redistribution rights, notices, kernel/guest obligations and supported acquisition path |
| Docker/k3d and OCI images | Review each image/runtime/backend separately; an image tag is not a licensing decision |
| MinIO | Verify selected server/client artifact/version license and distribution/use obligations; do not assume all storage integrations are permissively licensed |
| Optional Floci | Preserve applicable upstream license/notices; keep it an external integration, not misattributed original code |
| Historical copied handbook | HANDBOOK.md is LocalCloud reference material, not CoveStack implementation or an established license grant; verify origin/permission before redistribution |
| Generated models/vectors/assets | Record origin, generator and rights; generated output does not automatically remove upstream obligations |
| Release | SBOM and notices match actual binaries/images; signatures/provenance and vulnerability policy reviewed |

Floci's repository identifies an MIT license; verify the exact pinned artifact and retain required notices at integration time: [upstream](https://github.com/floci-io/floci). No license compatibility conclusion for the combined product is made here. Official software documentation links are research references, not permission to copy implementation or claim certification.

Do not publish an open-source/production badge without the corresponding rights/evidence. Designating a real security contact, governance/maintainer responsibilities, support policy and dependency-update process is release work. No invented legal text, legal advice, agreement, contact address or mentor confirmation is supplied.
