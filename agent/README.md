# CoveStack Cloud — Agent placeholder

| Item | Details                                                                                |
| ---- | -------------------------------------------------------------------------------------- |
| 1    | The active modules remain kan-agent/ and kan-server/. This directory is a placeholder. |

## Agent configuration

`kan-agent/internal/config.Load` reads process environment only; it does not source an environment file. The existing `auth-smoke` executable uses this loader. The full agent executable and HTTP client are still pending.

| Variable         | Requirement                                                                                                                                                                |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `KAN_SERVER_URL` | Required HTTP(S) origin. HTTPS is required outside localhost/loopback; URL credentials, paths (except `/`), queries and fragments are rejected.                            |
| `KAN_ACCESS_KEY` | Required nonempty service access key, without whitespace.                                                                                                                  |
| `KAN_SECRET_KEY` | Required nonempty signing secret, without whitespace. Values are preserved exactly.                                                                                        |
| `KAN_RULES_PATH` | Optional; defaults to `../contracts/sigv1-rules.json` when running from `kan-agent/`. Set explicitly from other working directories. The shared signer validates the file. |

After B/C provide a running server and active service credentials, load your trusted local file and run from the repository root:

```bash
set -a
source deploy/secrets/agent.env
set +a
(cd kan-agent && go run ./cmd/auth-smoke)
unset KAN_ACCESS_KEY KAN_SECRET_KEY KAN_SERVER_URL KAN_RULES_PATH
```

Configuration errors and normal Go formatting of the configuration redact credentials. Do not log individual credential fields or serialize the configuration. Unit tests use synthetic credentials; they do not establish live Java integration.
