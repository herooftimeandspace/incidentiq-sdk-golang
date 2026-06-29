# `analytics` Golden Namespace

Go client access: `client.Analytics`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Methods

### `get_asset_counts_by_audit_policy_coverage`

- Go wrapper: `client.Analytics.GetAssetCountsByAuditPolicyCoverage(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_audit_policy_coverage", opts, out)`
- HTTP route: `GET /analytics/assets/by-audit-policy-coverage`
- Source controller: `Analytics`

Get audit policy coverage

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset audit policy coverage
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-audit-policy-coverage
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_counts_by_audit_policy_schedule_status`

- Go wrapper: `client.Analytics.GetAssetCountsByAuditPolicyScheduleStatus(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_audit_policy_schedule_status", opts, out)`
- HTTP route: `GET /analytics/assets/by-audit-policy-schedule-status/{AssetAuditPolicyScheduleId}`
- Source controller: `Analytics`

Get status by schedule

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset audit status for a given audit policy schedule
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-audit-schedule-status/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetAuditPolicyScheduleId"]` | `AssetAuditPolicyScheduleId` | `path` | `yes` | `string` | `-` | Audit Policy Schedule ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_counts_by_audit_policy_status`

- Go wrapper: `client.Analytics.GetAssetCountsByAuditPolicyStatus(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_audit_policy_status", opts, out)`
- HTTP route: `GET /analytics/assets/by-audit-policy-status/{AssetAuditPolicyId}`
- Source controller: `Analytics`

Get status by policy

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset audit status for a given audit policy
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-audit-status/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetAuditPolicyId"]` | `AssetAuditPolicyId` | `path` | `yes` | `string` | `-` | Audit Policy ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_counts_by_audit_status`

- Go wrapper: `client.Analytics.GetAssetCountsByAuditStatus(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_audit_status", opts, out)`
- HTTP route: `GET /analytics/assets/by-audit-status`
- Source controller: `Analytics`

Get audit policy status

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset audit status
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-audit-status
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_counts_by_verification_location`

- Go wrapper: `client.Analytics.GetAssetCountsByVerificationLocation(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_verification_location", opts, out)`
- HTTP route: `GET /analytics/assets/by-verification-location`
- Source controller: `Analytics`

Get assets by verification location

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset verification location
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-verification-location
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_counts_by_verification_type`

- Go wrapper: `client.Analytics.GetAssetCountsByVerificationType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_asset_counts_by_verification_type", opts, out)`
- HTTP route: `GET /analytics/assets/by-verification-type`
- Source controller: `Analytics`

Get assets by verification type

#### Performs an analytics query to obtain a breakdown of asset counts grouped by asset verification type
#### Sample request:
```
GET /api/v1.0/analytics/assets/by-verification-type
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAnalyticsDataPoint` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAnalyticsDataPoint`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_report`

- Go wrapper: `client.Analytics.GetReport(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_report", opts, out)`
- HTTP route: `GET /analytics/reports/{ReportId}`
- Source controller: `Analytics`

Get report details

#### Get information related to a given report.  Meta information for the overall report is returned.  For a listing of defined report elements associated with this report call GET `/api/v1.0/analytics/reports/elements/{ReportId}`.
#### Sample request:
```
GET /api/v1.0/analytics/reports/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ReportId"]` | `ReportId` | `path` | `yes` | `string` | `-` | Report ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfReport` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfReport`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_report_elements`

- Go wrapper: `client.Analytics.GetReportElements(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_report_elements", opts, out)`
- HTTP route: `GET /analytics/reports/elements/{ReportId}`
- Source controller: `Analytics`

Get report elements

#### Get all elements defined for a given report.  Meta information for the overall report is not returned in this endpoint.  To obtain information regarding the overall report call GET `/api/v1.0/analytics/reports/{ReportId}`.
#### Sample request:
```
GET /api/v1.0/analytics/reports/elements/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ReportId"]` | `ReportId` | `path` | `yes` | `string` | `-` | Report ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfReportElement` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfReportElement`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_report_queries`

- Go wrapper: `client.Analytics.GetReportQueries(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_report_queries", opts, out)`
- HTTP route: `GET /analytics/reports/queries/{ReportId}`
- Source controller: `Analytics`

Get report queries

#### Get all queries defined for a given report.  Meta information for the overall report is not returned in this endpoint.  To obtain information regarding the overall report call GET `/api/v1.0/analytics/reports/{ReportId}`.
#### Sample request:
```
GET /api/v1.0/analytics/reports/queries/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ReportId"]` | `ReportId` | `path` | `yes` | `string` | `-` | Report ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfReportQuery` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfReportQuery`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_reports`

- Go wrapper: `client.Analytics.GetReports(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "analytics", "get_reports", opts, out)`
- HTTP route: `GET /analytics/reports`
- Source controller: `Analytics`

Get all reports

#### Get all currently active and defined reports
#### Sample request:
```
GET /api/v1.0/analytics/reports
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfReport` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfReport`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
