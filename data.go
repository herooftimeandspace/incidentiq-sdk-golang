package incidentiq

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed data/*.json data/postman/*.json data/stoplight/*.json data/stoplight/controllers/*.json testdata/contract/*.json
var embeddedData embed.FS

// GoldenOperation describes one Stoplight-derived SDK operation from the
// Python SDK's golden inventory.
type GoldenOperation struct {
	Method      string `json:"method"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	OperationID string `json:"operation_id"`
	Path        string `json:"path"`
}

// SilverOperation describes one HAR-derived route from the Python SDK's Silver
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

// GoldenInventory returns the bundled Python SDK Golden operation inventory.
func GoldenInventory() ([]GoldenOperation, error) {
	var inventory []GoldenOperation
	if err := readEmbeddedJSON("testdata/contract/golden_sdk_inventory.json", &inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}

// SilverInventory returns the bundled Python SDK Silver operation inventory.
func SilverInventory() ([]SilverOperation, error) {
	var inventory []SilverOperation
	if err := readEmbeddedJSON("testdata/contract/silver_sdk_inventory.json", &inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}

// StoplightControllerNames returns the bundled controller contract names.
func StoplightControllerNames() ([]string, error) {
	entries, err := fs.ReadDir(embeddedData, "data/stoplight/controllers")
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		names = append(names, strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())))
	}
	return names, nil
}

func readEmbeddedJSON(path string, out any) error {
	payload, err := embeddedData.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, out)
}
