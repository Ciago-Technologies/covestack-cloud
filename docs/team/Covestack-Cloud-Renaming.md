# CoveStack Cloud — Folder and GitHub Rename

The local checkout is now `Project-CDAC/covestack-cloud`, and `origin` points to `https://github.com/Ciago-Technologies/covestack-cloud.git`. No further folder rename is needed for this checkout.

| Item | Action |
|---|---|
| Product spelling | Use **CoveStack Cloud** in titles and descriptions; use `covestack-cloud` for the repository slug. |
| Existing checkout | Reopen the editor in `covestack-cloud` if it still displays the old workspace name. |
| Compatibility identifiers | Keep `kan-agent`, `kan-server`, Go module/import paths, Java packages, `KAN1`, `KAN_*`, `X-Kan-*`, API paths, ownership labels, database identity and Valkey prefixes unchanged. A technical rename requires a separate coordinated migration. |
| Older Claude artifacts | Patches and team ZIPs are historical delivery/backup files. Security, branding and lifecycle/documentation changes are already committed; do not reapply old patches or overwrite this checkout with a generated copy. |
| Team workflow | Commit and push changes on a feature branch, then raise a reviewed PR. |

## Verify this checkout

Run from the repository root:

```bash
pwd
git remote -v
python3 contracts/check_lifecycle.py
git status --short
```

The Claude artifacts are at `/run/media/anujtiwari/DevMount/Claude`. From this checkout, reach them with `cd ../../Claude`.

## Other checkouts still named kan-cloud

Close running tasks and the editor before moving another checkout. Run these commands only if `kan-cloud` exists and `covestack-cloud` does not exist in that parent directory:

```bash
cd /path/to/repository-parent
mv -- kan-cloud covestack-cloud
cd covestack-cloud
git remote set-url origin https://github.com/Ciago-Technologies/covestack-cloud.git
git remote -v
```

Update absolute paths in editor workspaces, scripts and service working directories. Keep the existing Docker Compose project name and volumes; do not use `down -v` for a rename. Check any GitHub Pages settings or callers of Actions hosted in the renamed repository separately.

| Reference | Link |
|---|---|
| GitHub repository renaming | https://docs.github.com/en/repositories/creating-and-managing-repositories/renaming-a-repository |
| Compose project names | https://docs.docker.com/compose/how-tos/project-name/ |
