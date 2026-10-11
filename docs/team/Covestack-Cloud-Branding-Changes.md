# Covestack Cloud — Branding and Documentation Changes

Updated 11 October 2026. Requested product name: **Covestack Cloud**.

## What changed

All 12 top-level Markdown guides have updated branding/names/references. Current
status is corrected to committed baseline 349a2f4 on published fix/security-baseline.
Old statements that the original checkout is untouched, has only two vectors or
only a secrets CI job are removed from active onboarding. Historical audit/rewrite
records are labeled. Anuj's config.go is correctly described as an empty placeholder.
Owner priorities, remaining work and team sharing instructions are updated.

A fresh source copy at `Claude/covestack-cloud` includes the latest original commit,
updated README/handoff/placeholder docs, all team guides under docs/team, Maven
project display name and OpenAPI display title. This copy contains no deploy
secrets/.env or Git history. Use the patch for the real repository.

## Apply to the existing branch

The original checkout is read-only to this assistant session. Original files,
branch, remote and commit were not changed by this branding task. From Fedora:

```bash
cd /run/media/anujtiwari/DevMount/Project-CDAC/kan-cloud
git apply --check /run/media/anujtiwari/DevMount/Claude/Covestack-Cloud-Branding.patch
git apply /run/media/anujtiwari/DevMount/Claude/Covestack-Cloud-Branding.patch
git diff --stat
git add README.md contracts kan-server/pom.xml docs agent/README.md control-plane/README.md deploy/README.md
git commit -m "Brand Covestack Cloud and refresh team documentation"
git push
```

Current branch is fix/security-baseline with upstream configured. These commands
add a follow-up commit to that PR; they do not push main. Review before committing.
Do not reapply the old security patch, which is already part of baseline 349a2f4.

## Compatibility and repository naming

Existing kan-agent/kan-server folders, Go module/imports, Java packages, KAN_*
configuration, X-Kan-* headers, KAN1-HMAC-SHA256, kan.principal, kan:nonce, kan.io
labels, database/volume names and /kan/v1 routes are preserved. No data migration
or resigning is necessary for this display-name update. A technical rename is a
separate coordinated change; preserve interoperability for currently configured clients.

The GitHub repository/local source directory remain kan-cloud. Intended future slug
is covestack-cloud; an owner must rename the GitHub repository before changing its
remote URL. The new URL has not been created/verified here. Product naming does not
establish domain/trademark availability.

## Guide renames and resulting lines

Lines refer to the renamed current guide. Archived originals are preserved in
Covestack-Cloud-Docs-Before-Rename-2026-10-11.zip. Earlier security patch/archive
filenames retain their historical identities.

| Previous filename | Current filename | Changed lines | Reason |
|---|---|---|---|
| `Kan-Cloud-Documentation-Change-Log.md` | `Covestack-Cloud-Documentation-Change-Log.md` | 1–4, 8–8, 12–12, 16–21, 29–80, 86–88, 97–102 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Kan-Cloud-Current-Setup-Fixes.md` | `Covestack-Cloud-Current-Setup-Fixes.md` | 1–1, 3–3, 5–12, 14–14, 16–16, 18–22, 24–27, 29–30, 43–43, 56–56, 58–58, 61–62, 188–188, 193–200, 202–205 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Kan-Cloud-Project-Workflow.md` | `Covestack-Cloud-Project-Workflow.md` | 1–1, 5–5, 9–23 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Kan-Cloud-Audit-2026-10-11.md` | `Covestack-Cloud-Audit-2026-10-11.md` | 1–6, 114–114, 152–152, 170–170, 176–176, 182–182, 194–194, 206–206, 212–212 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Kan-Cloud-Code-Changes.md` | `Covestack-Cloud-Code-Changes.md` | 1–1, 3–3, 5–12, 14–17, 19–21, 31–31, 55–55, 76–76, 79–79, 98–99 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Kan-Cloud-Complete-Working-Setup.md` | `Covestack-Cloud-Complete-Working-Setup.md` | 1–1, 3–19, 31–31, 47–47, 174–174, 270–270 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `DocsV4-DO_Team.md` | `Covestack-Cloud-Technical-Design.md` | 1–1, 5–5, 9–23 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Team_Roles_KANSP_CloudV2.md` | `Covestack-Cloud-Team-Roles.md` | 1–1, 3–3, 19–19, 38–38, 58–58, 124–124 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `tools.md` | `Covestack-Cloud-Tools.md` | 1–1, 3–3, 92–92, 108–108 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Folder_Stucture_Explanation.md` | `Covestack-Cloud-Repository-Layout.md` | 1–1, 3–3, 5–10, 13–14, 16–16, 18–18, 20–20, 22–32, 36–36, 38–42, 44–46, 48–50, 52–56 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Team-START-HERE.md` | `Covestack-Cloud-START-HERE.md` | 1–1, 3–19, 23–27, 29–29, 31–37, 39–42, 44–47, 49–49, 51–55 | Product branding, links and current baseline status; historical evidence labeled where applicable. |
| `Anuj-Next-Steps-2026-10-11.md` | `Covestack-Cloud-Anuj-Next-Steps.md` | 11–11, 20–22, 27–27, 33–33, 144–144 | Product branding, links and current baseline status; historical evidence labeled where applicable. |

## Repository display-name changes

| File | Problem | Result |
|---|---|---|
| README.md | Old product title and no current onboarding | Covestack Cloud title, compatible identifiers, current baseline and team guide link |
| contracts/internal-api.yaml:3 | Old API display title | Covestack Cloud — Internal Agent API |
| contracts/CANONICAL-REQUEST.md:1 | Product branding missing | Covestack Cloud heading with current kan sigv1 protocol identity |
| contracts/README.md:1 | Generic/old branding | Covestack Cloud contracts heading |
| kan-server/pom.xml:17 | Old Maven display name | Covestack Cloud Server; artifactId unchanged |
| docs/ANUJ-PERSON-C-HANDOFF.md | Unapplied-patch instructions were stale | Committed baseline, current work and compatibility note |
| docs/README.md | Empty | Onboarding/handoff links |
| agent/README.md and control-plane/README.md | Empty | Clearly labeled placeholders and actual active modules |
| deploy/README.md | Empty | Correct Compose env-file command and team setup link |
| docs/team/*.md | Team guides lived outside repo | All updated guides available to reviewers/teammates via Git |

## Checks

Patch dry-check and byte-exact application compared with the original baseline.
Markdown relative file links checked for existence, XML/YAML display-name edits
parsed, and shared Python vectors passed. Protocol/code semantics are unchanged.
Real Maven, Go 1.27.2, remote CI and live deployment acceptance retain their prior
pending status. No GitHub repository rename, commit or push performed by assistant.
