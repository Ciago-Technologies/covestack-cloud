# CoveStack — Teammate Start Guide

| Document control           | Value                                                               |
| -------------------------- | ------------------------------------------------------------------- |
| Updated                    | 15 October 2026                                                     |
| Audience                   | A — Anuj, B, C, D and E                                             |
| Submission deadline        | Monday, 25 January 2027                                             |
| Repository                 | https://github.com/Ciago-Technologies/covestack-cloud               |
| Current planning authority | V5 documentation and reviewed team decisions                        |
| Implementation status      | Planned targets; source presence is not service acceptance evidence |

## 1. What you are joining

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

| CORE-5 service     | Required boundary                                                                               |
| ------------------ | ----------------------------------------------------------------------------------------------- |
| Compute/Droplets   | Containers on one host; never VMs                                                               |
| Kubernetes         | Required; max 2 clusters per host; 1 server + max 2 agents per cluster; one tested version pair |
| Managed PostgreSQL | Backup/restore, scoped credentials and private networking                                       |
| Object storage     | S3-compatible MinIO; per-account credentials and bucket scoping                                 |
| VPC/IPAM           | One private network per account; simple allocator; no custom CIDR                               |

Security, Recovery, CLI, optional Terraform, one fixed PostgreSQL + MinIO + app Droplet stack, and Ansible Part 1 are cross-cutting work, not additional services. Optional Terraform covers Droplets/Kubernetes only. No UI work during CORE-5. Deferred requirements belong to the [extended roadmap](V5/CoveStack-Extended-Roadmap.md), not the SRS deliverables.

Clone the **whole repository**. Java developers work inside kan-server; A owns kan-agent; E owns the planned CLI. Nobody pushes into a folder: Git pushes repository commits on a branch.

## 2. Before cloning — confirm access and baseline

Anuj shares the repository URL, agreed baseline branch/revision, role letter and first task. Confirm GitHub access and the PR target. Do not infer merge status from this guide.

| Situation                               | Baseline to use                                                                                                                          |
| --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| Reviewed baseline is merged into main   | Updated origin/main                                                                                                                      |
| Baseline PR is still pending            | The exact shared branch Anuj confirms; feat/agent-config was used for this documentation work, but do not assume it remains the baseline |
| Your task depends on another pending PR | Agree the dependency branch and review target with its owner before coding                                                               |

Use one agreed baseline across the team. A missing V5 folder can mean you cloned a branch without the documentation; check the branch before recreating files or copying ZIP source over the checkout.

## 3. Tools to prepare

Install tools appropriate to your role using your operating system's normal setup. This guide does not select new versions or package-manager commands. Match [tools](V5/tools.md), the checked-out manifests and CI, and record actual versions.

| Who                               | Tools needed to start                                                    | Check                                                          |
| --------------------------------- | ------------------------------------------------------------------------ | -------------------------------------------------------------- |
| Everyone                          | Git, editor, repository/GitHub access                                    | git --version                                                  |
| A and E                           | Go matching reviewed go.mod/CI; Git                                      | go version                                                     |
| B, C and D                        | JDK matching pom.xml/CI; Maven                                           | java -version; javac -version; mvn -version                    |
| Contract/integration contributors | Python 3 and existing contract checker prerequisites                     | python3 --version                                              |
| Local infrastructure/integration  | Docker Engine/Desktop and Compose with reviewed network/VM configuration | docker version; docker compose version                         |
| A later                           | Tested k3d+k3s pair and Ansible toolchain                                | Select/pin only after host spike/installer review              |
| E later                           | Terraform/framework                                                      | Optional provider work after CLI; not required for W1 learning |

Inspected repository inputs are Go toolchain 1.27.2, Java 25 and control-plane PostgreSQL 17; verify the actual checkout before installing/changing versions. No Maven wrapper is assumed. If a dependency/toolchain cannot be downloaded, record the blocker; do not call an incomplete verification a pass.

Shell examples use Bash/Git Bash. Windows PowerShell users should use their equivalent syntax or Git Bash. Application resource creation remains a target until the relevant main/API/CLI exists.

## 4. Clone and inspect the repository

Run from the directory where you keep projects:

```bash
git clone https://github.com/Ciago-Technologies/covestack-cloud.git
cd covestack-cloud
git remote -v
git status -sb
git fetch origin
git branch -r
```

If you already have a checkout, start with git status -sb and preserve your work. Do not clone over an existing directory, reset changes, or reapply historical patches.

Set the baseline to the value Anuj confirmed. The following main value is correct **only after the baseline merge**; otherwise replace it with the agreed branch name before continuing:

```bash
COVESTACK_BASELINE_BRANCH=main
```

This variable is a local shell convenience, not a CoveStack application setting. Commands below use the fetched origin branch directly so an old local main does not become your starting point.

## 5. Create your own feature branch

Choose **one** first-task branch. These names are suggestions, not existing branches or protocol identifiers.

| Role | Suggested first branch   | First slice                                                        |
| ---- | ------------------------ | ------------------------------------------------------------------ |
| A    | feat/a-agent-client      | Shared-client module review and single-account agent client        |
| B    | feat/b-server-foundation | Server configuration/main and reviewed schema/migration foundation |
| C    | feat/c-auth-acceptance   | Concrete credentials/authentication and acceptance foundation      |
| D    | feat/d-api-contracts     | Java DTO/API inventory and reviewed fixtures                       |
| E    | feat/e-cli-foundation    | Cobra learning/client skeleton and packaging                       |

For example, B runs:

```bash
git switch -c feat/b-server-foundation "origin/$COVESTACK_BASELINE_BRANCH"
git status -sb
```

A, C, D and E substitute their own branch name. Do not run all five branch-creation examples, reuse someone else's feature branch, or push directly to main. If your branch already exists, switch to it instead of trying to create it again.

## 6. Read these files before implementation

| Order | Required reading                                               | What to take away                                            |
| ----- | -------------------------------------------------------------- | ------------------------------------------------------------ |
| 1     | [V5 index](V5/README.md)                                       | Current document map                                         |
| 2     | [SRS scope](V5/CoveStack-SRS-Scope.md)                         | Exactly five services and deliverables                       |
| 3     | [Main design](V5/DocsV5-DO_Team.md)                            | Authority, reconciliation and agreed host/account model      |
| 4     | [Team roles](V5/Team_Roles.md)                                 | Your boundaries and required reviewers                       |
| 5     | [Workflow](V5/CoveStack-Project-Workflow.md)                   | Dependencies, dated checkpoints and freeze                   |
| 6     | [Setup](V5/CoveStack-Complete-Working-Setup.md)                | Baseline commands, target installation and Known constraints |
| 7     | [Service gates](V5/CoveStack-Service-Gates.md)                 | Evidence before service promotion                            |
| 8     | [Acceptance](V5/CoveStack-AntiWrapper-Acceptance.md)           | Exact tests 1–15; C owns acceptance                          |
| 9     | [Implementation status](V5/CoveStack-Implementation-Status.md) | Open work and decisions; no assumed live platform            |

For integration, read [contract overview](contracts/README.md), [internal API](contracts/internal-api.yaml), [lifecycle reference](contracts/FROZEN-STATES.md) and [signing reference](contracts/CANONICAL-REQUEST.md). Contract 1.1.0 wins on wire strings. Missing values are TBD with B; new messages require coordinated schemas, fixtures and generated outputs.

Historical team/handoff ZIPs in [docs/archive](docs/archive/README.md) preserve earlier work. Their December schedule, role assignments, paths and commands are not active instructions. The old combined V5 HTML is also a historical snapshot.

## 7. Run baseline checks for your part

From the repository root, contract contributors can run:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
```

A/E run existing agent checks from the actual module; this does not assert that the future CLI shares that module:

```bash
cd kan-agent
go test -race ./...
go vet ./...
cd ..
```

B/C/D run Java verification:

```bash
cd kan-server
mvn --batch-mode --no-transfer-progress verify
cd ..
```

Record revision, command, actual tool versions and result. Race checks need a compatible local C compiler/platform setup; record environment blockers rather than silently skipping them. Builds may need dependency downloads. A library/test pass is not a server boot, Docker lifecycle or tenant-isolation pass.

**Do not run a fictional server or CLI.** B must first add and verify the application entry point/configuration. A must add the production agent executable. E must add the CLI in the reviewed module layout. Publish startup commands only after those artifacts exist and have been tested.

## 8. Prepare infrastructure only when your task needs it

A/B/C coordinate shared integration on the AlmaLinux VM. Teammates may use approved local dependencies for their tests; do not expose the shared DB/cache to the internet or assume a remote VM's localhost is your laptop.

| File / dependency                | What to do                                                                                       |
| -------------------------------- | ------------------------------------------------------------------------------------------------ |
| deploy/.env                      | Create from the tracked template only if absent; use reviewed actual image references            |
| deploy/secrets/postgres_password | Generate a private local database password                                                       |
| deploy/secrets/valkey.conf       | Copy the reviewed template and set a private credential                                          |
| deploy/secrets/agent.env         | A configures approved service credentials supplied through C's reviewed process                  |
| Java configuration               | B/C configure datasource and replay-store access separately using the reviewed private mechanism |

Follow [Setup sections 5–7](V5/CoveStack-Complete-Working-Setup.md#5-private-infrastructure-configuration) for exact preparation. Never overwrite existing secrets or put real values in Git, chat, screenshots or PRs.

After required files and exact image references exist, from repository root:

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
docker compose --env-file deploy/.env -f deploy/docker-compose.yml ps
```

Compose starts private dependencies, not the unfinished Java server/full agent. Preserve volumes; do not use down -v or blanket prune for repair. A changed password file does not rotate credentials in an initialized database volume. E can start against a mock API without running the infrastructure.

## 9. A — Anuj: start here

| Step | Work / location                                                    | Evidence and dependency                                                                                                |
| ---- | ------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------- |
| 1    | Review kan-agent/internal/config/, kan-agent/sigv1/ and contracts/ | Explain configuration, authenticated bytes and scope; agree credential provisioning/assignment with B+C                |
| 2    | Plan pkg/kanclient/ signer/HTTP client and module layout           | A+B+C review before relocation; no new root package assumption                                                         |
| 3    | Build client and production agent coordination in kan-agent/       | One account first; finite timeout/cancellation, bounded responses, no redirects, fresh signed retries                  |
| 4    | Generic container driver and reconcile                             | Tests 1–6 with C/B: one resource after crash, fenced stale owner, no outage cleanup                                    |
| 5    | Add a second account context                                       | One process/host, separate immutable credential/client/poll context per account; one auth failure does not stop others |
| 6    | Docker network helper inside container driver                      | C allocator/B records; test 14; no invented VPC wire kind                                                              |
| 7    | One-day host spike, then k3d CLI driver                            | Tested pair/flags; max 2 clusters/host; tests 7–10; scoped kubeconfig/owned cleanup                                    |
| 8    | PostgreSQL/MinIO process and volume lifecycle                      | Generic container inputs; tests 11–13; Java owns bucket/credential reconcile                                           |
| 9    | Fixed-stack execution, Ansible Part 1, CI and ITISS evidence       | Test 15, repeat install changed=0, failure/restore/security evidence                                                   |

Read [Anuj Next Steps](V5/CoveStack-Anuj-Next-Steps.md), [Go Handover](V5/CoveStack-Go-Handover.md), [Security](V5/CoveStack-Security-and-ITISS.md) and [Operations](V5/CoveStack-Operations-and-Recovery.md). The personal Anuj_Working guide may be shared separately; it is not required to exist in a fresh clone.

A owns all agent Go and reviews every Go change. AI-generated agent code is reference only; A types and explains it. A's first PR should be a small client/module slice, not both drivers and every account at once.

## 10. B — Java core/database: start here

| Step | Work / location                                                                                   | Review or acceptance                                                                                  |
| ---- | ------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| 1    | Read kan-server/pom.xml and existing Java auth/contract packages                                  | Run Java baseline; agree service boundaries with C/D                                                  |
| 2    | Review [Database Schema](V5/CoveStack-Database-Schema.md), including migration order              | Reference DDL is not executable/migration-ready; no live schema was established by the docs           |
| 3    | Add real application main/configuration and staged Flyway migrations in later implementation work | B owns all migration files/SQL; C tests; A agent-column review; D DTO review                          |
| 4    | Implement scoped repositories/core resource and operation transactions                            | Account predicates/composite ownership; atomic idempotency                                            |
| 5    | Implement work/lease/fencing/observations with C/A                                                | Exact generation/owner/epoch ack; clear both lease fields; progress never acknowledges                |
| 6    | Canonical host/assignment, quotas, IP/port records                                                | One host across accounts; reserve ports before create; third-cluster rejection under transaction lock |
| 7    | PostgreSQL/storage records, Java bucket/credential reconcile and fixed stack orchestration        | C acceptance; A runtime integration; no DAG engine                                                    |

Your Java root is kan-server/src/main/java/com/ciagotech/kan/server/. Coordinate exact new package names; do not copy proposed schema placeholders into a migration. Enable packaged startup only after main/configuration exist and are verified. D's endpoint mappings must call the same scoped core services.

Read [API](V5/CoveStack-API-and-Compatibility.md), [Resource Services](V5/CoveStack-Resource-Services.md) and [Stack](V5/CoveStack-Stack-Templates.md). First slice: reviewed schema foundation and server boot/configuration, coordinated with C's required beans.

## 11. C — Authentication/security/tests/IPAM: start here

| Step | Work / location                                                                  | Review or acceptance                                                                                    |
| ---- | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| 1    | Inspect existing auth classes and tests under kan-server/src/                    | Understand KeyService, BearerTokenService, filters, nonce store and current principal                   |
| 2    | Review signing rules/vectors with A and schema/security with B                   | Hash-only api_tokens; encrypted service-kind signing_credentials; key custody remains reviewed decision |
| 3    | Implement real active credential/token lookup, revocation/expiry and permissions | No development fallback or public unrestricted key issuer                                               |
| 4    | Own all acceptance tests from day 1                                              | Add cases while B/A build slices; do not wait until their last milestone                                |
| 5    | Run real Java/Valkey/Go smoke with A/B when server exists                        | Valid request, replay/tamper denial, no nonce poisoning, fail-closed store outage                       |
| 6    | Test assigned agent/account scope and credential lifecycle                       | Cross-account requests denied; separate account contexts and shared host quota                          |
| 7    | Implement Java IPAM allocator with B's transaction model                         | Test 14: ten IP allocations and ten account VPCs without overlap                                        |
| 8    | Review DB/S3 isolation, MinIO admin custody/path and audit policy                | Tests 11–13; no secrets in ordinary list/get/spec/observed                                              |

Read [Security](V5/CoveStack-Security-and-ITISS.md), [Schema](V5/CoveStack-Database-Schema.md), [Testing](V5/CoveStack-Testing-and-Acceptance.md) and [Networking](V5/CoveStack-VPC-IPAM-and-Networking.md). First slice: credential lookup design/implementation with test fixtures and negative authentication cases. B owns migrations/SQL; you review tenancy/secrets and implement Java allocator/auth logic.

## 12. D — Java APIs/DTOs: start here

| Step | Work / location                                                                                            | Review or acceptance                                                                                |
| ---- | ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| 1    | Read kan-server auth/contract classes, internal API and [API guide](V5/CoveStack-API-and-Compatibility.md) | Know existing schemas versus proposed endpoints                                                     |
| 2    | Review IDs/request/response mapping with B                                                                 | Internal uuid, CLI public_id and numeric /v2 provider_id; finalize mappings rather than invent them |
| 3    | Build Java /v2 endpoints and DTO mappings against B's scoped services                                      | Authenticated ownership and explicit unsupported behavior                                           |
| 4    | Include the bounded /v2/kubernetes/clusters subset                                                         | Coordinate execution shape with A and tenant tests with C                                           |
| 5    | Provide reviewed API/schema and optional Terraform test fixtures                                           | E writes Go provider; you write Java mappings/fixtures                                              |
| 6    | Exercise real HTTP behavior as endpoints land                                                              | Test inputs/errors/pagination and no secret/kubeconfig content in ordinary views                    |

Read [Resource Services](V5/CoveStack-Resource-Services.md), [Schema](V5/CoveStack-Database-Schema.md) and [Terraform/Ansible scope](V5/CoveStack-Terraform-and-Ansible.md). Existing examples are fixture evidence, not proof every service runs. First slice: DTO/endpoint inventory and agreed request/response fixtures with B/C. **D writes no Go.**

## 13. E — Go CLI learner: start here

| Stage | Work / planned location                                                   | A's review gate                                                        |
| ----- | ------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| W1    | Go basics/modules/structs/errors; cobra hello; planned cmd/covestack/     | Explain code and reviewed module layout; three target builds           |
| W2    | HTTP/JSON/context/timeouts against a mock API                             | Non-2xx/invalid JSON/timeout/cancellation handled                      |
| W3    | Droplet create/list/get/delete, --wait, -o json and automatic idempotency | API-only; mock-only VPC/K8s list/get practice                          |
| W4    | Table tests/httptest, then real-server CLI                                | C's account/idempotency checks with actual backend; packaging evidence |

These are learning stages, not new calendar promises. Production command order is **droplet → vpc → k8s → db → bucket → stack**; previous service's failure test must pass in CI before the next production service starts.

Read [Go Handover](V5/CoveStack-Go-Handover.md), [Folder Map](V5/Folder_Structure_Explanation.md), [Tools](V5/tools.md) and [API](V5/CoveStack-API-and-Compatibility.md).

| CLI boundary            | Requirement                                                                                   |
| ----------------------- | --------------------------------------------------------------------------------------------- |
| Access                  | Authenticated API through reviewed pkg/kanclient; never Docker/k3d/DB                         |
| Credentials             | User API authentication, not agent service secrets; env/keychain; no secrets in args/logs     |
| Local config            | Server URL, key ID, profile; no local resource database                                       |
| Mutation                | Idempotency key; create returns pending immediately                                           |
| Wait/get/list           | Finite --wait; desired/converged generation, lease and observed state; every list/get -o json |
| Packaging               | linux/amd64, darwin/arm64, windows/amd64; actual CLI demonstrations on two OSes               |
| Later optional provider | Thin Droplet/K8s Terraform layer after CLI; A Go review; D Java fixtures                      |

Do not add Go modules or move the signer independently; agree imports with A first. Do not touch agent/signing/drivers/reconcile or take Java management/mail work from the historical role guide. No UI work during CORE-5. First slice: reviewed module plan, cobra skeleton/mock client and meaningful client tests.

## 14. First shared integration

| Order | Lead       | Required result                                                                     |
| ----- | ---------- | ----------------------------------------------------------------------------------- |
| 1     | B+C+A      | Reviewed schema, trusted agent assignment and credential delivery model             |
| 2     | B+C        | Real server/private dependencies and concrete authenticated identity                |
| 3     | A+C        | Real auth-smoke against Java/Valkey; distinct account/user credentials              |
| 4     | B+D        | One authenticated resource create/read flow with durable desired intent             |
| 5     | A+B+C      | One container converges and reports; crash/fence/tenant/idempotency failures tested |
| 6     | E with A/C | Same resource flow through thin CLI once its reviewed client exists                 |

Mocks may unblock learning or local unit tests; they do not replace this integration or justify bypassing authentication. Do not run a fake and real agent against the same resource assignment.

## 15. Daily work, commit and PR

Before starting a task, agree behavior, owner, dependencies and acceptance with its reviewers. Check your branch and work tree:

```bash
git status -sb
git branch --show-current
```

Implement a small slice; run the checks for affected code and required contracts. Inspect your diff. Stage exact paths you changed; replace the path placeholder below:

```bash
git diff --check
git diff --stat
git add path/to/your/changed/files
git diff --cached --check
git diff --cached --stat
git commit -m "Describe the concrete behavior and tests"
```

Push your actual feature branch. Example for B:

```bash
git push -u origin feat/b-server-foundation
```

Open a PR against the agreed target, normally main after the baseline merge. Include behavior, tests/commands/results, contract/schema impact, dependencies and unresolved limitations. Ask the correct reviewers before merge.

| Change                               | Required reviewers                          |
| ------------------------------------ | ------------------------------------------- |
| Any Go                               | A                                           |
| pkg/kanclient or signing integration | A+B+C                                       |
| k3d driver                           | A+B+C                                       |
| Schema                               | B+C+A; D DTO alignment                      |
| Java core/API/auth                   | Relevant B/C/D owners; A for agent boundary |
| Acceptance                           | C owns tests; relevant implementer pairs    |

## 16. Get updates without losing work

Before a new task, fetch and use the latest agreed baseline. For an existing feature branch with committed work, this is a simple merge-based update:

```bash
git status --short
git fetch origin
git merge "origin/$COVESTACK_BASELINE_BRANCH"
```

Run the merge only with a clean work tree; the baseline variable must still be set in this terminal. In a new terminal, set it again to the agreed branch. Resolve conflicts with the affected owner and rerun checks before pushing. Do not reset, force-push or silently overwrite teammates' changes. If integration dependencies changed, confirm the PR target/baseline with Anuj first.

## 17. Before marking a service Gates passed

| Gate | Required evidence                                                                   |
| ---- | ----------------------------------------------------------------------------------- |
| G1   | DB intent and generation are authoritative                                          |
| G2   | Crash-safe reconcile; inspect/rebuild from DB without duplicates                    |
| G3   | Other-account API request returns 404; required protocol/network denial also tested |
| G4   | Relevant failure test runs in CI                                                    |

Use **Planned / In progress / Gates passed / Extended track**. Supported is permitted only after all four gates pass. Tests 1–15 remain mandatory; C owns acceptance. No new CORE-5 features after 10 January 2027. Extended work never delays CORE-5; if A is one week behind, stop all extended work.

Never commit actual .env/secret files, credentials, master/private keys, credential-bearing kubeconfigs, Terraform state or logs containing them. Keep CoveStack capitalization and kan-server/kan-agent/KAN1/KAN\__/X-Kan-_ compatibility identifiers. Never assume proposed kan.\* reserved labels already exist in the compatibility contract.

## 18. First-day checklist

- Confirm GitHub access, role, baseline and first task with Anuj.
- Clone the whole repository and create your own feature branch.
- Read SRS, design, roles, workflow, setup, gates and acceptance.
- Check your toolchain and run relevant baseline tests.
- Record blockers and actual results; do not bypass security to make a demo pass.
- Review assigned files and agree your first small slice with reviewers.
- Implement on your feature branch, verify, commit only intended files and push.
- Raise a PR and update task evidence; no direct main push.
