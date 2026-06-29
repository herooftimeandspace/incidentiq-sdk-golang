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
- Silver is a separate namespace for quasi-supported API calls derived from
  live site interaction HARs.
- Silver methods must stay under `client.Silver.<Namespace>.<Method>` or, for
  app routes, `client.Silver.Apps.<AppNamespace>.<Method>`. They must not be
  presented as the default API surface.
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
- Run native Go coverage for shared request behavior, generated wrappers, and CI
  changes:
  `GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test -covermode=atomic -coverprofile=coverage.out ./...`.
  Use `go tool cover -func=coverage.out -o coverage-summary.txt` for the text
  summary and `go tool cover -html=coverage.out -o coverage.html` for the HTML
  report. Use `go run scripts/build_badge_json.go coverage --coverage-file
  coverage.out --label "coverage local" --minimum 95.0 --output
  coverage-badge.json` to enforce the current CI coverage floor.
- Run `go vet ./...` after touching shared runtime code or generated wrappers.
- Run `go run scripts/build_docs_site.go` after touching docs or docs workflow
  behavior.
- Promotion, badge, and release automation lives in `.github/workflows/` and is
  intentionally aligned with `incident-py-q` while using Go helper scripts under
  `scripts/`.
- Keep the `dev -> staging -> main` promotion chain, `semver:*` labels, badge
  branch payload paths, main-branch docs build, Dependabot config, and
  main-branch GitHub Release behavior aligned with the documented workflow when
  editing CI.
- The automation uses `GITHUB_TOKEN` by default. Do not introduce a personal
  access token unless GitHub repository rules make it unavoidable and the docs
  are updated with the reason.
