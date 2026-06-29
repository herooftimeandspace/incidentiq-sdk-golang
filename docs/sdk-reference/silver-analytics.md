# `Silver.analytics` Namespace

Go client access: `client.Silver.Analytics`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_agent_current_stats`

- Go wrapper: `client.Silver.Analytics.GetAgentCurrentStats(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "analytics", "get_agent_current_stats", opts, out)`
- HTTP route: `GET /api/v1.0/analytics/agent-current-stats`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Analytics`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic. The April 22, 2026 resize HAR showed the upload plus a later `GET /img/...?...w=150&h=150`, but no separate persisted crop endpoint, so the SDK applies the avatar framing locally. It accepts common local raster image formats including JPG/JPEG, PNG, GIF, WEBP, and BMP. For non-square inputs it uses the largest centered square crop, then converts the result inside `client.Silver.Profiles.PostProfilePicture(...)` to PNG and downscales it until the uploaded PNG payload stays at or below 1 MB.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_agent_location_stats`

- Go wrapper: `client.Silver.Analytics.GetAgentLocationStats(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "analytics", "get_agent_location_stats", opts, out)`
- HTTP route: `GET /api/v1.0/analytics/agent-location-stats`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Analytics`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic. The April 22, 2026 resize HAR showed the upload plus a later `GET /img/...?...w=150&h=150`, but no separate persisted crop endpoint, so the SDK applies the avatar framing locally. For non-square inputs it uses the largest centered square crop, then converts the result inside `client.Silver.Profiles.PostProfilePicture(...)` to PNG and downscales it until the uploaded PNG payload stays at or below 1 MB.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_asset_summary_stats`

- Go wrapper: `client.Silver.Analytics.GetAssetSummaryStats(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "analytics", "get_asset_summary_stats", opts, out)`
- HTTP route: `GET /api/v1.0/analytics/asset/{asset_id}/summary-stats`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Analytics`.

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

### `get_requestor_summary_stats`

- Go wrapper: `client.Silver.Analytics.GetRequestorSummaryStats(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "analytics", "get_requestor_summary_stats", opts, out)`
- HTTP route: `GET /api/v1.0/analytics/requestor/{requestor_id}/summary-stats`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Analytics`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["requestor_id"]` | `requestor_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
