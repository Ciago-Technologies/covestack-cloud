# CoveStack Cloud — contracts/

| Item | Details                                                                                                                                        |
| ---- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | This folder is the agreement between kan-server (Java) and kan-agent (Go). Nobody writes server<->agent connection code before this is frozen. |

```
contracts/
├── README.md                          <- you are here
├── internal-api.yaml                  <- the 5 endpoints, request/response shapes
├── FROZEN-STATES.md                   <- exact allowed state strings, do not respell
├── CANONICAL-REQUEST.md               <- kan sigv1 signing contract, mandatory headers and raw byte rules
├── sign_example.py                    <- run this to check your Java/Go signature output matches
├── sigv1-vectors.json                 <- shared signing test — Java and Go both check against it
└── examples/
    ├── get-work-response-k8s.json     <- example: GET /work response (k8s.cluster resource)
    ├── put-observed-request-k8s.json  <- example: PUT /observed body (k8s.cluster resource)
    └── post-event-request-k8s.json    <- example: POST /events body (k8s.cluster resource)
```

## Freeze process

| Item | Details                                                                     |
| ---- | --------------------------------------------------------------------------- |
| 1    | Run the three computed vectors in Python, Java and Go. Review the mandatory |
| 2    | signed nonce/access-key revision with Person C before integrating.          |
| 3    | Anuj + the Java lead both review this whole folder as one pull request.     |
| 4    | Merge = frozen. Target: Day 5.                                              |
| 5    | After that, changing anything in `contracts/` requires a PR reviewed by     |
| 6    | both of them again. No solo edits.                                          |

## Why k8s.cluster was picked for every example

| Item | Details                                                                                                                                                                                                                                                                                                                           |
| ---- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Its desired states are `active`/`deleted`; observed states include `provisioning`, `running`, `degraded`, `error`, `deleting`, `deleted`. Do not confuse desired and observed state. The same shapes apply to `compute.instance` and `db.instance` — just swap the `kind`, `spec`, and use that kind's row from FROZEN-STATES.md. |

## Lifecycle source and compatibility

| Item | Details                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ---- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | Version1.1.0 has one source: lifecycle-rules.json. generate_lifecycle.py produces FROZEN-STATES.md, internal-api.yaml and Go/Java catalog data. internal-api.base.json is the endpoint scaffold, not a second lifecycle-state table. check_lifecycle.py checks generator drift and our emitted schema subset against all-service/negative examples. It is not a general-purpose OpenAPI validator. Shared lifecycle-test-vectors are executed by both Go and Java. Full kind-specific spec DTO validation and live controller/driver ordering/authorization checks remain required. |
| 2    | Seven resource kinds are defined; six are agent-owned. Resource includes groups for server use, while AgentResource and ObservedReport exclude groups/buckets and buckets. Operation and notification enums are separate server lifecycles; metadata services are inventoried without inventing resource kinds. Heartbeat supportedKinds must list only real drivers; empty is valid. Observed kind is required and must match the stored resource. DB resizing/migrating/forking are reserved and excluded from accepted schemas.                                                  |
| 3    | Sample specs illustrate planned services, not authoritative driver-ready specs. Coordinate mandatory kind/capability fields and updated signature vectors in the same Anuj/B/C reviewed PR. Keep signing algorithm/header names unchanged for this revision.                                                                                                                                                                                                                                                                                                                        |
