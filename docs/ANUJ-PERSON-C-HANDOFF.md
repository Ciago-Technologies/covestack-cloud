# Anuj and Person C: integration handoff

| Item | Details |
|---|---|
| 1 | The security/build baseline is committed as `349a2f4` (Fixed security issues) on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`. The lifecycle changes are currently applied but uncommitted. Remote main/PR merge status has not been verified; review and merge through the shared PR. Three signature vectors, negative tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates and auth-smoke exist. The config.go file currently contains only `package config`; it is a placeholder, not a working loader. Real credential/token lookup, application main, migrations/controllers, full agent and live deployment tests remain pending. |
| 2 | Product name: **CoveStack Cloud**; repository slug: `covestack-cloud`. The local checkout and origin URL now use `covestack-cloud`. Use that repository folder in commands. Existing module folders, Go import paths, Java packages, signing headers/algorithm, KAN_* environment names, API paths, labels, database identity and Valkey key prefixes remain compatibility identifiers. A future technical rename needs a separate coordinated migration. |
| 3 | Read [team start guide](team/Covestack-Cloud-START-HERE.md). |

## Breaking changes to integrate together

| Item | Details |
|---|---|
| 1 | `host`, `x-kan-access-key`, `x-kan-date`, `x-kan-nonce` are now mandatory signed |
| 2 | headers; nonempty bodies also require `content-type`. Old signatures fail. |
| 3 | `Sign` in Go now returns **four values**, including an error; canonical builder |
| 4 | returns three including an error. Always handle the error before sending. |
| 5 | Java `KeyService.getSecretKey` becomes `findActiveCredential`, returning |
| 6 | `Credential(accountId, accessKeyId, kind, secretKey)` or null. Agent kind is exactly `service`. Do not configure a development fallback credential. |
| 7 | New public `BearerTokenService.findActivePrincipal` returns an active user |
| 8 | principal or null. Person C owns hash-only lookup and expiry/revocation. |
| 9 | Observed reports require `observedAt`; observedGeneration is nullable/optional |
| 10 | for health reports. Report an applied generation only after successful convergence. K8s observed state is `running`, while its desired state remains `active`. |

## Who continues where

| Owner | Next code | Acceptance |
|---|---|---|
| Anuj | `kan-agent/cmd/kan-agent/main.go`, `internal/client/`, polling then Docker driver | Signed heartbeat/work/report cycle; retries use fresh requests; server outage cannot delete workloads |
| Person C | Implement `KeyService` and `BearerTokenService` beans | Active service keys only internally; active user tokens publicly; revoked/expired credentials fail; no secrets in logs/git |
| Person C | AES-GCM secret storage, token hashing, public permissions, auth integration tests | Wrong account cannot read/write another account; wrong credential kind fails |
| Java core owner | Application main, Postgres migrations, server configuration and controllers | Main starts with auth configuration; every controller scopes by `kan.principal.accountId` |
| Java core owner + Anuj | Observed report ordering and apply semantics | Reject future generation; ignore stale generation/timestamp; health reports do not acknowledge desired work |
| Anuj + Person C | Run Go smoke against Java and shared real Valkey | Valid 200; replay/tamper 401; invalid MAC does not consume nonce; Valkey outage fails closed |

| Item | Details |
|---|---|
| 1 | `AgentAuthConfiguration` registers agent auth on `/internal/v1/agent` and children, and Bearer auth on `/v2`, `/kan/v1` and their children. Place the application main under `com.ciagotech.kan.server` so it discovers this configuration. The rules default assumes starting from `kan-server`; otherwise set `kan.sigv1.rules-path` explicitly. Implement BOTH required service beans: startup should fail if either is missing. No application main or controllers have been invented in this baseline. Remove Boot plugin `<skip>true</skip>` once the application main is added. |

## First integration session

| Item | Details |
|---|---|
| 1 | Anuj and C review `CANONICAL-REQUEST.md`, rules and the three shared vectors. |
| 2 | Land contracts/Go/Java together; do not independently deploy the old signer. |
| 3 | Each runs the verification commands below and resolves CI before integration. |
| 4 | Java owner starts Postgres/Valkey using configured local secrets, then Java. |
| 5 | Configure datasource URL/user/password and Redis host/port/password outside git. On the host, Valkey is `127.0.0.1:6390`; on Compose backend network it is `valkey:6379`. |
| 6 | C provisions an active `service` credential for a test account, sharing the secret |
| 7 | privately with Anuj. Obtain a separate `user` token for public API tests. |
| 8 | From `kan-agent`, set `KAN_SERVER_URL`, `KAN_ACCESS_KEY`, `KAN_SECRET_KEY` in the |
| 9 | local environment without committing them, then run: |

```bash
go run ./cmd/auth-smoke
```

| Item | Details |
|---|---|
| 1 | The smoke command performs read-only GET `/internal/v1/agent/work` calls and expects: valid=200, identical replay=401, bad signature=401, valid request with that same nonce (after the rejected bad signature)=200, tampered nonce=401. It disables redirects, uses a 10-second timeout and requires HTTPS outside localhost. It prints only test names/statuses. It is NOT the full Docker agent. |
| 2 | C tests public routes with no token (401), revoked token (401), service identity |
| 3 | (401), valid user token (accepted), wrong account resource access (403/404). |
| 4 | Test two server instances against the same Valkey, restart/replay, nonce expiry, |
| 5 | and Valkey outage. These real-service tests remain pending. |

## Anuj: after authentication integration

| Item | Details |
|---|---|
| 1 | Build signed heartbeat -> work polling -> observed report first, against the Java contract. Start with one compute resource. Add bounded timeouts/backoff, immutable body signing, fresh nonces per retry and account-labelled Docker objects. Never reuse credentials on a redirected request; the provided smoke client rejects redirects. Do not implement automatic orphan deletion until known-ids is authenticated, account-scoped and safe during outages. Record VM/tool versions before the first demo. |

## Checks and what has actually passed

| Item | Details |
|---|---|
| 1 | From repository root: |

```bash
python3 contracts/sign_example.py
(cd kan-agent && go test -race ./... && go vet ./...)
(cd kan-server && mvn --batch-mode verify)
```

| Item | Details |
|---|---|
| 1 | Locally: Python three vectors PASS; Go race tests and vet PASS with installed Go 1.27.1 using `GOTOOLCHAIN=local`; all Java sources compile and 14 Java test methods PASS using cached dependencies and a temporary direct test runner. That runner is not Maven/JUnit-engine execution. Go smoke tests use an in-process HTTP transport, not a running Java server. YAML parses and Compose config validates using temporary placeholder image references/secrets, without deploying anything. |
| 2 | The pinned Go 1.27.2 toolchain download and Maven's Boot 3.5.16 parent download are blocked by network/DNS restrictions. Full pinned-toolchain/Maven verification, GitHub workflow execution, Docker deployment and live Go->Java/Valkey tests have NOT passed yet. Run them in CI or the development VM with network access. |

## Through 15 December 2026

| Item | Details |
|---|---|
| 1 | By 16 October finish auth/CI/VM baseline; by 23 October demonstrate one real compute create/report/delete. Extend the Go driver and provider APIs through November. Use 1 December as feature freeze, 8 December as release candidate, 11 December as final rehearsal, and 13–14 December as buffer. Track remaining scope in the existing `team/Covestack-Cloud-Project-Workflow.md`; this patch is the security/build baseline, not all 38 audit findings or a finished emulator. |

## Complete lifecycle contract integration

| Item | Details |
|---|---|
| 1 | Read [lifecycle reference](team/../../contracts/FROZEN-STATES.md). Source: contracts/lifecycle-rules.json; generated states/API/Go/Java must be committed as one coordinated version1.1.0 change. Required new payload fields: kind on observed reports and supportedKinds on heartbeat. Agent work/report schemas exclude groups and buckets. Anuj validates received resources/deleting intent against installed capabilities and Go lifecycle helpers before dispatch. B validates resource intent/state/acknowledgement with LifecycleValidator in transactions, scopes stored kind/account/agent, returns applied=false for stale reports, rejects future/invalid acknowledgements, and enforces readiness/time/lease ordering separately. C checks ownership/authorization and prevents user labels changing ownership. D maps internal resource/operation statuses to provider responses. E keeps credential/CLI metadata separate from agent resource lifecycles. |
| 2 | Progress/error reports omit observedGeneration or use null; success must match current desired target. Reconcile observations may skip reachable intermediate states, but terminal resources cannot resurrect. Operation failures are terminal for the attempt; retry is a new linked attempt. Deletion and retained-volume policy need explicit tests. DB resize/migrate/fork states are reserved, not accepted in current wire schemas. Capability advertisement does not implement a driver; do not advertise planned ones. Run Python lifecycle checks and the shared Go/Java lifecycle vector tests before merging. |
