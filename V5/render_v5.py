#!/usr/bin/env python3
"""Print HTML for this controlled Markdown subset; not a general Markdown parser."""
from pathlib import Path
import html
import re
from validate_v5 import validate, cells

ROOT = Path(__file__).resolve().parent

def anchor(name):
    return 'doc-' + re.sub(r'[^a-z0-9]+', '-', name.lower()).strip('-')

def inline(text):
    tokens = []
    def hold(value):
        tokens.append(value)
        return '\x00' + str(len(tokens) - 1) + '\x00'
    text = re.sub(r'`([^`]+)`', lambda m: hold('<code>' + html.escape(m[1]) + '</code>'), text)
    def link(m):
        target = m[2]
        if target.split('#')[0].endswith('.md') and (ROOT / target.split('#')[0]).exists():
            target = '#' + anchor(target.split('#')[0])
        return hold('<a href="' + html.escape(target, quote=True) + '">' + html.escape(m[1]) + '</a>')
    text = re.sub(r'\[([^\]]+)\]\(([^)]+)\)', link, text)
    text = html.escape(text)
    text = re.sub(r'\*\*([^*]+)\*\*', r'<strong>\1</strong>', text)
    # A link label may contain a prior code token; restore until no tokens remain.
    while re.search(r'\x00(\d+)\x00', text):
        text = re.sub(r'\x00(\d+)\x00', lambda m: tokens[int(m[1])], text)
    return text

def render(text):
    lines = text.splitlines()
    out = []
    i = 0
    while i < len(lines):
        line = lines[i]
        if not line.strip():
            i += 1
            continue
        if line.startswith('```'):
            language = line[3:].strip()
            i += 1
            code = []
            while i < len(lines) and not lines[i].startswith('```'):
                code.append(lines[i])
                i += 1
            out.append('<pre><code class="language-' + html.escape(language) + '">' + html.escape('\n'.join(code)) + '</code></pre>')
            i += 1
            continue
        if line.startswith('|'):
            rows = []
            while i < len(lines) and lines[i].startswith('|'):
                rows.append(cells(lines[i]))
                i += 1
            klass = ' class="audit"' if rows[0][0] == 'ID' else ''
            head = ''.join('<th scope="col">' + inline(c) + '</th>' for c in rows[0])
            body = ''.join('<tr>' + ''.join('<td>' + inline(c) + '</td>' for c in row) + '</tr>' for row in rows[2:])
            out.append('<table' + klass + '><thead><tr>' + head + '</tr></thead><tbody>' + body + '</tbody></table>')
            continue
        heading = re.match(r'^(#{1,6}) (.+)$', line)
        if heading:
            n = len(heading[1])
            out.append(f'<h{n}>' + inline(heading[2]) + f'</h{n}>')
            i += 1
            continue
        if line.startswith('- '):
            items = []
            while i < len(lines) and lines[i].startswith('- '):
                items.append('<li>' + inline(lines[i][2:]) + '</li>')
                i += 1
            out.append('<ul>' + ''.join(items) + '</ul>')
            continue
        paragraph = [line]
        i += 1
        while i < len(lines) and lines[i].strip() and not lines[i].startswith(('#', '|', '```', '- ')):
            paragraph.append(lines[i])
            i += 1
        out.append('<p>' + inline(' '.join(paragraph)) + '</p>')
    return '\n'.join(out)

if __name__ == '__main__':
    paths, expected_tables = validate()
    first = ['README.md', 'CoveStack-Team-START-HERE.md', 'DocsV5-DO_Team.md', 'CoveStack-Service-Catalog.md', 'CoveStack-Resource-Services.md', 'CoveStack-API-and-Compatibility.md', 'Team_Roles.md', 'CoveStack-Project-Workflow.md', 'CoveStack-Implementation-Status.md', 'Folder_Structure_Explanation.md', 'tools.md', 'CoveStack-Complete-Working-Setup.md', 'CoveStack-Anuj-Next-Steps.md', 'CoveStack-Stack-Templates.md', 'CoveStack-VPC-IPAM-and-Networking.md', 'CoveStack-Publication-Domains-and-Tunneling.md', 'CoveStack-Instance-Metadata.md', 'CoveStack-Terraform-and-Ansible.md', 'CoveStack-UI-CLI-and-Monitoring.md', 'CoveStack-Security-and-ITISS.md', 'CoveStack-Operations-and-Recovery.md', 'CoveStack-Testing-and-Acceptance.md', 'CoveStack-Current-Setup-Fixes.md', 'CoveStack-References-and-Decision-Log.md', 'CoveStack-PDF-and-Formatting.md']
    ordered = [ROOT / name for name in first]
    ordered += [p for p in paths if p not in ordered]
    body = '\n'.join('<section class="document" id="' + anchor(p.name) + '"><p class="document-source">' + html.escape(p.name) + '</p>' + render(p.read_text()) + '</section>' for p in ordered)
    assert body.count('<table') == expected_tables, 'HTML table coverage mismatch'
    result = '<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>CoveStack V5 — Complete documentation</title><style>' + (ROOT / 'pdf-print.css').read_text() + '</style></head><body><p class="screen-only">Print / Save as PDF: A4, scale 100%. Review pagination before sharing.</p>' + body + '</body></html>'
    output = ROOT / 'CoveStack-V5-Print.html'
    output.write_text(result)
    print(f'Created {output}: {len(ordered)} documents, {expected_tables} real HTML tables.')
