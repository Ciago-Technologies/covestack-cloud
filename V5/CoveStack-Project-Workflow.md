# CoveStack V5 — Delivery Workflow to 25 January 2027

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Schedule authority

**Submission is Monday, 25 January 2027.** This V5 workflow replaces the older December submission timetable for planning. Gates are acceptance targets, not predictions that unfinished dependencies will disappear. Review the extended scope with the mentor and revise assignments openly if a gate slips.

The first release slice is Droplets/Kubernetes. A complete web/data template requires working PostgreSQL, S3 and network/publication dependencies, so those cannot be postponed beyond the stack milestone.

## 2. Dated gates

| Gate/date | Accepted result | Lead pairing |
|---|---|---|
| G0 — 16 Oct 2026 | V5 scope, owners, contract delta and tool matrix agreed | Entire team |
| G1 — 23 Oct 2026 | Booting backend, real auth, typed client and first Droplet flow | A/B/C/D |
| G2 — 6 Nov 2026 | Droplet lifecycle, restart recovery, VPC/IPAM and initial CLI | A/B/C/E |
| G3 — 20 Nov 2026 | Kubernetes lifecycle, usable kubeconfig, public app ingress | A/B/C/D/E |
| G4 — 4 Dec 2026 | Persistent PostgreSQL, scoped S3 and first full stack template | A/B/C/E |
| G5 — 18 Dec 2026 | UI/CLI stack parity, Terraform vertical slice, Ansible example | D/E with A/B/C |
| G6 — 8 Jan 2027 | LB/ASG/snapshot acceptance; metadata and monitoring evidence | A/B/C/D/E |
| Freeze — 11 Jan 2027 | No new scope; supported-feature matrix fixed | Entire team |
| RC — 15 Jan 2027 | Packaged artifacts, full tests and clean-VM run | A/E plus owners |
| Rehearsal — 22 Jan 2027 | External-user demo, security and restoration evidence | Entire team |
| Buffer — 23–24 Jan 2027 | Critical repairs and submission assembly | Entire team |
| Submit — 25 Jan 2027 | Tested source/artifacts/docs/evidence delivered | Entire team |

## 3. Stage-specific work

| Period | Application work | Anuj / ITISS work |
|---|---|---|
| 14–23 Oct | Main/migrations/auth/client; real compute | Baseline VM, CI, signing and threat model |
| 24 Oct–6 Nov | Lifecycle/actions/leases; CLI compute | Restart/failure tests and address plan |
| 7–20 Nov | Kubernetes and ingress records | Cluster network, DNS/TLS and tunnel proof |
| 21 Nov–4 Dec | DB/S3 contracts and stack engine | Persistent data, policies and restore trial |
| 5–18 Dec | React, provider slice and playbook | Controlled automation runner and parity tests |
| 19 Dec–8 Jan | LB/ASG/snapshots/metadata | Health, isolation, outage and scan evidence |
| 9–15 Jan | Defects/package/install | Repeat install, backups and compatibility |
| 16–25 Jan | Rehearsal and final fixes | Security demonstration and handover |

## 4. Dependency order

```text
Auth + durable state + API + signed agent
                  |
       Droplet lifecycle and recovery
           /                   \
     VPC/IPAM             Kubernetes driver
           \                   /
          Public ingress + connectivity
                  |
         Persistent DB + scoped storage
                  |
          Stack dependency orchestration
                  |
       UI/CLI parity + Terraform/Ansible
                  |
         Expanded services + release tests
```

Simple UI layouts and test fixtures can be designed early. Enabled forms and successful results must wait for real APIs. Provider development starts with one proven resource, not every service schema at once.

## 5. Daily engineering workflow

| Step | Expected artifact |
|---|---|
| Plan one slice | Behavior and acceptance case written first |
| Branch | Agreed reviewed baseline; no direct main push |
| Implement | Small change respecting ownership/contracts |
| Verify | Appropriate tests and actual integration evidence |
| Review | Boundary owners inspect behavior and errors |
| Merge | Required checks and approved PR |
| Record | Feature matrix, evidence and known limitations updated |

Standup: last accepted result, today's acceptance target and current dependency. A blocker lasting a working day triggers pairing or reassignment, not silent schedule drift. Do not force-push teammates' branches or bypass policy to meet a gate.

## 6. Scope and gate failure policy

| Situation | Response |
|---|---|
| Backend/auth unavailable | Continue isolated client/infra tests; mark live gate pending |
| Kubernetes scope too large | Keep one tested shape/version; document unsupported actions |
| BYOD not ready | Ship team-domain publication; report BYOD gap explicitly |
| Full-template dependency incomplete | Do not offer a success-shaped stack with fake services |
| Nested Docker requested | Optional after release gates; do not weaken workload privileges |
| Gate missed | Revise plan with mentor; preserve honest feature status |

All agreed capabilities are targets. The team may negotiate reduced scope, but V5 does not silently remove a requested service. No feature is accepted only because it has a UI form, schema enum or screenshot.

## 7. Final demonstration

Create a template through UI, inspect its resources through CLI, publish its app to an external device, deny another account's action, restart the agent without duplication, change one owned resource through Terraform, configure through approved Ansible connectivity, and restore known data separately. Show the compatibility/limitation matrix alongside the demonstration.

## 8. Submission package

| Material | Requirement |
|---|---|
| Source | Reviewed tested revision with history |
| Binaries/images | Exact versions/digests and checksums |
| Documentation | V5 plus actual runtime command updates |
| Test evidence | Commit/date/environment/results and unresolved issues |
| Security | Threat model, policy matrix, remediation and residual risks |
| Recovery | Repeated restore evidence with measured scenario |
| Demonstration | Script and backup recording; no secrets visible |
