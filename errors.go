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

// ResponseTooLargeError reports a response that exceeded the configured
// bounded read limit before the SDK could safely decode or attach it to an API
// error.
type ResponseTooLargeError struct {
	Method string
	Path   string
	Limit  int64
}

func (e *ResponseTooLargeError) Error() string {
	return fmt.Sprintf("incidentiq %s %s response exceeded %d byte limit", e.Method, e.Path, e.Limit)
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
