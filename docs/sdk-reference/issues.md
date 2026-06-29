# `issues` Golden Namespace

Go client access: `client.Issues`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_issue` | `DELETE /issues/{IssueId}` |
| `get` | `get_issue` | `GET /issues/{IssueId}` |
| `update` | `update_issue` | `POST /issues/{IssueId}` |

## Methods

### `delete_issue`

- Go wrapper: `client.Issues.DeleteIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "delete_issue", opts, out)`
- HTTP route: `DELETE /issues/{IssueId}`
- Source controller: `Issues`
- Aliases: `delete`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueId"]` | `IssueId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_issue_type`

- Go wrapper: `client.Issues.DeleteIssueType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "delete_issue_type", opts, out)`
- HTTP route: `DELETE /issues/types/{IssueTypeId}`
- Source controller: `Issues`

Delete Issue Type

#### Delete a specific Issue Type
#### Sample request:
```
DELETE /api/v1.0/issues/types/{IssueTypeId:guid}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueTypeId"]` | `IssueTypeId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_available_issues`

- Go wrapper: `client.Issues.GetAvailableIssues(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "get_available_issues", opts, out)`
- HTTP route: `GET /issues/site`
- Source controller: `Issues`

Get Available Issues

#### Retrieve all available issues for the given Site
#### Sample request:
```
GET /api/v1.0/issues/site
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfIssueRoles` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfIssueRoles`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_issue`

- Go wrapper: `client.Issues.GetIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "get_issue", opts, out)`
- HTTP route: `GET /issues/{IssueId}`
- Source controller: `Issues`
- Aliases: `get`

Get Issue 

#### Retrieve a specific Issue via Issue ID
#### Sample request:
```
GET /api/v1.0/issues/53bf1b92-a533-4a0e-8894-059f05b38f41
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueId"]` | `IssueId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfIssue` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfIssue`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_issue_type`

- Go wrapper: `client.Issues.GetIssueType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "get_issue_type", opts, out)`
- HTTP route: `GET /issues/types/{IssueTypeId}`
- Source controller: `Issues`

Get Issue Type

#### Retrieve a specific Issue Type via Issue Type Id
#### Sample request:
```
GET /api/v1.0/issues/types/fb08be5e-0876-4204-acf1-8f6aad44797c
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueTypeId"]` | `IssueTypeId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfIssueType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfIssueType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_issue_types`

- Go wrapper: `client.Issues.GetIssueTypes(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "get_issue_types", opts, out)`
- HTTP route: `POST /issues/types`
- Source controller: `Issues`

Get Issue Types

#### Retrieves a list of issue types that can then be used to retrieve a specific custom field via _GetIssueType_ via ID
#### Sample request:
```
POST /api/v1.0/issues/types
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `request` | `body` | `yes` | `GetIssueTypesRequest` | `GetIssueTypesRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfIssueType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfIssueType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_issue_types_simple`

- Go wrapper: `client.Issues.GetIssueTypesSimple(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "get_issue_types_simple", opts, out)`
- HTTP route: `GET /issues/types`
- Source controller: `Issues`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `Params["ApplySiteVisibility"]` | `ApplySiteVisibility` | `query` | `no` | `bool` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfIssueType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfIssueType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_issue`

- Go wrapper: `client.Issues.UpdateIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "update_issue", opts, out)`
- HTTP route: `POST /issues/{IssueId}`
- Source controller: `Issues`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueId"]` | `IssueId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Model` | `body` | `yes` | `Issue` | `Issue` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfIssue` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfIssue`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_issue_type`

- Go wrapper: `client.Issues.UpdateIssueType(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "issues", "update_issue_type", opts, out)`
- HTTP route: `POST /issues/types/{IssueTypeId}`
- Source controller: `Issues`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["IssueTypeId"]` | `IssueTypeId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Model` | `body` | `yes` | `IssueType` | `IssueType` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfIssueType` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfIssueType`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
