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
- `INCIDENTIQ_AUTH_MODE` (optional, default `bearer`)
- `INCIDENTIQ_APP_HEADERS_JSON` (optional JSON object string for app-path calls)

`INCIDENTIQ_BASE_URL` may be either the tenant root, such as `https://your-tenant.incidentiq.com`, or an explicit API prefix such as `https://your-tenant.incidentiq.com/api/v1.0`. Bare tenant roots are normalized to `/api/v1.0` for Golden routes.

Future integration smoke tests should use:
- `INCIDENTIQ_TEST_BASE_URL`
- `INCIDENTIQ_TEST_API_TOKEN`
- `INCIDENTIQ_TEST_SITE_ID` (optional)
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

The SDK reads response bodies through a 4 MiB safety cap before JSON decoding
or non-2xx `APIError` construction. Use `Config.MaxResponseBodyBytes` for a
client-wide override or `RequestOptions.MaxResponseBodyBytes` for one request.
Oversized responses return `*incidentiq.ResponseSizeError`.

## Golden And Silver Helpers

Golden is the correct default SDK path. Golden methods are exposed directly on
the client:

```go
var statuses map[string]any
err := client.Tickets.GetTicketStatuses(ctx, incidentiq.RequestOptions{}, &statuses)
```

Silver is a separate namespace for quasi-supported API calls derived from live
site interaction HARs. The exported Go accessor is `client.Silver`:

```go
var status map[string]any
err := client.Silver.Tickets.GetTicketStatus(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"ticket_id": "ticket-guid"},
}, &status)
```
