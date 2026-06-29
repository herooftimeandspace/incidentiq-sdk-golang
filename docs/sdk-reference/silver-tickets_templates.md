# `silver.tickets_templates` Namespace

Go client access: `client.Silver.TicketsTemplates`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `post_endpoint`

- Go wrapper: `client.Silver.TicketsTemplates.PostEndpoint(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "tickets_templates", "post_endpoint", opts, out)`
- HTTP route: `POST /api/v1.0/tickets-templates`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.TicketsTemplates`.

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
