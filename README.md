# 1Claw Go SDK

Go client library for the [1Claw](https://1claw.xyz) Vault API — HSM-backed secret management for AI agents and humans.

## Install

```bash
go get github.com/1clawAI/1claw-go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    oneclaw "github.com/1clawAI/1claw-go-sdk"
)

func main() {
    client, _ := oneclaw.New(
        oneclaw.WithAPIKey("ocv_..."),
        oneclaw.WithAgentID("agent-uuid"),
    )

    ctx := context.Background()

    // List vaults
    vaults, _ := client.Vaults.List(ctx)
    fmt.Println(vaults)
}
```

## Authentication

```go
// API key (auto-exchanges for JWT)
client, _ := oneclaw.New(oneclaw.WithAPIKey("ocv_..."))

// Pre-authenticated JWT
client, _ := oneclaw.New(oneclaw.WithToken("eyJ..."))

// Agent credentials
client, _ := oneclaw.New(
    oneclaw.WithAPIKey("ocv_..."),
    oneclaw.WithAgentID("agent-uuid"),
)
```

## Services

| Service            | Description                                              |
| ------------------ | -------------------------------------------------------- |
| `Auth`             | Login, signup, agent tokens, federated tokens            |
| `Vaults`           | Create, list, get, delete vaults                         |
| `Secrets`          | Store, retrieve, rotate, delete secrets                  |
| `Agents`           | Manage agents, transactions, signing                     |
| `APIKeys`          | Create, list, revoke API keys                            |
| `Sharing`          | Share secrets via links or with users/agents              |
| `Access`           | Policy-based access control                              |
| `Org`              | Organization member management                           |
| `Chains`           | Blockchain chain registry                                |
| `Billing`          | Subscription, credits, usage                             |
| `Audit`            | Hash-chained audit event log                             |
| `X402`             | On-chain micropayments                                   |
| `Treasury`         | Safe multisig treasuries, wallets, delegations             |
| `Treasury.Wallets` | Multi-chain wallet management: generate, list, get, balance, send, swap, export, rotate, deactivate |
| `Treasury.Proposals` | Treasury proposals: create, list, get, sign, execute, cancel |
| `SigningKeys`      | Multi-chain signing key management (create, list, rotate, deactivate, export) |
| `Bindings`         | Execution Intents — bindings CRUD, execute, rotate credential, history |
| `Platform`         | Platform API — build multi-tenant apps on 1Claw          |
| `Webhooks`         | Register and manage event webhooks                       |
| `Risk`             | Risk events, verdicts, honeytokens (v0.36+)               |
| `Approvals`        | Human-in-the-loop approval workflow                      |
| `Automations`      | Cron-based scheduled tasks for agents                    |
| `Memory`           | Persistent vector memory — store and search              |
| `Runtimes`         | Managed runtime environments — deploy, scale, monitor    |
| `Discovery`        | Agent directory — publish, search, manage listings       |

## Execution Intents (Bindings)

```go
// Create a binding with an inline credential
binding, _ := client.Bindings.Create(ctx, agentID, oneclaw.CreateBindingParams{
    Name:        "httpbin",
    BindingType: "http",
    Config:      map[string]interface{}{"base_url": "https://httpbin.org"},
    Credential:  map[string]interface{}{"token": "secret"},
})

// Create a binding with a vault_ref credential (live-pointer to an existing secret)
binding2, _ := client.Bindings.Create(ctx, agentID, oneclaw.CreateBindingParams{
    Name:        "stripe-api",
    BindingType: "http",
    Config:      map[string]interface{}{"base_url": "https://api.stripe.com"},
    CredentialSource: &oneclaw.CredentialSource{
        Type:    "vault_ref",
        VaultID: vaultID,
        Path:    "integrations/stripe-key",
    },
})

result, _ := client.Bindings.Execute(ctx, agentID, oneclaw.ExecuteParams{
    Binding:    "httpbin",
    IntentType: "http",
    Params:     map[string]interface{}{"method": "GET", "path": "/get"},
})

_, _ = client.Bindings.RotateCredential(ctx, agentID, binding.ID, oneclaw.RotateCredentialParams{
    Credential: map[string]interface{}{"token": "new-secret"},
})

events, _ := client.Bindings.ListExecutions(ctx, agentID, nil, nil)
```

## Platform API

The Platform API lets developers build applications on top of 1Claw, provisioning users, vaults, agents, and policies on behalf of end-users via bootstrap templates.

```go
// Create a platform app
app, _ := client.Platform.CreateApp(ctx, oneclaw.CreatePlatformAppRequest{
    Name: "my-app",
})

// List platform apps
apps, _ := client.Platform.ListApps(ctx)

// Upsert a user
user, _ := client.Platform.UpsertUser(ctx, oneclaw.UpsertPlatformUserRequest{
    Email: "user@example.com",
})

// Bootstrap resources from a template
result, _ := client.Platform.BootstrapUser(ctx, connectionID)
```

## DPoP (Proof-of-Possession)

```go
client, _ := oneclaw.New(
    oneclaw.WithAPIKey("ocv_..."),
    oneclaw.WithDPoP(true),
)
```

> **Note:** For the full v0.42 API surface (non-EVM transaction signing, OAuth, email OTP, spend policies, deposit destinations, fiat ramps, and internal accounts), see the [TypeScript SDK](https://www.npmjs.com/package/@1claw/sdk) and the [OpenAPI spec](https://www.npmjs.com/package/@1claw/openapi-spec).

## Automations

```go
// Create a cron-based automation (workflow_spec required)
auto, _ := client.Automations.Create(ctx, oneclaw.CreateAutomationParams{
    Name:        "rotate-api-key",
    AgentID:     agentID,
    TriggerType: "cron",
    CronExpr:    "0 0 * * 0",
    Timezone:    "UTC",
    WorkflowSpec: map[string]interface{}{
        "steps": []map[string]interface{}{
            {"type": "log", "action": "run_agent_task", "message": "Rotate weekly API keys"},
        },
    },
})

// List automations in the org
list, _ := client.Automations.List(ctx)

// Trigger manually
_, _ = client.Automations.Trigger(ctx, auto.ID, nil)
```

## Agent Memory

```go
// Store a memory entry
entry, _ := client.Memory.Store(ctx, agentID, oneclaw.StoreMemoryParams{
    Content:   "User prefers JSON output",
    Namespace: "preferences",
})

// Semantic search
results, _ := client.Memory.Search(ctx, agentID, oneclaw.SearchMemoryParams{
    Query: "output format preferences",
    Limit: intPtr(5),
})

// Clear all memory
_ = client.Memory.Clear(ctx, agentID)
```

## Runtimes

```go
// Deploy a runtime
runtime, _ := client.Runtimes.Create(ctx, oneclaw.CreateRuntimeParams{
    AgentID: agentID,
    Name:    "my-agent",
    Image:   "ghcr.io/my-org/agent:latest",
    Env:     map[string]string{"MODEL": "gpt-4"},
})

// Restart
_ = client.Runtimes.Restart(ctx, runtime.ID)
```

## Discovery

```go
// Publish to directory
listing, _ := client.Discovery.Publish(ctx, agentID, oneclaw.PublishParams{
    Description: "Automated treasury management agent",
    Tags:        []string{"defi", "treasury"},
})

// Search
results, _ := client.Discovery.Search(ctx, "treasury management")
```

## Options

| Option             | Description                                |
| ------------------ | ------------------------------------------ |
| `WithBaseURL`      | Override API base URL                      |
| `WithToken`        | Set a pre-obtained JWT                     |
| `WithAPIKey`       | Set API key (auto-exchanges for JWT)       |
| `WithAgentID`      | Set agent ID for agent token flow          |
| `WithHTTPClient`   | Custom HTTP client                         |
| `WithUserAgent`    | Custom User-Agent header                   |
| `WithDebug`        | Enable debug logging                       |

## License

[MIT](./LICENSE)
