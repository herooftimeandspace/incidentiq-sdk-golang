# `silver.tickets` Namespace

Go client access: `client.Silver.Tickets`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_ticket_activities`

- Go wrapper: `client.Silver.Tickets.GetTicketActivities(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "get_ticket_activities", opts, out)`
- HTTP route: `GET /api/v1.0/tickets/{ticket_id}/activities`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_ticket_kb_articles`

- Go wrapper: `client.Silver.Tickets.GetTicketKbArticles(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "get_ticket_kb_articles", opts, out)`
- HTTP route: `GET /api/v1.0/tickets/{ticket_id}/kb-articles`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_ticket_next_steps`

- Go wrapper: `client.Silver.Tickets.GetTicketNextSteps(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "get_ticket_next_steps", opts, out)`
- HTTP route: `GET /api/v1.0/tickets/{ticket_id}/next-steps`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_ticket_status`

- Go wrapper: `client.Silver.Tickets.GetTicketStatus(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "get_ticket_status", opts, out)`
- HTTP route: `GET /api/v1.0/tickets/{ticket_id}/status`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_wizards_site`

- Go wrapper: `client.Silver.Tickets.GetWizardsSite(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "get_wizards_site", opts, out)`
- HTTP route: `GET /api/v1.0/tickets/wizards/site/{site_id}`
- Observed in: `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["site_id"]` | `site_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_endpoint`

- Go wrapper: `client.Silver.Tickets.PostEndpoint(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "post_endpoint", opts, out)`
- HTTP route: `POST /api/v1.0/tickets`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$o"]` | `$o` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any` | Request body observed in HAR traffic. The SDK uses a single `json_body` payload because the Silver route carries a complex undocumented schema. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_ticket_timeline`

- Go wrapper: `client.Silver.Tickets.PostTicketTimeline(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets", "post_ticket_timeline", opts, out)`
- HTTP route: `POST /api/v1.0/tickets/{ticket_id}/timeline`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Tickets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `list_current_user_assigned_tickets`

- Source SDK helper `tickets.list_current_user_assigned_tickets` is not exposed by the generated Go wrapper surface.
- Backing routes: `POST /services/tickets/-/-/AssignedToMe_Unassigned`, `POST /services/tickets/-/-/AssignedToMe_Unassigned with configured Client header`, `POST /services/tickets/-/-/AssignedToMe_Unassigned with legacy o/d sort query`, `POST /services/tickets with AssignedToMe_Unassigned schema body`

List the current user's UI-style assigned/open ticket queue through the `AssignedToMe_Unassigned` services route.

This helper exists because tenant analytics summaries and saved-view routes can return zero rows for current-user ticket queues even when the Incident IQ web UI shows assigned work. The bundled Postman corpus includes the UI-observed `/services/tickets/-/-/AssignedToMe_Unassigned` route for open assigned tickets, so the SDK exposes a narrow read-only helper around that route instead of asking callers to construct a services URL by hand. The route uses POST for query semantics and some tenants only expose it to the UI-style `Client: WebBrowser` header, so the helper sends that header with page size, sort field, and sort direction parameters and does not send a mutation body. If the UI-shaped request returns 404, the helper retries once with the caller's configured client header for older tenants that accepted the pre-0.2.5 SDK request shape, then tries the Postman-observed legacy sort-query spelling that uses `$s`, `o`, and `d`. Some tenants, including WUSD as validated for issue #73, do not expose the direct queue route but do expose the same queue through `POST /services/tickets` with `{"Schema": "AssignedToMe_Unassigned"}`, so that schema body is the final read-only fallback.

#### Examples

```go
var payload map[string]any
err := client.Request(ctx, "POST", "/services/tickets/-/-/AssignedToMe_Unassigned", incidentiq.RequestOptions{
	Params: map[string]string{
		"$s": "100",
		"$o": "TicketPriority",
		"$d": "Descending",
	},
}, &payload)
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$s"]` | `$s` | `query` | `no` | `int` | Maximum number of ticket rows to return from the queue. |
| `Params["$o"]` | `$o` | `query` | `no` | `string` | Incident IQ ticket field used for ordering returned rows. |
| `Params["$d"]` | `$d` | `query` | `no` | `string` | Sort direction, either `Ascending` or `Descending`. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this manual helper has no typed response model.

---
