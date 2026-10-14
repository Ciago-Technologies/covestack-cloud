# CoveStack V5 — Team Onboarding and Shared Working Rules

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Read first

| Order | Document | Purpose |
|---|---|---|
| 1 | [README](README.md) | V5 index and scope |
| 2 | [Design](DocsV5-DO_Team.md) | Components and boundaries |
| 3 | [Roles](Team_Roles.md) | Owners and pairing |
| 4 | [Workflow](CoveStack-Project-Workflow.md) | New dated gates |
| 5 | [Setup](CoveStack-Complete-Working-Setup.md) | Verified commands and pending prerequisites |
| 6 | [Status](CoveStack-Implementation-Status.md) | What exists versus planned |
| 7 | [Testing](CoveStack-Testing-and-Acceptance.md) | What marks behavior accepted |

Then read feature guides relevant to your task. Java developers work in `kan-server`, Go developers in `kan-agent`. Clone the entire repository; Git pushes repository commits, not a separate folder.

## 2. Branch workflow

```bash
git clone https://github.com/Ciago-Technologies/covestack-cloud.git
cd covestack-cloud
git switch main
git pull --ff-only origin main
git switch -c feat/your-task
```

Use main only after the agreed baseline PRs are merged. Otherwise branch from the exact agreed shared baseline; do not guess merge status from old docs.

```bash
git status --short
git add path/to/your/changed/files
git commit -m "Describe the accepted behavior"
git push -u origin feat/your-task
```

Raise a reviewed PR into main. Stage intended files explicitly. Do not add ignored credentials or someone else's unfinished changes.

## 3. Shared rules

| Rule | Reason |
|---|---|
| Keep server/agent identifiers | Compatibility and imports |
| Common API/services | UI/CLI/IaC must share policy/state |
| Contract changes reviewed | Go/Java/schema agreement |
| No direct main push | Review and required tests |
| Local secrets per teammate | No shared real secret bundles |
| Honest pending state | Mock/schema is not runtime success |
| Actual evidence | Date, commit, versions and result |

## 4. Documentation sharing

Share this V5 set and the source branch/PR. `V5` is currently ignored by the repository's `.gitignore`, so the files will not appear on GitHub automatically. Decide whether to distribute them as a document bundle or deliberately change that ignore policy in a reviewed change. This generation does not change it.

The older `docs/team` timetable is historical for planning; V5 uses 25 January 2027. Runtime `contracts` remains authoritative for current wire behavior. Do not reapply archived delivery patches.

## 5. First team task

Agree owners for new React/provider/Ansible work, confirm the bounded template/cluster/network shape, review current implementation, then accept one real Droplet flow. Complete it before broad success claims or a dashboard filled with mock statuses.
