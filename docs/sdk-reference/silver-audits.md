# `silver.audits` Namespace

Go client access: `client.Silver.Audits`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_policies_schedules_for_asset`

- Go wrapper: `client.Silver.Audits.GetPoliciesSchedulesForAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "audits", "get_policies_schedules_for_asset", opts, out)`
- HTTP route: `GET /api/v1.0/audits/policies/schedules/for-asset/{for_asset_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Audits`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["for_asset_id"]` | `for_asset_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
