# `notifications` Golden Namespace

Go client access: `client.Notifications`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `create` | `get_notifications` | `POST /notifications` |

## Methods

### `get_notifications`

- Go wrapper: `client.Notifications.GetNotifications(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "get_notifications", opts, out)`
- HTTP route: `POST /notifications`
- Source controller: `Emails & Notifications`
- Aliases: `create`

Query notifications

#### Queries the system for any notifications based on submitted request parameters
#### Sample request:
```
POST /api/v1.0/notifications
{
  "IncludeRead": true,
  "IncludeArchived": true,
  "IncludeUnarchived": true
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `RequestInfo` | `body` | `yes` | `GetNotificationsRequest` | `GetNotificationsRequest` | - |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfNotification` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfNotification`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_ticket_emails`

- Go wrapper: `client.Notifications.GetTicketEmails(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "get_ticket_emails", opts, out)`
- HTTP route: `GET /notifications/emails/for/ticket/{TicketId}`
- Source controller: `Emails & Notifications`

Get emails for ticket

#### For a specific, single ticket query the system for all related emails
#### Sample request:
```
GET /api/v1.0/notifications/emails/for/ticket/ac6cece8-e4f4-e511-a789-005056bb000e
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["TicketId"]` | `TicketId` | `path` | `yes` | `string` | `-` | Ticket ID of the record to query |

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfEmail` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfEmail`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_unarchived_notifications`

- Go wrapper: `client.Notifications.GetUnarchivedNotifications(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "get_unarchived_notifications", opts, out)`
- HTTP route: `GET /notifications/unarchived`
- Source controller: `Emails & Notifications`

Get unarchived

#### Queries the system for any unarchived notifications
#### Sample request:
```
GET /api/v1.0/notifications/unarchived
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfNotification` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfNotification`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_unread_notifications`

- Go wrapper: `client.Notifications.GetUnreadNotifications(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "get_unread_notifications", opts, out)`
- HTTP route: `GET /notifications/unread`
- Source controller: `Emails & Notifications`

Get unread

#### Queries the system for any unread notifications
#### Sample request:
```
GET /api/v1.0/notifications/unread
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfNotification` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfNotification`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_all_notifications_archived`

- Go wrapper: `client.Notifications.MarkAllNotificationsArchived(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "mark_all_notifications_archived", opts, out)`
- HTTP route: `POST /notifications/all-archive`
- Source controller: `Emails & Notifications`

Mark all notification as archived

#### Marks all notifications that are not yet marked "archived" as "archived"
#### Sample request:
```
POST /api/v1.0/notifications/all-archive
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_all_notifications_read`

- Go wrapper: `client.Notifications.MarkAllNotificationsRead(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "mark_all_notifications_read", opts, out)`
- HTTP route: `POST /notifications/all-read`
- Source controller: `Emails & Notifications`

Mark all notification as read

#### Marks all notifications that are not yet marked "read" as "read"
#### Sample request:
```
POST /api/v1.0/notifications/all-read
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_notification_archived`

- Go wrapper: `client.Notifications.MarkNotificationArchived(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "mark_notification_archived", opts, out)`
- HTTP route: `POST /notifications/{NotificationId}/archive`
- Source controller: `Emails & Notifications`

Mark notification as archived

#### Marks a provided single notification as "archived"
#### Sample request:
```
POST /api/v1.0/notifications/ac6cece8-e4f4-e511-a789-005056bb000e/archive
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["NotificationId"]` | `NotificationId` | `path` | `yes` | `string` | `-` | Notification ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `mark_notification_read`

- Go wrapper: `client.Notifications.MarkNotificationRead(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "notifications", "mark_notification_read", opts, out)`
- HTTP route: `POST /notifications/{NotificationId}/read`
- Source controller: `Emails & Notifications`

Mark notification as read

#### Marks a provided single notification as "read"
#### Sample request:
```
POST /api/v1.0/notifications/ac6cece8-e4f4-e511-a789-005056bb000e/read
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["NotificationId"]` | `NotificationId` | `path` | `yes` | `string` | `-` | Notification ID of the record to modify |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
