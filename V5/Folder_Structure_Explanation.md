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

Do not rename module paths, packages, database identities or ownership labels as a documentation cleanup. A technical rename needs a separate migration.

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

## 3. Proposed implementation folders

These paths are future placement decisions; they do not imply files exist today. Keep the current package roots and agree detailed naming before creating code.

```text
kan-agent/internal/
  client/                typed signed API client
  reconcile/             resource work coordination
  driver/compute/        Droplet implementation
  driver/kubernetes/     k3d operations
  network/               owned bridges and leases
  ingress/               approved proxy routes
  metadata/              self-metadata adapter
kan-agent/cmd/kan-agent/  production agent entry point
kan-server/src/main/resources/db/migration/
kan-cli/                 standalone API client
web/                     React SPA
terraform-provider-covestack/
ansible/                 inventory/playbooks/role examples
compat/                  provider compatibility tests
```

| Proposed area | Responsibility | Must not contain |
|---|---|---|
| Java core services | Transactions, ownership, stacks and actions | Duplicate business logic per UI |
| Provider adapters | Map core services to tested external DTOs | Unscoped direct repository calls |
| Go drivers | Inspect and converge owned resources | Arbitrary user shell execution |
| CLI | Authenticated API calls and useful output | Agent signing credentials |
| React | Forms, status, help and safe downloads | Tunnel/daemon credentials |
| Terraform provider | API resource lifecycle and state refresh | Hidden bypass of account APIs |
| Ansible | Approved inventory/configuration | Exported plaintext secrets |
| Evidence | Sanitized results and versions | Private keys/full tokens |

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

Runtime state belongs in PostgreSQL/approved durable stores, not a frontend local array or CLI-only file. Local config files may hold protected credentials but never become the platform's resource registry.

| Material | Placement/policy |
|---|---|
| Local `.env` and secret files | Ignored deployment paths; restrictive permissions |
| Master encryption key | Restricted external secret mechanism; protected recovery copy |
| Terraform state | Protected backend/local file with lock policy; contains sensitive values |
| Kubeconfig | Authorized export; restrictive permissions; preserve existing config |
| Provider/tunnel credentials | Operator-owned restricted secret store |
| Checked-in examples | Placeholder/synthetic data only |

## 6. V5 document map

The six requested replacement guides are this file, `DocsV5-DO_Team.md`, setup, workflow, `Team_Roles.md`, and `tools.md`. Read `README.md` for the additional guides and recommended order. Do not copy archived patches over the current repository.
