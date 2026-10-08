#!/usr/bin/env bash
set -euo pipefail

# Fail when the committed generated output is not what the generators produce.
#
# generated_wrappers.go and docs/sdk-reference are both committed and are built
# from the same bundled contract artifacts by the two //go:generate directives
# in doc.go. Regenerating must be a no-op; if it is not, one generator was run
# without the other, or an artifact sync was committed without regenerating.
#
# This lives in one file because every producer of the `unit` check has to apply
# the same gate: .github/workflows/quality.yml and both promotion-owned
# unit_script blocks in .github/workflows/promotion.yml.

paths=(generated_wrappers.go docs/sdk-reference)

go generate ./...

if [ -n "$(git status --porcelain -- "${paths[@]}")" ]; then
  echo "go generate ./... changed the generated output; commit it." >&2
  git status --porcelain -- "${paths[@]}" >&2
  # Stage intent-to-add so a newly generated page shows as a diff rather than
  # as a bare untracked line.
  git add --intent-to-add -- "${paths[@]}" >/dev/null 2>&1 || true
  git diff -- "${paths[@]}" >&2
  exit 1
fi

echo "Generated output is up to date."
