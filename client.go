package incidentiq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var tenantRootPathPrefixes = []string{"/api/", "/services/", "/apps/", "/img/", "/s/"}

// Client is the shared entry point for Incident IQ HTTP calls.
type Client struct {
	config Config
}

// NewClient validates configuration and returns an Incident IQ client.
func NewClient(config Config) (*Client, error) {
	normalized, err := config.normalized()
	if err != nil {
		return nil, err
	}
	return &Client{config: normalized}, nil
}

// NewClientFromEnv builds a Client from the same environment variables used by
// incident-py-q.
func NewClientFromEnv() (*Client, error) {
	config, err := ConfigFromEnv(false)
	if err != nil {
		return nil, err
	}
	return NewClient(config)
}

// Config returns the normalized client configuration. The API token is included
// so callers should avoid logging this value directly.
func (c *Client) Config() Config {
	return c.config
}

// Request sends one low-level API request and decodes a JSON object or array
// response into out when out is non-nil.
func (c *Client) Request(ctx context.Context, method string, path string, opts RequestOptions, out any) error {
	if strings.TrimSpace(method) == "" {
		return &ValidationError{Message: "method is required"}
	}
	if strings.TrimSpace(path) == "" {
		return &ValidationError{Message: "path is required"}
	}
	renderedPath, err := renderPath(path, opts.PathParams)
	if err != nil {
		return err
	}
	requestURL, err := c.buildURL(renderedPath, opts.Params)
	if err != nil {
		return err
	}
	body, err := encodeJSONBody(opts.JSON)
	if err != nil {
		return err
	}
	method = strings.ToUpper(strings.TrimSpace(method))
	return c.doHTTPRequest(ctx, method, renderedPath, requestURL, body, opts, out)
}

// RequestGolden looks up a bundled Golden operation by namespace and method
// name, then sends it through Request.
func (c *Client) RequestGolden(ctx context.Context, namespace string, name string, opts RequestOptions, out any) error {
	inventory, err := GoldenInventory()
	if err != nil {
		return err
	}
	for _, operation := range inventory {
		if operation.Namespace == namespace && operation.Name == name {
			return c.Request(ctx, operation.Method, operation.Path, opts, out)
		}
	}
	return &ValidationError{Message: fmt.Sprintf("unknown Golden operation %s.%s", namespace, name)}
}

// RequestSilver looks up a bundled Silver operation by namespace and method
// name, then sends it through Request.
func (c *Client) RequestSilver(ctx context.Context, namespace string, name string, opts RequestOptions, out any) error {
	inventory, err := SilverInventory()
	if err != nil {
		return err
	}
	for _, operation := range inventory {
		if operation.Namespace == namespace && operation.Name == name {
			return c.Request(ctx, operation.HTTPMethod, operation.Path, opts, out)
		}
	}
	return &ValidationError{Message: fmt.Sprintf("unknown Silver operation %s.%s", namespace, name)}
}

func (c *Client) buildURL(renderedPath string, params map[string]string) (string, error) {
	if strings.HasPrefix(renderedPath, "http://") || strings.HasPrefix(renderedPath, "https://") {
		return addQuery(renderedPath, params)
	}
	safePath := renderedPath
	if !strings.HasPrefix(safePath, "/") {
		safePath = "/" + safePath
	}
	base, err := url.Parse(c.config.BaseURL)
	if err != nil {
		return "", err
	}
	for _, prefix := range tenantRootPathPrefixes {
		if strings.HasPrefix(safePath, prefix) {
			root := base.Scheme + "://" + base.Host + safePath
			return addQuery(root, params)
		}
	}
	base.Path = strings.TrimRight(base.Path, "/") + "/" + strings.TrimLeft(safePath, "/")
	return addQuery(base.String(), params)
}

func addQuery(rawURL string, params map[string]string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	query := parsed.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

func (c *Client) doHTTPRequest(ctx context.Context, method string, path string, requestURL string, body []byte, opts RequestOptions, out any) error {
	canRetry := methodIsIdempotent(method)
	var lastErr error
	for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
		if attempt > 0 {
			if err := sleepWithContext(ctx, computeBackoff(attempt-1, c.config.BackoffBase)); err != nil {
				return err
			}
		}
		err := c.sendOnce(ctx, method, path, requestURL, body, opts, out)
		if err == nil {
			return nil
		}
		lastErr = err
		if !canRetry || !isRetriableError(err) || attempt == c.config.MaxRetries {
			return err
		}
	}
	return lastErr
}

func (c *Client) sendOnce(ctx context.Context, method string, path string, requestURL string, body []byte, opts RequestOptions, out any) error {
	request, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", authorizationValue(c.config.APIToken, c.config.AuthMode))
	request.Header.Set("Client", c.config.ClientHeader)
	if c.config.SiteID != "" {
		request.Header.Set("SiteId", c.config.SiteID)
	}
	for key, value := range c.config.AppHeaders {
		request.Header.Set(key, value)
	}
	for key, value := range opts.Headers {
		request.Header.Set(key, value)
	}
	if len(body) > 0 {
		request.Header.Set("Content-Type", "application/json")
	}
	httpClient := c.config.HTTPClient
	if opts.Timeout > 0 {
		cloned := *httpClient
		cloned.Timeout = opts.Timeout
		httpClient = &cloned
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return &APIError{StatusCode: response.StatusCode, Method: method, Path: path, Body: payload}
	}
	if len(bytes.TrimSpace(payload)) == 0 || response.StatusCode == http.StatusNoContent || out == nil {
		return nil
	}
	contentType := strings.ToLower(response.Header.Get("Content-Type"))
	if contentType != "" && !strings.Contains(contentType, "json") {
		return nil
	}
	if err := json.Unmarshal(payload, out); err != nil {
		return &ValidationError{Message: fmt.Sprintf("response for %s %s was not valid JSON: %v", method, path, err)}
	}
	return nil
}

func encodeJSONBody(body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, &ValidationError{Message: fmt.Sprintf("request body could not be encoded as JSON: %v", err)}
	}
	return payload, nil
}

func renderPath(pathTemplate string, pathParams map[string]any) (string, error) {
	rendered := pathTemplate
	missing := make([]string, 0)
	for {
		start := strings.Index(rendered, "{")
		if start < 0 {
			break
		}
		end := strings.Index(rendered[start:], "}")
		if end < 0 {
			break
		}
		name := rendered[start+1 : start+end]
		value, ok := pathParams[name]
		if !ok {
			missing = append(missing, name)
			rendered = strings.Replace(rendered, "{"+name+"}", "", 1)
			continue
		}
		rendered = strings.Replace(rendered, "{"+name+"}", url.PathEscape(fmt.Sprint(value)), 1)
	}
	if len(missing) > 0 {
		return "", &ValidationError{Message: "missing path parameters: " + strings.Join(missing, ", ")}
	}
	return rendered, nil
}

func isRetriableError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return shouldRetryStatus(apiErr.StatusCode)
	}
	return true
}

func sleepWithContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
