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

// DeleteApp deletes a platform app.
func (s *PlatformService) DeleteApp(ctx context.Context, appID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/platform/apps/%s", appID), nil, nil)
}

// UpsertPlatformUserRequest are parameters for provisioning or finding a user.
type UpsertPlatformUserRequest struct {
	Email        string `json:"email,omitempty"`
	SubjectToken string `json:"subject_token,omitempty"`
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
func (s *PlatformService) BootstrapUser(ctx context.Context, connectionID string) (*BootstrapUserResponse, error) {
	var result BootstrapUserResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/platform/connections/%s/bootstrap", connectionID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
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
