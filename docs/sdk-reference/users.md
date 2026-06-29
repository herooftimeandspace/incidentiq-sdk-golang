# `users` Golden Namespace

Go client access: `client.Users`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_user` | `DELETE /users/{UserId}` |
| `get` | `get_user` | `GET /users/{UserId}` |
| `list` | `get_users_legacy` | `GET /users` |
| `update` | `update_user` | `POST /users/{UserId}` |

## Methods

### `create_local_user`

- Go wrapper: `client.Users.CreateLocalUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "create_local_user", opts, out)`
- HTTP route: `POST /users/local/new`
- Source controller: `Users`

Create User

#### Creates the user(s) with attributes as specified.
#### Will need to include both a LocationId which can be queried at `GET /api/v1.0/locations/all`.
#### Along with a RoleId:
| RoleId | Name |
| ----------- | ----------- |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E501 | Student |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E502 | Faculty |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E503 | Staff |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E504 | Agent |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E507 | Parent / Guardian |
#### Sample request:
```
POST /api/v1.0/users/local/new
{
"PreventProviderUpdates": false,
"DataMappings": {},
"FirstName": "John",
"LastName": "Smith",
"Email": "johnsmith@smith.com",
"SchoolIdNumber": "SMT12",
"Phone": "555-5555",
"Password": "T3STP@SS",
"PasswordExpirationDate": "09/01/2022",
"InternalComments": "lorem ipsum dias",
"Grade": "4",
"LocationId": "1fd1574a-2243-e611-a93b-a45e60e682da",
"RoleId": "6d5fee76-e05e-43c0-b8e9-b8447746e504",
"UpdateCustomFields": true,
"CustomFieldValues": []
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `User` | `body` | `yes` | `User` | `User` | The User to be created |

#### Returns

- Go wrapper return: `error`; decoded `ItemCreateResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemCreateResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_user`

- Go wrapper: `client.Users.DeleteUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "delete_user", opts, out)`
- HTTP route: `DELETE /users/{UserId}`
- Source controller: `Users`
- Aliases: `delete`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_user_view`

- Go wrapper: `client.Users.DeleteUserView(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "delete_user_view", opts, out)`
- HTTP route: `DELETE /users/views/{ViewId}`
- Source controller: `Views`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ViewId"]` | `ViewId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_agents`

- Go wrapper: `client.Users.GetAgents(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_agents", opts, out)`
- HTTP route: `POST /users/agents`
- Source controller: `Users`

Get Agents

#### Retrieves a list of agent based on filters set within the request object.
#### Sample request:
```
POST /api/v1.0/users/agents
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `RequestInfo` | `body` | `yes` | `GetUsersRequestV1` | `GetUsersRequestV1` | Request object containing filters |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_agents_legacy`

- Go wrapper: `client.Users.GetAgentsLegacy(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_agents_legacy", opts, out)`
- HTTP route: `GET /users/agents`
- Source controller: `Users`

No contract summary provided.

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_user`

- Go wrapper: `client.Users.GetUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_user", opts, out)`
- HTTP route: `GET /users/{UserId}`
- Source controller: `Users`
- Aliases: `get`

Get User

#### Retrieve a User's information.
#### Sample request:
```
GET /api/v1.0/users/0b92ee98-edef-4423-9e15-98fe5726c1af
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | User Id of the user details to be retrieved |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_user_view`

- Go wrapper: `client.Users.GetUserView(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_user_view", opts, out)`
- HTTP route: `GET /users/views/{ViewId}`
- Source controller: `Views`

Get User View

#### Retrieves details about a single user view in IIQ.   User lists can then be retrieved via GET `api/v1.0/users/view/{ViewId:guid}`.
#### Sample request:
```
GET /api/v1.0/users/views/{ViewId:guid}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ViewId"]` | `ViewId` | `path` | `yes` | `string` | `-` | View Id which details are to be retrieved |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfView` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfView`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_user_views`

- Go wrapper: `client.Users.GetUserViews(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_user_views", opts, out)`
- HTTP route: `GET /users/views`
- Source controller: `Views`

Get User View List

#### Retrieves a list of all user views in IIQ.   User lists can then be retrieved via GET `api/v1.0/users/view/{ViewId:guid}`.
#### Sample request:
```
GET /api/v1.0/users/views
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfView` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfView`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users`

- Go wrapper: `client.Users.GetUsers(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users", opts, out)`
- HTTP route: `POST /users`
- Source controller: `Users`

Get Users

#### Retrieves a list of users. A specific user can be retrieved via GET `api/v1.0/users/cb347afd-e235-43aa-832c-cddc89680e3c` with documentation under the _Get User_ section.
#### Sample request:
```
POST /api/v1.0/users
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `RequestInfo` | `body` | `yes` | `GetUsersRequestV1` | `GetUsersRequestV1` | Request object containing filters |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users_for_view`

- Go wrapper: `client.Users.GetUsersForView(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users_for_view", opts, out)`
- HTTP route: `POST /users/view/{ViewId}`
- Source controller: `Users`

Get Users for View

#### Retrieves a list of all users in IIQ that fit the criteria of the view requested. A list of applicable User Views can be retrieved via GET `/api/v1.0/users/views`. More documentation available under the _Views_ section.
#### Sample request:
```
POST /api/v1.0/users/view/4e6817c9-86b5-4810-84d8-ae99dfb6ba81
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ViewId"]` | `ViewId` | `path` | `yes` | `string` | `-` | The ViewId to use |
| `JSON` | `BodyFilters` | `body` | `yes` | `[]any` | `-` | Other filtration to use to filter the returned view's user list |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users_legacy`

- Go wrapper: `client.Users.GetUsersLegacy(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users_legacy", opts, out)`
- HTTP route: `GET /users`
- Source controller: `Users`
- Aliases: `list`

No contract summary provided.

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users_per_grade`

- Go wrapper: `client.Users.GetUsersPerGrade(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users_per_grade", opts, out)`
- HTTP route: `GET /users/stats/grades`
- Source controller: `Users`

Get User Counts Per Grade

#### Retrieve a list of grades with assigned user counts
#### Sample request:
```
GET /api/v1.0/users/stats/grades
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfGradeStat` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfGradeStat`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users_per_location`

- Go wrapper: `client.Users.GetUsersPerLocation(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users_per_location", opts, out)`
- HTTP route: `GET /users/stats/locations`
- Source controller: `Users`

Get User Counts Per Location

#### Retrieve a list of locations with assigned user counts
#### Sample request:
```
GET /api/v1.0/users/stats/locations
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfLocationStat` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfLocationStat`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_users_with_location_permission`

- Go wrapper: `client.Users.GetUsersWithLocationPermission(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "get_users_with_location_permission", opts, out)`
- HTTP route: `GET /users/location/{LocationId}`
- Source controller: `Users`

Get Location User Permissions

#### Retrieves a the user(s) requested along with their permissions for the locationId specified.  Locations can be queried at `GET /api/v1.0/locations/all`.  More details in the Locations section under _Get Locations_.
#### Sample request:
```
GET /api/v1.0/users/location/1e997655-cbce-4cc3-85a6-5420aa9c6206
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["LocationId"]` | `LocationId` | `path` | `yes` | `string` | `-` | The LocationId of the Location to be queried |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfUserWithLocationPermission` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfUserWithLocationPermission`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_user`

- Go wrapper: `client.Users.UpdateUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "update_user", opts, out)`
- HTTP route: `POST /users/{UserId}`
- Source controller: `Users`
- Aliases: `update`

Update User

#### Updates specified user with attributes as specified.
#### Will need to include both a LocationId which can be queried at `GET /api/v1.0/locations/all`.
#### Along with a RoleId:
| RoleId | Name |
| ----------- | ----------- |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E501 | Student |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E502 | Faculty |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E503 | Staff |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E504 | Agent |
| 6D5FEE76-E05E-43C0-B8E9-B8447746E507 | Parent / Guardian |
#### Sample request:
```
POST /api/v1.0/users/88c33434-34c5-4e03-968c-db99b7d6606b
UpdateUserRequest: {
"PreventProviderUpdates": false,
"DataMappings": {},
"FirstName": "John",
"LastName": "Smith",
"Email": "johnsmith@gmail.com",
"SchoolIdNumber": "SMT12",
"Phone": "555-5555",
"Password": "T3STP@SS",
"PasswordExpirationDate": "09/01/2022",
"InternalComments": "lorem ipsum dias",
"Grade": "4",
"LocationId": "1fd1574a-2243-e611-a93b-a45e60e682da",
"RoleId": "6d5fee76-e05e-43c0-b8e9-b8447746e504",
"UpdateCustomFields": true,
"CustomFieldValues": []
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["UserId"]` | `UserId` | `path` | `yes` | `string` | `-` | The UserId of the user to be updated |
| `JSON` | `User` | `body` | `yes` | `UpdateUserRequest` | `UpdateUserRequest` | The user details to be updated |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfUser` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfUser`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_user_view`

- Go wrapper: `client.Users.UpdateUserView(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "users", "update_user_view", opts, out)`
- HTTP route: `POST /users/views/{ViewId}`
- Source controller: `Views`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["ViewId"]` | `ViewId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `UserView` | `body` | `yes` | `View` | `View` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfView` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfView`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
