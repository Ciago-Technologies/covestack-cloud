# CoveStack V5 — Setup, Deployment and Verification

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Scope of this setup

**The current repository is not yet a complete runnable cloud platform.** These instructions separate usable baseline checks from the target installation. Do not run a fictional CLI/main entry point and treat failure as an environment problem.

| Stage | What can be verified |
|---|---|
| Current baseline | Go tests/vet, signing vectors, lifecycle generation, Java library tests |
| Infrastructure | Compose configuration and private Postgres/Valkey/Mailpit readiness |
| Backend gate | Application entry point, migrations, concrete auth and API routing |
| Agent gate | Configuration, signed client, polling and real resource driver |
| Extended deployment | VPC, proxy/tunnel, UI, storage and Kubernetes integration |

## 2. Environment inventory

Record the exact tested versions before release. Existing pins are repository observations, not blanket recommendations for future installs.

| Tool | Current repository expectation | Verify |
|---|---|---|
| Go | Language directive 1.22; preferred toolchain 1.27.2 | Actual executable and CI selection |
| Java | 25 in `kan-server/pom.xml` | Both runtime and compiler |
| Spring Boot | Parent 3.5.16 in inspected POM | Dependency availability and real build |
| PostgreSQL | Compose requests a reviewed 17 image | Exact image/digest and mount layout |
| Valkey | Compose requests a reviewed 8 image | Authentication and replay persistence |
| Docker/Compose | Needed for current execution/infra | Engine, compose plugin and firewall backend |
| k3d/k3s | Target integration, not baseline requirement | One tested version pair and shape |

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

| Result | Interpretation |
|---|---|
| Vector/schema pass | Contract evidence only |
| Go unit/race pass | Exercised implementation paths, not all production behavior |
| Maven verify pass | Library/build result; application boot needs an entry point |
| Network/cache/dependency failure | Incomplete verification, not success |
| Local toolchain fallback | Label actual tested version; not pinned-version evidence |

The inspected Boot plugin has `skip=true` while the server lacks the application main. B must enable packaging after that entry point and configuration exist, then verify the packaged app actually boots.

## 5. Private infrastructure configuration

Create `deploy/.env` from the tracked template **only if absent**. Choose reviewed exact image references rather than leaving blanks or using `latest`. Preserve an existing configured file.

| Local file | Purpose |
|---|---|
| `deploy/.env` | Postgres, Valkey and Mailpit image references |
| `deploy/secrets/postgres_password` | Database password secret |
| `deploy/secrets/valkey.conf` | Authenticated Valkey configuration from template |
| `deploy/secrets/agent.env` | Agent server origin and service credential settings |

Generate secrets using a cryptographic source and store locally; do not paste them into repository docs. Restrict the secret directory to the owner and files to required readers. The agent signing secret and infrastructure passwords are different credentials.

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
docker compose --env-file deploy/.env -f deploy/docker-compose.yml ps
```

| Existing host listener | Intended access |
|---|---|
| `127.0.0.1:5432` | Control-plane PostgreSQL |
| `127.0.0.1:6390` | Valkey host access; container port 6379 |
| `127.0.0.1:1025` | Mailpit SMTP capture |
| `127.0.0.1:8025` | Private Mailpit UI |

Compose starts these dependencies, not Java or the full Go agent. Existing PostgreSQL volumes retain initialized credentials. Rotate intentionally; never delete volumes merely to make a new environment variable take effect.

## 6. Backend integration gate

| Step | Owner/evidence |
|---|---|
| Add application main and real configuration | B: package and boot |
| Apply fresh migrations | B: clean-database evidence |
| Implement active credential/token lookup | C: real store and revocation |
| Provision account-bound service key | C/E: restricted issuance |
| Register scoped APIs | B/C/D/E: real HTTP tests |
| Configure replay store | C/A: authenticated store and outage rejection |

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

The production agent must load settings once, reuse a finite-timeout/no-redirect client, handle shutdown, heartbeat/poll, inspect ownership and reconcile installed supported kinds. A main command becomes documentable only when implemented and tested.

## 8. Target environment installation sequence

| Order | Configuration | Verify before next step |
|---|---|---|
| 1 | Dedicated lab VM and private infra | Resource budget and protected services |
| 2 | Java authentication/core and Go agent | Real create/report/delete |
| 3 | VPC pool and bridge/IPAM policy | Concurrent allocation and denied paths |
| 4 | k3d driver and kubeconfig export | Cluster ready and scoped access |
| 5 | PostgreSQL workload and MinIO | Persistence and scoped credentials |
| 6 | Proxy/tunnel/domain | External HTTPS app; internal services denied |
| 7 | Stack orchestration | Dependencies, partial failure and retention |
| 8 | CLI/UI/provider/playbook | All use supported authenticated APIs |
| 9 | Monitoring/restore | Triggered alerts and recovered known data |

Do not attach user databases to the control-plane network. A proxy on the host can use host-loopback upstream ports; a proxy container cannot use its own localhost as if it were the host. Test the chosen topology explicitly.

## 9. Clean-machine release acceptance

A teammate follows these steps on a clean VM, records exact versions and starts the supported release without ad hoc auth bypasses. Verify external application access from mobile data, protected kubeconfig and known-data restoration. Preserve data on routine shutdown; `down -v` and blanket pruning are not repair instructions.

Review [Testing](CoveStack-Testing-and-Acceptance.md), [Security](CoveStack-Security-and-ITISS.md), and [Operations](CoveStack-Operations-and-Recovery.md) before declaring a release usable.
