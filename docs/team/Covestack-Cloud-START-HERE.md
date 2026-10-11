# Covestack Cloud — Team Start Guide

Updated 11 October 2026. Product: Covestack Cloud. Intended slug: covestack-cloud.

The security/build baseline is committed as `349a2f4` (Fixed security issues)
on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`.
The working tree was clean when inspected. Remote main/PR merge status has not been
verified; review and merge through the shared PR. Three signature vectors, negative
tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates
and auth-smoke exist. The config.go file currently contains only `package config`;
it is a placeholder, not a working loader. Real credential/token lookup, application
main, migrations/controllers, full agent and live deployment tests remain pending.

Product name: **Covestack Cloud**; intended repository slug: `covestack-cloud`.
The actual checkout and GitHub remote are still named `kan-cloud` until the repository
owner renames them. Use the real paths in commands. Existing module folders,
Go import paths, Java packages, signing headers/algorithm, KAN_* environment names,
API paths, labels, database identity and Valkey key prefixes remain compatibility
identifiers. A future technical rename needs a separate coordinated migration.

## Read first

1. [Roles](Covestack-Cloud-Team-Roles.md): responsibilities and acceptance.
2. [Workflow](Covestack-Cloud-Project-Workflow.md): deadlines through 15 December 2026.
3. [Technical design](Covestack-Cloud-Technical-Design.md): target architecture.
4. [Code changes](Covestack-Cloud-Code-Changes.md): committed baseline/remains.
5. Repository contracts/: executable wire agreement, coordinated review required.

## Start by owner

| Owner | Additional reading | First deliverable |
|---|---|---|
| Anuj/A | Tools, Working Setup, Anuj Next Steps, repo handoff | Implement empty config.go, typed signed client, heartbeat/work polling; CI/VM baseline |
| B | Current Setup Fixes, API/state contracts, handoff | Application/migrations/config, scoped agent controllers/report ordering |
| C | Code Changes, Current Setup Fixes, historical Audit, signing rules/vectors, handoff | Real key/token lookup, encryption/hashing/revocation, permissions and live auth tests |
| D | Design, Current Setup Fixes, examples | Provider schemas and first compute API through B services |
| E | Design, Working Setup, Tools | Secure local bootstrap with C, CLI configuration/doctor and management API |

Use the shared PR for fix/security-baseline. After merge everyone updates main and
creates an owned feature branch. Before merge explicitly base on the published
baseline branch. Do not remove protection or force-push main. This guide does not
assert the PR has merged or remote CI is green.

The documentation bundle contains no credentials. Obtain actual code and JSON/YAML
contracts from the reviewed repository. Do not distribute deploy/.env, deploy/secrets,
passwords/signing secrets or credential-bearing logs. Audit and previous-doc backup
are historical references; active guides describe current work and target design.

## Repository rename coordination

The repository owner can rename the GitHub repository to covestack-cloud. Only after
that succeeds, update remotes and clone paths in team instructions. The new URL is
not asserted to exist. This local branding patch does not rename the hosted repo.
GitHub repository rename, local folder rename and technical protocol migration are
separate operations; coordinate them before updating all team commands.
