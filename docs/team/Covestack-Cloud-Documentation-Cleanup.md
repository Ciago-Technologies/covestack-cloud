# Covestack Cloud — Documentation Cleanup

| Current status | Meaning |
|---|---|
| Security baseline | Commit `349a2f4` exists. Live authentication/authorization acceptance remains pending. |
| Branding baseline | Commit `7ac2719` exists on `fix/security-baseline`. |
| Lifecycle 1.1.0 | Applied to the working tree, not committed at inspection. Do not reapply previous patches. |
| Submission | Tuesday, 15 December 2026. |


| File | Current line | Problem | Solution |
|---|---|---|---|
| `README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `agent/README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `contracts/CANONICAL-REQUEST.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `contracts/FROZEN-STATES.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `contracts/README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `control-plane/README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `deploy/README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/ANUJ-PERSON-C-HANDOFF.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/README.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Anuj-Next-Steps.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Complete-Working-Setup.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Current-Setup-Fixes.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Project-Workflow.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Renaming.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-START-HERE.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Team-Roles.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Technical-Design.md` | 1 | Dense prose and mixed presentation | Tables for prose and tasks; code blocks and existing detailed tables preserved. |
| `docs/team/Covestack-Cloud-Tools.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Complete-Working-Setup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Repository-Layout.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Complete-Working-Setup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Audit-2026-10-11.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Documentation-Cleanup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Code-Changes.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Documentation-Cleanup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Branding-Changes.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Documentation-Cleanup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Documentation-Change-Log.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Documentation-Cleanup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Lifecycle-Changes.md` | Deleted | Duplicate or historical active guide | Use `Covestack-Cloud-Documentation-Cleanup.md`; original retained in history ZIP. |
| `docs/team/Covestack-Cloud-Lifecycle-Reference.md` | Deleted | Duplicate or historical active guide | Authoritative lifecycle is `contracts/FROZEN-STATES.md`. |
| `contracts/generate_lifecycle.py` | 149 | Regeneration could restore prose and duplicate reference | Generate the tabular authoritative reference only. |

| Verification / limitation | Result |
|---|---|
| Preservation | All original project Markdown retained in `docs/archive/Documentation-Before-Tables-2026-10-11.zip`. All 38 audit rows remain in Current-Setup-Fixes. |
| Actual project | Read-only to this session; apply the checked patch there. Existing lifecycle working changes are the patch baseline. |
| Scope | Documentation presentation and Markdown generator output only; runtime security/controller work remains pending. |

## Readability follow-up

| File | Line / section | Problem | Solution |
|---|---|---|---|
| All active Markdown | Table boundaries | Missing blank lines could prevent table rendering. | Add blank lines around tables without altering fenced examples. |
| `contracts/FROZEN-STATES.md` | Resource lifecycle vocabulary | Seven very wide rows mixed service identity and long observed-state lists. | Separate service/desired/terminal summary from observed states, one state per row. |
| `contracts/CANONICAL-REQUEST.md` | Whole guide | Numbered paragraph tables hid individual signing rules and lacked a worked example. | Separate terms, headers, canonical components, formulas, verification and checks; add exact validated GET example and body fixture reference. |
| `contracts/generate_lifecycle.py` | `build_markdown` / `tabular_markdown` | Regeneration must preserve improved layout. | Generate split vocabulary tables and blank-line boundaries. |
| Claude older copies | Deleted directories | Multiple generated source copies caused confusion. | Archive and remove `kan-cloud-fixed`, `covestack-cloud`, `covestack-cloud-lifecycle`; retain only `covestack-cloud-tabular`. |
| START-HERE | Where documentation belongs | Unclear whether to push or share. | Commit active documentation through the PR; use ZIP only as a convenience. |
