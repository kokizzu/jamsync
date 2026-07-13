# JamSync

File synchronization service and CLI based on JamSync, with server, web, and
client command components.

Original download reference: https://jamsync.dev/download

## Requirements

- Go 1.26+

## Tests

```bash
go test ./internal/server/clientauth
```

## Maintenance Checklist

- [x] Return clear OAuth token parsing errors instead of panicking on malformed responses.
- [x] Add unit tests for access-token response parsing.
- [x] Update the Go module directive for the current Go toolchain.
