# Anuj and Person C: integration handoff

Implemented baseline as of 11 October 2026. The original checkout is untouched
because it is outside this session's writable roots. Work is in `Claude/kan-cloud-fixed`;
apply `Claude/Kan-Cloud-Code-Fixes.patch` to the original as described in the report.

## Breaking changes to integrate together

- `host`, `x-kan-access-key`, `x-kan-date`, `x-kan-nonce` are now mandatory signed
  headers; nonempty bodies also require `content-type`. Old signatures fail.
- `Sign` in Go now returns **four values**, including an error; canonical builder
  returns three including an error. Always handle the error before sending.
- Java `KeyService.getSecretKey` becomes `findActiveCredential`, returning
  `Credential(accountId, accessKeyId, kind, secretKey)` or null. Agent kind is
  exactly `service`. Do not configure a development fallback credential.
- New public `BearerTokenService.findActivePrincipal` returns an active user
  principal or null. Person C owns hash-only lookup and expiry/revocation.
- Observed reports require `observedAt`; observedGeneration is nullable/optional
  for health reports. Report an applied generation only after successful convergence.
  K8s observed state is `running`, while its desired state remains `active`.

## Who continues where

| Owner | Next code | Acceptance |
|---|---|---|
| Anuj | `kan-agent/cmd/kan-agent/main.go`, `internal/client/`, polling then Docker driver | Signed heartbeat/work/report cycle; retries use fresh requests; server outage cannot delete workloads |
| Person C | Implement `KeyService` and `BearerTokenService` beans | Active service keys only internally; active user tokens publicly; revoked/expired credentials fail; no secrets in logs/git |
| Person C | AES-GCM secret storage, token hashing, public permissions, auth integration tests | Wrong account cannot read/write another account; wrong credential kind fails |
| Java core owner | Application main, Postgres migrations, server configuration and controllers | Main starts with auth configuration; every controller scopes by `kan.principal.accountId` |
| Java core owner + Anuj | Observed report ordering and apply semantics | Reject future generation; ignore stale generation/timestamp; health reports do not acknowledge desired work |
| Anuj + Person C | Run Go smoke against Java and shared real Valkey | Valid 200; replay/tamper 401; invalid MAC does not consume nonce; Valkey outage fails closed |

`AgentAuthConfiguration` registers agent auth on `/internal/v1/agent` and children,
and Bearer auth on `/v2`, `/kan/v1` and their children. Place the application main
under `com.ciagotech.kan.server` so it discovers this configuration. The rules default
assumes starting from `kan-server`; otherwise set `kan.sigv1.rules-path` explicitly.
Implement BOTH required service beans: startup should fail if either is missing.
No application main or controllers have been invented in this baseline.
Remove Boot plugin `<skip>true</skip>` once the application main is added.

## First integration session

1. Anuj and C review `CANONICAL-REQUEST.md`, rules and the three shared vectors.
   Land contracts/Go/Java together; do not independently deploy the old signer.
2. Each runs the verification commands below and resolves CI before integration.
3. Java owner starts Postgres/Valkey using configured local secrets, then Java.
   Configure datasource URL/user/password and Redis host/port/password outside git.
   On the host, Valkey is `127.0.0.1:6390`; on Compose backend network it is `valkey:6379`.
4. C provisions an active `service` credential for a test account, sharing the secret
   privately with Anuj. Obtain a separate `user` token for public API tests.
5. From `kan-agent`, set `KAN_SERVER_URL`, `KAN_ACCESS_KEY`, `KAN_SECRET_KEY` in the
   local environment without committing them, then run:

```bash
go run ./cmd/auth-smoke
```

The smoke command performs read-only GET `/internal/v1/agent/work` calls and expects:
valid=200, identical replay=401, bad signature=401, valid request with that same nonce
(after the rejected bad signature)=200, tampered nonce=401. It disables redirects,
uses a 10-second timeout and requires HTTPS outside localhost. It prints only test
names/statuses. It is NOT the full Docker agent.

6. C tests public routes with no token (401), revoked token (401), service identity
   (401), valid user token (accepted), wrong account resource access (403/404).
7. Test two server instances against the same Valkey, restart/replay, nonce expiry,
   and Valkey outage. These real-service tests remain pending.

## Anuj: after authentication integration

Build signed heartbeat -> work polling -> observed report first, against the Java
contract. Start with one compute resource. Add bounded timeouts/backoff, immutable
body signing, fresh nonces per retry and account-labelled Docker objects. Never
reuse credentials on a redirected request; the provided smoke client rejects redirects.
Do not implement automatic orphan deletion until known-ids is authenticated,
account-scoped and safe during outages. Record VM/tool versions before the first demo.

## Checks and what has actually passed

From repository root:

```bash
python3 contracts/sign_example.py
(cd kan-agent && go test -race ./... && go vet ./...)
(cd kan-server && mvn --batch-mode verify)
```

Locally: Python three vectors PASS; Go race tests and vet PASS with installed Go
1.27.1 using `GOTOOLCHAIN=local`; all Java sources compile and 14 Java test methods
PASS using cached dependencies and a temporary direct test runner. That runner is
not Maven/JUnit-engine execution. Go smoke tests use an in-process HTTP transport,
not a running Java server. YAML parses and Compose config validates using temporary
placeholder image references/secrets, without deploying anything.

The pinned Go 1.27.2 toolchain download and Maven's Boot 3.5.16 parent download are
blocked by network/DNS restrictions. Full pinned-toolchain/Maven verification,
GitHub workflow execution, Docker deployment and live Go->Java/Valkey tests have
NOT passed yet. Run them in CI or the development VM with network access.

## Through 15 December 2026

By 16 October finish auth/CI/VM baseline; by 23 October demonstrate one real compute
create/report/delete. Extend the Go driver and provider APIs through November.
Use 1 December as feature freeze, 8 December as release candidate, 11 December as
final rehearsal, and 13–14 December as buffer. Track remaining scope in the existing
`Kan-Cloud-Project-Workflow.md`; this patch is the security/build baseline, not all
38 audit findings or a finished emulator.
