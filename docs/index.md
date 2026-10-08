# Incident IQ Go SDK

`incidentiq-sdk-golang` is a contract-driven Incident IQ client for Go. It
imports as `incidentiq` and depends on the standard library only.

It provides:
- a single `*incidentiq.Client` with a `context.Context`-aware request path
- the Golden surface, generated from the bundled Incident IQ OpenAPI contract
- the Silver surface, generated from routes observed in live site interaction HARs
- deprecated forwarders for every pre-migration method name
- golden-surface tests that catch accidental API drift

## Design Highlights

- Bearer token auth by default.
- Tenant base URL is explicit per client; bare tenant roots are normalized to `/api/v1.0`.
- Nothing is downloaded at runtime: the contracts are embedded in the module.
- Integration tests read separate `INCIDENTIQ_TEST_*` variables.
- Every generated method shares one signature, so there is no per-method API to learn.

## Quick Example

```go
client, err := incidentiq.NewClientFromEnv()
if err != nil {
	return err
}

var asset map[string]any
if err := client.Assets.GetAssetById(ctx, incidentiq.RequestOptions{
	PathParams: map[string]any{"assetId": assetID},
}, &asset); err != nil {
	return err
}
```

## Where To Go Next

- [Getting started](getting-started.md)
- [SDK usage](sdk-usage.md)
- [Generated SDK reference](sdk-reference/index.md)
- [Go parity notes](go-parity.md)
