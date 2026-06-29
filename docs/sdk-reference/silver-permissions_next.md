# `Silver.permissions_next` Namespace

Go client access: `client.Silver.PermissionsNext`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_users_user_preview_applied_policies`

- Go wrapper: `client.Silver.PermissionsNext.GetUsersUserPreviewAppliedPolicies(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "permissions_next", "get_users_user_preview_applied_policies", opts, out)`
- HTTP route: `GET /api/v1.0/permissions-next/users/{user_id}/user-preview-applied-policies`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.PermissionsNext`.

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

### `post_check_permission`

- Go wrapper: `client.Silver.PermissionsNext.PostCheckPermission(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "permissions_next", "post_check_permission", opts, out)`
- HTTP route: `POST /api/v1.0/permissions-next/check-permission`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.PermissionsNext`.

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
