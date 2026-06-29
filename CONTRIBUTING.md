# Contributing

## Development Setup

1. Use the Go version declared in `go.mod`.
2. Clone this repository.
3. Run the local test suite:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

The repo-local cache variables are useful on workstations where the default Go cache path is sandboxed or shared across projects.

## Branch Promotion Flow

- Branch feature, bugfix, and chore work from `dev`.
- Merge reviewed work into `dev` after the `test` workflow passes.
- Promote `dev -> staging` after green pushes on `dev`.
- Promote `staging -> main` after green pushes on `staging`.
- Keep `dev`, `staging`, and `main` aligned with the source repository workflow unless a repository-specific decision changes that policy.

## Quality Gates

Run all required checks before opening a pull request:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

The current GitHub Actions workflow runs the same package test command.

## Release Labels and Mainline Releases

- Every `staging -> main` promotion PR should carry exactly one of:
  - `semver:patch`
  - `semver:minor`
  - `semver:major`
- Version tagging and release asset publication are not implemented yet.
- Do not document or rely on release artifacts that are not produced by the current workflows.

## Schema Sync Workflow

The runtime client and contract tests use bundled schema artifacts only. To refresh bundled upstream contracts from a sibling source checkout:

```bash
./scripts/sync_from_source_sdk.sh ../incident-py-q
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

If schema updates change generated SDK method inventory, commit the refreshed files under `testdata/contract/` in the same change.

## Public API Stability

The public SDK surface (`Client`, `Config`, `RequestOptions`, request methods, and inventory helpers) is semver-governed:

- Backward-compatible additive changes: minor version bump.
- Breaking method, type, or signature changes: major version bump.
- Bug fixes that preserve public surface: patch version bump.
