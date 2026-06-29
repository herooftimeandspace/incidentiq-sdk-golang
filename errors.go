package incidentiq

import (
	"fmt"
	"net/http"
)

// ConfigurationError reports invalid client setup such as unsafe base URLs or
// missing credentials.
type ConfigurationError struct {
	Message string
}

func (e *ConfigurationError) Error() string {
	return e.Message
}

// ValidationError reports a caller-side request construction problem.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// APIError reports a non-2xx HTTP response from Incident IQ.
type APIError struct {
	StatusCode int
	Method     string
	Path       string
	Body       []byte
	Headers    http.Header
}

func (e *APIError) Error() string {
	return fmt.Sprintf("incidentiq %s %s failed with HTTP %d", e.Method, e.Path, e.StatusCode)
}

// ResponseSizeError reports a response body that exceeds the configured SDK
// safety limit before it can be decoded or attached to APIError.
type ResponseSizeError struct {
	Method string
	Path   string
	Limit  int64
}

func (e *ResponseSizeError) Error() string {
	return fmt.Sprintf("incidentiq %s %s response body exceeded %d bytes", e.Method, e.Path, e.Limit)
}
