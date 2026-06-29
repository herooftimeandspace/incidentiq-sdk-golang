package incidentiq

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
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
		if got, want := r.Header.Get("Client"), "ApiClient"; got != want {
			t.Fatalf("Client header = %q, want %q", got, want)
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

func TestSilverProfilePictureWrapperSendsRawMultipartBody(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("File", "photo.jpg")
	if err != nil {
		t.Fatalf("CreateFormFile returned error: %v", err)
	}
	if _, err := part.Write([]byte("image-bytes")); err != nil {
		t.Fatalf("write multipart part: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/api/v1.0/profiles/user-123/picture"; got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer token"; got != want {
			t.Fatalf("Authorization = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Client"), "ApiClient"; got != want {
			t.Fatalf("Client = %q, want %q", got, want)
		}
		if got := r.Header.Get("SiteId"); got != "" {
			t.Fatalf("SiteId = %q, want omitted", got)
		}
		if got := r.Header.Get("Content-Type"); got == "" || !strings.HasPrefix(got, "multipart/form-data;") {
			t.Fatalf("Content-Type = %q, want multipart form data", got)
		}
		reader, err := r.MultipartReader()
		if err != nil {
			t.Fatalf("MultipartReader returned error: %v", err)
		}
		uploaded, err := reader.NextPart()
		if err != nil {
			t.Fatalf("NextPart returned error: %v", err)
		}
		if uploaded.FormName() != "File" || uploaded.FileName() != "photo.jpg" {
			t.Fatalf("multipart part = %s/%s, want File/photo.jpg", uploaded.FormName(), uploaded.FileName())
		}
		content, err := io.ReadAll(uploaded)
		if err != nil {
			t.Fatalf("read multipart content: %v", err)
		}
		if string(content) != "image-bytes" {
			t.Fatalf("multipart content = %q, want image-bytes", content)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"FileId": "file-123"})
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
	if err := client.Silver.Profiles.PostProfilePicture(context.Background(), RequestOptions{
		PathParams:       map[string]any{"user_id": "user-123"},
		Body:             body.Bytes(),
		ContentType:      writer.FormDataContentType(),
		OmitSiteIDHeader: true,
	}, &payload); err != nil {
		t.Fatalf("Request returned error: %v", err)
	}
	if payload["FileId"] != "file-123" {
		t.Fatalf("payload = %#v, want FileId", payload)
	}
}

func TestRequestRespectsExplicitClientHeader(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Client"), "CustomClient"; got != want {
			t.Fatalf("Client header = %q, want %q", got, want)
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
	if err := client.Request(context.Background(), "GET", "/users", RequestOptions{
		Headers: map[string]string{"Client": "CustomClient"},
	}, nil); err != nil {
		t.Fatalf("Request returned error: %v", err)
	}
}

func TestRequestSilverRetriesWithoutDefaultClientHeader(t *testing.T) {
	attempts := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		switch attempts {
		case 1:
			if got, want := r.Header.Get("Client"), "ApiClient"; got != want {
				t.Fatalf("first Client header = %q, want %q", got, want)
			}
			http.Error(w, "client header not accepted", http.StatusForbidden)
		case 2:
			if got := r.Header.Get("Client"); got != "" {
				t.Fatalf("second Client header = %q, want omitted", got)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"ok":true}`))
		default:
			t.Fatalf("unexpected attempt %d", attempts)
		}
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
	var payload map[string]any
	if err := client.RequestSilver(context.Background(), "tickets", "get_ticket_status", RequestOptions{
		PathParams: map[string]any{"ticket_id": "ticket-1"},
	}, &payload); err != nil {
		t.Fatalf("RequestSilver returned error: %v", err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
	if payload["ok"] != true {
		t.Fatalf("payload = %#v, want ok true", payload)
	}
}

func TestRequestRejectsJSONAndRawBodyTogether(t *testing.T) {
	client, err := NewClient(Config{
		BaseURL:  "https://example.incidentiq.com",
		APIToken: "token",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	err = client.Request(context.Background(), "POST", "/users", RequestOptions{
		JSON: map[string]any{"ok": true},
		Body: []byte("raw"),
	}, nil)
	if err == nil || !strings.Contains(err.Error(), "both JSON and Body") {
		t.Fatalf("Request error = %v, want JSON/raw-body validation", err)
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
