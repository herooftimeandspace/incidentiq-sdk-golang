package incidentiq

import "testing"

func TestAuthorizationValue(t *testing.T) {
	tests := []struct {
		name string
		mode AuthMode
		in   string
		want string
	}{
		{name: "bearer adds prefix", mode: AuthModeBearer, in: "abc", want: "Bearer abc"},
		{name: "bearer keeps prefix", mode: AuthModeBearer, in: "Bearer abc", want: "Bearer abc"},
		{name: "raw keeps token", mode: AuthModeRaw, in: "abc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authorizationValue(tt.in, tt.mode); got != tt.want {
				t.Fatalf("authorizationValue() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConfigFromEnvUsesPythonVariableNames(t *testing.T) {
	t.Setenv("INCIDENTIQ_BASE_URL", "https://example.incidentiq.com")
	t.Setenv("INCIDENTIQ_API_TOKEN", "token")
	t.Setenv("INCIDENTIQ_SITE_ID", "site")
	t.Setenv("INCIDENTIQ_AUTH_MODE", "raw")
	t.Setenv("INCIDENTIQ_APP_HEADERS_JSON", `{"X-App":"yes"}`)

	config, err := ConfigFromEnv(false)
	if err != nil {
		t.Fatalf("ConfigFromEnv returned error: %v", err)
	}
	if config.BaseURL != "https://example.incidentiq.com" {
		t.Fatalf("BaseURL = %q", config.BaseURL)
	}
	if config.AuthMode != AuthModeRaw {
		t.Fatalf("AuthMode = %q, want raw", config.AuthMode)
	}
	if got := config.AppHeaders["X-App"]; got != "yes" {
		t.Fatalf("AppHeaders[X-App] = %q, want yes", got)
	}
}
