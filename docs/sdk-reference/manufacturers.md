# `manufacturers` Golden Namespace

Go client access: `client.Manufacturers`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_manufacturer2` | `DELETE /manufacturers/{ManufacturerId}` |
| `get` | `get_manufacturer2` | `GET /manufacturers/{ManufacturerId}` |
| `update` | `update_manufacturer2` | `POST /manufacturers/{ManufacturerId}` |

## Methods

### `add_manufacturer_to_site3`

- Go wrapper: `client.Manufacturers.AddManufacturerToSite3(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "add_manufacturer_to_site3", opts, out)`
- HTTP route: `POST /manufacturers/{ManufacturerId}/site`
- Source controller: `Manufacturers`

Remove Manufacturer

#### Remove a Manufacturer from a Site
#### Sample request:
```
POST /api/v1.0/manufacturers/70fe08d5-e67e-4495-8ac4-d92f734774af/site/true
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | Manufacturer Id to be added |
| `Params["IncludeAllModels"]` | `IncludeAllModels` | `query` | `no` | `bool` | `-` | (default false) Add all Models from this manufacturer to site |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `add_manufacturer_to_site4`

- Go wrapper: `client.Manufacturers.AddManufacturerToSite4(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "add_manufacturer_to_site4", opts, out)`
- HTTP route: `POST /manufacturers/{ManufacturerId}/site/{IncludeAllModels}`
- Source controller: `Manufacturers`

Remove Manufacturer

#### Remove a Manufacturer from a Site
#### Sample request:
```
POST /api/v1.0/manufacturers/70fe08d5-e67e-4495-8ac4-d92f734774af/site/true
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | Manufacturer Id to be added |
| `PathParams["IncludeAllModels"]` | `IncludeAllModels` | `path` | `yes` | `bool` | `-` | (default false) Add all Models from this manufacturer to site |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_manufacturer2`

- Go wrapper: `client.Manufacturers.DeleteManufacturer2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "delete_manufacturer2", opts, out)`
- HTTP route: `DELETE /manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`
- Aliases: `delete`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_global_manufacturers3`

- Go wrapper: `client.Manufacturers.GetGlobalManufacturers3(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "get_global_manufacturers3", opts, out)`
- HTTP route: `GET /manufacturers/global`
- Source controller: `Manufacturers`

Get Manufacturers

#### Retrieves a list of all manufacturers.
#### Sample request:
```
GET /api/v1.0/manufacturers/global
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfManufacturer` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfManufacturer`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_global_manufacturers4`

- Go wrapper: `client.Manufacturers.GetGlobalManufacturers4(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "get_global_manufacturers4", opts, out)`
- HTTP route: `POST /manufacturers/global`
- Source controller: `Manufacturers`

Get Manufacturers

#### Retrieves a list of all manufacturers.
#### Sample request:
```
GET /api/v1.0/manufacturers/global
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `r` | `body` | `yes` | `GetManufacturersRequest` | `GetManufacturersRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfManufacturer` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfManufacturer`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_manufacturer2`

- Go wrapper: `client.Manufacturers.GetManufacturer2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "get_manufacturer2", opts, out)`
- HTTP route: `GET /manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`
- Aliases: `get`

Get Manufacturer

#### Retrieve a specific manufacturer by Manufacturer Id
#### Sample request:
```
GET /api/v1.0/parts/manufacturers/70fe08d5-e67e-4495-8ac4-d92f734774af/site
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | Manufacturer ID to be retrieved |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfManufacturer` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfManufacturer`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `remove_manufacturer_from_site2`

- Go wrapper: `client.Manufacturers.RemoveManufacturerFromSite2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "remove_manufacturer_from_site2", opts, out)`
- HTTP route: `DELETE /manufacturers/{ManufacturerId}/site`
- Source controller: `Manufacturers`

Remove Manufacturer

#### Remove a Manufacturer
#### Sample request:
```
DELETE /api/v1.0/manufacturers/70fe08d5-e67e-4495-8ac4-d92f734774af/site
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | Manufacturer Id to be removed |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_manufacturer2`

- Go wrapper: `client.Manufacturers.UpdateManufacturer2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "manufacturers", "update_manufacturer2", opts, out)`
- HTTP route: `POST /manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ManufacturerId"]` | `ManufacturerId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateManufacturerRequest` | `UpdateManufacturerRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
