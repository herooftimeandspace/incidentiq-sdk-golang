# `Silver.search` Namespace

Go client access: `client.Silver.Search`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `post_scan`

- Go wrapper: `client.Silver.Search.PostScan(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "search", "post_scan", opts, out)`
- HTTP route: `POST /api/v1.0/search/scan`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Search`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `Limit` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `Query` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
