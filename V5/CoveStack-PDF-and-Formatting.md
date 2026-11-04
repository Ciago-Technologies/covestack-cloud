# CoveStack V5 — Markdown, PDF and Export Boundaries

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Authoring policy

| Element | Required format |
| --- | --- |
| Comparisons/owners/gates/schema | Actual pipe table with separator, constant column count and blank boundaries |
| Table size | Match actual content; use landscape/reviewed wrapping for wide tables, no claim every converter paginates identically |
| Headings | One active H1 and numbered hierarchy; dated document-control table |
| Examples | Fenced language, labelled actual versus proposal; no fictional executable startup/schema |
| Diagrams | Mermaid source with clear target/evidence labels; requires compatible renderer |
| Links | Relative repository links and heading anchors; stable historical snapshot anchors for old text |
| Status | Exact source/evidence vocabulary, no service support from prose or mock output |

Markdown can represent tables correctly, but a viewer needs GFM/pipe-table support. HTML/PDF render and page layout must be inspected separately; no source text guarantees every renderer's identical result.

## 2. Existing export scripts are not the current authority

The old generated CoveStack-V5-Print.html is a stale historical export, not this newly aligned source bundle. render_v5.py/validate_v5.py/pdf-print.css and HTML remain unchanged in this documentation-only task. No export was regenerated and no PDF/visual rendering claim is made.

Current validator accepts 2–8 columns, but still requires an obsolete fixed academic deadline in every file and exactly 38 old audit IDs twice. The aligned product docs intentionally do not satisfy those retired checks. Renderer calls that validator, so the legacy bundle path remains blocked until a separately authorized tooling update. Do not restore false dates/audit rows or bypass safety validation merely to print.

## 3. Renderer limitations and later tooling task

| Surface | Observed source behavior | Later requirement |
| --- | --- | --- |
| validate_v5.py | Fixed old date/audit semantics and limited link checks | Validate current product/status/authority, complete local links/anchors and structural tables/fences; no false academic requirements |
| render_v5.py | Controlled Markdown subset; Mermaid code display; document link rewriting can discard fragments | Preserve section anchors and external/root targets, render or clearly label diagrams, verify every output link |
| pdf-print.css / browser print | Existing styles and browser-dependent pagination | Inspect real header repetition/wide schema tables/code/diagrams and avoid clipped content |
| Existing generated HTML | Earlier document content | Mark historical until rebuilt from aligned docs with tested tooling |

These are documented implementation/tooling requirements, not script changes performed here. Use installed general Markdown parsing only for structural source checks; it does not certify a final PDF.

## 4. Conversion when separately authorized

After tooling is corrected and output writes are approved, build from the reviewed exact source revision and inspect browser/renderer output. An external installed GFM/Mermaid-aware converter is an alternative; no install or command claiming it is already available is supplied here. Render one large catalogue/schema/gate table first to qualify width/page behavior.

## 5. Sharing with teammates

Share active repository Markdown links first: V5 README, root Starter/team checkpoint and canonical product specs. Include code/runtime revision/evidence where relevant. Do not send stale HTML/PDF as the new architecture or include secrets in exported metadata/screenshots. Historical snapshot is for traceability only. Root Apache-2.0 and upstream notices still apply to distributed material.

## 6. Source verification

Check table column counts, fences, local file/section targets, canonical catalogue classification, exact lifecycle strings, status/roles/ownership, current VM/RustFS scope and historical preservation. Application/runtime acceptance is separate and was not run by this documentation update. Full source-validation details belong in the audit/checkpoint, not a fabricated successful legacy renderer transcript.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-pdf-and-formatting) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
