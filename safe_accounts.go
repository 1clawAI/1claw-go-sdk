package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// AgentAccount is an on-chain account record for an agent (EOA or Safe).
type AgentAccount struct {
	ID             string                 `json:"id"`
	OrgID          string                 `json:"org_id"`
	AgentID        string                 `json:"agent_id"`
	Chain          string                 `json:"chain"`
	AccountType    string                 `json:"account_type"`
	Address        string                 `json:"address,omitempty"`
	SafeVersion    string                 `json:"safe_version,omitempty"`
	ModulesEnabled []string               `json:"modules_enabled,omitempty"`
	DeployStatus   string                 `json:"deploy_status,omitempty"`
	CosignEnabled  bool                   `json:"cosign_enabled,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at,omitempty"`
}

// AgentAccountList wraps list accounts response.
type AgentAccountList struct {
	Accounts []AgentAccount `json:"accounts"`
}

// ProvisionAgentAccountParams configures account provisioning.
type ProvisionAgentAccountParams struct {
	Chain       string `json:"chain"`
	AccountType string `json:"account_type,omitempty"`
	Address     string `json:"address,omitempty"`
}

// MigrateToSafeParams configures EOA→Safe migration.
type MigrateToSafeParams struct {
	Chain        string `json:"chain"`
	DeprecateEOA bool   `json:"deprecate_eoa,omitempty"`
}

// MigrationPlan is the counterfactual Safe migration plan.
type MigrationPlan struct {
	AgentID            string                   `json:"agent_id"`
	Chain              string                   `json:"chain"`
	SafeAddress        string                   `json:"safe_address"`
	SafeVersion        string                   `json:"safe_version"`
	Modules            []string                 `json:"modules"`
	EOAAddress         string                   `json:"eoa_address,omitempty"`
	SweepInstructions  []map[string]interface{} `json:"sweep_instructions"`
	RolesConfigHash    string                   `json:"roles_config_hash"`
	AllowanceConfigHash string                  `json:"allowance_config_hash"`
	Warnings           []string                 `json:"warnings"`
	DeployStatus       string                   `json:"deploy_status"`
}

// AllowanceReconcileReport summarizes org Safe allowance drift.
type AllowanceReconcileReport struct {
	OrgID         string                   `json:"org_id"`
	AgentsChecked int                      `json:"agents_checked"`
	Compiled      []map[string]interface{} `json:"compiled"`
	DriftDetected []map[string]interface{} `json:"drift_detected"`
	OnchainSync   string                   `json:"onchain_sync"`
}

// SafeModuleRegistry lists pinned module addresses for a chain.
type SafeModuleRegistry struct {
	Chain   string           `json:"chain"`
	Modules []SafeModuleInfo `json:"modules"`
}

// SafeModuleInfo is a pinned Safe or Zodiac module.
type SafeModuleInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Version string `json:"version,omitempty"`
}

// NotImplementedResponse is returned by Phase 5 stub endpoints.
type NotImplementedResponse struct {
	Error   string `json:"error"`
	Phase   string `json:"phase"`
	Message string `json:"message"`
}

// ListAccounts lists agent on-chain accounts.
func (s *AgentsService) ListAccounts(ctx context.Context, agentID string) (*AgentAccountList, error) {
	var resp AgentAccountList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/accounts", agentID), nil, &resp)
	return &resp, err
}

// ProvisionAccount creates an agent account record (human-only).
func (s *AgentsService) ProvisionAccount(ctx context.Context, agentID string, params ProvisionAgentAccountParams) (*AgentAccount, error) {
	var resp AgentAccount
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/accounts", agentID), params, &resp)
	return &resp, err
}

// MigrateToSafe builds an EOA→Safe migration plan (human-only).
func (s *AgentsService) MigrateToSafe(ctx context.Context, agentID string, params MigrateToSafeParams) (*MigrationPlan, error) {
	var resp MigrationPlan
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/accounts/migrate", agentID), params, &resp)
	return &resp, err
}

// DeprecateEOAAccount marks the agent EOA deprecated for a chain (human-only).
func (s *AgentsService) DeprecateEOAAccount(ctx context.Context, agentID, chain string) (*AgentAccount, error) {
	var resp AgentAccount
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/accounts/%s/deprecate-eoa", agentID, url.PathEscape(chain)), nil, &resp)
	return &resp, err
}

// ReplayGuardrails dry-runs draft guardrails against recent transactions.
func (s *AgentsService) ReplayGuardrails(ctx context.Context, agentID string, params map[string]interface{}) (map[string]interface{}, error) {
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/guardrails/replay", agentID), params, &resp)
	return resp, err
}

// GetSafeModuleRegistry returns pinned Safe module addresses (public).
func (s *OrgService) GetSafeModuleRegistry(ctx context.Context, chain string) (*SafeModuleRegistry, error) {
	var resp SafeModuleRegistry
	err := s.client.doJSONPublic(ctx, "GET", fmt.Sprintf("/v1/safe/module-registry/%s", url.PathEscape(chain)), nil, &resp)
	return &resp, err
}

// SyncOrgSafeAllowances reconciles org Safe allowance configs (owner/admin).
func (s *OrgService) SyncOrgSafeAllowances(ctx context.Context) (*AllowanceReconcileReport, error) {
	var resp AllowanceReconcileReport
	err := s.client.doJSON(ctx, "POST", "/v1/org/safe/sync-allowances", nil, &resp)
	return &resp, err
}

// GetGuardrailShadowReport returns Convention 6 shadow violations.
func (s *OrgService) GetGuardrailShadowReport(ctx context.Context, since, until string) (map[string]interface{}, error) {
	path := "/v1/org/guardrail-shadow-report"
	q := url.Values{}
	if since != "" {
		q.Set("since", since)
	}
	if until != "" {
		q.Set("until", until)
	}
	if enc := q.Encode(); enc != "" {
		path += "?" + enc
	}
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "GET", path, nil, &resp)
	return resp, err
}

// ListGuardrailRevisions lists guardrail revision history.
func (s *OrgService) ListGuardrailRevisions(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "GET", "/v1/org/guardrail-revisions", nil, &resp)
	return resp, err
}

// OnboardingStatus tracks welcome bundle and MCP readiness.
type OnboardingStatus struct {
	HasVault              bool   `json:"has_vault"`
	HasAgent              bool   `json:"has_agent"`
	HasPolicy             bool   `json:"has_policy"`
	HasSampleSecret       bool   `json:"has_sample_secret"`
	FirstSecretRead       bool   `json:"first_secret_read"`
	WelcomeBundleComplete bool   `json:"welcome_bundle_complete"`
	DefaultVaultID        string `json:"default_vault_id,omitempty"`
}

// OnboardingProvisionRequest configures MCP onboarding provision.
type OnboardingProvisionRequest struct {
	AgentName string `json:"agent_name,omitempty"`
	Client    string `json:"client,omitempty"`
}

// OnboardingProvisionResponse is returned by POST /v1/onboarding/provision.
type OnboardingProvisionResponse struct {
	AgentID        string                 `json:"agent_id"`
	APIKey         string                 `json:"api_key"`
	VaultID        string                 `json:"vault_id"`
	McpStdioConfig map[string]interface{} `json:"mcp_stdio_config"`
	VerifyPrompt   string                 `json:"verify_prompt"`
}

// GetOnboardingStatus returns org onboarding progress (human-only).
func (s *OrgService) GetOnboardingStatus(ctx context.Context) (*OnboardingStatus, error) {
	var resp OnboardingStatus
	err := s.client.doJSON(ctx, "GET", "/v1/org/onboarding/status", nil, &resp)
	return &resp, err
}

// ProvisionOnboarding creates welcome vault, sample secret, agent, and default policy.
func (s *OrgService) ProvisionOnboarding(ctx context.Context, params *OnboardingProvisionRequest) (*OnboardingProvisionResponse, error) {
	var body interface{}
	if params != nil {
		body = params
	} else {
		body = map[string]interface{}{}
	}
	var resp OnboardingProvisionResponse
	err := s.client.doJSON(ctx, "POST", "/v1/onboarding/provision", body, &resp)
	return &resp, err
}
