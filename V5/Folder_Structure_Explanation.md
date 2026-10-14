# CoveStack V5 — Folder and File Responsibilities

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Naming rules

| Name | Meaning |
|---|---|
| CoveStack | Product name; preserve capitalization |
| `covestack-cloud` | Repository slug |
| `kan-server` | Existing Java server module |
| `kan-agent` | Existing Go agent module |
| `KAN_*`, `KAN1`, `X-Kan-*` | Existing configuration/signing compatibility identifiers |
| V5 | New documentation set; not a protocol bump |

Do not rename module paths, packages, database identities as a documentation cleanup. Existing kan.io/* fixture labels are observations; kan.* is a proposed reserved-label contract, not existing compatibility. A technical rename needs a separate migration.

## 2. Existing tree

```text
covestack-cloud/
  V5/                    new documentation set
  docs/team/             older active guides; historical schedule
  Version Docs/          archived copies and delivery patches
  contracts/             authoritative wire/lifecycle fixtures
  kan-server/            Java authentication/core libraries
  kan-agent/             Go config/signing/lifecycle/smoke
  deploy/                local infrastructure configuration
  .github/workflows/     builds and secret scanning
  agent/                 placeholder documentation
  control-plane/         placeholder documentation
```

| Location | Current purpose | Rule |
|---|---|---|
| `kan-agent/internal/config` | Local environment validation | Never log credential fields |
| `kan-agent/sigv1` | Shared rules and outgoing signing | Preserve authenticated bytes |
| `kan-agent/internal/lifecycle` | Generated catalog and validation | Edit generator source, not output |
| `kan-agent/cmd/auth-smoke` | Real-server signing checks | Mock tests are separate evidence |
| `kan-server/.../auth` | Authentication interfaces/filters | Concrete stores still required |
| `kan-server/.../contract` | Shared lifecycle helpers | Runtime integration is additional |
| `contracts` | Schemas, vectors, generated-state inputs | Coordinated Go/Java changes |
| `deploy` | Private infra and secret templates | Local secrets untracked |

## 3. Planned implementation folders — no code moves

These are target folders, not source-presence or gate claims. This batch creates no directories, .gitkeep files, module files or code. Module layout and signer relocation need A+B+C review before implementation.

```text
covestack-cloud/
  contracts/                     existing authoritative fixtures
  pkg/kanclient/                 planned signer + HTTP API client (A)
  cmd/covestack/                 planned thin cobra CLI (E)
  kan-agent/                     existing module; A container/k3d/reconcile
  kan-server/                    existing Java module; B/C/D
  terraform-provider-covestack/  planned thin provider (E, after CLI)
  ansible/                       planned Part 1 site.yml/roles; inventory later
  deploy/                        existing private infrastructure
  docs/                          existing documentation
```

| Target area | Responsibility | Boundary |
|---|---|---|
| pkg/kanclient | Shared signer and HTTP API client; A owns changes | Agent/CLI/provider import reviewed client; module layout TBD |
| cmd/covestack | E cobra thin API client, JSON and packaging | No Docker/k3d/DB; no local resource DB |
| kan-agent | A container driver, k3d driver and reconcile | Docker network helper inside container driver; only two workload drivers |
| kan-server | B core/all SQL; C auth/tests/IPAM; D API/DTO | Account predicates; bucket/credential reconcile in Java |
| terraform-provider-covestack | E later thin API lifecycle layer | Optional Droplet/K8s only; A review |
| ansible | A Part 1 reviewed operator installer | roles/docker, infra, server, agent; no tenant inventory in CORE-5 |
| Extended folders | Read-only UI/inventory and proxy-config helpers | Placement TBD; [roadmap](CoveStack-Extended-Roadmap.md), after entry gates |
| Evidence | Sanitized tests, actual versions and restore results | No private keys/full tokens |

The former internal/client, kan-cli and separate network-driver layout are superseded by the reviewed shared-client/thin-CLI/two-driver target. No existing code is moved. Production agent entry point is a planned task, not a fictional startup instruction. Java migrations belong to B's later reviewed work; this docs task creates none.

## 4. Generated versus handwritten files

| File group | How to update |
|---|---|
| `lifecycle-rules.json` | Reviewed source-of-truth change |
| `internal-api.base.json` / generator | Coordinated schema/description change |
| Go/Java catalog outputs | Regenerate together |
| `FROZEN-STATES.md` / API output | Regenerate and compare |
| Signing vectors | Cross-language agreed fixture update |
| V5 Markdown | Documentation edits; do not change runtime by implication |

## 5. Secrets and state placement

Runtime state belongs in PostgreSQL/approved durable stores, not a frontend local array or CLI-only file. CLI config contains server URL/key ID/profile only; secrets come from env or OS keychain. No client file becomes the resource registry.

| Material | Placement/policy |
|---|---|
| Local `.env` and secret files | Ignored deployment paths; restrictive permissions |
| Master encryption key | Restricted external secret mechanism; protected recovery copy |
| Terraform state | Protected backend/local file with lock policy; contains sensitive values |
| Kubeconfig | Authorized export; restrictive permissions; preserve existing config |
| Provider/tunnel credentials | Operator-owned restricted secret store |
| Checked-in examples | Placeholder/synthetic data only |

## 6. V5 document map

Read the main design, SRS scope, service gates and acceptance first; then schema, workflow, roles, setup and Go handover. Topic guides now separate CORE-5 from roadmap pointers; preserved historical text must not override CORE-5. Do not copy archived patches over the current repository.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
