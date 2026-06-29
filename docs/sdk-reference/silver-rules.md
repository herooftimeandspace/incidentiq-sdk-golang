# `silver.rules` Namespace

Go client access: `client.Silver.Rules`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_log`

- Go wrapper: `client.Silver.Rules.GetLog(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "rules", "get_log", opts, out)`
- HTTP route: `GET /api/v1.0/rules/logs/{log_id}/{n_888891ac_91aa_e711_80c2_100dffa00003_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Rules`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["log_id"]` | `log_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_888891ac_91aa_e711_80c2_100dffa00003_id"]` | `n_888891ac_91aa_e711_80c2_100dffa00003_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `Params["$s"]` | `$s` | `query` | `yes` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_ids`

- Go wrapper: `client.Silver.Rules.PostIds(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "rules", "post_ids", opts, out)`
- HTTP route: `POST /api/v1.0/rules/ids`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Rules`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any | []any | string` | Request body observed in HAR traffic. The SDK keeps it as a single payload parameter because the undocumented route did not expose a stable object shape. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
