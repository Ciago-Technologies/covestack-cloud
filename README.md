# CoveStack Cloud

| Item | Details                                                                                                                                                                                                                                                                                                    |
| ---- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Local DigitalOcean emulator under construction. Currently this repository contains shared contracts and Go/Java authentication libraries, not a runnable server/agent.                                                                                                                                     |
| 2    | Product name: **CoveStack Cloud**; repository slug: `covestack-cloud`. Existing kan-server/kan-agent folders, KAN_* configuration and signing identifiers remain compatible. The local checkout and origin URL now use `covestack-cloud`. See [team start guide](V5/CoveStack-Team-START-HERE.md). |
| 3    | Baseline 349a2f4 is published on fix/security-baseline. Agent configuration loader/tests exist; the signed HTTP client and polling executable remain pending; credential/token persistence, app/controllers and full agent remain pending. Review the baseline PR before claiming runtime readiness.   |

## Current documentation

Use [V5](V5/README.md) for the current CORE-5 scope, team roles and schedule to **Monday, 25 January 2027**. Historical team guides and the older integration handoff are preserved as ZIPs in [docs/archive](docs/archive/README.md); their commands, role assignments and December schedule are not current instructions. Documentation describes targets, not service acceptance evidence.

## Verify the current baseline

| Item | Details                                                                                          |
| ---- | ------------------------------------------------------------------------------------------------ |
| 1    | Use Java 25, Maven 3.9+, Go 1.27.2 (pinned by toolchain and CI), Python 3. From repository root: |

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
(cd kan-agent && go test -race ./... && go vet ./...)
(cd kan-server && mvn --batch-mode verify)
```

## Local infrastructure

| Item | Details                                                                                                                                                                                                                                                                                                                                 |
| ---- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Create `deploy/.env` from `.env.example`; select exact tested image references. Create `deploy/secrets/postgres_password` containing a random password and copy `deploy/valkey.conf.example` to `deploy/secrets/valkey.conf`, replacing its password. Protect local secrets with mode 600. These files are gitignored. Start from root: |

```bash
docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/.env -f deploy/docker-compose.yml up -d
```

| Item | Details                                                                                                                                                                                                                                                                                                                                                                                                   |
| ---- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Compose provides Postgres, authenticated persistent Valkey and Mailpit. MinIO is not included in this Compose configuration. Existing Postgres volumes retain their original password; changing the secret file does not rotate an initialized database's credentials. Do not delete data volumes to rotate credentials. Coordinate database password rotation with the Java owner and update connection settings together. |

## Integration

| Item | Details                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Read `contracts/CANONICAL-REQUEST.md` and [V5 agent work](V5/CoveStack-Anuj-Next-Steps.md) and [security integration](V5/CoveStack-Security-and-ITISS.md). Go's `sigv1.NewSignedRequest` signs exact body bytes, raw path/query and fresh nonces. Use HTTPS outside localhost and a client with redirects disabled. Person C implements `KeyService.findActiveCredential`, with service keys scoped to an account, and the Bearer token lookup implementation. Public filters are registered but require these implementations and account-scoped controllers. Agent auth configuration uses Valkey; there is no production fallback to the in-memory store. |
| 2    | Person B still needs application main/configuration, migrations and account-scoped controllers. Boot repackage is disabled while this is a library; enable it when the application entry point lands. Anuj still needs the agent executable, polling and Docker operations. VM/Docker deployment and real Valkey integration are separate acceptance checks; unit tests do not establish those are working.                                                                                                                                                                  |

## Complete lifecycle contract

| Item | Details                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| ---- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Version1.1.0 is defined in [lifecycle-rules.json](contracts/lifecycle-rules.json). [State reference](contracts/FROZEN-STATES.md) covers all seven resource kinds and server operations/notifications/metadata policies. Five agent kinds are capability- filtered; compute.group and storage.bucket are server-only. New observed reports require kind, and heartbeat requires supportedKinds. Planned drivers are not operational by declaration. Go/Java lifecycle helpers and shared vectors exist; B/Anuj must call them from the future scoped services/client dispatch. Regenerate with `python3 contracts/generate_lifecycle.py`; CI runs generation/schema fixture checks. |
