# `Silver.api` Namespace

Go client access: `client.Silver.Api`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_search_v2_expression`

- Go wrapper: `client.Silver.Api.GetSearchV2Expression(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "api", "get_search_v2_expression", opts, out)`
- HTTP route: `GET /api/search/v2/expression`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented GET route for `client.Silver.Api`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["s"]` | `s` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_search_v2_recent_user`

- Go wrapper: `client.Silver.Api.PostSearchV2RecentUser(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "api", "post_search_v2_recent_user", opts, out)`
- HTTP route: `POST /api/search/v2/recent/user`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Api`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["searchExpressionAsString"]` | `searchExpressionAsString` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["source"]` | `source` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["sourceUrl"]` | `sourceUrl` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["sourceUserAgent"]` | `sourceUserAgent` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["userId"]` | `userId` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
