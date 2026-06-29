# `alerts` Golden Namespace

Go client access: `client.Alerts`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `create` | `queue_notification` | `POST /alerts/new` |

## Methods

### `queue_notification`

- Go wrapper: `client.Alerts.QueueNotification(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "alerts", "queue_notification", opts, out)`
- HTTP route: `POST /alerts/new`
- Source controller: `Emails & Notifications`
- Aliases: `create`

Send a new notification

#### Allows creation / sending of notifications within Incident IQ.  Notifications can be directed at users with varying severity.  Content may include rich HTML and links to other resources.
#### Sample request:
```
POST /api/v1.0/alerts/new
{
  "AlertKey": "Custom-Notification-1",
  "EntityTypeId": "ac6cece8-e4f4-e511-a789-005056bb000e",
  "EntityId": "ac6cece8-e4f4-e511-a789-005056bb000e",
  "Severity": 20,
  "Subject": "Test notification 1",
  "Details": "New notification content - click &lt;a href="#"&gt;here&lt;/a&gt;"
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `notification` | `body` | `yes` | `Alert` | `Alert` | Notification details / parameters |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
