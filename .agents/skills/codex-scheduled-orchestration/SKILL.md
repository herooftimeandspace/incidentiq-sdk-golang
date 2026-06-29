---
name: codex-scheduled-orchestration
description: Use when scheduled or heartbeat-style Codex automation selects GitHub issues, polls PRs for Codex feedback, prepares isolated branches, or coordinates promotion work for incidentiq-sdk-golang.
---

# Codex Scheduled Orchestration

This skill defines repo-local policy for external scheduled Codex automation and
manual heartbeat ticks. It is not a local runner implementation. Do not add a
daemon, dispatcher, worker registry, local status database, or package script to
execute these rules unless the user explicitly asks for that runtime.

## Source Order

1. Read `README.md` for SDK goals and parity requirements.
2. Read `AGENTS.md` for repo-specific Golden/Silver, verification, coverage, and
   promotion rules.
3. Read `CONTRIBUTING.md` for branch promotion, release labels, and quality
   gates.
4. Read `.agents/skills/codex-review-safety/SKILL.md` before pushing, opening,
   updating, merging, or resolving review feedback on a PR.
5. Read the issue body, issue labels, issue comments, linked PRs, and current
   review-thread state before selecting work.
6. Read the most specific SDK docs, workflow files, scripts, tests, and
   generated-artifact rules for the selected files.

## Queue Eligibility

- Use GitHub issues as the durable work queue.
- Prefer existing issues over duplicates. Create a new issue only for real,
  actionable work with clear acceptance criteria and references.
- A scheduled tick may inspect all open issues and PRs, but should select only
  one safe unit of work unless the user explicitly asks for batching.
- Skip issues or PRs with `blocked`, `agent-blocked`, `human-review`,
  `human-only`, `security-sensitive`, or similarly blocking labels.
- Treat missing acceptance criteria, ambiguous release policy, credentials,
  live-write uncertainty, broad source-SDK parity decisions, and merge conflicts
  requiring judgment as blockers.
- If selected work overlaps with an open PR, merged commit, or already
  implemented behavior, update the existing issue or PR instead of
  reimplementing.

## Branch And Workspace Rules

- Branch implementation work from the latest `origin/dev` unless a checked-in
  repo document or human instruction says otherwise.
- Use branch names shaped like `issue-<number>-<short-slug>` for issue work.
- Keep one branch per issue and avoid mixing unrelated fixes.
- Do not edit a dirty user checkout for scheduled work. Use an isolated worktree
  when concurrent work or a dirty checkout would make ownership ambiguous.
- Record the selected issue, base branch, base commit, worktree path, and changed
  files in the final report or issue/PR body.
- Before opening a PR, run the internal Codex review gate from
  `.agents/skills/codex-review-safety/SKILL.md`.
- Open implementation PRs against `dev`.
- Promotion PRs should flow `dev -> staging -> main` through the repo workflows.
- If branch protection says a promotion head is behind its base, sync the base
  history back into the source branch with a dedicated sync branch and PR.
- Do not bypass branch protection, required checks, review requirements, release
  labels, or coverage ratchets.

## Heartbeat Tick Behavior

A heartbeat tick is a bounded inspection-and-action pass. Run dry inspection
before write actions.

1. Inspect open PRs:
   - `gh pr list --state open --json number,title,headRefName,baseRefName,url,statusCheckRollup,reviewDecision,updatedAt`
   - Identify PRs with failing, pending, missing, or stale checks.
   - Identify PRs with Codex review comments, unresolved review threads, or stale
     clean signals.
2. Inspect recent merged PRs when the promotion path may have advanced:
   - Confirm `dev`, `staging`, and `main` are aligned as expected.
   - Confirm release labels and releases were created for `main` promotions.
   - Confirm docs, quality, integration, and release workflows are green.
3. Inspect open issues:
   - Select one issue that is safe, labeled or described clearly enough, and not
     already covered by a PR.
   - Leave factual comments only when they add durable state, such as a blocker,
     linked PR, or implemented-in-main evidence.
4. Inspect Codex Review state on relevant PRs:
   - Use thread-aware GraphQL queries from `codex-review-safety`.
   - Compare `headRefOid`, PR update timestamps, review-thread timestamps, review
     request timestamps, remediation pushes, and Codex thumbs-up reactions.
   - Do not mark PRs clear on stale reactions or old clean comments.
5. Act on one safe unit:
   - Fix actionable feedback on a branch from `dev`.
   - Open or update a PR after internal review and local checks.
   - Merge only when required checks, current Codex review state, release labels,
     and branch protection are satisfied.

## Codex Feedback Polling

When the heartbeat is specifically polling for Codex feedback:

- Query every open PR and any recently merged PR in the active promotion chain.
- For each PR, collect:
  - `number`, `baseRefName`, `headRefName`, `headRefOid`, `updatedAt`
  - status checks and workflow URLs
  - review decision
  - unresolved review threads
  - unresolved outdated threads that may still need explicit resolution
  - comments and reviews from `chatgpt-codex-connector[bot]`
  - `THUMBS_UP` reaction presence and whether it is current enough to count
- Resolve GitHub review threads only after code or docs in the source branch make
  the feedback fixed or obsolete.
- If feedback appears on `dev -> staging` or `staging -> main`, fix the source
  branch first, then let the fix promote forward.
- If a clean Codex signal predates the latest source-branch update, request a
  fresh review or wait for the automated review instead of merging.
- Report stale, ambiguous, or missing review signals as blockers.

## Worker Completion Contract

Every scheduled worker or manual heartbeat pass should report:

- issue number or PR number
- base branch and working branch
- changed files
- checks run and whether they passed
- Codex review sources inspected, including thread and reaction freshness
- open blockers or residual risk
- promotion state, if the work merged beyond `dev`

Workers should not resolve review threads, close issues, or merge PRs unless the
branch contains the fix, checks are current, and the repo's promotion policy says
the state is complete.

## Verification

- Run the narrowest local checks relevant to the files changed.
- For SDK code or workflow changes, prefer:
  - `GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go test -covermode=atomic -coverprofile=coverage.out ./...`
  - `go tool cover -func=coverage.out -o coverage-summary.txt`
  - `go run scripts/build_badge_json.go coverage --coverage-file coverage.out --label "coverage local" --minimum 95.0 --output coverage-badge.json`
  - `GOCACHE="$(pwd)/.gocache" GOMODCACHE="$(pwd)/.gomodcache" go vet ./...`
  - `go run scripts/build_docs_site.go` after docs or docs workflow changes
- Do not claim checks passed unless they were actually run in the current branch.
