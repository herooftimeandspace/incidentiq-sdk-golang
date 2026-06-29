# `custom_fields` Golden Namespace

Go client access: `client.CustomFields`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_custom_field` | `DELETE /custom-fields/{CustomFieldId}` |
| `get` | `get_custom_field` | `GET /custom-fields/{CustomFieldId}` |
| `list` | `get_custom_fields` | `GET /custom-fields` |
| `create` | `get_custom_fields2` | `POST /custom-fields` |
| `update` | `update_custom_field` | `POST /custom-fields/{CustomFieldId}` |

## Methods

### `delete_custom_field`

- Go wrapper: `client.CustomFields.DeleteCustomField(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "delete_custom_field", opts, out)`
- HTTP route: `DELETE /custom-fields/{CustomFieldId}`
- Source controller: `Custom Fields`
- Aliases: `delete`

Delete Custom Field 

#### Delete a Custom Field
#### Sample request:
```
DELETE /api/v1.0/custom-fields/f10f540f-2a9f-47fe-acbd-8ca82dc73e7c
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldId"]` | `CustomFieldId` | `path` | `yes` | `string` | `-` | Custom Field ID to be deleted |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_custom_field_type`

- Go wrapper: `client.CustomFields.DeleteCustomFieldType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "delete_custom_field_type", opts, out)`
- HTTP route: `DELETE /custom-fields/types/{CustomFieldTypeId}`
- Source controller: `Custom Fields`

Delete Custom Field Type

#### Delete a Custom Field Type
#### Sample request:
```
DELETE /api/v1.0/custom-fields/types/108d5344-012f-4380-b865-3eaa72f7c42f
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldTypeId"]` | `CustomFieldTypeId` | `path` | `yes` | `string` | `-` | Custom Field Type ID to be deleted |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_custom_fields`

- Go wrapper: `client.CustomFields.DeleteCustomFields(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "delete_custom_fields", opts, out)`
- HTTP route: `DELETE /custom-fields`
- Source controller: `Custom Fields`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `CustomFieldIds` | `body` | `yes` | `[]any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field`

- Go wrapper: `client.CustomFields.GetCustomField(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field", opts, out)`
- HTTP route: `GET /custom-fields/{CustomFieldId}`
- Source controller: `Custom Fields`
- Aliases: `get`

Get Custom Field 

#### Retrieve a specific Custom Field
#### Sample request:
```
GET /api/v1.0/custom-fields/e60c4fe8-2236-48bb-8dd0-18a6f5b4c7ae
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldId"]` | `CustomFieldId` | `path` | `yes` | `string` | `-` | Custom Field ID to be retrieved |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfCustomField` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfCustomField`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_type`

- Go wrapper: `client.CustomFields.GetCustomFieldType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_type", opts, out)`
- HTTP route: `GET /custom-fields/types/{CustomFieldTypeId}`
- Source controller: `Custom Fields`

Get Custom Field Type

#### Retrieve a specific Custom Field Type
#### Sample request:
```
GET /api/v1.0/custom-fields/types/108d5344-012f-4380-b865-3eaa72f7c42f
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldTypeId"]` | `CustomFieldTypeId` | `path` | `yes` | `string` | `-` | Custom Field Type ID to be deleted |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfCustomFieldType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfCustomFieldType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_types`

- Go wrapper: `client.CustomFields.GetCustomFieldTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_types", opts, out)`
- HTTP route: `GET /custom-fields/types`
- Source controller: `Custom Fields`

Get Custom Field Types

#### Retrieves a list of CustomFieldTypes that can then be used to retrieve the custom field type via _custom-fields/types/{CustomFieldTypeId:guid}_
#### Sample request:
```
GET /api/v1.0/custom-fields/types
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfCustomFieldType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfCustomFieldType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_types2`

- Go wrapper: `client.CustomFields.GetCustomFieldTypes2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_types2", opts, out)`
- HTTP route: `POST /custom-fields/types`
- Source controller: `Custom Fields`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `CustomFieldTypesRequest` | `body` | `yes` | `GetCustomFieldTypesRequest` | `GetCustomFieldTypesRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfCustomFieldType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfCustomFieldType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_values_for_asset`

- Go wrapper: `client.CustomFields.GetCustomFieldValuesForAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_values_for_asset", opts, out)`
- HTTP route: `GET /custom-fields/values/for/asset/{AssetId}`
- Source controller: `Custom Fields`

Retrieve for Asset 

#### Retrieve Custom Fields for a specific Asset
#### Sample request:
```
GET /api/v1.0/custom-fields/values/for/asset/4c35d2ef-2bb1-457f-b123-dea74ee826ba
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAssetCustomFieldValue` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAssetCustomFieldValue`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_values_for_ticket`

- Go wrapper: `client.CustomFields.GetCustomFieldValuesForTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_values_for_ticket", opts, out)`
- HTTP route: `GET /custom-fields/values/for/ticket/{TicketId}`
- Source controller: `Custom Fields`

Retrieve for Ticket 

#### Retrieve Custom Fields for a specific Ticket
#### Sample request:
```
GET /api/v1.0/custom-fields/values/for/ticket/9933a6e3-d98b-4c5c-be04-043cf23e838e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfTicketCustomFieldValue` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfTicketCustomFieldValue`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_field_values_for_user`

- Go wrapper: `client.CustomFields.GetCustomFieldValuesForUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_field_values_for_user", opts, out)`
- HTTP route: `GET /custom-fields/values/for/user/{UserId}`
- Source controller: `Custom Fields`

Retrieve for User 

#### Retrieve Custom Fields for a specific User
#### Sample request:
```
GET /api/v1.0/custom-fields/values/for/user/634d18fb-6c6e-44e6-a993-8444466910e2
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUserCustomFieldValue` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUserCustomFieldValue`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_fields`

- Go wrapper: `client.CustomFields.GetCustomFields(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_fields", opts, out)`
- HTTP route: `GET /custom-fields`
- Source controller: `Custom Fields`
- Aliases: `list`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["ProductFilterType"]` | `ProductFilterType` | `query` | `no` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfCustomField` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfCustomField`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_custom_fields2`

- Go wrapper: `client.CustomFields.GetCustomFields2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "get_custom_fields2", opts, out)`
- HTTP route: `POST /custom-fields`
- Source controller: `Custom Fields`
- Aliases: `create`

Get Custom Fields

#### Retrieves a list of custom fields that can then be used to retrieve a specific custom field via _Get Custom Field_
#### Sample request:
```
POST /api/v1.0/custom-fields
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `CustomFieldsRequest` | `body` | `yes` | `GetCustomFieldsRequest` | `GetCustomFieldsRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfCustomField` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfCustomField`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_custom_field`

- Go wrapper: `client.CustomFields.UpdateCustomField(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "update_custom_field", opts, out)`
- HTTP route: `POST /custom-fields/{CustomFieldId}`
- Source controller: `Custom Fields`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldId"]` | `CustomFieldId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `CustomField` | `body` | `yes` | `CustomField` | `CustomField` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfCustomField` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfCustomField`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_custom_field_type`

- Go wrapper: `client.CustomFields.UpdateCustomFieldType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "custom_fields", "update_custom_field_type", opts, out)`
- HTTP route: `POST /custom-fields/types/{CustomFieldTypeId}`
- Source controller: `Custom Fields`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["CustomFieldTypeId"]` | `CustomFieldTypeId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `CustomFieldType` | `body` | `yes` | `CustomFieldType` | `CustomFieldType` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfCustomFieldType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfCustomFieldType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
