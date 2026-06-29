package incidentiq

import "time"

var retryableStatuses = map[int]bool{
	408: true,
	429: true,
	500: true,
	502: true,
	503: true,
	504: true,
}

var idempotentMethods = map[string]bool{
	"GET":     true,
	"HEAD":    true,
	"OPTIONS": true,
	"DELETE":  true,
	"PUT":     true,
}

func shouldRetryStatus(statusCode int) bool {
	return retryableStatuses[statusCode]
}

func methodIsIdempotent(method string) bool {
	return idempotentMethods[method]
}

func computeBackoff(attempt int, base time.Duration) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	return time.Duration(1<<attempt) * base
}
