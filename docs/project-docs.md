# Project Docs

Core repository documents:

- `README.md`: user-facing overview and quickstart.
- `CONTRIBUTING.md`: local setup, validation gates, and contract sync workflow.
- `SECURITY.md`: private reporting and secret handling expectations.
- `IMPLEMENTATION_PLAN.md`: committed execution plan for this repository.
- `AGENTS.md`: repository rules for any AI coding agent. `CLAUDE.md`,
  `GEMINI.md`, `CONVENTIONS.md`, `.clinerules`, and `.windsurfrules` are
  symlinks to it; `.cursor/rules/agents.mdc` and
  `.github/copilot-instructions.md` are pointer files, because Cursor requires
  frontmatter and GitHub's API serves a symlink blob as its path text.
- `.agents/skills/`: policy skills for PR review safety and scheduled
  orchestration, routed from `AGENTS.md`.
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

This section is the single owner of how promotion works. `AGENTS.md` states the
rules an agent must not break and links here for the reasoning.

Promotion PRs (`dev` → `staging` → `main`) and the release-prep version bump are
authored with the built-in `github.token`. There is **no stored credential** — no
personal access token and no GitHub App. An earlier `PROMOTION_PR_TOKEN` PAT
expired silently and broke promotion with `HTTP 401: Bad credentials`; nothing
can expire this way now.

### Promotion branch shape

Both legs run through a dedicated branch rather than promoting a long-lived
branch as the PR head:

| Leg | PR head | Built from |
|---|---|---|
| `dev -> staging` | `promote/dev-to-staging` | `staging` tip, then merges the `dev` tip |
| `staging -> main` | `promote/staging-to-main` | `main` tip, then merges the `staging` tip |

`staging` and `main` both enforce `strict_required_status_checks_policy`. That
policy is **topological**: it asks whether the head branch contains the base
tip, not whether the two trees agree. A base branch accumulates a merge commit
every time a promotion PR lands, and that commit never flows back to the source
branch, so promoting `dev` directly opened every `dev -> staging` PR `BEHIND`
even when `git diff dev...staging` was empty. Building the head from the base
tip removes the problem at its source.

### How the required checks are satisfied

GitHub deliberately does not fire `pull_request` workflows for pull requests
opened with `github.token`, to prevent recursive runs, and a branch pushed with
`github.token` produces no `push` run either. So
`.github/workflows/promotion.yml` reports the required contexts itself, as API
check-runs against the promotion branch head: `unit` and `integration` for
`dev -> staging`; `unit`, `integration`, `docs-build`, and `release-prep` for
`staging -> main`.

The ordinary PR workflows expose the same context names. They are **not** gated
and do **not** wait for the promotion-owned check-run. If the `pull_request`
runs stay parked at `action_required`, the promotion-owned results carry the
contexts. If someone approves them, they run the same validation and report it
honestly. Either path merges, and a real failure still blocks.

An earlier design had those workflows block until the promotion-owned check-run
appeared. That deadlocked the first time the parked runs were approved: the
waiters never observed the promotion-owned check-run — it was present on the
head commit but absent from that commit's check-runs listing — so every
required context ended up carrying both a FAILURE and a SUCCESS, and GitHub
blocked on the failure. Raising the timeout only made it fail slower.

The `dev -> staging` leg also reports `promotion-head`. It is **advisory**: it
is not a required context on `staging`, so it cannot block a merge. It exists to
make one case loud — if two `dev` pushes race and the promotion branch is
force-pushed back to the older tip, `promotion-head` fails because the head no
longer contains `origin/dev`. Merge safety itself is covered by the ruleset,
which refuses a head behind `staging`, and by `unit` being computed on that same
head.

`promotion-head` and `release-prep` re-check branch ancestry live, because the
base branch can move without producing a new promotion head commit.

### Live integration tests

`integration.yml` runs live tests only on pushes to `staging` and `main`, on
manual dispatch, and on the weekly schedule. Elsewhere it records an
`integration` check without calling the tenant. Integration badges are published
only from runs that actually executed live tests.

### Consequences to keep in mind

- Do not add broad push-triggered unit runs for feature, bugfix, chore, or sync
  branches. Pull requests validate those. Push-triggered runs are reserved for
  `dev`, `staging`, and `main`, which publish badges and drive promotion.
- `promotion.yml` is `workflow_run`-triggered and `release-prep.yml` is
  `pull_request_target`-triggered, so GitHub always runs the copy of those files
  on the **default branch**. Changes to them take effect only once merged to
  `main`.
- The promotion-owned `unit` check regenerates on a merge tree. If `dev` and
  `staging` ever regenerate the committed generated output independently, a
  textually clean merge can still disagree with the generators and fail `unit`
  on the promotion PR. Fix it on `dev` and let the promotion branch be rebuilt
  from the base tip, which each cycle already does.
