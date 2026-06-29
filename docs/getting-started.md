# Getting Started

## Requirements

- Go version declared in `go.mod`

## Install

```bash
go get github.com/herooftimeandspace/incidentiq-sdk-golang
```

## Environment Variables

Runtime:
- `INCIDENTIQ_BASE_URL`
- `INCIDENTIQ_API_TOKEN`
- `INCIDENTIQ_SITE_ID` (optional)
- `INCIDENTIQ_CLIENT_HEADER` (optional, default `ApiClient`)
- `INCIDENTIQ_AUTH_MODE` (optional, default `bearer`)
- `INCIDENTIQ_APP_HEADERS_JSON` (optional JSON object string for app-path calls)

`INCIDENTIQ_BASE_URL` may be either the tenant root, such as `https://your-tenant.incidentiq.com`, or an explicit API prefix such as `https://your-tenant.incidentiq.com/api/v1.0`. Bare tenant roots are normalized to `/api/v1.0` for Golden routes.

Future integration smoke tests should use:
- `INCIDENTIQ_TEST_BASE_URL`
- `INCIDENTIQ_TEST_API_TOKEN`
- `INCIDENTIQ_TEST_SITE_ID` (optional)
- `INCIDENTIQ_TEST_CLIENT_HEADER` (optional)
- `INCIDENTIQ_TEST_AUTH_MODE` (optional)
- `INCIDENTIQ_TEST_APP_HEADERS_JSON` (optional JSON object string)

## Client From Env

```go
client, err := incidentiq.NewClientFromEnv()
if err != nil {
	return err
}
```

## Low-Level Request

```go
var users map[string]any
err := client.Request(ctx, "GET", "/users/{UserId}", incidentiq.RequestOptions{
	PathParams: map[string]any{
		"UserId": "00000000-0000-0000-0000-000000000000",
	},
}, &users)
```

## Golden And Silver Helpers

Golden methods are available through `RequestGolden` while typed wrappers are still pending:

```go
var statuses map[string]any
err := client.RequestGolden(ctx, "tickets", "get_ticket_statuses", incidentiq.RequestOptions{}, &statuses)
```

Silver routes are available through `RequestSilver`:

```go
var assigned map[string]any
err := client.RequestSilver(ctx, "tickets", "list_current_user_assigned_tickets", incidentiq.RequestOptions{}, &assigned)
```
