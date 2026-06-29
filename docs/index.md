# Incident IQ SDK

`incidentiq-sdk-golang` is a Go client/SDK for Incident IQ.

It combines:
- low-level request client
- contract and inventory artifacts copied from the source repository
- Golden inventory helpers for documented Stoplight routes
- Silver inventory helpers for HAR-observed routes
- tests to prevent accidental runtime behavior drift

## Design Highlights

- Bearer token auth by default.
- Tenant base URL is explicit and configurable per client instance.
- Bare tenant roots are normalized to `/api/v1.0`.
- Tenant-root paths beginning with `/api/`, `/services/`, `/apps/`, `/img/`, or `/s/` are sent from the tenant origin.
- Runtime never downloads schema documents.
- Future integration tests should use separate `INCIDENTIQ_TEST_*` variables for smoke usage.

## Quick Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	incidentiq "github.com/herooftimeandspace/incidentiq-sdk-golang"
)

func main() {
	client, err := incidentiq.NewClient(incidentiq.Config{
		BaseURL:  "https://your-tenant.incidentiq.com",
		APIToken: "your-token",
	})
	if err != nil {
		log.Fatal(err)
	}

	var payload map[string]any
	err = client.Request(context.Background(), "GET", "/users/{UserId}", incidentiq.RequestOptions{
		PathParams: map[string]any{"UserId": "00000000-0000-0000-0000-000000000000"},
	}, &payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T\n", payload)
}
```
