# `Silver.sites` Namespace

Go client access: `client.Silver.Sites`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_deployments`

- Go wrapper: `client.Silver.Sites.GetDeployments(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "sites", "get_deployments", opts, out)`
- HTTP route: `GET /api/v1.0/sites/deployments`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented GET route for `client.Silver.Sites`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_roles`

- Go wrapper: `client.Silver.Sites.PostRoles(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "sites", "post_roles", opts, out)`
- HTTP route: `POST /api/v1.0/sites/roles`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Sites`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["$filter"]` | `$filter` | `query` | `no` | `string` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$s"]` | `$s` | `query` | `no` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |
| `Params["$v"]` | `$v` | `query` | `no` | `int` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
