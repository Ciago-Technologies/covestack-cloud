# Product strategy and user validation

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. Status: **Proposed product hypotheses; VM target selected, implementation Planned.**

## Positioning

A local-first infrastructure control plane for small teams that want account-owned environments, intelligent resource scheduling, reliable provisioning, built-in security and recoverable infrastructure across their own machines. CoveStack-owned material is licensed under Apache-2.0; this product direction is not a claim of implemented multi-host support. It unifies API-driven ownership, durable intent, scheduling, safe reconciliation and operations. Primary compute is a guest-kernel VM; App Containers remain an explicit lower-trust alternative. Read the [resource model](resource-model.md).

The value proposition is an observable request-to-ready-to-recovery workflow, not a large service count or a claim that other tools lack these mechanisms. Emulators can run inside provisioned environments as optional application-test dependencies. Incus itself already manages infrastructure; CoveStack must demonstrate useful additional integration rather than rename its API.

## User hypotheses

| Potential user | Problem to validate | Small pilot | Evidence against adoption |
| --- | --- | --- | --- |
| Small development team | Separate DB/storage and predictable cleanup without repeated manual setup | Two accounts provision a fixed app/DB/bucket environment and recover interrupted creation | Existing Compose/Incus/platform workflow is simpler or sufficiently reliable |
| Organization infrastructure operator | Account/project quotas, private environments and an audit trail on owned servers | Provision two owned application environments, inspect denied operations and safe cleanup | Host capacity/operations burden exceeds the benefit of existing tools |
| DevOps engineer | API-driven test infrastructure and failure/recovery visibility | Script lifecycle requests and verify idempotent retry and restored data | Needs AWS compatibility, remote CI integration or HA CoveStack cannot provide |
| Self-hosting/homelab operator | Own hardware, API/CLI control, several apps/databases with understandable ownership | Clean install, one VM and trusted App Container, inspect and remove owned data | Direct Incus/Coolify/Dokku provides the desired workflow with less administration |

These are proposed pilots, not current capabilities or proof of customers. No lab-management, assignment, instructor dashboard, learning/assessment or built-in CI/CD product is included. Prefer one initial segment after interviews; none is confirmed.

## Proving a reason to adopt

| Hypothesis | Required product behavior | Comparison experiment |
| --- | --- | --- |
| Safer interrupted operations | Persist request/result; inspect deterministic ownership; block uncertain takeover | Kill during VM create, reconnect, prove exactly one owned VM and no unrelated deletion |
| Useful shared ownership | Consistent account/RBAC at API, network, disk and backup boundaries | Two accounts with same friendly names; unauthorized requests and packets fail |
| Lower operational burden | Reproducible install, visible prerequisites and explicit retention | Independent operator installs and recovers without author assistance; record errors and manual steps |
| Better recovery understanding | Revision, attempt, freshness and sanitized failure evidence | Operator diagnoses an induced failure and restores known data; measure task completion |
| Integrated VM/application workflow | VM plus private database/storage under capacity admission | Compare equivalent scenario using direct Incus and existing tools; include maintenance cost |

Interview potential users before expanding scope. Record their current tool, pain, required isolation, hardware, willingness to pilot and reasons to refuse. Separate stated enthusiasm from an actual repeat use/pilot commitment. Decide advance/revise/stop after evidence; no arbitrary customer-count or success-rate claim is imposed.

## Non-goals

No full cloud API parity, public-cloud billing/SLA, arbitrary host-command service, mandatory external cloud account, compulsory Floci dependency, or immediate HA. Single-host hardening does not survive loss of that host. A CLI/UI form, HMAC or choice of language is not independently a competitive advantage. [Competitive analysis](competitive-analysis.md) and [core/future dependencies](future-services.md) bound these claims.
