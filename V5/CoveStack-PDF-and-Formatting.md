# CoveStack V5 — Markdown and PDF Export

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Formatting policy

| Element                     | Convention                                                                                   |
| --------------------------- | -------------------------------------------------------------------------------------------- |
| Comparison/ownership/status | Real pipe table with header separator                                                        |
| Table boundary              | Blank line before and after                                                                  |
| Table width                 | Match required comparison/gate/schema columns; wide tables need landscape or reviewed layout |
| Highlight                   | Bold decisions, deadline and critical status                                                 |
| Commands/examples           | Fenced code with language; examples labeled                                                  |
| Diagram                     | Text layout or Mermaid ERD; use Mermaid-aware rendering for the schema                       |
| Headings                    | Hierarchical titles with whitespace                                                          |
| Dates                       | Full year and consistent schedule                                                            |

Tables are not enclosed in code fences. A Markdown viewer must enable pipe-table support (GFM/Pandoc style). If the viewer treats a valid pipe table as plain text, that is a renderer configuration issue rather than a different table syntax.

## 2. Included print path

**The existing generated HTML is a historical snapshot, not the current reviewed bundle.** Scripts/HTML/CSS remain unchanged in this docs-only task.

`render_v5.py` uses the controlled Markdown subset in these documents and creates `CoveStack-V5-Print.html`. It renders actual HTML tables, combines the documents in reading order, and embeds the print stylesheet. It does not fetch remote pages, execute code blocks or convert every arbitrary Markdown extension.

The existing validator enforces at most four columns, while approved gate/acceptance matrices have more. Therefore the unchanged render command is blocked by its validator until a separately reviewed tooling change. Do not claim a fresh current PDF/HTML was generated here. The renderer also displays Mermaid as code and discards section fragments when rewriting document links.

Historical command, for use only after those tooling limitations are addressed:

```bash
python3 V5/render_v5.py
```

Open the generated HTML in your browser, then Print / Save as PDF. Choose A4, scale 100%, and review preview. The stylesheet wraps table/code content, repeats table headers where the browser supports it, and starts each document on a new page. A very long cell/row may still require renderer-specific pagination adjustment; no Markdown can guarantee identical output in every converter.

## 3. Alternative conversion

If Pandoc is installed, individual documents can be converted to printable HTML:

```bash
pandoc -f gfm -t html5 -s V5/DocsV5-DO_Team.md \
  --css pdf-print.css -o V5/design-print.html
```

Keep the CSS next to the generated HTML. Print that page through a browser. Direct PDF generation requires an installed compatible PDF engine; none is assumed by this bundle. The schema ERD is Mermaid relationships-only; choose an existing Mermaid-aware converter, or render that diagram separately in a later authorized task. No converter/install is assumed here.

## 4. Validation

```bash
python3 V5/validate_v5.py
```

The unchanged script checks table row shapes/header delimiters, blank boundaries, fences, local Markdown links, metadata and deadline consistency, but its four-column cap predates the required wider tables. Use independent Markdown shape/link/fence checks for this review; a script failure on approved width is a tooling limitation, not a passed check. It is document validation, not application testing. The HTML renderer also checks its generated table count against the Markdown inventory.

## 5. Sharing

Send the V5 Markdown directory or generated print HTML/PDF as documentation. Do not include deployment secrets, Terraform state or real kubeconfigs. V5 is tracked in Git; share reviewed repository links to the published branch/revision, or a dated Markdown bundle. Do not share the stale HTML as current. Ignore rules remain untouched. Publication requires a deliberate commit/push of reviewed changes.
