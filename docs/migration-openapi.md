# Migrating from the pre-OpenAPI SDK

The Golden contract moved from the retired Stoplight controller specs to the published
[Incident IQ OpenAPI contract](https://scopousiiq.github.io/iiq-docusaurus-docs/docs/api/).
That grew the SDK from 173 to 863 documented operations, renamed most methods, and left a
set of routes undocumented.

## What happened to your code

Go code keeps compiling on upgrade except for one name that changed meaning
(see [Requires manual review](#requires-manual-review)) and the Silver methods
listed in [Removed Silver methods](#removed-silver-methods). Every previous
Golden method name still exists: 50 were renamed and 64 moved to the Silver
surface, and both kinds are generated as forwarding methods so you can migrate
incrementally.

Go has no runtime deprecation warning, so each forwarder carries a
`// Deprecated:` comment instead. `go doc`, editors, and `staticcheck` report
it; nothing is logged at runtime. To find every call site to update:

```bash
staticcheck -checks SA1019 ./...
```

Forwarders are excluded from `incidentiq.GoldenWrapperInventory()` and
`incidentiq.SilverWrapperInventory()`, which describe the current contract only.

## Requires manual review

One legacy name is now used by a **different** operation, so it could not be aliased without
shadowing a real method. Calling it still works but now reaches a different route. Check these
call sites by hand; `incidentiq.LegacyAliasConflicts()` returns the same list at runtime.
This one compiles unchanged and silently does something different.

| Legacy call | Used to call | Now calls | Previous behavior moved to |
| --- | --- | --- | --- |
| `client.Tickets.AssignTicket` | `POST /tickets/{TicketId}/sla` | `POST /api/v1.0/tickets/{ticketId}/assign` | `client.Tickets.AssignTicketSla` |

## Routes migrated to the Silver surface

The published contract is a locked allowlist profile and no longer documents these 64
routes. Rather than drop them, they moved to `client.Silver.*`, so behavior is
unchanged apart from the namespace. The pruned copy of the previous contract at
`data/legacy/contract.json` is embedded for parity with the source SDK; this SDK
performs no response validation against it.

| Legacy call | New call |
| --- | --- |
| `client.Alerts.QueueNotification` | `client.Silver.Alerts.QueueNotification` |
| `client.Analytics.GetReport` | `client.Silver.Analytics.GetReport` |
| `client.Analytics.GetReportElements` | `client.Silver.Analytics.GetReportElements` |
| `client.Analytics.GetReportQueries` | `client.Silver.Analytics.GetReportQueries` |
| `client.Analytics.GetReports` | `client.Silver.Analytics.GetReports` |
| `client.Assets.AddManufacturerToSite2` | `client.Silver.Assets.AddManufacturerToSite2` |
| `client.Assets.CreateAssetStatusType` | `client.Silver.Assets.CreateAssetStatusType` |
| `client.Assets.DeleteAssetFundingType` | `client.Silver.Assets.DeleteAssetFundingType` |
| `client.Assets.DeleteAssetStatusType` | `client.Silver.Assets.DeleteAssetStatusType` |
| `client.Assets.GetAssetFavorites2` | `client.Silver.Assets.GetAssetFavorites2` |
| `client.Assets.GetAssetFundingType` | `client.Silver.Assets.GetAssetFundingType` |
| `client.Assets.GetAssetFundingTypes2` | `client.Silver.Assets.GetAssetFundingTypes2` |
| `client.Assets.GetAssetsByAssetStatusType` | `client.Silver.Assets.GetAssetsByAssetStatusType` |
| `client.Assets.GetAssetsByAssetTag` | `client.Silver.Assets.GetAssetsByAssetTag` |
| `client.Assets.GetSpareAssetsByAssetTag` | `client.Silver.Assets.GetSpareAssetsByAssetTag` |
| `client.Assets.GetUserAssets2` | `client.Silver.Assets.GetUserAssets2` |
| `client.Assets.SearchAssetsByAssetTag` | `client.Silver.Assets.SearchAssetsByAssetTag` |
| `client.Assets.UpdateAssetFundingType` | `client.Silver.Assets.UpdateAssetFundingType` |
| `client.Assets.UpdateAssetStatusType` | `client.Silver.Assets.UpdateAssetStatusType` |
| `client.CustomFields.DeleteCustomFields` | `client.Silver.CustomFields.DeleteCustomFields` |
| `client.CustomFields.GetCustomFields` | `client.Silver.CustomFields.GetCustomFields` |
| `client.Forms.SubmitForm` | `client.Silver.Forms.SubmitForm` |
| `client.Locations.DeleteLocation` | `client.Silver.Locations.DeleteLocation` |
| `client.Locations.GetAllLocationRooms` | `client.Silver.Locations.GetAllLocationRooms` |
| `client.Locations.GetLocationRooms` | `client.Silver.Locations.GetLocationRooms` |
| `client.Locations.GetLocationType` | `client.Silver.Locations.GetLocationType` |
| `client.Locations.GetLocationTypes` | `client.Silver.Locations.GetLocationTypes` |
| `client.Locations.UpdateLocation` | `client.Silver.Locations.UpdateLocation` |
| `client.Manufacturers.AddManufacturerToSite3` | `client.Silver.Manufacturers.AddManufacturerToSite3` |
| `client.Manufacturers.AddManufacturerToSite4` | `client.Silver.Manufacturers.AddManufacturerToSite4` |
| `client.Manufacturers.DeleteManufacturer2` | `client.Silver.Manufacturers.DeleteManufacturer2` |
| `client.Manufacturers.GetGlobalManufacturers3` | `client.Silver.Manufacturers.GetGlobalManufacturers3` |
| `client.Manufacturers.GetGlobalManufacturers4` | `client.Silver.Manufacturers.GetGlobalManufacturers4` |
| `client.Manufacturers.GetManufacturer2` | `client.Silver.Manufacturers.GetManufacturer2` |
| `client.Manufacturers.RemoveManufacturerFromSite2` | `client.Silver.Manufacturers.RemoveManufacturerFromSite2` |
| `client.Manufacturers.UpdateManufacturer2` | `client.Silver.Manufacturers.UpdateManufacturer2` |
| `client.Metrics.DeleteMetricType` | `client.Silver.Metrics.DeleteMetricType` |
| `client.Metrics.GetMetric` | `client.Silver.Metrics.GetMetric` |
| `client.Metrics.GetMetricsForSla` | `client.Silver.Metrics.GetMetricsForSla` |
| `client.Metrics.UpdateMetric` | `client.Silver.Metrics.UpdateMetric` |
| `client.Metrics.UpdateMetricType` | `client.Silver.Metrics.UpdateMetricType` |
| `client.Notifications.GetNotifications` | `client.Silver.Notifications.GetNotifications` |
| `client.Notifications.GetTicketEmails` | `client.Silver.Notifications.GetTicketEmails` |
| `client.Notifications.GetUnarchivedNotifications` | `client.Silver.Notifications.GetUnarchivedNotifications` |
| `client.Notifications.GetUnreadNotifications` | `client.Silver.Notifications.GetUnreadNotifications` |
| `client.Notifications.MarkAllNotificationsArchived` | `client.Silver.Notifications.MarkAllNotificationsArchived` |
| `client.Notifications.MarkAllNotificationsRead` | `client.Silver.Notifications.MarkAllNotificationsRead` |
| `client.Notifications.MarkNotificationArchived` | `client.Silver.Notifications.MarkNotificationArchived` |
| `client.Notifications.MarkNotificationRead` | `client.Silver.Notifications.MarkNotificationRead` |
| `client.Parts.DeletePart` | `client.Silver.Parts.DeletePart` |
| `client.Parts.DeletePartSupplier` | `client.Silver.Parts.DeletePartSupplier` |
| `client.Parts.GetPart` | `client.Silver.Parts.GetPart` |
| `client.Parts.GetPartSupplier` | `client.Silver.Parts.GetPartSupplier` |
| `client.Parts.GetPartSuppliers` | `client.Silver.Parts.GetPartSuppliers` |
| `client.Parts.GetParts` | `client.Silver.Parts.GetParts` |
| `client.Parts.UpdatePart` | `client.Silver.Parts.UpdatePart` |
| `client.Parts.UpdatePartSupplier` | `client.Silver.Parts.UpdatePartSupplier` |
| `client.Purchaseorders.DeletePurchaseOrder` | `client.Silver.Purchaseorders.DeletePurchaseOrder` |
| `client.Purchaseorders.GetPurchaseOrder` | `client.Silver.Purchaseorders.GetPurchaseOrder` |
| `client.Purchaseorders.GetPurchaseOrders` | `client.Silver.Purchaseorders.GetPurchaseOrders` |
| `client.Purchaseorders.UpdatePurchaseOrder` | `client.Silver.Purchaseorders.UpdatePurchaseOrder` |
| `client.Slas.GetSla` | `client.Silver.Slas.GetSla` |
| `client.Slas.UpdateSla` | `client.Silver.Slas.UpdateSla` |
| `client.Users.UpdateUserView` | `client.Silver.Users.UpdateUserView` |

## Renamed Golden methods

These 50 routes are still documented, but the new contract renamed their operations.

| Legacy call | New call |
| --- | --- |
| `client.Analytics.GetAssetCountsByVerificationLocation` | `client.Analytics.GetAssetVerificationCountsByLocation` |
| `client.Analytics.GetAssetCountsByVerificationType` | `client.Analytics.GetAssetVerificationCountsByType` |
| `client.Assets.AddUserFavoriteAsset` | `client.Assets.AddAssetFavorite` |
| `client.Assets.GetAsset` | `client.Assets.GetAssetById` |
| `client.Assets.GetAssetFavorites` | `client.Assets.GetUserFavoriteAssets` |
| `client.Assets.GetAssetFundingTypes` | `client.Assets.ListAssetFundingTypes` |
| `client.Assets.GetAssetStatusTypes` | `client.Assets.ListAssetStatusTypes` |
| `client.Assets.GetAssetStatusTypes2` | `client.Assets.ListAssetStatusTypesPost` |
| `client.Assets.GetAssets` | `client.Assets.SearchAssets` |
| `client.Assets.GetAssetsByLocationRoom` | `client.Assets.ListAssetsByRoom` |
| `client.Assets.GetAssetsBySerial` | `client.Assets.GetAssetBySerial` |
| `client.Assets.GetAssetsByStorageUnitNumber` | `client.Assets.ListAssetsByStorageUnit` |
| `client.Assets.GetAssetsCount` | `client.Assets.CountAssets` |
| `client.Assets.GetGlobalManufacturers2` | `client.Assets.SearchGlobalManufacturers` |
| `client.Assets.GetManufacturer` | `client.Assets.GetManufacturerById` |
| `client.Assets.GetUserAssets` | `client.Assets.GetAssetsForUserCurrent` |
| `client.Assets.RemoveUserFavoriteAsset` | `client.Assets.RemoveAssetFavorite` |
| `client.CustomFields.GetCustomField` | `client.CustomFields.GetCustomFieldById` |
| `client.CustomFields.GetCustomFieldType` | `client.CustomFields.GetCustomFieldTypeById` |
| `client.CustomFields.GetCustomFieldTypes` | `client.CustomFields.ListCustomFieldTypes` |
| `client.CustomFields.GetCustomFieldTypes2` | `client.CustomFields.SearchCustomFieldTypes` |
| `client.CustomFields.GetCustomFields2` | `client.CustomFields.SearchCustomFields` |
| `client.Issues.GetIssue` | `client.Issues.GetIssueById` |
| `client.Issues.GetIssueType` | `client.Issues.GetIssueTypeById` |
| `client.Issues.GetIssueTypes` | `client.Issues.SearchIssueTypes` |
| `client.Issues.GetIssueTypesSimple` | `client.Issues.ListIssueTypes` |
| `client.Locations.GetLocation` | `client.Locations.GetLocationByIdV2` |
| `client.Locations.GetLocationRoom` | `client.Locations.GetLocationRoomById` |
| `client.Locations.GetLocations` | `client.Locations.GetAllSiteLocationsV2` |
| `client.Metrics.DeleteMetric` | `client.Metrics.DeleteMetricMetrics` |
| `client.Metrics.GetMetricType` | `client.Metrics.GetMetricById` |
| `client.Metrics.GetMetricTypes` | `client.Metrics.ListMetrics` |
| `client.Metrics.GetMetrics` | `client.Metrics.ListMetricMetrics` |
| `client.Metrics.UpdateMetricsForSla` | `client.Metrics.CreateMetricById` |
| `client.Slas.DeleteSla` | `client.Slas.DeleteSlaById` |
| `client.Slas.GetSlas` | `client.Slas.ListSlas` |
| `client.Tickets.ChangeTicketToRequestorResponded` | `client.Tickets.SetTicketRequestorResponded` |
| `client.Tickets.ChangeTicketToWaitingOnRequestor` | `client.Tickets.SetTicketWaitingOnRequestor` |
| `client.Tickets.GetTicketStatuses` | `client.Tickets.ListTicketStatuses` |
| `client.Tickets.UnAssignTicketFromTeam` | `client.Tickets.UnassignTicketFromTeam` |
| `client.Tickets.UnAssignTicketFromUser` | `client.Tickets.UnassignTicketFromUser` |
| `client.Tickets.UnAssignTicketSla` | `client.Tickets.UnassignTicketSla` |
| `client.Tickets.UnConfirmTicketIssue` | `client.Tickets.UnconfirmTicketIssue` |
| `client.Users.DeleteUserView` | `client.Users.DeleteUserById` |
| `client.Users.GetAgents` | `client.Users.SearchAgents` |
| `client.Users.GetAgentsLegacy` | `client.Users.SearchAgentsLegacyGet` |
| `client.Users.GetUser` | `client.Users.GetUserById` |
| `client.Users.GetUserViews` | `client.Users.ListUserViews` |
| `client.Users.GetUsers` | `client.Users.SearchUsers` |
| `client.Users.GetUsersLegacy` | `client.Users.SearchUsersLegacyGet` |

## Removed Silver methods

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
