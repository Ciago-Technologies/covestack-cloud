# Covestack Cloud — Actual Repository Layout

Updated 11 October 2026. Baseline 349a2f4 is published on fix/security-baseline.

Product name: **Covestack Cloud**; intended repository slug: `covestack-cloud`.
The actual checkout and GitHub remote are still named `kan-cloud` until the repository
owner renames them. Use the real paths in commands. Existing module folders,
Go import paths, Java packages, signing headers/algorithm, KAN_* environment names,
API paths, labels, database identity and Valkey key prefixes remain compatibility
identifiers. A future technical rename needs a separate coordinated migration.

```text
kan-cloud/                         # Actual checkout until owner rename
├── contracts/                     # API schema, rules, three vectors, examples
├── kan-server/
│   ├── pom.xml                    # Boot 3.5.16 / Java 25 / PostgreSQL Flyway
│   └── src/{main,test}/java/com/ciagotech/kan/server/auth/
│       # Agent/Bearer filters, registration, credential SPIs, Valkey store/tests
├── kan-agent/
│   ├── go.mod                     # Go 1.27.2 toolchain pin
│   ├── go.sum
│   ├── sigv1/                     # Signer, request helper and regression tests
│   ├── cmd/auth-smoke/             # Runnable integration smoke command
│   └── internal/config/config.go  # Placeholder: package config only
├── deploy/
│   ├── docker-compose.yml         # Postgres, Valkey, Mailpit
│   ├── .env.example
│   └── valkey.conf.example
├── .github/workflows/             # Go/Java/contracts CI and secrets scan
├── docs/ANUJ-PERSON-C-HANDOFF.md
├── agent/                         # Empty README placeholder
├── control-plane/                 # Empty README placeholder
└── README.md
```

## Next additions

Anuj implements config loading, typed API client, cmd/kan-agent, polling/reconcile
and Docker drivers. C implements KeyService/BearerTokenService persistence and
permissions. B adds app/main/profiles/migrations/scoped controllers/report ordering.
D adds provider DTOs/endpoints via B services; E adds management/bootstrap/CLI.
No directory/module migration is required for the branding change.

Shared contracts already sign nonce/key ID, require observedAt, allow nullable
health generations and use running for observed K8s state. Future agent lease,
report sequencing and kind-specific schemas still require coordinated B/A/C work.

Go and Python tests pass locally on installed Go 1.27.1; exact pinned Go/Maven/VM
acceptance remains pending. CI definitions exist but actual remote runs are not
verified here. MinIO, workload isolation and release packaging remain future work.

Read [design](Covestack-Cloud-Technical-Design.md),
[roles](Covestack-Cloud-Team-Roles.md),
[setup](Covestack-Cloud-Complete-Working-Setup.md),
[workflow](Covestack-Cloud-Project-Workflow.md),
[remaining fixes](Covestack-Cloud-Current-Setup-Fixes.md).
