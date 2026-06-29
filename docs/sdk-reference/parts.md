# `parts` Golden Namespace

Go client access: `client.Parts`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_part` | `DELETE /parts/{PartId}` |
| `get` | `get_part` | `GET /parts/{PartId}` |
| `list` | `get_parts` | `GET /parts` |
| `update` | `update_part` | `POST /parts/{PartId}` |

## Methods

### `delete_part`

- Go wrapper: `client.Parts.DeletePart(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "delete_part", opts, out)`
- HTTP route: `DELETE /parts/{PartId}`
- Source controller: `Parts`
- Aliases: `delete`

Delete a Part

#### Delete a specific Part
#### Sample request:
```
DELETE /api/v1.0/parts/c94f81dc-8fae-4e82-8014-b5e5b5e86575
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartId"]` | `PartId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_part_supplier`

- Go wrapper: `client.Parts.DeletePartSupplier(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "delete_part_supplier", opts, out)`
- HTTP route: `DELETE /parts/suppliers/{PartSupplierId}`
- Source controller: `Parts`

Delete a Part Supplier

#### Delete a specific part supplier.
#### Sample request:
```
DELETE /api/v1.0/parts/suppliers/e67018f1-3815-449c-a09f-e66dc6b83202
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartSupplierId"]` | `PartSupplierId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_part`

- Go wrapper: `client.Parts.GetPart(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "get_part", opts, out)`
- HTTP route: `GET /parts/{PartId}`
- Source controller: `Parts`
- Aliases: `get`

Get Part

#### Retrieve a specific Part based on PartID
#### Sample request:
```
GET /api/v1.0/parts/bbfdf941-7bbe-4cc2-a9d7-9b7fbeed2358
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartId"]` | `PartId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfPart` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfPart`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_part_supplier`

- Go wrapper: `client.Parts.GetPartSupplier(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "get_part_supplier", opts, out)`
- HTTP route: `GET /parts/suppliers/{PartSupplierId}`
- Source controller: `Parts`

Get Part Supplier

#### Retrieve a specific part supplier on PartSupplierId
#### Sample request:
```
GET /api/v1.0/parts/suppliers/0baa50cf-2037-43a4-9b64-391737f58d41
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartSupplierId"]` | `PartSupplierId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfPartSupplier` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfPartSupplier`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_part_suppliers`

- Go wrapper: `client.Parts.GetPartSuppliers(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "get_part_suppliers", opts, out)`
- HTTP route: `GET /parts/suppliers`
- Source controller: `Parts`

Get Part Suppliers

#### Retrieves a list of parts suppliers. A specific location type via GET `api/v1.0/parts/suppliers/{PartSupplierId:guid}`.
#### Sample request:
```
GET /api/v1.0/parts/suppliers
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfPartSupplier` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfPartSupplier`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_parts`

- Go wrapper: `client.Parts.GetParts(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "get_parts", opts, out)`
- HTTP route: `GET /parts`
- Source controller: `Parts`
- Aliases: `list`

Get Parts

#### Retrieves a list of parts. A specific part can be retrieved via GET `api/v1.0/parts/{PartId:guid}`.
#### Sample request:
```
GET /api/v1.0/parts
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfPart` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfPart`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_part`

- Go wrapper: `client.Parts.UpdatePart(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "update_part", opts, out)`
- HTTP route: `POST /parts/{PartId}`
- Source controller: `Parts`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartId"]` | `PartId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Part` | `body` | `yes` | `Part` | `Part` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfPart` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfPart`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_part_supplier`

- Go wrapper: `client.Parts.UpdatePartSupplier(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "parts", "update_part_supplier", opts, out)`
- HTTP route: `POST /parts/suppliers/{PartSupplierId}`
- Source controller: `Parts`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PartSupplierId"]` | `PartSupplierId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `PartSupplier` | `body` | `yes` | `PartSupplier` | `PartSupplier` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfPartSupplier` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfPartSupplier`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
