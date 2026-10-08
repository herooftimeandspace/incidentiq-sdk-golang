---
name: pr-review-safety
description: Use before pushing code, opening or updating a PR, resolving review feedback, or deciding whether an incidentiq-sdk-golang PR can merge.
---

# PR Review Safety

Review is a gate beside tests, coverage, docs builds, promotion checks, and
release labels. This skill is reviewer-agnostic: it reads the same whether the
reviewer is a person or any review bot.

Read `AGENTS.md` first for repository rules, then the most specific docs,
workflow files, or scripts for the files being changed, plus the issue body,
labels, comments, linked PRs, and current review-thread state.

## Before Pushing

Run `scripts/verify.sh` and fix what it reports. Do not push on a failing gate.

An **independent pre-push review is required** before pushing code to GitHub,
opening a PR, or updating a PR branch. It is part of verification, not a
summary step: a reviewer that did not write the change catches defects the
author's own context hides.

This applies to human- and agent-authored branches. It does **not** apply to
the Actions-authored promotion chain (`promote/dev-to-staging`,
`promote/staging-to-main`), which has no agent in the loop and could not comply;
that content is gated by required checks, the live label and ancestry checks in
`release-prep` and `promotion-head`, and the branch ruleset.

**The minimum compliant form is a fresh context with no implementation history
reading the PR-ready diff.** A separate agent or a different model is better,
but a new session that did not write the change already satisfies this. Every
harness can do that, which is what makes this gate different from a hosted
review product: it needs no vendor, no installation, and no CI check. Do not
encode it as a GitHub Actions check-run; see **Merge Readiness** for why an
AI-review signoff bridge is forbidden.

Because the floor is that low, declining the gate is reserved for cases where
even a fresh context is impossible. If that happens, say so in the **PR body**
— not only in your own transcript — and treat the PR as needing the human
approver to accept the absence explicitly before it can merge.

Running it:

- The reviewer must be independent of the implementation work.
- Ask it to find reasons the branch should **not** be pushed yet.
- Give it the PR-ready diff, not a summary: the issue, branch base, changed
  files, nearby unchanged code, docs, workflow rules, and local verification
  evidence.
- Fix actionable findings before pushing. If a finding raises a product,
  security, release, or operations decision that cannot be safely inferred,
  stop and record the blocker rather than guessing.

Target the review at what changed:

- **Workflow, promotion, coverage, release, branch, or review-policy changes:**
  test stale-state, timestamp-ordering, race-condition, missing-evidence,
  ambiguous-signal, and bypass scenarios.
- **SDK runtime changes:** check Golden and Silver request behavior; header
  shape (the default `Client: ApiClient`, a caller-supplied `Client`,
  `OmitClientHeader`, `SiteId`, auth mode, content type); the Silver
  retry-without-`Client` fallback; tenant-root path handling; retry and
  idempotency behavior, including retry storms and response-status preservation;
  bounded response reads; generated-wrapper drift; and parity with the bundled
  source artifacts.
- **Docs-only changes:** check that commands, branch names, release labels,
  coverage gates, and workflow behavior match the actual implementation.

## Merge Readiness

Every PR except a promotion PR is merge-ready when **all** of these hold:

- Required status checks pass on the **current head commit**.
- A **human has approved** the PR.
- No open review thread is current and unaddressed.
- No reviewer has requested changes.
- The independent pre-push review ran on the current head, and its actionable
  findings are fixed or recorded as accepted risk **in the PR body or a PR
  comment**, where the human approver can see them. A record that exists only
  in an agent's transcript does not count, and leaves a later reviewer unable
  to confirm the gate ran at all. If the review was declined under the rule in
  **Before Pushing**, the human approver must accept that absence explicitly.

**Promotion PRs are not human-gated.** A promotion PR is one whose head is
`promote/dev-to-staging` or `promote/staging-to-main`; judge it by that head
ref, not by its base, so a hotfix opened directly against `staging` or `main` is
an ordinary PR and still needs an approval.

The chain is designed to run unattended: promotion PRs are opened by Actions
with `GITHUB_TOKEN`, and `.github/workflows/promotion.yml` exists so they can
satisfy required checks at all. Requiring a human approval there would be a
policy the automation cannot satisfy, and such policies get routed around.

A promotion PR is merge-ready when its required checks pass on the current head,
it carries exactly one `semver:*` label where the branch rules require one, and
no reviewer has raised an open thread on it.

What a promotion PR carries is *mostly* work already reviewed into `dev`, but
not entirely, so do not treat promotion as proof of review:

- `dev` permits the repository owner to bypass the PR gate, so a direct push to
  `dev` reaches `main` through two exempt PRs.
- `prepare-release-promotion` rewrites the promotion branch with a version bump
  that never passed through `dev` at all — the same commit `AGENTS.md` requires
  you to sync back.

Those paths are covered by required checks, the live label and ancestry checks
in `release-prep` and `promotion-head`, and the branch ruleset — not by review.
If a promotion PR ever carries substantive unreviewed work, treat that as the
defect and review it.

Do not merge on missing evidence, ambiguous review status, blocked labels, or a
stale approval that predates the latest substantive push. Silence is not
approval.

No hosted review product is required, and none should become required. A bot
signoff gate breaks the moment that product is unavailable, not installed, or
replaced — and a gate that cannot be satisfied is worse than no gate, because it
gets routed around. If a review bot does comment here, treat it as one more
reviewer under the thread rules below, never as the thing that unblocks a merge.
Do not add a GitHub Actions check-run bridge for AI review signoff.

## Thread-Aware Review Inspection

Flat PR comments hide whether a thread is resolved, outdated, or tied to the
current head. Use a query that preserves thread state:

```bash
gh api graphql \
  -f owner=herooftimeandspace \
  -f name=incidentiq-sdk-golang \
  -F number=<pr-number> \
  -f query='
query($owner:String!,$name:String!,$number:Int!){
  repository(owner:$owner,name:$name){
    pullRequest(number:$number){
      number
      headRefOid
      updatedAt
      reviewDecision
      commits(last:1){nodes{commit{oid committedDate}}}
      reviewThreads(first:100){
        nodes{
          id
          isResolved
          isOutdated
          path
          line
          comments(first:20){
            nodes{author{login} body createdAt url}
          }
        }
      }
      reviews(last:50){
        nodes{author{login} state submittedAt url}
      }
    }
  }
}'
```

For each thread:

- Unresolved and not outdated: **blocking**.
- Unresolved and outdated: **blocking** until the branch contains a fix or the
  thread is demonstrably obsolete for the current diff.
- Resolved: confirm it corresponds to a commit or a documented decision that
  actually addresses the feedback.

Resolve or reply to a thread only after the branch update makes the feedback
fixed or obsolete.

If feedback lands on a promotion PR, fix the source branch and let the change
flow through `dev -> staging -> main`. Do not patch only the promotion branch
unless repository docs explicitly require it.

## Freshness

An approval or clean signal counts only when it is newer than the work it
claims to clear. Record and compare:

- the current `headRefOid` and its commit time
- the latest review request or review-triggering comment
- the latest pushed remediation commit
- the latest unresolved review thread activity

A removed reaction, a stale comment, an old approval, or a missing comment is
not current signoff. When a signal predates the current head, request a fresh
review instead of treating it as clearance.

The same test applies to the pre-push review. A push that changes behavior,
interfaces, or policy re-triggers it. A push that cannot change what a reviewer
would conclude — a rebase with no content change, a commit-message edit, a
formatting-only fixup — does not; say which case it is when you push on top of
a completed review.

## Review Output

- Lead with actionable findings by severity.
- Include file and line references when available.
- Distinguish blocking defects from optional improvements.
- State which sources, threads, timestamps, and checks were inspected.
- If nothing is wrong, say so plainly and list residual risk and checks not run.
