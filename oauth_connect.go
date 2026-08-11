package oneclaw

import (
	"context"
	"fmt"
)

// OAuthProviderScope represents a single scope offered by an OAuth provider.
type OAuthProviderScope struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// OAuthProvider represents a registered OAuth provider.
type OAuthProvider struct {
	Slug     string               `json:"slug"`
	Name     string               `json:"name"`
	AuthURL  string               `json:"auth_url,omitempty"`
	TokenURL string               `json:"token_url,omitempty"`
	Scopes   []OAuthProviderScope `json:"scopes,omitempty"`
	LogoURL  string               `json:"logo_url,omitempty"`
}

// OAuthProviderListResponse is the response from listing OAuth providers.
type OAuthProviderListResponse struct {
	Providers []OAuthProvider `json:"providers"`
}

// OAuthConnection represents an active OAuth connection for an agent.
type OAuthConnection struct {
	ID                  string   `json:"id"`
	AgentID             string   `json:"agent_id"`
	ProviderSlug        string   `json:"provider_slug"`
	Status              string   `json:"status,omitempty"`
	Scopes              []string `json:"scopes,omitempty"`
	ExternalAccountID   string   `json:"external_account_id,omitempty"`
	ExternalAccountName string   `json:"external_account_name,omitempty"`
	CreatedAt           string   `json:"created_at,omitempty"`
	ExpiresAt           string   `json:"expires_at,omitempty"`
}

// OAuthConnectionListResponse is the response from listing OAuth connections.
type OAuthConnectionListResponse struct {
	Connections []OAuthConnection `json:"connections"`
}

// ConnectOAuthRequest are parameters for initiating an OAuth connection.
type ConnectOAuthRequest struct {
	ProviderSlug  string   `json:"provider_slug"`
	Scopes        []string `json:"scopes,omitempty"`
	RedirectAfter string   `json:"redirect_after,omitempty"`
}

// ConnectOAuthResponse is the response from initiating an OAuth connection.
type ConnectOAuthResponse struct {
	AuthorizeURL string `json:"authorize_url,omitempty"`
	ConnectionID string `json:"connection_id,omitempty"`
	State        string `json:"state,omitempty"`
}

// OAuthAppCredential represents saved OAuth app credentials for a provider.
type OAuthAppCredential struct {
	ProviderSlug string `json:"provider_slug"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

// OAuthAppCredentialListResponse is the response from listing OAuth app credentials.
type OAuthAppCredentialListResponse struct {
	Credentials []OAuthAppCredential `json:"credentials"`
}

// SaveOAuthAppCredentialsRequest are parameters for saving OAuth app credentials.
type SaveOAuthAppCredentialsRequest struct {
	ProviderSlug string `json:"provider_slug"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

// ListOAuthProviders lists all available OAuth providers in the registry.
func (s *OAuthConnectService) ListOAuthProviders(ctx context.Context) (*OAuthProviderListResponse, error) {
	var result OAuthProviderListResponse
	err := s.client.doJSON(ctx, "GET", "/v1/oauth/providers", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListOAuthConnections lists all OAuth connections for an agent.
func (s *OAuthConnectService) ListOAuthConnections(ctx context.Context, agentID string) (*OAuthConnectionListResponse, error) {
	var result OAuthConnectionListResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/oauth/connections", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConnectOAuth initiates an OAuth connection for an agent.
func (s *OAuthConnectService) ConnectOAuth(ctx context.Context, agentID string, req ConnectOAuthRequest) (*ConnectOAuthResponse, error) {
	var result ConnectOAuthResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/oauth/connect", agentID), req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DisconnectOAuth disconnects (revokes) an OAuth connection for an agent.
func (s *OAuthConnectService) DisconnectOAuth(ctx context.Context, agentID, bindingID string) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/oauth/disconnect/%s", agentID, bindingID), nil, nil)
}

// SaveOAuthAppCredentials saves custom OAuth app credentials for a provider.
func (s *OAuthConnectService) SaveOAuthAppCredentials(ctx context.Context, agentID string, req SaveOAuthAppCredentialsRequest) (*OAuthAppCredential, error) {
	var result OAuthAppCredential
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/oauth/app-credentials", agentID), req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListOAuthAppCredentials lists saved OAuth app credentials for an agent.
func (s *OAuthConnectService) ListOAuthAppCredentials(ctx context.Context, agentID string) (*OAuthAppCredentialListResponse, error) {
	var result OAuthAppCredentialListResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/oauth/app-credentials", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteOAuthAppCredentials deletes saved OAuth app credentials for a provider.
func (s *OAuthConnectService) DeleteOAuthAppCredentials(ctx context.Context, agentID, providerSlug string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/oauth/app-credentials/%s", agentID, providerSlug), nil, nil)
}
