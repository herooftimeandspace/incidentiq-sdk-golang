# SDK Usage

The SDK surface is split into Golden and Silver paths:
- Golden: generated from the bundled Incident IQ OpenAPI contract, reached as
  `client.<Namespace>.<Method>`. This is the correct default path.
- Silver: generated from routes observed in live site interaction HARs, reached
  as `client.Silver.<Namespace>.<Method>`. App routes nest one level deeper, as
  `client.Silver.Apps.<AppNamespace>.<Method>`.

Full generated method and route documentation lives under the
[SDK reference](sdk-reference/index.md).

## One Call Shape

Every generated method has the same signature:

```go
func (s *GoldenTicketsService) GetTicket(ctx context.Context, opts RequestOptions, out any) error
```

There is no per-method argument list to learn. The reference page for a method
names the `RequestOptions` field that carries each documented parameter.

```go
type RequestOptions struct {
	PathParams           map[string]any
	Params               map[string]string
	JSON                 any
	Body                 []byte
	ContentType          string
	Headers              map[string]string
	Timeout              time.Duration
	MaxResponseBodyBytes int64
	OmitClientHeader     bool
	OmitSiteIDHeader     bool
}
```

- `PathParams` fills `{placeholders}` in the route. A missing placeholder is a
  `*ValidationError`, not a malformed URL.
- `Params` adds query parameters.
- `JSON` is marshaled as the request body; `Body` with `ContentType` sends bytes
  as-is, which is how multipart uploads are made.
- `Headers` wins over the client-wide headers for that one call.

## Responses

`out` receives the decoded JSON response:

```go
var ticket map[string]any
err := client.Tickets.GetTicket(ctx, opts, &ticket)
```

Pass a pointer to your own struct when you want typed access. The SDK does not
generate response structs; the reference page for each method names the contract
response schema to model yours on.

```go
type ticketResponse struct {
	Item struct {
		TicketId     string `json:"TicketId"`
		TicketNumber int    `json:"TicketNumber"`
	} `json:"Item"`
}

var parsed ticketResponse
err := client.Tickets.GetTicket(ctx, opts, &parsed)
```

Pass `nil` to discard the body entirely.

## Pagination

Incident IQ paging is driven by query parameters on the routes that support it,
and the SDK does not wrap it. Loop on the parameters the reference page lists
for the method and stop when the tenant stops returning items:

```go
for page := 0; ; page++ {
	var batch map[string]any
	err := client.Assets.SearchAssets(ctx, incidentiq.RequestOptions{
		Params: map[string]string{"$p": strconv.Itoa(page), "$s": "100"},
		JSON:   searchBody,
	}, &batch)
	if err != nil {
		return err
	}
	items, _ := batch["Items"].([]any)
	if len(items) == 0 {
		break
	}
	// ...
}
```

## Low-Level Request API

For a route the bundled contracts do not cover:

```go
err := client.Request(ctx, "GET", "/api/v1.0/users/{UserId}", incidentiq.RequestOptions{
	PathParams: map[string]any{"UserId": userID},
}, &user)
```

To resolve a bundled route by its contract namespace and operation name, without
naming the Go method:

```go
err := client.RequestGolden(ctx, "tickets", "get_ticket", opts, &ticket)
err = client.RequestSilver(ctx, "analytics", "get_agent_current_stats", opts, &stats)
```

Both read the embedded inventories and return a `*ValidationError` for an
unknown operation. The inventories are also readable directly:

```go
operations, err := incidentiq.GoldenInventory()
silver, err := incidentiq.SilverInventory()
```

## Silver Examples

```go
err := client.Silver.AppRegistry.GetApp(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"app_key": appKey},
}, &app)

err = client.Silver.Apps.MicrosoftIntune.GetSyncStatusLast(ctx, incidentiq.RequestOptions{}, &status)
```

App routes usually need tenant app headers. Set them once on `Config.AppHeaders`
(or `INCIDENTIQ_APP_HEADERS_JSON`), or per call in `opts.Headers`.

Silver `POST` routes observed from the browser sometimes reject the default
`Client: ApiClient` header. The client retries those once without it; set
`opts.OmitClientHeader` to skip it from the start, and `opts.OmitSiteIDHeader`
to drop the `SiteId` header for a route that rejects it.

## Deprecated Forwarders

Pre-migration method names still compile and forward to their new location. Go
has no runtime deprecation warning, so each one carries a `// Deprecated:`
comment that `go doc`, editors, and `staticcheck` report. To find every call
site, run a linter that reports deprecated identifiers:

```bash
staticcheck -checks SA1019 ./...
```

`incidentiq.LegacyAliasConflicts()` reports the one pre-migration name that now
reaches a different operation. See [migration-openapi.md](migration-openapi.md).
