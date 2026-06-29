# `apps` Silver Namespace

Go client access: `client.Silver.Apps.<AppNamespace>` where `<AppNamespace>` is the generated app service listed below.



These methods are Silver because Stoplight does not publish Golden contracts for them.

| Service | Manual Helpers | Generic Silver Methods | Access Path |
| --- | ---: | ---: | --- |
| `aeries_sis` | 0 | 7 | `client.Silver.Apps.AeriesSis` |
| `file_wave` | 0 | 2 | `client.Silver.Apps.FileWave` |
| `google_device_data` | 6 | 11 | `client.Silver.Apps.GoogleDeviceData` |
| `infinite_campus` | 0 | 1 | `client.Silver.Apps.InfiniteCampus` |
| `jamf` | 0 | 1 | `client.Silver.Apps.Jamf` |
| `jamf_school` | 0 | 1 | `client.Silver.Apps.JamfSchool` |
| `local_asset_manager` | 0 | 1 | `client.Silver.Apps.LocalAssetManager` |
| `lock_n_charge2` | 0 | 3 | `client.Silver.Apps.LockNCharge2` |
| `meraki_mdm` | 0 | 1 | `client.Silver.Apps.MerakiMdm` |
| `microsoft_intune` | 6 | 23 | `client.Silver.Apps.MicrosoftIntune` |
| `mosyle` | 4 | 0 | Source SDK helper only; use `Client.Request` |
| `mosyle_manager` | 0 | 4 | `client.Silver.Apps.MosyleManager` |
| `password_reset` | 0 | 1 | `client.Silver.Apps.PasswordReset` |
| `payments_fee_tracker` | 0 | 3 | `client.Silver.Apps.PaymentsFeeTracker` |
| `payments_in_touch` | 0 | 3 | `client.Silver.Apps.PaymentsInTouch` |
| `payments_my_school_bucks` | 0 | 3 | `client.Silver.Apps.PaymentsMySchoolBucks` |
| `payments_square` | 0 | 3 | `client.Silver.Apps.PaymentsSquare` |
| `payments_stripe` | 0 | 3 | `client.Silver.Apps.PaymentsStripe` |
| `payments_vanco` | 0 | 3 | `client.Silver.Apps.PaymentsVanco` |
| `policy_agreements` | 0 | 1 | `client.Silver.Apps.PolicyAgreements` |
| `registry` | 2 | 0 | Source SDK helper only; use `Client.Request` |
| `remote_beyond_trust` | 0 | 1 | `client.Silver.Apps.RemoteBeyondTrust` |
| `remote_chrome` | 0 | 1 | `client.Silver.Apps.RemoteChrome` |
| `remote_team_viewer` | 0 | 1 | `client.Silver.Apps.RemoteTeamViewer` |
| `spare_pool_management` | 0 | 4 | `client.Silver.Apps.SparePoolManagement` |
| `trafera` | 0 | 2 | `client.Silver.Apps.Trafera` |
| `widgets` | 0 | 1 | `client.Silver.Apps.Widgets` |
| `workspace_one` | 0 | 1 | `client.Silver.Apps.WorkspaceOne` |

## `aeries_sis`

Aeries Sis service available at `client.Silver.Apps.AeriesSis`.

### `get_auth_oneroster_validate`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetAuthOnerosterValidate(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_auth_oneroster_validate", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/auth/oneroster/validate`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_settings_options`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetSettingsOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_settings_options", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/settings/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_settings_options_sync`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetSettingsOptionsSync(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_settings_options_sync", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/settings/options/sync`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_history`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetSyncHistory(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_sync_history", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/data/sync/history`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_locations`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetSyncLocations(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_sync_locations", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/sync/data/locations`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_status_last`

- Go wrapper: `client.Silver.Apps.AeriesSis.GetSyncStatusLast(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "get_sync_status_last", opts, out)`
- HTTP route: `GET /apps/aeriesSis/api/aeriesSis/sync/status/last`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_users_search`

- Go wrapper: `client.Silver.Apps.AeriesSis.PostUsersSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.aeries_sis", "post_users_search", opts, out)`
- HTTP route: `POST /apps/aeriesSis/api/aeriesSis/data/users/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.AeriesSis`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `Limit` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `Skip` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `Sort` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `file_wave`

File Wave service available at `client.Silver.Apps.FileWave`.

### `get_assets_usage`

- Go wrapper: `client.Silver.Apps.FileWave.GetAssetsUsage(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.file_wave", "get_assets_usage", opts, out)`
- HTTP route: `GET /apps/fileWave/api/fileWave/data/assets/usage/{usage_id}/{n_883d10f1_e4a0_4268_a319_3d36d1948030_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.FileWave`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["usage_id"]` | `usage_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_883d10f1_e4a0_4268_a319_3d36d1948030_id"]` | `n_883d10f1_e4a0_4268_a319_3d36d1948030_id` | `path` | `yes` | `int` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_assets_lookup`

- Go wrapper: `client.Silver.Apps.FileWave.PostAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.file_wave", "post_assets_lookup", opts, out)`
- HTTP route: `POST /apps/fileWave/api/fileWave/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.FileWave`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `google_device_data`

Google Device Data service available at `client.Silver.Apps.GoogleDeviceData`.

### `lookup_asset`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.lookup_asset` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/data/assets/get-google-device`

Look up an Incident IQ asset against Google Device Data.

Posts the asset lookup payload to the Google Device Data endpoint.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |
| `JSON` | `Query` | `body` | `no` | `string` | Optional search query. |
| `JSON` | `Skip` | `body` | `no` | `int` | Result offset for the Google endpoint. |
| `JSON` | `Limit` | `body` | `no` | `int` | Maximum results requested. |

#### Returns

- Go wrapper return: `error`; decoded `AppLookupResponse | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response model: `AppLookupResponse`

---

### `lookup_asset_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.lookup_asset_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/data/assets/get-google-device`

Look up an asset against Google Device Data and return raw JSON.

Same request as `lookup_asset`, but returns validated raw JSON.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |
| `JSON` | `Query` | `body` | `no` | `string` | Optional search query. |
| `JSON` | `Skip` | `body` | `no` | `int` | Result offset for the Google endpoint. |
| `JSON` | `Limit` | `body` | `no` | `int` | Maximum results requested. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response schema: `lookup_response`

---

### `list_remote_actions`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.list_remote_actions` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/remoteactions`

List available Google Device Data remote actions.

Calls the Google Device Data remote actions endpoint and returns typed action records.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[AppRemoteAction]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response model: `AppRemoteAction`

---

### `list_remote_actions_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.list_remote_actions_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/remoteactions`

List available Google Device Data remote actions and return raw JSON.

Same request as `list_remote_actions`, but returns validated raw JSON.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[map[string]any]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response schema: `remote_actions_response`

---

### `get_sync_options`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.get_sync_options` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/sync/options`

Fetch Google Device Data sync options.

Calls the sync options endpoint and returns the typed options model.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `GoogleSyncOptionsResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response model: `GoogleSyncOptionsResponse`

---

### `get_sync_options_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.google_device_data.get_sync_options_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/sync/options`

Fetch Google Device Data sync options and return raw JSON.

Same request as `get_sync_options`, but returns validated raw JSON.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response schema: `google_sync_options_response`

---

### `get_assets_usage`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetAssetsUsage(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_assets_usage", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/data/assets/usage/{usage_id}/{n_883d10f1_e4a0_4268_a319_3d36d1948030_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["usage_id"]` | `usage_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_883d10f1_e4a0_4268_a319_3d36d1948030_id"]` | `n_883d10f1_e4a0_4268_a319_3d36d1948030_id` | `path` | `yes` | `int` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_assignment_suggestions`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetAssignmentSuggestions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_assignment_suggestions", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/data/assignment/suggestions`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_auth_token_check`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetAuthTokenCheck(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_auth_token_check", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/auth/token-check`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_models_distinct`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetModelsDistinct(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_models_distinct", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/data/models/distinct`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_status_last`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetStatusLast(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_status_last", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/{google_device_data_key}/status/last`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["google_device_data_key"]` | `google_device_data_key` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_history`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.GetSyncHistory(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "get_sync_history", opts, out)`
- HTTP route: `GET /apps/googleDeviceData/api/googleDeviceData/data/sync/history`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_device_sync_push_schedule`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.PostDeviceSyncPushSchedule(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "post_device_sync_push_schedule", opts, out)`
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/device/sync/push/schedule/{schedule_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["schedule_id"]` | `schedule_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_jobs_logs`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.PostJobsLogs(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "post_jobs_logs", opts, out)`
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/data/jobs/{job_id}/logs`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["job_id"]` | `job_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `Limit` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_deprovision_device_execute`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.PostRemoteactionsDeprovisionDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "post_remoteactions_deprovision_device_execute", opts, out)`
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/remoteactions/DeprovisionDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `json_body` | `body` | `yes` | `map[string]any` | Request body observed in HAR traffic. The SDK uses a single `json_body` payload because the Silver route carries a complex undocumented schema. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_disable_device_execute`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.PostRemoteactionsDisableDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "post_remoteactions_disable_device_execute", opts, out)`
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/remoteactions/DisableDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_wipe_device_execute`

- Go wrapper: `client.Silver.Apps.GoogleDeviceData.PostRemoteactionsWipeDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.google_device_data", "post_remoteactions_wipe_device_execute", opts, out)`
- HTTP route: `POST /apps/googleDeviceData/api/googleDeviceData/remoteactions/WipeDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.GoogleDeviceData`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `infinite_campus`

Infinite Campus service available at `client.Silver.Apps.InfiniteCampus`.

### `get_settings_options_sync`

- Go wrapper: `client.Silver.Apps.InfiniteCampus.GetSettingsOptionsSync(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.infinite_campus", "get_settings_options_sync", opts, out)`
- HTTP route: `GET /apps/infiniteCampus/api/infiniteCampus/settings/options/sync`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.InfiniteCampus`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `jamf`

Jamf service available at `client.Silver.Apps.Jamf`.

### `post_assets_lookup`

- Go wrapper: `client.Silver.Apps.Jamf.PostAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.jamf", "post_assets_lookup", opts, out)`
- HTTP route: `POST /apps/jamf/api/jamf/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.Jamf`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `jamf_school`

Jamf School service available at `client.Silver.Apps.JamfSchool`.

### `post_assets_lookup`

- Go wrapper: `client.Silver.Apps.JamfSchool.PostAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.jamf_school", "post_assets_lookup", opts, out)`
- HTTP route: `POST /apps/jamfSchool/api/jamfSchool/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.JamfSchool`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `local_asset_manager`

Local Asset Manager service available at `client.Silver.Apps.LocalAssetManager`.

### `post_microsoft_sccm_assets_lookup`

- Go wrapper: `client.Silver.Apps.LocalAssetManager.PostMicrosoftSccmAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.local_asset_manager", "post_microsoft_sccm_assets_lookup", opts, out)`
- HTTP route: `POST /apps/localAssetManager/api/microsoftSccm/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.LocalAssetManager`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `lock_n_charge2`

Lock N Charge2 service available at `client.Silver.Apps.LockNCharge2`.

### `get_bays`

- Go wrapper: `client.Silver.Apps.LockNCharge2.GetBays(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.lock_n_charge2", "get_bays", opts, out)`
- HTTP route: `GET /apps/lockNCharge2/api/bays`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.LockNCharge2`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `Params["refresh"]` | `refresh` | `query` | `yes` | `bool` | Query parameter inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_stations`

- Go wrapper: `client.Silver.Apps.LockNCharge2.GetStations(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.lock_n_charge2", "get_stations", opts, out)`
- HTTP route: `GET /apps/lockNCharge2/api/stations`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.LockNCharge2`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_stations_licensed`

- Go wrapper: `client.Silver.Apps.LockNCharge2.GetStationsLicensed(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.lock_n_charge2", "get_stations_licensed", opts, out)`
- HTTP route: `GET /apps/lockNCharge2/api/stations/licensed`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.LockNCharge2`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `meraki_mdm`

Meraki Mdm service available at `client.Silver.Apps.MerakiMdm`.

### `post_assets_lookup`

- Go wrapper: `client.Silver.Apps.MerakiMdm.PostAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.meraki_mdm", "post_assets_lookup", opts, out)`
- HTTP route: `POST /apps/merakiMdm/api/merakiMdm/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MerakiMdm`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `microsoft_intune`

Microsoft Intune service available at `client.Silver.Apps.MicrosoftIntune`.

### `lookup_asset`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.lookup_asset` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/data/assets/lookup`

Look up an Incident IQ asset against Microsoft Intune.

Posts the asset lookup payload to the Intune app endpoint and returns the typed lookup response when available.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number sent to the Intune lookup endpoint. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |

#### Returns

- Go wrapper return: `error`; decoded `AppLookupResponse | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response model: `AppLookupResponse`

---

### `lookup_asset_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.lookup_asset_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/data/assets/lookup`

Look up an asset against Microsoft Intune and return raw JSON.

Same request as `lookup_asset`, but returns validated raw JSON.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number sent to the Intune lookup endpoint. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response schema: `lookup_response`

---

### `list_remote_actions`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.list_remote_actions` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/remoteactions`

List available Intune remote actions.

Calls the Intune remote actions endpoint and returns typed action records.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[AppRemoteAction]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response model: `AppRemoteAction`

---

### `list_remote_actions_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.list_remote_actions_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/remoteactions`

List available Intune remote actions and return raw JSON.

Same request as `list_remote_actions`, but returns validated raw JSON.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[map[string]any]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response schema: `remote_actions_response`

---

### `classify_owner_type_from_lookup`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.classify_owner_type_from_lookup` is not exposed by the generated Go wrapper surface.
- HTTP route: Utility helper (no HTTP request)

Classify Intune owner type from a lookup payload.

Utility helper that derives owner type and optional external-id match state from a lookup response; no HTTP request is made.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `RequestOptions` | `lookup_response` | `go` | `yes` | `map[string]any` | Lookup payload or model to classify. |
| `RequestOptions` | `expected_external_id` | `go` | `no` | `string` | Optional external id used to flag mismatches. |

#### Returns

- Go wrapper return: `error`; decoded `IntuneOwnerClassification` responses are written into `out`.
- Decoded response: caller-provided `out` receives `IntuneOwnerClassification` when the route returns JSON.
- Response model: `IntuneOwnerClassification`

---

### `partition_assets_by_owner_type`

Provenance: Silver manual helper

- Source SDK helper `apps.microsoft_intune.partition_assets_by_owner_type` is not exposed by the generated Go wrapper surface.
- HTTP route: Utility helper (no HTTP request)

Partition Intune-linked assets by owner type.

Utility helper that performs lookups as needed and groups assets into company, personal, and unknown partitions.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `RequestOptions` | `assets` | `go` | `yes` | `[]map[string]any` | Asset payloads containing Intune app mapping data. |

#### Returns

- Go wrapper return: `error`; decoded `IntuneOwnershipPartition` responses are written into `out`.
- Decoded response: caller-provided `out` receives `IntuneOwnershipPartition` when the route returns JSON.
- Response model: `IntuneOwnershipPartition`

---

### `get_assets_usage`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.GetAssetsUsage(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "get_assets_usage", opts, out)`
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/data/assets/usage/{usage_id}/{n_883d10f1_e4a0_4268_a319_3d36d1948030_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["usage_id"]` | `usage_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["n_883d10f1_e4a0_4268_a319_3d36d1948030_id"]` | `n_883d10f1_e4a0_4268_a319_3d36d1948030_id` | `path` | `yes` | `int` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_auth_token_check`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.GetAuthTokenCheck(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "get_auth_token_check", opts, out)`
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/auth/token-check`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_history`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.GetSyncHistory(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "get_sync_history", opts, out)`
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/data/sync/history`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_sync_status_last`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.GetSyncStatusLast(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "get_sync_status_last", opts, out)`
- HTTP route: `GET /apps/microsoftIntune/api/microsoftIntune/sync/status/last`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_jobs_logs`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostJobsLogs(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_jobs_logs", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/data/jobs/{job_id}/logs`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["job_id"]` | `job_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `Limit` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_clean_windows_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsCleanWindowsDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_clean_windows_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/CleanWindowsDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `KeepUserData` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_delete_user_from_shared_apple_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsDeleteUserFromSharedAppleDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_delete_user_from_shared_apple_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/DeleteUserFromSharedAppleDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `UserPrincipalName` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_disable_lost_mode_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsDisableLostModeExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_disable_lost_mode_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/DisableLostMode/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_locate_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsLocateDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_locate_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/LocateDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_logout_shared_apple_device_active_user_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsLogoutSharedAppleDeviceActiveUserExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_logout_shared_apple_device_active_user_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/LogoutSharedAppleDeviceActiveUser/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_reboot_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsRebootDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_reboot_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/RebootDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_recover_passcode_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsRecoverPasscodeExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_recover_passcode_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/RecoverPasscode/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_remote_lock_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsRemoteLockExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_remote_lock_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/RemoteLock/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_request_remote_assistance_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsRequestRemoteAssistanceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_request_remote_assistance_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/RequestRemoteAssistance/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_reset_passcode_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsResetPasscodeExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_reset_passcode_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/ResetPasscode/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_retire_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsRetireDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_retire_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/RetireDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_shut_down_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsShutDownExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_shut_down_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/ShutDown/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_sync_device_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsSyncDeviceExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_sync_device_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/SyncDevice/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_update_windows_device_account_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsUpdateWindowsDeviceAccountExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_update_windows_device_account_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/UpdateWindowsDeviceAccount/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_windows_defender_scan_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsWindowsDefenderScanExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_windows_defender_scan_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/WindowsDefenderScan/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `QuickScan` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_windows_defender_update_signatures_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsWindowsDefenderUpdateSignaturesExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_windows_defender_update_signatures_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/WindowsDefenderUpdateSignatures/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_wipe_execute`

- Go wrapper: `client.Silver.Apps.MicrosoftIntune.PostRemoteactionsWipeExecute(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.microsoft_intune", "post_remoteactions_wipe_execute", opts, out)`
- HTTP route: `POST /apps/microsoftIntune/api/microsoftIntune/remoteactions/Wipe/execute/{execute_id}`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `windows-asset-intune-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MicrosoftIntune`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["execute_id"]` | `execute_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `JSON` | `KeepEnrollmentData` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `KeepUserData` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `MacOsUnlockCode` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `mosyle`

Mosyle helpers are source SDK helpers only in the current Go reference. They
are not exposed by the generated Go wrapper surface; use `Client.Request` with
the documented route when implementing these workflows in Go.

### `lookup_asset`

Provenance: Silver manual helper

- Source SDK helper `apps.mosyle.lookup_asset` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/data/assets/lookup`

Look up an Incident IQ asset against Mosyle.

Posts the asset lookup payload to the Mosyle app endpoint.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |

#### Returns

- Go wrapper return: `error`; decoded `AppLookupResponse | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response model: `AppLookupResponse`

---

### `lookup_asset_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.mosyle.lookup_asset_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/data/assets/lookup`

Look up an asset against Mosyle and return raw JSON.

Same request as `lookup_asset`, but returns validated raw JSON.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Incident IQ asset identifier. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Serial number. |
| `JSON` | `AssetTag` | `body` | `no` | `string` | Optional asset tag. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response schema: `lookup_response`

---

### `list_remote_actions`

Provenance: Silver manual helper

- Source SDK helper `apps.mosyle.list_remote_actions` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/mosyleManager/api/mosyleManager/remoteactions`

List available Mosyle remote actions.

Calls the Mosyle remote actions endpoint and returns typed action records.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[AppRemoteAction]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response model: `AppRemoteAction`

---

### `list_remote_actions_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.mosyle.list_remote_actions_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /apps/mosyleManager/api/mosyleManager/remoteactions`

List available Mosyle remote actions and return raw JSON.

Same request as `list_remote_actions`, but returns validated raw JSON.

#### Parameters

This method does not define parameters.

#### Returns

- Go wrapper return: `error`; decoded `list[map[string]any]` responses are written into `out`.
- Decoded response: caller-provided `out` receives `list[map[string]any]` when the route returns JSON.
- Response schema: `remote_actions_response`

---

## `mosyle_manager`

Mosyle Manager service available at `client.Silver.Apps.MosyleManager`.

### `post_remoteactions_clear_commands_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d`

- Go wrapper: `client.Silver.Apps.MosyleManager.PostRemoteactionsClearCommandsExecuteMac268d3c3f77455eb19299A92e984d642d(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.mosyle_manager", "post_remoteactions_clear_commands_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d", opts, out)`
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/remoteactions/ClearCommands/execute/mac268D3C3F-7745-5EB1-9299-A92E984D642D`
- Observed in: `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MosyleManager`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_restart_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d`

- Go wrapper: `client.Silver.Apps.MosyleManager.PostRemoteactionsRestartDeviceExecuteMac268d3c3f77455eb19299A92e984d642d(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.mosyle_manager", "post_remoteactions_restart_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d", opts, out)`
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/remoteactions/RestartDevice/execute/mac268D3C3F-7745-5EB1-9299-A92E984D642D`
- Observed in: `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MosyleManager`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_shutdown_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d`

- Go wrapper: `client.Silver.Apps.MosyleManager.PostRemoteactionsShutdownDeviceExecuteMac268d3c3f77455eb19299A92e984d642d(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.mosyle_manager", "post_remoteactions_shutdown_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d", opts, out)`
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/remoteactions/ShutdownDevice/execute/mac268D3C3F-7745-5EB1-9299-A92E984D642D`
- Observed in: `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MosyleManager`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_remoteactions_wipe_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d`

- Go wrapper: `client.Silver.Apps.MosyleManager.PostRemoteactionsWipeDeviceExecuteMac268d3c3f77455eb19299A92e984d642d(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.mosyle_manager", "post_remoteactions_wipe_device_execute_mac268d3c3f_7745_5eb1_9299_a92e984d642d", opts, out)`
- HTTP route: `POST /apps/mosyleManager/api/mosyleManager/remoteactions/WipeDevice/execute/mac268D3C3F-7745-5EB1-9299-A92E984D642D`
- Observed in: `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Apps.MosyleManager`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `pin_code` | `body` | `yes` | `int` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `password_reset`

Password Reset service available at `client.Silver.Apps.PasswordReset`.

### `get_user_setup`

- Go wrapper: `client.Silver.Apps.PasswordReset.GetUserSetup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.password_reset", "get_user_setup", opts, out)`
- HTTP route: `GET /apps/passwordReset/api/user/{user_id}/setup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PasswordReset`.

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

## `payments_fee_tracker`

Payments Fee Tracker service available at `client.Silver.Apps.PaymentsFeeTracker`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsFeeTracker.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_fee_tracker", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsFeeTracker/api/sync/options`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsFeeTracker`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsFeeTracker.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_fee_tracker", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsFeeTracker/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsFeeTracker`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsFeeTracker.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_fee_tracker", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsFeeTracker/api/data/line-items/search`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsFeeTracker`.

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

## `payments_in_touch`

Payments In Touch service available at `client.Silver.Apps.PaymentsInTouch`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsInTouch.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_in_touch", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsInTouch/api/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsInTouch`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsInTouch.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_in_touch", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsInTouch/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsInTouch`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsInTouch.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_in_touch", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsInTouch/api/data/line-items/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsInTouch`.

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

## `payments_my_school_bucks`

Payments My School Bucks service available at `client.Silver.Apps.PaymentsMySchoolBucks`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsMySchoolBucks.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_my_school_bucks", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsMySchoolBucks/api/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsMySchoolBucks`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsMySchoolBucks.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_my_school_bucks", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsMySchoolBucks/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsMySchoolBucks`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsMySchoolBucks.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_my_school_bucks", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsMySchoolBucks/api/data/line-items/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsMySchoolBucks`.

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

## `payments_square`

Payments Square service available at `client.Silver.Apps.PaymentsSquare`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsSquare.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_square", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsSquare/api/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsSquare`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsSquare.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_square", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsSquare/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsSquare`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsSquare.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_square", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsSquare/api/data/line-items/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsSquare`.

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

## `payments_stripe`

Payments Stripe service available at `client.Silver.Apps.PaymentsStripe`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsStripe.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_stripe", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsStripe/api/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsStripe`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsStripe.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_stripe", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsStripe/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsStripe`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsStripe.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_stripe", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsStripe/api/data/line-items/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsStripe`.

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

## `payments_vanco`

Payments Vanco service available at `client.Silver.Apps.PaymentsVanco`.

### `get_sync_options`

- Go wrapper: `client.Silver.Apps.PaymentsVanco.GetSyncOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_vanco", "get_sync_options", opts, out)`
- HTTP route: `GET /apps/paymentsVanco/api/sync/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PaymentsVanco`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_item_types_search`

- Go wrapper: `client.Silver.Apps.PaymentsVanco.PostLineItemTypesSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_vanco", "post_line_item_types_search", opts, out)`
- HTTP route: `POST /apps/paymentsVanco/api/data/line-item-types/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsVanco`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_line_items_search`

- Go wrapper: `client.Silver.Apps.PaymentsVanco.PostLineItemsSearch(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.payments_vanco", "post_line_items_search", opts, out)`
- HTTP route: `POST /apps/paymentsVanco/api/data/line-items/search`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.PaymentsVanco`.

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

## `policy_agreements`

Policy Agreements service available at `client.Silver.Apps.PolicyAgreements`.

### `get_users_acceptances_including_cleared`

- Go wrapper: `client.Silver.Apps.PolicyAgreements.GetUsersAcceptancesIncludingCleared(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.policy_agreements", "get_users_acceptances_including_cleared", opts, out)`
- HTTP route: `GET /apps/policyAgreements/api/users/{user_id}/acceptances/including-cleared`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.PolicyAgreements`.

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

## `registry`

App registry helpers are source SDK helpers only in the current Go reference.
They are not exposed by the generated Go wrapper surface; use `Client.Request`
with the documented route when implementing these workflows in Go.

### `list_apps`

Provenance: Silver manual helper

- Source SDK helper `apps.registry.list_apps` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /api/v1.0/app-registry/apps/{include_hidden}`

List registered tenant apps.

Calls the tenant app registry endpoint and returns the typed registry response envelope.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["include_hidden"]` | `include_hidden` | `path` | `no` | `bool` | Whether to include hidden app registrations. |

#### Returns

- Go wrapper return: `error`; decoded `AppRegistryResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response model: `AppRegistryResponse`

---

### `list_apps_raw`

Provenance: Silver manual helper

- Source SDK helper `apps.registry.list_apps_raw` is not exposed by the generated Go wrapper surface.
- HTTP route: `GET /api/v1.0/app-registry/apps/{include_hidden}`

List registered tenant apps and return raw JSON.

Same request as `list_apps`, but returns validated raw JSON.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["include_hidden"]` | `include_hidden` | `path` | `no` | `bool` | Whether to include hidden app registrations. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | nil` when the route returns JSON.
- Response schema: `registry_response`

---

## `remote_beyond_trust`

Remote Beyond Trust service available at `client.Silver.Apps.RemoteBeyondTrust`.

### `get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu`

- Go wrapper: `client.Silver.Apps.RemoteBeyondTrust.GetDbBb6cece8E4f4E511A789005056bb000eStatu(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.remote_beyond_trust", "get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu", opts, out)`
- HTTP route: `GET /api/v1.0/apps/remoteBeyondTrust/db/bb6cece8-e4f4-e511-a789-005056bb000e-Status/{bb6cece8_e4f4_e511_a789_005056bb000e_statu_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.RemoteBeyondTrust`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["bb6cece8_e4f4_e511_a789_005056bb000e_statu_id"]` | `bb6cece8_e4f4_e511_a789_005056bb000e_statu_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `remote_chrome`

Remote Chrome service available at `client.Silver.Apps.RemoteChrome`.

### `get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu`

- Go wrapper: `client.Silver.Apps.RemoteChrome.GetDbBb6cece8E4f4E511A789005056bb000eStatu(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.remote_chrome", "get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu", opts, out)`
- HTTP route: `GET /api/v1.0/apps/remoteChrome/db/bb6cece8-e4f4-e511-a789-005056bb000e-Status/{bb6cece8_e4f4_e511_a789_005056bb000e_statu_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.RemoteChrome`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["bb6cece8_e4f4_e511_a789_005056bb000e_statu_id"]` | `bb6cece8_e4f4_e511_a789_005056bb000e_statu_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `remote_team_viewer`

Remote Team Viewer service available at `client.Silver.Apps.RemoteTeamViewer`.

### `get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu`

- Go wrapper: `client.Silver.Apps.RemoteTeamViewer.GetDbBb6cece8E4f4E511A789005056bb000eStatu(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.remote_team_viewer", "get_db_bb6cece8_e4f4_e511_a789_005056bb000e_statu", opts, out)`
- HTTP route: `GET /api/v1.0/apps/remoteTeamViewer/db/bb6cece8-e4f4-e511-a789-005056bb000e-Status/{bb6cece8_e4f4_e511_a789_005056bb000e_statu_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.RemoteTeamViewer`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["bb6cece8_e4f4_e511_a789_005056bb000e_statu_id"]` | `bb6cece8_e4f4_e511_a789_005056bb000e_statu_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `spare_pool_management`

Spare Pool Management service available at `client.Silver.Apps.SparePoolManagement`.

### `get_pool_stats_today`

- Go wrapper: `client.Silver.Apps.SparePoolManagement.GetPoolStatsToday(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.spare_pool_management", "get_pool_stats_today", opts, out)`
- HTTP route: `GET /apps/sparePoolManagement/api/pool/stats/today`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.SparePoolManagement`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_ticket_pool`

- Go wrapper: `client.Silver.Apps.SparePoolManagement.GetTicketPool(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.spare_pool_management", "get_ticket_pool", opts, out)`
- HTTP route: `GET /apps/sparePoolManagement/api/ticket/{ticket_id}/pools/{pool_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.SparePoolManagement`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["ticket_id"]` | `ticket_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |
| `PathParams["pool_id"]` | `pool_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_assets_deployments`

- Go wrapper: `client.Silver.Apps.SparePoolManagement.PostAssetsDeployments(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.spare_pool_management", "post_assets_deployments", opts, out)`
- HTTP route: `POST /apps/sparePoolManagement/api/assets/deployments`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.SparePoolManagement`.

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

### `post_groups`

- Go wrapper: `client.Silver.Apps.SparePoolManagement.PostGroups(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.spare_pool_management", "post_groups", opts, out)`
- HTTP route: `POST /apps/sparePoolManagement/api/groups`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.SparePoolManagement`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `trafera`

Trafera service available at `client.Silver.Apps.Trafera`.

### `get_parts`

- Go wrapper: `client.Silver.Apps.Trafera.GetParts(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.trafera", "get_parts", opts, out)`
- HTTP route: `GET /apps/trafera/api/parts`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.Trafera`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_settings_options`

- Go wrapper: `client.Silver.Apps.Trafera.GetSettingsOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.trafera", "get_settings_options", opts, out)`
- HTTP route: `GET /apps/trafera/api/settings/options`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.Trafera`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `widgets`

Widgets service available at `client.Silver.Apps.Widgets`.

### `get_endpoint`

- Go wrapper: `client.Silver.Apps.Widgets.GetEndpoint(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.widgets", "get_endpoint", opts, out)`
- HTTP route: `GET /api/v1.0/apps/widgets/{widget_id}`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Apps.Widgets`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `PathParams["widget_id"]` | `widget_id` | `path` | `yes` | `string` | Path parameter inferred from HAR observations. This route remains on the Silver surface because Stoplight does not publish a Golden contract for it. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

## `workspace_one`

Workspace One service available at `client.Silver.Apps.WorkspaceOne`.

### `post_assets_lookup`

- Go wrapper: `client.Silver.Apps.WorkspaceOne.PostAssetsLookup(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "apps.workspace_one", "post_assets_lookup", opts, out)`
- HTTP route: `POST /apps/workspaceOne/api/workspaceOne/data/assets/lookup`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented POST route for `client.Silver.Apps.WorkspaceOne`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `AssetId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `AssetTag` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SerialNumber` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
