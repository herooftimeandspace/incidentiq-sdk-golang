# Schema and Validation

## Bundled Contracts

Primary contract corpus:
- Incident IQ Stoplight Swagger 2.0 controller specs (`/controllers/*.json`)

Secondary compatibility corpus:
- Incident IQ APIHub Postman collection

Bundled assets live under:
- `data/stoplight/controllers/`
- `data/postman/`
- `data/source_manifest.json`
- `data/app_schemas.json`
- `data/silver_inventory.json`
- `testdata/contract/golden_sdk_inventory.json`
- `testdata/contract/silver_sdk_inventory.json`
- `testdata/contract/merged_sdk_inventory.json`

## Runtime Flow

1. Render path parameters.
2. Build the tenant-aware URL.
3. Merge auth, `Client`, optional `SiteId`, app, and request headers.
4. Encode an optional JSON body.
5. Retry idempotent methods for configured retryable statuses.
6. Decode successful JSON object or array responses into the caller-provided output value.

The current Go runtime does not yet perform contract-backed response schema validation. The bundled schema and inventory files are embedded so validation and generated typed wrappers can be added without changing request semantics.

## Schema Sync Workflow

```bash
./scripts/sync_from_source_sdk.sh ../incident-py-q
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

`sync_from_source_sdk.sh` refreshes the bundled contract artifacts and inventory snapshots from the source repository. If schema updates change the SDK method inventory, commit the refreshed `testdata/contract/` files in the same change.
