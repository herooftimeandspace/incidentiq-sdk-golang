# Go SDK Parity Notes

This repo is the Go companion to `herooftimeandspace/incident-py-q`.

The copied Markdown files under this repo are intentionally retained so the Go
SDK can be reviewed against the same product, contract, validation, and release
documentation as the source SDK. When the source repo changes a shared contract
or user-facing behavior, run `scripts/sync_from_source_sdk.sh` and then update
the Go runtime or generated wrappers until the tests prove parity again.

## Shared Runtime Behavior

The Go client currently matches the source SDK for the following runtime rules:

- `INCIDENTIQ_*` environment variable names
- bearer and raw authorization modes
- optional `SiteId` header
- HTTPS-only base URL validation
- tenant root normalization to `/api/v1.0`
- tenant-root handling for `/api/`, `/services/`, `/apps/`, `/img/`, and `/s/`
- JSON object and array response decoding
- retry behavior for idempotent methods and retryable HTTP statuses
- Golden wrappers directly on `client.<Namespace>.<Method>` as the correct default SDK path
- Silver wrappers under `client.Silver.<Namespace>.<Method>` for quasi-supported API calls derived from live site interaction HARs, with app routes nested under `client.Silver.Apps.<AppNamespace>.<Method>`

## Contract Artifacts

The Go repo embeds these synced source SDK artifacts:

- `data/stoplight/controllers/*.json`
- `data/postman/collection.json`
- `data/source_manifest.json`
- `data/app_schemas.json`
- `data/silver_inventory.json`
- `testdata/contract/golden_sdk_inventory.json`
- `testdata/contract/silver_sdk_inventory.json`
- `testdata/contract/merged_sdk_inventory.json`

## Generated Wrapper Surface

The Go repo generates wrappers from the bundled Golden and Silver SDK inventory
snapshots. Golden is the golden SDK path, so Golden wrappers are promoted
directly onto `Client` as `client.<Namespace>.<Method>`. Silver is a separate
namespace for quasi-supported API calls derived from live site interaction HARs,
so Silver wrappers stay under `Client.Silver`. App-specific Silver routes are
nested under `Client.Silver.Apps`.

Regenerate wrappers after refreshing inventories:

```bash
go generate ./...
```
