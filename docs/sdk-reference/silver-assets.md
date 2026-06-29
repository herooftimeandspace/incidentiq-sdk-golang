# `Silver.assets` Namespace

Go client access: `client.Silver.Assets`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_asset_by_serial`

- Go wrapper: `client.Silver.Assets.GetAssetBySerial(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "get_asset_by_serial", opts, out)`
- HTTP route: `GET /assets/serial/{serial}`
- Observed in: `synthetic_required_route`

Silver path for asset lookup by serial number.

This route is treated as Silver because bundled Stoplight contracts do not define it. The SDK exposes it separately so Golden Stoplight methods remain the authoritative contract whenever they exist.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["serial"]` | `serial` | `path` | `yes` | `string` | Serial number path segment. This Silver route is added explicitly because it is known to exist even when the HAR does not capture it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_asset_files`

- Go wrapper: `client.Silver.Assets.GetAssetFiles(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "get_asset_files", opts, out)`
- HTTP route: `GET /api/v1.0/assets/{asset_id}/files`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Assets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["asset_id"]` | `asset_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_asset_verifications`

- Go wrapper: `client.Silver.Assets.GetAssetVerifications(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "get_asset_verifications", opts, out)`
- HTTP route: `GET /api/v1.0/assets/{asset_id}/verifications`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Assets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["asset_id"]` | `asset_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_stats_locations`

- Go wrapper: `client.Silver.Assets.GetStatsLocations(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "get_stats_locations", opts, out)`
- HTTP route: `GET /api/v1.0/assets/stats/locations`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Assets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_type`

- Go wrapper: `client.Silver.Assets.GetType(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "get_type", opts, out)`
- HTTP route: `GET /api/v1.0/assets/types/{type_id}`
- Observed in: `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Assets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["type_id"]` | `type_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_checkouts_transactions_query_get`

- Go wrapper: `client.Silver.Assets.PostCheckoutsTransactionsQueryGet(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "assets", "post_checkouts_transactions_query_get", opts, out)`
- HTTP route: `POST /api/v1.0/assets/checkouts/transactions/query/get`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Assets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any` | Request body observed in HAR traffic. The SDK uses a single `json_body` payload because the Silver route carries a complex undocumented schema. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
