package oneclaw

import (
	"context"
	"fmt"
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
