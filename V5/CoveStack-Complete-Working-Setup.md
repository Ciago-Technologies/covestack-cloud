# CoveStack V5 — Setup, Deployment and Verification

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Scope and current evidence

The retained filename is not a claim the full product works. Existing source has Go config/signing/client/no-execution polling, Java auth/replay/contract libraries and partial extended migrations. No runnable full server/foundational schema, VM/container/k3d driver, RustFS integration, public CLI or supported deployment profile is established. Read [status](CoveStack-Implementation-Status.md) before attempting installation.

## 2. Environment inventory and prerequisites

| Area | Existing input / target | Evidence required |
| --- | --- | --- |
| Source tools | Current go.mod/CI toolchain pin 1.27.2; pom/CI Java 25, Maven; Python contract tools | Check exact checkout/version; historical local tool output is not a pin change |
| Dependency services | Compose defines platform PostgreSQL, authenticated Valkey and Mailpit | Not Java server, agent assignment, tenant PostgreSQL, RustFS or VM execution |
| VM host | Approved Linux/kernel/Incus/QEMU/KVM/storage/network/image | Hardware/nested virtualization, privileges, guest boot, measurements and hardening |
| App Container/Kubernetes | Reviewed Docker/OCI and tested k3d+k3s pair | Actual runtime/ports/policy/capacity; no package/tag chosen by a diagram |
| Object storage | RustFS preferred Planned provider | Approved immutable artifact, operator volumes/private administration, Java API and real byte recovery |
| Secrets | Protected key/DB/replay/admin delivery and rotation | No committed default keys, real .env/secrets/kubeconfigs or secret command args |

No install, dependency service start, application or infrastructure test was run for this documentation task. Current sandbox restrictions remain; get a separately authorized capable environment before provisioning.

## 3. Clone and branch

Use [Starter.md](../Starter.md) for exact clone/baseline/feature-branch/PR procedure. Confirm Anuj's reviewed baseline before branching; never assume all prerequisite PRs are merged. Preserve local/user changes; no reset/clean/force push or secret staging. The canonical AI handoff is [checkpoint.md](../checkpoint.md), not a duplicate personal status file.

## 4. Existing baseline checks

From kan-agent/ on an authorized dev environment:

```bash
go test -race ./...
go vet ./...
```

From kan-server/:

```bash
mvn --batch-mode --no-transfer-progress verify
```

From repository root for contract work:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
python3 contracts/generate_lifecycle.py --check
```

Report actual result/version/dependency blocker. No Maven wrapper or runnable Spring application main is assumed. Build success is not startup/migration/VM proof. CLI packaging/dependency checks apply only once its module exists: linux/amd64, darwin/arm64, windows/amd64; dependency graph must not include runtime drivers or direct DB. Do not run a fictional cmd/covestack module command today.

## 5. Private infrastructure and custody

Read existing Compose/example config and approved secret delivery outside version control. Exact image tag/digest must be tested and recorded; never substitute latest to imply support. Platform DB/Valkey bindings remain loopback/private. No `down -v` or volume deletion to fix configuration; preserve data and diagnose owner/state. Mailpit is development capture only, not production mail.

RustFS is not in current Compose. A future operator deployment needs a persistent data volume with reviewed permissions, private management/TLS/console access and protected administrator identity. Java tenant bucket/key/policy reconcile is separate from the process lifecycle; tenant keys never receive root/admin access. [Provider gaps and tests](../docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance).

## 6. Backend host evaluation gates

| Spike | Check before publishing install commands |
| --- | --- |
| Incus VM | Supported host/architecture/kernel, KVM access/nesting, image/firmware provenance, CPU/RAM/disk overhead, disk/NIC/firewall ownership, guest readiness, restart and uncertain async task recovery |
| App Container | Trusted shared-kernel policy, safe capabilities/mounts/devices, network/volume persistence and health |
| Kubernetes | Exact tested Docker+k3d+k3s pair, CLI flags, cgroup/security/firewall behavior, ports, real workload and capacity under combined load |
| RustFS | Exact admin/data API compatibility, scoped policy/key lifecycle, volumes/restart, actual object bytes and independent restoration |

No RAM/density guarantee or version pin is inferred. Fail if host lacks approved capability; never silently fall back from VM to container or an unqualified emulator. Root socket authority remains privileged even with service hardening.

## 7. Current agent integration

Existing main loads environment config, creates signed client and runs cancellable single-account read-only polling. Current names: KAN_SERVER_URL, KAN_ACCESS_KEY, KAN_SECRET_KEY, optional KAN_RULES_PATH. Loader does not source .env; default rules path assumes kan-agent working directory. Remote origin requires HTTPS; no credentials/path/query in origin.

B+C must provide the real server endpoint, protected test service key and trusted account/host assignment; a contract's localhost example is not a verified listener. Current GET /internal/v1/agent/work returning [] is no work. Returned work deliberately stops with no claim/ack/execute/delete. General transport-outage retry policy is not proven by timeout-only tests. [Team checkpoint](../TEAM-FIRST-CHECKPOINT.md).

## 8. Future supported installation sequence — Planned

| Step | Operator/server/agent requirement | Acceptance |
| --- | --- | --- |
| Preflight | Detect approved OS/backend/virtualization/disk/routes/capacity | Explicit refusal of unsupported host; no silent fallback |
| Bootstrap | Protected operator identity, key custody, DB roles/replay store and initial assignments | No anonymous/default secret bootstrap |
| Schema/server | Real app entry point/config and reviewed forward-only migration foundation | Empty DB/upgrade/startup and real auth, not fictional main |
| Backends | Agent-only privileged sockets, owned storage/network policy and measured reserves | Guest/tenant cannot reach host/admin routes |
| RustFS | Operator service/volumes/private admin and reviewed Java adapter | Known bytes, scoped denial, restart/restore |
| Agent | Trusted signed account contexts and approved future capabilities/claims | Actual auth/no-work first, then bounded runtime failure evidence |
| Data/environment | Owned ready child services with secret/retention policy | Actual SQL/S3/app workflow and safe retry/cleanup |
| Operator handover | Verified artifacts, runbooks/health/audit/key recovery | Second operator repeatable clean install and recovery/upgrade proof |

### 8.1 Ansible installer and service hardening

Planned operator roles cover runtime/dependencies/server/agent; configure reviewed secrets via Vault/no_log or approved custody, validate signatures/checksums, lint, and prove second-run idempotence. No playbook/systemd file is created here. NoNewPrivileges/ProtectSystem suggestions need actual compatibility testing with required runtime/socket/volume access; they do not remove agent privilege. Future inventory reads authenticated API/CLI JSON, never backend/DB.

## 9. Clean-machine and demo evidence

Record exact artifacts/profile, sanitized setup/commands, fresh test results and limitations. Immediate live demo can prove real signed no-work/auth/cancel when server exists. Later show real VM interrupted create, tenant denial/idempotency/generation/freshness and SQL/object/network/restore evidence from CI/video if needed. Test actual CLI runs on supported OS targets rather than infer from cross-compile. Have backup video and local evidence; do not depend on phone hotspot/router control.

## 10. Known constraints

| Constraint | Current fact / target | Scope effect | Status |
| --- | --- | --- | --- |
| Hardware | Old 8 GB AlmaLinux VM is historical dev/lab context | No production minimum or VM/cluster density; benchmark profiles | Qualification Planned |
| VM capability | Host/KVM/Incus execution not verified here | Dedicated authorized spike before VM commands/support | Blocked on environment/evidence |
| Cluster/node limits | No current supported density/profile | Reject by measured physical/account reservation; old two-cluster/node caps not product ceilings | Planned |
| Failure domain | Single-host target is not HA; multi-host future | Host failure interrupts services; independent backup and safe recovery required | Planned |
| Network access | Historical phone-hotspot/institute-router restrictions | No assumed router DNS/DHCP or inbound internet route; backup recording | Environment-specific |
| Agent state | Signed polling/no-execution exists; server/credential/assignment foundation incomplete | Real integration before claims/drivers; no fake success | Source-confirmed / integration pending |
| Team capacity | A runtime Go plus E thin CLI learner; B/C/D Java/security/API | Protect current reviews/tests; no implied full product deadline | Retained ownership |
| Object backend | RustFS preferred Planned integration, exact artifact/admin/backup unverified | No supported S3/profile claim; alternatives via adapter | Planned |
| License | Root Apache-2.0 adopted; third-party distribution/provenance still reviewed | License is not runtime support or blanket artifact clearance | Adopted / review pending |

Follow [operations](CoveStack-Operations-and-Recovery.md), [capacity](../docs/platform/capacity.md) and [readiness](../docs/platform/readiness.md). No mentor confirmation, new deployment permission or deadline is invented.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-complete-working-setup) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
