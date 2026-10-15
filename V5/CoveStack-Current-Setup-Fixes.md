# CoveStack V5 — Baseline Audit and Extension Checklist

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Preserved baseline findings

The 38 baseline audit IDs remain review inputs. Their older target dates are superseded by the V5 workflow. The table preserves the required remedy and acceptance but does **not** mark every finding unimplemented or resolved. Link each ID to current commits/tests before setting its status. **Historical remedy text is not a live implementation claim, version/tool decision or extra SRS requirement.** Current role/scope/contract decisions override it; items 23, 29–32 and other broader portions are deferred where outside CORE-5.

| ID  | Required remedy                                                                                                                                                                                    | Acceptance                                                                                   |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| 1   | Unsigned nonce permits modified replay → require/sign nonce and key ID, collect them in Java and regenerate signatures                                                                             | Change only nonce/key ID without resigning: rejected                                         |
| 2   | Nonces stored before authentication → verify first, then atomic authenticated SET NX EX; invalid ingress separately limited                                                                        | Bad signature does not consume nonce; concurrent identical valid requests execute once       |
| 3   | readAllBytes unbounded → enforce 1 MiB declared and actual-byte cap, including chunked input, with 413                                                                                             | Oversized known-key bad-MAC/chunked input rejected without full allocation                   |
| 4   | Public/Kan routes skipped and no wiring → register filters and explicitly cover all routes except minimal liveness                                                                                 | Real HTTP routing rejects unsigned/invalid requests on protected paths                       |
| 5   | Secret-only lookup lacks identity/kind → active account/key/kind/permissions record, trusted principal and scoped reads/updates                                                                    | User credential denied on agent route; cross-account IDs denied; principal exists            |
| 6   | Memory store ignores TTL/restarts → shared atomic store; bounded expiring test-only fallback                                                                                                       | Replay rejected across instances/restart policy; store outage fails closed                   |
| 7   | Missing required signed headers silently omitted → reject missing/duplicate/ambiguous input and define normalization                                                                               | Missing host/content type/nonce fails; raw-path/query edge vectors agree                     |
| 8   | Python signs all supplied headers → shared selection policy; load real vectors and UTF-8 bytes                                                                                                     | Python/Java/Go compare same complete header inputs                                           |
| 9   | Startup commands target nonexistent apps → implement entry points; fix Compose path and runnable Go command                                                                                        | Clean checkout runs documented app and real agent                                            |
| 10  | Boot 3.3.5 outside Java 25 support → reviewed maintained compatible parent + one tested matrix                                                                                                     | Java 25 build/package/boot passes; dependency scan triaged                                   |
| 11  | Postgres Flyway plugin absent → add aligned flyway-database-postgresql; mail starter before email                                                                                                  | Actual Postgres migration succeeds from empty database                                       |
| 12  | Only Gitleaks exists → Java verify, Go test/vet, contract validation and integration checks                                                                                                        | Real required CI jobs execute and fail on regressions                                        |
| 13  | Unverified archive/moving tags → pinned action commits/tools/images and checksum verification                                                                                                      | Tampered download fails; release artifacts/digests reproducible                              |
| 14  | Default secrets/implicit networking → separate secrets/config, explicit backend/storage and workload-account policies                                                                              | Backend inaccessible from client/workload IP; no default passwords                           |
| 15  | AOF without managed persistence policy → named /data volume; authenticated durable store and fail-closed loss handling                                                                             | Store replacement/restart preserves valid-window protection or blocks until recovery         |
| 16  | Missing timestamp/nonnullable generation → required observedAt, nullable int64 generation, report ordering and applied response                                                                    | Null health reports/schema examples pass; stale report ignored                               |
| 17  | Conflicting kinds/states → corrected per-kind vocabulary from design section 5; add DB stopped state                                                                                               | Every accepted state/example valid for its kind                                              |
| 18  | Planned features outside contract → complete seven-kind lifecycle catalog including storage.bucket; groups excluded from agent work; extra ssh-keys/lb-members APIs still require separate schemas | Unsupported kinds reject; extensions validated before drivers merge                          |
| 19  | Auth only prose, unconstrained values → required signing header/security declarations, bounded typed specs, complete scoped snapshots                                                              | Examples/negative payloads validated; partial known-ID response never deletes objects        |
| 20  | Failure acknowledged as success → preserve successful generation, record attempt/error and schedule bounded retry; action errors truthfully                                                        | Failed create/stop/delete never completes successfully                                       |
| 21  | Old generation/timestamp can corrupt state → scoped locked validation, monotonic success, generation/sequence/epoch and timestamp bounds                                                           | Newer intent survives old-generation/newer-time reports                                      |
| 22  | active→active cannot express reboot → explicit restart intent and executed marker                                                                                                                  | Process/container start actually changes; retry does not blindly repeat                      |
| 23  | Only active/recent count creates excess replacements → pending reservations, parent binding, deadlines, serialized scaling, scale-down and outage pause                                            | Slow pull/agent outage stays within desired/max; failed children retired                     |
| 24  | User DB joins control-plane network → account workload network, narrow allowed client routes                                                                                                       | User DB cannot reach backend IP/ports                                                        |
| 25  | Shared root password/shell interpolation → key-only SSH, safe file write/argument arrays, bind/firewall policy                                                                                     | Real SSH works; malicious key/name cannot execute shell data                                 |
| 26  | Plain Ubuntu/no tested sshd capability setup → pinned SSH-ready image with reviewed minimum and real session tests                                                                                 | Key login and session work under exact release capabilities                                  |
| 27  | Recreated DB loses attached-data identity; MinIO mounts wrong path → labeled named DB volumes and miniodata:/data, restore/deletion policy                                                         | Tables/objects survive replacement; backup restore passes                                    |
| 28  | Prefixes/root creds do not isolate S3 → per-account credential policies and scoped physical-bucket metadata                                                                                        | A credential cannot read/write/list B bucket                                                 |
| 29  | Forwarding reaches guest interface, loopback-only services inaccessible → correct guest listeners/proxy, DNS TCP+UDP, laptop-address answer                                                        | Two-device DNS/TLS and user service endpoints pass                                           |
| 30  | Signed Host rewritten; internal publicly proxied; Mailpit HTTP treated as HTTPS → authority preservation, internal denial and explicit restricted TLS/tunnel                                       | Public internal request blocked; private signed Host test succeeds                           |
| 31  | Invalid top-level fragment/stale member updates → valid conf.d include, nginx -t, safe directory replacement and health/IP refresh                                                                 | LB boot/reload and member removal/empty pool behavior pass                                   |
| 32  | Invented snapshot endpoint/full-disk claim → type snapshot action, scoped public image IDs and documented volume exclusion                                                                         | godo snapshot action works; volumes not claimed as captured                                  |
| 33  | Concurrent token race/undefined columns/weak audit → atomic fingerprinted claim, consistent lock/revocation model and append-only identity/outcome audit                                           | Concurrent create yields one resource; altered-token payload rejects; audit immutable to app |
| 34  | Unrestricted credentials/overwritten kubeconfig/auth mismatch → local bootstrap, restrictive ACL/mode, atomic writes, Bearer CLI and separate/merged kubeconfig                                    | Credential mode passes; original kubeconfig preserved                                        |
| 35  | Ignored errors/timeouts/always-zero script → checked errors and bounded assertions, nonzero failed checks, accurate limited-compliance claims                                                      | Deliberate bad state/failed control makes test/script fail                                   |
| 36  | Converged crashes missed/terminal rows churn/partial list dangerous → retry/health protocol, fencing, complete scoped snapshot, deletion acknowledgement/deleted_at                                | Crash heals, stopped resources stay stopped, partial responses never destroy live data       |
| 37  | Conflicting algorithm/scope/layout/versions → one corrected design/workflow and actual executable contract                                                                                         | No active contradictory recipes; docs match reviewed entry-point/feature evidence            |
| 38  | Incorrect weekdays → corrected 2026 dates and explicit submission buffer                                                                                                                           | Calendar dates verified; Dec 15 is Tuesday                                                   |

## 2. Location register

| ID  | Original code/configuration area                                                               |
| --- | ---------------------------------------------------------------------------------------------- |
| 1   | contracts/sigv1-rules.json; Java KanSigV1AuthFilter/KanSigV1; Go sigv1/sign.go; vectors/Python |
| 2   | KanSigV1AuthFilter.doFilterInternal; production NonceStore                                     |
| 3   | CachedBodyHttpServletRequest constructor; ingress settings                                     |
| 4   | New Spring app/filter registration, Bearer verifier and route policy                           |
| 5   | KeyService credential model, filter principal and repositories                                 |
| 6   | New ValkeyNonceStore; InMemoryNonceStore; profiles                                             |
| 7   | Java/Go canonicalization and rules validation                                                  |
| 8   | contracts/sign_example.py; Java tests; vectors                                                 |
| 9   | README, Java application/config/migrations, kan-agent/cmd/kan-agent/main.go                    |
| 10  | kan-server/pom.xml; wrapper/CI/runtime images/versions.env                                     |
| 11  | kan-server/pom.xml and migrations                                                              |
| 12  | .github/workflows/                                                                             |
| 13  | secrets-scan workflow; versions/install/image manifests                                        |
| 14  | deploy/docker-compose.yml; local secrets/networks; VM firewall                                 |
| 15  | Valkey Compose volume/config and readiness/recovery                                            |
| 16  | contracts/internal-api.yaml ObservedReport; examples; Java/Go DTOs                             |
| 17  | contracts/FROZEN-STATES.md; YAML/examples; all fakes/drivers                                   |
| 18  | YAML kind/spec enum and additional agent paths                                                 |
| 19  | OpenAPI security/headers/specs/errors and known-ID protocol                                    |
| 20  | Agent reconcile.handle; Java observed/action/retry service                                     |
| 21  | Observed service/repository; schema/DTO epoch/sequence                                         |
| 22  | Work spec, OperationService, agent restart driver                                              |
| 23  | AsgScheduler/group repo; periodic agent inspection/lease                                       |
| 24  | Managed DB driver/workload network policy                                                      |
| 25  | SSH image build, key injection and host command runners                                        |
| 26  | Image catalog, Dockerfile and HostConfig capability profile                                    |
| 27  | Managed DB volume lifecycle; MinIO Compose                                                     |
| 28  | SpacesService, token/policy issuer, S3 client mapping                                          |
| 29  | VM NAT/DNS/firewall/proxy configs and client metadata                                          |
| 30  | Nginx/host policy and private agent URL                                                        |
| 31  | LB config generator/persistent directory/membership loop; web workload image                   |
| 32  | Droplet action API, SnapshotService/image mapping/docs                                         |
| 33  | Idempotency/operation/credential/event migrations and services                                 |
| 34  | CLI bootstrap/config/kubeconfig export; auth matrix                                            |
| 35  | SDK tests, hardening/health scripts, isolation evidence                                        |
| 36  | Work predicate, periodic inspect, agent lease, orphan cleanup and GC                           |
| 37  | All Claude guides plus repo README/contracts/runtime config                                    |
| 38  | Workflow/calendar                                                                              |

## 3. Current scope and decision register

| CORE-5/cross-cutting item | Required review                                                                                                   |
| ------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| VPC/IPAM                  | C Java allocator; B SQL; A helper; ten IP allocations/ten account VPCs; no custom CIDR                            |
| Polling                   | Verified credential account scope; assignment/multi-account model TBD B+C+A                                       |
| Stack                     | One fixed template, idempotent children and retained data; no DAG engine                                          |
| CLI/optional Terraform    | E thin API client/provider; D Java fixtures; A Go review; no runtime imports                                      |
| Installer                 | A Part 1; checksum/vault/no_log/lint and second run changed=0                                                     |
| Schema/auth               | No schema exists yet; Flyway declared; B finalizes staged migrations; service signing versus hash-only API tokens |

Broader gateway/domain, metadata, React/metrics and DAG requirements are preserved in the [roadmap](CoveStack-Extended-Roadmap.md). The 38 IDs remain intact in both historical tables; none is an invented new requirement ID. Do not reapply a historical remedy to undo a corrected baseline or select an old tool/version from this audit.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
