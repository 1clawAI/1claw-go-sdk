package oneclaw

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// PlatformService handles Platform API operations.
type PlatformService struct {
	client *Client
}

// CreateApp registers a new platform app. Returns the app and its one-time API key.
func (s *PlatformService) CreateApp(ctx context.Context, params CreatePlatformAppRequest) (*PlatformAppCreated, error) {
	var result PlatformAppCreated
	err := s.client.doJSON(ctx, "POST", "/v1/platform/apps", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListApps lists platform apps for the current organization.
func (s *PlatformService) ListApps(ctx context.Context) (*PlatformAppList, error) {
	var result PlatformAppList
	err := s.client.doJSON(ctx, "GET", "/v1/platform/apps", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApp retrieves a specific platform app by ID.
func (s *PlatformService) GetApp(ctx context.Context, appID string) (*PlatformApp, error) {
	var result PlatformApp
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/apps/%s", appID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateApp updates a platform app.
func (s *PlatformService) UpdateApp(ctx context.Context, appID string, params UpdatePlatformAppRequest) (*PlatformApp, error) {
	var result PlatformApp
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/platform/apps/%s", appID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteApp soft-deletes a platform app and releases its slug.
func (s *PlatformService) DeleteApp(ctx context.Context, appID string) (*PlatformAppDeleteResponse, error) {
	var result PlatformAppDeleteResponse
	err := s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/platform/apps/%s", appID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PlatformAppDeleteResponse is returned when a platform app is soft-deleted.
type PlatformAppDeleteResponse struct {
	ID           string `json:"id"`
	Slug         string `json:"slug"`
	DeletedAt    string `json:"deleted_at"`
	SlugReleased bool   `json:"slug_released"`
}

// TransferAppOwnershipRequest moves a platform app to another org.
type TransferAppOwnershipRequest struct {
	TargetOrgID       string `json:"target_org_id,omitempty"`
	TargetUserEmail   string `json:"target_user_email,omitempty"`
}

// TransferAppOwnershipResponse confirms an ownership transfer.
type TransferAppOwnershipResponse struct {
	AppID        string `json:"app_id"`
	FormerOrgID  string `json:"former_org_id"`
	NewOrgID     string `json:"new_org_id"`
}

// TransferAppOwnership moves a platform app to another organization (step-up required).
func (s *PlatformService) TransferAppOwnership(ctx context.Context, appID string, params TransferAppOwnershipRequest, confirmToken string) (*TransferAppOwnershipResponse, error) {
	var result TransferAppOwnershipResponse
	headers := map[string]string{}
	if confirmToken != "" {
		headers["X-Auth-Confirm"] = confirmToken
	}
	err := s.client.doJSONWithHeaders(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/transfer-ownership", appID), params, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSpendPolicy returns a spend policy by ID.
func (s *PlatformService) GetSpendPolicy(ctx context.Context, appID, policyID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/apps/%s/spend-policies/%s", appID, policyID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConnectionSpendPolicy returns the effective spend policy for a connection.
func (s *PlatformService) GetConnectionSpendPolicy(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/spend-policy", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListConnectionApprovals lists approvals for a connected user (plt_ auth).
func (s *PlatformService) ListConnectionApprovals(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/approvals", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConnectionApproval returns a single approval for a connection.
func (s *PlatformService) GetConnectionApproval(ctx context.Context, connectionID, approvalID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/approvals/%s", connectionID, approvalID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListConnectionPendingApprovals lists consensus pending approvals for a connection.
func (s *PlatformService) ListConnectionPendingApprovals(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/pending-approvals", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpsertPlatformUserRequest are parameters for provisioning or finding a user.
type UpsertPlatformUserRequest struct {
	Email              string `json:"email,omitempty"`
	SubjectToken       string `json:"subject_token,omitempty"`
	SubjectTokenType   string `json:"subject_token_type,omitempty"`
	SiweMessage        string `json:"siwe_message,omitempty"`
	SiweSignature      string `json:"siwe_signature,omitempty"`
}

// BootstrapUserRequest are parameters for bootstrapping a connected user.
type BootstrapUserRequest struct {
	TemplateID string                 `json:"template_id,omitempty"`
	ReturnTo   string                 `json:"return_to,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// UpsertPlatformUserResponse is the response from upserting a platform user.
type UpsertPlatformUserResponse struct {
	UserHandle   string `json:"user_handle"`
	ConnectionID string `json:"connection_id"`
	IsNew        bool   `json:"is_new"`
}

// UpsertUser provisions or finds a user via email or subject token.
func (s *PlatformService) UpsertUser(ctx context.Context, params UpsertPlatformUserRequest) (*UpsertPlatformUserResponse, error) {
	var result UpsertPlatformUserResponse
	err := s.client.doJSON(ctx, "POST", "/v1/platform/users/upsert", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// BootstrapUserResponse is the response from bootstrapping resources for a connected user.
type BootstrapUserResponse struct {
	ClaimURL     string                 `json:"claim_url"`
	ClaimToken   string                 `json:"claim_token"`
	ConnectionID string                 `json:"connection_id"`
	Summary      map[string]interface{} `json:"summary"`
}

// BootstrapUser bootstraps resources for a connected user from a template.
func (s *PlatformService) BootstrapUser(ctx context.Context, connectionID string, params *BootstrapUserRequest) (*BootstrapUserResponse, error) {
	var result BootstrapUserResponse
	body := params
	if body == nil {
		body = &BootstrapUserRequest{}
	}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/connections/%s/bootstrap", connectionID), body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SiweChallengeResponse is returned from POST /v1/platform/siwe/challenge.
type SiweChallengeResponse struct {
	Nonce     string `json:"nonce"`
	ExpiresIn int    `json:"expires_in"`
	Domain    string `json:"domain"`
}

// SiweChallenge issues a SIWE nonce for wallet-native provisioning.
func (s *PlatformService) SiweChallenge(ctx context.Context, domain string) (*SiweChallengeResponse, error) {
	var result SiweChallengeResponse
	body := map[string]string{}
	if domain != "" {
		body["domain"] = domain
	}
	err := s.client.doJSON(ctx, "POST", "/v1/platform/siwe/challenge", body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetConnection returns connection details including claim and entitlement status.
func (s *PlatformService) GetConnection(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConnectionUsage returns per-connection inference spend for the current month.
func (s *PlatformService) GetConnectionUsage(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/usage", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEntitlements lists on-chain entitlement watches for a connection.
func (s *PlatformService) ListEntitlements(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/entitlements", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PreviewTemplate resolves template placeholders without provisioning resources.
func (s *PlatformService) PreviewTemplate(ctx context.Context, appID, templateID string, params map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/templates/%s/preview", appID, templateID), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListConnectedApps lists platform apps connected to the calling user.
func (s *PlatformService) ListConnectedApps(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.client.doJSON(ctx, "GET", "/v1/platform/connected-apps", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DisconnectApp disconnects a user from a platform app.
func (s *PlatformService) DisconnectApp(ctx context.Context, connectionID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/platform/connected-apps/%s", connectionID), nil, nil)
}

// MarketplaceResponse is the response from browsing the platform marketplace.
type MarketplaceResponse struct {
	Apps []map[string]interface{} `json:"apps"`
}

// Marketplace browses the public platform marketplace (no auth required).
func (s *PlatformService) Marketplace(ctx context.Context) (*MarketplaceResponse, error) {
	var result MarketplaceResponse
	err := s.client.doJSON(ctx, "GET", "/v1/platform/marketplace", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PlatformAppStatsResponse contains aggregate statistics for a platform app.
type PlatformAppStatsResponse struct {
	TotalConnections  int `json:"total_connections"`
	ActiveConnections int `json:"active_connections"`
	ClaimedConnections int `json:"claimed_connections"`
	TotalBootstraps   int `json:"total_bootstraps"`
	TotalGrants       int `json:"total_grants"`
}

// GetAppStats returns aggregate statistics for a platform app.
func (s *PlatformService) GetAppStats(ctx context.Context, appID string) (*PlatformAppStatsResponse, error) {
	var result PlatformAppStatsResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/apps/%s/stats", appID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RotateWebhookSecretResponse is the response from rotating a platform app's
// webhook signing secret.
type RotateWebhookSecretResponse struct {
	WebhookSecret string `json:"webhook_secret"`
}

// RotateWebhookSecret rotates a platform app's webhook signing secret.
// The new secret is returned once and cannot be retrieved again.
func (s *PlatformService) RotateWebhookSecret(ctx context.Context, appID string) (*RotateWebhookSecretResponse, error) {
	var result RotateWebhookSecretResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/rotate-webhook-secret", appID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetConnectionRuntime returns a runtime provisioned on a connection (plt_ auth).
func (s *PlatformService) GetConnectionRuntime(ctx context.Context, connectionID, runtimeID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/connections/%s/runtimes/%s", connectionID, runtimeID), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConnectionPasskeyEnrollBegin starts WebAuthn registration for a connected end-user (plt_ auth).
func (s *PlatformService) ConnectionPasskeyEnrollBegin(ctx context.Context, connectionID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/connections/%s/passkeys/enroll/begin", connectionID), map[string]interface{}{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ── Fleets ──────────────────────────────────────────────────────────────
//
// A fleet is every agent one bootstrap template provisioned. Each call below
// acts on all of them at once, which is why the surface is narrower than the
// per-agent API rather than wider: guardrails and capability flags are not
// bulk-patchable, and one bad field refuses the whole patch.

// FleetVersionBucket is how many agents were provisioned from one template version.
type FleetVersionBucket struct {
	TemplateVersion *int  `json:"template_version"`
	Agents          int64 `json:"agents"`
}

// FleetSummary describes a template's cohort.
type FleetSummary struct {
	TemplateID   string `json:"template_id"`
	TemplateName string `json:"template_name"`
	// CurrentVersion is the template's version; agents provisioned from an
	// earlier one are behind.
	CurrentVersion int `json:"current_version"`
	// SpecHash distinguishes a version bump that changed nothing from one that
	// did. Empty on templates written before migration 245.
	SpecHash               string               `json:"spec_hash,omitempty"`
	TotalAgents            int64                `json:"total_agents"`
	VersionSkew            []FleetVersionBucket `json:"version_skew"`
	AgentsOnCurrentVersion int64                `json:"agents_on_current_version"`
	AgentsBehind           int64                `json:"agents_behind"`
	// DriftedAgents were changed outside fleet control; a rollout skips them.
	DriftedAgents int64 `json:"drifted_agents"`
	// BulkPatchableFields is the authoritative allowlist. Read it from here
	// rather than hard-coding it: it excludes guardrails and capability flags,
	// and may narrow further.
	BulkPatchableFields []string `json:"bulk_patchable_fields"`
}

// FleetAgent is one member of a cohort.
type FleetAgent struct {
	AgentID                string `json:"agent_id"`
	Name                   string `json:"name"`
	OrgID                  string `json:"org_id"`
	PlatformConnectionID   string `json:"platform_connection_id,omitempty"`
	ProvisionedFromVersion *int   `json:"provisioned_from_version"`
	LastFleetSyncAt        string `json:"last_fleet_sync_at,omitempty"`
	// DriftFields are the fields a rollout skipped because a human changed them.
	DriftFields []string `json:"drift_fields"`
	IsActive    bool     `json:"is_active"`
	IsCurrent   bool     `json:"is_current"`
}

// ListFleetAgentsResponse is a page of a fleet's agents.
type ListFleetAgentsResponse struct {
	Agents         []FleetAgent `json:"agents"`
	Limit          int64        `json:"limit"`
	Offset         int64        `json:"offset"`
	CurrentVersion int          `json:"current_version"`
}

// BulkPatchFleetResponse reports what a bulk patch touched.
type BulkPatchFleetResponse struct {
	FieldsApplied []string `json:"fields_applied"`
	AgentsMatched int64    `json:"agents_matched"`
	AgentsUpdated int64    `json:"agents_updated"`
}

// FleetRolloutOutcome is what a rollout decided about one agent. Outcome is one
// of "already_current", "synced", or "skipped_drifted".
type FleetRolloutOutcome struct {
	Outcome     string   `json:"outcome"`
	AgentID     string   `json:"agent_id"`
	Fields      []string `json:"fields,omitempty"`
	DriftFields []string `json:"drift_fields,omitempty"`
}

// FleetRolloutRequest controls a rollout.
type FleetRolloutRequest struct {
	// Force overwrites hand edits. It still cannot carry a guardrail.
	Force bool `json:"force"`
	// DryRun reports the plan without applying it, and claims no job.
	DryRun bool `json:"dry_run"`
}

// FleetRolloutResponse is the result of a rollout.
type FleetRolloutResponse struct {
	// JobID is empty for a dry run, which claims no job.
	JobID          string                `json:"job_id,omitempty"`
	ToVersion      int                   `json:"to_version"`
	DryRun         bool                  `json:"dry_run"`
	Forced         bool                  `json:"forced"`
	TotalAgents    int64                 `json:"total_agents"`
	Synced         int                   `json:"synced"`
	AlreadyCurrent int                   `json:"already_current"`
	SkippedDrifted int                   `json:"skipped_drifted"`
	Outcomes       []FleetRolloutOutcome `json:"outcomes"`
}

// PauseFleetResponse reports how many agents were deactivated.
type PauseFleetResponse struct {
	AgentsPaused int64 `json:"agents_paused"`
}

// GetFleet returns a template's cohort: size, version skew, and drift.
func (s *PlatformService) GetFleet(ctx context.Context, appID, templateID string) (*FleetSummary, error) {
	var result FleetSummary
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/platform/apps/%s/fleets/%s", appID, templateID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFleetAgents lists the agents in a fleet. Pass limit/offset <= 0 to omit them.
func (s *PlatformService) ListFleetAgents(ctx context.Context, appID, templateID string, limit, offset int) (*ListFleetAgentsResponse, error) {
	path := fmt.Sprintf("/v1/platform/apps/%s/fleets/%s/agents", appID, templateID)
	q := url.Values{}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result ListFleetAgentsResponse
	if err := s.client.doJSON(ctx, "GET", path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkPatchFleet applies one patch to every agent in the cohort.
//
// Guardrails and capability flags (intents_api_enabled,
// execution_intents_enabled) are refused with a 400 naming the field, and one
// bad field refuses the whole patch rather than applying it in part. Read the
// current allowlist from GetFleet().BulkPatchableFields.
func (s *PlatformService) BulkPatchFleet(ctx context.Context, appID, templateID string, patch map[string]interface{}) (*BulkPatchFleetResponse, error) {
	body := map[string]interface{}{"patch": patch}
	var result BulkPatchFleetResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/fleets/%s/bulk-patch", appID, templateID), body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RolloutFleet brings the cohort up to the template's current version.
//
// An agent changed outside fleet control is skipped rather than corrected.
// Force overrides that but still cannot carry a guardrail. DryRun reports the
// plan, claims no job (JobID comes back empty), and so never blocks the real
// rollout that follows it. Only one rollout runs per template at a time; a
// second returns 409.
func (s *PlatformService) RolloutFleet(ctx context.Context, appID, templateID string, params FleetRolloutRequest) (*FleetRolloutResponse, error) {
	var result FleetRolloutResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/fleets/%s/rollout", appID, templateID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PauseFleet deactivates every agent in the cohort.
func (s *PlatformService) PauseFleet(ctx context.Context, appID, templateID string) (*PauseFleetResponse, error) {
	var result PauseFleetResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/apps/%s/fleets/%s/pause", appID, templateID), map[string]interface{}{}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
