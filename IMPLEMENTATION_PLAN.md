# Incident IQ Go SDK Build Plan (`incidentiq-sdk-golang`)

## Summary
Build a production-ready Go module (`github.com/herooftimeandspace/incidentiq-sdk-golang`,
package `incidentiq`) with a single context-aware client, generated wrappers for
the bundled Incident IQ contracts, generated reference documentation,
contract-driven tests, and CI/CD. The contract artifacts are shared with
`herooftimeandspace/incident-py-q`; the API shape, documentation, and tooling
are this repository's own.

## Key Constraints and Public Interfaces
- Go target is the toolchain version in `go.mod` (**1.26.4+**).
- Standard library only: no third-party dependencies.
- Default auth behavior is **Bearer token** (`Authorization: Bearer <token>`).
- Tenant base URL is explicit per client, with a separate variable set for
  integration smoke tests.
- Public entry points:
  - `incidentiq.NewClient(Config) (*Client, error)`
  - `incidentiq.NewClientFromEnv() (*Client, error)`
- Unified request method:
  - `(*Client).Request(ctx context.Context, method string, path string, opts RequestOptions, out any) error`
  - plus `RequestGolden` / `RequestSilver` for inventory-resolved routes
- Generated wrappers share one signature:
  - `func (s *Service) Method(ctx context.Context, opts RequestOptions, out any) error`
- Config surface (`incidentiq.Config`):
  - `BaseURL` (tenant-specific, required unless provided by env)
  - `APIToken`
  - `SiteID` (optional)
  - `AuthMode` (default `bearer`; optional `raw`)
  - `AppHeaders`, `Timeout`, `MaxResponseBytes`, `MaxRetries`, `BackoffBase`, `HTTPClient`
- Tenant URL env vars:
  - `INCIDENTIQ_BASE_URL` for normal SDK usage
  - `INCIDENTIQ_TEST_BASE_URL` for integration/smoke tests
  - integration tests skip cleanly with clear messaging when credentials are missing

## Implementation Changes
- Flat package at the module root:
  - transport (`client.go`, `request.go`, `retry.go`) with retries and bounded responses
  - auth/header policy (bearer default, optional `SiteId` and app headers)
  - configuration validation (`config.go`) with HTTPS-only base URLs and header-value checks
  - embedded contract artifacts and accessors (`data.go`)
  - hand-written typed Silver helpers where a generic wrapper would be unsafe
- Code generation from the bundled artifacts, both run by `go generate ./...`:
  - `scripts/generate_wrappers.go` -> `generated_wrappers.go`
  - `scripts/generate_sdk_reference.go` -> `docs/sdk-reference/`
- Contract sync:
  - `scripts/sync_from_source_sdk.sh` copies machine-readable artifacts from a
    local `incident-py-q` checkout; HAR classification and OpenAPI fetching stay
    upstream
  - documentation is never synced
- Surface policy:
  - Golden wrappers directly on `client.<Namespace>.<Method>`
  - Silver wrappers under `client.Silver.<Namespace>.<Method>`, app routes under
    `client.Silver.Apps.<AppNamespace>.<Method>`
  - deprecated forwarders for every pre-migration Golden name, with
    `// Deprecated:` comments and `LegacyAliasConflicts()`

## Test Plan
- Unit tests:
  - auth/header behavior, base URL normalization, tenant-root path handling
  - request construction, retries, timeouts, size limits, error propagation
  - configuration validation and environment parsing
- Contract tests:
  - generated wrapper inventory against the bundled golden snapshots
  - generated reference against the generated wrappers
    (`TestGeneratedReferenceMatchesWrappers`)
  - reference pages stay Go-shaped (`TestGeneratedReferenceIsGoShaped`)
- Integration tests:
  - use `INCIDENTIQ_TEST_BASE_URL` explicitly
  - read-only smoke endpoints only
  - skip cleanly when credentials or base URL are absent
- Coverage:
  - `go test -covermode=atomic -coverprofile=coverage.out ./...`
  - ratcheted per branch, with a permanent `95.0%` floor

## Docs, CI/CD, and Artifact
- Docs:
  - `README.md`, `CHANGELOG.md`, `CONTRIBUTING.md`, `SECURITY.md`
  - `docs/` guides plus the generated `docs/sdk-reference/`
  - static site built by `scripts/build_docs_site.go` and published to GitHub Pages
  - Go doc comments as the package API reference
- CI:
  - `quality.yml`: `go vet`, tests, coverage ratchet, badge payloads
  - `integration.yml`: live tenant tests, secret-gated
  - `docs.yml`: docs build and Pages publish from `main`
  - `promotion.yml`, `release-prep.yml`, `release-prep-check.yml`, `release.yml`:
    `dev -> staging -> main` promotion, semver labels, tagged GitHub Releases
- Plan artifact:
  - `IMPLEMENTATION_PLAN.md` (this document)

## Assumptions and Defaults
- Module path `github.com/herooftimeandspace/incidentiq-sdk-golang`; package `incidentiq`.
- Releases are `vX.Y.Z` Git tags, which is what `go get` resolves.
- Bearer token is the default auth mode.
- Tenant URL is configurable per client instance and separately for integration smoke tests.
