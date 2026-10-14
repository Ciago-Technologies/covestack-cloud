# CoveStack V5 — One Stack Integration Template

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. One integration template — CORE-5 cross-cutting demo

**Status: Planned.** ONE fixed-sequence template creates **PostgreSQL + MinIO + app Droplet** using the account's private network and owned credentials. This is test 15 integration, not a sixth service or a DAG engine. Broader template choices/customization remain preserved in the [roadmap](CoveStack-Extended-Roadmap.md#12-optional-future-work--not-additional-extended-services).

Each child remains a durable owned resource with its own gate evidence. The approved app must accept documented endpoint/secret references; invalid application assumptions produce explicit errors. No arbitrary uploaded shell/YAML/IaC job executes on the platform host.

## 2. End-to-end CORE-5 example

| Step | User action | Platform target |
|---|---|---|
| Select | CLI selects the ONE approved template | Read reviewed immutable definition/version |
| Configure | Validated app/image inputs | Check account quota and documented JSON input shape |
| Submit | One mutating command with idempotency key | Store owned stack/operation and child references |
| Provision | Inspect pending progress or --wait | Ensure private network, DB and storage dependencies; then app |
| Connect | Use scoped references | No passwords or signing secrets in spec/observed/list/get |
| Inspect | CLI list/get with -o json | Same child public IDs/generations/errors from API |
| Retry | Repeat durable request after failure | Resume without duplicate children |
| Remove | Explicit retained-data policy | Owned cleanup; no hidden deletion of persistent data |

CLI command/schema spelling is finalized by E+B; no new template ID, API version or wire field is invented. The old HCL/YAML/public-route examples are historical future designs in the roadmap, not runnable schemas.

## 3. Template record and JSON shapes

| Record | Required documented shape |
|---|---|
| stack_templates | Approved template identity and ownership/visibility |
| stack_template_versions.definition | Fixed ordered component list, approved input validation and references; schema_hash |
| stacks.inputs | Validated nonsecret values and protected secret references |
| stack_resources | Stack/child account ownership, role and ordinal; no duplicate child binding |

Definition and inputs are explicit jsonb exceptions; document their shapes before B writes migrations. No business ownership keys are hidden in JSON. Broader DAG/component customization is deferred, not silently removed.

## 4. Stack state and failure policy

Stack state values are **TBD with B**; do not invent them from human progress labels. Operations use queued, running, completed, failed, cancelled. Never rewrite a child state to make the stack appear converged.

| Failure | Required response |
|---|---|
| DB/storage dependency unavailable | Record failed attempt; do not start dependent app with fake endpoint |
| Agent/server restart | Resume durable child references; inspect before create |
| Duplicate request | Same idempotent result; changed-body same key → 422 |
| Partial deletion | Keep retryable owned cleanup records |
| App failure | Preserve user data under reviewed retention policy |

Record sanitized error/attempt/timestamps. Compensation is not unconditional data rollback; retention defaults remain B+A decision.

## 5. Single ownership model

CLI reads server-owned stack/resource records; no local registry. Optional CORE-5 Terraform manages only Droplet and Kubernetes resources, not stacks/children. Broader IaC ownership, adoption and customization requirements are preserved in the roadmap; clients never bypass API policy.

## 6. Acceptance

C owns test 15 with B/E and A execution pairing: one CLI command provisions PostgreSQL + MinIO + app Droplet; partial failure recorded; retry resumes without duplicate children. Also verify account denial, restart during provisioning, defined readiness and retained-data cleanup. Production CLI stack starts only after bucket failure evidence passes in CI.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
