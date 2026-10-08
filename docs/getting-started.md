# Getting Started

## Requirements

- Go `1.26.4+` (the toolchain version in `go.mod`)

The module has no third-party dependencies.

## Install

```bash
go get github.com/herooftimeandspace/incidentiq-sdk-golang
```

```go
import incidentiq "github.com/herooftimeandspace/incidentiq-sdk-golang"
```

## Environment Variables

Runtime, read by `incidentiq.ConfigFromEnv(false)` and `incidentiq.NewClientFromEnv()`:
- `INCIDENTIQ_BASE_URL`
- `INCIDENTIQ_API_TOKEN`
- `INCIDENTIQ_SITE_ID` (optional)
- `INCIDENTIQ_AUTH_MODE` (optional, default `bearer`)
- `INCIDENTIQ_APP_HEADERS_JSON` (optional JSON object string for app-path calls)

`INCIDENTIQ_BASE_URL` may be either the tenant root, such as
`https://your-tenant.incidentiq.com`, or an explicit API prefix such as
`https://your-tenant.incidentiq.com/api/v1.0`. Bare tenant roots are normalized
to `/api/v1.0` for Golden routes.

Integration smoke tests, read by `incidentiq.ConfigFromEnv(true)`:
- `INCIDENTIQ_TEST_BASE_URL`
- `INCIDENTIQ_TEST_API_TOKEN`
- `INCIDENTIQ_TEST_SITE_ID` (optional)
- `INCIDENTIQ_TEST_AUTH_MODE` (optional)
- `INCIDENTIQ_TEST_APP_HEADERS_JSON` (optional JSON object string)

## Creating a Client

From explicit configuration:

```go
client, err := incidentiq.NewClient(incidentiq.Config{
	BaseURL:  "https://your-tenant.incidentiq.com",
	APIToken: os.Getenv("INCIDENTIQ_API_TOKEN"),
	SiteID:   "optional-site-id",
	Timeout:  30 * time.Second,
})
```

From the environment:

```go
client, err := incidentiq.NewClientFromEnv()
```

`Client` holds no resources of its own, so there is nothing to close. Supply
your own `*http.Client` through `Config.HTTPClient` when you need custom
transport behavior.

## Making a Call

```go
var ticket map[string]any
err := client.Tickets.GetTicket(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"ticketId": ticketID},
}, &ticket)
```

Cancellation and deadlines come from `ctx`. `RequestOptions.Timeout` bounds a
single call independently of the client-wide timeout.

## Handling Errors

```go
var apiErr *incidentiq.APIError
if errors.As(err, &apiErr) {
	log.Printf("incident iq returned %d: %s", apiErr.StatusCode, apiErr.Message)
}
```

- `*incidentiq.APIError` — a non-2xx response
- `*incidentiq.ValidationError` — a bad argument caught before the request
- `*incidentiq.ConfigurationError` — invalid client configuration
- `*incidentiq.ResponseTooLargeError` — the response exceeded the size limit

## Silver Namespace

Golden methods stay on `client.<Namespace>.<Method>`. Quasi-supported routes
derived from live site interaction HARs are exposed under `client.Silver`, with
app routes nested one level deeper:

```go
err := client.Silver.AppRegistry.GetApp(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"app_key": appKey},
}, &app)

err = client.Silver.Apps.GoogleDeviceData.GetSyncHistory(ctx, incidentiq.RequestOptions{}, &history)
```

App routes usually need tenant app headers. Set them once on `Config.AppHeaders`
(or `INCIDENTIQ_APP_HEADERS_JSON`), or per call in `opts.Headers`.

A few Silver routes have hand-written typed helpers instead of generic wrappers:

```go
err := client.Silver.Users.SetUserRooms(ctx, userID, []string{roomID}, incidentiq.RequestOptions{}, nil)
```

See [user room associations](user-room-associations.md).
