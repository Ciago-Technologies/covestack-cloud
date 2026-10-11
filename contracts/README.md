# Covestack Cloud — contracts/

This folder is the agreement between kan-server (Java) and kan-agent (Go).
Nobody writes server<->agent connection code before this is frozen.

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

1. Run the three computed vectors in Python, Java and Go. Review the mandatory
   signed nonce/access-key revision with Person C before integrating.
2. Anuj + the Java lead both review this whole folder as one pull request.
3. Merge = frozen. Target: Day 5.
4. After that, changing anything in `contracts/` requires a PR reviewed by
   both of them again. No solo edits.

## Why k8s.cluster was picked for every example

Its desired states are `active`/`deleted`; observed states include `provisioning`,
`running`, `degraded`, `error`, `deleting`, `deleted`. Do not confuse desired and
observed state. The same shapes
apply to `compute.instance` and `db.instance` — just swap the `kind`,
`spec`, and use that kind's row from FROZEN-STATES.md.
