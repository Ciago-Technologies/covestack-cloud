# CoveStack V5 — React UI, CLI Parity and Monitoring

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Single source of state

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#10-react-ui). CORE-5 has priority; no capability is claimed here.

## 2. UI screens

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#10-react-ui). CORE-5 has priority; no capability is claimed here.

## 3. Forms and secret handling

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#10-react-ui). CORE-5 has priority; no capability is claimed here.

## 4. CLI behavior — CORE-5 interface

E owns the thin cobra CLI; A owns pkg/kanclient and reviews Go. CLI calls only the authenticated API, never Docker/k3d or the DB. Config: server URL/key ID/profile; secrets env/keychain; no local DB. Every mutation sends an idempotency key; create returns pending; --wait polls with finite timeout. Get shows desired/converged generation, lease and observed state; every list/get supports -o json.

Build order: droplet → vpc → k8s → db → bucket → stack; next production service starts only after the previous failure test passes in CI. W3 K8s learning is mock-only. No UI work during CORE-5. [Go handover](CoveStack-Go-Handover.md) defines E's plan; prior broader parity details remain in the roadmap.

## 5. Monitoring model

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#10-react-ui). CORE-5 has priority; no capability is claimed here.

## 6. User workflow acceptance

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#10-react-ui). CORE-5 has priority; no capability is claimed here.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
