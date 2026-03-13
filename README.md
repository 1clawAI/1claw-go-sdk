# 1claw-go-sdk

Go SDK for **1Claw Vault** — HSM-backed secret management for AI agents and humans.

## Install

```bash
go get github.com/1clawAI/1claw-go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/1clawAI/1claw-go-sdk"
)

func main() {
    client, err := oneclaw.New(
        oneclaw.WithBaseURL("https://api.1claw.xyz"),
        oneclaw.WithAPIKey("ocv_..."),  // auto-exchanges for JWT
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // List vaults
    vaults, err := client.Vaults.ListVaults(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Get a secret
    secret, err := client.Secrets.GetSecret(ctx, "vault-id", "path/to/secret")

    // Create an API key
    created, err := client.APIKeys.CreateAPIKey(ctx, "my-key", []string{"vault:read"})
}
```

## Authentication

```go
// 1. API key (auto-exchanges for JWT on first call)
client, _ := oneclaw.New(
    oneclaw.WithBaseURL("https://api.1claw.xyz"),
    oneclaw.WithAPIKey("ocv_..."),
)

// 2. Agent credentials
client, _ := oneclaw.New(
    oneclaw.WithBaseURL("https://api.1claw.xyz"),
    oneclaw.WithAPIKey("ocv_..."),
    oneclaw.WithAgentID("agent-uuid"),
)

// 3. Pre-obtained JWT
client, _ := oneclaw.New(
    oneclaw.WithBaseURL("https://api.1claw.xyz"),
    oneclaw.WithToken("eyJ..."),
)
```

## API Resources

| Resource | Methods |
|----------|---------|
| `client.Auth` | ApiKeyToken, AgentToken, Login, GetMe |
| `client.Vaults` | Create, List, Get, Delete |
| `client.Secrets` | Put, Get, Delete, List |
| `client.Agents` | Create, Get, List, Update, Delete |
| `client.APIKeys` | Create, List, Revoke |
| `client.Sharing` | CreateShare, ListOutbound, ListInbound, Revoke |
| `client.Access` | CreatePolicy, ListPolicies, UpdatePolicy, DeletePolicy |
| `client.Org` | ListMembers, InviteMember, UpdateMemberRole, RemoveMember |
| `client.Chains` | List, Get |
| `client.Billing` | GetSubscription, GetCreditBalance |
| `client.Audit` | QueryAuditEvents |
| `client.X402` | Payment protocol (X402Signer interface) |

## CMEK (Customer-Managed Encryption Keys)

```go
import "github.com/1clawAI/1claw-go-sdk/cmek"

// Generate a key
key, err := cmek.GenerateKey()

// Compute fingerprint
fingerprint := cmek.Fingerprint(key)

// Encrypt before storing
encrypted, _ := cmek.Encrypt(key, "my-secret")
client.Secrets.PutSecret(ctx, vaultID, "path", encrypted, "generic")

// Decrypt after retrieving
secret, _ := client.Secrets.GetSecret(ctx, vaultID, "path")
plaintext, _ := cmek.Decrypt(key, secret.Value)
```

## Using in Other Projects

Add to your `go.mod`:

```go
require github.com/1clawAI/1claw-go-sdk v0.1.0
```

For local development:

```go
replace github.com/1clawAI/1claw-go-sdk => /path/to/1claw-go-sdk
```

## Regenerating the Client

From the repo root:

```bash
make generate
# or
SPEC_PATH=/path/to/openapi.yaml make generate
```

## Development

```bash
go build ./...
go test ./...
```

## Versioning

SDK versions track compatible OpenAPI spec versions. Regenerate when the spec changes.

## License

PolyForm Noncommercial 1.0.0
