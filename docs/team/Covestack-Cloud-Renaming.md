# CloveStack Cloud — Folder and GitHub Rename

| Change | Effect | Required action |
|---|---|---|
| Local `kan-cloud` → `clovestack-cloud` | Moving the entire folder keeps `.git`, commits, branches and files. Remote URL does not change. | Close running tasks; move the whole folder; reopen editor; update absolute paths, scripts and systemd working directories. |
| GitHub repository rename | GitHub redirects normal old repository URLs and Git operations. Admin permission is required. | Rename in repository Settings; each teammate updates `origin`. Do not reuse the old repository name. |
| Protected main | Renaming does not grant push rights or remove branch protections. | Continue feature branches and reviewed PRs. |
| GitHub Pages / hosted custom Actions | Pages project URLs and callers of Actions hosted in the renamed repository need attention. | Check Pages settings and workflows using the old owner/repository path. |
| Docker Compose | Project name controls container/network/volume grouping. With the current `-f deploy/docker-compose.yml`, the default folder is `deploy`; root-folder rename alone normally preserves it. | Inspect `docker compose ls`; preserve the actual existing project name with `-p` if necessary. Never use `down -v` to fix a rename. |
| Go/Java folders and protocol | Branding does not require changing `kan-agent`, `kan-server`, module/package names, `KAN_*` variables or `X-Kan-*` headers. | Keep compatibility names; coordinate separately if changing public interfaces. |

## Apply documentation first
```bash
cd /run/media/anujtiwari/DevMount/Project-CDAC/kan-cloud
git apply --check /run/media/anujtiwari/DevMount/Claude/CloveStack-Docs-Readability.patch
git apply /run/media/anujtiwari/DevMount/Claude/CloveStack-Docs-Readability.patch
python3 contracts/check_lifecycle.py
git status --short
```

| Step | Action |
|---|---|
| 1 | Review and commit lifecycle plus documentation changes on your current feature branch. |
| 2 | Push the feature branch and open/update the PR; do not push to protected main. |
| 3 | Ask the repository admin to rename GitHub to `clovestack-cloud` when the team agrees. |
| 4 | Run the following local rename and remote-update commands after the GitHub rename succeeds. |
```bash
cd /run/media/anujtiwari/DevMount/Project-CDAC
mv -- kan-cloud clovestack-cloud
cd clovestack-cloud
git remote set-url origin https://github.com/Ciago-Technologies/clovestack-cloud.git
git remote -v
```

| Official reference | Link |
|---|---|
| GitHub rename behavior | https://docs.github.com/en/repositories/creating-and-managing-repositories/renaming-a-repository |
| Compose project-name rules | https://docs.docker.com/compose/how-tos/project-name/ |
