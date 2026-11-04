# Historical environment evidence — scope clarification

> Original version/check results below remain unchanged and dated; they are not current supported host profiles or VM runtime acceptance. Existing sandbox restrictions remain in force. VM/KVM/Incus capability was not evaluated by those records. See [implementation evidence](platform/implementation-status.md) and [capacity qualification](platform/capacity.md).

<!-- Historical body preserved below; see supersession notice above. -->

# CoveStack Development Toolchain Status

Updated: **4 November 2026**

Status is based on executable checks inside the OmniRush sandbox. “Available” means the executable ran; “Tested” describes the specific capability check, not merely installation.

| Tool | Version | Available inside sandbox | Tested | Result |
| --- | --- | --- | --- | --- |
| Go | 1.27.3 | Yes | `go test -race ./...`, `go vet ./...` | Tests passed; vet produced no diagnostic |
| gofmt | Go 1.27.3 | Yes | `gofmt -w` requested files; `gofmt -l .` | Requested files formatted; 13 pre-existing CRLF files remain listed |
| golangci-lint | 2.15.0 | Yes | `golangci-lint run` | PASS: 0 issues after eight code-level fixes |
| Java/JDK | Temurin 25.0.5 | Yes | `java -version`, `javac -version` | Passed |
| Maven | 3.8.7 | Yes | `mvn test`, `mvn verify` in kan-server | 15 tests passed; build success |
| Python | 3.11.2 | Yes | AST parse and `PYTHONPYCACHEPREFIX=... python3 -m compileall -q .` | Passed for 5 Python files |
| pip / venv | pip 23.0.1 / venv | Yes | Created and removed disposable venv outside project | Passed |
| Docker CLI | 20.10.24 | Yes | `docker --version` | Client works |
| Docker daemon | — | No | `docker info` | Permanently accepted sandbox limitation: `/var/run/docker.sock` unavailable |
| Docker Compose | 5.7.0 | Yes | `docker compose config --quiet` | Config validation passed; runtime permanently unavailable in this sandbox |
| kubectl | v1.37.2 | Yes | `kubectl version --client` | Client works; no cluster configured |
| k3d | v5.9.0 | Yes | `k3d version`, `k3d cluster list` | Binary works; cluster runtime permanently unavailable; no Kubernetes manifests found |
| k3s | v1.35.5-k3s1 bundled by k3d | No host binary | k3d version only | No separate install needed yet |
| PostgreSQL client/server | 17.11 | Yes | `pg_isready`, local `select 1` | Passed on sandbox port 5433 |
| Valkey server/CLI | 9.1.3 | Yes | `PING`, disposable SET/GET/DEL | Passed on 127.0.0.1:6379 |
| MinIO client | RELEASE.2025-08-13T08-35-41Z | Yes | `mc --version`, help | Client works; no S3 endpoint is intentionally configured |
| Ansible | core 2.14.18 / package 7.7.0 | Yes | `ansible --version` | Passed |
| ansible-lint | 6.13.1 | Yes | `ansible-lint --version` | Passed; no playbook run requested |
| ansible-vault | core 2.14.18 | Yes | `ansible-vault --version` | Passed |
| Git | 2.39.5 | Yes | `git --version` | Passed |
| GitHub CLI | 2.23.0 | Yes | `gh auth status` | Installed; authentication intentionally left unset |
| curl | 7.88.1 | Yes | HTTPS probes to go.dev, pypi.org, npm registry | Network returned HTTP 200 |
| Terraform | — | No | `terraform version` | Permanently deferred for this setup; optional in V5 |
| terraform-plugin-framework | — | No | Not applicable | Permanently deferred with Terraform provider |

## Project verification summary

| Check | Result | Evidence scope |
| --- | --- | --- |
| `python3 contracts/sign_example.py` | PASS | Signing vectors/examples |
| `python3 contracts/check_lifecycle.py` | PASS | Generated lifecycle files and fixture coverage |
| `python3 V5/validate_v5.py` | PASS | 31 Markdown files, 336 tables, 38 audit IDs |
| `(cd kan-agent && gofmt -l .)` | PRE-EXISTING FORMAT NOISE | 13 tracked CRLF files listed; unrelated line endings preserved |
| `(cd kan-agent && go test -race ./...)` | PASS | Agent unit tests in sandbox |
| `(cd kan-agent && go vet ./...)` | PASS | No vet diagnostic observed |
| `(cd kan-agent && golangci-lint run)` | PASS | 0 issues; no suppressions or configuration changes |
| `(cd kan-server && mvn test)` | PASS | 15 Java tests |
| `(cd kan-server && mvn verify)` | PASS | Java compile, tests, jar/repackage |
| `docker compose ... config --quiet` | PASS | Compose syntax/interpolation only |
| Docker Compose startup/shutdown | ACCEPTED LIMITATION | No Docker daemon/socket in this sandbox |
| k3d/kubectl cluster validation | ACCEPTED LIMITATION | No Docker daemon/cluster in this sandbox |
| MinIO/S3 object operations | ACCEPTED LIMITATION | No MinIO/S3 endpoint is configured |

## Go lint remediation record

The eight original findings were corrected as follows:

| Original location and finding | Code-level correction |
| --- | --- |
| `cmd/auth-smoke/main.go:42:22` `errcheck`: `response.Body.Close` | Check the close error and return a contextual error. |
| `internal/client/client.go:113:23` `errcheck`: deferred `resp.Body.Close` | Check the deferred close and surface it when no earlier error exists. |
| `internal/client/client_test.go:54:13` `errcheck`: `fmt.Fprint` | Check the write result and report errors through `testing.T`. |
| `internal/client/client_test.go:87:107` `errcheck`: `fmt.Fprint` | Expand the handler and check the write result. |
| `internal/client/client_test.go:112:15` `errcheck`: `fmt.Fprint` | Check the write result and omit the invalid body for HTTP 204. |
| `internal/agent/loop.go:136:65` `staticcheck SA1019`: `net.Error.Temporary` | Use the supported timeout signal; the deprecated method is ill-defined. |
| `internal/client/client.go:82:28` `staticcheck QF1001` | Apply the equivalent De Morgan simplification. |
| `internal/config/config.go:68:28` `staticcheck QF1001` | Apply the equivalent De Morgan simplification. |

Verification after remediation:

- `gofmt -w internal/agent/loop.go internal/agent/loop_test.go cmd/kan-agent/main.go` passed.
- `gofmt -l .` lists 13 pre-existing tracked CRLF files; no unrelated line-ending rewrite was performed.
- `go test -race ./...` passed.
- `go vet ./...` passed with no diagnostic.
- `golangci-lint run` passed with `0 issues`.

## Accepted permanent environment restrictions

The following are user-directed permanent decisions for this environment. They are not application defects, test passes, or service-gate evidence:

- Do not provide Docker daemon/socket access; Compose runtime, k3d clusters, MinIO containers and container-backed acceptance remain unavailable.
- Keep the repository without Kubernetes manifests.
- Do not configure a MinIO/S3 endpoint.
- Keep Terraform and `terraform-plugin-framework` deferred.
- Keep GitHub CLI unauthenticated.
