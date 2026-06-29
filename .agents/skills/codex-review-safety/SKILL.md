---
name: codex-review-safety
description: Use before pushing code, opening or updating a PR, resolving Codex review feedback, or deciding whether an incidentiq-sdk-golang PR can merge.
---

# Codex Review Safety

Use this skill as part of verification for `incidentiq-sdk-golang`. Treat Codex
review as a gate beside Go tests, coverage, docs builds, promotion checks, and
release labels.

## Source Order

1. Read `README.md` for SDK behavior, public surface, and parity goals.
2. Read `AGENTS.md` for repo-specific Golden/Silver, coverage, promotion, and
   verification rules.
3. Read `CONTRIBUTING.md` for branch promotion, quality gates, coverage
   ratcheting, and release-label rules.
4. Read the issue body, issue labels, issue comments, linked PRs, and current
   review-thread state before changing or merging a PR.
5. Read the most specific SDK docs, generated reference docs, workflow files,
   or scripts for the files being changed.

## Internal Codex Review Gate

- Run an internal Codex review before pushing code to GitHub, opening a PR, or
  updating a PR branch.
- Treat the internal review as part of verification, not as a summary step.
- The internal reviewer must be independent of the implementation work. Use a
  separate agent, a fresh review context, or an equivalent local Codex review
  workflow that did not author the change.
- Ask the reviewer to find reasons the branch should not be pushed yet.
- The reviewer must inspect the PR-ready diff, not only the implementer's
  summary. Include the issue, branch base, changed files, nearby unchanged code,
  docs, workflow rules, and local verification evidence.
- For workflow, promotion, coverage, release, branch, or review-policy changes,
  the reviewer must actively test stale-state, timestamp-ordering,
  race-condition, missing-evidence, ambiguous-signal, and bypass scenarios.
- For SDK runtime changes, the reviewer must check Golden and Silver request
  behavior, default headers, Silver fallback behavior, tenant-root path handling,
  retry/idempotency behavior, bounded response reads, generated-wrapper drift,
  and parity with bundled source artifacts where relevant.
- For docs-only changes, the reviewer must check that commands, branch names,
  release labels, coverage gates, and workflow behavior match the repo's actual
  implementation.
- Fix actionable internal review findings before pushing, opening, or updating a
  PR. If the reviewer identifies a product, security, release, or operations
  decision that cannot be safely inferred, stop and record the blocker.

## GitHub Codex Review Gate

- Inspect thread-aware GitHub review context before declaring a PR clear.
- A PR is not merge-ready while it has open current GitHub Codex Review threads.
- A PR is not merge-ready until Codex Review has an explicit current clean signal
  after the latest observed PR head-state change.
- A current clean signal can be a `chatgpt-codex-connector[bot]` thumbs-up
  reaction on the PR body, provided the reaction was created after all of:
  - the current `headRefOid`
  - the latest review request or review-triggering comment
  - the latest pushed remediation commit
  - the latest unresolved Codex review thread creation or update
- Do not treat a removed reaction, stale comment, old clean signal, missing
  comment, or stale thread as current signoff.
- Do not merge PRs with open current review threads, requested changes, failing
  required checks, missing evidence, blocked labels, ambiguous review status, or
  missing semver labels on promotion PRs.
- Resolve or reply to review conversations only after the branch update makes
  the feedback fixed or obsolete.

## Thread-Aware Review Inspection

Use GraphQL or a helper that preserves review-thread state. Flat PR comments are
not enough because they hide whether a thread is resolved, outdated, or tied to
the current head.

Minimum thread query:

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
      reviewThreads(first:100){
        nodes{
          id
          isResolved
          isOutdated
          path
          line
          updatedAt
          comments(first:20){
            nodes{
              author{login}
              body
              createdAt
              updatedAt
              url
            }
          }
        }
      }
    }
  }
}'
```

For each `chatgpt-codex-connector` thread:

- Treat unresolved, non-outdated threads as blocking.
- Treat unresolved outdated threads as blocking until the branch contains a fix
  or the thread is explicitly obsolete for the current diff.
- Verify that any resolved thread corresponds to a commit or documented decision
  that actually addresses the feedback.
- If the feedback appears on a promotion PR, fix the source branch first and let
  the change flow through `dev -> staging -> main`. Do not patch only the
  promotion branch unless the repo docs explicitly require that.

## Reaction And Timestamp Inspection

Use reactions as a freshness signal only when their timestamps are newer than
the current head-state evidence.

Minimum reaction query:

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
      comments(last:50){
        nodes{
          author{login}
          body
          createdAt
          updatedAt
          url
          reactionGroups{
            content
            users(first:20){nodes{login}}
          }
        }
      }
      reviews(last:50){
        nodes{
          author{login}
          state
          submittedAt
          body
          url
        }
      }
    }
  }
}'
```

When deciding whether Codex Review is current:

- Record the current `headRefOid`.
- Record the latest commit time, latest PR update time, latest Codex thread
  update time, latest review request/review-triggering comment, and latest
  remediation push.
- Confirm any `THUMBS_UP` reaction from `chatgpt-codex-connector[bot]` is newer
  than those relevant timestamps.
- If the clean reaction predates the current head or remediation push, request or
  wait for a fresh review instead of treating it as signoff.
- If a Codex comment has no current thumbs-up, inspect review threads and commit
  state directly; do not assume silence means approval.

## Historical Feedback Classes For This SDK

Internal reviews should actively check these classes when relevant:

- **Header and request-shape drift:** default `Client: ApiClient`, explicit
  caller-provided `Client`, `OmitClientHeader`, `SiteId`, auth mode, content
  type, and Silver retry-without-client behavior.
- **Retry and body-size safety:** idempotent retry policy, oversized response
  behavior, response status preservation, and retry storms.
- **Golden/Silver namespace shape:** Golden remains `client.<Namespace>.<Method>`;
  Silver remains under `client.Silver`, including `client.Silver.Apps`.
- **Generated artifact drift:** generated wrappers, inventories, source manifests,
  copied Markdown, docs site output, and bundled source SDK artifacts stay aligned.
- **Coverage and branch gates:** coverage ratchets do not allow drops, PR checks
  attach to promotion PRs, semver labels exist, and branch protection is not
  bypassed.
- **Promotion-path safety:** fixes land in `dev`, promote to `staging`, then
  promote to `main`; direct-to-main changes must be synchronized back to `dev`.
- **Docs and command accuracy:** documented commands use Go tooling and match
  actual workflows, labels, branch names, and coverage behavior.

## Review Output

- Lead with actionable findings by severity.
- Include file and line references when available.
- Clearly distinguish blocking defects from optional improvements.
- State which sources, threads, reactions, timestamps, and checks were inspected.
- If no issues are found, say that clearly and list residual risk or checks not
  run.
