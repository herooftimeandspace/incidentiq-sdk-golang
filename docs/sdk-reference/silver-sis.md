# `Silver.sis` Namespace

Go client access: `client.Silver.Sis`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_classes_for__user`

- Go wrapper: `client.Silver.Sis.GetClassesForUser(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "sis", "get_classes_for__user", opts, out)`
- HTTP route: `GET /api/v1.0/sis/classes/for/user/{user_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Sis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["user_id"]` | `user_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
