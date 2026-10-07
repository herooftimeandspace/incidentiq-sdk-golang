#!/usr/bin/env bash
set -euo pipefail

source_repo="${1:-../incident-py-q}"

if [[ ! -d "${source_repo}/.git" ]]; then
  echo "source repo not found: ${source_repo}" >&2
  exit 1
fi

# The source SDK adds and removes reference pages as the contract changes, so
# the local copy is replaced wholesale rather than merged. The Stoplight and
# Postman trees are retired; drop them so a re-sync over an old checkout is
# idempotent.
rm -rf docs/sdk-reference data/stoplight data/postman
mkdir -p docs/sdk-reference data/openapi data/legacy testdata/contract

cp "${source_repo}/"*.md .
cp "${source_repo}/LICENSE" .
cp "${source_repo}/docs/"*.md docs/
cp "${source_repo}/docs/sdk-reference/"*.md docs/sdk-reference/
cp "${source_repo}/src/incident_py_q/data/app_schemas.json" data/
cp "${source_repo}/src/incident_py_q/data/silver_inventory.json" data/
cp "${source_repo}/src/incident_py_q/data/source_manifest.json" data/
cp "${source_repo}/src/incident_py_q/data/openapi/"*.json data/openapi/
cp "${source_repo}/src/incident_py_q/data/legacy/"*.json data/legacy/
cp "${source_repo}/tests/contract/"*_sdk_inventory.json testdata/contract/

echo "synced docs and contract artifacts from ${source_repo}"
