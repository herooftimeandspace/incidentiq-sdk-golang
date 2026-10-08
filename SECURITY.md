# Security Policy

## Supported Versions

Only the latest minor/patch release line is actively supported.

## Reporting a Vulnerability

Do not open public issues for undisclosed vulnerabilities.

Report security concerns privately to the project maintainers with:
- affected version
- impact summary
- reproduction details
- any known mitigation

Maintainers will acknowledge receipt, assess severity, and coordinate a fix/release.

## Secret Handling

- Never commit `INCIDENTIQ_API_TOKEN` or `INCIDENTIQ_TEST_API_TOKEN`.
- The SDK emits no logs, so it cannot leak an `Authorization` header on its own.
  If you wrap `Config.HTTPClient` with your own `RoundTripper`, redact the
  `Authorization` header before logging a request.
- `Config.APIToken` is returned by `Client.Config()`. Do not log that value or
  the struct that carries it.
- Keep tenant credentials in the environment (`INCIDENTIQ_*`) rather than in
  source, and keep integration credentials in repository secrets.
