package incidentiq

import "time"

// RequestOptions describes one low-level Incident IQ API request.
//
// Most callers should use generated or inventory-backed helpers as they are
// added. Request remains the drop-in escape hatch that matches incident-py-q's
// Client.request method.
type RequestOptions struct {
	PathParams  map[string]any
	Params      map[string]string
	JSON        any
	Body        []byte
	ContentType string
	Headers     map[string]string
	Timeout     time.Duration
	// OmitClientHeader skips the SDK-provided Client: ApiClient header for
	// this request. Default requests should leave this false; it exists for
	// HAR-validated Silver routes that intentionally omit the header.
	OmitClientHeader bool
	OmitSiteIDHeader bool
}
