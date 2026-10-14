# CoveStack V5 — Technical Design and Team Architecture

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Purpose and scope

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

CORE-5 is the SRS/deliverable scope. CLI is the first interface; no UI work during CORE-5. Every service defaults to **Planned**. Replaceability is an architecture target, not evidence of interchangeable drivers today. The single integration template has PostgreSQL, MinIO and an app Droplet; it is not a DAG engine.

| User problem | Platform response | Evidence |
|---|---|---|
| Repeated manual environment setup | One reviewed integration template | Test 15: retry without duplicates |
| Shared lab users need scoped access | Account-owned resources and action permissions | Cross-account denial tests |
| Workloads need scoped connectivity | Account-owned private network and credentials | Tests 11–14 |
| Failures cause uncertain state | Desired/observed state and durable operations | Restart without duplication |
| Infrastructure needs credible evaluation | ITISS configuration, tests and runbooks | Clean install and measured restore |

### 1.1 Why this is not a wrapper

The comparison concerns the account-scoped platform contract, not unique ownership of runtime mechanisms. These are target assertions, not claims of current capability.

| Docker/k3d/MinIO alone cannot do X | CoveStack does Y — target | Evidence |
|---|---|---|
| Enforce our account/resource API policy | Multi-tenancy with other-account 404 | Test 3 |
| Persist our desired/observed generations | DB-authoritative intent and convergence | Tests 1, 5 |
| Fence our competing agent identities | Lease owner/epoch and exact-generation ack | Test 2 |
| Resume our container/cluster operation records | Inspect deterministic identity before create | Tests 1, 7 |
| Atomically reserve our API request keys | One resource/result transaction | Test 4 |
| Supply our measured owned-data restore result | Backup, restore and canonical dataset verification | Test 11 |
| Enforce our scoped kubeconfig API export | No kubeconfig in list/get/observations | Test 8 |
| Issue our account-scoped bucket credentials | Java bucket/policy reconcile | Tests 12, 13 |
| Allocate according to our account/host policy | Collision-free owned subnet/IP reservations | Test 14 |
| Expose our Terraform-usable /v2 contract | Optional API-only Droplet/Kubernetes provider | Provider lifecycle/JSON gate |

### 1.2 Two scope layers

| Layer | Scope | Status |
|---|---|---|
| CORE-5 SRS/deliverables | Droplets; required Kubernetes; Managed PostgreSQL; Object storage; VPC/IPAM | Planned |
| Cross-cutting, not services | Security, recovery, CLI, optional Terraform, one stack demo, Ansible Part 1 | Planned |
| Extended track, not SRS | [Roadmap](CoveStack-Extended-Roadmap.md); dependency gates first | Extended track |

### 1.3 Service Gate Rule

**Supported** is permitted only at status **Gates passed** after G1 DB desired state/generation, G2 crash-safe rebuild from DB without duplicates, G3 other-account 404, and G4 one failure test in CI. DB rows are authoritative, never execution objects/configuration. Status vocabulary: **Planned | In progress | Gates passed | Extended track**. No promotion is asserted here. See [Service Gates](CoveStack-Service-Gates.md).

## 2. Architecture

```text
Go/cobra CLI / optional thin Terraform provider
                         |
                  Authenticated API
                         |
            kan-server services and policy
                   /               \
       PostgreSQL desired state   Replay store
                         |
            Signed outbound agent polling
                         |
               kan-agent reconciliation
                 /             \
       Generic container       k3d CLI
       (network helper)        driver
                 \             /
             Observations and operation results
```

| Component | Responsibility | Authority boundary |
|---|---|---|
| `kan-server` | Accounts, resource/stack records, actions, leases, API policy | Does not grant public clients daemon access |
| `kan-agent` | Validate work, inspect/converge owned infrastructure | Trusted high-authority host component |
| PostgreSQL | Durable desired intent, ownership and action records | Private; migrations and backups reviewed |
| Valkey | Authenticated replay-related storage | Private; outage policy fails closed |
| Docker | Container execution, bridge/IPAM primitives | Never mounted into user workloads |
| k3d / k3s | Local Kubernetes execution | Protected kubeconfig and bounded ports |
| Java bucket/VPC reconcile | Server-owned desired records and allocator | Not new bucket/VPC agent kinds |

CLI, Terraform provider and Ansible inventory talk **ONLY to the authenticated API**; none import or call Docker/k3d. Ansible Part 1 is the operator installation playbook, a distinct infrastructure activity. Only two workload drivers are planned: generic container and k3d. PostgreSQL/MinIO processes use container image/env/volume/labels/health; network execution is a helper inside it.

## 3. State and reconciliation

| Concept | Rule |
|---|---|
| Resource identity | Stable logical ID plus account, host/agent and kind |
| Desired generation | Version of requested specification; increment transactionally |
| Observation | Actual state, time, generation acknowledgment and lease identity |
| Progress | Separate guarded statement; never changes acked_generation |
| Convergence | Acknowledge only after actual requested state/readiness is verified |
| Stale report | Ignore safely; cannot overwrite newer desired intent |
| Failed attempt | Record failure; retry is bounded and traceable |
| Agent restart | Inspect reserved ownership identifiers before creating |
| Server outage | Pause new work; do not infer deletion from failed/empty polls |

Convergence acknowledgment requires id, lease_owner, lease_epoch match **and generation = :gen**. Zero rows means rejected/stale. Epoch is the fence; lease-expiry checking is an open decision for B. See [reference schema/queries](CoveStack-Database-Schema.md).

A controller checks what exists and brings it toward the requested state. Starting a fresh container on every poll is not reconciliation. Serialize work for the same logical resource; introduce bounded concurrency only after lease/fencing and per-resource deduplication work.

The current contract source remains `contracts/lifecycle-rules.json` version 1.1.0. **V5 is the documentation version, not a wire-protocol version.** New stack/network/ingress messages require coordinated schema changes, fixtures and generated Go/Java outputs. Do not invent undocumented strings in an agent work payload.

## 4. Ownership and consistency

| Boundary | Required behavior |
|---|---|
| Account scope | Derive identity from authentication; validate every object action |
| Database writes | Account predicate and optimistic/version checks; avoid whole-row observation overwrite |
| Idempotency | Scope request key, store fingerprint and result atomically |
| Lease/fencing | Old agent cannot acknowledge or mutate work after losing authority |
| Runtime labels | Proposed kan.account/kan.resource/kan.agent contract; review before use; not present compatibility |
| Garbage collection | Complete scoped inventory plus explicit deletion intent and grace policy |
| Secret data | Separate authorized retrieval from ordinary metadata listings |

## 5. Resource services and stack composition

The Service Catalog and stack topic files will be aligned in Batch 3. [CORE-5 SRS](CoveStack-SRS-Scope.md) and [Service Gates](CoveStack-Service-Gates.md) govern this batch. A stack records child identities and fixed order; it exposes each resource's safe metadata and records partial failure/retry, not a new orchestration engine.

### 5.1 Droplets / compute instances — Planned

Execution backend only (Docker). Platform value: tenancy, desired/observed generation, lease and fencing, crash-safe restart.

Container-backed, single host, never called a VM. Generic container driver inspects deterministic owned identity before create; API-only CLI create/delete/get evidence is tests 1–6. No host daemon socket in user workloads.

### 5.2 Kubernetes — Planned, REQUIRED

Execution backend only (k3d). Platform value: ownership, scoping, crash-safe lifecycle, quota, scoped kubeconfig.

| Logic we own | k3d alone cannot |
|---|---|
| Account-owned clusters, proposed labels kan.account, kan.resource, kan.agent | k3d has no tenants |
| Desired/observed cluster state (name, k3s version, node count, generation) | no desired state |
| Crash-safe create: kill agent mid-create => exactly 1 cluster | create is not idempotent |
| API-port allocator, reserved BEFORE create, keyed by resource ID | takes any port |
| Scoped kubeconfig export on demand; never in list/get/observations | prints it to anyone |
| Node-pool scale via reconcile | manual |
| Quota: 2-cluster host cap, per-account share inside it, node limits | none |
| Delete removes all labeled containers, volumes, networks | can leave leftovers |
| DO-shaped /v2/kubernetes/clusters subset | not in k3d |

This table describes the missing CoveStack account/policy semantics at the execution boundary; it does not claim unauthenticated users can access a protected host's k3d command. Labels are proposed, and node-pool/attachment message changes need coordinated contract work.

| Limit/prerequisite | Target |
|---|---|
| Host | One 8 GB AlmaLinux VM |
| Cluster cap | Maximum 2 per host; third rejected with documented quota error TBD with B+D |
| Shape | 1 server + maximum 2 agent nodes each |
| Version | One pinned tested k3d+k3s pair; values TBD after spike |
| Memory estimate | 0.5–1 GB/cluster: unverified estimate, confirm in K8s host spike |
| November host spike | One day: cgroup v2, SELinux, firewalld, RAM budget and tested pair |
| Execution | Shell out to pinned k3d CLI, not library; verify each flag before documenting |

Ensure finds a half-created cluster by proposed labels/deterministic name and finishes it or cleans owned partial objects before recreating under the same name. Port is reserved first; tests 7–10 prove crash, tenancy, cleanup and quota.

### 5.3 Managed PostgreSQL — Planned

Execution backend only (container). Platform value: volume ownership, backup/restore, credential scoping.

Acceptance requires backup, measured restore, scoped credentials and private networking together. Generic container driver reattaches owned volume; test 11 verifies rows/checksum and cross-account denial. No service is labeled managed without those acceptance items.

### 5.4 Object storage — Planned

Execution backend only (MinIO). Platform value: per-account credentials, bucket scoping, cross-account denial.

MinIO process uses the container driver. Java server reconciles physical buckets and scoped credential policy from DB rows; storage.bucket remains server-owned. Tests 12–13 prove denied key usage and rebuild. Root storage credentials remain private.

### 5.5 VPC/IPAM — Planned

Execution backend only (Docker bridge). Platform value: allocator, ownership, collision-free subnets and IPs.

One private network/account; no custom CIDR. C owns Java allocator; B owns SQL; A owns container-driver network helper. Test 14 includes ten parallel IP allocations within one VPC and ten accounts allocating networks concurrently. Pool values TBD A+B. VPC has no kind in 1.1.0: server-owned design, state and compute attachment decision B+A; no invented wire fields.

## 6. API layers

| Interface | Policy |
|---|---|
| Internal agent routes | Shared `KAN1` signing contract; private route |
| Provider-style `/v2` routes | Tested bounded DigitalOcean subset; Bearer permissions |
| Management routes | CORE-5 stack/network/account records; scoped authorization |
| Minimal liveness | Small non-sensitive response only |
| Detailed doctor | Authenticated account/admin scope; no sensitive public status |

Every list/get supports stable JSON for automation. CLI mutations send an idempotency key; create returns pending immediately and --wait polls convergence. Get exposes desired/converged generation, lease and observed state, never kubeconfig or secret credentials. Kubeconfig is authorized export on demand.

Provider compatibility is an endpoint-by-endpoint target matrix. Same-looking JSON and paths do not imply full provider parity. SSH key validation and action completion need their exact tested mappings; unknown new message strings are TBD with B. IDs: internal uuid, prefixed CLI public_id, numeric /v2 provider_id; B+D finalize.

## 7. Extended track (not in SRS)

Detailed design is preserved in [Extended Roadmap](CoveStack-Extended-Roadmap.md), not CORE-5 deliverables. B/C/D/E server-side work may start from 1 December only if CORE-5 duties are not blocked and dependency gates pass. A's extended Go work starts only after all CORE-5 gates pass. A one-week delay stops all extended work. The track continues after submission; demo/viva mention is bonus only after Gates passed.

## 8. Decision records required before code

| CORE-5 decision | Owner |
|---|---|
| Go module layout and planned signer move to pkg/kanclient | A+B+C |
| CLI/agent boundary | A+E |
| k3d+k3s pinned pair, API-port range, kubeconfig delivery | A+B+C |
| Subnet pool, Java allocator and compute attachment | B+A; C allocator |
| MinIO credential model and key custody | B+C |
| Backup storage format and retention | B+A |
| ID mapping and missing contract vocabulary | B+D |

Extended tool/policy decisions move to the roadmap; values are not invented. See [Known constraints](CoveStack-Complete-Working-Setup.md#10-known-constraints), added in Batch 2. Limits communicated to mentor on <date: fill in>. Mentor confirmation: pending.

### 8.1 Planned folder layout — no code move in this task

```text
covestack-cloud/
  contracts/
  pkg/kanclient/                  signer + HTTP API client (planned)
  cmd/covestack/                  cobra CLI (E)
  kan-agent/                      container/k3d drivers, network helper, reconcile (A)
  kan-server/                     Java core/security/API (B, C, D)
  terraform-provider-covestack/    thin optional Go layer (E, after CLI)
  ansible/                        Part 1 roles; inventory later
  deploy/
  docs/
```

Root/client/agent Go module layout is a decision record; the current module cannot import a new root package merely because this tree is drawn. No go.mod or signer relocation occurs here.

## 9. Reading order

Read [SRS scope](CoveStack-SRS-Scope.md), [Service Gates](CoveStack-Service-Gates.md), [Acceptance](CoveStack-AntiWrapper-Acceptance.md), [Schema](CoveStack-Database-Schema.md), [Go Handover](CoveStack-Go-Handover.md), then [Extended Roadmap](CoveStack-Extended-Roadmap.md) for deferred design. Workflow/roles/setup and remaining topic files are updated in Batches 2–3; their older scope is transitional, not authority over this approved CORE-5 decision.
