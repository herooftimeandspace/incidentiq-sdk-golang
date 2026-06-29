package incidentiq

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const defaultGoldenAPIPrefix = "/api/v1.0"

// AuthMode controls how API tokens are written into the Authorization header.
type AuthMode string

const (
	// AuthModeBearer writes tokens as "Bearer <token>" unless the token already
	// includes that prefix.
	AuthModeBearer AuthMode = "bearer"
	// AuthModeRaw writes the token exactly as supplied.
	AuthModeRaw AuthMode = "raw"
)

// Config contains normalized runtime settings for Client.
type Config struct {
	BaseURL           string
	APIToken          string
	SiteID            string
	AuthMode          AuthMode
	AppHeaders        map[string]string
	Timeout           time.Duration
	MaxResponseBytes  int64
	ValidateResponses bool
	MaxRetries        int
	BackoffBase       time.Duration
	HTTPClient        *http.Client
}

// ConfigFromEnv builds Config from the same runtime variables used by
// incident-py-q. When testMode is true it reads INCIDENTIQ_TEST_* variables.
func ConfigFromEnv(testMode bool) (Config, error) {
	prefix := "INCIDENTIQ_"
	if testMode {
		prefix = "INCIDENTIQ_TEST_"
	}
	config := Config{
		BaseURL:           strings.TrimSpace(os.Getenv(prefix + "BASE_URL")),
		APIToken:          strings.TrimSpace(os.Getenv(prefix + "API_TOKEN")),
		SiteID:            strings.TrimSpace(os.Getenv(prefix + "SITE_ID")),
		AuthMode:          AuthMode(strings.ToLower(strings.TrimSpace(os.Getenv(prefix + "AUTH_MODE")))),
		ValidateResponses: true,
	}
	rawAppHeaders := strings.TrimSpace(os.Getenv(prefix + "APP_HEADERS_JSON"))
	if rawAppHeaders != "" {
		appHeaders := map[string]string{}
		if err := json.Unmarshal([]byte(rawAppHeaders), &appHeaders); err != nil {
			return Config{}, &ConfigurationError{Message: prefix + "APP_HEADERS_JSON must contain a JSON object with string values"}
		}
		config.AppHeaders = appHeaders
	}
	return config, nil
}

func (c Config) normalized() (Config, error) {
	if strings.TrimSpace(c.BaseURL) == "" {
		return Config{}, &ConfigurationError{Message: "base_url is required"}
	}
	if strings.TrimSpace(c.APIToken) == "" {
		return Config{}, &ConfigurationError{Message: "api_token is required"}
	}
	baseURL, err := normalizeBaseURL(c.BaseURL)
	if err != nil {
		return Config{}, err
	}
	c.BaseURL = baseURL
	if c.AuthMode == "" {
		c.AuthMode = AuthModeBearer
	}
	if c.AuthMode != AuthModeBearer && c.AuthMode != AuthModeRaw {
		return Config{}, &ConfigurationError{Message: "auth_mode must be bearer or raw"}
	}
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}
	if c.Timeout < 0 {
		return Config{}, &ConfigurationError{Message: "timeout must be greater than zero"}
	}
	if c.MaxResponseBytes == 0 {
		c.MaxResponseBytes = 4 << 20
	}
	if c.MaxResponseBytes < 0 {
		return Config{}, &ConfigurationError{Message: "max_response_bytes must be greater than zero"}
	}
	if c.BackoffBase == 0 {
		c.BackoffBase = 250 * time.Millisecond
	}
	if c.BackoffBase < 0 {
		return Config{}, &ConfigurationError{Message: "backoff_base must be greater than zero"}
	}
	if c.MaxRetries < 0 {
		return Config{}, &ConfigurationError{Message: "max_retries must be zero or positive"}
	}
	if err := validateHeaderValue(c.SiteID, "site_id"); err != nil {
		return Config{}, err
	}
	for key, value := range c.AppHeaders {
		if err := validateHeaderValue(key, "app_headers key"); err != nil {
			return Config{}, err
		}
		if err := validateHeaderValue(value, "app_headers value"); err != nil {
			return Config{}, err
		}
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: c.Timeout}
	}
	return c, nil
}

func normalizeBaseURL(raw string) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", err
	}
	if strings.ToLower(parsed.Scheme) != "https" {
		return "", &ConfigurationError{Message: "base_url must use https"}
	}
	if parsed.Host == "" {
		return "", &ConfigurationError{Message: "base_url must include a network location"}
	}
	if parsed.User != nil {
		return "", &ConfigurationError{Message: "base_url must not contain credentials"}
	}
	if parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", &ConfigurationError{Message: "base_url must not include query parameters or fragments"}
	}
	parsed.Path = strings.TrimRight(parsed.Path, "/")
	if parsed.Path == "" {
		parsed.Path = defaultGoldenAPIPrefix
	}
	return parsed.String(), nil
}

func validateHeaderValue(value string, fieldName string) error {
	if strings.ContainsAny(value, "\r\n") {
		return &ConfigurationError{Message: fieldName + " must not contain CR or LF characters"}
	}
	return nil
}

func authorizationValue(token string, mode AuthMode) string {
	if mode == AuthModeRaw {
		return token
	}
	if strings.HasPrefix(strings.ToLower(token), "bearer ") {
		return token
	}
	return "Bearer " + token
}
