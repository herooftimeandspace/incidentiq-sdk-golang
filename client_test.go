package incidentiq

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClientNormalizesTenantRoot(t *testing.T) {
	client, err := NewClient(Config{
		BaseURL:  "https://example.incidentiq.com",
		APIToken: "token",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if got, want := client.Config().BaseURL, "https://example.incidentiq.com/api/v1.0"; got != want {
		t.Fatalf("BaseURL = %q, want %q", got, want)
	}
}

func TestNewClientRejectsUnsafeBaseURL(t *testing.T) {
	_, err := NewClient(Config{
		BaseURL:  "http://example.incidentiq.com",
		APIToken: "token",
	})
	if err == nil {
		t.Fatal("NewClient accepted an http base URL")
	}
}

func TestRequestSendsIncidentIQHeadersAndJSON(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/api/v1.0/users/abc%20123"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		if got, want := r.URL.Query().Get("$s"), "10"; got != want {
			t.Fatalf("query $s = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer token"; got != want {
			t.Fatalf("Authorization = %q, want %q", got, want)
		}
		if got := r.Header.Get("Client"); got != "" {
			t.Fatalf("Client header = %q, want absent", got)
		}
		if got, want := r.Header.Get("SiteId"), "site-1"; got != want {
			t.Fatalf("SiteId = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:    server.URL,
		APIToken:   "token",
		SiteID:     "site-1",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	var payload map[string]any
	err = client.Request(context.Background(), "GET", "/users/{UserId}", RequestOptions{
		PathParams: map[string]any{"UserId": "abc 123"},
		Params:     map[string]string{"$s": "10"},
	}, &payload)
	if err != nil {
		t.Fatalf("Request returned error: %v", err)
	}
	if payload["ok"] != true {
		t.Fatalf("payload = %#v, want ok true", payload)
	}
}

func TestTenantRootPathBypassesGoldenPrefix(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/services/tickets/-/-/AssignedToMe_Unassigned"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		w.WriteHeader(http.StatusNoContent)
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
	if err := client.Request(context.Background(), "GET", "/services/tickets/-/-/AssignedToMe_Unassigned", RequestOptions{}, nil); err != nil {
		t.Fatalf("Request returned error: %v", err)
	}
}

func TestIdempotentRequestsRetryRetryableStatus(t *testing.T) {
	attempts := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			http.Error(w, "try again", http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		BaseURL:     server.URL,
		APIToken:    "token",
		HTTPClient:  server.Client(),
		MaxRetries:  1,
		BackoffBase: time.Nanosecond,
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	var payload map[string]any
	if err := client.Request(context.Background(), "GET", "/users", RequestOptions{}, &payload); err != nil {
		t.Fatalf("Request returned error: %v", err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}

func TestInventoriesLoadFromPythonSDKArtifacts(t *testing.T) {
	golden, err := GoldenInventory()
	if err != nil {
		t.Fatalf("GoldenInventory returned error: %v", err)
	}
	if len(golden) == 0 {
		t.Fatal("GoldenInventory returned no operations")
	}
	silver, err := SilverInventory()
	if err != nil {
		t.Fatalf("SilverInventory returned error: %v", err)
	}
	if len(silver) == 0 {
		t.Fatal("SilverInventory returned no operations")
	}
	controllers, err := StoplightControllerNames()
	if err != nil {
		t.Fatalf("StoplightControllerNames returned error: %v", err)
	}
	if len(controllers) == 0 {
		t.Fatal("StoplightControllerNames returned no controllers")
	}
}
