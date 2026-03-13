package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// PutSecret creates or updates a secret.
func (s *SecretsService) PutSecret(ctx context.Context, vaultID, path, value string, secretType string) (*SecretMetadata, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := openapi.PutSecretRequest{Value: value}
	if secretType != "" {
		req.Type = &secretType
	}
	resp, _, err := s.client.api.SecretsAPI.PutSecret(authCtx, vaultID, path).
		PutSecretRequest(req).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	sm := secretMetadataFromAPI(resp)
	return &sm, nil
}

// GetSecret retrieves a decrypted secret.
func (s *SecretsService) GetSecret(ctx context.Context, vaultID, path string) (*Secret, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SecretsAPI.GetSecret(authCtx, vaultID, path).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return secretFromAPI(resp), nil
}

// DeleteSecret deletes a secret.
func (s *SecretsService) DeleteSecret(ctx context.Context, vaultID, path string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.SecretsAPI.DeleteSecret(authCtx, vaultID, path).Execute()
	return wrapAPIError(err)
}

// ListSecrets lists secrets in a vault.
func (s *SecretsService) ListSecrets(ctx context.Context, vaultID string) (*SecretList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SecretsAPI.ListSecrets(authCtx, vaultID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return secretListFromAPI(resp), nil
}
