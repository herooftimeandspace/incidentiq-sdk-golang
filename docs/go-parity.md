# Go SDK Parity Notes

This repo is the Go companion to `herooftimeandspace/incident-py-q`.

Parity is defined at the **contract and wire** level, not at the documentation
level. The two SDKs share the bundled contract artifacts, so they call the same
routes with the same environment variables, authentication headers, URL
normalization rules, and retry policy. Everything above that line — the API
shape, the documentation, and the tooling — is this repository's own and is
written for Go.

When the source repo changes a shared contract, run
`scripts/sync_from_source_sdk.sh`, then `go generate ./...`, and update the Go
runtime until the tests prove parity again. The sync script copies
machine-readable artifacts only; it never overwrites Markdown.

This page records where the Go SDK deliberately differs.

## Shared Runtime Behavior

The Go client currently matches the source SDK for the following runtime rules:

- `INCIDENTIQ_*` environment variable names
- bearer and raw authorization modes
- optional `SiteId` header
- HTTPS-only base URL validation
- tenant root normalization to `/api/v1.0`
- tenant-root handling for `/api/`, `/services/`, `/apps/`, `/img/`, `/s/`, and `/pub/`
- JSON object and array response decoding
- retry behavior for idempotent methods and retryable HTTP statuses
- Golden wrappers directly on `client.<Namespace>.<Method>` as the correct default SDK path
- Silver wrappers under `client.Silver.<Namespace>.<Method>` for quasi-supported API calls derived from live site interaction HARs, with app routes nested under `client.Silver.Apps.<AppNamespace>.<Method>`

## Contract Artifacts

The Go module embeds the same contract artifacts as the source SDK. The list is
in [schema-validation.md](schema-validation.md#bundled-artifacts).

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

## OpenAPI Contract Migration

The Golden contract source moved from the Stoplight controller sync and the
APIHub Postman collection to the published Incident IQ OpenAPI 3.0 document.
`CHANGELOG.md` describes that migration in the source SDK's terms. The Go SDK
mirrors it with these differences:

- **Deprecation signalling.** Go has no runtime deprecation warning, so every
  legacy name is emitted as a forwarding method carrying a `// Deprecated:` doc
  comment. `go doc`, editors, and `staticcheck` surface it; nothing is logged at
  runtime.
- **Alias coverage.** All 114 aliases from `data/legacy/aliases.json` are
  generated: 50 renamed Golden methods and 64 routes that moved to Silver.
  The alias layer covers pre-migration **Golden** names only. Silver method
  names are not aliased; see *Removed Silver methods* below. The
  six legacy namespaces the new contract dropped (`alerts`, `forms`,
  `manufacturers`, `notifications`, `parts`, `purchaseorders`) still exist on
  `Client` and hold only deprecated forwarders into `client.Silver`.
- **Aliases never shadow generated methods.** The generator emits contract
  operations first and skips any alias whose exported name is already taken.
- **Conflict discovery.** `LegacyAliasConflicts()` is the Go equivalent of
  `incident_py_q.legacy_alias_conflicts()`.
- **No response validation.** `data/legacy/contract.json` and
  `data/app_schemas.json` are embedded for parity, but the Go SDK unmarshals
  into `out any` and performs no response-schema validation, so those bundles
  currently have no functional consumer here. `Config.ValidateResponses` is
  reserved for that behavior and is inert today.
- **No logging.** The source SDK logs through the standard `logging` module with
  header redaction. The Go SDK emits nothing and installs no logger; wrap
  `Config.HTTPClient` with your own `RoundTripper` to observe requests, and do
  your own redaction.
- **No configurable client header.** The source SDK reads
  `INCIDENTIQ_CLIENT_HEADER`. The Go client always sends the constant
  `Client: ApiClient`, and a caller overrides it per request through
  `RequestOptions.Headers` or suppresses it with
  `RequestOptions.OmitClientHeader`.
- **No typed response models or pagination helper.** The source SDK returns
  Pydantic models and offers `iter_pages(...)`. Every Go wrapper decodes into
  the caller's `out` value, and paging is driven by the route's own query
  parameters.

### `Tickets.AssignTicket` changed meaning

This is the one name that could not be aliased, because the new contract gives
it to a different operation:

| | |
|---|---|
| Before | `POST /tickets/{TicketId}/sla` — assign an **SLA** |
| Now | `POST /api/v1.0/tickets/{ticketId}/assign` — assign the **ticket** |
| Previous behavior moved to | `Tickets.AssignTicketSla` |

This compiles unchanged and silently does something different at runtime.
Callers of `client.Tickets.AssignTicket` that meant to assign an SLA must move
to `client.Tickets.AssignTicketSla`.

### Removed Silver methods

The alias layer does not cover Silver. 45 `client.Silver.<Namespace>.<Method>`
methods that existed before the migration are gone, and every one of them is a
compile error for existing callers. 40 are still reachable after a hand edit,
because the published contract now documents the route and it moved onto the
Golden surface; 5 are gone outright. That removal happened upstream in
`incident-py-q`, not here. The full table is in
[migration-openapi.md](migration-openapi.md#removed-silver-methods).

## Promotion Branch Shape

Promotion mechanics are not a Go-versus-Python parity topic. They are documented
in [project-docs.md](project-docs.md#promotion-automation).

## Documentation Is Not Synced

The documentation in this repository used to be copied verbatim from the source
SDK, which meant the README, contributor docs, and SDK reference described a
Python package with sync and async clients, `.raw(...)` calls, and `Python Arg`
parameter tables — none of which exist here. Go-specific edits were silently
reverted by the next sync.

That is no longer the case:

- `scripts/sync_from_source_sdk.sh` copies contract artifacts only.
- `docs/sdk-reference/` is generated from those artifacts by
  `scripts/generate_sdk_reference.go`, in Go terms.
- `TestGeneratedReferenceIsGoShaped` fails if the Python call shape reappears in
  the reference pages, and `TestGeneratedReferenceMatchesWrappers` fails if the
  reference and the wrappers disagree about the method surface.

Write Go-specific notes wherever they belong. Nothing is overwritten by a sync.
