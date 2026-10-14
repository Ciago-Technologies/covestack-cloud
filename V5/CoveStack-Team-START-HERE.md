# CoveStack V5 — Team Onboarding and Shared Working Rules

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Read first

Start from the [V5 index and teammate reading order](README.md#4-reading-order-and-teammate-sharing). Everyone reads:

| Order | File | Purpose |
|---|---|---|
| 1 | [SRS scope](CoveStack-SRS-Scope.md) | Exactly five services and cross-cutting deliverables |
| 2 | [Design](DocsV5-DO_Team.md) | Control plane, not direct runtime handlers |
| 3 | [Roles](Team_Roles.md) | A agent Go; E thin CLI; B/C/D Java boundaries |
| 4 | [Workflow](CoveStack-Project-Workflow.md) | CORE-5 schedule to Monday, 25 January 2027 |
| 5 | [Setup](CoveStack-Complete-Working-Setup.md) | Baseline checks, target installation and Known constraints |
| 6 | [Gates](CoveStack-Service-Gates.md) | G1–G4 promotion evidence |
| 7 | [Acceptance](CoveStack-AntiWrapper-Acceptance.md) | Tests 1–15 and CORE-5 viva |
| 8 | [Status](CoveStack-Implementation-Status.md) | No service promoted from source presence |

Then read assigned schema/API/Go/security guides from the index. Clone the entire repository. Java developers work in kan-server; A works in kan-agent/pkg/kanclient; E in planned cmd/covestack and later thin provider. Git pushes repository commits, not into a service folder.

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
| Common API/services | CLI/provider/inventory authenticated API only; no Docker/k3d/DB |
| Contract changes reviewed | Go/Java/schema agreement |
| No direct main push | Review and required tests |
| Local secrets per teammate | No shared real secret bundles |
| Honest service status | Planned / In progress / Gates passed / Extended track only; mocks do not pass gates |
| Actual evidence | Date, commit, versions and result; C owns all acceptance tests |
| CLI order | droplet → vpc → k8s → db → bucket → stack; previous failure test in CI first |
| Kubernetes | Required; max 2 clusters/host, 1 server + max 2 agents |
| Credentials/polling | service signing versus hash-only API tokens; multi-account polls TBD B+C+A |
| Go learning | E CLI only; A types/explains agent Go; any Go change reviewed by A |

## 4. Documentation sharing

Share the reviewed whole V5 directory and exact repository branch/revision; README lists essential files by owner. V5 is tracked in checkpoint 40291e6; subsequent edits are local until deliberately committed/pushed. No ignore policy change is needed or made here. Do not tell teammates to clone an unpublished local revision.

Older docs/team schedules are historical; V5 deadline is Monday, 25 January 2027. Contracts retain lifecycle 1.1.0 authority. Use source/template links, not separate stale ZIP source copies; each teammate creates private configuration. Never share real credentials, kubeconfigs, Terraform state or deployment secrets. The existing print HTML is stale.

## 5. First team task

| Owner | First agreed slice |
|---|---|
| A | Review config/client/module boundary; signed poll main and real auth-smoke; agent account/assignment decision with B+C |
| B | Server main/configuration; schema review and staged Flyway proposal; all lease/fence/idempotency SQL |
| C | Concrete credential/auth lookup; acceptance tests from day 1; allocator/security review |
| D | Java API/DTO schema review and /v2 subset fixtures; no Go |
| E | W1–W4 Go handover; thin cobra CLI/API mocks and cross-builds; no agent work |

Confirm the shared baseline and reviewers, then deliver one real Droplet lifecycle through API/CLI before advancing. No UI work during CORE-5. Extend only under the roadmap entry rule; if A is one week behind, stop all extended work. Mentor ITISS coding confirmation remains pending.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
