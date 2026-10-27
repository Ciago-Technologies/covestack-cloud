# CoveStack V5 — Anuj Implementation and ITISS Work Order

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **27 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Current starting point

Config/signing/lifecycle source and auth-smoke are inspection evidence only; no service is Gates passed from file presence. The next A work is reviewed client/polling, production agent and generic container reconciliation. Do not reapply old patches or rename compatibility packages. No code changes happen in this documentation task.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. Immediate implementation sequence

| Order | A work                                                  | Acceptance/review                                                     |
| ----- | ------------------------------------------------------- | --------------------------------------------------------------------- |
| 1     | Review local config/tests; schema and signed-poll scope | Implement agreed one-process/per-account-client model after assignment/custody review; no cross-account shortcut |
| 2     | Plan module layout and signer move into pkg/kanclient   | A+B+C review; move not performed now                                  |
| 3     | Typed HTTP client and injected-transport tests          | Finite timeout, cancellation, no redirects, bounded body              |
| 4     | Production heartbeat/poll/shutdown main                 | Real server/auth-smoke; advertise only installed reviewed kinds       |
| 5     | Generic container driver and reconcile                  | Tests 1–6 with C; B owns all lease/fence/idempotency SQL              |
| 6     | Docker network helper inside container driver           | C Java allocator; B records; test 14; attachment contract change B+A  |
| 7     | k3d CLI driver after early host spike; revalidate in November                | Tests 7–10; max 2 clusters/host; reserved port and scoped kubeconfig  |
| 8     | PostgreSQL and MinIO process/volume lifecycle           | Tests 11–13; Java reconciles buckets/credentials                      |
| 9     | One stack integration; Ansible Part 1 and evidence      | Test 15; installer second run changed=0                               |

Only two workload drivers: container and k3d. Generic container inputs are image/env/volume/labels/health. Proposed kan.* reserved labels require approval. Unknown states/wire fields remain TBD with B; contract 1.1.0 wins.

A types and can explain agent Go; AI-generated code is reference only. E owns thin CLI, never agent/signing/drivers/reconcile. Follow CLI service order droplet → vpc → k8s → db → bucket → stack, waiting for the previous failure test in CI.

### 2.1 HLD/LLD changes to your first slices

Run the total-host capacity spike at first host preparation, not after committing the combined workload shape. Before mutation, pair with B+D/C on claim/renew/report and separate signed network-control schemas; existing 1.1.0 fields do not define the full protocol. Do not add wire fields alone.

Acquire host OS exclusion, serialize resource mutations, retain child-operation exclusion and record durable attempts with B. A timeout does not prove a backend stopped. Ambiguous old execution blocks conflicting takeover until settled; DB epoch rejection alone is insufficient. Renew only current unexpired claims and stop new mutations when authorization is lost.

Default VPC execution is needed for the first real Droplet. Add periodic inspection of converged rows, using fresh claims and no health acknowledgment. Repair drift with unchanged desired generation; stale inspection cannot overwrite newer intent. For MinIO, execute one operator-owned shared process/volume from reviewed intent/host secret references; Java owns tenant buckets and the stack owns a bucket child, not the shared process. Preserve the fixed stack/test 15 and installer; optional/extended work is cut first.

These are Planned implementation duties. [HLD](DocsV5-DO_Team.md#22-hld--control-execution-and-data-paths) and [LLD](CoveStack-Database-Schema.md#7-concurrency-and-reference-queries) define the precise dependencies and failure evidence.

### 2.2 User experience responsibilities — Planned

Your immediate milestone remains one-account signed communication, no-work handling, bounded retries and clean shutdown against the real server. Add execution only after reviewed claim/network schemas. These experience checks accompany the existing order; do not start extended features to satisfy them.

| When | Your work | Pairing/evidence |
| --- | --- | --- |
| Container reconciliation | Run a reviewed app; preserve resource identity across restart; emit bounded sanitized observations/errors with freshness | B/D API and E display; C tests 1–6 |
| Network helper | Execute allocated network attachments; prove intended app/DB paths and cross-account denial | C allocator/tests, B records; test 14 |
| Kubernetes driver | Ensure usable scoped API access; deploy/read a bounded test workload after recovery | B/D/E integration; C tests 7–10 |
| Persistence and stack | Preserve volumes; app reads/writes DB and object bytes; help restore known data | B owns bucket reconcile/stack records; C tests 11–15; E CLI |
| Installer and ITISS | Teammate clean-host walkthrough, outage timeline, traffic matrix and measured restore | C reviews sanitized evidence; Ansible Part 1 |

You do not take over CLI output, SQL, bucket administration or the acceptance suite. Agree image, endpoint/secret delivery and API representation with their owners. See [experience assertions](CoveStack-AntiWrapper-Acceptance.md#33-user-experience-assertions--planned-existing-tests).

### 2.3 First milestone decision — single-account communication, Planned

Selected starting plan on 27 October 2026: one agent process, one assigned account and one dedicated service signing credential; short polling and no workload execution. Later add isolated account contexts under the existing host model. This is an implementation target authorized for planning, not a claim of server readiness, completed security review or a new wire release. Existing contract rules win; operational defaults below are selected starting values to test with B+C, not protocol constants.

| Item | Starting decision | Authority / owner |
| --- | --- | --- |
| Same-host origin | http://localhost:4700; real listening address must be confirmed | Existing internal-api.yaml development origin; B starts/configures server |
| Remote origin | Verified HTTPS; no TLS-verification bypass | Existing config rejects remote plain HTTP; B+C provide actual endpoint |
| Settings | KAN_SERVER_URL, KAN_ACCESS_KEY, KAN_SECRET_KEY, KAN_RULES_PATH | Existing environment loader; no automatic .env loading or invented config fields |
| Credential | Dedicated service key for one account/agent assignment; no default/shared key | C securely generates/issues; B stores trusted assignment; A loads secret |
| Stored secret | Signing ciphertext plus key version; Bearer tokens remain hash-only | B+C finalize master-key custody, algorithm implementation and rotation; no key stored alongside ciphertext in ordinary DB rows |
| Work request | GET /internal/v1/agent/work; read-only, not a claim | Current contract; verified credential determines account |
| No work | HTTP 200, JSON [] | Existing array schema; reject malformed/null responses and do not introduce 204 as the normal result |
| Capability | supportedKinds: [] until an installed driver passes review | Existing heartbeat supports empty list; server must return no executable work |
| Request timeout | 10 seconds per attempt | Selected initial client behavior; finite and cancellation-aware |
| Normal poll | 5-second base interval, with bounded small jitter | Selected initial client behavior; A+C test bounds; never busy-loop |
| Temporary failures | Backoff from 1 second, doubling to a 30-second cap, with bounded jitter; reset after valid success | Network/timeout and transient 5xx; rebuild/sign each attempt |
| Rate limit | For 429 respect valid Retry-After, with cancellable waiting; invalid value uses backoff | B+C+A review supported parsing/wait bounds; no immediate retry loop |
| Authentication/config errors | 401/403 or invalid config stops this account loop with sanitized diagnostic | No endless retry with invalid credentials; later other account contexts remain independent |
| Redirects | Refuse all redirects; do not forward signed credentials | Existing client boundary target |
| Shutdown | Ctrl+C/termination cancels HTTP request, polling wait and backoff; no next request | A implements context/signal handling |

Use response-size bounds and validate decoded data; exact local bound is A+B/C review, not the signed request-body limit. Other nonretryable 4xx and malformed/oversized responses are explicit errors, not empty successful polls. Never log keys, signatures, raw secret-bearing payloads or kubeconfigs. HMAC authenticates bytes; HTTPS protects confidentiality.

#### Start discussion and delivery checklist

| Order | Owner | Required action / result |
| --- | --- | --- |
| 1 | B+C+A | Confirm actual listening URL, account identity and trusted credential-to-host lookup; KeyService currently supplies account/key/kind, not proof of host assignment |
| 2 | B | Boot real Java server with reviewed foundational migrations and required replay store; existing extended migrations alone do not bootstrap the DB |
| 3 | C+B | Provision dedicated service key, store encrypted verifier secret and trusted account/agent assignment; deliver secret once by protected channel, never chat/source/logs; test revoke/unknown key |
| 4 | A | Review existing config, sigv1 and auth-smoke; reuse signer; agree module layout before moving into pkg/kanclient |
| 5 | A+B+C | Valid signed GET receives 200 []; stop/restart server to prove backoff/recovery; invalid credentials stop; Ctrl+C interrupts request and wait |
| 6 | C with A | Record unit/mock versus real-server results separately; no-work check does not pass any resource gate |
| 7 | B+D+A/C | Before execution: finalize separate signed idempotent claim, renewal and progress/convergence schemas, fixtures and generated outputs |

Heartbeat currently requires agentVersion, dockerVersion and supportedKinds. Agree a truthful dockerVersion source before sending it; do not invent a fake version or advertise unavailable drivers. The initial signed GET/auth-smoke can establish communication before heartbeat integration. With empty capabilities, unexpected nonempty work is a protocol/capability error: no dispatch, no success acknowledgment and no cleanup. Empty/failed polls never authorize deleting local objects.

The agent never writes platform SQL or selects its account from a request/label. B validates assignment from verified service identity; C controls eligibility/custody; A implements isolated clients. Exact assignment storage/provisioning is still B+C+A review, not silently solved by these defaults.

#### Exit criteria and next step

A real server accepts the valid signature and returns 200 []; an outage triggers bounded backoff and recovery; invalid credentials produce a safe stop; cancellation exits promptly; credentials are absent from logs. C records the revision/environment/results. Claim and execute nothing in this milestone. Then implement one owned Droplet lifecycle only after claim/network protocol review, with C's crash/fencing evidence; E continues the CLI without agent ownership.

## 3. Client rules

| Concern         | Requirement                                                        |
| --------------- | ------------------------------------------------------------------ |
| Context/timeout | Caller cancellation, finite HTTP timeout                           |
| Redirects       | Refuse forwarding signed credentials                               |
| Request bytes   | Marshal once; sign/transmit identical bytes                        |
| Response size   | Bounded read; detect overflow                                      |
| Status/JSON     | Non-2xx and malformed data are errors                              |
| Retry           | Fresh signed request; mutations need agreed idempotency            |
| Logging         | No secret structs, canonical real-key payloads or full kubeconfigs |

Coordinate endpoints, retry/idempotency, agent leases and readiness with B. Test using synthetic credentials; a fake verifier is not the deployed auth/store.

## 4. First compute acceptance

Create one owned container-backed Droplet, inspect by approved labels before create, restart after mid-create crash, converge without duplicates and delete only the owned resource. Preserve volumes/unrelated resources. Never describe Droplets as VMs. Server outage or incomplete/empty poll must not trigger cleanup.

Epoch is the fence; convergence ack requires exact generation/owner/epoch and clears lease_owner plus lease_expires_at together, leaving epoch unchanged. Progress is separate and never changes acked_generation. A reviews agent integration; B owns SQL; C owns tests. Lifecycle success mapping uses existing contract strings only.

Agreed target: one agent process per physical host; separate service credential and authenticated polling context per assigned account. B owns server assignment validation; C credential provisioning/rotation/revocation; A isolated clients and shared bounded execution. Provisioning/storage details still need B+C+A review; no new wire fields.

## 5. ITISS parallel track

| CORE-5 phase | Evidence                                                                             |
| ------------ | ------------------------------------------------------------------------------------ |
| Foundation   | AlmaLinux VM/version/resource budget, threat model, real signing/replay tests and CI |
| Compute      | Privilege/mount review, crash/fence/isolation tests with C                           |
| Network      | Pool/allocator decision, private firewall and actual denied paths                    |
| Kubernetes   | One-day host spike: cgroup v2, SELinux, firewalld, memory and tested k3d+k3s flags   |
| Data         | Owned volumes/credentials; measured backup/restore and cross-account denial          |
| Install      | Vault/no_log/checksums/systemd hardening; second run changed=0                       |

Publication/TLS/tunnel, metadata, metrics/alerts/incidents and dynamic inventory remain Extended track evidence in the roadmap. A begins extended Go only after all CORE-5 gates pass. If A is one week behind, all extended work stops.

Ask mentor in writing whether ITISS students may own agent/driver Go; if not, name a developer and record handover. Confirmation remains pending.

## 6. Learning sequence

Study structs/pointers/interfaces/errors, HTTP/JSON/readers, contexts/timers/signals, tests and then concurrency. Inspect one fixture/result before coding. Use Go documentation and the prior learning guide if present; V5 does not replace code-level walkthroughs.

## 7. Dated targets

| Period                          | A checkpoint                                                      |
| ------------------------------- | ----------------------------------------------------------------- |
| 14–31 October 2026              | Client/main/poll/auth-smoke and schema review                     |
| 1–20 November 2026              | Container/reconcile/fencing integration; revalidate early capacity spike   |
| 21 November–5 December 2026     | Tests 1–6 in CI and network helper/allocator integration          |
| 6–20 December 2026              | Required bounded K8s and tests 7–10                               |
| 21 December 2026–5 January 2027 | PostgreSQL/MinIO process lifecycle and restore with C/B           |
| 6–10 January 2027               | One stack and Ansible Part 1                                      |
| 11–25 January 2027              | Freeze, clean-VM test, two-OS CLI evidence, viva and backup video |

No new CORE-5 features after 10 January. Never cut tests 1–15; cut extended work first and report any unmet gate honestly. Final deadline: Monday, 25 January 2027.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
