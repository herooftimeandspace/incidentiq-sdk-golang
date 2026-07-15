package incidentiq

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
	"unicode"
)

func TestGeneratedWrapperInventoriesMatchSourceInventories(t *testing.T) {
	golden, err := GoldenInventory()
	if err != nil {
		t.Fatalf("GoldenInventory returned error: %v", err)
	}
	wantGolden := make([]string, 0, len(golden))
	for _, operation := range golden {
		wantGolden = append(wantGolden, operation.Namespace+"."+operation.Name)
	}
	sort.Strings(wantGolden)
	gotGolden := GoldenWrapperInventory()
	sort.Strings(gotGolden)
	if !reflect.DeepEqual(gotGolden, wantGolden) {
		t.Fatalf("GoldenWrapperInventory mismatch: got %d entries, want %d", len(gotGolden), len(wantGolden))
	}

	silver, err := SilverInventory()
	if err != nil {
		t.Fatalf("SilverInventory returned error: %v", err)
	}
	wantSilver := make([]string, 0, len(silver))
	for _, operation := range silver {
		wantSilver = append(wantSilver, operation.Namespace+"."+operation.Name)
	}
	sort.Strings(wantSilver)
	gotSilver := SilverWrapperInventory()
	sort.Strings(gotSilver)
	if !reflect.DeepEqual(gotSilver, wantSilver) {
		t.Fatalf("SilverWrapperInventory mismatch: got %d entries, want %d", len(gotSilver), len(wantSilver))
	}
}

func TestGeneratedWrapperMethodCountsMatchSourceInventories(t *testing.T) {
	client, err := NewClient(Config{
		BaseURL:  "https://example.incidentiq.com",
		APIToken: "token",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	golden, err := GoldenInventory()
	if err != nil {
		t.Fatalf("GoldenInventory returned error: %v", err)
	}
	wantGolden := countByExportedNamespace(golden)
	goldenServices := reflect.TypeOf(client.generatedClientServices)
	for i := 0; i < goldenServices.NumField(); i++ {
		field := goldenServices.Field(i)
		got := field.Type.NumMethod()
		if wantGolden[field.Name] != got {
			t.Fatalf("golden namespace %s has %d generated methods, want %d", field.Name, got, wantGolden[field.Name])
		}
		delete(wantGolden, field.Name)
	}
	if len(wantGolden) != 0 {
		t.Fatalf("missing golden generated namespaces: %#v", wantGolden)
	}

	silver, err := SilverInventory()
	if err != nil {
		t.Fatalf("SilverInventory returned error: %v", err)
	}
	wantSilver := countGeneratedSilverByExportedNamespace(t, silver)
	silverServices := reflect.TypeOf(*client.Silver)
	for i := 0; i < silverServices.NumField(); i++ {
		field := silverServices.Field(i)
		if field.Name == "Apps" {
			appServices := field.Type.Elem()
			for j := 0; j < appServices.NumField(); j++ {
				appField := appServices.Field(j)
				got := generatedWrapperMethodCount(appField.Type)
				key := field.Name + "." + appField.Name
				if wantSilver[key] != got {
					t.Fatalf("silver namespace %s has %d generated methods, want %d", key, got, wantSilver[key])
				}
				delete(wantSilver, key)
			}
			continue
		}
		got := generatedWrapperMethodCount(field.Type)
		if wantSilver[field.Name] != got {
			t.Fatalf("silver namespace %s has %d generated methods, want %d", field.Name, got, wantSilver[field.Name])
		}
		delete(wantSilver, field.Name)
	}
	if len(wantSilver) != 0 {
		t.Fatalf("missing silver generated namespaces: %#v", wantSilver)
	}
}

func TestGeneratedSilverAppsWrapperUsesNestedAppsNamespace(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/apps/googleDeviceData/api/googleDeviceData/site-app/status/last"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Client"), "ApiClient"; got != want {
			t.Fatalf("Client = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:    server.URL,
		APIToken:   "token",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if client.Silver == nil || client.Silver.Apps == nil || client.Silver.Apps.GoogleDeviceData == nil {
		t.Fatal("client.Silver.Apps.GoogleDeviceData was nil")
	}
	var payload map[string]any
	if err := client.Silver.Apps.GoogleDeviceData.GetStatusLast(context.Background(), RequestOptions{
		PathParams: map[string]any{"google_device_data_key": "site-app"},
	}, &payload); err != nil {
		t.Fatalf("GetStatusLast returned error: %v", err)
	}
	if payload["ok"] != true {
		t.Fatalf("payload = %#v, want ok true", payload)
	}
}

func TestGeneratedGoldenWrapperUsesDirectClientNamespace(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/api/v1.0/users"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:    server.URL,
		APIToken:   "token",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if client.Users == nil {
		t.Fatal("client.Users was nil")
	}
	var payload map[string]any
	if err := client.Users.GetUsersLegacy(context.Background(), RequestOptions{}, &payload); err != nil {
		t.Fatalf("GetUsersLegacy returned error: %v", err)
	}
	if payload["ok"] != true {
		t.Fatalf("payload = %#v, want ok true", payload)
	}
}

func TestGeneratedSilverWrapperUsesSilverNamespace(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/api/v1.0/tickets/ticket-1/status"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:    server.URL,
		APIToken:   "token",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if client.Silver == nil || client.Silver.Tickets == nil {
		t.Fatal("client.Silver.Tickets was nil")
	}
	var payload map[string]any
	if err := client.Silver.Tickets.GetTicketStatus(context.Background(), RequestOptions{
		PathParams: map[string]any{"ticket_id": "ticket-1"},
	}, &payload); err != nil {
		t.Fatalf("GetTicketStatus returned error: %v", err)
	}
	if payload["ok"] != true {
		t.Fatalf("payload = %#v, want ok true", payload)
	}
}

func TestGeneratedWrappersInvokeAllInventoryOperations(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:    server.URL,
		APIToken:   "token",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	golden, err := GoldenInventory()
	if err != nil {
		t.Fatalf("GoldenInventory returned error: %v", err)
	}
	for _, operation := range golden {
		service := reflect.ValueOf(client.generatedClientServices).FieldByName(testExportedName(operation.Namespace))
		invokeGeneratedWrapper(t, "Golden."+operation.Namespace+"."+operation.Name, service, operation.Name, operation.Path)
	}

	silver, err := SilverInventory()
	if err != nil {
		t.Fatalf("SilverInventory returned error: %v", err)
	}
	reserved := typedSilverMethodKeys(t)
	for _, operation := range silver {
		if reserved[operation.Namespace+"."+operation.Name] {
			continue
		}
		service := silverServiceValue(t, client.Silver, operation.Namespace)
		invokeGeneratedWrapper(t, "Silver."+operation.Namespace+"."+operation.Name, service, operation.Name, operation.Path)
	}
}

// TestGeneratorOmitsFutureReservedTypedMethods simulates a source SDK sync that
// introduces a method already implemented by a handwritten typed helper. It
// runs the real generator in an isolated directory and proves that the method
// remains in the public inventory without producing a duplicate generic method.
func TestGeneratorOmitsFutureReservedTypedMethods(t *testing.T) {
	temporaryRoot := t.TempDir()
	for _, directory := range []string{"scripts", "data", "testdata/contract"} {
		if err := os.MkdirAll(filepath.Join(temporaryRoot, directory), 0o755); err != nil {
			t.Fatalf("create generator fixture directory %s: %v", directory, err)
		}
	}
	for _, path := range []string{
		"scripts/generate_wrappers.go",
		"data/typed_silver_methods.json",
		"testdata/contract/golden_sdk_inventory.json",
		"testdata/contract/silver_sdk_inventory.json",
	} {
		copyGeneratorFixtureFile(t, path, filepath.Join(temporaryRoot, path))
	}

	var silver []SilverOperation
	fixturePath := filepath.Join(temporaryRoot, "testdata/contract/silver_sdk_inventory.json")
	payload, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf("read synthetic Silver inventory: %v", err)
	}
	if err := json.Unmarshal(payload, &silver); err != nil {
		t.Fatalf("parse synthetic Silver inventory: %v", err)
	}
	silver = append(silver, SilverOperation{
		HTTPMethod: "POST",
		Name:       "add_user_room",
		Namespace:  "users",
		Path:       "/api/v1.0/users/{user_id}/rooms/{location_room_id}",
		Provenance: "silver",
		Sources:    []string{"synthetic-source-sync"},
	})
	payload, err = json.MarshalIndent(silver, "", "  ")
	if err != nil {
		t.Fatalf("encode synthetic Silver inventory: %v", err)
	}
	if err := os.WriteFile(fixturePath, append(payload, '\n'), 0o644); err != nil {
		t.Fatalf("write synthetic Silver inventory: %v", err)
	}

	command := exec.Command("go", "run", "scripts/generate_wrappers.go")
	command.Dir = temporaryRoot
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("run generator with synthetic source sync: %v\n%s", err, output)
	}
	generated, err := os.ReadFile(filepath.Join(temporaryRoot, "generated_wrappers.go"))
	if err != nil {
		t.Fatalf("read synthetic generated wrappers: %v", err)
	}
	source := string(generated)
	if strings.Contains(source, "func (s *SilverUsersService) AddUserRoom(ctx context.Context, opts RequestOptions") {
		t.Fatal("generator emitted a generic AddUserRoom method reserved for the typed helper")
	}
	if !strings.Contains(source, `"users.add_user_room"`) {
		t.Fatal("generator removed the reserved operation from SilverInventory")
	}
}

// copyGeneratorFixtureFile copies one repository-owned generator input into an
// isolated test directory so the real generator cannot alter the checkout.
func copyGeneratorFixtureFile(t *testing.T, sourcePath, destinationPath string) {
	t.Helper()
	payload, err := os.ReadFile(sourcePath)
	if err != nil {
		t.Fatalf("read generator fixture %s: %v", sourcePath, err)
	}
	if err := os.WriteFile(destinationPath, payload, 0o644); err != nil {
		t.Fatalf("write generator fixture %s: %v", destinationPath, err)
	}
}

type inventoryOperation interface {
	GetNamespace() string
}

func (op GoldenOperation) GetNamespace() string {
	return op.Namespace
}

func (op SilverOperation) GetNamespace() string {
	return op.Namespace
}

func countByExportedNamespace[T inventoryOperation](operations []T) map[string]int {
	counts := map[string]int{}
	for _, operation := range operations {
		namespace := operation.GetNamespace()
		if parent, child, ok := strings.Cut(namespace, "."); ok {
			counts[testExportedName(parent)+"."+testExportedName(child)]++
			continue
		}
		counts[testExportedName(namespace)]++
	}
	return counts
}

// countGeneratedSilverByExportedNamespace excludes methods whose names are
// reserved for handwritten typed helpers. The generator consumes the same
// registry, preventing a future source inventory sync from emitting duplicate
// methods with unsafe generic signatures.
func countGeneratedSilverByExportedNamespace(t *testing.T, operations []SilverOperation) map[string]int {
	t.Helper()
	reserved := typedSilverMethodKeys(t)
	counts := map[string]int{}
	for _, operation := range operations {
		if reserved[operation.Namespace+"."+operation.Name] {
			continue
		}
		if parent, child, ok := strings.Cut(operation.Namespace, "."); ok {
			counts[testExportedName(parent)+"."+testExportedName(child)]++
			continue
		}
		counts[testExportedName(operation.Namespace)]++
	}
	return counts
}

func typedSilverMethodKeys(t *testing.T) map[string]bool {
	t.Helper()
	var methods []struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	}
	if err := readEmbeddedJSON("data/typed_silver_methods.json", &methods); err != nil {
		t.Fatalf("read typed Silver method registry: %v", err)
	}
	keys := make(map[string]bool, len(methods))
	for _, method := range methods {
		keys[method.Namespace+"."+method.Name] = true
	}
	return keys
}

// generatedWrapperMethodCount counts only the uniform inventory-generated
// wrapper signature. Handwritten typed helpers may share a generated service
// type without being mistaken for generated inventory operations.
func generatedWrapperMethodCount(serviceType reflect.Type) int {
	count := 0
	for index := 0; index < serviceType.NumMethod(); index++ {
		method := serviceType.Method(index)
		if method.Type.NumIn() == 4 && method.Type.NumOut() == 1 {
			count++
		}
	}
	return count
}

var testNonIdentifier = regexp.MustCompile(`[^0-9A-Za-z]+`)
var testPathParamPattern = regexp.MustCompile(`\{([^}]+)\}`)

func silverServiceValue(t *testing.T, silver *SilverClient, namespace string) reflect.Value {
	t.Helper()
	service := reflect.ValueOf(silver).Elem()
	if parent, child, ok := strings.Cut(namespace, "."); ok {
		parentValue := service.FieldByName(testExportedName(parent))
		if !parentValue.IsValid() || parentValue.IsNil() {
			t.Fatalf("missing Silver parent namespace %s", parent)
		}
		service = parentValue.Elem()
		return service.FieldByName(testExportedName(child))
	}
	return service.FieldByName(testExportedName(namespace))
}

func invokeGeneratedWrapper(t *testing.T, label string, service reflect.Value, operationName string, path string) {
	t.Helper()
	if !service.IsValid() || service.IsNil() {
		t.Fatalf("%s service was nil", label)
	}
	method := service.MethodByName(testExportedName(operationName))
	if !method.IsValid() {
		t.Fatalf("%s wrapper method was missing", label)
	}

	var payload map[string]any
	results := method.Call([]reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(RequestOptions{PathParams: pathParamsForTest(path)}),
		reflect.ValueOf(&payload),
	})
	if len(results) != 1 {
		t.Fatalf("%s returned %d values, want 1", label, len(results))
	}
	if err, ok := results[0].Interface().(error); ok && err != nil {
		t.Fatalf("%s returned error: %v", label, err)
	}
	if payload["ok"] != true {
		t.Fatalf("%s payload = %#v, want ok true", label, payload)
	}
}

func pathParamsForTest(path string) map[string]any {
	params := map[string]any{}
	for _, match := range testPathParamPattern.FindAllStringSubmatch(path, -1) {
		params[match[1]] = "value"
	}
	return params
}

func testExportedName(value string) string {
	parts := strings.Fields(testNonIdentifier.ReplaceAllString(value, " "))
	for i, part := range parts {
		runes := []rune(part)
		for j, r := range runes {
			if j == 0 {
				runes[j] = unicode.ToUpper(r)
				continue
			}
			runes[j] = unicode.ToLower(r)
		}
		parts[i] = string(runes)
	}
	name := strings.Join(parts, "")
	if name == "" {
		return "Value"
	}
	if unicode.IsDigit(rune(name[0])) {
		return "N" + name
	}
	return name
}
