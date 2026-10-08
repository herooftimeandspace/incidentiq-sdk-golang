# Contracts and Validation

## Bundled Contracts

Primary contract corpus:
- The published Incident IQ OpenAPI 3.0 contract, rendered as the
  [API reference](https://scopousiiq.github.io/iiq-docusaurus-docs/docs/api/)

This single document replaces the former Stoplight Swagger 2.0 controller specs
and the APIHub Postman compatibility corpus.

Legacy contract corpus:
- `data/legacy/contract.json`, a pruned Swagger 2.0 bundle covering the routes
  the published contract no longer documents. Those routes moved to the Silver
  surface during the OpenAPI migration.

## Bundled Artifacts

This list is the single owner; other documents link here rather than repeating
it. Everything below is embedded into the module with `go:embed`:

| Path | Contents |
| --- | --- |
| `data/openapi/openapi-spec.json` | Golden contract, primary |
| `data/openapi/metadata.json` | Sync provenance |
| `data/source_manifest.json` | Source manifest |
| `data/silver_inventory.json` | Silver inventory: HAR-derived and migrated routes |
| `data/legacy/contract.json` | Schemas for routes migrated off Golden |
| `data/legacy/aliases.json` | Deprecated method-name aliases |
| `data/app_schemas.json` | HAR-derived app-path schemas |
| `data/typed_silver_methods.json` | Routes with hand-written typed helpers |
| `testdata/contract/golden_sdk_inventory.json` | Golden-surface drift snapshot |
| `testdata/contract/silver_sdk_inventory.json` | Silver-surface drift snapshot |
| `testdata/contract/merged_sdk_inventory.json` | Combined drift snapshot |
| `testdata/contract/user_room_mutation_observation.json` | Observed request shape for the typed user-room helpers |

Nothing is fetched over the network at runtime.

## What The Contracts Drive

In this SDK the contracts are a **build-time** input, not a runtime validator:

1. `scripts/generate_wrappers.go` turns the Golden and Silver inventories into
   `generated_wrappers.go`, so every documented route has a Go method on the
   right surface, with the right HTTP verb and path.
2. `scripts/generate_sdk_reference.go` turns the same artifacts plus the OpenAPI
   document into `docs/sdk-reference/`, so each method's parameters, response
   schema, and status codes are documented.
3. `testdata/contract/*_sdk_inventory.json` snapshots gate semver-significant
   drift in the generated surface.

## Response Validation

The Go SDK does **not** validate response payloads against the contracts. A
success response is decoded with `encoding/json` into the `out` argument, and
whatever the tenant returns is what you get.

- `Config.ValidateResponses` defaults to `true` but is currently inert. It is
  reserved for schema validation and has no effect today.
- `data/legacy/contract.json` and `data/app_schemas.json` are embedded for
  parity with the source SDK and for future validation work. They have no
  runtime consumer here.
- What the client does enforce is transport-level: JSON object and array
  decoding, a bounded response body (`Config.MaxResponseBytes` and
  `RequestOptions.MaxResponseBodyBytes`), and `*APIError` for every non-2xx
  response.

Model your own structs on the response schema each reference page names, and
treat unexpected fields as possible, because nothing rejects them.

## Golden Versus Silver

Golden routes under `client.<Namespace>.*` come from the published contract and
are the correct default path. Silver routes under `client.Silver.*` are kept
separate precisely because the published contract does not document them. When a
contract sync starts documenting a Silver route, the Silver twin is dropped
upstream so Golden owns it, and the generated surface here follows on the next
sync.

## Contract Sync Workflow

```bash
scripts/sync_from_source_sdk.sh ../incident-py-q
go generate ./...
go test ./...
```

The script copies the machine-readable artifacts from a local checkout of the
`incident-py-q` SDK, where the HAR classification and OpenAPI fetch tooling
lives. It copies no documentation. `go generate ./...` then rebuilds both the
wrappers and the SDK reference from the refreshed artifacts.
