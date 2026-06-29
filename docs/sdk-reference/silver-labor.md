# `Silver.labor` Namespace

Go client access: `client.Silver.Labor`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_rates_user`

- Go wrapper: `client.Silver.Labor.GetRatesUser(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "labor", "get_rates_user", opts, out)`
- HTTP route: `GET /api/v1.0/labor/rates/user/{user_id}/{n_61757b5f_dd31_f111_8ef2_000d3a7cb1a2_id}/{n_88df910c_91aa_e711_80c2_0004ffa00010_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Labor`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["user_id"]` | `user_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_61757b5f_dd31_f111_8ef2_000d3a7cb1a2_id"]` | `n_61757b5f_dd31_f111_8ef2_000d3a7cb1a2_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_88df910c_91aa_e711_80c2_0004ffa00010_id"]` | `n_88df910c_91aa_e711_80c2_0004ffa00010_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_rates_user_2`

- Go wrapper: `client.Silver.Labor.GetRatesUser2(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "labor", "get_rates_user_2", opts, out)`
- HTTP route: `GET /api/v1.0/labor/rates/user/{user_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Labor`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["user_id"]` | `user_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_types`

- Go wrapper: `client.Silver.Labor.PostTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "labor", "post_types", opts, out)`
- HTTP route: `POST /api/v1.0/labor/types`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Labor`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
