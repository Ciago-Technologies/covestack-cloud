# Development and contribution guide

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. Documentation is open to review; the repository has no root license grant yet. Maintainers must resolve contribution/distribution terms in [license and attribution](license-and-attribution.md) before assuming rights.

## Start safely

Read root AGENTS.md, checkpoint.md, current git status, [status](implementation-status.md), [decisions](decisions.md) and relevant contract sources. Preserve unrelated local changes and never display secret config files. Current task permissions govern edits; documentation does not authorize infrastructure access or implementation.

Create a bounded issue with user outcome, owner, dependencies, public contract/fields/errors, persistent state/ownership, backend capability, trust class, idempotency/timeout/cancellation, failure/cleanup, actual acceptance and known exclusions. Source presence or a green health check is not a service. A new adapter must integrate its consumer, not just expose a standalone helper.

| Role | Retained responsibility |
| --- | --- |
| A/Anuj | Agent, signing/shared client, runtime drivers, reconcile, CI, installer/ITISS; types and explains agent Go |
| B | Java core, all SQL/migrations, leases/fences/idempotency, reservations/quotas/stacks |
| C | Security/auth and all acceptance tests; Java IPAM allocator |
| D | Java API/DTO/schema and compatibility fixtures; no Go |
| E | Thin CLI/JSON/packaging, later optional provider; not agent/drivers/signing |

Future Go scheduler/provider ownership is an explicit team decision, not a silent transfer of B's current Java coordination. Any Go needs A review; shared-client/VM backend boundaries A+B+C; schema B+C+A, DTO D; C security/test review.

## Review sequence

1. Read affected callers, tests, generated/base schema and docs; define small change scope.
2. Agree missing contracts and migration policy before code. New VM/type/claim interfaces require coordinated fixtures and generated outputs, not a solo wire string.
3. Implement in approved scope; no new dependency or service when an existing reviewed backend capability suffices.
4. Run relevant unit/static checks, then integration/real runtime evidence on separately authorized infrastructure; label each evidence class and retain failures.
5. Check clean install/upgrade/rollback/cleanup/security and document supported subset, release/profile limits and unresolved work.
6. Open PR, attach commands/actual results and reviewers; no direct main push or fabricated green CI.
7. Update checkpoint.md with concrete next action and source/evidence scope, not guessed completion percent. GitHub issues own live status; this guide owns process.

Do not run restricted Docker/Kubernetes/MinIO/Terraform/GitHub operations in the current sandbox to satisfy a gate. A unit fake cannot replace real VM/packet/data evidence. No application tests were run for this docs redesign.

## Documentation contribution rules

Current architecture belongs in docs/platform. Keep V5 historical body and dated environment/checkpoint/test evidence intact; add explicit supersession or correction, never rewrite historical test outcomes. If a source feature changes, update implementation ledger, API compatibility, feature spec, readiness and affected operating/upgrade steps together. A release must not become production-supported through wording alone. Cite official sources for technical/competitor claims and distinguish absence of evidence from evidence of absence.

No new code license, CLA, DCO, governance contact or security-reporting email is chosen by this guide; maintainers must decide and publish the actual policy.
