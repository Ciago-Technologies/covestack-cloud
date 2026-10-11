# Covestack Cloud — Technical Design, corrected 11 October 2026

| Item | Details |
|---|---|
| 1 | Submission: **Tuesday, 15 December 2026**. This replaces contradictory earlier starters; it describes the target, not features already implemented. |
| 2 | Read [current setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) for exact repository tasks, [workflow](Covestack-Cloud-Project-Workflow.md) for dates, and [audit](Covestack-Cloud-Documentation-Cleanup.md) for the original evidence. The audit remains a historical record; its old document line numbers refer to the preserved originals. |

## 1. Actual baseline and document authority

| Item | Details |
|---|---|
| 1 | The security/build baseline is committed as `349a2f4` (Fixed security issues) on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`. The lifecycle changes are currently applied but uncommitted. Remote main/PR merge status has not been verified; review and merge through the shared PR. Three signature vectors, negative tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates and auth-smoke exist. The config.go file currently contains only `package config`; it is a placeholder, not a working loader. Real credential/token lookup, application main, migrations/controllers, full agent and live deployment tests remain pending. |
| 2 | Product name: **Covestack Cloud**; intended repository slug: `covestack-cloud`. The actual checkout and GitHub remote are still named `kan-cloud` until the repository owner renames them. Use the real paths in commands. Existing module folders, Go import paths, Java packages, signing headers/algorithm, KAN_* environment names, API paths, labels, database identity and Valkey key prefixes remain compatibility identifiers. A future technical rename needs a separate coordinated migration. |
| 3 | Keep the actual layout: `kan-server/`, `kan-agent/`, `contracts/`, `deploy/`, `docs/`. `agent/` and `control-plane/` are empty placeholders, not active modules. Avoid a module migration before the deadline. Add `kan-cli/` and `compat/` when needed. |
| 4 | After coordinated review, versioned `contracts/` files define wire behavior. This design defines intended changes until those contracts are corrected. A documentation update alone does not change signing rules or make an unsafe server safe. Any protocol change updates Java, Go, examples and negative tests in the same PR. |

## 2. Architecture and release scope

| Item | Details |
|---|---|
| 1 | Java receives API requests, authenticates them, authorizes account access, stores desired state and actions in Postgres, and handles scaling decisions. The Go agent runs beside Docker, polls signed work and makes actual containers match desired state. It never queries Postgres. The CLI calls APIs; its local `up/down` commands invoke Compose. |
| 2 | The agent is trusted host-management software. Running it on the host rather than in a container does not reduce the authority of access to the Docker socket. Restrict that authority to the agent user and keep the VM boundary intact. |

| Release component | December target | Boundaries |
|---|---|---|
| Droplets, actions, account, catalogs, tags, SSH keys | Required foundation | Numeric Droplet/action IDs; key-only SSH |
| ASG, CLI, health, Mailpit | Required foundation | Account-scoped; real healing evidence |
| Managed Postgres | Required application stack | Persistent data and reachable connection metadata |
| Spaces | Required application stack | Persistent MinIO and scoped S3 credentials |
| HTTP load balancer | Planned bounded implementation | One HTTP rule, healthy Droplet targets |
| Kubernetes | Planned bounded implementation | One supported k3s version and node-pool shape |
| Snapshots | Planned bounded implementation | Container filesystem snapshots; mounted data excluded |
| MySQL, additional DBs, advanced LB/DOKS features | After required paths are stable | Implement explicitly or return documented 501 |
| AWS/GCP/Azure and registry | Outside this submission | No new provider work |

| Item | Details |
|---|---|
| 1 | The workflow schedules the planned features; a feature missing its stability gate must be clearly identified as unsupported, never disguised as operational. No full DigitalOcean parity claim is made. |

## 3. Toolchain, builds and configuration

| Item | Details |
|---|---|
| 1 | Use Java 25 throughout the team and a maintained Spring Boot release explicitly supporting it. Prefer the compatible Boot 3.5 line to a major-framework migration, subject to support/security verification when pinning the exact patch. Boot 3.3.5 is not the target. Keep Spring MVC, JdbcClient, constructor injection, records and Flyway; no WebFlux/JPA migration is needed. |
| 2 | Pin Maven wrapper, Go toolchain, dependencies, action commits, tools and image digests. Record actual tested versions in `deploy/versions.env` and CI. Use a supported Go release; 1.22 is a historical language minimum, not a recommended installed toolchain. Upgrade the inspected 1.27.1 installation to the verified security patch 1.27.2 or a later reviewed patch before release. |
| 3 | Add PostgreSQL Flyway support (`org.flywaydb:flyway-database-postgresql`) aligned with Boot, and add the mail starter when implementing email. Record unsupported database/version combinations as build blockers. Never use unresolved placeholder versions or digests in a runnable file. |
| 4 | Local host configuration uses Postgres `127.0.0.1:5432`, Valkey `127.0.0.1:6390`, Mailpit SMTP `127.0.0.1:1025`. Container configuration uses service names and container ports, especially Valkey `6379` and Mailpit `1025`. |

## 4. Authentication and replay protection
### 4.1 Authentication matrix

| Route | Authentication | Authorization |
|---|---|---|
| `/internal/v1/agent/*` | kan sigv1 | Active service credential, assigned account/agent |
| `/v2/*` | Opaque Bearer token | Account-scoped resource/function permissions |
| `/kan/v1/*` | Opaque Bearer token | Account-scoped management permissions |
| `/kan/v1/health` | Explicit public exception | Minimal liveness only; no secrets or detailed topology |
| Detailed doctor/readiness API | Bearer token | Authorized account/admin scope |

| Item | Details |
|---|---|
| 1 | For December, CLI management calls use Bearer tokens. CLI does not sign DO endpoints with sigv1. Its agent credentials are separate from user Bearer tokens. Public access-key IDs are identifiers, not Bearer secrets. |
| 2 | Bootstrap is a one-time local administrative command that creates the initial account and random token, writes them to a restricted file and reports only the file path. It is not a public unauthenticated key-creation API. Later key issuance/revocation requires authenticated account permission. Test revoked credentials immediately, including invalidation of caches. |

### 4.2 kan sigv1 canonical request

| Item | Details |
|---|---|
| 1 | Use **simple HMAC-SHA256** with the raw signing secret. No four-step derivation and no AWS-style Authorization parsing. |

```text
CanonicalRequest = METHOD + "\n" + RAW_PATH + "\n" + RAW_QUERY + "\n"
                 + CANONICAL_HEADERS + "\n" + SIGNED_HEADER_NAMES + "\n"
                 + SHA256_HEX(RAW_BODY)
StringToSign = "KAN1-HMAC-SHA256\n" + X_KAN_DATE + "\n" + SHA256_HEX(CanonicalRequest)
Signature = HMAC_SHA256_HEX(SecretKey, StringToSign)
```

| Item | Details |
|---|---|
| 1 | Mandatory signed headers: `host`, `x-kan-access-key`, `x-kan-date`, **`x-kan-nonce`**. Sign and require `content-type:application/json` for a non-empty JSON body. Sort names, normalize ASCII header names and specified whitespace, and use UTF-8. Reject missing/duplicate mandatory headers and invalid signatures instead of silently omitting them. Define raw escaping/query behavior in the contract; sign the bytes sent over HTTP. |
| 2 | Transmit `X-Kan-Access-Key`, `X-Kan-Date`, `X-Kan-Nonce`, `X-Kan-Signature`. Require fixed-format UTC dates, a bounded nonce such as 32 lowercase hex characters from 16 random bytes, and a 64-character hex signature. Existing signatures change when nonce/key headers become signed; regenerate rather than copy the old expected values. |

### 4.3 Verification order and request limits

| Item | Details |
|---|---|
| 1 | Register the filter explicitly. Cover the route matrix and test its actual servlet mappings. |
| 2 | Enforce ingress rate/read-time limits and validate bounded header syntax. |
| 3 | Check Host and timestamp (maximum skew 300 seconds). |
| 4 | Check declared body length; read at most `MAX_BYTES + 1` and return 413 if above 1 MiB. Bound chunked input too. |
| 5 | Resolve an active credential with account ID, kind and permissions; reject user keys on agent routes. |
| 6 | Build the canonical request from the actual bytes and verify with a constant-time primitive. |
| 7 | Atomically claim the authenticated nonce using Valkey `SET NX EX 600`; replay or store failure rejects the request before side effects. |
| 8 | Attach `kan.principal`, enforce ownership and invoke the controller with the cached body. |
| 9 | Invalid signatures and unknown keys never consume nonce storage. No `X-Dev-Skip-Auth`, arbitrary nonempty-token acceptance or TODO-pass branch exists in an integration/demo profile. Fakes belong to explicit test profiles. |
| 10 | Use a named Valkey data volume, AOF, `noeviction`, bounded storage and monitoring. Configure and test restart persistence for the validity window. AOF must meet the required durability policy; loss of replay history blocks agent authentication until recovery or until all pre-loss timestamps have expired. Boot readiness includes replay-store readiness. In-memory replay storage is bounded/expiring and test-only. |

### 4.4 Secrets, permissions and logging

| Item | Details |
|---|---|
| 1 | Keep AES-GCM master keys outside tracked YAML: use a restricted secret file or deployment secret mechanism. Use fresh random IVs, versioned ciphertext and authenticated account/key context. Protect master-key backups and rotation separately from database backups. |
| 2 | Use a credential record, not just `String getSecretKey()`. Enforce account/agent ownership for work, observations, events, SSH keys, known IDs and LB members. Reserve ownership labels; users cannot change `kan.io/account`, resource ID or managed flags. |
| 3 | Store opaque Bearer token verification hashes; signing secrets require encryption because verification needs the secret. Never log body secrets, signature credentials, database URIs with passwords, kubeconfigs or decrypted keys. Audit accesses with account/key/action/resource/outcome/request ID. Make audit writes append-only through DB permissions; an unconstrained `events` table is not inherently immutable. |

## 5. Contract and state vocabulary
### 5.1 Shared lifecycle contract

| Item | Details |
|---|---|
| 1 | Version1.1.0 is defined in contracts/lifecycle-rules.json, with generated FROZEN-STATES.md, kind-discriminated OpenAPI schemas and Go/Java catalogs. Read the [complete lifecycle reference](../../contracts/FROZEN-STATES.md) for every resource, desired/observed state, transition, success mapping and metadata policy. |
| 2 | Seven resource kinds cover compute.instance, db.instance, k8s.cluster, compute.group, compute.snapshot, network.load-balancer and storage.bucket. compute.group and storage.bucket are server-managed; five execution kinds may be agent work only when the agent advertises an installed/tested driver in heartbeat supportedKinds. Defined contracts do not mean those drivers/controllers are implemented. |
| 3 | Operation statuses are queued/running/completed/failed/cancelled; cancellation is queued-only. A failed attempt is terminal and retry creates a linked new attempt. Notification delivery is queued/delivering/delivered/failed with bounded outbox retry. Accounts, credentials, SSH keys, tags, catalogs, CLI/health and S3 objects have separate metadata/authorization policies rather than invented agent kinds. |
| 4 | Deletion intent requires deleting=true: archive for instances, deleted otherwise. Do not acknowledge generations for progress/degraded/error. Require matching kind, UTC observedAt and current generation on successful convergence. DB resize/migrate/ fork states are reserved and rejected until separately implemented/contracted. Public provider status mappings are D's reviewed DTO boundary, not raw internal enums. |

### 5.2 Endpoints and schemas

| Item | Details |
|---|---|
| 1 | Keep the five agent endpoints: heartbeat, work, observed, known-ids and events. Freeze heartbeat at `200 {"ok":true}`. Additional signed `/ssh-keys` and `/lb-members/{id}` endpoints require schemas and ownership rules before use. |
| 2 | Work excludes server-only groups and terminal resources already acknowledged. Describe kind-specific specs, int64 nonnegative generations, bounded arrays/strings and reserved labels. Use a discriminator or equivalent validation. Publish required signing headers, security requirements and 401/403/413/429/error responses in OpenAPI. |
| 3 | Observations require `kind`, `observedAt` as UTC date-time and a kind-valid state. Agent epoch and monotonically increasing report sequence are additional server/agent ordering requirements that B/A/C must implement and add to the wire schema together; they are not claimed present in version 1.1.0. `observedGeneration` is optional/nullable for health-only reports. Null does not acknowledge work. Response is `200 {"applied":true&#124;false}`; ignored stale input is not a worker error. |
| 4 | Known IDs are a complete account/agent-scoped live-resource snapshot. If paginated, include snapshot identity/completeness and fetch all pages. A timeout, partial response or decoding failure never triggers orphan deletion. Require grace periods and managed/account/agent labels before cleanup. |
| 5 | Validate every example against the schema and validate all state transitions in both fakes and real paths. Bind labels, resource ID, snapshots and SSH keys to the authorized account. |

## 6. Persistence, generations and actions
### 6.1 Database model and transactions

| Item | Details |
|---|---|
| 1 | Use Postgres with resources, accounts, credentials, operations, idempotency, events and agent heartbeat/lease records. Existing schema prose is a design, not a migration that has already run. Add real Flyway migrations and constraints. |
| 2 | Use internal ULIDs; Droplet/action/public image IDs are numeric in DO responses. Kubernetes/DB/LB IDs use provider-shaped UUIDs. Never expose internal IDs through DO DTOs. Store explicit deleted_at rather than using updated_at for garbage collection. Define parent/child retirement and deletion order. |
| 3 | Server writes desired state/spec/generation/version with account predicate and optimistic locking. Agent observation writes touch only observed columns and include account/agent scope. Use transactions for resource/action creation, action locking, acknowledgements and idempotency; do not issue whole-row updates. |
| 4 | Derive Droplet `locked` from open operations; do not write a nonexistent `resources.locked` column. Use `access_keys.status='revoked'` consistently until an intentional migration changes the model. Put tokens in the idempotency table rather than assuming a resources.client_token column exists. |
| 5 | Atomically reserve `(account_id, scope, token)` and store a request fingerprint plus replayable result in the same transaction as creation. Concurrent identical requests return one result; token reuse with a different payload is rejected. |

### 6.2 Observation acceptance

| Item | Details |
|---|---|
| 1 | Lock/read the scoped resource and validate agent epoch/sequence and timestamp skew. Reject future/negative/impossible generations. Reject older-generation reconcile reports from changing current state/detail; timestamp ordering alone cannot make obsolete generation data current. |
| 2 | Health reports with null generation inspect current actual state and do not advance acknowledged generation. Successful reconcile reports monotonically advance acknowledged generation only after the requested physical state is reached. Store attempted failures separately with error/retry metadata. |
| 3 | A failed Ensure never sets successful observedGeneration to the target. Mark the matching operation attempt failed and schedule bounded retries/backoff for retryable failures. Complete an action only when its target generation and expected physical state/action result are confirmed. Test older-generation reports with newer timestamps, equal timestamps, delayed observations, failure and intentional shutdown. |

### 6.3 Reboot and deletion

| Item | Details |
|---|---|
| 1 | Active-to-active changes do not express reboot. Carry a restart generation/operation intent and persist its executed marker so retries acknowledge the already performed restart. Verify actual start-time/process change and define crash/retry semantics; do not claim universal exactly-once execution across arbitrary failures. |
| 2 | Soft delete requests archive/deleted by kind; the agent removes the workload and acknowledges its terminal state. Hide accepted deletions in public lists consistently. Garbage collection waits for acknowledged deletion, uses deleted_at and retains operation history. Removing a row before teardown acknowledgement is forbidden. |

## 7. Agent correctness and scaling

| Item | Details |
|---|---|
| 1 | One assigned host agent manages the December deployment. Enforce a renewable agent lease and fencing epoch; a second agent cannot independently mutate the same resource. Containers carry managed, account, agent, internal resource and external ID labels. |
| 2 | Rebuild ports from every managed running/stopped container's HostConfig.PortBindings before allocating. Port ownership includes resource, container port and protocol. Probe the intended bind address; handle Docker bind failure/reallocate safely, release only after teardown. A closed probe socket is not a permanent reservation. |
| 3 | Use sharded workers, in-flight deduplication and non-blocking dispatch. Docker events mark dirty resources for their own shard; they never race a direct observed-state write. Preserve dirty state until inspection/report succeeds; coalesce observe jobs and retry on failure. Events are informational audit input, not authoritative state mutations. |
| 4 | Poll changed work every 2 seconds and inspect converged live workloads periodically, such as every 15 seconds. Reconnect event streams with backoff. Apply timeouts to API calls, pulls, driver operations and readiness checks. Distinguish intentional stops/deletes from crashes. |
| 5 | ASG counts healthy plus pending reservations within a startup deadline, binds children to parent/account, and serializes decisions per group. Enforce min/desired/max and scale-down. During agent outage, stop replacements and surface degraded health; a stale heartbeat is not proof all workloads died. Confirm failures and retire failed children before replacement. Test slow pulls across scheduler periods and recovery after outage. |

## 8. Workload isolation, SSH and limits

| Item | Details |
|---|---|
| 1 | Backend Postgres, Valkey and server use a dedicated backend network. User Droplets, databases, LBs and k3d use workload/account networks. User DBs never join the backend. Shared workload networks allow peer reachability; use per-account isolation and allow only intended LB/backend connections. |
| 2 | Disallow privileged Droplets, Docker socket mounts, host-network/host-PID modes and unapproved bind mounts. Enforce CPU, memory, PID and resource-count quotas; reject zero/unbounded limits for workload requests. Drop capabilities and add back a reviewed minimum for the chosen image, keep no-new-privileges and default confinement. Conventional sshd needs a tested capability policy; NET_BIND_SERVICE alone is not assumed sufficient. |
| 3 | Build/map supported image slugs to immutable SSH-ready images; plain ubuntu is not a running SSH server. Disable password authentication and default root passwords. Validate public keys, write authorized_keys safely with correct permissions and regenerate per-instance host keys. Invoke host tools with validated argument arrays, never interpolate request data into a shell. |
| 4 | Bind published ports explicitly and allow the demo subnet through narrow firewall rules. Test isolation by backend IP, host gateway, published ports and another account's workloads; DNS lookup failure alone is not evidence of blocked access. A container is not a full VM security boundary. |

## 9. Databases, Spaces, load balancers and snapshots
### 9.1 Managed database and Spaces data

| Item | Details |
|---|---|
| 1 | Attach a labeled named volume to each database's actual data directory and retain it across reconciliation/recreation. Define explicit user deletion retention and backup/restore. Configure the returned user, password and database in the actual container; setting only POSTGRES_PASSWORD does not create doadmin/defaultdb. |
| 2 | Encrypt stored DB credentials; authorized read/connection APIs may return credentials intentionally. Other lists/logs mask them. Report a configurable client-reachable hostname and mapped port rather than always localhost. Gate online state on actual connection/readiness. |
| 3 | Mount MinIO's volume at `/data` when its command is `server /data`. Keep root S3 credentials server-side and issue separate account-scoped credentials/policies. storage.bucket lifecycle is server-managed, never agent work; B/E reconcile owned buckets through trusted MinIO administration and C enforces policy. Prefix internal bucket names and verify ownership; prefixes alone do not authorize S3 access. Clients use the scoped physical bucket name supplied by Kan, or a translating gateway if logical names must be retained. |

### 9.2 HTTP load balancers

| Item | Details |
|---|---|
| 1 | Start with one HTTP forwarding rule. Generate a valid include under `/etc/nginx/conf.d/` or a complete nginx.conf with events/http. Mount a persistent configuration directory so atomic file replacement is visible. Validate config with nginx -t before reload; preserve the previous valid version on failure. |
| 2 | Reconcile membership on tagged/explicit member creation, removal, deletion, health or IP changes. Choose only ready members and define empty-pool behavior. Validate numeric ports/addresses/identifiers. Use HTTP-ready workload images for demo responses; SSH alone does not provide port 80. |

### 9.3 Kubernetes and snapshots

| Item | Details |
|---|---|
| 1 | Use one pinned tested k3s version, validated cluster names, bounded node counts and a workload network separated from backend services. k3d's required elevated privileges are a separate trusted feature boundary, not compliant with the ordinary Droplet capability profile. Describe that limitation. Export a client-reachable API endpoint, protect kubeconfig and merge or save separately; never silently overwrite a user's existing file. |
| 2 | Initiate snapshots through `POST /v2/droplets/{id}/actions` with `{"type":"snapshot","name":"..."}`. Map snapshot public image IDs to scoped immutable images. Container commit excludes mounted-volume data and VM memory; snapshots are not database backups. Prevent cloned instance host keys/credentials from becoming shared identity. |

## 10. DNS, TLS and public routing

| Item | Details |
|---|---|
| 1 | Use Nginx on the VM as the only public API proxy. Keep the app on guest loopback for the host-agent private connection and proxy; deny `/internal/` in public routing. Keep control-plane DB/cache private. A loopback-published service cannot receive VMware forwarding addressed to the VM interface. |
| 2 | dnsmasq listens on VM loopback plus the actual NAT interface IP, with DNS-only mode and subnet-restricted TCP/UDP 53. Forward DNS TCP and UDP and HTTPS 443 to the correct guest IP; a DHCP reservation/static guest address prevents stale forwarding rules. Resolve the demo domain to the laptop's hotspot IP, which clients can actually reach. |
| 3 | Preserve the validated request authority in proxying, including nondefault ports when signed. Use fixed configured host allowlists, trusted proxy headers and private internal routes. Publish Mailpit through a separate restricted TLS listener/location rather than assuming its HTTP port already speaks HTTPS. |
| 4 | Issue a DNS-01 certificate before rehearsal and configure renewal/reload; certificates expire. Protect DNS provider tokens in credential files, never shell history or tracked YAML. mkcert fallback requires installing only the public CA certificate on owned client devices; never distribute its private root key. |

## 11. CI, tests, operations and truthful evidence

| Item | Details |
|---|---|
| 1 | CI builds Java and Go on every PR, validates contracts/examples, runs signing positive/negative vectors, filter tests and integration tests, then scans secrets/dependencies/images. Pin action commits and verify downloaded tools; grant minimal workflow permissions. Pin images by digest and schedule reviewed updates. Supply-chain scans/signatures are evidence, not proof no vulnerability exists. |
| 2 | Required negative tests: missing/duplicate signed headers; mutated nonce/key/path/query/body; identical replay; invalid signature does not poison nonce; concurrent replay; stale/future dates; revoked/user keys; oversized chunked body; missing principal; cross-account access; store outage/restart; path routing and signed Host through proxy. |
| 3 | Lifecycle tests assert bounded readiness/action completion and check every SDK error. No noisy 5% whole-request timing test; use vetted constant-time comparison. Hardening scripts accumulate failures and exit nonzero, and report only the controls actually checked. Do not call a limited script full CIS compliance. |
| 4 | Detailed doctor checks are server-side for remote clients, not localhost DB/socket assumptions. Generate request IDs and clear logging context after each request. Notification delivery happens after commit via a retryable outbox, so SMTP failure does not roll back a valid create or claim success before commit. |
| 5 | Keep operational runbooks for key rotation, agent outage, stuck work, backup restore, network isolation and certificate renewal. Preserve data; no default docker compose down -v or blanket prune of snapshot images. Rehearse on a clean VM and second device. See the workflow's dated gates. |

## 12. Official references and implementation status

| Item | Details |
|---|---|
| 1 | [Spring Boot 3.5 Java support](https://docs.spring.io/spring-boot/3.5/system-requirements.html). |
| 2 | [Go releases and supported-version policy](https://go.dev/doc/devel/release). |
| 3 | [Flyway PostgreSQL module](https://github.com/flyway/flyway/blob/main/documentation/Reference/Database%20Driver%20Reference/PostgreSQL%20Database.md). |
| 4 | [Docker socket authority](https://docs.docker.com/engine/install/linux-postinstall/). |
| 5 | [Docker published-port behavior](https://docs.docker.com/engine/network/port-publishing/). |
| 6 | [DO snapshot actions](https://docs.digitalocean.com/products/snapshots/how-to/snapshot-droplets/). |
| 7 | [Docker commit limitations](https://docs.docker.com/reference/cli/docker/container/commit/). |
| 8 | All remedies above remain implementation requirements until repository PRs and acceptance checks pass. The current setup fixes document tracks each of the 38 audit items without marking documentation changes as code fixes. |
