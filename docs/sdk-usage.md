# SDK Usage

The SDK surface is split into Golden and Silver paths:
- Golden: bundled Stoplight controller contract inventory exposed directly on `client.<Namespace>.<Method>`.
- Silver: HAR-observed undocumented route inventory exposed under `client.Silver.<Namespace>.<Method>`.

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
