# Project Docs

Core repository documents:

- `README.md`: user-facing overview and quickstart.
- `CONTRIBUTING.md`: local setup, validation gates, and contract sync workflow.
- `SECURITY.md`: private reporting and secret handling expectations.
- `IMPLEMENTATION_PLAN.md`: committed execution plan for this repository.
- `AGENTS.md`: repository-specific rules for automated contributors.
- `docs/go-parity.md`: where this SDK deliberately differs from the source SDK.

These files are maintained in the repository root and versioned with code
changes. None of them are synced from another repository: documentation here
describes the Go API, the Go toolchain, and this repository's own CI.

## Generated Documentation

| Output | Generator | Committed |
| --- | --- | --- |
| `docs/sdk-reference/` | `scripts/generate_sdk_reference.go` (`go generate ./...`) | yes |
| `site/` | `scripts/build_docs_site.go` | no (gitignored; built in CI and published to Pages) |

`build_docs_site.go` copies every Markdown file in the repository into `site/`
and writes a static `index.html` listing them, so GitHub Pages publishes the
same documentation tree contributors read in the checkout.

## Promotion automation

Promotion PRs (`dev` → `staging` → `main`) and the release-prep version bump are
authored with the built-in `github.token`. There is **no stored credential** — no
personal access token and no GitHub App. An earlier `PROMOTION_PR_TOKEN` PAT
expired silently and broke promotion with `HTTP 401: Bad credentials`; nothing
can expire this way now.

### How the required checks are satisfied

`staging` and `main` require the promotion checks, and GitHub deliberately does
not fire `pull_request` workflows for pull requests opened with `github.token`,
to prevent recursive runs. A branch pushed with `github.token` produces no
`push` run either.

So `.github/workflows/promotion.yml` reports the required contexts itself, as
API check-runs against the promotion branch head: `unit` and `integration` for
`dev -> staging`; `unit`, `integration`, `docs-build`, and `release-prep` for
`staging -> main`.

The ordinary PR workflows expose the same context names. They are **not** gated
and do **not** wait for the promotion-owned check-run. If the `pull_request`
runs stay parked at `action_required`, the promotion-owned results carry the
contexts. If someone approves them, they run the same validation and report it
honestly. Either path merges, and a real failure still blocks.

`promotion-head` is also reported on the `dev -> staging` leg. It is advisory
rather than required, so it diagnoses an incomplete promotion without blocking a
merge the ruleset already guards.

### Live integration tests

`integration.yml` runs live tests only on pushes to `staging` and `main`, on
manual dispatch, and on the weekly schedule. Elsewhere it records an
`integration` check without calling the tenant. Integration badges are published
only from runs that actually executed live tests.

### Consequences to keep in mind

- Promotion runs through a dedicated branch on both legs (`promote/dev-to-staging`
  and `promote/staging-to-main`), each built from the base tip. Promoting a
  long-lived branch as the PR head opens it `BEHIND` under
  `strict_required_status_checks_policy`. See [go-parity.md](go-parity.md).
- Do not add broad push-triggered unit runs for feature, bugfix, chore, or sync
  branches. Pull requests validate those. Push-triggered runs are reserved for
  `dev`, `staging`, and `main`, which publish badges and drive promotion.
- `promotion.yml` is `workflow_run`-triggered and `release-prep.yml` is
  `pull_request_target`-triggered, so GitHub always runs the copy of those files
  on the **default branch**. Changes to them take effect only once merged to
  `main`.
