# 1claw-go-sdk

Go SDK for **1Claw Vault** — HSM-backed secret management for AI agents and humans.

## Quick Start

```go
import "github.com/1clawAI/1claw-go-sdk"

client, err := oneclaw.New(
    oneclaw.WithBaseURL("https://api.1claw.xyz"),
    oneclaw.WithAPIKey("ocv_..."),  // auto-exchanges for JWT
)
if err != nil {
    log.Fatal(err)
}

// List vaults
vaults, err := client.Vaults.ListVaults(ctx)

// Get a secret
secret, err := client.Secrets.GetSecret(ctx, "vault-id", "path/to/secret")

// Create an API key
created, err := client.APIKeys.CreateAPIKey(ctx, "my-key", []string{"vault:read"})
```

## Status

Early development. Core resources (Auth, Vaults, Secrets, Agents, APIKeys) implemented.

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
