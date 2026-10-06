package incidentiq

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// TestSilverUserRoomMutationMethods verifies the observed routes, headers, and
// replacement payload without making a live tenant write.
func TestSilverUserRoomMutationMethods(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantBody any
		invoke   func(*SilverUsersService, any) error
	}{
		{
			name: "add", method: http.MethodPost, path: "/api/v1.0/users/user-one/rooms/room-one",
			invoke: func(users *SilverUsersService, out any) error {
				return users.AddUserRoom(context.Background(), "user-one", "room-one", RequestOptions{}, out)
			},
		},
		{
			name: "replace", method: http.MethodPost, path: "/api/v1.0/users/user-one/rooms",
			wantBody: []any{"room-one", "room-two"},
			invoke: func(users *SilverUsersService, out any) error {
				return users.SetUserRooms(context.Background(), "user-one", []string{"room-one", "room-two"}, RequestOptions{}, out)
			},
		},
		{
			name: "remove", method: http.MethodDelete, path: "/api/v1.0/users/user-one/rooms/room-one",
			invoke: func(users *SilverUsersService, out any) error {
				return users.RemoveUserRoom(context.Background(), "user-one", "room-one", RequestOptions{}, out)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != test.method || r.URL.Path != test.path {
					t.Errorf("request = %s %s, want %s %s", r.Method, r.URL.Path, test.method, test.path)
				}
				for name, want := range map[string]string{"Authorization": "Bearer token", "Client": "ApiClient", "SiteId": "site-id"} {
					if got := r.Header.Get(name); got != want {
						t.Errorf("%s = %q, want %q", name, got, want)
					}
				}
				payload, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("read request body: %v", err)
				}
				if test.wantBody == nil && len(payload) != 0 {
					t.Errorf("body = %q, want empty", payload)
				}
				if test.wantBody != nil {
					var got any
					if err := json.Unmarshal(payload, &got); err != nil {
						t.Fatalf("decode request body: %v", err)
					}
					if !reflect.DeepEqual(got, test.wantBody) {
						t.Errorf("body = %#v, want %#v", got, test.wantBody)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"Item":{"IsSuccessful":true}}`))
			}))
			defer server.Close()
			client, err := NewClient(Config{BaseURL: server.URL, APIToken: "token", SiteID: "site-id", HTTPClient: server.Client()})
			if err != nil {
				t.Fatalf("NewClient returned error: %v", err)
			}
			var response map[string]any
			if err := test.invoke(client.Silver.Users, &response); err != nil {
				t.Fatalf("mutation returned error: %v", err)
			}
			if response["Item"] == nil {
				t.Fatalf("response = %#v, want decoded raw response", response)
			}
		})
	}
}

// TestSilverUserRoomMutationValidation proves that ambiguous identifiers and
// caller-controlled mutation bodies are rejected before transport.
func TestSilverUserRoomMutationValidation(t *testing.T) {
	client, err := NewClient(Config{BaseURL: "https://example.incidentiq.com", APIToken: "token"})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	tests := []struct {
		name   string
		invoke func() error
	}{
		{"add missing user", func() error {
			return client.Silver.Users.AddUserRoom(context.Background(), "", "room", RequestOptions{}, nil)
		}},
		{"add missing room", func() error {
			return client.Silver.Users.AddUserRoom(context.Background(), "user", "", RequestOptions{}, nil)
		}},
		{"add body", func() error {
			return client.Silver.Users.AddUserRoom(context.Background(), "user", "room", RequestOptions{JSON: map[string]any{}}, nil)
		}},
		{"set missing user", func() error {
			return client.Silver.Users.SetUserRooms(context.Background(), "", []string{}, RequestOptions{}, nil)
		}},
		{"set nil rooms", func() error {
			return client.Silver.Users.SetUserRooms(context.Background(), "user", nil, RequestOptions{}, nil)
		}},
		{"set empty room ID", func() error {
			return client.Silver.Users.SetUserRooms(context.Background(), "user", []string{" "}, RequestOptions{}, nil)
		}},
		{"set caller body", func() error {
			return client.Silver.Users.SetUserRooms(context.Background(), "user", []string{"room"}, RequestOptions{Body: []byte("[]")}, nil)
		}},
		{"remove body", func() error {
			return client.Silver.Users.RemoveUserRoom(context.Background(), "user", "room", RequestOptions{JSON: []string{}}, nil)
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.invoke(); err == nil {
				t.Fatal("mutation returned nil error, want validation error")
			} else if _, ok := err.(*ValidationError); !ok {
				t.Fatalf("error = %T %v, want *ValidationError", err, err)
			}
		})
	}
}

// TestSilverUserRoomMutationsDisableClientFallback proves that a rejected POST
// is not automatically replayed without Client.
func TestSilverUserRoomMutationsDisableClientFallback(t *testing.T) {
	requests := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Header.Get("Client") == "ApiClient" {
			http.Error(w, "client header rejected", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	client, err := NewClient(Config{BaseURL: server.URL, APIToken: "token", HTTPClient: server.Client()})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	err = client.Silver.Users.AddUserRoom(context.Background(), "user", "room", RequestOptions{}, nil)
	if err == nil || requests != 1 {
		t.Fatalf("error = %v, requests = %d; want one rejected POST without fallback", err, requests)
	}
	if err := client.Silver.Users.AddUserRoom(context.Background(), "user", "room", RequestOptions{OmitClientHeader: true}, nil); err != nil {
		t.Fatalf("explicit browser header shape returned error: %v", err)
	}
	if requests != 2 {
		t.Fatalf("requests = %d, want one explicit second call", requests)
	}
}

// TestUserRoomMutationObservationIsBundled keeps the evidence and source digest
// available from the checkout that owns the typed helpers.
func TestUserRoomMutationObservationIsBundled(t *testing.T) {
	var observation struct {
		SourceSHA256 string            `json:"source_sha256"`
		Contract     map[string]string `json:"service_contract"`
	}
	if err := readEmbeddedJSON("testdata/contract/user_room_mutation_observation.json", &observation); err != nil {
		t.Fatalf("read observation: %v", err)
	}
	if len(observation.SourceSHA256) != 64 || len(observation.Contract) != 3 {
		t.Fatalf("observation = %#v, want digest and three mutation contracts", observation)
	}
}

// TestUserRoomMutationMethodsRemainTyped verifies that the shared generator
// reservation registry points to the handwritten helpers and that none regress
// to the generic generated wrapper signature.
func TestUserRoomMutationMethodsRemainTyped(t *testing.T) {
	reserved := typedSilverMethodKeys(t)
	want := map[string]string{
		"users.add_user_room":    "AddUserRoom",
		"users.remove_user_room": "RemoveUserRoom",
		"users.set_user_rooms":   "SetUserRooms",
	}
	serviceType := reflect.TypeOf((*SilverUsersService)(nil))
	for key, methodName := range want {
		if !reserved[key] {
			t.Errorf("typed Silver registry is missing %q", key)
			continue
		}
		method, ok := serviceType.MethodByName(methodName)
		if !ok {
			t.Errorf("typed helper %s is missing", methodName)
			continue
		}
		if method.Type.NumIn() == 4 {
			t.Errorf("%s has generic generated signature", methodName)
		}
		delete(reserved, key)
	}
	if len(reserved) != 0 {
		t.Fatalf("unexpected typed Silver reservations: %#v", reserved)
	}
}

// TestResponseTooLargeErrorMessages covers both diagnostic forms used when a
// bounded room-association response is rejected before or after an HTTP status.
func TestResponseTooLargeErrorMessages(t *testing.T) {
	withStatus := (&ResponseTooLargeError{StatusCode: http.StatusOK, Method: http.MethodGet, Path: "/users/user/rooms", Limit: 10}).Error()
	if !strings.Contains(withStatus, "HTTP 200") || !strings.Contains(withStatus, "10 byte limit") {
		t.Fatalf("status error = %q, want status and limit", withStatus)
	}
	withoutStatus := (&ResponseTooLargeError{Method: http.MethodGet, Path: "/users/user/rooms", Limit: 10}).Error()
	if strings.Contains(withoutStatus, "HTTP") || !strings.Contains(withoutStatus, "10 byte limit") {
		t.Fatalf("transport error = %q, want limit without HTTP status", withoutStatus)
	}
}
