# CoveStack Cloud — Start Here

| Current status | Meaning |
|---|---|
| Security baseline | Commit `349a2f4` exists. Live authentication/authorization acceptance remains pending. |
| Branding baseline | Commit `7ac2719` exists on `fix/security-baseline`. |
| Lifecycle 1.1.0 | Committed in `189df4f` with documentation improvements. Do not reapply previous patches. |
| Submission | Tuesday, 15 December 2026. |

## Send this package to the team

| Reader | Files to read | Next action |
|---|---|---|
| Everyone | Team-Roles, Project-Workflow, Technical-Design, Current-Setup-Fixes | Agree ownership, gate dates and acceptance evidence. |
| Everyone | `contracts/FROZEN-STATES.md`, `CANONICAL-REQUEST.md`, `internal-api.yaml` | Use the same states, request shapes and signing rules. |
| Anuj (A) | Anuj-Next-Steps, Complete-Working-Setup, `docs/ANUJ-PERSON-C-HANDOFF.md` | Implement Go config/client/polling and integrate C's auth. |
| B | Technical-Design, contracts, Current-Setup-Fixes | Implement main/migrations/scoped services and call Java lifecycle validation. |
| C | Handoff, canonical request, Current-Setup-Fixes | Implement real key/token stores, permissions, revocation and live negative tests. |
| D | Technical-Design, frozen states, API examples | Map internal states into provider DTOs and test actual clients. |
| E | Complete-Working-Setup, Technical-Design, Current-Setup-Fixes | Implement secure bootstrap/CLI and remote doctor. |

## Active guides

| Guide | Purpose |
|---|---|
| [Covestack-Cloud-Team-Roles.md](Covestack-Cloud-Team-Roles.md) | Team Roles |
| [Covestack-Cloud-Project-Workflow.md](Covestack-Cloud-Project-Workflow.md) | Project Workflow |
| [Covestack-Cloud-Technical-Design.md](Covestack-Cloud-Technical-Design.md) | Technical Design |
| [Covestack-Cloud-Complete-Working-Setup.md](Covestack-Cloud-Complete-Working-Setup.md) | Complete Working Setup |
| [Covestack-Cloud-Current-Setup-Fixes.md](Covestack-Cloud-Current-Setup-Fixes.md) | Current Setup Fixes |
| [Covestack-Cloud-Anuj-Next-Steps.md](Covestack-Cloud-Anuj-Next-Steps.md) | Anuj Next Steps |
| [Covestack-Cloud-Renaming.md](Covestack-Cloud-Renaming.md) | Renaming |
| [Covestack-Cloud-Documentation-Cleanup.md](Covestack-Cloud-Documentation-Cleanup.md) | Documentation Cleanup |

## Sharing precautions

| Send | Keep private |
|---|---|
| Team ZIP, reviewed Git branch/PR, contracts and sanitized test results | `.env`, actual access/secret keys, master keys, tokens, database passwords, private SSH keys and credential-bearing kubeconfigs |
| Current active guides | Historical ZIP is reference only; do not follow old setup instructions. |

## Where documentation belongs

| Material | Git / sharing policy |
|---|---|
| Active `docs/`, root README and `contracts/` | Commit and push on the feature branch; review through a PR. Git is the authoritative team source. |
| Team ZIP | Share as a convenience; it can become stale, so point everyone to the branch/PR too. |
| Historical ZIPs | Optional archival evidence; never treat old instructions as current setup. |
| Claude prepared copy and patch ZIPs | Local delivery artifacts; do not copy the entire Claude directory into the project. |
| Credentials | Never commit/share actual `.env`, secrets, master keys or private tokens. |
| Current product spelling | CoveStack. The repository folder and origin URL use `covestack-cloud`; compatibility folders `kan-agent` and `kan-server` remain unchanged. |
