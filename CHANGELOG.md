# Changelog

All notable changes to `incidentiq-sdk-golang` are documented in this file.

The format is based on Keep a Changelog and this project follows Semantic Versioning.

## [0.1.0] - 2026-06-29

### Added
- Initial `github.com/herooftimeandspace/incidentiq-sdk-golang` module.
- Low-level `Client.Request` API.
- `RequestGolden` and `RequestSilver` helpers backed by bundled SDK inventories.
- Generated Golden wrappers directly on `Client`.
- Generated Silver fallback wrappers under `Client.Silver`.
- Bearer-token auth by default, with optional raw mode.
- Tenant base URL support for runtime and future integration tests.
- Embedded Stoplight, Postman, app schema, Silver inventory, and SDK inventory artifacts.
- Sync tooling via `scripts/sync_from_source_sdk.sh`.
- Unit tests for config, headers, path rendering, URL construction, retries, and embedded inventory loading.
- GitHub Actions workflow for `go test ./...`.
