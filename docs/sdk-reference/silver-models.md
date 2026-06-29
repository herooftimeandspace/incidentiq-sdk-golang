# `silver.models` Namespace

Go client access: `client.Silver.Models`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_all`

- Go wrapper: `client.Silver.Models.GetAll(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "get_all", opts, out)`
- HTTP route: `GET /api/v1.0/models/all`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$filter"]` | `$filter` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$o"]` | `$o` | `query` | `yes` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps_aeries_sis`

- Go wrapper: `client.Silver.Models.GetAppsAeriesSis(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "get_apps_aeries_sis", opts, out)`
- HTTP route: `GET /api/v1.0/models/apps/aeriesSis`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps_google_device_data`

- Go wrapper: `client.Silver.Models.GetAppsGoogleDeviceData(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "get_apps_google_device_data", opts, out)`
- HTTP route: `GET /api/v1.0/models/apps/googleDeviceData`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps_microsoft_intune`

- Go wrapper: `client.Silver.Models.GetAppsMicrosoftIntune(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "get_apps_microsoft_intune", opts, out)`
- HTTP route: `GET /api/v1.0/models/apps/microsoftIntune`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps_subtickets_for_i_t`

- Go wrapper: `client.Silver.Models.GetAppsSubticketsForIT(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "get_apps_subtickets_for_i_t", opts, out)`
- HTTP route: `GET /api/v1.0/models/apps/subticketsForIT`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_available_to_site`

- Go wrapper: `client.Silver.Models.PostAvailableToSite(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "post_available_to_site", opts, out)`
- HTTP route: `POST /api/v1.0/models/available/to/site`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any` | Request body observed in HAR traffic. The SDK uses a single `json_body` payload because the Silver route carries a complex undocumented schema. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_endpoint`

- Go wrapper: `client.Silver.Models.PostEndpoint(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "models", "post_endpoint", opts, out)`
- HTTP route: `POST /api/v1.0/models`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Models`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$filter"]` | `$filter` | `query` | `no` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$s"]` | `$s` | `query` | `no` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$v"]` | `$v` | `query` | `no` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any` | Request body observed in HAR traffic. The SDK uses a single `json_body` payload because the Silver route carries a complex undocumented schema. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
