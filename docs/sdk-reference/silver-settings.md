# `Silver.settings` Namespace

Go client access: `client.Silver.Settings`


These methods are Silver because Stoplight does not publish direct Golden contracts for them, or because the SDK intentionally wraps a narrower Silver workflow around existing Golden operations. They remain separate so undocumented or convenience behavior never overrides the documented SDK surface.

## Methods

### `get_apps__remote_beyond_trust__remote_connection_url`

- Go wrapper: `client.Silver.Settings.GetAppsRemoteBeyondTrustRemoteConnectionUrl(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_apps__remote_beyond_trust__remote_connection_url", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Apps.RemoteBeyondTrust.RemoteConnectionUrl`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps__remote_team_viewer__api_token`

- Go wrapper: `client.Silver.Settings.GetAppsRemoteTeamViewerApiToken(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_apps__remote_team_viewer__api_token", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Apps.RemoteTeamViewer.ApiToken`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps__remote_team_viewer__download_link`

- Go wrapper: `client.Silver.Settings.GetAppsRemoteTeamViewerDownloadLink(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_apps__remote_team_viewer__download_link", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Apps.RemoteTeamViewer.DownloadLink`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps__remote_team_viewer__primary_group_id`

- Go wrapper: `client.Silver.Settings.GetAppsRemoteTeamViewerPrimaryGroupId(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_apps__remote_team_viewer__primary_group_id", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Apps.RemoteTeamViewer.PrimaryGroupId`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_apps__sis__custom_field_mapping__enabled`

- Go wrapper: `client.Silver.Settings.GetAppsSisCustomFieldMappingEnabled(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_apps__sis__custom_field_mapping__enabled", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Apps.Sis.CustomFieldMapping.Enabled`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_grids__user_options`

- Go wrapper: `client.Silver.Settings.GetGridsUserOptions(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_grids__user_options", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Grids.UserOptions`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`, `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_preferred_mappings`

- Go wrapper: `client.Silver.Settings.GetPreferredMappings(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_preferred_mappings", opts, out)`
- HTTP route: `GET /api/v1.0/settings/preferred-mappings`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_shortcuts__recent_ids`

- Go wrapper: `client.Silver.Settings.GetShortcutsRecentIds(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_shortcuts__recent_ids", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Shortcuts.RecentIds`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_tickets__status_lock`

- Go wrapper: `client.Silver.Settings.GetTicketsStatusLock(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_tickets__status_lock", opts, out)`
- HTTP route: `GET /api/v1.0/settings/Tickets.StatusLock`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `get_undefined`

- Go wrapper: `client.Silver.Settings.GetUndefined(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "get_undefined", opts, out)`
- HTTP route: `GET /api/v1.0/settings/undefined`
- Observed in: `demo.incidentiq.com.har`

HAR-derived undocumented GET route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

This Silver route does not define inferred parameters.

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---

### `post_search__recent_user_searches`

- Go wrapper: `client.Silver.Settings.PostSearchRecentUserSearches(ctx, opts, out)`
- Dynamic helper: `client.RequestSilver(ctx, "settings", "post_search__recent_user_searches", opts, out)`
- HTTP route: `POST /api/v1.0/settings/Search.RecentUserSearches`
- Observed in: `Chromebook-asset-actions.har`, `apple-asset-actions.har`

HAR-derived undocumented POST route for `client.Silver.Settings`.

This method is intentionally kept on the Silver surface because bundled Stoplight controller contracts do not define this route. Golden Stoplight operations remain the preferred contract source whenever they exist, so Silver only supplements gaps observed in tenant HAR traffic.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Description |
| --- | --- | --- | --- | --- | --- |
| `JSON` | `ProductId` | `body` | `yes` | `any | nil` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `Scope` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SettingTypeId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `SiteId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `UserId` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |
| `JSON` | `Value` | `body` | `yes` | `string` | Body field inferred from HAR observations for this undocumented Silver route. |

#### Returns

- Go wrapper return: `error`; decoded `map[string]any | []any | nil` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: Raw JSON payload only; this Silver route has no Golden schema contract.

---
