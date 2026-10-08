#!/usr/bin/env bash
set -euo pipefail

# Sync the contract artifacts this SDK shares with herooftimeandspace/incident-py-q.
#
# Only machine-readable artifacts are copied. Documentation is NOT synced: the
# Markdown in this repository describes the Go API, the Go toolchain, and this
# repository's own CI, and docs/sdk-reference is generated from the artifacts
# below by scripts/generate_sdk_reference.go.

source_repo="${1:-../incident-py-q}"

if [[ ! -d "${source_repo}/.git" ]]; then
  echo "source repo not found: ${source_repo}" >&2
  exit 1
fi

# The Stoplight and Postman trees are retired; drop them so a re-sync over an
# old checkout is idempotent.
rm -rf data/stoplight data/postman

mkdir -p data/openapi data/legacy testdata/contract

cp "${source_repo}/src/incident_py_q/data/app_schemas.json" data/
cp "${source_repo}/src/incident_py_q/data/silver_inventory.json" data/
cp "${source_repo}/src/incident_py_q/data/source_manifest.json" data/
cp "${source_repo}/src/incident_py_q/data/openapi/"*.json data/openapi/
cp "${source_repo}/src/incident_py_q/data/legacy/"*.json data/legacy/
cp "${source_repo}/tests/contract/"*_sdk_inventory.json testdata/contract/

echo "synced contract artifacts from ${source_repo}"
echo "next: go generate ./... && go test ./..."
