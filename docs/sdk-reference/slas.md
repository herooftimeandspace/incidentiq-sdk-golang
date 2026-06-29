# `slas` Golden Namespace

Go client access: `client.Slas`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_sla` | `DELETE /slas/{SlaId}` |
| `get` | `get_sla` | `GET /slas/{SlaId}` |
| `list` | `get_slas` | `GET /slas` |
| `update` | `update_sla` | `POST /slas/{SlaId}` |

## Methods

### `activate_sla`

- Go wrapper: `client.Slas.ActivateSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "activate_sla", opts, out)`
- HTTP route: `POST /slas/{SlaId}/activate`
- Source controller: `SLAs`

Activate SLA

#### Activates an SLA
#### Sample request:
```
POST /api/v1.0/slas/bd64e104-4c83-4744-a888-eeb760c03bfe/activate
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | SLA ID to be activated |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `deactivate_sla`

- Go wrapper: `client.Slas.DeactivateSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "deactivate_sla", opts, out)`
- HTTP route: `POST /slas/{SlaId}/deactivate`
- Source controller: `SLAs`

Deactivate SLA

#### Deactivates an SLA
#### Sample request:
```
POST /api/v1.0/slas/bd64e104-4c83-4744-a888-eeb760c03bfe/deactivate
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | SLA ID to be deactivated |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_sla`

- Go wrapper: `client.Slas.DeleteSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "delete_sla", opts, out)`
- HTTP route: `DELETE /slas/{SlaId}`
- Source controller: `SLAs`
- Aliases: `delete`

Delete SLA 

#### Delete a specific SLA
#### Sample request:
```
DELETE /api/v1.0/slas/bd64e104-4c83-4744-a888-eeb760c03bfe
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | SLA ID to be activated |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_sla`

- Go wrapper: `client.Slas.GetSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "get_sla", opts, out)`
- HTTP route: `GET /slas/{SlaId}`
- Source controller: `SLAs`
- Aliases: `get`

Get SLA

#### Retrieve a specific SLA  by SlaId
#### Sample request:
```
GET /api/v1.0/slas/{SlaId:guid}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | SlaId of SLA being requested |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the SlaId |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfSla` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfSla`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_slas`

- Go wrapper: `client.Slas.GetSlas(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "get_slas", opts, out)`
- HTTP route: `GET /slas`
- Source controller: `SLAs`
- Aliases: `list`

Get SLAs

#### Retrieves a list of available SLAs to apply to a ticket
#### Sample request:
```
GET /api/v1.0/slas
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | Request Options specified for the Sla |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfSla` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfSla`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_sla`

- Go wrapper: `client.Slas.UpdateSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "slas", "update_sla", opts, out)`
- HTTP route: `POST /slas/{SlaId}`
- Source controller: `SLAs`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["SlaId"]` | `SlaId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateSlaRequest` | `UpdateSlaRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
