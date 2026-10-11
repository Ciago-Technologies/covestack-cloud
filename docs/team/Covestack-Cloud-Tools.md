# Covestack Cloud — Tools and Version Guide

Corrected 11 October 2026. Use [setup](Covestack-Cloud-Complete-Working-Setup.md) for project configuration and [workflow](Covestack-Cloud-Project-Workflow.md) for delivery dates.

## 1. One tested version matrix

| Component | Decision | Verification |
|---|---|---|
| Java | JDK 25 for everyone | java, javac and Maven all report 25 |
| Spring Boot | Maintained Java-25-compatible patch, preferably compatible 3.5 line | Pin exact version in pom and CI after support/security review |
| Maven | Reviewed 3.9 release/wrapper | Commit wrapper; resolve and verify |
| Go | Supported patched toolchain; inspected 1.27.1 needs 1.27.2 security update | Record exact toolchain and align CI |
| Control-plane Postgres | Reviewed Postgres 17 digest | Same tested patch on all dev/VM machines |
| Valkey | Reviewed Valkey 8 digest | Auth, SET NX EX, AOF durability/restart tests |
| Mailpit / MinIO | Reviewed pinned image digests | SMTP/S3/persistence/secret-file acceptance |
| Docker / Compose | Tested daemon and Compose v2 | SDK API negotiation and secret mount tests |
| k3d / k3s / kubectl | One compatible tested matrix | Pin versions; do not pair old k3s with arbitrary latest kubectl |

No example claims JDK 21 can compile a Java-25-targeted project. A go.mod language minimum is distinct from the actual toolchain. No moving @latest installation is part of the repeatable release.

[Spring's Java support](https://docs.spring.io/spring-boot/3.5/system-requirements.html) explicitly includes Java 25 for the 3.5 line. [Go's release page](https://go.dev/doc/devel/release) lists 1.27.2 security fixes and the supported-release policy. Record installation date and exact versions; recheck updates before release.

## 2. Who installs what

| Machine | Required | Optional/phase-specific |
|---|---|---|
| Fedora editor host | Git, JDK 25, Maven, IDE, Python, VMware | Go for local tests, local Docker |
| AlmaLinux integration VM | Git, JDK 25, Maven, Go, Docker/Compose, Nginx, DNS tools | k3d/kubectl, psql, S3 client, certificate tooling |
| Windows Java teammates | Git, Temurin JDK 25, Maven, IDE, Docker Desktop/WSL2 | Python, psql, doctl, SSH/S3 clients |
| Mac Java teammate | Git, JDK 25, Maven, IDE, Docker Desktop | Same clients, Go if running real agent |

Windows/Mac developers can run the explicit FakeAgent test profile while the real agent is built. Integration evidence always distinguishes fake and real agent runs.

## 3. JDK and Maven installation

Fedora: use available trusted repository packages for JDK 25 and Maven, then select Java/javac alternatives and set the IDE SDK to 25. AlmaLinux: confirm JDK 25 package availability in its configured repositories; otherwise install a verified Temurin 25 build from the vendor. Do not assume every AlmaLinux release has the same package names.

Windows: download Temurin **JDK 25**, enable JAVA_HOME and PATH during setup, open a new terminal and confirm versions. Install a reviewed Maven 3.9 release and add its bin directory to PATH, or use the repository wrapper once added.

Mac: install a reviewed JDK 25 using a trusted package manager/vendor and point JAVA_HOME at the actual JDK Home, not an unrelated parent directory. Configure Maven/IDE to use the same JDK.

```bash
java -version
javac -version
mvn -version
```

All three must agree. Older runtime compatibility is not inferred from the source looking simple.

## 4. Docker and VMware installation

Use [Docker's RHEL installation guide](https://docs.docker.com/engine/install/rhel/) for the integration VM and [Fedora guide](https://docs.docker.com/engine/install/fedora/) for a Fedora daemon. Add the vendor repository correctly before attempting docker-ce installation. AlmaLinux compatibility must be tested on the chosen release. Windows/Mac use Docker Desktop with the appropriate virtualization backend.

Keep firewalld/SELinux active. Do not put docker0 in a broadly trusted firewall zone as a generic networking fix. Review precise forwarding policies. [Docker documents](https://docs.docker.com/engine/install/linux-postinstall/) that membership of the docker group grants root-level privileges; host execution is not a security exemption.

VMware NAT configuration needs the actual reserved guest IP. Allow HTTPS plus DNS TCP/UDP and narrowly controlled SSH/workload forwards. Do not forward control-plane Postgres/Valkey. Record the actual CPU/RAM/disk allocation and test it against measured workload consumption.

## 5. Go, SDK and executable tools

Download a specific supported Go patch from [go.dev/dl](https://go.dev/dl/), verify its published checksum and install through the trusted OS/vendor process. Preserve local configuration; do not blindly remove an existing installation or unpack a new tree over an old one. Record toolchain selection and restart the terminal/IDE as needed.

Add the reviewed Docker SDK version appropriate to the daemon and available client API. Commit imports first or in the same change, then run tidy and tests. Unused requirements disappear from go.mod; fetching a dependency alone does not implement an agent.

```bash
cd kan-agent
go mod tidy
go test ./...
go vet ./...
```

Pin k3d, kubectl, doctl, mc, Gitleaks, Trivy, Syft and Cosign to reviewed versions. Verify checksums/signatures and architecture. Download to a file, verify, inspect/install; do not execute an unchecked moving-branch curl pipe.

## 6. Authenticated clients and certificates

Use an emulator-issued scoped token. Give doctl the emulator API URL **during auth initialization as well as later commands**, so credential validation does not go to real DigitalOcean. Verify the flags for the pinned doctl version; never use a real cloud token in a demo.

CLI credentials use directory mode 0700 and file mode 0600 on Unix, equivalent restricted ACLs on Windows, and atomic writes. Sensitive output is masked. Scope S3 credentials per account; root MinIO credentials stay with administrators/server.

DNS-01 certificate provisioning uses a protected provider credential file and a tested renewal hook. mkcert fallback requires installing its public CA certificate on the controlled clients. Keep the private CA key local. Prefer SSH tunnels for Mailpit and MinIO admin UIs.

## 7. Build and tool verification

Commands from repository root:

```bash
python3 contracts/sign_example.py
(cd kan-server && mvn -B verify)
(cd kan-agent && go test ./...)
(cd kan-agent && go vet ./...)
```

Python, Go and Java fixtures now use the corrected signed headers and three shared vectors. Passing these does not replace real-server/replay/tenant acceptance. Maven may fail until the actual application/plugin/dependencies exist; fix the error and record the result rather than treating it as a pass.

Go tests run with the package directory as their working directory; changing the shell's current directory is not a universal solution for a wrong fixture path. Package/embed runtime contract data or configure explicit paths; do not rely on ../contracts being present beside a released binary.

Check versions directly and compare them to the pinned matrix. Do not use a single --version flag for every tool; flags differ. If scripting checks, use checked exit codes and pipefail for pipelines. Hardening scripts return nonzero when a check fails.

## 8. Storage and maintenance

Use named volumes for backend Postgres, Valkey, MinIO and every managed database. Mount MinIO at /data. Back up data and master keys through separate protected procedures and test restoration.

Ordinary Compose shutdown does not remove volumes. Do not recommend down -v or blanket docker system prune -a as harmless maintenance; stopped containers/snapshot images may be part of the demo's data. Inspect disk usage and remove only confirmed disposable artifacts.

## 9. CI tools

CI must actually include Java/Go builds, contract validation, negative auth tests and integration tests. Add Gitleaks, dependency/SAST scans and image scans with pinned artifacts/checksums and least-privilege job permissions. Sign immutable release images and keep SBOM/test outputs as artifacts. Scan findings require triage by actual reachability and release policy; unsupported stacks still require an upgrade plan.

See [current setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) for all required files and acceptance tests and [change log](Covestack-Cloud-Documentation-Change-Log.md) for what changed in this guide.
