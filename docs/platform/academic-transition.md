# Historical academic scope reference — not product authority

As of 4 November 2026, the user explicitly removed academic CORE-5 as a constraint on the product. The [product catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../product/CoveStack-Product-SRS.md) define active scope independently. The body below records the previous academic transition discussion, including old limits and approval dependencies; none is a gate on future product functionality. Past approvals, dates and evidence are not rewritten.

<!-- Historical academic transition body preserved below. -->

# Academic obligations and VM architecture transition

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **No mentor approval, requirement removal or revised delivery schedule is recorded.** This register preserves obligations while the product target changes.

## Distinguish three agreements

| Agreement                  | Current interpretation                                                                                                                                                                          |
| -------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Existing academic V5       | Five categories: compute, required Kubernetes, Managed PostgreSQL, object storage, VPC/IPAM; tests 1–15 plus fixed stack, CLI, security/recovery and Ansible Part 1; optional bounded Terraform |
| New product architecture   | User requires true VM Instance as primary compute and a separate App Container; preferred Incus evaluation; measured capacity, no fixed two-driver ceiling                                      |
| Actual wire/implementation | Contract 1.1.0 remains container-shaped compute; agent communicates without execution; VM schema/driver and full server/foundation absent                                                       |

The original V5 compute definition is superseded as a forward architecture, but its acceptance obligations are not erased. A VM pivot changes image/spec/network/storage/readiness/security/capacity assumptions. B+C+A/D and mentor must review revised SRS mapping and scope feasibility before claiming academic acceptance. Preserve old requirement numbers and historical results; do not present old container proof as new VM acceptance.

## Obligation mapping

| Existing obligation    | Product interpretation / required review                                                                                                                                                        |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Compute tests 1–6      | Keep idempotency/tenancy/stale report/outage/fencing/crash intent; VM create proves one guest-kernel VM; separate App Container needs its own cases; runtime-uncertainty test remains mandatory |
| Kubernetes tests 7–10  | Required academic scope remains; retain 2-cluster/1-server/max-2-agent safety ceiling pending measured admission; no guarantee 8 GB fits it                                                     |
| PostgreSQL test 11     | Actual private scoped SQL, backup and known-row restore, regardless of selected VM/container hosting profile                                                                                    |
| Object tests 12–13     | Preserve bucket/key/retained-byte evidence; map historical MinIO test 13 to a selected provider only through team/mentor review                                                                                                        |
| VPC/IPAM test 14       | Allocation races plus actual intended/denied VM/container paths; old Docker bridge helper does not automatically support VM NICs                                                                |
| Fixed stack test 15    | One app + PG + bucket workflow; map app VM versus distinct trusted App Container explicitly before implementation                                                                               |
| Installer/CLI/security | Still required cross-cutting evidence; VM host prerequisites and trust boundary change need new verification                                                                                    |
| Terraform              | Optional and bounded; existing planned Droplet compatibility must not silently reinterpret old specs                                                                                            |

Kubernetes is core in the proposed product catalogue and remains an academic requirement; the former product-release grouping is superseded historical planning. No new SRS IDs or extra service-count claim is invented. VM/App Container product expansion is not a covert sixth academic deliverable; assess/revise the formal scope explicitly.

## Schedule conflict

V5 records 25 January 2027 and freeze after 10 January. The latest user statement gives 15 December as the deadline without restating the year; current context suggests 2026 but confirmation is needed before a replacement calendar. This task sets no new deadline, mentor approval or promise that the VM pivot fits either plan. Old schedules remain historical records, not current feasibility evidence.

## Team responsibilities retained

A: all agent/client/driver/signing/reconcile Go, CI, installer/ITISS. B: Java core/all SQL/leases/idempotency/quotas/stacks. C: auth/security, all acceptance tests, Java allocator. D: Java API/schema/DTO/fixtures, no Go. E: thin CLI/JSON/packaging, later optional provider, not agent. A reviews Go; A+B+C review shared-client/VM backend boundaries, B+C+A schema, D DTO. Mentor ITISS coding permission remains unresolved.

Academic deliverable changes, wire evolution, backend pinning, host-profile support and new scheduler ownership require explicit reviews in [decisions](decisions.md). No names, usernames or handover approvals are invented.
