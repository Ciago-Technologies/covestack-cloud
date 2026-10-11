# CoveStack Cloud — Corrected Setup Guide

| Item | Details |
|---|---|
| 1 | Updated 11 October 2026. Read [current setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) before using this guide. The committed security/build baseline is present. The extended infrastructure examples below are proposals; the actual Compose file provides only Postgres/Valkey/Mailpit. Java and the full agent still need entry points and implementations. |
| 2 | The security/build baseline is committed as `349a2f4` (Fixed security issues) on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`. The lifecycle changes are currently applied but uncommitted. Remote main/PR merge status has not been verified; review and merge through the shared PR. Three signature vectors, negative tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates and auth-smoke exist. The config.go file currently contains only `package config`; it is a placeholder, not a working loader. Real credential/token lookup, application main, migrations/controllers, full agent and live deployment tests remain pending. |
| 3 | Product name: **CoveStack Cloud**; repository slug: `covestack-cloud`. The local checkout and origin URL now use `covestack-cloud`. Use that repository folder in commands. Existing module folders, Go import paths, Java packages, signing headers/algorithm, KAN_* environment names, API paths, labels, database identity and Valkey key prefixes remain compatibility identifiers. A future technical rename needs a separate coordinated migration. |

## 1. Keep the existing project layout

| Item | Details |
|---|---|
| 1 | Use `kan-server/` for Java and `kan-agent/` for Go. Keep contracts beside them and Compose in `deploy/`. Add a standalone `kan-cli/` later. Do not move everything into control-plane/agent during deadline work. |
| 2 | From the current Claude directory: |

```bash
cd ../Project-CDAC/covestack-cloud
```

| Item | Details |
|---|---|
| 1 | Read AGENTS/CONTRIBUTING instructions if present. Preserve uncommitted edits and use reviewed branches. Security/build changes are now committed on fix/security-baseline. This branding update adds documentation and display-name changes; live VM deployment remains unverified. |

## 2. Verify the machine before installing again

| Item | Details |
|---|---|
| 1 | Anuj uses Fedora for editing and the AlmaLinux VM for integration. Windows/Mac teammates run Java plus local Docker infrastructure and an explicit test FakeAgent. The VM needs enough headroom for the demonstrated workloads: start with 8 GiB RAM, 4 vCPUs and at least 60 GiB disk if the host supports it; measure docker stats and disk usage rather than assuming that allocation guarantees capacity. |

```bash
java -version
javac -version
mvn -version
go version
docker version
docker compose version
python3 --version
```

| Item | Details |
|---|---|
| 1 | Java and javac must be 25; Maven must use Java 25. Record exact tested versions. Upgrade the inspected Go 1.27.1 to the verified 1.27.2 security patch or a later reviewed supported patch. Use checksums from the official release, not an unverified download or old Go 1.22 install recipe. Follow [tools](Covestack-Cloud-Complete-Working-Setup.md) for platform-specific setup. |
| 2 | Keep SELinux and firewalld enabled. Fix narrowly scoped permissions/policies instead of disabling them. Docker-group membership grants root-level daemon authority; only the trusted agent/developer account should have it. |

## 3. Prepare secrets and versions

| Item | Details |
|---|---|
| 1 | Create gitignored `deploy/local/` with directory mode 0700 and secret files mode 0600. Generate separate random Postgres, Valkey, MinIO and AES-GCM master secrets. Do not print them into logs, command history or shared screenshots. Add restrictive Windows ACLs for developer equivalents. |
| 2 | Record reviewed image references in `deploy/versions.env` as complete `name@sha256:digest` values for Postgres 17, Valkey 8, Mailpit and MinIO. These major lines match the existing deployment; verify the exact patches and feature support before pinning. Use one tested matrix, not mixed 16/17 and 7/8 examples. This file is not automatically loaded by Compose; pass `--env-file deploy/versions.env` explicitly. |
| 3 | Required local files: |

```text
local/postgres-password.txt
local/valkey.conf
local/minio-user.txt
local/minio-password.txt
local/master-key.bin
```

| Item | Details |
|---|---|
| 1 | Valkey config must include `appendonly yes`, a reviewed `appendfsync` durability setting, `maxmemory 128mb`, `maxmemory-policy noeviction`, and a non-default `requirepass`. Configure protected mode/listening for the controlled backend network and authenticate clients. Test nonce durability across graceful restart and abrupt loss; fail closed after replay-history loss. |
| 2 | A new password file does not automatically rotate the password in an existing Postgres volume. Rotate the database role and update client configuration together; preserve/back up existing data first. Apply the same coordinated rotation principle to cache/S3 credentials. |

## 4. Proposed infrastructure Compose replacement

| Item | Details |
|---|---|
| 1 | Put this in `deploy/docker-compose.yml` only after creating the required files and resolving image variables. It starts infrastructure, not a server/agent that has not yet been implemented. |

```yaml
services:
  postgres:
    image: ${POSTGRES_IMAGE:?Set a reviewed image digest}
    environment:
      POSTGRES_DB: kan
      POSTGRES_USER: kan
      POSTGRES_PASSWORD_FILE: /run/secrets/postgres_password
    secrets: [postgres_password]
    ports: ["127.0.0.1:5432:5432"]
    volumes: [pgdata:/var/lib/postgresql/data]
    networks: [kan-backend]
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U kan -d kan"]
      interval: 5s
      timeout: 3s
      retries: 20

  valkey:
    image: ${VALKEY_IMAGE:?Set a reviewed image digest}
    command: ["valkey-server", "/run/secrets/valkey_config"]
    secrets: [valkey_config]
    ports: ["127.0.0.1:6390:6379"]
    volumes: [valkeydata:/data]
    working_dir: /data
    networks: [kan-backend]

  mailpit:
    image: ${MAILPIT_IMAGE:?Set a reviewed image digest}
    environment:
      MP_MAX_MESSAGES: "500"
    ports: ["127.0.0.1:1025:1025", "127.0.0.1:8025:8025"]
    networks: [kan-backend]

  minio:
    image: ${MINIO_IMAGE:?Set a reviewed image digest}
    command: ["server", "/data", "--console-address", ":9001"]
    environment:
      MINIO_ROOT_USER_FILE: /run/secrets/minio_user
      MINIO_ROOT_PASSWORD_FILE: /run/secrets/minio_password
    secrets: [minio_user, minio_password]
    ports: ["127.0.0.1:9000:9000", "127.0.0.1:9001:9001"]
    volumes: [miniodata:/data]
    networks: [kan-storage]

secrets:
  postgres_password:
    file: ./local/postgres-password.txt
  valkey_config:
    file: ./local/valkey.conf
  minio_user:
    file: ./local/minio-user.txt
  minio_password:
    file: ./local/minio-password.txt

volumes:
  pgdata:
  valkeydata:
  miniodata:

networks:
  kan-backend:
    driver: bridge
  kan-storage:
    driver: bridge
```

| Item | Details |
|---|---|
| 1 | Local Compose secrets are file mounts, not an encrypted secret vault. Verify the selected image actually supports these `_FILE` settings and can read the mounted files under its runtime UID. Secret file permissions/SELinux labels and volume ownership must be tested on the VM. The backend bridge is not by itself an egress firewall. |
| 2 | From repository root, after editing the repository files: |

```bash
docker compose --env-file deploy/versions.env -f deploy/docker-compose.yml config --quiet
docker compose --env-file deploy/versions.env -f deploy/docker-compose.yml up -d
docker compose --env-file deploy/versions.env -f deploy/docker-compose.yml ps
```

| Item | Details |
|---|---|
| 1 | Use `down` without `-v` for ordinary restarts. Never use blanket pruning as a routine fix: snapshot images and stopped-workload data may be valuable. |

## 5. Java server changes before startup

| Item | Details |
|---|---|
| 1 | In `kan-server/pom.xml`, retain Java 25, replace Boot 3.3.5 with a reviewed maintained version explicitly supporting Java 25, and add the aligned Flyway PostgreSQL module. Add mail support before using JavaMailSender. Keep ordinary signing unit tests independent of a live database. |
| 2 | Create the Spring application entry point, configuration, Flyway migrations, credential implementations, registered filters and controllers. Set server port 4700; while running Java directly on the VM host, bind to 127.0.0.1. Connect to the loopback infrastructure ports with secret-file configuration. |
| 3 | If later packaging Java in Compose, build context is `../kan-server` relative to `deploy/docker-compose.yml`; create an actual Dockerfile first. Use service hostnames inside the container, join only backend/storage as required, publish API on guest loopback and never mount the Docker socket. Prefer an infrastructure profile and an app profile so early development does not try to build a missing Dockerfile. |
| 4 | Run only once these files exist: |

```bash
cd kan-server
mvn -B verify
mvn spring-boot:run
```

| Item | Details |
|---|---|
| 1 | Use the Maven wrapper after the team adds it. Packaging must find the real entry point. A cache/network dependency failure must be fixed as build infrastructure, not declared a successful test. |

## 6. Fix and wire authentication before API demos

| Item | Details |
|---|---|
| 1 | Follow [design section 4](Covestack-Cloud-Technical-Design.md#4-authentication-and-replay-protection): signed nonce/key ID, required headers, bounded body, active service credentials, constant-time verify, atomic nonce claim after verification and trusted principal. |
| 2 | Implement Bearer auth for both `/v2/` and `/kan/v1/`, with only minimal liveness explicitly public. Bootstrap the first account/token through a local command and restrict its credential file. Reject unsigned internal calls and invalid/revoked tokens. No dev-skip header or arbitrary fake Bearer token is accepted in integration mode. |
| 3 | Confirm the real filter is registered. Test oversized chunked requests, mutated nonce, identical replay, bad-signature nonce poisoning, missing principal, store outage, cross-account IDs and proxy Host handling. Passing old vectors alone is insufficient. |

## 7. Go agent and driver setup

| Item | Details |
|---|---|
| 1 | Create `kan-agent/cmd/kan-agent/main.go`, configuration loader, signed API client, leased/fenced polling loop and Docker drivers. Add a reviewed Docker SDK dependency compatible with the daemon API; pin it and commit go.sum. Run go mod tidy after imports exist; tidy removes unused dependencies. |
| 2 | Test command available today: |

```bash
cd kan-agent
go test ./...
```

| Item | Details |
|---|---|
| 1 | Future command once entry point/config exists: |

```bash
go run ./cmd/kan-agent --config /path/to/restricted/agent-config.json
```

| Item | Details |
|---|---|
| 1 | Do not use `go run .` against the current package-less root. Agent config supplies private API URL, signing credential file, account/agent assignment, port ranges, trusted images and quotas. Rebuild ports before workers start. Inspect healthy containers periodically; process dirty events on their own shard. |
| 2 | User workloads get account networks and ownership labels; user databases do not join kan-backend. Only authorized agents can return observations. Failures do not advance successful observed generation. Reboot carries explicit operation intent. |

## 8. VM routing, DNS and TLS

| Item | Details |
|---|---|
| 1 | Reserve the AlmaLinux NAT interface IP. Record actual laptop hotspot IP, guest IP/interface and demo subnet; example addresses are not universal. Forward laptop HTTPS 443 to guest 443, DNS TCP and UDP 53 to guest 53, and optionally laptop 2222 to guest 22 for controlled administration. |
| 2 | Do not forward control-plane Postgres or Valkey to the hotspot. Do not forward raw API 4700; the public API goes through HTTPS. Use SSH tunnels for private administrative interfaces when possible. |
| 3 | Example dnsmasq configuration, with the example IPs/interface replaced: |

```conf
port=53
listen-address=127.0.0.1,192.168.100.10
bind-interfaces
no-dhcp-interface=ens33
address=/anuj-kan.duckdns.org/192.168.43.100
server=1.1.1.1
```

| Item | Details |
|---|---|
| 1 | The guest interface listener receives NAT-forwarded traffic; the DNS answer is the laptop address reachable by clients. Restrict resolver access to the intended NAT/demo traffic in the firewall; UDP-only forwarding is incomplete. For the simpler fallback, set a hosts entry on the two controlled clients and keep the same certificate hostname. |
| 2 | Provision TLS using DNS-01 before the demo. Use a restricted provider credential file, a tested renewal/reload mechanism and fixed domain. mkcert is the fallback on owned clients; only its public CA certificate is distributed. Do not assume a real certificate lasts forever. |
| 3 | Public Nginx API location example, after setting the actual domain/certificate paths: |

```nginx
server {
    listen 443 ssl;
    server_name anuj-kan.duckdns.org;
    ssl_certificate /etc/letsencrypt/live/anuj-kan.duckdns.org/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/anuj-kan.duckdns.org/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    client_max_body_size 1m;
    location = /internal { return 404; }
    location ^~ /internal/ { return 404; }
    location / {
        proxy_pass http://127.0.0.1:4700;
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Request-Id $request_id;
    }
}
```

| Item | Details |
|---|---|
| 1 | Add a default virtual host rejecting unrecognized hosts and configure application host allowlists. Keep agent traffic on the private signed loopback route. On SELinux systems permit only the required Nginx backend connection under a reviewed policy; diagnose AVC denials rather than disabling SELinux. |
| 2 | Mailpit HTTP :8025 is not HTTPS. Use an SSH tunnel for its UI, or add an explicit restricted TLS listener with an HTTP upstream. Use a separate suitable S3 TLS hostname/proxy for direct scoped S3 access; don't accidentally alter SigV4 path/Host semantics. Keep MinIO admin console private. |

## 9. Workload connectivity and persistence

| Item | Details |
|---|---|
| 1 | Configure a finite tested set of workload port forwards on VMware for SSH, user DB, HTTP LB and Kubernetes API. Do not assume a GUI supports a port-range rule. If necessary, maintain individual rules for the demo's small fixed port budget or a documented host forwarding mechanism. |
| 2 | Bind Docker workload ports to the intended guest interface and apply workload/host firewall restrictions. Return the external hostname plus forwarded port in client metadata. `localhost` is correct only for clients running on the same machine. |
| 3 | Build key-only SSH images and test actual login under the selected capabilities. Create a labeled data volume for each user DB. Test that a recreated DB container reads the same table. For Spaces, test that objects survive MinIO container recreation, and account A cannot access account B's physical bucket with scoped credentials. |
| 4 | Test all components by IP as well as DNS from a real second device. An HTTP-ready web workload is required before demonstrating a load balancer. |

## 10. Acceptance checklist

| Item | Details |
|---|---|
| 1 | [ ] Pinned Java/Boot/Go matrix builds on all developer machines. |
| 2 | [ ] Infrastructure loads from repository root; secret files excluded from Git. |
| 3 | [ ] Real server and agent entry points run; filters/principal implementations are wired. |
| 4 | [ ] Auth failure, replay/body limits and cross-account tests pass. |
| 5 | [ ] Create/power-off/power-on/reboot/delete reach truthful completed/failed operation outcomes. |
| 6 | [ ] Slow create and agent outage do not trigger unlimited ASG replacements. |
| 7 | [ ] DB/object data survive recreation; restore is rehearsed. |
| 8 | [ ] Public proxy blocks internal paths; backend is inaccessible from workloads/other devices. |
| 9 | [ ] DNS/TLS/workload ports work on the second OS without disabling security controls. |
| 10 | [ ] CI, clean-machine installation and the dated release gates pass. |
| 11 | The [current setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) assigns all 38 original findings to concrete files, owners and acceptance checks. The [workflow](Covestack-Cloud-Project-Workflow.md) schedules them through December 15. |
