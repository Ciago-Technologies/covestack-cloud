# CoveStack V5 — Setup, Deployment and Verification

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **15 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Scope of this setup

**This guide describes target deployment and required evidence.** These instructions separate usable baseline checks from the target installation. Do not run a fictional CLI/main entry point and treat failure as an environment problem.

| Stage               | What can be verified                                                           |
| ------------------- | ------------------------------------------------------------------------------ |
| Current baseline    | Go tests/vet, signing vectors, lifecycle generation, Java library tests        |
| Infrastructure      | Compose configuration and private Postgres/Valkey/Mailpit readiness            |
| Backend gate        | Application entry point, migrations, concrete auth and API routing             |
| Agent gate          | Configuration, signed client, polling and real resource driver                 |
| CORE-5 target       | Droplets, required Kubernetes, Managed PostgreSQL, Object storage and VPC/IPAM |
| Extended deployment | Separate [roadmap](CoveStack-Extended-Roadmap.md); no CORE-5 UI work           |

## 2. Environment inventory

Record the exact tested versions before release. Existing pins are repository observations, not blanket recommendations for future installs.

| Tool           | Current repository expectation                      | Verify                                                     |
| -------------- | --------------------------------------------------- | ---------------------------------------------------------- |
| Go             | Language directive 1.22; preferred toolchain 1.27.2 | Actual executable and CI selection                         |
| Java           | 25 in `kan-server/pom.xml`                          | Both runtime and compiler                                  |
| Spring Boot    | Parent 3.5.16 in inspected POM                      | Dependency availability and real build                     |
| PostgreSQL     | Compose requests a reviewed 17 image                | Exact image/digest and mount layout                        |
| Valkey         | Compose requests a reviewed 8 image                 | Authentication and replay persistence                      |
| Docker/Compose | Needed for current execution/infra                  | Engine, compose plugin and firewall backend                |
| k3d/k3s        | Required CORE-5 integration                         | One tested pair; pin after host spike; max 2 clusters/host |

Use the [tool matrix](tools.md). Review support/security updates before choosing release versions; do not change pins only because a course uses another release.

## 3. Clone and work on a feature branch

```bash
git clone https://github.com/Ciago-Technologies/covestack-cloud.git
cd covestack-cloud
git status -sb
git remote -v
```

Use the agreed reviewed baseline branch. If baseline PRs are merged:

```bash
git switch main
git pull --ff-only origin main
git switch -c feat/your-task
```

Do not share deployment secrets. Read V5 roles and status before accepting implementation tasks.

## 4. Baseline checks

From the repository root:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
(cd kan-agent && go test -race ./...)
(cd kan-agent && go vet ./...)
(cd kan-server && mvn --batch-mode --no-transfer-progress verify)
```

| Result                           | Interpretation                                              |
| -------------------------------- | ----------------------------------------------------------- |
| Vector/schema pass               | Contract evidence only                                      |
| Go unit/race pass                | Exercised implementation paths, not all production behavior |
| Maven verify pass                | Library/build result; application boot needs an entry point |
| Network/cache/dependency failure | Incomplete verification, not success                        |
| Local toolchain fallback         | Label actual tested version; not pinned-version evidence    |

The inspected Boot plugin has `skip=true` while the server lacks the application main. B must enable packaging after that entry point and configuration exist, then verify the packaged app actually boots.

### 4.1 Planned CLI dependency and packaging gates

Run after the reviewed Go module layout and cmd/covestack exist, from the owning module root. These are verification targets, not commands for a fictional current main:

```bash
go list -deps ./cmd/covestack | grep -i -e docker -e k3d
GOOS=linux GOARCH=amd64 go build -o /tmp/covestack-linux-amd64 ./cmd/covestack
GOOS=darwin GOARCH=arm64 go build -o /tmp/covestack-darwin-arm64 ./cmd/covestack
GOOS=windows GOARCH=amd64 go build -o /tmp/covestack-windows-amd64.exe ./cmd/covestack
```

Dependency grep must print nothing; CI must separately check go list succeeded, since grep's no-match exit alone cannot prove a build. CLI imports no Docker/k3d, calls only authenticated API, and uses no local DB. Cross-builds require actual CLI runs on two OSes; they do not prove runtime portability.

## 5. Private infrastructure configuration

Create `deploy/.env` from the tracked template **only if absent**. Choose reviewed exact image references rather than leaving blanks or using `latest`. Preserve an existing configured file.

| Local file                         | Purpose                                             |
| ---------------------------------- | --------------------------------------------------- |
| `deploy/.env`                      | Postgres, Valkey and Mailpit image references       |
| `deploy/secrets/postgres_password` | Database password secret                            |
| `deploy/secrets/valkey.conf`       | Authenticated Valkey configuration from template    |
| `deploy/secrets/agent.env`         | Agent server origin and service credential settings |

Generate secrets using a cryptographic source and store locally; do not paste them into repository docs. Restrict the secret directory to the owner and files to required readers. The agent signing secret and infrastructure passwords are different credentials.

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
docker compose --env-file deploy/.env -f deploy/docker-compose.yml ps
```

| Existing host listener | Intended access                         |
| ---------------------- | --------------------------------------- |
| `127.0.0.1:5432`       | Control-plane PostgreSQL                |
| `127.0.0.1:6390`       | Valkey host access; container port 6379 |
| `127.0.0.1:1025`       | Mailpit SMTP capture                    |
| `127.0.0.1:8025`       | Private Mailpit UI                      |

Compose starts these dependencies, not Java or the full Go agent. Existing PostgreSQL volumes retain initialized credentials. Rotate intentionally; never delete volumes merely to make a new environment variable take effect.

## 6. Backend integration gate

| Step                                        | Owner/evidence                                |
| ------------------------------------------- | --------------------------------------------- |
| Add application main and real configuration | B: package and boot                           |
| Apply fresh migrations                      | B: clean-database evidence                    |
| Implement active credential/token lookup    | C: real store and revocation                  |
| Provision account-bound service key         | C/B: restricted issuance and custody review   |
| Register scoped APIs                        | B/C/D: Java HTTP tests; E consumes API        |
| Configure replay store                      | C/A: authenticated store and outage rejection |

Start the server using the entry point/artifact B has actually added. Until then, no production startup command is asserted. Keep internal agent traffic on a private host route and preserve signed authority/path/query bytes.

## 7. Agent configuration and smoke

The local config loader reads process environment, not `.env` automatically. `KAN_RULES_PATH` defaults to `../contracts/sigv1-rules.json` when executed from `kan-agent`.

```bash
set -a
source deploy/secrets/agent.env
set +a
(cd kan-agent && go run ./cmd/auth-smoke)
unset KAN_ACCESS_KEY KAN_SECRET_KEY KAN_SERVER_URL KAN_RULES_PATH
```

Source only your trusted file. Real server-side credentials are required. Expected smoke statuses are valid 200, replay 401, invalid MAC 401, preserved valid request 200 and nonce tampering 401. A mocked handler test is not this deployment test.

The planned production agent must load settings once, reuse a finite-timeout/no-redirect client, handle shutdown, heartbeat/poll, inspect ownership and reconcile installed supported kinds. Publish an actual main command only with reviewed build/runtime evidence. Agreed target: one agent process per physical host; separate service credential and authenticated polling context per assigned account. B owns server assignment validation; C credential provisioning/rotation/revocation; A isolated clients and shared bounded execution. Provisioning/storage details still need B+C+A review; no new wire fields. Do not add caller-selected account wire fields.

## 8. Target environment installation sequence

| Order    | Configuration                                                       | Verify before next step                                     |
| -------- | ------------------------------------------------------------------- | ----------------------------------------------------------- |
| 1        | Single AlmaLinux lab VM; private control-plane PostgreSQL 17/Valkey | Resource budget and protected listeners                     |
| 2        | Java auth/core, signed Go agent, thin CLI                           | Real auth; tests 1–6 through API                            |
| 3        | Server-owned VPC/IPAM and container-driver network helper           | Test 14; network attachment schema decision B+A             |
| Before 4 | Early one-day K8s host spike; repeat integration check in November                                  | cgroup v2, SELinux, firewalld, RAM; one tested k3d+k3s pair |
| 4        | Required k3d driver, reserved API port, scoped kubeconfig           | Tests 7–10; hard cap 2 clusters                             |
| 5        | PostgreSQL and MinIO on generic container driver                    | Tests 11–13; server-side bucket/credential rebuild          |
| 6        | One fixed-sequence Postgres + MinIO + app Droplet stack             | Test 15; partial failure/retry; no DAG engine               |
| 7        | Ansible Part 1 installer and optional thin Terraform                | Repeat install; API-only provider for Droplet/K8s           |

### 8.1 K8s host spike prerequisite

One day at first host preparation, before committing driver capacity and before installation step 4; revalidate in November: test AlmaLinux cgroup v2, SELinux and firewalld; record the combined Java/DB/cache/MinIO/application RAM/CPU/disk plus one and two clusters and cleanup results. The **0.5–1 GB per cluster is an unverified estimate, confirm in K8s host spike**. Hard cap: 2 clusters/host; each 1 server + max 2 agents. Record one tested/pinned k3d+k3s pair, API-port range and every CLI flag against that version. Fixtures are not release pins.

The k3d driver shells out to the CLI, not the library. After a crash, Ensure finds a half-created cluster using the proposed reserved-label contract, then finishes it or deletes/recreates under the same deterministic name. Third cluster request needs an explicit documented quota error; its wire value remains TBD with B. Kubeconfig export is separately authorized and absent from list/get/observations.

### 8.1.1 Startup, claim protocol and shared storage prerequisites

Do not launch multiple agent processes for different accounts. A holds an OS host exclusion lock, uses per-resource serialization and recovers outstanding children/backend requests before mutations. Losing a DB lease or timing out does not prove a Docker/k3d request stopped; uncertain outcomes block conflicting work until proven settled. No local lock file is desired-state authority. The reviewed claim/renew/report protocol must exist before mutation; current 1.1.0 does not encode the full fence exchange.

Default-VPC intent and the separate signed network-control exchange must converge before Droplet/DB attachment. Provision one operator-owned shared MinIO process/volume from durable platform-owned intent; install protected host secret references out of band. Java needs reviewed private administration connectivity; tenant apps get explicit scoped S3 access, not root keys or console access. The stack creates a tenant bucket in shared MinIO and never owns that process. No bootstrap command/path is asserted until actually added and tested.

Periodic inspection is DB scheduled after convergence and uses fresh claims; health completion does not advance acked_generation. Test drift removal, expired renewal, stale inspection and host takeover, not only initial create/ack. Preserve reservations while cluster/network cleanup is pending.

### 8.2 Ansible install playbook — Part 1

| Target layout/control                                                  | Acceptance                                                                        |
| ---------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| ansible/site.yml; roles/docker, roles/infra, roles/server, roles/agent | Reviewed host install, not tenant inventory                                       |
| systemd NoNewPrivileges and ProtectSystem=strict                       | Explicit writable paths; test daemon behavior without weakening settings silently |
| ansible-vault and no_log for sensitive tasks                           | Custody/access B+C review; no plaintext secret logs                               |
| Checksum-verified binaries                                             | Record source/version/checksum and fail on mismatch                               |
| ansible-lint                                                           | Pass on actual reviewed playbook                                                  |
| Repeatability                                                          | Second run changed=0; record both results                                         |

No playbook or systemd file is created in this docs batch. Inventory Part 2 and its SSH/Python/API-key-injection prerequisites are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#11-ansible-inventory-part-2). The operator installer is separate from tenant workload configuration.

### 8.3 Planned folder layout — no code moves

```text
covestack-cloud/
  contracts/
  pkg/kanclient/                  Go signer + HTTP API client (planned)
  cmd/covestack/                  Go cobra CLI (E; planned)
  kan-agent/                      Go (A): container, k3d, reconcile
  kan-server/                     Java (B, C, D)
  terraform-provider-covestack/   Go (E; after CLI)
  ansible/                        Part 1 site.yml/roles; inventory later
  deploy/
  docs/
```

Signer relocation and module layout are Planned workflow decisions, not actions in this batch. Network execution is a helper within the generic container driver; only two workload drivers exist in the target architecture. PostgreSQL/MinIO processes use image/env/volume/labels/health configuration on that driver; Java owns bucket/credential reconciliation.

Do not attach user databases to the control-plane network. Historical proxy-topology warnings remain in the extended roadmap: a proxy container's localhost is not the host. Tool choices are TBD decision records.

## 9. Clean-machine release acceptance

A teammate performs a clean-VM installation without ad hoc auth bypasses and records source revision, actual tools, tests and limits. Preserve data on routine shutdown; **down -v and blanket pruning are not repair instructions**. Existing PostgreSQL volumes retain initialized credentials; an env change is not a credential rotation.

### 9.1 Five-minute CORE-5 evidence script — CLI only

All commands below are target behavior, not asserted current executables. Live highlights: **tests 1, 3, 4 and get view**. Show the remaining evidence as CI output or recorded video.

| Mode/test                  | Exact evidence                                                                                                                                |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| Live 4                     | create --wait; repeat the same idempotency key → same ID; parallel same-key test in CI                                                        |
| Live 1                     | kill -9 kan-agent after container create before ack; restart; docker ps --filter label=kan.resource shows exactly one owned instance          |
| Live 3                     | Bob get Alice public ID → 404; no existence leak                                                                                              |
| Live get                   | Desired versus converged generation, lease and observed state                                                                                 |
| CI/video 2, 5, 6           | Expired-owner ack rejected by epoch; stale generation ignored; server down makes agent create/delete nothing                                  |
| CI/video dependency outage | Stop control-plane PostgreSQL → covestack droplet list fails clearly; stop agent → create remains pending and --wait times out clearly        |
| CI/video 7–10              | Cluster mid-create crash → one cluster; Bob kubeconfig → 404; list/get omit kubeconfig; delete cleans labeled objects; third cluster rejected |
| CI/video 11                | PostgreSQL backup, delete known rows, restore → rows and checksum return; B cannot reach A DB                                                 |
| CI/video 12–13             | A key denied on B bucket; restart MinIO, Java server recreates buckets/credentials from DB                                                    |
| CI/video 14                | Ten parallel IP allocations in one VPC → no duplicate; ten accounts create VPCs → no overlapping CIDR; tenant denial                          |
| CI/video 15                | One stack command; partial failure recorded; retry resumes without duplicates                                                                 |
| Installer/package          | Second Ansible run changed=0; cross-build three targets; CLI demonstrated on two actual OSes                                                  |

kan.* inspection labels are a proposed reserved-label contract, not established compatibility; approve them before using these assertions. CLI never calls Docker/k3d; the docker ps inspection is operator test evidence. Record a backup video; do not rely on the phone hotspot alone.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).

## 10. Known constraints

| Constraint          | Value                                 | Effect on scope                                                  | Status  |
| ------------------- | ------------------------------------- | ---------------------------------------------------------------- | ------- |
| Host RAM            | 8 GB, single AlmaLinux VM             | Host spike measures actual workload budget                       | Planned |
| Kubernetes host cap | Max 2 clusters per host               | Transaction lock and derived count; explicit rejection for third | Planned |
| Nodes per cluster   | 1 server + max 2 agents               | One bounded tested shape/version                                 | Planned |
| Host/network model  | Single host; no multi-host networking | Multi-host is Optional future work                               | Planned |
| Droplets            | Containers, not VMs                   | No VM/kernel isolation claim                                     | Planned |
| Institute network   | Phone hotspot; no LAN/WiFi            | No router-level DNS/DHCP; backup demo video required             | Planned |
| Go capacity         | One Go developer plus one Go learner  | A agent ownership; E thin CLI; review bottleneck controlled      | Planned |
| Priority            | Extended track never delays CORE-5    | Stop/cut extended work first                                     | Planned |

Limits communicated to mentor on <date: fill in>. Mentor confirmation: pending.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
