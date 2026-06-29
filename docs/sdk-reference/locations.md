# `locations` Golden Namespace

Go client access: `client.Locations`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_location` | `DELETE /locations/{LocationId}` |
| `get` | `get_location` | `GET /locations/{LocationId}` |
| `update` | `update_location` | `POST /locations/{LocationId}` |

## Methods

### `delete_location`

- Go wrapper: `client.Locations.DeleteLocation(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "delete_location", opts, out)`
- HTTP route: `DELETE /locations/{LocationId}`
- Source controller: `Locations`
- Aliases: `delete`

Delete location

#### Delete a specific location
#### Sample request:
```
DELETE /api/v1.0/locations/d344d88d-d201-4c52-8d23-4371fa7179bb
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_location_room`

- Go wrapper: `client.Locations.DeleteLocationRoom(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "delete_location_room", opts, out)`
- HTTP route: `DELETE /locations/rooms/{LocationRoomId}`
- Source controller: `Locations`

Delete location room

#### Delete a specific location room
#### Sample request:
```
DELETE /api/v1.0/locations/rooms/dba9553e-c473-4863-b99f-2c5c54c73a1b
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationRoomId"]` | `LocationRoomId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_all_location_rooms`

- Go wrapper: `client.Locations.GetAllLocationRooms(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_all_location_rooms", opts, out)`
- HTTP route: `GET /locations/rooms`
- Source controller: `Locations`

Get All Location Rooms

#### Retrieve a list of rooms for all locations
#### Sample request:
```
GET /api/v1.0/locations/rooms
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfLocationRoom` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfLocationRoom`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_location`

- Go wrapper: `client.Locations.GetLocation(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_location", opts, out)`
- HTTP route: `GET /locations/{LocationId}`
- Source controller: `Locations`
- Aliases: `get`

Get Location

#### Retrieve a specific Location via Location ID
#### Sample request:
```
GET /api/v1.0/locations/4fc0dc90-c40a-4012-b4e2-224ca02bdfb7
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfLocation` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfLocation`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_location_room`

- Go wrapper: `client.Locations.GetLocationRoom(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_location_room", opts, out)`
- HTTP route: `GET /locations/rooms/{LocationRoomId}`
- Source controller: `Locations`

Get Location Room

#### Retrieve a specific Location Room via that Room's ID
#### Sample request:
```
GET /api/v1.0/locations/rooms/6983376d-8906-4c3b-95de-f252a11d9164
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationRoomId"]` | `LocationRoomId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfLocationRoom` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfLocationRoom`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_location_rooms`

- Go wrapper: `client.Locations.GetLocationRooms(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_location_rooms", opts, out)`
- HTTP route: `GET /locations/{LocationId}/rooms`
- Source controller: `Locations`

Get Location Room for Location

#### Retrieve a list of Location Rooms for a specific Location ID
#### Sample request:
```
GET /api/v1.0/locations/4fc0dc90-c40a-4012-b4e2-224ca02bdfb7/rooms
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfLocationRoom` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfLocationRoom`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_location_type`

- Go wrapper: `client.Locations.GetLocationType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_location_type", opts, out)`
- HTTP route: `GET /locations/types/{LocationTypeId}`
- Source controller: `Locations`

Get Location Type

#### Retrieve a specific Location Type via LocationTypeID
#### Sample request:
```
GET /api/v1.0/locations/types/27c81cbe-fdc4-4d7f-9c1c-e0311de5f882
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationTypeId"]` | `LocationTypeId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfLocationType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfLocationType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_location_types`

- Go wrapper: `client.Locations.GetLocationTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_location_types", opts, out)`
- HTTP route: `GET /locations/types`
- Source controller: `Locations`

Get Location Types

#### Retrieves a list of location types that can then be used to retrieve a specific location type via GET `api/locations/types/{LocationTypeId:guid}`.
#### Sample request:
```
GET /api/v1.0/locations/types
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfLocationType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfLocationType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_locations`

- Go wrapper: `client.Locations.GetLocations(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "get_locations", opts, out)`
- HTTP route: `GET /locations/all`
- Source controller: `Locations`

Get All Locations

#### Retrieve a list of all locations
#### Sample request:
```
GET /api/v1.0/locations/all
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfLocation` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfLocation`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_location`

- Go wrapper: `client.Locations.UpdateLocation(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "update_location", opts, out)`
- HTTP route: `POST /locations/{LocationId}`
- Source controller: `Locations`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Location` | `body` | `yes` | `Location` | `Location` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfLocation` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfLocation`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_location_room`

- Go wrapper: `client.Locations.UpdateLocationRoom(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "locations", "update_location_room", opts, out)`
- HTTP route: `POST /locations/rooms/{LocationRoomId}`
- Source controller: `Locations`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationRoomId"]` | `LocationRoomId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `LocationRoom` | `body` | `yes` | `LocationRoom` | `LocationRoom` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfLocationRoom` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfLocationRoom`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
