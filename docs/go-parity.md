# Go SDK Parity Notes

This repo is the Go companion to `herooftimeandspace/incident-py-q`.

The copied Markdown files under this repo are intentionally retained so the Go
SDK can be reviewed against the same product, contract, validation, and release
documentation as the Python SDK. When the Python repo changes a shared contract
or user-facing behavior, run `scripts/sync_from_incident_py_q.sh` and then update
the Go runtime or generated wrappers until the tests prove parity again.

## Shared Runtime Behavior

The Go client currently matches the Python SDK for the following runtime rules:

- `INCIDENTIQ_*` environment variable names
- bearer and raw authorization modes
- `Client` header default value of `ApiClient`
- optional `SiteId` header
- HTTPS-only base URL validation
- tenant root normalization to `/api/v1.0`
- tenant-root handling for `/api/`, `/services/`, `/apps/`, `/img/`, and `/s/`
- JSON object and array response decoding
- retry behavior for idempotent methods and retryable HTTP statuses

## Contract Artifacts

The Go repo embeds these synced Python SDK artifacts:

- `data/stoplight/controllers/*.json`
- `data/postman/collection.json`
- `data/source_manifest.json`
- `data/app_schemas.json`
- `data/silver_inventory.json`
- `testdata/contract/golden_sdk_inventory.json`
- `testdata/contract/silver_sdk_inventory.json`
- `testdata/contract/merged_sdk_inventory.json`

## Remaining Parity Work

The Go repo still needs generated typed wrappers for each Golden and Silver SDK
inventory entry. Until those wrappers are generated, callers can use
`RequestGolden`, `RequestSilver`, or `Request` to reach the same routes with the
same runtime semantics.
