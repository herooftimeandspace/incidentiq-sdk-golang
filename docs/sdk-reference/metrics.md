# `metrics` Golden Namespace

Go client access: `client.Metrics`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_metric` | `DELETE /metrics/{MetricId}` |
| `get` | `get_metric` | `GET /metrics/{MetricId}` |
| `list` | `get_metrics` | `GET /metrics` |
| `update` | `update_metric` | `POST /metrics/{MetricId}` |

## Methods

### `delete_metric`

- Go wrapper: `client.Metrics.DeleteMetric(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "delete_metric", opts, out)`
- HTTP route: `DELETE /metrics/{MetricId}`
- Source controller: `SLAs`
- Aliases: `delete`

Delete SLA Metric

#### Delete a specific SLA Metric
#### Sample request:
```
DELETE /api/v1.0/metrics/2c6101d2-1ac8-4320-b234-74f51a3b2e58
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricId"]` | `MetricId` | `path` | `yes` | `string` | `-` | MetricID to be deleted |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_metric_type`

- Go wrapper: `client.Metrics.DeleteMetricType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "delete_metric_type", opts, out)`
- HTTP route: `DELETE /metrics/types/{MetricTypeId}`
- Source controller: `SLAs`

Delete Metric Type

#### Delete a specific SLA Metric Type
#### Sample request:
```
DELETE /api/v1.0/metrics/types/67a39334-d778-487c-95ae-07a776ed8201
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricTypeId"]` | `MetricTypeId` | `path` | `yes` | `string` | `-` | MetricTypeID of the MetricType to be deleted |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_metric`

- Go wrapper: `client.Metrics.GetMetric(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "get_metric", opts, out)`
- HTTP route: `GET /metrics/{MetricId}`
- Source controller: `SLAs`
- Aliases: `get`

Get SLA Metric

#### Retrieve a specific metric type by MetricId
#### Sample request:
```
GET /api/v1.0/metrics/2c6101d2-1ac8-4320-b234-74f51a3b2e58
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricId"]` | `MetricId` | `path` | `yes` | `string` | `-` | MetricId of Metric being requested |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the Metric(s) |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfMetric` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfMetric`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_metric_type`

- Go wrapper: `client.Metrics.GetMetricType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "get_metric_type", opts, out)`
- HTTP route: `GET /metrics/types/{MetricTypeId}`
- Source controller: `SLAs`

Get Metric Type

#### Retrieve a specific metric type by MetricTypeId
#### Sample request:
```
GET /api/v1.0/metrics/types/67a39334-d778-487c-95ae-07a776ed8201
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricTypeId"]` | `MetricTypeId` | `path` | `yes` | `string` | `-` | MetricTypeId of the MetricType being requested |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the MetricType(s) |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfMetricType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfMetricType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_metric_types`

- Go wrapper: `client.Metrics.GetMetricTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "get_metric_types", opts, out)`
- HTTP route: `GET /metrics/types`
- Source controller: `SLAs`

Get Metric Types

#### Retrieves a list of metric types. A specific metric type can be retrieved via GET `api/v1.0/metrics/types/{MetricTypeId}`.
#### Sample request:
```
GET /api/v1.0/metrics/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the MetricType(s) |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfMetricType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfMetricType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_metrics`

- Go wrapper: `client.Metrics.GetMetrics(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "get_metrics", opts, out)`
- HTTP route: `GET /metrics`
- Source controller: `SLAs`
- Aliases: `list`

Get All SLA Metrics

#### Retrieves a list of metrics. A specific metric can be retrieved via GET `api/v1.0/metrics/{MetricId}`.
#### Sample request:
```
GET /api/v1.0/metrics/
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the Metric(s) |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfMetric` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfMetric`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_metrics_for_sla`

- Go wrapper: `client.Metrics.GetMetricsForSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "get_metrics_for_sla", opts, out)`
- HTTP route: `GET /metrics/for/sla/{SlaId}`
- Source controller: `SLAs`

Get Metrics for an SLA

#### Retrieves a list of metrics for a specific SLA.
#### Sample request:
```
GET /api/v1.0/metrics/for/sla/bd64e104-4c83-4744-a888-eeb760c03bfe
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | SlaId of SLA containing the Metrics being requested |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the Sla Metrics |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfMetric` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfMetric`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_metric`

- Go wrapper: `client.Metrics.UpdateMetric(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "update_metric", opts, out)`
- HTTP route: `POST /metrics/{MetricId}`
- Source controller: `SLAs`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricId"]` | `MetricId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateMetricRequest` | `UpdateMetricRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_metric_type`

- Go wrapper: `client.Metrics.UpdateMetricType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "update_metric_type", opts, out)`
- HTTP route: `POST /metrics/types/{MetricTypeId}`
- Source controller: `SLAs`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["MetricTypeId"]` | `MetricTypeId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateMetricTypeRequest` | `UpdateMetricTypeRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_metrics_for_sla`

- Go wrapper: `client.Metrics.UpdateMetricsForSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "metrics", "update_metrics_for_sla", opts, out)`
- HTTP route: `POST /metrics/for/sla/{SlaId}`
- Source controller: `SLAs`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Items` | `body` | `yes` | `[]any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
