# Runtime Behavior

## Logging

The current Go runtime does not install a logger or emit structured logs.

Callers should log request context in their own application code and must avoid logging credentials, tokens, cookies, or raw Authorization header values.

## Retry Policy

Conservative retry defaults:
- Retryable statuses: `408`, `429`, `500`, `502`, `503`, `504`
- Retry methods: idempotent methods only (`GET`, `HEAD`, `OPTIONS`, `DELETE`, `PUT`)
- Backoff: exponential based on `BackoffBase`

Non-idempotent writes are not retried by default.

## Auth Behavior

Default auth mode is bearer:
- token `abc` -> `Authorization: Bearer abc`
- pre-prefixed token `Bearer abc` is preserved

Optional raw mode is available for compatibility.
