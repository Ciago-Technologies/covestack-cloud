# CoveStack V5 — Markdown and PDF Export

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Formatting policy

| Element | Convention |
|---|---|
| Comparison/ownership/status | Real pipe table with header separator |
| Table boundary | Blank line before and after |
| Table width | Usually two/three columns; long matrices split |
| Highlight | Bold decisions, deadline and critical status |
| Commands/examples | Fenced code with language; examples labeled |
| Diagram | Plain-text code diagram for portable printing |
| Headings | Hierarchical titles with whitespace |
| Dates | Full year and consistent schedule |

Tables are not enclosed in code fences. A Markdown viewer must enable pipe-table support (GFM/Pandoc style). If the viewer treats a valid pipe table as plain text, that is a renderer configuration issue rather than a different table syntax.

## 2. Included print path

`render_v5.py` uses the controlled Markdown subset in these documents and creates `CoveStack-V5-Print.html`. It renders actual HTML tables, combines the documents in reading order, and embeds the print stylesheet. It does not fetch remote pages, execute code blocks or convert every arbitrary Markdown extension.

From the repository root:

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

Keep the CSS next to the generated HTML. Print that page through a browser. Direct PDF generation requires an installed compatible PDF engine; none is assumed by this bundle. Diagrams use plain text so a Mermaid extension is unnecessary.

## 4. Validation

```bash
python3 V5/validate_v5.py
```

This checks table row shapes/header delimiters, blank boundaries, fences, local Markdown links, metadata and deadline consistency. It is document validation, not application testing. The HTML renderer also checks its generated table count against the Markdown inventory.

## 5. Sharing

Send the V5 Markdown directory or generated print HTML/PDF as documentation. Do not include deployment secrets, Terraform state or real kubeconfigs. V5 is currently ignored by Git; share it deliberately or change that policy through a reviewed task if wanted. No automatic commit/push is performed.
