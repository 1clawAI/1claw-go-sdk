# 1Claw Go SDK

> ⭐ **Star [1clawAI/agent-templates](https://github.com/1clawAI/agent-templates)** — ready-to-run agent templates wired to 1Claw. It is our single starred repo.

Go client for the [1Claw](https://1claw.co) Vault API.

Built for backend services and infrastructure code written in Go. You get typed structs for vaults, secrets, agents, policies, treasury, Intents API signing, execution bindings, and billing. Agent API keys exchange for JWTs automatically, same as the TypeScript and Python SDKs.

Reach for this when your stack is already Go (microservices, operators, CI workers) and you want compile-time types instead of raw HTTP. For AI agent frameworks, use the MCP server or a language-specific integration package.

## Graduated HITL (v0.54–0.55)

`CreateAgentParams` / `UpdateAgentParams` / `Agent` include v0.55 guardrail fields: `TxApprovalPolicy`, `TypedDataPolicy`, `SimulationFailurePolicy`, `RawSigningPolicy`, USD caps, per-recipient limits, `AllowErc4337`, `AllowEip7702`, and `AutoSuspended`. Org emergency freeze: `POST /v1/org/freeze`.

## v0.56 — Safe accounts, guardrail governance, HFA

```go
// Agent on-chain accounts
accounts, _ := client.Agents.ListAccounts(ctx, agentID)
plan, _ := client.Agents.MigrateToSafe(ctx, agentID, oneclaw.MigrateToSafeParams{Chain: "ethereum"})
registry, _ := client.Org.GetSafeModuleRegistry(ctx, "ethereum") // public
report, _ := client.Org.SyncOrgSafeAllowances(ctx)

// Guardrail governance
shadow, _ := client.Org.GetGuardrailShadowReport(ctx, "", "")
revisions, _ := client.Org.ListGuardrailRevisions(ctx)
replay, _ := client.Agents.ReplayGuardrails(ctx, agentID, map[string]interface{}{
    "draft_guardrails": map[string]interface{}{"tx_max_value_eth": "0.1"},
})

// Human Factor Auth
hfa, _ := client.Auth.GetHumanFactorAuth(ctx)
authPolicy, _ := client.Treasury.GetWalletAuthPolicy(ctx)
```

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
| `Agents`           | Manage agents, transactions, signing, delegations        |
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

## Agent Delegation

```go
// Create a delegation (human-only)
delegation, _ := client.Agents.CreateDelegation(ctx, orchestratorID, oneclaw.CreateDelegationParams{
    DelegateID:     subAgentID,
    AllowedTools:   []string{"delegate_task"},
    DelegationMode: "caller",
})

// List delegations
list, _ := client.Agents.ListDelegations(ctx, agentID)

// Get effective delegations (for runtime tool discovery)
effective, _ := client.Agents.GetEffectiveDelegations(ctx, agentID)

// Revoke a delegation
_ = client.Agents.RevokeDelegation(ctx, agentID, delegationID)
```

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

// Browse the public marketplace (no auth required)
marketplace, _ := client.Platform.Marketplace(ctx)

// Get app stats (connections, bootstraps, grants)
stats, _ := client.Platform.GetAppStats(ctx, appID)

// Rotate webhook signing secret (returns new secret once)
secret, _ := client.Platform.RotateWebhookSecret(ctx, appID)

// Platform expansion (v0.57+)
challenge, _ := client.Platform.SiweChallenge(ctx, nil)
conn, _ := client.Platform.GetConnection(ctx, connectionID)
usage, _ := client.Platform.GetConnectionUsage(ctx, connectionID)
entitlements, _ := client.Platform.ListEntitlements(ctx, connectionID)
preview, _ := client.Platform.PreviewTemplate(ctx, appID, templateID, &oneclaw.PreviewTemplateParams{
    Parameters: map[string]interface{}{"agent_name": "demo"},
})

// Platform control plane (v0.58+)
_, _ = client.Platform.TransferAppOwnership(ctx, appID, oneclaw.TransferAppOwnershipRequest{
    TargetOrgID: targetOrgID,
}, confirmToken)
policy, _ := client.Platform.GetSpendPolicy(ctx, appID, policyID)
connPolicy, _ := client.Platform.GetConnectionSpendPolicy(ctx, connectionID)
approvals, _ := client.Platform.ListConnectionApprovals(ctx, connectionID)
approval, _ := client.Platform.GetConnectionApproval(ctx, connectionID, approvalID)
pending, _ := client.Platform.ListConnectionPendingApprovals(ctx, connectionID)
_, _ = client.Platform.GetConnectionRuntime(ctx, connectionID, runtimeID)
_, _ = client.Platform.ConnectionPasskeyEnrollBegin(ctx, connectionID)
```

**v0.59.4 connection endpoints** (`portfolio`, `pending-approvals` create, connection-scoped `automations`/`memory`, `inspect-content`) are in the OpenAPI spec and TypeScript/Python SDKs; Go client methods are pending — call the REST API directly until the next Go SDK release.

## DPoP (Proof-of-Possession)

```go
client, _ := oneclaw.New(
    oneclaw.WithAPIKey("ocv_..."),
    oneclaw.WithDPoP(true),
)
```

## OAuth Token & Consent Management

```go
// Revoke an OAuth token (RFC 7009)
resp, _ := client.OAuthConnect.RevokeToken(ctx, "eyJ...", "access_token")

// Revoke consent for a platform app (deletes consent + revokes all tokens)
_ = client.OAuthConnect.RevokeConsent(ctx, appID)
```

> **Note:** For the full v0.42 API surface (non-EVM transaction signing, email OTP, spend policies, deposit destinations, fiat ramps, and internal accounts), see the [TypeScript SDK](https://www.npmjs.com/package/@1claw/sdk) and the [OpenAPI spec](https://www.npmjs.com/package/@1claw/openapi-spec).

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

## Channels

```go
// Register a Telegram channel for an agent
ch, _ := client.Channels.Create(ctx, agentID, oneclaw.CreateChannelParams{
    ChannelType: "telegram",
    ChannelName: "Support Bot",
})

// Send a message
msg, _ := client.Channels.SendMessage(ctx, agentID, ch.ID, "Hello!", nil)

// List message history
messages, _ := client.Channels.ListMessages(ctx, agentID, ch.ID)
```

## Webhooks

```go
// Register a webhook
wh, _ := client.Webhooks.Create(ctx, oneclaw.CreateWebhookParams{
    URL:    "https://example.com/hooks",
    Events: []string{"agent.transaction.broadcast", "policy.created"},
})

// List webhooks
hooks, _ := client.Webhooks.List(ctx)
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
