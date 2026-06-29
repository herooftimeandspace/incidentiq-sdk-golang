# `Silver.products` Namespace

Go client access: `client.Silver.Products`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_available_features`

- Go wrapper: `client.Silver.Products.GetAvailableFeatures(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "products", "get_available_features", opts, out)`
- HTTP route: `GET /api/v1.0/products/available-features`
- Observed in: `apple-asset-actions.har`

HAR-derived undocumented GET route for `client.Silver.Products`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
