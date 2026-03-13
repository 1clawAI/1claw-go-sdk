package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// CreateAPIKey creates a personal API key.
func (s *APIKeysService) CreateAPIKey(ctx context.Context, name string, scopes []string) (*openapi.ApiKeyCreatedResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := openapi.CreateApiKeyRequest{Name: name}
	if len(scopes) > 0 {
		req.Scopes = scopes
	}
	resp, _, err := s.client.api.APIKeysAPI.CreateApiKey(authCtx).
		CreateApiKeyRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// ListAPIKeys lists personal API keys.
func (s *APIKeysService) ListAPIKeys(ctx context.Context) (*openapi.ApiKeyListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.APIKeysAPI.ListApiKeys(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// RevokeAPIKey revokes an API key.
func (s *APIKeysService) RevokeAPIKey(ctx context.Context, keyID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.APIKeysAPI.RevokeApiKey(authCtx, keyID).Execute()
	return wrapAPIError(err)
}
