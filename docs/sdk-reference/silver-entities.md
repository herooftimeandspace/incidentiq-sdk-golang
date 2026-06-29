# `Silver.entities` Namespace

Go client access: `client.Silver.Entities`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_types_fields_withoptions`

- Go wrapper: `client.Silver.Entities.GetTypesFieldsWithoptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "entities", "get_types_fields_withoptions", opts, out)`
- HTTP route: `GET /api/v1.0/entities/types/{type_id}/fields/withoptions`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Entities`.

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
