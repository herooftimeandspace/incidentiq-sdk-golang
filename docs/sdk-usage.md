# SDK Usage

The SDK surface distinguishes the Golden path from the Silver namespace:
- Golden: the golden SDK path and correct default API surface, exposed directly on `client.<Namespace>.<Method>`.
- Silver: quasi-supported API calls derived from live site interaction HARs. Silver methods are exposed under `client.Silver.<Namespace>.<Method>`, with app routes nested under `client.Silver.Apps.<AppNamespace>.<Method>`.

Full generated route documentation lives under the SDK reference pages. The generated Go wrappers are reproduced from the bundled inventory snapshots.

## Golden Wrapper Pattern

```go
var payload map[string]any
err := client.Tickets.GetTicketStatuses(ctx, incidentiq.RequestOptions{}, &payload)
```

## Silver Wrapper Pattern

```go
var payload map[string]any
err := client.Silver.Tickets.GetTicketStatus(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"ticket_id": "ticket-guid"},
}, &payload)
```

App-specific Silver routes use the nested app namespace:

```go
err := client.Silver.Apps.GoogleDeviceData.GetStatusLast(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"google_device_data_key": "site-app-key"},
}, &payload)
```

All requests send `Client: ApiClient` by default. Silver wrappers and
`RequestSilver` retry once without the SDK-provided `Client` header if the
ApiClient-shaped request fails, because Silver routes are reverse-engineered
from live site traffic and may not comply with the documented Postman header
contract.

## Request Options

Use `RequestOptions` for path parameters, query parameters, JSON bodies, headers, and per-request timeout:

```go
err := client.Request(ctx, "GET", "/users/{UserId}", incidentiq.RequestOptions{
	PathParams: map[string]any{"UserId": "00000000-0000-0000-0000-000000000000"},
	Params:     map[string]string{"$s": "100"},
}, &payload)
```

## Low-Level Request API

```go
err := client.Request(ctx, "POST", "/services/tickets/-/-/AssignedToMe_Unassigned", incidentiq.RequestOptions{
	JSON: map[string]any{"OnlyOpen": true},
}, &payload)
```

Use `RequestGolden` and `RequestSilver` when a caller needs to resolve a route dynamically by inventory namespace and method name instead of calling a generated Go method.
