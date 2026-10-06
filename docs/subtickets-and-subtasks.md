# Subticket and Subtask Creation Status

The SDK does not currently expose a verified method for creating an Incident IQ
subticket or subtask. This page records the checked-in evidence, prevents
callers from treating read helpers as write capabilities, and defines what is
needed to add the missing operation safely.

## Current SDK Evidence

| SDK method | HTTP capability | Classification | Creation status |
| --- | --- | --- | --- |
| `client.Tickets.CreateTicket` | `POST /api/v1.0/tickets/new` | Golden | Creates an ordinary ticket. The bundled `UpdateTicketRequest` schema has assignee, team, affected-user, `Subject`, `IssueDescription`, product, issue, and location fields, but no parent-ticket linkage field. |
| `client.Tickets.CopyTicket` | `POST /api/v1.0/tickets/{TicketId}/copy` | Golden | Creates ordinary ticket copies and can request a named workpackage. Its `CopyTicketOptions` contract has no subtask or parent-child linkage fields, and its response is `ListGetResponseOfTicket`, not a subtask model. |
| `client.Silver.Subtasks.GetSubtask` | `GET /api/v1.0/subtasks/{subtask_id}` | Silver | Read-only. It does not establish a POST route or creation payload. |
| `client.Silver.Tasks.GetEndpoint` | `GET /api/v1.0/tasks` | Silver | Read-only task listing observed in HAR traffic. It does not establish a creation route or payload. |
| `client.Silver.Models.GetAppsSubticketsForIT` | `GET /api/v1.0/models/apps/subticketsForIT` | Silver | Read-only application metadata. It is not a subticket creation operation. |

The published `CreateTicket` response is `ItemCreateResponseOfTicket`, whose
`Item` is a ticket record. That contract does not define a parent-ticket field
that a caller could verify after creation. Passing guessed `ParentTicketId` or
`ParentTicketNumber` properties through `RequestOptions.JSON` would only prove
that the Go client can serialize arbitrary JSON; it would not prove Incident IQ
accepts the fields or creates the requested relationship.

Accordingly, callers such as accountsdot must keep related-work creation
blocked. They must not call `CreateTicket` and assume it created a subticket,
change the known Silver GET route to POST, or construct a custom Incident IQ
HTTP request outside this SDK.

## Evidence Required to Add the Capability

The source SDK needs a sanitized capture from a non-production tenant or an
updated vendor contract for each creation workflow. A usable capture must
establish:

- the exact HTTP method and route;
- whether subtickets and subtasks share a route or use separate routes;
- required authorization, `Client`, `SiteId`, product, or application headers;
- parent ticket ID and parent ticket number fields;
- affected-user or owner fields;
- assigned user and assigned team fields;
- subject, body, and initial comment fields;
- the success status code and response envelope;
- the created item ID and any relationship fields used for verification;
- behavior when the same logical request is sent more than once; and
- the follow-up read that proves parent linkage, affected-user ownership,
  assignment, and sanitized body content.

Do not add secrets, authorization values, cookies, raw personal data, or a full
unsanitized HAR to the repository. Reduce the evidence to the relevant request
and response, replace tenant and entity identifiers with stable placeholders,
and preserve field names, HTTP semantics, and response structure exactly.

Once that evidence exists, add the operation to the source SDK inventory first,
sync the bundled artifacts into this repository, generate the Go wrapper, and
add request/response contract tests. A vendor-documented route belongs on the
Golden surface. A route supported only by sanitized live-site traffic belongs
under `client.Silver`; it must not be presented as Golden.

## Safety Contract for a Future Write Wrapper

Any future wrapper and caller workflow must preserve these requirements:

- Set the affected user or owner to the person whose work is represented. Do
  not silently make the integration account the affected user.
- Send only a sanitized subject, body, and comment. Omit credentials, tokens,
  recovery codes, session data, raw audit payloads, and unrelated personal
  information.
- Do not automatically retry a creation POST unless the endpoint provides a
  verified idempotency mechanism. The current Go client retries only idempotent
  HTTP methods.
- If Incident IQ has no idempotency key, derive and persist a caller-side
  operation key, reconcile unknown outcomes before retrying, and record the
  created item ID against that key.
- After creation, use the verified follow-up read to confirm the parent link,
  affected user, assigned user or team, and sanitized content. Treat a mismatch
  as a failed workflow requiring reconciliation.
- Exercise the write path in a non-production tenant before enabling it in
  production.

Until the required contract evidence is available, there is no response shape,
header set, or duplicate-prevention behavior the SDK can truthfully promise for
subticket or subtask creation.
