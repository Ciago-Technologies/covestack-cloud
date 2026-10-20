#!/usr/bin/env python3
"""Validate the controlled V5 Markdown bundle without third-party packages."""
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parent

def cells(line):
    return [s.strip().replace(r'\|', '|') for s in re.split(r'(?<!\\)\|', line.strip().strip('|'))]

def validate():
    errors = []
    tables = 0
    paths = sorted(ROOT.glob('*.md'))
    for path in paths:
        text = path.read_text()
        lines = text.splitlines()
        if not text.startswith('# '):
            errors.append(f'{path.name}: missing title')
        if 'Monday, 25 January 2027' not in text:
            errors.append(f'{path.name}: missing full deadline')
        if 'CloveStack' in text:
            errors.append(f'{path.name}: incorrect branding')
        if '15 December 2026' in text or 'submit Dec 15' in text:
            errors.append(f'{path.name}: stale submission date')
        fenced = False
        i = 0
        while i < len(lines):
            line = lines[i]
            if line.startswith('```'):
                fenced = not fenced
                i += 1
                continue
            if not fenced and line.startswith('|'):
                start = i
                rows = []
                while i < len(lines) and lines[i].startswith('|'):
                    rows.append(cells(lines[i]))
                    i += 1
                tables += 1
                if start > 0 and lines[start - 1].strip():
                    errors.append(f'{path.name}:{start+1}: blank line required before table')
                if i < len(lines) and lines[i].strip():
                    errors.append(f'{path.name}:{i+1}: blank line required after table')
                if len(rows) < 2 or not all(re.fullmatch(r':?-{3,}:?', x) for x in rows[1]):
                    errors.append(f'{path.name}:{start+1}: missing table header delimiter')
                width = len(rows[0])
                if not 2 <= width <= 8:
                    errors.append(f'{path.name}:{start+1}: table too wide for bundle policy')
                if any(len(row) != width for row in rows):
                    errors.append(f'{path.name}:{start+1}: inconsistent table columns')
                continue
            if not fenced:
                for target in re.findall(r'\]\(([^)]+)\)', line):
                    if target.startswith(('https://', 'http://', '#')):
                        continue
                    local = target.split('#')[0]
                    if not (path.parent / local).exists() and not target.startswith('../'):
                        errors.append(f'{path.name}:{i+1}: broken local link {target}')
            i += 1
        if fenced:
            errors.append(f'{path.name}: unclosed code fence')
    audit = (ROOT / 'CoveStack-Current-Setup-Fixes.md').read_text()
    ids = re.findall(r'^\|\s*(\d+)\s*\|', audit, re.M)
    if set(map(int, ids)) != set(range(1, 39)) or len(ids) != 76:
        errors.append('baseline audit: expected all 38 IDs in both remedy and location registers')
    if errors:
        raise SystemExit('\n'.join(errors))
    return paths, tables

if __name__ == '__main__':
    paths, tables = validate()
    words = sum(len(p.read_text().split()) for p in paths)
    print(f'PASS: {len(paths)} Markdown files, {tables} tables, {words:,} words')
    print('Deadline, branding, fences, local links and all 38 baseline audit IDs checked.')
