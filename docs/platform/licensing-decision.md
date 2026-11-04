# Licensing decision — Apache-2.0 adopted

Updated: 4 November 2026. **Adopted with explicit maintainer approval in the current task.** License: Apache License 2.0; SPDX identifier: `Apache-2.0`. The repository-root [LICENSE](../../LICENSE) contains the complete official text downloaded from apache.org, unchanged, including its appendix. No copyright holder or author identity has been invented. The appendix's example placeholders are part of the official text, not a fabricated project copyright statement.

The license applies to CoveStack-owned code and documentation. Third-party material, dependencies, guest images and external projects retain their own licenses and attribution. Project license adoption does not establish provenance for copied material or final clearance for all redistributable artifacts.

## Rationale and MIT comparison

| Topic | Apache-2.0 | MIT |
| --- | --- | --- |
| Adoption and reuse | Permissive commercial/modification/distribution model | Permissive commercial/modification/distribution model |
| Patent terms | Express contributor patent grant with specified patent-litigation termination | No express patent grant in the license text |
| Redistribution | Preserve license and applicable notices; identify changed files; retain relevant NOTICE attribution when present | Preserve copyright and permission notice in copies/substantial portions |
| Trademarks / warranty | No general trademark grant; warranty/liability disclaimers | Warranty/liability disclaimers; do not infer brand rights |
| Operational tradeoff | More notice discipline; useful explicit terms for a multi-contributor infrastructure product | Shorter and simpler attribution text |

**Apache-2.0 is adopted for CoveStack-owned code and documentation following the maintainer's explicit approval.** Explicit patent terms are the main reason; this is not a claim it prevents all patent risk. MIT remains a reasonable simpler alternative. Neither choice licenses third-party material or guarantees compatibility of every combined distribution. Sources: [Apache-2.0 text](https://www.apache.org/licenses/LICENSE-2.0.txt), [MIT text](https://opensource.org/license/mit).

## Dependency and redistribution review

Read-only evidence: kan-agent/go.mod declares no external module dependencies; Go runtime/standard-library notices still matter. kan-server/pom.xml declares Spring Boot/web/JDBC/Redis/validation/test, pgjdbc, Flyway and Jackson modules. This is a direct-dependency inventory, NOT a resolved transitive/license bill of materials. No package resolution, build or scanner was run.

| Artifact / category | Evidence reviewed | Action before redistribution |
| --- | --- | --- |
| Go runtime/standard library | [Upstream BSD-style license](https://go.dev/LICENSE) | Include applicable copyright/license/disclaimer with distributed binaries; review exact toolchain and bundled material |
| Spring Boot / Framework | [Boot contribution/license statement](https://github.com/spring-projects/spring-boot/blob/main/CONTRIBUTING.adoc) | Apache-2.0 upstream; preserve relevant notices and examine exact resolved/transitive artifacts |
| PostgreSQL JDBC | [Upstream LICENSE](https://github.com/pgjdbc/pgjdbc/blob/master/LICENSE) | BSD-style terms; retain required notices for exact jar and distribution |
| Flyway | [Upstream license file](https://github.com/flyway/flyway/blob/main/LICENSE.txt) | Review exact resolved edition/module and transitive notices; repository licensing is not permission for every commercial package |
| Jackson, Redis client, validation, tests and other transitives | Declared or transitively expected from pom.xml; full artifact set not resolved here | Generate exact artifact inventory/SBOM and check each license/NOTICE; test dependencies matter if shipped |
| RustFS candidate | [Upstream Apache-2.0 license](https://github.com/rustfs/rustfs/blob/main/LICENSE) | Review selected image/binary and bundled dependencies, preserve applicable license/NOTICE; no inherited CoveStack license grant |
| Incus/QEMU/KVM, Docker/k3d, guest/OCI images, PostgreSQL/Valkey/Mailpit | External runtimes/services/images, not all statically linked libraries | Distinguish API use from bundling/modification; inspect exact package/image notices, source obligations, distribution rights and trademarks |
| Historical MinIO / copied handbook / generated assets | Historical references do not establish copying or redistribution rights | Preserve attribution; exclude or obtain permission where needed; never relabel upstream content as CoveStack Apache-2.0 |

External process separation does not erase obligations when binaries/images are bundled or modified. Do not relicense GPL/AGPL or other upstream material through a CoveStack root LICENSE. Exact distribution and compatibility review remains pending. [Apache guidance on applying the license](https://apache.org/legal/apply-license.html) informs packaging practice; CoveStack is not an Apache Software Foundation project.

## Adoption record and remaining distribution work

| Item | Current disposition |
| --- | --- |
| Previous proposal | Earlier documentation recommended Apache-2.0 versus MIT and correctly recorded approval as pending at that time |
| Approval | The latest user instruction explicitly states that the maintainer approved Apache License 2.0; no additional author identities or consents are inferred |
| License file | Complete official Apache-2.0 text added unchanged as root LICENSE; no source headers, manifests or upstream files changed |
| Documentation | README, contribution guide, active product/SRS/status and agent handoff now use SPDX Apache-2.0 consistently |
| Third-party rights | Exact resolved dependency/image inventory, required notices/source obligations and copied-material provenance remain separate review tasks |

## NOTICE decision

No project-specific NOTICE is added in this task: no approved original attribution text or verified bundled dependency NOTICE content was supplied. Do not invent copyright holders, copy an ASF project NOTICE or imply CoveStack is an ASF project. Apache-2.0 section 4(d) requires retaining applicable upstream NOTICE attribution when present; it does not require fabricating an empty or generic project NOTICE. During packaging, inspect actual distributed artifacts and include required notices in the permitted form, creating an appropriate NOTICE if that review calls for one. Existing upstream notices/copyrights are untouched.

## Remaining maintainer responsibilities

1. Maintain contribution provenance and resolve any copied-reference, institute/employer or other third-party ownership questions without assuming this approval relicenses their work.
2. Review exact resolved dependencies, runtime packages, firmware/guest images and shipped binaries; retain licenses/notices and fulfill applicable source/distribution obligations. No package resolution, SBOM generation or legal clearance ran here.
3. Define any additional governance, CLA/DCO or disclosure process explicitly; this task does not invent one. Contributor guidance follows the root license and preserves upstream material.
4. Review product packaging metadata/headers later under separate authorization; implementation/dependency files were intentionally untouched.
5. Keep service maturity distinct: an open-source license neither verifies a feature nor makes the platform production-ready.
