# incidentiq-sdk-golang

| Main | Staging | Dev | License |
| --- | --- | --- | --- |
| [![Main coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-coverage%2Fmain%2Fcoverage.json)](https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/main/coverage.json) | [![Staging coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-coverage%2Fstaging%2Fcoverage.json)](https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/staging/coverage.json) | [![Dev coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-coverage%2Fdev%2Fcoverage.json)](https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/dev/coverage.json) | [![License repo](https://img.shields.io/github/license/herooftimeandspace/incidentiq-sdk-golang?label=license%20repo)](LICENSE) |
| [![Main unit](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-status%2Fmain%2Funit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Amain) | [![Staging unit](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-status%2Fstaging%2Funit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Astaging) | [![Dev unit](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-status%2Fdev%2Funit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Adev) |  |
| [![Main integration](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-status%2Fmain%2Fintegration.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/integration.yml?query=branch%3Amain) | [![Staging integration](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Fherooftimeandspace%2Fincidentiq-sdk-golang%2Fbadges%2Fbranch-status%2Fstaging%2Fintegration.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/integration.yml?query=branch%3Astaging) |  |  |

Contract-driven Incident IQ SDK for Go (module: `github.com/herooftimeandspace/incidentiq-sdk-golang`, package: `incidentiq`).

Coverage and phase-status badges are published by CI to a dedicated `badges` branch so protected branches (`main`, `staging`, `dev`) never require bot commits for badge refreshes.

The module ships:
- one `*incidentiq.Client` with a `context.Context`-aware request path
- generated wrappers for every operation in the bundled Incident IQ contracts
- the Golden surface (`client.<Namespace>.<Method>`) for supported, contract-documented calls
- the Silver surface (`client.Silver.<Namespace>.<Method>`) for quasi-supported routes derived from live site interaction HARs
- deprecated forwarders for every pre-migration method name, carrying `// Deprecated:` comments
- generated SDK reference docs, a static docs site builder, and CI workflows

## Requirements

- Go `1.26.4+` (the toolchain version in `go.mod`)

No third-party dependencies: the module uses the standard library only.

## Install

```bash
go get github.com/herooftimeandspace/incidentiq-sdk-golang
```

```go
import incidentiq "github.com/herooftimeandspace/incidentiq-sdk-golang"
```

## Authentication and Tenant URL

Default auth mode is bearer token:

```text
Authorization: Bearer <token>
```

Each client requires a tenant-specific base URL. You may pass either the tenant root
(`https://your-tenant.incidentiq.com`) or an explicit API prefix such as
`https://your-tenant.incidentiq.com/api/v1.0`. Bare tenant roots are normalized to
`/api/v1.0`. Golden contract paths are tenant-absolute (`/api/v1.0/...`), and Silver
routes that include an absolute tenant path such as `/api/v1.0/...`, `/services/...`,
`/apps/...`, or `/pub/...` are sent from the tenant origin so they do not accidentally
inherit the base URL prefix twice.

Runtime environment variables, read by `incidentiq.ConfigFromEnv(false)` and
`incidentiq.NewClientFromEnv()`:
- `INCIDENTIQ_BASE_URL` (required unless set on `Config`)
- `INCIDENTIQ_API_TOKEN` (required unless set on `Config`)
- `INCIDENTIQ_SITE_ID` (optional)
- `INCIDENTIQ_AUTH_MODE` (optional, default `bearer`, supported: `bearer`, `raw`)
- `INCIDENTIQ_APP_HEADERS_JSON` (optional JSON object string for app-path calls)

Security hardening rules:
- `INCIDENTIQ_BASE_URL` must use `https`
- base URLs with embedded credentials, query strings, or fragments are rejected
- `site_id` and app header keys/values cannot contain CR/LF characters
- timeout and retry tuning values must stay within safe positive/non-negative bounds

Integration/smoke environment variables, read by `incidentiq.ConfigFromEnv(true)`:
- `INCIDENTIQ_TEST_BASE_URL` (required for integration tests)
- `INCIDENTIQ_TEST_API_TOKEN` (required for integration tests)
- `INCIDENTIQ_TEST_SITE_ID` (optional)
- `INCIDENTIQ_TEST_AUTH_MODE` (optional, default `bearer`)
- `INCIDENTIQ_TEST_APP_HEADERS_JSON` (optional JSON object string for app-path integration calls)

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	incidentiq "github.com/herooftimeandspace/incidentiq-sdk-golang"
)

func main() {
	client, err := incidentiq.NewClient(incidentiq.Config{
		BaseURL:  "https://your-tenant.incidentiq.com",
		APIToken: "your-token",
	})
	if err != nil {
		log.Fatal(err)
	}

	var users map[string]any
	if err := client.Users.ListUsersSimple(context.Background(), incidentiq.RequestOptions{}, &users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users["Items"])
}
```

`incidentiq.NewClientFromEnv()` builds the same client from `INCIDENTIQ_*`.

### One Call Shape

Every generated method has the same signature:

```go
func (s *GoldenAssetsService) GetAssetById(ctx context.Context, opts RequestOptions, out any) error
```

- `opts` carries the request: `PathParams`, `Params` (query), `Headers`, `JSON` or
  `Body` with `ContentType`, plus per-call `Timeout` and size limits.
- `out` receives the decoded JSON response. Pass a pointer to your own struct,
  a `*map[string]any`, or `nil` to discard the body.
- Non-2xx responses return `*incidentiq.APIError`. Configuration and argument
  problems return `*incidentiq.ConfigurationError` or `*incidentiq.ValidationError`.

The SDK does not generate response structs. Model the types you care about on
the response schemas named in the [SDK reference](docs/sdk-reference/index.md).

### Path and Query Parameters

```go
var asset map[string]any
err := client.Assets.GetAssetById(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"assetId": "00000000-0000-0000-0000-000000000000"},
}, &asset)
```

### Request Bodies

```go
var created map[string]any
err := client.Tickets.CreateTicket(ctx, incidentiq.RequestOptions{
	JSON: map[string]any{"Subject": "Broken screen"},
}, &created)
```

### Low-Level Request API

For a route the bundled contracts do not cover:

```go
err := client.Request(ctx, "GET", "/api/v1.0/users/{UserId}", incidentiq.RequestOptions{
	PathParams: map[string]any{"UserId": "00000000-0000-0000-0000-000000000000"},
}, &user)
```

`client.RequestGolden(ctx, namespace, name, opts, out)` and
`client.RequestSilver(...)` resolve a route from the bundled inventory by its
contract namespace and operation name.

### Silver and App Routes

```go
err := client.Silver.AppRegistry.GetApp(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"app_key": appKey},
}, &app)
err = client.Silver.Apps.MicrosoftIntune.GetSyncStatusLast(ctx, incidentiq.RequestOptions{}, &status)
```

App routes usually need tenant app headers. Set them once on `Config.AppHeaders`
(or `INCIDENTIQ_APP_HEADERS_JSON`), or per call in `opts.Headers`.

Silver `POST` routes observed from the browser sometimes reject the default
`Client: ApiClient` header. The client retries those once without the header.
Set `opts.OmitClientHeader` to skip it from the start.

### Typed Silver Helpers

A few Silver routes have hand-written helpers instead of generic wrappers,
because they validate arguments and own their request body:

```go
err := client.Silver.Users.SetUserRooms(ctx, userID, []string{roomID}, incidentiq.RequestOptions{}, nil)
```

See [`client.Silver.Users`](docs/sdk-reference/silver-users.md).

## Response Handling

- Successful JSON objects and arrays are decoded with `encoding/json` into `out`.
- The SDK performs no response-schema validation. `Config.ValidateResponses` is
  reserved for that behavior and currently has no effect.
- `Config.MaxResponseBytes` (and per-call `RequestOptions.MaxResponseBodyBytes`)
  bound the response body; exceeding the limit returns
  `*incidentiq.ResponseTooLargeError`.
- Retries cover idempotent methods (`GET`, `HEAD`, `OPTIONS`, `DELETE`, `PUT`)
  on `408`, `429`, `500`, `502`, `503`, and `504`, with exponential backoff.

## Development Commands

```bash
go generate ./...
go vet ./...
go test -covermode=atomic -coverprofile=coverage.out ./...
go tool cover -func=coverage.out -o coverage-summary.txt
go run scripts/build_badge_json.go coverage --coverage-file coverage.out --label "coverage local" --minimum 95.0 --output coverage-badge.json
go run scripts/build_docs_site.go
```

Set `GOCACHE="$(pwd)/.gocache"` and `GOMODCACHE="$(pwd)/.gomodcache"` to keep
build caches inside the checkout.

Branch targets map to the GitHub Actions gates:
- `dev`: `unit` (vet, tests, coverage ratchet)
- `staging`: `unit` plus `integration` against a live tenant
- `main`: `unit`, `integration`, `docs-build`, and `release-prep`

## Generated Code

Two generators read the bundled contract artifacts, and `go generate ./...` runs both:

| Generator | Output |
| --- | --- |
| `scripts/generate_wrappers.go` | `generated_wrappers.go` |
| `scripts/generate_sdk_reference.go` | `docs/sdk-reference/*.md` |

Both are checked in, and `TestGeneratedReferenceMatchesWrappers` fails if one is
regenerated without the other.

## Contract Artifacts

The contracts are shared with the `herooftimeandspace/incident-py-q` SDK and are
embedded in this module:

- `data/openapi/openapi-spec.json` (Golden contract, primary)
- `data/openapi/metadata.json` (sync provenance)
- `data/source_manifest.json` (source manifest)
- `data/silver_inventory.json` (Silver inventory: HAR-derived + migrated routes)
- `data/legacy/contract.json` (schemas for routes migrated off Golden)
- `data/legacy/aliases.json` (deprecated method-name aliases)
- `data/app_schemas.json` (HAR-derived app-path schemas)
- `data/typed_silver_methods.json` (routes with hand-written typed helpers)
- `testdata/contract/*_sdk_inventory.json` (golden-surface drift snapshots)

Refresh them from a local checkout of the source SDK and regenerate:

```bash
scripts/sync_from_source_sdk.sh ../incident-py-q
go generate ./...
go test ./...
```

The script copies artifacts only. Documentation in this repository is written
for Go and is never overwritten by a sync. See [docs/go-parity.md](docs/go-parity.md).

## Upgrading from the pre-OpenAPI SDK

The Golden contract moved to the published OpenAPI document, which renamed most
operations and stopped documenting 64 routes. Those routes moved to
`client.Silver.*` rather than being dropped, and every previous method name
still compiles as a deprecated forwarder.

Go has no runtime deprecation warning, so each forwarder carries a
`// Deprecated:` comment that `go doc`, editors, and `staticcheck` surface.

One legacy name could not be preserved: `client.Tickets.AssignTicket` now reaches
a different operation than it used to. `incidentiq.LegacyAliasConflicts()` reports
it at runtime, and it compiles unchanged while doing something different — see
[docs/migration-openapi.md](docs/migration-openapi.md) for the full mapping and
for the Silver methods that were removed outright.

## Versioning and Stability

- The module follows semantic versioning, published as `vX.Y.Z` Git tags.
- The generated SDK surface is semver-significant and protected by golden tests.
- Promotion into `main` requires exactly one release label: `semver:patch`, `semver:minor`, or `semver:major`.
- Promotion workflows propagate the source PR's semver label when one is present and otherwise default the promotion PR to `semver:patch`.

## Documentation

- [SDK reference](docs/sdk-reference/index.md), generated from the bundled contracts
- [Getting started](docs/getting-started.md)
- [SDK usage](docs/sdk-usage.md)
- [Go parity notes](docs/go-parity.md)
- `go doc github.com/herooftimeandspace/incidentiq-sdk-golang`

Build the published static site:

```bash
go run scripts/build_docs_site.go
```
