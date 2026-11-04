# Historical environment evidence — scope clarification

> Original version/check results below remain unchanged and dated; they are not current supported host profiles or VM runtime acceptance. Existing sandbox restrictions remain in force. VM/KVM/Incus capability was not evaluated by those records. See [implementation evidence](platform/implementation-status.md) and [capacity qualification](platform/capacity.md).

<!-- Historical body preserved below; see supersession notice above. -->

# CoveStack Development Environment

Updated: **4 November 2026**

This document records the development environment inside the OmniRush shell sandbox, not the Fedora host. It is evidence of the checks actually run; installed tools do not imply that a daemon or production service is available.

## Environment

- Working directory: `/run/media/anujtiwari/DevMount/Project-CDAC/covestack-cloud`
- OS inside sandbox: Debian GNU/Linux 12
- Kernel/architecture: Linux 7.2.11-200.fc44.x86_64, x86_64
- Shell identity: root inside the sandbox; no `sudo` was used
- Disk at audit: project mount 46 GB free; overlay filesystem 164 GB free
- Network: outbound HTTPS succeeded against go.dev, pypi.org and registry.npmjs.org
- Package managers: apt/apt-get, dnf, apk, pacman, pip/pip3, uv, npm/npx and cargo are exposed through OmniRush wrappers
- PATH: OmniRush wrappers first, then `/usr/local/bin` and system paths

## Installed tool locations

- Go 1.27.3: `/usr/local/go`, with `go` and `gofmt` in `/usr/local/bin`
- Temurin JDK 25.0.5: `/usr/lib/jvm/temurin-25`; `java` and `javac` are in `/usr/local/bin`
- Maven 3.8.7: `/usr/bin/mvn`
- PostgreSQL 17.11: `/usr/lib/postgresql/17`; client tools are on PATH
- Valkey 9.1.3: `/usr/local/bin/valkey-server` and `/usr/local/bin/valkey-cli`
- Docker CLI 20.10.24: `/usr/bin/docker`
- Docker Compose v5.7.0: `/usr/local/lib/docker/cli-plugins/docker-compose`
- kubectl v1.37.2: `/usr/local/bin/kubectl`
- k3d v5.9.0: `/usr/local/bin/k3d`
- MinIO client mc RELEASE.2025-08-13T08-35-41Z: `/usr/local/bin/mc`
- Ansible core 2.14.18 / ansible package 7.7.0: system Python packages
- ansible-lint 6.13.1 and ansible-vault: system PATH
- golangci-lint 2.15.0: `/usr/local/bin/golangci-lint`
- Git 2.39.5, GitHub CLI 2.23.0 and curl 7.88.1: system PATH

## Project checks

Run from the repository root unless noted:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
(cd kan-agent && gofmt -w internal/agent/loop.go internal/agent/loop_test.go cmd/kan-agent/main.go)
(cd kan-agent && gofmt -l .)
(cd kan-agent && go test -race ./...)
(cd kan-agent && go vet ./...)
(cd kan-agent && golangci-lint run)
(cd kan-server && mvn --batch-mode --no-transfer-progress test)
(cd kan-server && mvn --batch-mode --no-transfer-progress verify)
python3 V5/validate_v5.py
docker compose -f deploy/docker-compose.yml config --quiet
```

Observed results on 4 November 2026:

- Contract signing and lifecycle checks passed.
- V5 validation passed: 31 Markdown files, 336 tables, 72,849 words and all 38 baseline audit IDs.
- Go race tests passed for all agent packages, including the new agent package.
- `gofmt -w internal/agent/loop.go internal/agent/loop_test.go cmd/kan-agent/main.go` completed successfully. `gofmt -l .` still lists 13 pre-existing tracked files with CRLF line endings; those unrelated line-ending changes were preserved.
- `go vet ./...` produced no diagnostic.
- Maven test and verify passed; 15 Java tests passed with zero failures/errors.
- Compose configuration validation passed.
- `golangci-lint run` passes with `0 issues` after the eight code-level corrections documented in `docs/development-toolchain-status.md`.

## Go lint remediation

The original eight findings were inspected and fixed without disabling linters or changing contracts:

| Original finding | Correction |
| --- | --- |
| `cmd/auth-smoke/main.go:42:22` — unchecked `response.Body.Close` | Check the close error and return a contextual error from the smoke-check helper. |
| `internal/client/client.go:113:23` — unchecked deferred `resp.Body.Close` | Use a named result and a checked deferred close; surface a close failure when no earlier request/parse error exists. |
| `internal/client/client_test.go:54:13` — unchecked `fmt.Fprint` | Check the test response write and report writer errors through `testing.T`. |
| `internal/client/client_test.go:87:107` — unchecked `fmt.Fprint` | Expand the handler and check the response write error. |
| `internal/client/client_test.go:112:15` — unchecked `fmt.Fprint` | Check the response write error, while correctly omitting a body for HTTP 204. |
| `internal/agent/loop.go:136:65` — deprecated `net.Error.Temporary` | Remove the deprecated, ill-defined probe and retry portable network timeout errors only. |
| `internal/client/client.go:82:28` — De Morgan simplification | Rewrite the equivalent HTTPS/local HTTP condition in simplified form. |
| `internal/config/config.go:68:28` — De Morgan simplification | Apply the same equivalent condition simplification to configuration validation. |

Final Go results: `go test -race ./...` passed for all packages, `go vet ./...` produced no diagnostic, and `golangci-lint run` reported `0 issues`.

## Local services

### PostgreSQL

PostgreSQL 17.11 was installed from the official PostgreSQL APT repository. The sandbox cluster is online on port **5433** (the package selected this port), and this read-only probe passed:

```bash
pg_isready -h 127.0.0.1 -p 5433
runuser -u postgres -- psql -p 5433 -d postgres -tAc 'select 1;'
```

The expected V5 Compose port 5432 was not serving in this sandbox. No production database was contacted.

### Valkey

Valkey 9.1.3 was built from the official valkey-io/valkey repository at tag 9.1.3. A disposable local instance on 127.0.0.1:6379 passed `PING` and a `SET/GET/DEL` probe. Its data directory is outside the repository under `/tmp/opencode/valkey`.

### Docker and Compose

Docker CLI and Compose are available, and the Compose file parses successfully. The Docker daemon is not reachable:

```text
Cannot connect to the Docker daemon at unix:///var/run/docker.sock
```

Therefore Docker Compose startup/shutdown, k3d cluster creation and MinIO container testing are blocked. No socket or host service configuration was changed.

### Kubernetes

kubectl and k3d are installed. k3d reports its bundled default k3s version as v1.35.5-k3s1, but no cluster can be listed because Docker is unavailable. A host-level k3s binary was not installed; k3d supplies the required k3s workload image.

No Kubernetes manifest containing an `apiVersion:` field was found in the repository during this audit; the YAML inventory contains the OpenAPI contract and Docker Compose configuration.

### Object storage

The MinIO client `mc` is installed. No MinIO/S3 endpoint is running in the sandbox, so no object operation was attempted.

## Accepted permanent setup limitations

Per the environment decision, these remain intentionally unresolved and must not trigger further setup work:

- Docker daemon/socket access is not provided; Compose runtime, k3d clusters, MinIO containers and container-backed acceptance remain unavailable.
- The repository has no Kubernetes manifests.
- No MinIO/S3 endpoint is configured.
- Terraform and `terraform-plugin-framework` remain deferred.
- GitHub CLI remains unauthenticated.

## Credentials and host limits

No credentials, API keys, SSH keys, kubeconfigs or production endpoints were read or recorded. GitHub CLI is installed but `gh auth status` reports no authenticated host. Terraform is optional in V5 and was not installed before the CLI/provider work begins.
