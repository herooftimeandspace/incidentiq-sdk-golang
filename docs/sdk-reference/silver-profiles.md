# `Silver.profiles` Namespace

Go client access: `client.Silver.Profiles`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `post_profile_picture`

- Go wrapper: `client.Silver.Profiles.PostProfilePicture(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "profiles", "post_profile_picture", opts, out)`
- HTTP route: `POST /api/v1.0/profiles/{user_id}/picture`
- Observed in: `iiq-profile-picture.har`

HAR-derived undocumented POST route for `client.Silver.Profiles`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

The Go wrapper sends the multipart body supplied through `RequestOptions.Body` with the caller-provided multipart `ContentType`. The caller is responsible for preparing the image bytes and multipart form body before calling the wrapper. Browser-observed profile photo uploads omit the SDK `Client` header, so callers that need to match the HAR exactly can set `RequestOptions.OmitClientHeader`.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["user_id"]` | `user_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `Body` and `ContentType` | `File` | `file` | `yes` | `multipart/form-data bytes` | Multipart image field inferred from HAR observations for this undocumented Silver route. Build a multipart form with a `File` part and pass the encoded bytes plus the writer's form-data content type. |
| `OmitClientHeader` | `Client` | `header` | `no` | `bool` | Set to `true` when the upload must match HAR-observed browser traffic that does not include the SDK-provided `Client: ApiClient` header. |

#### Examples

```go
var body bytes.Buffer
writer := multipart.NewWriter(&body)
part, err := writer.CreateFormFile("File", "photo.png")
if err != nil {
	return err
}
if _, err := part.Write(photoBytes); err != nil {
	return err
}
if err := writer.Close(); err != nil {
	return err
}

var payload map[string]any
err = client.Silver.Profiles.PostProfilePicture(ctx, incidentiq.RequestOptions{
	PathParams:       map[string]any{"user_id": userID},
	Body:             body.Bytes(),
	ContentType:      writer.FormDataContentType(),
	OmitClientHeader: true,
	OmitSiteIDHeader: true,
}, &payload)
```

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `remove_profile_picture`

- Source SDK helper `profiles.remove_profile_picture` is not exposed by the generated Go wrapper surface.
- Backing routes: `GET /api/v1.0/users/{user_id}`, `POST /api/v1.0/users/{user_id}`

Remove a user's profile picture by clearing `PhotoId` through the documented Golden user update route.

This source SDK helper exists because Incident IQ exposes profile-picture removal as a user update workflow rather than as its own published profile route. The current generated Go surface does not wrap that source SDK helper. Use the Golden user read/update wrappers or `Client.Request` directly when implementing the same workflow in Go.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["user_id"]` | `user_id` | `path` | `yes` | `string` | User identifier whose profile picture should be removed. |
| `RequestOptions` | `wait_for_consistency` | `client` | `no` | `bool` | When true, poll user readback until `PhotoId` becomes `nil` or return a timeout error if the tenant does not converge in time. |
| `RequestOptions` | `consistency_timeout` | `client` | `no` | `float64` | Maximum number of seconds to wait for readback convergence when `wait_for_consistency=true`. |
| `RequestOptions` | `consistency_poll_interval` | `client` | `no` | `float64` | Polling interval in seconds between user readback checks when `wait_for_consistency=true`. |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfUser` responses are written into `out`.
- Internal raw fallback return: `map[string]any | []any | nil`
- Response model: `ItemUpdateResponseOfUser`

---
