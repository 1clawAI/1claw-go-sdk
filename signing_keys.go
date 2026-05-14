package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// Create provisions a signing key for the given agent and chain.
func (s *SigningKeysService) Create(ctx context.Context, agentID string, params CreateSigningKeyParams) (*SigningKey, error) {
	var result SigningKey
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/signing-keys", agentID), params, &result)
	return &result, err
}

// List returns all active signing keys for the given agent.
func (s *SigningKeysService) List(ctx context.Context, agentID string) (*SigningKeyList, error) {
	var result SigningKeyList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/signing-keys", agentID), nil, &result)
	return &result, err
}

// Rotate generates a new keypair for the given chain, replacing the current key.
func (s *SigningKeysService) Rotate(ctx context.Context, agentID, chain string) (*SigningKey, error) {
	var result SigningKey
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/signing-keys/%s/rotate", agentID, url.PathEscape(chain)), nil, &result)
	return &result, err
}

// Deactivate marks the signing key for the given chain as inactive.
func (s *SigningKeysService) Deactivate(ctx context.Context, agentID, chain string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/signing-keys/%s", agentID, url.PathEscape(chain)), nil, nil)
}

// Export retrieves the private key for the given chain's signing key.
// Requires re-authentication via the user's account password.
func (s *SigningKeysService) Export(ctx context.Context, agentID, chain, password string) (*SigningKeyExport, error) {
	var result SigningKeyExport
	headers := map[string]string{"X-Auth-Confirm": password}
	err := s.client.doJSONWithHeaders(ctx, "POST", fmt.Sprintf("/v1/agents/%s/signing-keys/%s/export", agentID, url.PathEscape(chain)), nil, &result, headers)
	return &result, err
}
