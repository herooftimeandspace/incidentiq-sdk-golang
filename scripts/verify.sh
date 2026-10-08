#!/usr/bin/env bash
set -euo pipefail

# The local verification gate. This script is the single owner of the command
# list: AGENTS.md, CONTRIBUTING.md, and the skills under .agents/skills point
# here instead of each carrying their own copy to drift.
#
# It mirrors what the `unit` and `docs-build` checks do in CI. CI keeps its own
# sequencing because it also publishes badges and coverage ratchets.
#
# Usage:
#   scripts/verify.sh           # generated output, vet, tests, coverage, docs site
#   scripts/verify.sh --quick   # skip the docs site build

# This script is the documented entry point, so it must work from anywhere in
# the checkout.
cd "$(dirname "$0")/.."

export GOCACHE="${GOCACHE:-$(pwd)/.gocache}"
export GOMODCACHE="${GOMODCACHE:-$(pwd)/.gomodcache}"

# The permanent floor. Coverage is also ratcheted per branch: when the published
# badge for the base branch is higher than this, keep coverage at or above the
# badge instead of letting it drift down to the floor.
coverage_minimum="${COVERAGE_MINIMUM:-95.0}"

quick=0
for arg in "$@"; do
  case "$arg" in
    --quick) quick=1 ;;
    *) echo "unknown argument: $arg" >&2; exit 2 ;;
  esac
done

step() { printf '\n==> %s\n' "$1"; }

step "Generated output is up to date"
scripts/check_generated.sh

step "go vet"
go vet ./...

step "Tests with native coverage"
go test -covermode=atomic -coverprofile=coverage.out ./...

step "Coverage reports"
go tool cover -func=coverage.out -o coverage-summary.txt
go tool cover -html=coverage.out -o coverage.html
tail -1 coverage-summary.txt

step "Coverage floor (${coverage_minimum}%)"
go run scripts/build_badge_json.go coverage \
  --coverage-file coverage.out \
  --label "coverage local" \
  --minimum "${coverage_minimum}" \
  --output coverage-badge.json

if [ "$quick" -eq 0 ]; then
  step "Docs site"
  go run scripts/build_docs_site.go
fi

printf '\nAll local checks passed.\n'
