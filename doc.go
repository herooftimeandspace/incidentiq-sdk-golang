// Package incidentiq provides a Go client for the Incident IQ API.
//
// The package is designed to stay functionally aligned with the
// herooftimeandspace/incident-py-q source SDK. It uses the same bundled Golden
// Stoplight contracts, Silver HAR-derived inventory, authentication headers,
// URL normalization rules, retry policy, and environment variable names.
package incidentiq

//go:generate go run scripts/generate_wrappers.go
