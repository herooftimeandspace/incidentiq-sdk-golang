# API Reference

Public Go package documentation is available through the standard Go toolchain:

```bash
go doc ./...
```

Runtime inventory helpers are documented in code and in the generated SDK reference pages:

- Golden Stoplight methods are documented under the Golden namespace pages.
- Silver HAR-derived methods are documented under the generated [`silver` overview](sdk-reference/silver.md).
- The Go API exposes Golden methods directly on `client.<Namespace>.<Method>` and Silver fallback methods under `client.Silver.<Namespace>.<Method>`.

Build and test locally:

```bash
GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...
```

Then review:
- [SDK reference index](sdk-reference/index.md)
- [Silver overview](sdk-reference/silver.md)
- [Go parity notes](go-parity.md)
