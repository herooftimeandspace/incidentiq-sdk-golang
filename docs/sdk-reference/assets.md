# `assets` Golden Namespace

Go client access: `client.Assets`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_asset` | `DELETE /assets/{AssetId}` |
| `get` | `get_asset` | `GET /assets/{AssetId}` |
| `update` | `update_asset` | `POST /assets/{AssetId}` |

## Methods

### `add_manufacturer_to_site`

- Go wrapper: `client.Assets.AddManufacturerToSite(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "add_manufacturer_to_site", opts, out)`
- HTTP route: `POST /assets/manufacturers/{ManufacturerId}/site`
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

### `add_manufacturer_to_site2`

- Go wrapper: `client.Assets.AddManufacturerToSite2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "add_manufacturer_to_site2", opts, out)`
- HTTP route: `POST /assets/manufacturers/{ManufacturerId}/site/{IncludeAllModels}`
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

### `add_user_favorite_asset`

- Go wrapper: `client.Assets.AddUserFavoriteAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "add_user_favorite_asset", opts, out)`
- HTTP route: `POST /assets/favorites/{AssetId}/{UserId}`
- Source controller: `Assets`

Mark as favorite

#### Marks a given asset for a user as a favorite
#### Sample request:
```
POST /api/v1.0/assets/favorites/ac6cece8-e4f4-e511-a789-005056bb000e/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | Asset ID to mark as favorite |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID of the user to be linked to the provided asset |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `create_asset_status_type`

- Go wrapper: `client.Assets.CreateAssetStatusType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "create_asset_status_type", opts, out)`
- HTTP route: `POST /assets/status/types/new`
- Source controller: `Assets`

Create a new asset status type

#### Creates a new Asset Status Type with provided attributes.  A AssetStatusID, as well a complete record of the updated Asset Status Type is returned for every successful creation of Asset Status Types.
#### Sample request:
```
POST /api/v1.0/assets/status/types/new
{
  "Update":{
		"AssetStatusTypeId":"a102cced-419d-4102-aadf-461f1e96b07b"
		"Name":"Example New Status Type",
		"IsRetired":false
		}
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `Item` | `body` | `yes` | `UpdateAssetStatusTypeRequest` | `UpdateAssetStatusTypeRequest` | Asset Status Type to be updated, including all attributes |

#### Returns

- Go wrapper return: `error`; decoded `ItemCreateResponseOfAssetStatusType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemCreateResponseOfAssetStatusType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_asset`

- Go wrapper: `client.Assets.DeleteAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "delete_asset", opts, out)`
- HTTP route: `DELETE /assets/{AssetId}`
- Source controller: `Assets`
- Aliases: `delete`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_asset_funding_type`

- Go wrapper: `client.Assets.DeleteAssetFundingType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "delete_asset_funding_type", opts, out)`
- HTTP route: `DELETE /assets/funding/types/{AssetFundingTypeId}`
- Source controller: `Assets`

Delete Asset Funding Type 

#### Delete an Asset Funding Type
#### Sample request:
```
DELETE /api/v1.0/assets/funding/types/a443fec7-45fe-4d08-8daf-11f4dd86df79
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetFundingTypeId"]` | `AssetFundingTypeId` | `path` | `yes` | `string` | `-` | Asset Funding Type Id to be removed |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_asset_status_type`

- Go wrapper: `client.Assets.DeleteAssetStatusType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "delete_asset_status_type", opts, out)`
- HTTP route: `DELETE /assets/status/types/{AssetStatusTypeId}`
- Source controller: `Assets`

Delete Asset Status Type

#### Delete an asset status type.
#### Sample request:
```
DELETE /api/v1.0/assets/status/types/a102cced-419d-4102-aadf-461f1e96b07b
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetStatusTypeId"]` | `AssetStatusTypeId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_manufacturer`

- Go wrapper: `client.Assets.DeleteManufacturer(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "delete_manufacturer", opts, out)`
- HTTP route: `DELETE /assets/manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`

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

### `get_asset`

- Go wrapper: `client.Assets.GetAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset", opts, out)`
- HTTP route: `GET /assets/{AssetId}`
- Source controller: `Assets`
- Aliases: `get`

Get an asset

#### Get information related to a given asset.  Properties and meta information for the asset is returned.  For a listing of associated tickets call GET `/api/v1.0/tickets`.  For timeline / historical data call GET `/api/v1.0/assets/{AssetId:guid}/activities`
#### Sample request:
```
GET /api/v1.0/assets/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | Asset ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_activities`

- Go wrapper: `client.Assets.GetAssetActivities(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_activities", opts, out)`
- HTTP route: `GET /assets/{AssetId}/activities`
- Source controller: `Assets`

Get asset change history

#### Gets timeline / change history information related to a given asset.  Timestamps and associated data related to the activity or update are included.  For a listing of associated tickets call GET `/api/v1.0/tickets`.  For high-level asset information call GET `/api/v1.0/assets/{AssetId:guid}`.
#### Sample request:
```
GET /api/v1.0/assets/ac6cece8-e4f4-e511-a789-005056bb000e/activities
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | Asset ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfActivity` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfActivity`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_favorites`

- Go wrapper: `client.Assets.GetAssetFavorites(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_favorites", opts, out)`
- HTTP route: `GET /assets/favorites/{UserId}`
- Source controller: `Assets`

Get user favorites

#### Performs a query to obtain assets favorited by a given user.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/favorites/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID to use when searching assets |
| `Params["All"]` | `All` | `query` | `no` | `any` | `-` | Include assets within the provided user's assigned location rooms |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_favorites2`

- Go wrapper: `client.Assets.GetAssetFavorites2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_favorites2", opts, out)`
- HTTP route: `GET /assets/favorites/{UserId}/{All}`
- Source controller: `Assets`

Get user favorites

#### Performs a query to obtain assets favorited by a given user.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/favorites/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID to use when searching assets |
| `PathParams["All"]` | `All` | `path` | `yes` | `string` | `-` | Include assets within the provided user's assigned location rooms |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_funding_type`

- Go wrapper: `client.Assets.GetAssetFundingType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_funding_type", opts, out)`
- HTTP route: `GET /assets/funding/types/{AssetFundingTypeId}`
- Source controller: `Assets`

Get Asset Funding Type

#### Retrieve a specific asset funding type by AssetFundingTypeId
#### Sample request:
```
GET /api/v1.0/assets/funding/types/a443fec7-45fe-4d08-8daf-11f4dd86df79
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetFundingTypeId"]` | `AssetFundingTypeId` | `path` | `yes` | `string` | `-` | Asset Funding Type Id to be retrieved, including all attributes |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfAssetFundingType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfAssetFundingType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_funding_types`

- Go wrapper: `client.Assets.GetAssetFundingTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_funding_types", opts, out)`
- HTTP route: `GET /assets/funding/types`
- Source controller: `Assets`

Get Asset Funding Types

#### Retrieves a list of asset funding types. A specific location type via GET `api/v1.0/assets/funding/types/a443fec7-45fe-4d08-8daf-11f4dd86df79`.
#### Sample request:
```
GET /api/v1.0/assets/funding/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAssetFundingType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAssetFundingType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_funding_types2`

- Go wrapper: `client.Assets.GetAssetFundingTypes2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_funding_types2", opts, out)`
- HTTP route: `POST /assets/funding/types`
- Source controller: `Assets`

Get Asset Funding Types

#### Retrieves a list of asset funding types. A specific location type via GET `api/v1.0/assets/funding/types/a443fec7-45fe-4d08-8daf-11f4dd86df79`.
#### Sample request:
```
GET /api/v1.0/assets/funding/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `r` | `body` | `yes` | `GetAssetFundingTypesRequest` | `GetAssetFundingTypesRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAssetFundingType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAssetFundingType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_status_type`

- Go wrapper: `client.Assets.GetAssetStatusType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_status_type", opts, out)`
- HTTP route: `GET /assets/status/types/{AssetStatusTypeId}`
- Source controller: `Assets`

Get Asset Status Type

#### Retrieve a specific asset status type based on supplied AssetStatusTypeId
#### Sample request:
```
GET /api/v1.0/assets/status/types/a102cced-419d-4102-aadf-461f1e96b07b
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetStatusTypeId"]` | `AssetStatusTypeId` | `path` | `yes` | `string` | `-` | - |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfAssetStatusType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfAssetStatusType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_status_types`

- Go wrapper: `client.Assets.GetAssetStatusTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_status_types", opts, out)`
- HTTP route: `GET /assets/status/types`
- Source controller: `Assets`

Get Asset Status Types

#### Retrieves a list of asset status types. A specific asset status type can be retrieved via GET `api/v1.0/assets/status/types/{AssetStatusTypeId:guid}`.
#### Sample request:
```
GET /api/v1.0/assets/status/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["r"]` | `r` | `query` | `yes` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAssetStatusType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAssetStatusType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_asset_status_types2`

- Go wrapper: `client.Assets.GetAssetStatusTypes2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_asset_status_types2", opts, out)`
- HTTP route: `POST /assets/status/types`
- Source controller: `Assets`

Get Asset Status Types

#### Retrieves a list of asset status types. A specific asset status type can be retrieved via GET `api/v1.0/assets/status/types/{AssetStatusTypeId:guid}`.
#### Sample request:
```
GET /api/v1.0/assets/status/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `r` | `body` | `yes` | `GetAssetStatusTypesRequest` | `GetAssetStatusTypesRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAssetStatusType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAssetStatusType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets`

- Go wrapper: `client.Assets.GetAssets(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets", opts, out)`
- HTTP route: `POST /assets`
- Source controller: `Assets`

Get assets by filter

#### Performs a query to obtain assets which match the provided filter conditions.  This allows for a very flexible search as filter conditions may be expressed in AND / OR groups to facilitate the query.  This API call returns full asset details.  To obtain high-level asset counts for a given query / filter set call POST `/api/v1.0/assets/counts` and provide the same filters.
#### Sample request:
```
POST /api/v1.0/assets
{
   "Filters":[
      {
         "Facet":"location",
         "Id":"ca6ffef3-cb32-40cb-a62c-7b1068b5cc21",
         "Negative":false
      },
      {
         "Facet":"AssetType",
         "Id":"2a1561e5-34ff-4fcf-87de-2a146f0e1c01"
      }
   ]
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `RequestInfo` | `body` | `yes` | `GetAssetsRequest` | `GetAssetsRequest` | Search parameters including optional filter definitions |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_by_asset_status_type`

- Go wrapper: `client.Assets.GetAssetsByAssetStatusType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_by_asset_status_type", opts, out)`
- HTTP route: `GET /assets/assetstatustype/{AssetStatusTypeId}`
- Source controller: `Assets`

Query by status

#### Queries / searches for assets which match a given status.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/assetstatustype/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetStatusTypeId"]` | `AssetStatusTypeId` | `path` | `yes` | `string` | `-` | Asset status ID to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_by_asset_tag`

- Go wrapper: `client.Assets.GetAssetsByAssetTag(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_by_asset_tag", opts, out)`
- HTTP route: `GET /assets/assettag/{AssetTag}`
- Source controller: `Assets`

Query by asset tag ( exact )

#### Queries / searches for assets which exactly match a given asset tag.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/assettag/100345
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetTag"]` | `AssetTag` | `path` | `yes` | `string` | `-` | Asset tag to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_by_location_room`

- Go wrapper: `client.Assets.GetAssetsByLocationRoom(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_by_location_room", opts, out)`
- HTTP route: `GET /assets/rooms/{LocationRoomId}`
- Source controller: `Assets`

Query by room

#### Queries / searches for assets which match a given location room.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/rooms/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationRoomId"]` | `LocationRoomId` | `path` | `yes` | `string` | `-` | Room ID to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_by_serial`

- Go wrapper: `client.Assets.GetAssetsBySerial(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_by_serial", opts, out)`
- HTTP route: `GET /assets/serial/{Serial}`
- Source controller: `Assets`

Query by serial number ( exact )

#### Queries / searches for assets which exactly match a given serial number.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/serial/MT500-234A1000-4300
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["Serial"]` | `Serial` | `path` | `yes` | `string` | `-` | Serial number to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_by_storage_unit_number`

- Go wrapper: `client.Assets.GetAssetsByStorageUnitNumber(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_by_storage_unit_number", opts, out)`
- HTTP route: `GET /assets/storageunit/{LocationId}/{StorageUnitNumber}`
- Source controller: `Assets`

Query by storage unit

#### Queries / searches for assets at a given storage unit number.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/storageunit/ac6cece8-e4f4-e511-a789-005056bb000e/1050
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | Location ID of the location used to search |
| `PathParams["StorageUnitNumber"]` | `StorageUnitNumber` | `path` | `yes` | `string` | `-` | Storage unit number to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_assets_count`

- Go wrapper: `client.Assets.GetAssetsCount(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_assets_count", opts, out)`
- HTTP route: `POST /assets/count`
- Source controller: `Assets`

Get counts by filter

#### Performs a query to obtain asset counts which match the provided filter conditions.  This allows for a very flexible search as filter conditions may be expressed in AND / OR groups to facilitate the query.  This API call does not return assets and only returns high-level asset counts.  To return a list of assets which represent the counts returned call POST `/api/v1.0/assets` and provide the same filters.
#### Sample request:
```
POST /api/v1.0/assets/count
{
   "Filters":[
      {
         "Facet":"location",
         "Id":"ca6ffef3-cb32-40cb-a62c-7b1068b5cc21",
         "Negative":false
      },
      {
         "Facet":"AssetType",
         "Id":"2a1561e5-34ff-4fcf-87de-2a146f0e1c01"
      }
   ]
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `RequestInfo` | `body` | `yes` | `GetAssetsRequest` | `GetAssetsRequest` | Search parameters including optional filter definitions |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_global_manufacturers`

- Go wrapper: `client.Assets.GetGlobalManufacturers(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_global_manufacturers", opts, out)`
- HTTP route: `GET /assets/manufacturers/global`
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

### `get_global_manufacturers2`

- Go wrapper: `client.Assets.GetGlobalManufacturers2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_global_manufacturers2", opts, out)`
- HTTP route: `POST /assets/manufacturers/global`
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

### `get_manufacturer`

- Go wrapper: `client.Assets.GetManufacturer(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_manufacturer", opts, out)`
- HTTP route: `GET /assets/manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`

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

### `get_spare_assets_by_asset_tag`

- Go wrapper: `client.Assets.GetSpareAssetsByAssetTag(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_spare_assets_by_asset_tag", opts, out)`
- HTTP route: `GET /assets/spares/assettag/{AssetTag}`
- Source controller: `Assets`

Query spares by asset tag

#### Queries / searches only for spare assets which exactly match a given asset tag.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/assettag/search/100345
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetTag"]` | `AssetTag` | `path` | `yes` | `string` | `-` | Asset tag to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfSpareAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfSpareAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_user_assets`

- Go wrapper: `client.Assets.GetUserAssets(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_user_assets", opts, out)`
- HTTP route: `GET /assets/for/{UserId}`
- Source controller: `Assets`

Query by user

#### Queries / searches for assets which match a given user.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/for/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID to search for |
| `Params["All"]` | `All` | `query` | `no` | `any` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_user_assets2`

- Go wrapper: `client.Assets.GetUserAssets2(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "get_user_assets2", opts, out)`
- HTTP route: `GET /assets/for/{UserId}/{All}`
- Source controller: `Assets`

Query by user

#### Queries / searches for assets which match a given user.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/for/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID to search for |
| `PathParams["All"]` | `All` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `remove_manufacturer_from_site`

- Go wrapper: `client.Assets.RemoveManufacturerFromSite(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "remove_manufacturer_from_site", opts, out)`
- HTTP route: `DELETE /assets/manufacturers/{ManufacturerId}/site`
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

### `remove_user_favorite_asset`

- Go wrapper: `client.Assets.RemoveUserFavoriteAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "remove_user_favorite_asset", opts, out)`
- HTTP route: `POST /assets/favorites/remove/{AssetId}/{UserId}`
- Source controller: `Assets`

Unmark as favorite

#### Unmarks a given asset for a user as a favorite
#### Sample request:
```
POST /api/v1.0/assets/favorites/remove/ac6cece8-e4f4-e511-a789-005056bb000e/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | Asset ID to unmark as favorite |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User ID of the user to be unlinked to the provided asset |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `search_assets_by_asset_tag`

- Go wrapper: `client.Assets.SearchAssetsByAssetTag(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "search_assets_by_asset_tag", opts, out)`
- HTTP route: `GET /assets/assettag/search/{AssetTag}`
- Source controller: `Assets`

Query by asset tag ( wildcard )

#### Queries / searches for assets which loosely match a given asset tag by means of a "contains text" search.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/assettag/search/100345
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetTag"]` | `AssetTag` | `path` | `yes` | `string` | `-` | Asset tag to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `search_assets_by_serial`

- Go wrapper: `client.Assets.SearchAssetsBySerial(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "search_assets_by_serial", opts, out)`
- HTTP route: `GET /assets/serial/search/{Serial}`
- Source controller: `Assets`

Query by serial number ( wildcard )

#### Queries / searches for assets which loosely match a given serial number by means of a "contains text" search.  Search is case-insensitive.  For a more generalized / flexible asset query call POST `/api/v1.0/assets` to search by filter.
#### Sample request:
```
GET /api/v1.0/assets/serial/MT500-234A1000-4300
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["Serial"]` | `Serial` | `path` | `yes` | `string` | `-` | Serial number to search for |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_asset`

- Go wrapper: `client.Assets.UpdateAsset(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "update_asset", opts, out)`
- HTTP route: `POST /assets/{AssetId}`
- Source controller: `Assets`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetId"]` | `AssetId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Asset` | `body` | `yes` | `UpdateAssetRequest` | `UpdateAssetRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_asset_funding_type`

- Go wrapper: `client.Assets.UpdateAssetFundingType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "update_asset_funding_type", opts, out)`
- HTTP route: `POST /assets/funding/types/{AssetFundingTypeId}`
- Source controller: `Assets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetFundingTypeId"]` | `AssetFundingTypeId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateAssetFundingTypeRequest` | `UpdateAssetFundingTypeRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_asset_status_type`

- Go wrapper: `client.Assets.UpdateAssetStatusType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "update_asset_status_type", opts, out)`
- HTTP route: `POST /assets/status/types/{AssetStatusTypeId}`
- Source controller: `Assets`

Update an existing asset status type

#### Updates a previously submitted Asset Status Type with provided changes.  A AssetStatusID, as well a complete record of the updated Asset Status Type is returned for every successful update Asset Status Type call made to the API.  
#### Sample request:
```
POST /api/v1.0/assets/status/types/a102cced-419d-4102-aadf-461f1e96b07b
{
  "Update":{
		"Name":"New Status Type",
		"IsRetired":false
		}
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["AssetStatusTypeId"]` | `AssetStatusTypeId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Item` | `body` | `yes` | `UpdateAssetStatusTypeRequest` | `UpdateAssetStatusTypeRequest` | Asset Status Type to be created, including all necessary attributes |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfGuid` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfGuid`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_manufacturer`

- Go wrapper: `client.Assets.UpdateManufacturer(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "assets", "update_manufacturer", opts, out)`
- HTTP route: `POST /assets/manufacturers/{ManufacturerId}`
- Source controller: `Manufacturers`

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
