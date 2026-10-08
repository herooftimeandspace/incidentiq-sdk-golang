# AGENTS.md

Repository-wide instructions for **any AI coding agent** working in
`incidentiq-sdk-golang`. This file is model- and vendor-neutral: it applies to
Claude, Codex, Gemini, Copilot, Cursor, Aider, Cline, Windsurf, locally hosted
open-weight models, and anything else driving changes in this checkout.

Follow your global or harness instructions first, then apply these narrower
rules here. Where the two conflict on repository specifics, this file wins.

`CLAUDE.md`, `GEMINI.md`, `CONVENTIONS.md`, `.clinerules`, `.windsurfrules`,
`.cursor/rules/agents.mdc`, and `.github/copilot-instructions.md` all resolve
here. Do not maintain separate instructions in them.

## Source Of Truth

Read only what the task needs:

| Document | Read it for |
| --- | --- |
| This file | Repository rules. Always loaded. |
| `CONTRIBUTING.md` | Branch promotion, release labels, contract sync workflow |
| `docs/go-parity.md` | Where this SDK deliberately differs from the source SDK |
| `docs/sdk-usage.md` | Call shape, `RequestOptions`, Silver usage |
| `docs/schema-validation.md` | What the bundled contracts do and do not drive |
| `docs/sdk-reference/` | Per-method routes and parameters. Generated; large |
| `IMPLEMENTATION_PLAN.md` | Committed execution plan |
| `SECURITY.md` | Secret handling |

This table is a routing aid, not an exhaustive index; `ls docs/*.md` lists the
rest.

**Do not read `README.md` end to end.** It is the package front page for
humans: install, auth, and quickstart examples. Its repository-rule content is
owned by this file and `CONTRIBUTING.md`.

Durable project state lives in checked-in Markdown, not in chat history.

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

## Source SDK Boundary

- The Go SDK is kept functionally aligned with `herooftimeandspace/incident-py-q`
  at the contract and wire level only. When bundled contract artifacts change in
  the source SDK, run `scripts/sync_from_source_sdk.sh`, then `go generate ./...`,
  and update Go behavior, tests, and generated wrappers until parity is restored.
- Documentation is NOT synced and must never be copied from the source SDK.
  Every Markdown file here describes the Go API, the Go toolchain, and this
  repository's CI. `scripts/sync_from_source_sdk.sh` copies machine-readable
  artifacts only; do not reintroduce Markdown copying.

## Generated Artifacts

- `generated_wrappers.go` (from `scripts/generate_wrappers.go`) and
  `docs/sdk-reference/` (from `scripts/generate_sdk_reference.go`) are both
  produced by `go generate ./...` and both are committed. Never hand-edit either; change the
  generator and regenerate.
- `scripts/check_generated.sh` enforces this and every producer of the `unit`
  check calls it: `.github/workflows/quality.yml` and both promotion-owned
  `unit_script` blocks in `.github/workflows/promotion.yml`. The gate lives in
  one script rather than three inline copies so it cannot drift.

## Verification

Run `scripts/verify.sh` before pushing. It is the single owner of the local
command list — generated-output check, `go vet`, tests with native coverage, the
coverage floor, and the docs site build — so no document restates those
commands. `scripts/verify.sh --quick` skips the docs site build.

Scope it when a full run is not warranted: tests for code, generator, contract,
or public-surface changes; `go vet` for shared runtime code or generated
wrappers; the docs site build for docs or docs-workflow changes.

Coverage is ratcheted by branch above a permanent `95.0%` floor. If the
published badge for the relevant base branch is higher, keep coverage at or
above that higher value rather than letting it drift toward the floor. The
script checks the static floor only; the per-branch ratchet is enforced in CI,
which fetches the base branch's published badge. A regression that still clears
`95.0%` therefore passes locally and fails in CI — compare against the badge
yourself when coverage drops.

Never claim a check passed unless it actually ran on the current branch.

## Skills

Skills hold the sequence for their domain. If your harness has no skill-loading
mechanism, read the file directly.

| Skill | Read it when |
| --- | --- |
| `.agents/skills/pr-review-safety/SKILL.md` | before pushing, opening or updating a PR, resolving review feedback, or judging merge readiness |
| `.agents/skills/scheduled-orchestration/SKILL.md` | scheduled or heartbeat-style issue/PR sweeps, isolated branch selection, promotion coordination |

This repository has policy skills only, not a local scheduled-runner
implementation.

## CI And Promotion

Promotion, badge, and release automation lives in `.github/workflows/` and uses
Go helper scripts under `scripts/`. **`docs/project-docs.md` owns how promotion
works and why**; read it before editing any workflow. `CONTRIBUTING.md` owns the
contributor-facing flow. The rules below are the invariants that are easy to
break and expensive to debug, so they are restated here deliberately:

- Keep the `dev -> staging -> main` promotion chain, `semver:*` labels, badge
  branch payload paths, main-branch docs build, Dependabot config, and
  main-branch GitHub Release behavior aligned when editing CI.
- Keep the promotion-owned `unit`, `integration`, `docs-build`, and
  `release-prep` check reports in `.github/workflows/promotion.yml` aligned with
  the ordinary required PR checks. They exist so Actions-authored promotion PRs
  can satisfy required checks at all. Do not replace them with a push-triggered
  run on `promote/*`; that run never fires.
- Never make one producer of a required context wait for another. A previous
  design did, and deadlocked: every required context carried both a FAILURE and
  a SUCCESS. Let both producers do real work and agree.
- Do not add broad push-triggered unit runs for feature, bugfix, chore, or sync
  branches. Push-triggered runs are reserved for `dev`, `staging`, and `main`.
- Promote through a dedicated branch on both legs (`promote/dev-to-staging`,
  `promote/staging-to-main`), each built from the base tip. Never promote a
  long-lived branch as the PR head; `strict_required_status_checks_policy` is
  topological and will open it `BEHIND`.
- Keep a live label and ancestry check in `release-prep`, and a live ancestry
  check in `promotion-head`, because those inputs can change without a new
  promotion head commit.
- If branch protection reports a promotion head as behind its base, sync the
  base history back into the source branch with a dedicated sync branch and PR.
  Do not force-push the promotion branch to work around it.
- Fixes land in `dev` and promote forward. Anything committed directly to
  `main` — including a release-prep version bump — must be synchronized back to
  `dev`, or the next promotion diverges.
- The automation uses `GITHUB_TOKEN` by default. Do not introduce a personal
  access token unless GitHub repository rules make it unavoidable and the docs
  are updated with the reason.

## Agent Capability Assumptions

Nothing here requires a specific harness, hosted review product, or editor
integration. Where a rule mentions tooling, it also names the command that
verifies the result, so an agent without that tooling can still comply. If you
cannot run a required check, say so plainly rather than asserting it passed.
