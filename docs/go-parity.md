# Go SDK Parity Notes

This repo is the Go companion to `herooftimeandspace/incident-py-q`.

The copied Markdown files under this repo are intentionally retained so the Go
SDK can be reviewed against the same product, contract, validation, and release
documentation as the source SDK. When the source repo changes a shared contract
or user-facing behavior, run `scripts/sync_from_source_sdk.sh` and then update
the Go runtime or generated wrappers until the tests prove parity again.

## Shared Runtime Behavior

The Go client currently matches the source SDK for the following runtime rules:

- `INCIDENTIQ_*` environment variable names
- bearer and raw authorization modes
- optional `SiteId` header
- HTTPS-only base URL validation
- tenant root normalization to `/api/v1.0`
- tenant-root handling for `/api/`, `/services/`, `/apps/`, `/img/`, `/s/`, and `/pub/`
- JSON object and array response decoding
- retry behavior for idempotent methods and retryable HTTP statuses
- Golden wrappers directly on `client.<Namespace>.<Method>` as the correct default SDK path
- Silver wrappers under `client.Silver.<Namespace>.<Method>` for quasi-supported API calls derived from live site interaction HARs, with app routes nested under `client.Silver.Apps.<AppNamespace>.<Method>`

## Contract Artifacts

The Go repo embeds these synced source SDK artifacts:

- `data/openapi/openapi-spec.json`
- `data/openapi/metadata.json`
- `data/legacy/aliases.json`
- `data/legacy/contract.json`
- `data/source_manifest.json`
- `data/app_schemas.json`
- `data/silver_inventory.json`
- `testdata/contract/golden_sdk_inventory.json`
- `testdata/contract/silver_sdk_inventory.json`
- `testdata/contract/merged_sdk_inventory.json`

## Generated Wrapper Surface

The Go repo generates wrappers from the bundled Golden and Silver SDK inventory
snapshots. Golden is the golden SDK path, so Golden wrappers are promoted
directly onto `Client` as `client.<Namespace>.<Method>`. Silver is a separate
namespace for quasi-supported API calls derived from live site interaction HARs,
so Silver wrappers stay under `Client.Silver`. App-specific Silver routes are
nested under `Client.Silver.Apps`.

Regenerate wrappers after refreshing inventories:

```bash
go generate ./...
```

## OpenAPI Contract Migration

The Golden contract source moved from the Stoplight controller sync and the
APIHub Postman collection to the published Incident IQ OpenAPI 3.0 document.
`CHANGELOG.md` describes that migration in the source SDK's terms. The Go SDK
mirrors it with these differences:

- **Deprecation signalling.** Go has no runtime deprecation warning, so every
  legacy name is emitted as a forwarding method carrying a `// Deprecated:` doc
  comment. `go doc`, editors, and `staticcheck` surface it; nothing is logged at
  runtime.
- **Alias coverage.** All 114 aliases from `data/legacy/aliases.json` are
  generated: 50 renamed Golden methods and 64 routes that moved to Silver.
  The alias layer covers pre-migration **Golden** names only. Silver method
  names are not aliased; see *Removed Silver methods* below. The
  six legacy namespaces the new contract dropped (`alerts`, `forms`,
  `manufacturers`, `notifications`, `parts`, `purchaseorders`) still exist on
  `Client` and hold only deprecated forwarders into `client.Silver`.
- **Aliases never shadow generated methods.** The generator emits contract
  operations first and skips any alias whose exported name is already taken.
- **Conflict discovery.** `LegacyAliasConflicts()` is the Go equivalent of
  `incident_py_q.legacy_alias_conflicts()`.
- **No response validation.** `data/legacy/contract.json` is embedded for
  parity, but the Go SDK unmarshals into `out any` and performs no
  response-schema validation, so that bundle currently has no functional
  consumer here.

### `Tickets.AssignTicket` changed meaning

This is the one name that could not be aliased, because the new contract gives
it to a different operation:

| | |
|---|---|
| Before | `POST /tickets/{TicketId}/sla` — assign an **SLA** |
| Now | `POST /api/v1.0/tickets/{ticketId}/assign` — assign the **ticket** |
| Previous behavior moved to | `Tickets.AssignTicketSla` |

This compiles unchanged and silently does something different at runtime.
Callers of `client.Tickets.AssignTicket` that meant to assign an SLA must move
to `client.Tickets.AssignTicketSla`.

### Removed Silver methods

The alias layer does not cover Silver. 45 `client.Silver.<Namespace>.<Method>`
methods that existed before the migration are gone, and every one of them is a
compile error for existing callers.

40 are still reachable after a hand edit, because the published contract now
documents the route and it moved onto the Golden surface (sometimes under a
different name). 5 are gone outright: the route is in neither contract. That
removal happened upstream in `incident-py-q`, not here.

| Removed | Now |
|---|---|
| `Silver.Analytics.GetAssetSummaryStats` | `Analytics.GetAssetSummaryStats` |
| `Silver.Analytics.GetRequestorSummaryStats` | `Analytics.GetRequestorSummaryStats` |
| `Silver.Assets.GetAssetBySerial` | `Assets.GetAssetBySerial` |
| `Silver.Assets.GetAssetFiles` | `Assets.GetAssetFiles` |
| `Silver.Assets.GetAssetVerifications` | `Assets.GetAssetVerificationsForAsset` |
| `Silver.Assets.GetStatsLocations` | `Assets.GetAssetStatsByLocation` |
| `Silver.Assets.GetType` | `Assets.GetAssetType` |
| `Silver.Assets.PostCheckoutsTransactionsQueryGet` | `Assets.GetAssetCheckoutTransactions` |
| `Silver.Audits.GetPoliciesSchedulesForAsset` | `Audits.GetAssetAuditPolicySchedulesForAsset` |
| `Silver.Categories.GetOfFilters` | `Categories.ListFilterCategories` |
| `Silver.Categories.GetOfModels` | `Categories.ListModelCategories` |
| `Silver.CustomFields.PostForAsset` | `CustomFields.GetCustomFieldsForAsset` |
| `Silver.CustomFields.PostForTicket` | `CustomFields.GetCustomFieldsForTicket` |
| `Silver.CustomFields.PostForUser` | `CustomFields.GetCustomFieldsForUser` |
| `Silver.Files.GetEntity` | `Files.GetFilesForEntity` |
| `Silver.Filters.GetForEntitytype` | `Filters.ListFiltersForEntityType` |
| `Silver.Filters.GetSet` | `Filters.GetFilterSet` |
| `Silver.Labor.GetRatesUser2` | `Labor.GetUserLaborRates` |
| `Silver.Labor.PostTypes` | `Labor.QueryLaborTypes` |
| `Silver.Models.GetAll` | `Models.ListAllModelsGet` |
| `Silver.Models.GetAppsAeriesSis` | *(route no longer in either contract)* |
| `Silver.Models.GetAppsGoogleDeviceData` | *(route no longer in either contract)* |
| `Silver.Models.GetAppsMicrosoftIntune` | *(route no longer in either contract)* |
| `Silver.Models.GetAppsSubticketsForIT` | *(route no longer in either contract)* |
| `Silver.Models.PostAvailableToSite` | `Models.GetModelsAvailableToSite` |
| `Silver.Models.PostEndpoint` | `Models.SearchModels` |
| `Silver.Sites.GetDeployments` | *(route no longer in either contract)* |
| `Silver.Sites.PostRoles` | `Sites.ListRolesFiltered` |
| `Silver.Subtasks.GetSubtask` | `Subtasks.GetSubtasksForTicket` |
| `Silver.Surveys.GetResponsesTicket` | `Surveys.GetSurveyResponseForTicket` |
| `Silver.Teams.GetEndpoint` | `Teams.ListTeams` |
| `Silver.Tickets.GetTicketActivities` | `Tickets.GetTicketActivities` |
| `Silver.Tickets.GetTicketKbArticles` | `Tickets.GetTicketKbArticles` |
| `Silver.Tickets.GetTicketNextSteps` | `Tickets.GetTicketNextSteps` |
| `Silver.Tickets.GetTicketStatus` | `Tickets.GetTicketStatus` |
| `Silver.Tickets.PostEndpoint` | `Tickets.SearchTickets` |
| `Silver.Tickets.PostTicketTimeline` | `Tickets.ListTicketTimeline` |
| `Silver.Users.GetMyShortcuts` | `Users.GetMyShortcuts` |
| `Silver.Users.GetShortcutsAvailable` | `Users.GetAvailableShortcuts` |
| `Silver.Users.GetSimple` | `Users.GetSimpleUser` |
| `Silver.Users.GetUserOptions` | `Users.GetUserOptions` |
| `Silver.Users.GetUserRelationships` | `Users.GetUserRelationships` |
| `Silver.Users.GetUserRooms` | `Users.GetUserRooms` |
| `Silver.Views.GetView` | `Views.GetViewDefinition` |
| `Silver.Views.GetView2` | `Views.GetViewDefinition` |

To check a method yourself rather than trusting this table, look the route up by
HTTP method and path in `testdata/contract/golden_sdk_inventory.json` and
`testdata/contract/silver_sdk_inventory.json`; Golden paths gained the
`/api/v1.0` prefix in this migration, so compare paths with that prefix stripped.

## Promotion Branch Shape

Both promotion legs run through a dedicated branch rather than promoting a
long-lived branch as the PR head:

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

Because PRs opened with `GITHUB_TOKEN` do not trigger `pull_request` workflow
runs, `.github/workflows/promotion.yml` reports the required checks itself,
against the promotion branch head: `unit` and `integration` for
`dev -> staging`; `unit`, `integration`, `docs-build`, and `release-prep` for
`staging -> main`. The ordinary PR workflows expose the same context names
and fail closed if the matching promotion-owned check-run is missing or failed,
so they can never substitute a weaker result for a promotion-owned one.

The `dev -> staging` leg also reports `promotion-head`. It is **advisory**: it
is not a required context on `staging`, so it cannot block a merge. It exists to
make one case loud — if two `dev` pushes race and the promotion branch is
force-pushed back to the older tip, `promotion-head` fails because the head no
longer contains `origin/dev`. Merge safety itself is covered by the ruleset,
which refuses a head behind `staging`, and by `unit` being computed on that same
head.

`promotion-head` and `release-prep` re-check branch ancestry live, because the
base branch can move without producing a new promotion head commit.

Note that `CONTRIBUTING.md`, `CHANGELOG.md`, and `README.md` are synced verbatim
from the source SDK by `scripts/sync_from_source_sdk.sh`, which copies both
`<source>/*.md` and `<source>/docs/*.md`. Go-specific workflow notes belong here
or in `AGENTS.md`, not in those files, or the next sync silently reverts them.
This file survives only because the source SDK has no `docs/go-parity.md`; the
filename is the protection, so do not rename it to something the source SDK
might also use.
