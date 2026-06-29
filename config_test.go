package incidentiq

import (
	"strings"
	"testing"
)

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

func TestConfigFromEnvUsesTestPrefixAndRejectsBadAppHeaders(t *testing.T) {
	t.Setenv("INCIDENTIQ_TEST_BASE_URL", "https://test.incidentiq.com")
	t.Setenv("INCIDENTIQ_TEST_API_TOKEN", "test-token")
	t.Setenv("INCIDENTIQ_TEST_APP_HEADERS_JSON", `{"X-Test":"yes"}`)

	config, err := ConfigFromEnv(true)
	if err != nil {
		t.Fatalf("ConfigFromEnv(true) returned error: %v", err)
	}
	if got, want := config.BaseURL, "https://test.incidentiq.com"; got != want {
		t.Fatalf("BaseURL = %q, want %q", got, want)
	}
	if got, want := config.AppHeaders["X-Test"], "yes"; got != want {
		t.Fatalf("AppHeaders[X-Test] = %q, want %q", got, want)
	}

	t.Setenv("INCIDENTIQ_TEST_APP_HEADERS_JSON", `["bad"]`)
	if _, err := ConfigFromEnv(true); err == nil {
		t.Fatal("ConfigFromEnv accepted non-object app headers JSON")
	}
}

func TestConfigNormalizationRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   string
	}{
		{name: "missing base url", config: Config{APIToken: "token"}, want: "base_url is required"},
		{name: "missing token", config: Config{BaseURL: "https://example.incidentiq.com"}, want: "api_token is required"},
		{name: "invalid auth", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", AuthMode: "basic"}, want: "auth_mode must be bearer or raw"},
		{name: "negative timeout", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", Timeout: -1}, want: "timeout must be greater than zero"},
		{name: "negative backoff", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", BackoffBase: -1}, want: "backoff_base must be greater than zero"},
		{name: "negative retries", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", MaxRetries: -1}, want: "max_retries must be zero or positive"},
		{name: "negative response limit", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", MaxResponseBodyBytes: -1}, want: "max_response_body_bytes must be zero or positive"},
		{name: "site header injection", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", SiteID: "site\r\nbad"}, want: "site_id must not contain CR or LF characters"},
		{name: "app header key injection", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", AppHeaders: map[string]string{"X-Test\nBad": "yes"}}, want: "app_headers key must not contain CR or LF characters"},
		{name: "app header value injection", config: Config{BaseURL: "https://example.incidentiq.com", APIToken: "token", AppHeaders: map[string]string{"X-Test": "yes\nbad"}}, want: "app_headers value must not contain CR or LF characters"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.config)
			if err == nil || err.Error() != tt.want {
				t.Fatalf("NewClient error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestConfigNormalizationSetsDefaultResponseBodyLimit(t *testing.T) {
	client, err := NewClient(Config{
		BaseURL:  "https://example.incidentiq.com",
		APIToken: "token",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if got, want := client.Config().MaxResponseBodyBytes, defaultMaxResponseBodyBytes; got != want {
		t.Fatalf("MaxResponseBodyBytes = %d, want %d", got, want)
	}
}

func TestNormalizeBaseURLRejectsUnsafeShapes(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "parse error", raw: "https://%zz", want: "invalid URL escape"},
		{name: "missing host", raw: "https:///path", want: "base_url must include a network location"},
		{name: "credentials", raw: "https://user:pass@example.incidentiq.com", want: "base_url must not contain credentials"},
		{name: "query", raw: "https://example.incidentiq.com?x=1", want: "base_url must not include query parameters or fragments"},
		{name: "fragment", raw: "https://example.incidentiq.com#top", want: "base_url must not include query parameters or fragments"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := normalizeBaseURL(tt.raw)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("normalizeBaseURL error = %v, want containing %q", err, tt.want)
			}
		})
	}
}
