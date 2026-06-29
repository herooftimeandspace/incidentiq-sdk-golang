# incidentiq-sdk-golang

[![dev unit](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/dev/unit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Adev)
[![dev coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/dev/coverage.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Adev)
[![staging unit](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/staging/unit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Astaging)
[![staging coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/staging/coverage.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Astaging)
[![main unit](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/main/unit.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Amain)
[![main coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-coverage/main/coverage.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/quality.yml?query=branch%3Amain)
[![staging integration](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/staging/integration.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/integration.yml?query=branch%3Astaging)
[![main integration](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/main/integration.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/integration.yml?query=branch%3Amain)
[![main docs](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/herooftimeandspace/incidentiq-sdk-golang/badges/branch-status/main/docs.json)](https://github.com/herooftimeandspace/incidentiq-sdk-golang/actions/workflows/docs.yml?query=branch%3Amain)

`incidentiq-sdk-golang` is the Go SDK companion to
[`herooftimeandspace/incident-py-q`](https://github.com/herooftimeandspace/incident-py-q).

The project goal is functional parity with the source SDK:

- same Incident IQ runtime environment variables
- same tenant URL normalization rules
- same bearer/raw authorization behavior
- same auth and optional `SiteId` header behavior
- same Golden Stoplight contract artifacts
- same Silver HAR-derived inventory artifacts
- same retry policy for idempotent requests
- generated Go wrappers for every bundled Golden and Silver inventory entry
- same low-level request escape hatch for routes that need direct access
- same Markdown documentation set, kept in this repo so Go-specific updates can be made beside the shared source material

## Install

```bash
go get github.com/herooftimeandspace/incidentiq-sdk-golang
```

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

	var payload map[string]any
	err = client.Request(context.Background(), "GET", "/users/{UserId}", incidentiq.RequestOptions{
		PathParams: map[string]any{
			"UserId": "00000000-0000-0000-0000-000000000000",
		},
	}, &payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(payload)
}
```

## Environment Configuration

`NewClientFromEnv` uses the same runtime variables as `incident-py-q`:

- `INCIDENTIQ_BASE_URL`
- `INCIDENTIQ_API_TOKEN`
- `INCIDENTIQ_SITE_ID`
- `INCIDENTIQ_AUTH_MODE`
- `INCIDENTIQ_APP_HEADERS_JSON`

Tenant roots such as `https://your-tenant.incidentiq.com` are normalized to
`https://your-tenant.incidentiq.com/api/v1.0`. Absolute tenant paths beginning
with `/api/`, `/services/`, `/apps/`, `/img/`, or `/s/` are sent from the tenant
origin so Silver routes do not accidentally inherit the Golden `/api/v1.0`
prefix.

The client caps response bodies at 4 MiB by default before decoding JSON or
building an `APIError`. Set `Config.MaxResponseBodyBytes` to use a different
client-wide limit, or set `RequestOptions.MaxResponseBodyBytes` for a single
request. When a response exceeds the active limit, `Request` returns
`*incidentiq.ResponseSizeError` instead of buffering the full response.

## Golden And Silver Calls

The SDK embeds the same inventories generated by `incident-py-q`. Golden refers
to the golden SDK path and is the correct default API surface for supported SDK
calls. Golden methods are exposed directly on
`client.<Namespace>.<Method>`.

Silver is a separate namespace for quasi-supported API calls derived from
live site interaction HARs. Silver methods are exposed under
`client.Silver.<Namespace>.<Method>` and app-specific routes are exposed under
`client.Silver.Apps.<AppNamespace>.<Method>`.

All requests send `Client: ApiClient` by default. Silver requests first try that
same header shape; if the Silver route rejects it, the SDK retries once without
the SDK-provided `Client` header because HAR-derived routes may not follow the
documented Postman requirement. Set `RequestOptions.OmitClientHeader` only for
HAR-validated Silver calls that must intentionally match browser traffic without
the SDK-provided `Client` header.

```go
var tickets map[string]any
err := client.Tickets.GetTicketStatuses(ctx, incidentiq.RequestOptions{}, &tickets)
```

```go
var status map[string]any
err := client.Silver.Tickets.GetTicketStatus(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"ticket_id": "ticket-guid"},
}, &status)
```

Use `Request` directly when you already know the route or need a temporary escape
hatch:

```go
err := client.Request(ctx, "POST", "/services/tickets/-/-/AssignedToMe_Unassigned", incidentiq.RequestOptions{
	JSON:                 map[string]any{"OnlyOpen": true},
	MaxResponseBodyBytes: 8 * 1024 * 1024,
}, &queue)
```

## Sync Contract With incident-py-q

`incident-py-q` remains the source repo for the shared docs and contract
artifacts until the Go SDK has its own generator parity.

Refresh this repo from a sibling checkout:

```bash
./scripts/sync_from_source_sdk.sh ../incident-py-q
```

That command copies:

- root Markdown files
- `docs/**/*.md`
- `LICENSE`
- Golden Stoplight controller JSON
- Postman collection JSON
- source manifest JSON
- app schema JSON
- Silver inventory JSON
- Golden, Silver, and merged SDK inventory snapshots

After every sync, run:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

## Current Implementation Scope

This Go SDK is usable for generated Golden wrappers, generated Silver namespace
wrappers, inventory-backed calls, and low-level route calls.

The current parity surface is protected by tests for:

- config and environment variable names
- HTTPS-only base URL normalization
- auth and optional `SiteId` headers
- path parameter escaping
- Golden-prefix versus tenant-root URL construction
- idempotent retry handling for retryable status codes
- embedded Golden and Silver inventory loading
- generated wrapper inventory coverage

## Development

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

Run the native Go coverage workflow locally:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test -covermode=atomic -coverprofile=coverage.out ./...
go tool cover -func=coverage.out -o coverage-summary.txt
go tool cover -html=coverage.out -o coverage.html
go run scripts/build_badge_json.go coverage --coverage-file coverage.out --label "coverage local" --minimum 95.0 --output coverage-badge.json
```

The repo uses only the Go standard library for the runtime client.

Regenerate wrappers after refreshing inventory snapshots:

```bash
go generate ./...
```

Build the static documentation site that backs the `docs-build` workflow:

```bash
go run scripts/build_docs_site.go
```

## Promotion And Release Automation

The Go SDK follows the same branch promotion shape as `incident-py-q`:

- `dev` is the normal integration branch.
- successful `quality` push checks on `dev` create or refresh a `dev -> staging`
  promotion PR.
- successful `quality` push checks on `staging` create or refresh a
  `promote/staging-to-main -> main` promotion PR.
- every promotion PR carries exactly one release label:
  `semver:patch`, `semver:minor`, or `semver:major`.
- merges to `main` create a GitHub Release and source archive tagged as the next
  semantic version.
- `main` runs the docs workflow and publishes the generated Markdown site to
  GitHub Pages.
- Dependabot checks both Go modules and GitHub Actions weekly.

The automation uses the repository `GITHUB_TOKEN`. It does not require a
personal access token by default. Repository rules must allow GitHub
Actions-authored branches and pull requests for fully automatic promotion.
