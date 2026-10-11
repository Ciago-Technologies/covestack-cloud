# kan-cloud

Local free emulator for DigitalOcean — runs entirely on your laptop with Docker.

## Architecture

```
kan-server  (Java / Spring Boot)  — brains: REST API, reconcile loop, Postgres, Valkey
kan-agent   (Go)                  — hands: talks to Docker, creates containers/volumes/networks
contracts/                        — shared API spec + signing rules (frozen before Day 5)
```

**Flow:** User calls `kan-server` → server stores desired state in Postgres → agent polls `/internal/v1/agent/work` → agent provisions Docker resources → agent reports observed state back.

## Quickstart

### Prerequisites
- Java 21, Maven 3.9+
- Go 1.22+
- Docker Engine 24+
- `docker compose` v2

### Start infrastructure
```bash
docker compose up -d   # Postgres, Valkey, Mailpit, MinIO
```

### Run kan-server
```bash
cd kan-server
mvn spring-boot:run
```

### Run kan-agent
```bash
cd kan-agent
go run .
```

## Repo layout
```
kan-cloud/
├── contracts/          ← API contract between server and agent (READ BEFORE CODING)
│   ├── internal-api.yaml
│   ├── FROZEN-STATES.md
│   ├── CANONICAL-REQUEST.md
│   ├── sigv1-rules.json
│   ├── sigv1-vectors.json
│   ├── sign_example.py
│   └── examples/
├── kan-server/         ← Java / Spring Boot
│   ├── pom.xml
│   └── src/
├── kan-agent/          ← Go
│   ├── go.mod
│   └── sigv1/
└── docker-compose.yml  ← (to be added)
```

## Signing (kan sigv1)

All agent↔server traffic is signed. Before writing any HTTP client or server code, read `contracts/CANONICAL-REQUEST.md` and run the test vectors:

```bash
# Java
cd kan-server && mvn test

# Go
cd kan-agent && go test ./sigv1/...

# Python sanity check
python3 contracts/sign_example.py
```

Both tests read from `contracts/sigv1-rules.json` and `contracts/sigv1-vectors.json` and must produce identical signatures.
