---
name: scheduled-orchestration
description: Use when scheduled or heartbeat-style agent automation selects GitHub issues, polls PRs for review feedback, prepares isolated branches, or coordinates promotion work for incidentiq-sdk-golang.
---

# Scheduled Orchestration

Repo-local policy for external scheduled automation and manual heartbeat ticks.
It is not a runner implementation. Do not add a daemon, dispatcher, worker
registry, local status database, or package script to execute these rules unless
the user explicitly asks for that runtime.

Read `AGENTS.md` first for repository rules, then
`.agents/skills/pr-review-safety/SKILL.md` before any push, PR, or merge
decision, then the issue body, labels, comments, linked PRs, and current
review-thread state for the work you select.

## Queue Eligibility

- Use GitHub issues as the durable work queue.
- Prefer existing issues over duplicates. Create a new issue only for real,
  actionable work with clear acceptance criteria and references.
- A tick may inspect all open issues and PRs, but should select only one safe
  unit of work unless the user explicitly asks for batching.
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
- Open implementation PRs against `dev`. Promotion flows `dev -> staging -> main`
  through the repo workflows; `AGENTS.md` owns the promotion constraints.
- Do not bypass branch protection, required checks, review requirements, release
  labels, or coverage ratchets.

## Heartbeat Tick Behavior

A heartbeat tick is a bounded inspection-and-action pass. Run dry inspection
before write actions.

1. Inspect open PRs:
   - `gh pr list --state open --json number,title,headRefName,baseRefName,url,statusCheckRollup,reviewDecision,updatedAt`
   - Identify PRs with failing, pending, missing, or stale checks.
   - Identify PRs with unresolved review threads or stale approvals, using the
     thread-aware queries and freshness rules in `pr-review-safety`.
2. Inspect recent merged PRs when the promotion path may have advanced:
   - Confirm `dev`, `staging`, and `main` are aligned as expected.
   - Confirm release labels and releases were created for `main` promotions.
   - Confirm docs, quality, integration, and release workflows are green.
3. Inspect open issues:
   - Select one issue that is safe, clear enough, and not already covered by a PR.
   - Leave factual comments only when they add durable state, such as a blocker,
     a linked PR, or implemented-in-main evidence.
4. Act on one safe unit:
   - Fix actionable feedback on a branch from `dev`.
   - Run `scripts/verify.sh`, then open or update a PR.
   - Merge only when `pr-review-safety` says the PR is merge-ready.

## Reporting

Every scheduled worker or manual heartbeat pass should report:

- issue number or PR number
- base branch and working branch
- changed files
- checks run and whether they passed
- review sources inspected, including thread and approval freshness
- open blockers or residual risk
- promotion state, if the work merged beyond `dev`

Do not resolve review threads, close issues, or merge PRs unless the branch
contains the fix, checks are current, and the repository's promotion policy says
the state is complete.

## Verification

Run `scripts/verify.sh`, or the narrowest part of it relevant to the files
changed. `AGENTS.md` owns the command list and the coverage ratchet rule. Do not
claim checks passed unless they were actually run on the current branch.
