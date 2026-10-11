# Covestack Cloud — Documentation Change Log

Historical documentation rewrite record: line references/hashes below describe the earlier
revision. See Covestack-Cloud-Branding-Changes.md for the current rename/status changes.

Updated 11 October 2026. Six existing guides were consolidated and rewritten to eliminate competing algorithms, unsafe starters and inconsistent timelines. This is a full-file revision, not a claim that only isolated lines changed. Original six documents are preserved byte-for-byte in `Kan-Cloud-Docs-Before-Fixes-2026-10-11.zip`; the unchanged audit preserves its original evidence.

**Implementation status:** documentation changes are applied. No kan-cloud source files, runtime services, credentials, VM settings or contracts were changed. [Current setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) records the required implementation for all 38 findings.

## 1. Complete file changes

Line numbers refer to the original ZIP contents and the final revised documents. The exact unified diff (historical patch in Claude) lists every removed/added line, including condensed repetitive sections. Old line ranges describe the primary affected topic, not an exhaustive statement that only that topic changed.

| File | Original lines replaced | Final lines | Reason for consolidation |
|---|---|---|---|
| [Covestack-Cloud-Technical-Design.md](Covestack-Cloud-Technical-Design.md) | 1–1818 | 1–234 | One corrected technical design; remove unsafe code starters and false implemented/compliance claims. |
| [Covestack-Cloud-Complete-Working-Setup.md](Covestack-Cloud-Complete-Working-Setup.md) | 1–665 | 1–254 | Make current paths, prerequisites, infra proposal, secrets, persistence and networking actionable. |
| [Covestack-Cloud-Project-Workflow.md](Covestack-Cloud-Project-Workflow.md) | 1–1887 | 1–139 | Replace obsolete mixed timeline with actual 2026 gates, owner acceptance and December buffer. |
| [Covestack-Cloud-Team-Roles.md](Covestack-Cloud-Team-Roles.md) | 1–1401 | 1–124 | Keep five owners while removing repeated contradictory algorithms/examples and assigning tests. |
| [Covestack-Cloud-Tools.md](Covestack-Cloud-Tools.md) | 1–1492 | 1–108 | Replace drifting version/platform/install advice with supported pinned versions and verified tooling. |
| [Covestack-Cloud-Repository-Layout.md](Covestack-Cloud-Repository-Layout.md) | 1–68 | 1–74 | Distinguish actual skeleton from planned modules and correct descriptions of existing code. |

## 2. Section changes: file, lines, problem and solution

All new sections below are part of the corresponding full-file replacement. Links jump to the final section; numeric ranges refer to this final revision. Audit IDs link the reason back to the historical report.

| File / final section | Original topic lines | Final lines | Problem | Solution | Audit IDs |
|---|---|---|---|---|---|
| [Covestack-Cloud-Technical-Design.md: 1. Actual baseline and document authority](Covestack-Cloud-Technical-Design.md) | 276–308 | 7–14 | Docs claimed starter classes/features were complete and proposed different active folders. | Distinguished skeleton from target, retained actual folders and defined contract authority. | 4–5, 9, 37 |
| [Covestack-Cloud-Technical-Design.md: 2. Architecture and release scope](Covestack-Cloud-Technical-Design.md) | 19–308 | 15–34 | Security was attributed to host execution and release scope conflicted across guides. | Defined Docker authority and one bounded required/planned feature scope. | 18, 24, 37 |
| [Covestack-Cloud-Technical-Design.md: 3. Toolchain, builds and configuration](Covestack-Cloud-Technical-Design.md) | 152–171 | 35–44 | Java/Boot and Go/version guidance was inconsistent; missing dependencies were overlooked. | Added supported version pinning, Flyway/mail requirements and host/container endpoint distinctions. | 10–11, 13, 37 |
| [Covestack-Cloud-Technical-Design.md: 4. Authentication and replay protection](Covestack-Cloud-Technical-Design.md) | 924–1007 | 45–100 | Starter routing, signed headers, nonce order and storage did not provide the claimed controls. | Specified registered route matrix, signed nonce/key, bounded reads, verify-before-store, principal and durable replay policy. | 1–8, 14–15, 34 |
| [Covestack-Cloud-Technical-Design.md: 5. Contract and state vocabulary](Covestack-Cloud-Technical-Design.md) | 492–644 | 101–129 | State/examples and observation schema disagreed; new kinds/auth/completeness were unspecified. | Defined proposed kind states, nullable generation/timestamp/order, extensions and complete known-ID snapshots. | 16–19, 36 |
| [Covestack-Cloud-Technical-Design.md: 6. Persistence, generations and actions](Covestack-Cloud-Technical-Design.md) | 327–419 | 130–157 | Whole state/action logic could acknowledge failures, regress generation or ignore idempotency races. | Defined scoped transactions, validated reports, error/retry outcomes, explicit reboot and deleted_at/GC. | 20–22, 33, 36 |
| [Covestack-Cloud-Technical-Design.md: 7. Agent correctness and scaling](Covestack-Cloud-Technical-Design.md) | 1008–1105 | 158–169 | Event/health, pending counts, single-agent ownership and port reservation were incomplete. | Required fencing, periodic inspection, dirty retries, bounded calls and pending-aware ASG. | 20, 23, 36 |
| [Covestack-Cloud-Technical-Design.md: 8. Workload isolation, SSH and limits](Covestack-Cloud-Technical-Design.md) | 1443–1524 | 170–179 | Default root password, unsafe shell injection and assumed capability/network isolation. | Required key-only tested SSH, safe argument/file handling, quotas and account/backend policies. | 24–26 |
| [Covestack-Cloud-Technical-Design.md: 9. Databases, Spaces, load balancers and snapshots](Covestack-Cloud-Technical-Design.md) | 1142–1612 | 180–201 | Data/advanced-feature starters did not define persistence, S3 IAM or correct provider endpoints. | Required persistent volumes, scoped S3, valid LB reload, bounded k3d and correct snapshot action/limitations. | 27–28, 31–32, 34 |
| [Covestack-Cloud-Technical-Design.md: 10. DNS, TLS and public routing](Covestack-Cloud-Technical-Design.md) | 1232–1348 | 202–211 | Loopback listener/NAT and signed Host/public internal routes conflicted. | Defined actual guest listener, TCP+UDP DNS, private agent path, TLS expiry and UI routing. | 29–30 |
| [Covestack-Cloud-Technical-Design.md: 11. CI, tests, operations and truthful evidence](Covestack-Cloud-Technical-Design.md) | 1613–1732 | 212–223 | Tests/scripts and claimed compliance could report success without proof. | Required real CI/negative tests, checked SDK failures, truthful control evidence and operational recovery. | 12–13, 35–36 |
| [Covestack-Cloud-Technical-Design.md: 12. Official references and implementation status](Covestack-Cloud-Technical-Design.md) | 1804–1818 | 224–234 | Design remedies were portrayed as implemented and repeated obsolete summaries persisted. | Linked primary references and explicitly kept all remedies pending repository acceptance. | 10–11, 37 |
| [Covestack-Cloud-Complete-Working-Setup.md: 1. Keep the existing project layout](Covestack-Cloud-Complete-Working-Setup.md) | 124–151 | 5–16 | Layout/build paths assumed nonexistent active modules. | Used actual folders, root path and preserved edits. | 9, 37 |
| [Covestack-Cloud-Complete-Working-Setup.md: 2. Verify the machine before installing again](Covestack-Cloud-Complete-Working-Setup.md) | 64–123 | 17–34 | Reinstallation, old versions and unmeasured VM allocation caused drift. | Verify before installing; pin Java/Go and measure VM capacity; retain security controls. | 10, 13–14, 37 |
| [Covestack-Cloud-Complete-Working-Setup.md: 3. Prepare secrets and versions](Covestack-Cloud-Complete-Working-Setup.md) | 152–257 | 35–54 | Static passwords/untracked version variables and password rotation semantics unclear. | Protected local files, explicit env-file, reviewed images and coordinated existing-volume credential rotation. | 13–15, 27 |
| [Covestack-Cloud-Complete-Working-Setup.md: 4. Proposed infrastructure Compose replacement](Covestack-Cloud-Complete-Working-Setup.md) | 154–212 | 55–137 | Compose examples had wrong MinIO mount and no managed replay storage/isolation. | Provided infra proposal with named data volumes, file secrets, backend/storage networks and correct /data. | 9, 14–15, 27–28 |
| [Covestack-Cloud-Complete-Working-Setup.md: 5. Java server changes before startup](Covestack-Cloud-Complete-Working-Setup.md) | 152–257 | 138–155 | Server startup lacked main/migrations/Flyway module and wrong Docker build context. | Listed prerequisites, aligned dependencies and corrected host/container build/runtime configuration. | 9–11 |
| [Covestack-Cloud-Complete-Working-Setup.md: 6. Fix and wire authentication before API demos](Covestack-Cloud-Complete-Working-Setup.md) | 258–312 | 156–163 | Four-step signing and absent security wiring were presented as working. | Referenced corrected simple-HMAC route/auth flow and required negative acceptance tests. | 1–8, 34, 37 |
| [Covestack-Cloud-Complete-Working-Setup.md: 7. Go agent and driver setup](Covestack-Cloud-Complete-Working-Setup.md) | 398–432 | 164–184 | go run . targeted no runnable main and drivers/SDK were assumed installed. | Specified main/config path, pinned imported SDK, proper startup and behavior checks. | 9, 20–26, 36 |
| [Covestack-Cloud-Complete-Working-Setup.md: 8. VM routing, DNS and TLS](Covestack-Cloud-Complete-Working-Setup.md) | 433–455 | 185–230 | DNS/NAT/TLS configuration exposed or failed to reach intended interfaces. | Reserved guest address, controlled forwards, valid DNS/proxy examples and restricted admin/S3 endpoints. | 29–30 |
| [Covestack-Cloud-Complete-Working-Setup.md: 9. Workload connectivity and persistence](Covestack-Cloud-Complete-Working-Setup.md) | 623–640 | 231–240 | External metadata/port forwarding and data recreation were not tested. | Finite per-port plan, explicit binding, client hostnames and real persistence/account tests. | 24–28, 31 |
| [Covestack-Cloud-Complete-Working-Setup.md: 10. Acceptance checklist](Covestack-Cloud-Complete-Working-Setup.md) | 595–622 | 241–254 | Demo checklist treated missing implementations as ready and repeated inconsistent timelines. | Replaced with actual security/lifecycle/data/remote acceptance and linked one dated workflow. | 20–38 |
| [Covestack-Cloud-Project-Workflow.md: 1. Baseline and rules for finishing](Covestack-Cloud-Project-Workflow.md) | 1324–1887 | 7–16 | Complete/do-not-modify claims and dual layouts hid existing blockers. | Declared skeleton baseline, one layout and non-negotiable security acceptance. | 1–19, 37 |
| [Covestack-Cloud-Project-Workflow.md: 2. Scope and order](Covestack-Cloud-Project-Workflow.md) | 11–54 | 17–30 | Different documents promised incompatible services and unlimited depth. | Ordered foundation/data/bounded advanced paths and explicit unsupported behavior. | 18, 31–32, 37 |
| [Covestack-Cloud-Project-Workflow.md: 3. First 72 hours: 12–14 October](Covestack-Cloud-Project-Workflow.md) | 91–307 | 31–42 | Early tasks permitted auth TODO-pass/dev-skip and unprotected key bootstrap. | Assigned first 72 hours to real security/build/contracts with measurable owners. | 1–19, 34 |
| [Covestack-Cloud-Project-Workflow.md: 4. Weekly gates](Covestack-Cloud-Project-Workflow.md) | 55–90 | 43–61 | Dates/weekday labels wrong and packaging/security landed too late. | Corrected 2026 gates; Oct/Nov integration, Dec 1 freeze, Dec 8 RC and protected submission buffer. | 38; all scheduling |
| [Covestack-Cloud-Project-Workflow.md: 5. Work allocation by phase](Covestack-Cloud-Project-Workflow.md) | 308–620 | 62–119 | Weekly samples could report failed actions complete, over-scale or lose data. | Scheduled correct lifecycle/ASG/data/advanced fixes with evidence by phase. | 20–36 |
| [Covestack-Cloud-Project-Workflow.md: 6. Daily execution and escalation](Covestack-Cloud-Project-Workflow.md) | 1301–1313 | 120–127 | Ownership and checked boxes could claim completion without tested result. | Acceptance statuses, pairing/escalation and actual reviewed integration demonstrations. | 35, 37 |
| [Covestack-Cloud-Project-Workflow.md: 7. Submission checklist](Covestack-Cloud-Project-Workflow.md) | 1133–1160 | 128–139 | Final list overclaimed parity/compliance and relied on unstable delivery. | Required real negative/persistence/remote/artifact evidence and explicit feature limitations. | 12–15, 35–38 |
| [Covestack-Cloud-Team-Roles.md: 1. Team responsibilities](Covestack-Cloud-Team-Roles.md) | 45–59 | 5–16 | Role ownership referenced alternate inactive module paths. | Mapped five owners to actual files/packages and single service boundaries. | 9, 37 |
| [Covestack-Cloud-Team-Roles.md: 2. Anuj: agent, infrastructure and review](Covestack-Cloud-Team-Roles.md) | 60–387 | 17–35 | Agent duties omitted correct failure/health/storage and overstated host/network safety. | Assigned fencing/retries/ports, key SSH, persistence, driver and verified CI acceptance. | 1, 13–15, 20–32, 36 |
| [Covestack-Cloud-Team-Roles.md: 3. B: core state, SQL and agent API](Covestack-Cloud-Team-Roles.md) | 388–598 | 36–55 | Observed SQL/ASG/idempotency starters were incomplete or unsafe. | Assigned scoped generation/action/reboot/idempotency/ASG and correct fake isolation. | 16–23, 33, 36 |
| [Covestack-Cloud-Team-Roles.md: 4. C: authentication and authorization](Covestack-Cloud-Team-Roles.md) | 599–770 | 56–72 | Four-step crypto, nonce-first validation and secret-only identity were conflicting. | Assigned one contract, verify-before-store, principal, route coverage, bounded input and negative tests. | 1–8, 14–15, 28, 34–35 |
| [Covestack-Cloud-Team-Roles.md: 5. D: DigitalOcean adapters and real clients](Covestack-Cloud-Team-Roles.md) | 771–1068 | 73–88 | Wrong snapshot path, contradictory 501 scope and unchecked SDK loops. | Correct DO action path, bounded supported API set and asserted SDK results. | 17–19, 31–32, 35, 37 |
| [Covestack-Cloud-Team-Roles.md: 6. E: management API, CLI and mail](Covestack-Cloud-Team-Roles.md) | 1069–1300 | 89–106 | CLI auth/bootstrap, localhost doctor, credential and kubeconfig handling disagreed. | Assigned local bootstrap, Bearer management, remote checks, scoped S3 and safe config/mail. | 9, 28–30, 34 |
| [Covestack-Cloud-Team-Roles.md: 7. Pairing and reviews](Covestack-Cloud-Team-Roles.md) | 1375–1387 | 107–119 | Scary boundaries lacked exact cross-owner acceptance responsibilities. | Assigned reviews/evidence for signing, lifecycle, scaling, storage and public routing. | 1–38 |
| [Covestack-Cloud-Team-Roles.md: 8. Schedule and completion evidence](Covestack-Cloud-Team-Roles.md) | 1301–1374 | 120–124 | Old lifecycle repeated wrong algorithm and multiple timelines. | Linked one dated plan and audit-based implementation acceptance rather than restating crypto. | 37–38 |
| [Covestack-Cloud-Tools.md: 1. One tested version matrix](Covestack-Cloud-Tools.md) | 336–418 | 5–22 | JDK 25 install text expected 21; old Go/image/kubectl examples drifted. | One recorded supported matrix with specific Go security update and official references. | 10–11, 13, 37 |
| [Covestack-Cloud-Tools.md: 2. Who installs what](Covestack-Cloud-Tools.md) | 110–137 | 23–33 | Host versus guest/developer requirements caused unnecessary installation. | Separated actual roles and explicit fake/real integration. | 9, 37 |
| [Covestack-Cloud-Tools.md: 3. JDK and Maven installation](Covestack-Cloud-Tools.md) | 419–470 | 34–49 | Platform instructions pointed at inconsistent JDK/version paths. | Aligned JDK/javac/Maven/IDE verification with trusted vendor installs. | 10, 37 |
| [Covestack-Cloud-Tools.md: 4. Docker and VMware installation](Covestack-Cloud-Tools.md) | 577–709 | 50–57 | Broad trusted firewall zone/root-authority assumptions and backend forwards were unsafe. | Use official installation, restricted daemon authority and narrowly tested NAT/firewall settings. | 14, 24, 29–30 |
| [Covestack-Cloud-Tools.md: 5. Go, SDK and executable tools](Covestack-Cloud-Tools.md) | 508–544 | 58–72 | Old tarballs/@latest scripts and tidy without imports gave misleading SDK setup. | Pinned supported checksum-verified tools, compatible SDK and actual imports/tests. | 9, 13, 37 |
| [Covestack-Cloud-Tools.md: 6. Authenticated clients and certificates](Covestack-Cloud-Tools.md) | 794–842 | 73–80 | Arbitrary fake tokens/auth default cloud URL and CA/credential protection insufficient. | Emulator URL during initialization, scoped real tokens and safe secret/CA handling. | 4–5, 28, 34 |
| [Covestack-Cloud-Tools.md: 7. Build and tool verification](Covestack-Cloud-Tools.md) | 1240–1492 | 81–97 | Generic version flags/pipelines/fixture explanations could hide failures. | Checked proper build commands, actual fixture-path semantics and explicit error handling. | 8–12, 35 |
| [Covestack-Cloud-Tools.md: 8. Storage and maintenance](Covestack-Cloud-Tools.md) | 1154–1239 | 98–103 | Blanket prune called safe; persistence/master-key restore unaddressed. | Protected named volumes and data-aware maintenance/restore. | 15, 27, 32, 36 |
| [Covestack-Cloud-Tools.md: 9. CI tools](Covestack-Cloud-Tools.md) | 1008–1064 | 104–108 | Security tool list implied CI already ran builds/scans and moving binaries were trusted. | Required actual pinned jobs/artifacts and defensible evidence. | 12–13, 35 |
| [Covestack-Cloud-Repository-Layout.md: 1. Present repository](Covestack-Cloud-Repository-Layout.md) | 1–68 | 5–33 | Inventory called an incomplete skeleton a complete active architecture. | Listed present files/empty placeholders and preserved actual layout. | 9, 37 |
| [Covestack-Cloud-Repository-Layout.md: 2. Contracts: executable agreement](Covestack-Cloud-Repository-Layout.md) | 1–68 | 34–47 | Two vectors called 100 percent compatibility and contracts called undisputed law despite conflicts. | Assigned concrete contract fixes and limited what vector success proves. | 1, 7–8, 16–19 |
| [Covestack-Cloud-Repository-Layout.md: 3. Java application](Covestack-Cloud-Repository-Layout.md) | 1–68 | 48–55 | Java starter/filter dependencies presented as complete application. | Identified incomplete files and missing app/config/migrations/auth; removed module relocation. | 3–6, 9–11 |
| [Covestack-Cloud-Repository-Layout.md: 4. Go application](Covestack-Cloud-Repository-Layout.md) | 1–68 | 56–61 | Sign falsely described as injecting HTTP headers; rules.go called normalizer. | Corrected actual function responsibilities and planned entry/client/drivers. | 1, 7–9 |
| [Covestack-Cloud-Repository-Layout.md: 5. Deploy and CI](Covestack-Cloud-Repository-Layout.md) | 1–68 | 62–69 | Root compose/production parity and nonexistent CI testing overstated. | Corrected Compose context/startup, planned secret/volume/network and real CI jobs. | 9, 12–15, 27 |
| [Covestack-Cloud-Repository-Layout.md: 6. Documentation ownership](Covestack-Cloud-Repository-Layout.md) | 1–68 | 70–74 | Multiple guides competed as sources with no traceable update record. | Defined guide purposes, all-findings checklist, line log and historical archive. | 37–38 |

## 3. New documents and artifacts

| File | Purpose | Status |
|---|---|---|
| [Covestack-Cloud-Current-Setup-Fixes.md](Covestack-Cloud-Current-Setup-Fixes.md) | All 38 audit IDs mapped to repository paths, owners, dates, acceptance and VM changes | Generated; repository implementation pending |
| Covestack-Cloud-Documentation-Change-Log.md | This before/after line and reason ledger | Generated |
| `Kan-Cloud-Documentation-Changes.patch` (historical artifact in Claude) | Exact original-to-revised six-file unified diff | Generated |
| Kan-Cloud-Docs-Before-Fixes-2026-10-11.zip | Original six documents for review/recovery | Created before editing |

## 4. Integrity and review

The audit report was not edited. No earlier verification result was rewritten into a new claim. Hashes below identify the original and corrected six documents; this log can be regenerated after later edits if line references change.

| File | Original SHA-256 | Revised SHA-256 |
|---|---|---|
| Covestack-Cloud-Technical-Design.md | `2c9d731055950cb3d50a2659a55e7b5184b0e80b942272a735c562d174bae910` | `877c9f9f8a5ab73775c6b10c9ef6462c1a9beb78389abf72f238f8dd43d478f9` |
| Covestack-Cloud-Complete-Working-Setup.md | `d5d34fec9be4e3fa109d23d85b31bd59c047c20a2f9c9737d1f6a528c9039c9d` | `124a04be04fc76f8b72e87ceb34813989b2be02e8872cf7b0ef5019cbc474bb3` |
| Covestack-Cloud-Project-Workflow.md | `45e9577a37216ea94ac432fa3f07a0108d6fffa432e78a4147d9c27dd98304d1` | `393cd80cfa46beff20625a24616ea3f02780a447d7890a26f0f1ee049d6985b1` |
| Covestack-Cloud-Team-Roles.md | `c2df81b95181367beab0d047a4848d392e619815c5cb28fb4fc85690c456f3ec` | `9fecf862deb3cfdc29910be4da6c25e6caa7e52334917f88014363da92d4a81c` |
| Covestack-Cloud-Tools.md | `292245e398e099fb6353ef601edfac428bacf6606f5f774e5817eddd8bb5375d` | `5d95cf2ef656f291ec99d5bb0c220135df3e9f70e274a533f8fe17e128d16f08` |
| Covestack-Cloud-Repository-Layout.md | `955fd27e0fe7159528dd5d44f93b7cda539b5947f9034673da2d42dcdde7b763` | `2ecd4433b861367242961641994c2af65d85b79242ec6303c589bafd3c9c2bfd` |

## 5. How to use the result

1. Follow the workflow first-72-hours table for immediate owner assignments.
2. Open Current Setup Fixes for exact repository/config work and acceptance for each audit ID.
3. Use the corrected setup/technical guides during implementation; the historical audit is evidence, not the current recipe.
4. Review the patch/ZIP if you need any original detail; duplicated learning scripts and obsolete weekly instructions were consolidated intentionally.
5. Close each code finding only with a reviewed implementation and actual acceptance evidence. Documentation correction alone is not a security fix.

## 6. Validation performed

- Internal file links and section anchors resolve; fenced code blocks are balanced.
- Current Setup Fixes contains every audit ID 1–38 exactly once in its implementation tables.
- Dated Friday gates, Saturday half demo and Tuesday freeze/release-candidate/submission dates were checked against the 2026 calendar.
- Proposed Compose YAML parses, including explicit Valkey persistence and MinIO storage at /data. No runtime deployment was performed.
- Applying the exact diff to the archived original six guides in a temporary directory reconstructs the final files byte-for-byte.
- Repository git status retains the previously existing local edits; this task did not modify project source files.
