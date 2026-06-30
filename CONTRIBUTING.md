# Contributing

## Development Setup

1. Use the Go version declared in `go.mod`.
2. Clone this repository.
3. Run the local test suite:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

The repo-local cache variables are useful on workstations where the default Go cache path is sandboxed or shared across projects.

Run native Go coverage before changing shared request behavior, generated
wrappers, or test infrastructure:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test -covermode=atomic -coverprofile=coverage.out ./...
go tool cover -func=coverage.out -o coverage-summary.txt
go tool cover -html=coverage.out -o coverage.html
go run scripts/build_badge_json.go coverage --coverage-file coverage.out --label "coverage local" --minimum 95.0 --output coverage-badge.json
```

The `95.0%` value is the permanent minimum, not the target to drift back
toward. Coverage is ratcheted by branch. If the current published branch badge
shows a higher percentage, new work must keep coverage at or above that higher
percentage. For example, if `dev` is currently at `97.10%`, a new pull request
into `dev` must keep the Go coverage result at `97.10%` or higher even though
that is above the permanent `95.0%` floor. Add focused tests for new or changed
behavior instead of accepting a lower percentage.

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

The `quality` workflow runs `go vet ./...` plus the same package test command
with native Go coverage enabled. It uploads `coverage.out`,
`coverage-summary.txt`, and `coverage.html` as workflow artifacts. The workflow
also publishes branch status and coverage badge payloads to the `badges` branch
for `dev`, `staging`, and `main`.

Feature, bugfix, and chore branches are validated through their pull requests.
The workflow intentionally does not run full unit tests on every non-protected
branch push, which avoids running the same Go test suite once for the branch
push and again for the pull request. Push checks remain enabled for `dev`,
`staging`, and `main` because those runs validate the integrated branch tip,
publish branch badges, update the coverage ratchet source, and trigger the
automated promotion workflow.

Coverage must stay at or above the effective branch floor. The effective floor
is the greater of:

- the permanent minimum of `95.0%`
- the current coverage percentage published for the pull request base branch or,
  for direct branch pushes, the current coverage percentage published for that
  same branch

This means coverage should never drop. When adding code lowers the percentage,
add tests until the result is at least the formerly published percentage for
that branch.

Build the static docs site locally before changing contributor or reference
Markdown:

```bash
go run scripts/build_docs_site.go
```

The `docs` workflow runs the same docs build for `main` pull requests and
publishes the generated site plus the `main docs` badge after `main` pushes.

## Release Labels and Mainline Releases

- Every `staging -> main` promotion PR should carry exactly one of:
  - `semver:patch`
  - `semver:minor`
  - `semver:major`
- Automated promotion PRs copy the semver label from the associated source PR
  when one exists and otherwise default to `semver:patch`.
- The `dev -> staging` promotion PR reuses the authoritative `dev` push `unit`
  result instead of re-running the full unit suite on the promotion PR.
- The promotion workflow reports the required `unit`, `integration`,
  `docs-build`, and `release-prep` checks directly on the Actions-authored
  `promote/staging-to-main` branch. This keeps the automated promotion path
  usable even when GitHub does not attach ordinary `pull_request` workflow runs
  to the workflow-created promotion branch. The ordinary pull request workflows
  do not duplicate that work on the promotion PR; they wait for the matching
  promotion-owned check-run and fail closed if it is missing or failed. The
  ordinary `release-prep` workflow also rechecks current labels and branch
  ancestry because those inputs can change without a new promotion head commit.
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
