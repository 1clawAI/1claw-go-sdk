# 1claw-go-sdk

Go SDK for **1Claw Vault** — HSM-backed secret management for AI agents and humans.

## Status

Early development. See [milestones](.cursor/plans/go_sdk_milestones_688fb63f.plan.md) for the implementation plan.

## Requirements

- Go 1.21+
- OpenAPI spec at `../1claw-openapi-spec/openapi.yaml` (for regeneration)

## Regenerating the Client

From the repo root:

```bash
make generate
# or
SPEC_PATH=/path/to/openapi.yaml make generate
```

Or use the script:

```bash
./scripts/generate.sh
```

## Development

```bash
go build ./...
go test ./...
```

## License

PolyForm Noncommercial 1.0.0
