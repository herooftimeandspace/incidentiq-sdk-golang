package incidentiq

import (
	"embed"
	"encoding/json"
)

//go:embed data/*.json data/openapi/*.json data/legacy/*.json testdata/contract/*.json
var embeddedData embed.FS

// GoldenOperation describes one OpenAPI-derived SDK operation from the
// source SDK's golden inventory.
type GoldenOperation struct {
	Method      string `json:"method"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	OperationID string `json:"operation_id"`
	Path        string `json:"path"`
}

// SilverOperation describes one HAR-derived route from the source SDK's Silver
// inventory. These routes are useful live API helpers but are intentionally
// marked as inferred rather than official documented contracts.
type SilverOperation struct {
	HTTPMethod string   `json:"http_method"`
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Path       string   `json:"path"`
	Provenance string   `json:"provenance"`
	Sources    []string `json:"sources"`
}

// GoldenInventory returns the bundled Golden operation inventory.
func GoldenInventory() ([]GoldenOperation, error) {
	var inventory []GoldenOperation
	if err := readEmbeddedJSON("testdata/contract/golden_sdk_inventory.json", &inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}

// SilverInventory returns the bundled Silver operation inventory.
func SilverInventory() ([]SilverOperation, error) {
	var inventory []SilverOperation
	if err := readEmbeddedJSON("testdata/contract/silver_sdk_inventory.json", &inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}

// OpenAPIMetadata describes the provenance of the bundled Golden OpenAPI
// contract, as recorded by the source SDK at sync time.
type OpenAPIMetadata struct {
	APITitle         string         `json:"api_title"`
	APIVersion       string         `json:"api_version"`
	DocumentationURL string         `json:"documentation_url"`
	OpenAPIVersion   string         `json:"openapi_version"`
	OperationCount   int            `json:"operation_count"`
	PathCount        int            `json:"path_count"`
	SchemaCount      int            `json:"schema_count"`
	SpecURL          string         `json:"spec_url"`
	SyncedAt         string         `json:"synced_at"`
	UpstreamProfile  map[string]any `json:"upstream_profile"`
}

// OpenAPIContractMetadata returns the bundled Golden contract provenance.
func OpenAPIContractMetadata() (OpenAPIMetadata, error) {
	var metadata OpenAPIMetadata
	if err := readEmbeddedJSON("data/openapi/metadata.json", &metadata); err != nil {
		return OpenAPIMetadata{}, err
	}
	return metadata, nil
}

// LegacyAlias records one pre-migration operation name and the operation it
// now forwards to. Surface is either "golden" or "silver".
type LegacyAlias struct {
	LegacyName        string `json:"legacy_name"`
	LegacyNamespace   string `json:"legacy_namespace"`
	LegacyOperationID string `json:"legacy_operation_id"`
	Surface           string `json:"surface"`
	TargetName        string `json:"target_name"`
	TargetNamespace   string `json:"target_namespace"`
	TargetOperationID string `json:"target_operation_id"`
}

// LegacyAliasConflict records a pre-migration operation name that the new
// contract reassigned to a different operation. These names are deliberately
// not aliased: the new meaning wins, so existing callers change behavior
// silently and must be migrated by hand.
type LegacyAliasConflict struct {
	LegacyName              string `json:"legacy_name"`
	LegacyNamespace         string `json:"legacy_namespace"`
	LegacyOperationID       string `json:"legacy_operation_id"`
	LegacyRoute             string `json:"legacy_route"`
	NowResolvesToOperation  string `json:"now_resolves_to_operation_id"`
	NowResolvesToRoute      string `json:"now_resolves_to_route"`
	PreviousBehaviorMovedTo string `json:"previous_behavior_moved_to"`
	Surface                 string `json:"surface"`
}

type legacyAliasBundle struct {
	Aliases   []LegacyAlias         `json:"aliases"`
	Conflicts []LegacyAliasConflict `json:"conflicts"`
}

// LegacyAliases returns the bundled deprecated-alias map.
func LegacyAliases() ([]LegacyAlias, error) {
	var bundle legacyAliasBundle
	if err := readEmbeddedJSON("data/legacy/aliases.json", &bundle); err != nil {
		return nil, err
	}
	return bundle.Aliases, nil
}

// LegacyAliasConflicts returns the pre-migration names that could not be
// aliased because the new contract reassigned them to a different operation.
// Callers can use this to audit their own code for silent behavior changes.
func LegacyAliasConflicts() ([]LegacyAliasConflict, error) {
	var bundle legacyAliasBundle
	if err := readEmbeddedJSON("data/legacy/aliases.json", &bundle); err != nil {
		return nil, err
	}
	return bundle.Conflicts, nil
}

func readEmbeddedJSON(path string, out any) error {
	payload, err := embeddedData.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, out)
}
