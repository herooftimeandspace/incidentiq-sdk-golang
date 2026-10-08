// Package incidentiq provides a Go client for the Incident IQ API.
//
// Supported calls live on the Golden surface, generated from the bundled
// Incident IQ OpenAPI contract and reached as client.<Namespace>.<Method>.
// Quasi-supported routes derived from live site interaction HARs live on the
// Silver surface, under client.Silver.<Namespace>.<Method>, with app routes
// nested under client.Silver.Apps.<AppNamespace>.<Method>.
//
// Every generated method shares one signature:
//
//	func (s *GoldenAssetsService) GetAssetById(ctx context.Context, opts RequestOptions, out any) error
//
// RequestOptions carries path, query, header, and body values; out receives the
// decoded JSON response, or may be nil to discard it. Client.Request is the
// escape hatch for a route the bundled contracts do not cover.
//
// The contract artifacts under data/ are shared with the
// herooftimeandspace/incident-py-q SDK, so both clients speak to the same
// routes under the same environment variables, authentication headers, URL
// normalization rules, and retry policy. The Go API shape, documentation, and
// tooling are this repository's own.
package incidentiq

//go:generate go run scripts/generate_wrappers.go
//go:generate go run scripts/generate_sdk_reference.go
