# AGENTS.md

This file adds repository-specific guidance for Codex work in `incidentiq-sdk-golang`.
Follow the global instructions first, then apply these narrower rules when working
in this repository.

## Source Of Truth

- Read `README.md` first.
- Use `docs/`, `IMPLEMENTATION_PLAN.md`, `CONTRIBUTING.md`, `SECURITY.md`, and
  generated SDK reference Markdown as durable project documentation.
- The Go SDK is kept functionally aligned with `herooftimeandspace/incident-py-q`.
  When shared Markdown or bundled contract artifacts change in the source SDK,
  use `scripts/sync_from_source_sdk.sh` and then update Go behavior, tests, and
  generated wrappers until parity is restored.

## Golden And Silver Terminology

- Golden refers to the golden SDK path. It is the correct default API surface for
  supported SDK calls.
- Golden methods must be exposed directly as `client.<Namespace>.<Method>`.
- Silver is a separate namespace for quasi-supported API calls derived from live
  site interaction HARs.
- Silver methods must stay under `client.Silver.<Namespace>.<Method>` and must
  not be presented as the default API surface.
- Do not split Golden into a separate `client.Golden` namespace. The direct
  `client.<Namespace>.<Method>` shape is the Golden path.
- Preserve this distinction in docs, generated wrapper text, examples, tests, and
  PR descriptions.

## Verification

- Run `go generate ./...` after changing wrapper generation or bundled inventory
  files.
- Run `GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test ./...`
  after code, generator, contract, or documentation changes that affect the
  public SDK surface.
