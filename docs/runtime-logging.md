# Logging and Runtime Behavior

## Logging

The SDK emits no logs and installs no logger. It never writes to `stdout`,
`stderr`, or `log.Default()`, so it cannot leak an `Authorization` header into
your application's output.

Observe requests from your own code instead, by supplying an `http.Client` with
a wrapping `RoundTripper`:

```go
type loggingTransport struct{ base http.RoundTripper }

func (t loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := t.base.RoundTrip(req)
	// Log req.Method, req.URL.Path, res.StatusCode, time.Since(start).
	// Never log req.Header.Get("Authorization").
	return res, err
}

client, err := incidentiq.NewClient(incidentiq.Config{
	BaseURL:    baseURL,
	APIToken:   token,
	HTTPClient: &http.Client{Transport: loggingTransport{base: http.DefaultTransport}},
})
```

Redaction is then your responsibility: the SDK sets `Authorization`, `Client`,
`SiteId`, and any configured app headers on every request.

## Retry Policy

Conservative retry defaults:
- Retryable statuses: `408`, `429`, `500`, `502`, `503`, `504`
- Retry methods: idempotent methods only (`GET`, `HEAD`, `OPTIONS`, `DELETE`, `PUT`)
- Backoff: exponential, `Config.BackoffBase` doubled per attempt, with no jitter
- Attempt count: `Config.MaxRetries` retries after the first attempt

Non-idempotent writes are never retried. A backoff wait respects the caller's
`ctx`, so a cancelled context ends the retry loop immediately.

Silver has one extra behavior: a Silver route rejected for its `Client` header
is retried once without that header. This is not governed by `MaxRetries`, it
applies to non-idempotent Silver methods too, and the typed user-room helpers
opt out of it so an ambiguous write is never replayed. Set
`RequestOptions.OmitClientHeader` to skip the header from the first attempt.

## Timeouts

- `Config.Timeout` bounds every request made through the client.
- `RequestOptions.Timeout` overrides it for a single call, by cloning the
  underlying `*http.Client`.
- `ctx` deadlines and cancellation apply on top of both.

## Auth Behavior

Default auth mode is bearer:
- token `abc` -> `Authorization: Bearer abc`
- pre-prefixed token `Bearer abc` is preserved

`AuthModeRaw` sends the token exactly as supplied, for tenants that expect a
bare value.

## Headers

Every request sets `Accept: application/json` and `Authorization`. The client
also sets:
- `Client`, always the constant `ApiClient`, unless
  `RequestOptions.OmitClientHeader` is set or `Headers` already carries one
- `SiteId`, when `Config.SiteID` is set and `RequestOptions.OmitSiteIDHeader` is not
- every entry in `Config.AppHeaders`, then every entry in `RequestOptions.Headers`

`RequestOptions.Headers` is applied last, so it wins.
