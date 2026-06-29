package incidentiq

import (
	"context"
	"net/http"
	"net/http/httptest"
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
	wantSilver := countByExportedNamespace(silver)
	silverServices := reflect.TypeOf(*client.Silver)
	for i := 0; i < silverServices.NumField(); i++ {
		field := silverServices.Field(i)
		got := field.Type.NumMethod()
		if wantSilver[field.Name] != got {
			t.Fatalf("silver namespace %s has %d generated methods, want %d", field.Name, got, wantSilver[field.Name])
		}
		delete(wantSilver, field.Name)
	}
	if len(wantSilver) != 0 {
		t.Fatalf("missing silver generated namespaces: %#v", wantSilver)
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
		counts[testExportedName(operation.GetNamespace())]++
	}
	return counts
}

var testNonIdentifier = regexp.MustCompile(`[^0-9A-Za-z]+`)

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
