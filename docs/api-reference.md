# API Reference

There are two layers of reference documentation.

## Package API

The hand-written Go API — `Client`, `Config`, `RequestOptions`, the error types,
and the embedded-contract accessors — is documented in the source as Go doc
comments:

```bash
go doc github.com/herooftimeandspace/incidentiq-sdk-golang
go doc github.com/herooftimeandspace/incidentiq-sdk-golang.RequestOptions
```

The same pages render on [pkg.go.dev](https://pkg.go.dev/github.com/herooftimeandspace/incidentiq-sdk-golang)
for published tags.

## Generated Namespace Methods

The generated wrappers — `client.Tickets.GetTicket(...)` and the rest — are
documented under the generated SDK reference pages, which carry each route, its
parameters, and the `RequestOptions` field that supplies them.

- [SDK reference index](sdk-reference/index.md)
- [Silver overview](sdk-reference/silver.md)

Regenerate them with:

```bash
go generate ./...
```

Golden methods are documented on their own namespace pages and are the correct
default SDK path. Silver HAR-derived methods are documented under the
[`silver` overview](sdk-reference/silver.md), with app routes on their nested
`client.Silver.Apps.<AppNamespace>` pages.
