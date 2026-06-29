# Incident IQ SDK Build Plan (`incidentiq-sdk-golang`)

## Summary

Build a production-ready Go module (`github.com/herooftimeandspace/incidentiq-sdk-golang`) that stays functionally aligned with the `incident-py-q` source repository.

Primary contract source is the bundled Incident IQ Stoplight Swagger 2.0 controller specs copied from the source repository. The APIHub Postman collection and Silver HAR-derived inventory are bundled as secondary contract corpora.

## Key Constraints and Public Interfaces

- Go module target is the version declared in `go.mod`.
- Default auth behavior is **Bearer token** (`Authorization: Bearer <token>`).
- Client supports explicit tenant URL configuration for SDK/runtime and integration smoke tests.
- Public entry points:
  - `incidentiq.NewClient`
  - `incidentiq.NewClientFromEnv`
  - `incidentiq.ConfigFromEnv`
  - `incidentiq.GoldenInventory`
  - `incidentiq.SilverInventory`
- Unified request method:
  - `Request(ctx, method, path, opts, out) error`
- Inventory-backed request helpers:
  - `RequestGolden(ctx, namespace, name, opts, out) error`
  - `RequestSilver(ctx, namespace, name, opts, out) error`
- Config surface:
  - `BaseURL` (tenant-specific, required unless provided by env)
  - `APIToken`
  - `SiteID` (optional/required depending on endpoint policy)
  - `AuthMode` (default `bearer`; optional `raw` for compatibility)
  - `AppHeaders`
  - `Timeout`
  - `ValidateResponses`
  - `MaxRetries`
  - `BackoffBase`
  - `HTTPClient`
- Tenant URL env vars:
  - `INCIDENTIQ_BASE_URL` for normal SDK usage
  - `INCIDENTIQ_TEST_BASE_URL` for integration/smoke tests

## Implementation Changes

- Package layout at repository root:
  - `client.go`: transport, request construction, retries, and inventory-backed calls
  - `config.go`: config loading, auth mode, and URL/header validation
  - `data.go`: embedded contract and inventory access
  - `errors.go`: exported error types
  - `request.go`: request options
  - `retry.go`: retry policy helpers
- Contract artifacts:
  - `data/stoplight/controllers/*.json`
  - `data/postman/collection.json`
  - `data/source_manifest.json`
  - `data/app_schemas.json`
  - `data/silver_inventory.json`
  - `testdata/contract/*_sdk_inventory.json`
- Sync tooling:
  - `scripts/sync_from_source_sdk.sh`
- Current runtime behavior:
  - method/path request construction
  - path parameter escaping
  - tenant-root URL handling for Golden and Silver paths
  - JSON body encoding
  - JSON object/array response decoding
  - retry handling for idempotent methods and retryable statuses

## Test Plan

- Unit tests:
  - auth/header behavior
  - optional `SiteId` header
  - request construction
  - tenant URL normalization
  - tenant-root path handling
  - retries and error propagation
  - bundled inventory loading
- Contract tests:
  - verify Golden and Silver inventories load from embedded files
  - future wrapper-generation tests must compare generated Go method coverage against bundled inventory snapshots
- Integration tests:
  - not implemented yet
  - future tests should use `INCIDENTIQ_TEST_BASE_URL` explicitly
  - future tests should stay read-only unless a checked-in test plan documents a write-safe scenario

## Docs, CI/CD, and Artifact

- Docs:
  - `README.md`, `CHANGELOG.md`, `CONTRIBUTING.md`, `SECURITY.md`
  - `docs/go-parity.md`
  - copied source-reference Markdown kept for parity tracking until Go-generated docs replace it
- CI:
  - GitHub Actions runs `go test ./...` on `main`, `dev`, `staging`, and pull requests.
- Plan artifact:
  - `IMPLEMENTATION_PLAN.md` (this document)

## Assumptions and Defaults

- Module path: `github.com/herooftimeandspace/incidentiq-sdk-golang`.
- Bearer token is default auth mode.
- Tenant URL is always configurable per client instance and separately for integration smoke tests.
- One generated Go method per Golden and Silver inventory entry remains tracked in GitHub issue #1.
