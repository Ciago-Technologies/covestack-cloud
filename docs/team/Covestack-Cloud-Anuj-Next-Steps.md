# Anuj: verified setup and next work — 11 October 2026

Checked original `/run/media/anujtiwari/DevMount/Project-CDAC/kan-cloud` without
printing or sourcing secrets. This is the current state, superseding the earlier
handoff's historical statement that the patch had not been applied.

## Verified findings

| Item | Current result |
|---|---|
| Security/build patch | Committed as349a2f4; published fix/security-baseline branch |
| deploy/secrets/agent.env | Exists, mode 600; parent directory mode 700 |
| Agent file Git handling | Ignored; no deploy secret files tracked |
| KAN_SERVER_URL | Set to a localhost HTTP origin; value not displayed |
| KAN_ACCESS_KEY / KAN_SECRET_KEY | Both EMPTY |
| deploy/.env | Missing |
| deploy/secrets/postgres_password | Missing |
| deploy/secrets/valkey.conf | Missing |
| Python shared vectors | All 3 PASS in original checkout |
| Go race tests and vet | PASS with installed Go 1.27.1, GOTOOLCHAIN=local |
| Pinned Go toolchain | go.mod requests Go 1.27.2; exact pinned-version verification pending |
| Fedora / Java / Compose | Fedora44; Java 25.0.4.1; Compose5.6.0 |
| Java application/controllers | No application main or agent controllers found |
| Credential/token persistence | Interfaces exist; concrete implementations absent |
| Actual Docker/server status | Not verified; sandbox denied Docker API and socket inspection |

The earlier configuration check did not modify the project. Current branding changes are prepared separately in Claude/covestack-cloud. A permission error in this
assistant sandbox does NOT establish your Fedora terminal lacks Docker permissions.
The Compose config check fails immediately because deploy/.env does not exist.

## Work order for Anuj

1. Your baseline is committed/published on fix/security-baseline. Review/merge its PR; do not reapply the
   patch. Check git status and staged changes. Explicitly stage code/contracts and
   documentation; never force-add ignored credentials. Run normal pinned-toolchain
   tests in your Fedora terminal/network or CI. Local fallback tests on1.27.1 are
   supporting evidence, not verification of1.27.2.
2. Build config loading in kan-agent/internal/config/config.go. Read server URL,
   access key, secret and rules path. Fail on missing credentials. Validate origin,
   require HTTPS outside localhost, and never print secrets. The executable reads
   process environment; it does not automatically source agent.env.
3. Build typed HTTP client in kan-agent/internal/client/client.go and types.go.
   Reuse sigv1.NewSignedRequest, disable redirects, use finite HTTP timeouts and
   bounded response reads. Follow contracts/internal-api.yaml for heartbeat, work,
   observed, events and known-ids. Generate a fresh signed request for each retry.
   Do not blindly retry mutating operations until B confirms idempotency semantics.
4. Test the client without Java using an injected HTTP transport or httptest on your
   local machine. Verify non-2xx handling, timeout/cancellation, malformed JSON,
   response body bounds and fresh nonces. A mocked successful response is not a
   real Java/Docker integration pass. Do not implement a production auth bypass.
5. Add kan-agent/cmd/kan-agent/main.go. Load config, use signal cancellation, send
   signed heartbeat, then poll work with bounded backoff and jitter. No inbound
   agent port is required. Initially observe/log IDs without changing Docker.
6. Pair with C/B for live Java integration. Provision active account-bound service
   credentials, run auth-smoke, validate public user-token separation and revoked
   keys. Confirm lease/agent fencing before concurrent Docker execution.
7. Define the first compute spec with B and D before implementing the Docker driver:
   trusted image, ports, reserved ownership labels, desired states and readiness.
   Add internal/driver/compute/instance.go and internal/reconcile/worker.go only for
   one resource type first. Ensure by inspect/update instead of creating a new
   container on every poll. Avoid simultaneous work for the same resource.
8. Report observedAt on observations. Acknowledge observedGeneration only after
   successful convergence. On Ensure failure, keep generation unacknowledged and
   report error according to the contract. Never delete on failed/empty work poll.
9. Acceptance demo: create one resource -> ready -> stop/start -> agent restart
   without duplication -> intentional deletion of that owned resource. Include
   server/Valkey outage and wrong-account tests before expanding drivers.

## Infrastructure to finish alongside the client

From repository root, create deploy/.env from deploy/.env.example if absent.
Populate POSTGRES_IMAGE, VALKEY_IMAGE and MAILPIT_IMAGE with verified/tested exact
image references (prefer digest pins). Do not leave them blank or use latest.
Create random local postgres_password and Valkey config from the provided template;
replace its password. Set secret file modes600 and directory700. These credentials
are different from the agent signing secret.

Validate with:

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
```

After validation and image review, start the infrastructure:

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
docker compose --env-file deploy/.env -f deploy/docker-compose.yml ps
```

Starting Compose does not start Java or the Go agent. Java must be configured with
Postgres and Valkey passwords separately. Host addresses are Postgres127.0.0.1:5432
and Valkey127.0.0.1:6390; containers on backend use postgres:5432 and valkey:6379.
The current network publishes to localhost only. Do not expose these ports to the
LAN to make a remote test work. A VM's localhost and Fedora's localhost differ.
Existing Postgres volumes retain their original initialized password; rotate via
Postgres rather than deleting volumes. Avoid docker compose down -v.

## Person C / Person B handoff

| Owner | Required work |
|---|---|
| C | KeyService.findActiveCredential implementation; active service key/account identity; AES-GCM secret storage |
| C | BearerTokenService implementation; hash-only tokens, expiry/revocation and permissions |
| C + Anuj | Shared signing contract/vector review; tampering/replay/poisoning/outage tests |
| B | Spring application, migrations/configuration, typed/scoped agent controllers and report generation ordering |
| B + Anuj | Agent lease/fencing, retries/idempotency and resource convergence semantics |
| Anuj | Go client/executable/reconciler, Docker drivers, infrastructure, CI and VM baseline |

After C provisions the credential and B starts Java, fill the two blank fields in
agent.env. From repo root load your trusted local file and run the smoke command:

```bash
set -a
source deploy/secrets/agent.env
set +a
(cd kan-agent && go run ./cmd/auth-smoke)
unset KAN_ACCESS_KEY KAN_SECRET_KEY KAN_SERVER_URL
```

Expected statuses: valid200, replay401, invalid MAC401, valid request after poisoning
attempt200, nonce tampering401. Real credentials must exist on Java's side.

## Precautions while building

- Keep frontend/public user tokens separate from agent service keys; auth filters
  do not replace account predicates and resource-level authorization.
- Never print env dumps, credential structs, decrypted secrets or full kubeconfigs.
- Keep raw body bytes and escaped path/raw query identical after signing; proxies
  must preserve signed Host. Keep system clock synchronized.
- Never pass unchecked user input into shell commands; prefer Docker SDK methods.
- An agent with Docker socket access can control its host. Use the project VM for
  workload experiments; do not bind the host socket into user workload containers.
- Operate only on resources whose reserved ownership labels match the authenticated
  account/agent. Do not trust user-supplied labels to establish ownership.
- Preserve volumes; deletion must be explicit and scoped. Network/server failures
  must pause work, not trigger orphan cleanup or destructive recovery.
- Do not disable SELinux/firewalls or authentication just to unblock development.
- Record host and VM versions separately. Recheck CI and real VM behavior before
  claiming a feature complete.

## Near-term milestones

Today: review the baseline PR, implement the empty config.go, then add the typed client.
Next session: client tests and read-only polling executable.
With C/B: secure real smoke pass and agreed compute spec.
By16Oct: auth/CI/VM baseline; by23Oct: first real compute flow.
Feature freeze1Dec; RC8Dec; rehearsal11Dec; submit15Dec2026.

References: [Compose environment handling](https://docs.docker.com/compose/how-tos/environment-variables/variable-interpolation/),
[Go HTTP client](https://pkg.go.dev/net/http#Client),
[Docker daemon security](https://docs.docker.com/engine/security/).
