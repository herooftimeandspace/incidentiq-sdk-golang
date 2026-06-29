# SDK Usage

The SDK surface is split into Golden and Silver paths:
- Golden: bundled Stoplight controller contract inventory.
- Silver: HAR-observed undocumented route inventory.

Full generated route documentation lives under the SDK reference pages. Typed Go wrappers for every inventory entry are tracked separately; until those wrappers exist, use the inventory-backed request helpers.

## Golden Inventory Pattern

```go
var payload map[string]any
err := client.RequestGolden(ctx, "tickets", "get_ticket_statuses", incidentiq.RequestOptions{}, &payload)
```

## Silver Inventory Pattern

```go
var payload map[string]any
err := client.RequestSilver(ctx, "tickets", "list_current_user_assigned_tickets", incidentiq.RequestOptions{
	Params: map[string]string{"$s": "100"},
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

`list_current_user_assigned_tickets` uses the UI-observed read-only assigned/open queue route. It is useful when analytics or saved-view routes return zero rows while the web UI still shows current-user assigned work.
