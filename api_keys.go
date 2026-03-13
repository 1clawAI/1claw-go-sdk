package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// Create creates a personal API key.
func (s *APIKeysService) Create(ctx context.Context, name string, scopes []string) (*APIKeyCreated, error) {
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
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return apiKeyCreatedFromAPI(resp), nil
}

// List lists personal API keys.
func (s *APIKeysService) List(ctx context.Context) (*APIKeyList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.APIKeysAPI.ListApiKeys(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return apiKeyListFromAPI(resp), nil
}

// Revoke revokes an API key.
func (s *APIKeysService) Revoke(ctx context.Context, keyID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.APIKeysAPI.RevokeApiKey(authCtx, keyID).Execute()
	return wrapAPIError(err)
}
