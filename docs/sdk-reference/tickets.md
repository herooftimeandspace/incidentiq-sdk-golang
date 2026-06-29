# `tickets` Golden Namespace

Go client access: `client.Tickets`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `create` | `create_ticket` | `POST /tickets/new` |
| `delete` | `delete_ticket` | `DELETE /tickets/{TicketId}` |
| `get` | `get_ticket` | `GET /tickets/{TicketId}` |
| `update` | `update_ticket` | `POST /tickets/{TicketId}` |

## Methods

### `assign_ticket`

- Go wrapper: `client.Tickets.AssignTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "assign_ticket", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/sla`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `Sla` | `body` | `yes` | `Sla` | `Sla` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `cancel_ticket`

- Go wrapper: `client.Tickets.CancelTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "cancel_ticket", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/cancel`
- Source controller: `Tickets`

Cancel a ticket

#### Takes a CloseReasonTypeId from:

| CloseReasonTypeId  | Name			|
| :------------- | ----------:	|
| C98AE44E-ADA0-4AE7-AE30-817785C00001		| Canceled   |
| C98AE44E-ADA0-4AE7-AE30-817785C00002		| Resolved |
| C98AE44E-ADA0-4AE7-AE30-817785C00003		| Canceled - No Longer Occurring |
| C98AE44E-ADA0-4AE7-AE30-817785C00004		| Canceled - Duplicate Ticket |
| C98AE44E-ADA0-4AE7-AE30-817785C00005		| Resolved - No Issue Found |
| C98AE44E-ADA0-4AE7-AE30-817785C00006		| Resolved - No Response From Requestor |

#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/cancel
{
"CloseReasonTypeId":"C98AE44E-ADA0-4AE7-AE30-817785C00001",
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be canceling |
| `JSON` | `CancelOptions` | `body` | `yes` | `TicketCancelOptions` | `TicketCancelOptions` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `change_ticket_issue`

- Go wrapper: `client.Tickets.ChangeTicketIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "change_ticket_issue", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/issue`
- Source controller: `Tickets`

Change Issue

#### Specify a new Issue for an existing ticket   
#### Sample request:
```
POST /api/v1.0/tickets/EDE88D85-D6E4-E711-80C3-0003FF685BE7/issue
{
"TicketId":"EDE88D85-D6E4-E711-80C3-0003FF685BE7",
"IssueId":"DD66756C-6DA8-E711-80C2-0003FF001002"
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that will have the team assigned to it removed |
| `JSON` | `Request` | `body` | `yes` | `SetTicketIssueRequest` | `SetTicketIssueRequest` | Ticket Issue Request Object |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `change_ticket_to_requestor_responded`

- Go wrapper: `client.Tickets.ChangeTicketToRequestorResponded(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "change_ticket_to_requestor_responded", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/status/requestor-responded`
- Source controller: `Tickets`

Mark as 'requestor responded'

#### Change ticket status to 'requestor responded' and apply any applicable rules.
#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/status/requestor-responded
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be changing status of |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `change_ticket_to_waiting_on_requestor`

- Go wrapper: `client.Tickets.ChangeTicketToWaitingOnRequestor(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "change_ticket_to_waiting_on_requestor", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/status/waiting-on-requestor`
- Source controller: `Tickets`

Mark as 'waiting on requestor'

#### Change ticket status to 'waiting on requestor' and apply any applicable rules.
#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/status/waiting-on-requestor
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be changing status of |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `close_ticket`

- Go wrapper: `client.Tickets.CloseTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "close_ticket", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/close`
- Source controller: `Tickets`

Close a ticket

#### Takes a CloseReasonTypeId from:

| CloseReasonTypeId  | Name			|
| :------------- | ----------:	|
| C98AE44E-ADA0-4AE7-AE30-817785C00001		| Canceled   |
| C98AE44E-ADA0-4AE7-AE30-817785C00002		| Resolved |
| C98AE44E-ADA0-4AE7-AE30-817785C00003		| Canceled - No Longer Occurring |
| C98AE44E-ADA0-4AE7-AE30-817785C00004		| Canceled - Duplicate Ticket |
| C98AE44E-ADA0-4AE7-AE30-817785C00005		| Resolved - No Issue Found |
| C98AE44E-ADA0-4AE7-AE30-817785C00006		| Resolved - No Response From Requestor |

#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/close
{
"CloseReasonTypeId":"C98AE44E-ADA0-4AE7-AE30-817785C00001",
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be closing |
| `JSON` | `CloseOptions` | `body` | `yes` | `TicketCloseOptions` | `TicketCloseOptions` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `confirm_ticket_issue`

- Go wrapper: `client.Tickets.ConfirmTicketIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "confirm_ticket_issue", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/confirm-issue`
- Source controller: `Tickets`

Confirm a tickets issue

#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/confirm-issue
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be confirming issue of |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `copy_ticket`

- Go wrapper: `client.Tickets.CopyTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "copy_ticket", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/copy`
- Source controller: `Tickets`

Copy Ticket

#### Create a copy of a given ticket.    
#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/copy
{
"CopyRequestor":true,
"CopyPriority":true,
"CopyAssignedTo":true,
"CopyCustomFieldValues":true,
"RunRules":false,
"SendEmails":false,
"CreateWorkpackage":false,
"WorkpackageName":"New Name",
"Quantity":1
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `CopyOptions` | `body` | `yes` | `CopyTicketOptions` | `CopyTicketOptions` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `create_ticket`

- Go wrapper: `client.Tickets.CreateTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "create_ticket", opts, out)`
- HTTP route: `POST /tickets/new`
- Source controller: `Tickets`
- Aliases: `create`

Submit a new ticket

#### Creates a newly submitted ticket to the Incident IQ platform.  Based on inputs the new ticket may have additional rules fire which may subsequently update the new ticket.  A ticket ID, as well a complete record of the ticket, is returned for every successful new ticket call made to the API.  Please note that rules fire asynchronously after ticket submission so ticket data initially returned may subquently be updated.  To obtain the latest data using the provided ticket ID, call GET `/api/v1.0/tickets/{TicketId}`. For more flexible searching of tickets in general call POST `/api/v1.0/tickets` and provide filters for your search criteria.
#### Sample request:
```
POST /api/v1.0/tickets/new
{
   "Assets":[
      {
         "AssetId":"c8aef633-a009-48b1-9187-15621defbba8"
      }
   ],
   "IsUrgent":false,
   "TicketFollowers":null,
   "Locations":null,
   "Users":null,
   "Teams":null,
   "Roles":null,
   "AssetGroups":null,
   "IssueDescription":"Issue description",
   "CustomFieldValues":[
      {
         "SiteId":"bb6cece8-e4f4-e511-a789-005056bb000e",
         "CustomFieldId":"4112aa9a-db65-4e49-ad67-a5423010b344",
         "Scope":"Site",
         "CustomFieldTypeId":"e7c8b30d-d899-4193-a70b-21994d2a8089",
         "EntityTypeId":"888891ac-91aa-e711-80c2-100dffa00001",
         "IsRequired":false,
         "IsReadOnly":false,
         "IsHidden":false,
         "Filters":[
            {
               "FilterId":"888861e5-34ff-4fcf-87de-100d6f0e0043",
               "Exclude":false,
               "Value":"d5d91f20-2269-e611-80f1-000c29ab80b0"
            }
         ],
         "DisplayOrder":39,
         "EditorTypeId":12,
         "Value":"",
         "IsValid":true
      }
   ],
   "LocationId":"1fd1574a-2243-e611-a93b-a45e60e682da",
   "Attachments":[
            
   ],
   "LocationDetails":"Room details",
   "HasSensitiveInformation":false,
   "Subject":"Acer R13 CB5-312T - Application / Operating System &gt; Can\\'t access website",
   "TicketWizardCategoryId":"d5d91f20-2269-e611-80f1-000c29ab80b0",
   "ForId":"4dd6574a-2243-e611-a93b-a45e60e682da",
   "LocationRoomId":"d1186c41-7e78-e611-80f1-000c29ab80b0",
   "IssueCategoryId":"3cf36201-f97d-40e0-a713-c6cfac2c0b6e",
   "IssueId":"927f2fbb-4db7-e611-80ff-000c29ab80b0",
   "IssueTypeId":"c8add287-0ca2-4fcc-b027-8d0980a84519",
   "IsTraining":false,
   "ProductId":"88df910c-91aa-e711-80c2-0004ffa00010"
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `Ticket` | `body` | `yes` | `UpdateTicketRequest` | `UpdateTicketRequest` | Ticket detail parameters including optional associated entities such as assets or attachments |

#### Returns

- Go wrapper return: `error`; decoded `ItemCreateResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemCreateResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `delete_ticket`

- Go wrapper: `client.Tickets.DeleteTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "delete_ticket", opts, out)`
- HTTP route: `DELETE /tickets/{TicketId}`
- Source controller: `Tickets`
- Aliases: `delete`

Delete a ticket

#### Deletes an existing ticket
#### Sample request:
```
DELETE /api/v1.0/tickets/9F326CEF-8390-4707-BA2F-EBD7A397AD65
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be changing status of |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_ticket`

- Go wrapper: `client.Tickets.GetTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "get_ticket", opts, out)`
- HTTP route: `GET /tickets/{TicketId}`
- Source controller: `Tickets`
- Aliases: `get`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_ticket_assets`

- Go wrapper: `client.Tickets.GetTicketAssets(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "get_ticket_assets", opts, out)`
- HTTP route: `GET /tickets/{TicketId}/assets`
- Source controller: `Tickets`

Get assets for ticket

#### Get all assets currently linked to a given ticket.  Properties and meta information for each asset is returned.  To obtain the latest high-level ticket data provided a valid ticket ID, call GET `/api/v1.0/tickets/{TicketId}`.
#### Sample request:
```
GET /api/v1.0/tickets/ac6cece8-e4f4-e511-a789-005056bb000e/assets
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the ticket to update |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfTicketAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfTicketAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_ticket_sla`

- Go wrapper: `client.Tickets.GetTicketSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "get_ticket_sla", opts, out)`
- HTTP route: `GET /tickets/{TicketId}/sla`
- Source controller: `Tickets`

Get SLA for ticket

#### Get the current service level agreement ( SLA ) linked to a given ticket.  Properties and meta information for the SLA is returned.  To obtain the latest high-level ticket data provided a valid ticket ID, call GET `/api/v1.0/tickets/{TicketId}`.
#### Sample request:
```
GET /api/v1.0/tickets/ac6cece8-e4f4-e511-a789-005056bb000e/sla
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the ticket to update |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfTicketSla` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfTicketSla`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_ticket_statuses`

- Go wrapper: `client.Tickets.GetTicketStatuses(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "get_ticket_statuses", opts, out)`
- HTTP route: `GET /tickets/statuses`
- Source controller: `Tickets`

Retrieve list of ticket status

#### Returns a list of all ticket statuses available 
#### Sample request:
```
PUT /api/v1.0/tickets/statuses
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfTicketStatus` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfTicketStatus`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_ticket_as_duplicate`

- Go wrapper: `client.Tickets.MarkTicketAsDuplicate(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "mark_ticket_as_duplicate", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/mark-as-duplicate/{OriginalTicketId}`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |
| `PathParams["OriginalTicketId"]` | `OriginalTicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_ticket_not_sensitive`

- Go wrapper: `client.Tickets.MarkTicketNotSensitive(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "mark_ticket_not_sensitive", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/mark-not-sensitive`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_ticket_not_urgent`

- Go wrapper: `client.Tickets.MarkTicketNotUrgent(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "mark_ticket_not_urgent", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/mark-not-urgent`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_ticket_sensitive`

- Go wrapper: `client.Tickets.MarkTicketSensitive(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "mark_ticket_sensitive", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/mark-sensitive`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_ticket_urgent`

- Go wrapper: `client.Tickets.MarkTicketUrgent(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "mark_ticket_urgent", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/mark-urgent`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `un_assign_ticket_from_team`

- Go wrapper: `client.Tickets.UnAssignTicketFromTeam(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "un_assign_ticket_from_team", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/unassign/team`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `un_assign_ticket_from_user`

- Go wrapper: `client.Tickets.UnAssignTicketFromUser(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "un_assign_ticket_from_user", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/unassign/user`
- Source controller: `Tickets`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `un_assign_ticket_sla`

- Go wrapper: `client.Tickets.UnAssignTicketSla(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "un_assign_ticket_sla", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/unassign-sla`
- Source controller: `Tickets`

Remove SLA

#### Removes the currently assigned SLA from a ticket
#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/unassign-sla
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket IDof the ticket that will have the SLA removed |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `un_confirm_ticket_issue`

- Go wrapper: `client.Tickets.UnConfirmTicketIssue(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "un_confirm_ticket_issue", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/unconfirm-issue`
- Source controller: `Tickets`

Reverse ticket issue confirmation

#### Will reverse an already confirmed issue on a ticket
#### Sample request:
```
POST /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/unconfirm-issue
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be reversing the confirm of |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `undelete_ticket`

- Go wrapper: `client.Tickets.UndeleteTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "undelete_ticket", opts, out)`
- HTTP route: `PUT /tickets/{TicketId}/undelete`
- Source controller: `Tickets`

Undelete a ticket

#### Undeletes an existing deleted ticket
#### Sample request:
```
PUT /api/v1.0/tickets/c8aef633-a009-48b1-9187-15621defbba8/undelete
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we will be changing status of |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_ticket`

- Go wrapper: `client.Tickets.UpdateTicket(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "update_ticket", opts, out)`
- HTTP route: `POST /tickets/{TicketId}`
- Source controller: `Tickets`
- Aliases: `update`

Update an existing ticket

#### Updates a previously submitted ticket with provided changes.  Please note that based on inputs the updated ticket may have additional update rules fire which may affect final values.  A ticket ID, as well a complete record of the ticket, is returned for every successful update ticket call made to the API.  Please note that rules fire asynchronously so ticket data initially returned may subquently be updated.  To obtain the latest data using the provided ticket ID, call GET `/api/v1.0/tickets/{TicketId}`. For more flexible searching of tickets in general call POST `/api/v1.0/tickets` and provide filters for your search criteria.
#### Sample request:
```
POST /api/v1.0/tickets/ac6cece8-e4f4-e511-a789-005056bb000e
{
   "Assets":[
      {
         "AssetId":"c8aef633-a009-48b1-9187-15621defbba8"
      }
   ],
   "IsUrgent":false,
   "TicketFollowers":null,
   "Locations":null,
   "Users":null,
   "Teams":null,
   "Roles":null,
   "AssetGroups":null,
   "IssueDescription":"Issue description",
   "CustomFieldValues":[
      {
         "SiteId":"bb6cece8-e4f4-e511-a789-005056bb000e",
         "CustomFieldId":"4112aa9a-db65-4e49-ad67-a5423010b344",
         "Scope":"Site",
         "CustomFieldTypeId":"e7c8b30d-d899-4193-a70b-21994d2a8089",
         "EntityTypeId":"888891ac-91aa-e711-80c2-100dffa00001",
         "IsRequired":false,
         "IsReadOnly":false,
         "IsHidden":false,
         "Filters":[
            {
               "FilterId":"888861e5-34ff-4fcf-87de-100d6f0e0043",
               "Exclude":false,
               "Value":"d5d91f20-2269-e611-80f1-000c29ab80b0"
            }
         ],
         "DisplayOrder":39,
         "EditorTypeId":12,
         "Value":"",
         "IsValid":true
      }
   ],
   "LocationId":"1fd1574a-2243-e611-a93b-a45e60e682da",
   "Attachments":[
            
   ],
   "LocationDetails":"Room details",
   "HasSensitiveInformation":false,
   "Subject":"Acer R13 CB5-312T - Application / Operating System &gt; Can\\'t access website",
   "TicketWizardCategoryId":"d5d91f20-2269-e611-80f1-000c29ab80b0",
   "ForId":"4dd6574a-2243-e611-a93b-a45e60e682da",
   "LocationRoomId":"d1186c41-7e78-e611-80f1-000c29ab80b0",
   "IssueCategoryId":"3cf36201-f97d-40e0-a713-c6cfac2c0b6e",
   "IssueId":"927f2fbb-4db7-e611-80ff-000c29ab80b0",
   "IssueTypeId":"c8add287-0ca2-4fcc-b027-8d0980a84519",
   "IsTraining":false,
   "ProductId":"88df910c-91aa-e711-80c2-0004ffa00010"
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the ticket to update |
| `JSON` | `Ticket` | `body` | `yes` | `UpdateTicketRequest` | `UpdateTicketRequest` | Ticket detail parameters including optional associated entities such as assets or attachments |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfTicket` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfTicket`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_ticket_assets`

- Go wrapper: `client.Tickets.UpdateTicketAssets(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "update_ticket_assets", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/assets`
- Source controller: `Tickets`

Update assets

#### Updates a previously submitted ticket's assets with provided values.  Please note that based on inputs the updated ticket may have additional update rules fire which may affect final values.  Please note that rules fire asynchronously so ticket data initially returned may subquently be updated.  To obtain the latest data using the provided ticket ID, call GET `/api/v1.0/tickets/{TicketId}`.
#### Sample request:
```
POST /api/v1.0/tickets/ac6cece8-e4f4-e511-a789-005056bb000e
[
   {
      "AssetId":"4112aa9a-db65-4e49-ad67-a5423010b344"
   }
]
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the ticket to update |
| `JSON` | `Assets` | `body` | `yes` | `[]any` | `-` | Asset ID values for the ticket being updated |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfTicketAsset` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfTicketAsset`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_ticket_custom_fields`

- Go wrapper: `client.Tickets.UpdateTicketCustomFields(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "update_ticket_custom_fields", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/custom-fields`
- Source controller: `Tickets`

Update custom fields

#### Updates a previously submitted ticket's custom fields with provided values.  Please note that based on inputs the updated ticket may have additional update rules fire which may affect final values.  Please note that rules fire asynchronously so ticket data initially returned may subquently be updated.  To obtain the latest data using the provided ticket ID, call GET `/api/v1.0/tickets/{TicketId}`.
#### Sample request:
```
POST /api/v1.0/tickets/ac6cece8-e4f4-e511-a789-005056bb000e
[
   {
      "SiteId":"bb6cece8-e4f4-e511-a789-005056bb000e",
      "CustomFieldId":"4112aa9a-db65-4e49-ad67-a5423010b344",
      "Scope":"Site",
      "CustomFieldTypeId":"e7c8b30d-d899-4193-a70b-21994d2a8089",
      "EntityTypeId":"888891ac-91aa-e711-80c2-100dffa00001",
      "IsRequired":false,
      "IsReadOnly":false,
      "IsHidden":false,
      "Filters":[
         {
            "FilterId":"888861e5-34ff-4fcf-87de-100d6f0e0043",
            "Exclude":false,
            "Value":"d5d91f20-2269-e611-80f1-000c29ab80b0"
         }
      ],
      "DisplayOrder":39,
      "EditorTypeId":12,
      "Value":"",
      "IsValid":true
   }
]
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the ticket to update |
| `JSON` | `Values` | `body` | `yes` | `[]any` | `-` | Custom field values for the ticket being updated |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_ticket_subject`

- Go wrapper: `client.Tickets.UpdateTicketSubject(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "tickets", "update_ticket_subject", opts, out)`
- HTTP route: `POST /tickets/{TicketId}/subject`
- Source controller: `Tickets`

Update the subject of a ticket

#### Changes the Subject of an existing ticket
#### Sample request:
```
POST /api/v1.0/tickets/{TicketId:guid}/subject
{
   "TicketId":"c8aef633-a009-48b1-9187-15621defbba8",
   "Subject":"Example New Subject",
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID that we need to update |
| `JSON` | `Update` | `body` | `yes` | `UpdateTicketSubjectRequest` | `UpdateTicketSubjectRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
