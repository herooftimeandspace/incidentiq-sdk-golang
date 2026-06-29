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
- Merge reviewed work into `dev` after the `quality` workflow passes.
- Promote `dev -> staging` with the automated promotion PR created after green
  pushes on `dev`.
- Promote `staging -> main` with the automated `promote/staging-to-main`
  promotion PR created after green pushes on `staging`.
- Keep `dev`, `staging`, and `main` aligned with the source repository workflow unless a repository-specific decision changes that policy.

## Quality Gates

Run all required checks before opening a pull request:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

The `quality` workflow runs the same package test command. The workflow also
publishes branch status badge payloads to the `badges` branch for `dev`,
`staging`, and `main`.

## Release Labels and Mainline Releases

- Every `staging -> main` promotion PR should carry exactly one of:
  - `semver:patch`
  - `semver:minor`
  - `semver:major`
- Automated promotion PRs copy the semver label from the associated source PR
  when one exists and otherwise default to `semver:patch`.
- Merging the prepared `promote/staging-to-main -> main` PR creates a GitHub
  Release, a `vMAJOR.MINOR.PATCH` tag, and a source archive.
- Go module releases are tag-based. There is no package metadata file to bump in
  this repository.
- The workflows use the repository `GITHUB_TOKEN` by default. Do not add a
  personal access token unless repository rules prove that GitHub
  Actions-authored PRs cannot satisfy required checks.

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
