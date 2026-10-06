# User Room Associations

Incident IQ's web client exposes a Silver capability for reading and changing
the location rooms assigned to a user. These routes are not present in the
bundled Golden Stoplight contracts, so they remain under `client.Silver.Users`
and must not be presented as Golden `client.Users` methods.

## SDK methods and endpoints

| Operation | Go method | HTTP request | Request body |
| --- | --- | --- | --- |
| Read current associations | `GetUserRooms` | `GET /api/v1.0/users/{user_id}/rooms` | None; the generated Silver wrapper also requires the observed `$s` query parameter. |
| Add one association | `AddUserRoom` | `POST /api/v1.0/users/{user_id}/rooms/{location_room_id}` | None. |
| Replace all associations | `SetUserRooms` | `POST /api/v1.0/users/{user_id}/rooms` | A JSON array of location room ID strings, for example `["room-id-1", "room-id-2"]`. |
| Remove one association | `RemoveUserRoom` | `DELETE /api/v1.0/users/{user_id}/rooms/{location_room_id}` | None. |

The handwritten mutation helpers accept stable IDs as typed arguments and own
their route parameters. `AddUserRoom` and `RemoveUserRoom` reject any caller
body. `SetUserRooms` owns its JSON body and rejects caller-provided `JSON` or
`Body` values, so ordinary user fields cannot be sent accidentally.

All methods use the SDK's configured `Authorization`, `Client: ApiClient`, and,
when configured, `SiteId` headers. Configure `SiteID` when the target tenant
requires it. `SetUserRooms` sends `Content-Type: application/json` because it
has a JSON body; the bodyless add and remove requests do not require a content
type. To prevent an ambiguous POST from being replayed, these mutation helpers
do not use the usual Silver automatic retry-without-Client fallback. They send
`Client: ApiClient` by default. A caller may explicitly set
`OmitClientHeader` only when tenant-specific validation proves that the browser
header shape is required.

## Response contract

Incident IQ does not publish a Golden response schema for these routes. Any
successful JSON object or array is decoded into the caller-provided `out` value;
an empty successful response is also valid. Callers must not treat a particular
mutation response field as proof that the final room state is correct. A
non-2xx response is returned as `*incidentiq.APIError` with the bounded response
body and response headers.

## Safe single-association change

Read the current associations immediately before planning the write. Use the
stable `LocationRoomId`, not a room display name, and confirm that the target
room belongs to the expected site. For one-room changes, prefer `AddUserRoom`
or `RemoveUserRoom`; these routes cannot carry ordinary user profile fields and
therefore cannot overwrite names, roles, locations, custom fields, or other
unrelated user data.

```go
var before map[string]any
err := client.Silver.Users.GetUserRooms(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"user_id": userID},
	Params:     map[string]string{"$s": "100"},
}, &before)
if err != nil {
	return err
}

err = client.Silver.Users.AddUserRoom(
	ctx,
	userID,
	locationRoomID,
	incidentiq.RequestOptions{},
	nil,
)
if err != nil {
	return err
}

var after map[string]any
err = client.Silver.Users.GetUserRooms(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"user_id": userID},
	Params:     map[string]string{"$s": "100"},
}, &after)
if err != nil {
	return err
}
// Validate that the expected stable room ID is present exactly once before
// recording the operation as complete.
```

Use `SetUserRooms` only when the caller intentionally owns the complete desired
association set. It replaces the entire list, so a stale read can erase a room
added concurrently by another operator or integration. The helper rejects a
nil desired-room slice. A non-nil empty slice serializes as `[]` and is an
explicit request to remove every room association, so it should require the
caller's strongest destructive-write gate.

## Retry and idempotency expectations

`AddUserRoom` and `SetUserRooms` use POST and are not retried automatically by
the SDK, including no automatic retry without the `Client` header. If the
response is lost, read the room list again before deciding whether another
write is necessary. Do not blindly replay either POST.

`RemoveUserRoom` uses DELETE and follows the SDK's idempotent retry policy for
transient transport failures and retryable status codes. It does not use the
Silver retry-without-Client fallback. Even after a successful DELETE, perform a
read-after-write and confirm that the target stable room ID is absent.

The SDK does not create a provider idempotency key for these undocumented
routes. Workflow callers must reserve their own deterministic operation key,
record the write outcome, and make read-after-write verification part of the
same durable workflow. Production must not be the first environment where the
write path is exercised.

## Caller safety checklist

- Resolve the user, site, and room to stable provider IDs before the write.
- Fail closed on missing or ambiguous users, sites, or rooms.
- Read before writing and preserve associations the workflow does not own.
- Prefer the one-association add/remove methods over full replacement.
- Never send a full user object through `client.Users.UpdateUser` to change room
  assignments; room associations have separate Silver routes.
- Validate the full response status and then verify final state with
  `GetUserRooms`.
- Keep tokens, authorization headers, and raw sensitive provider payloads out
  of logs, audit events, fixtures, and error messages.
