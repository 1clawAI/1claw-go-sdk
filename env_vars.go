package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// EnvVarsService handles environment variable operations within vaults.
type EnvVarsService struct {
	client *Client
}

// List lists environment variables in a vault, optionally filtered by environment.
func (s *EnvVarsService) List(ctx context.Context, vaultID string, environment string) (*EnvVarListResponse, error) {
	path := fmt.Sprintf("/v1/vaults/%s/env-vars", vaultID)
	if environment != "" {
		path += "?environment=" + url.QueryEscape(environment)
	}
	var result EnvVarListResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an environment variable in a vault.
func (s *EnvVarsService) Create(ctx context.Context, vaultID string, req CreateEnvVarRequest) (*EnvVarResponse, error) {
	var result EnvVarResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/vaults/%s/env-vars", vaultID), req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a single environment variable by key.
func (s *EnvVarsService) Get(ctx context.Context, vaultID, key, environment, gitBranch string) (*EnvVarResponse, error) {
	path := fmt.Sprintf("/v1/vaults/%s/env-vars/%s", vaultID, url.PathEscape(key))
	q := url.Values{}
	if environment != "" {
		q.Set("environment", environment)
	}
	if gitBranch != "" {
		q.Set("git_branch", gitBranch)
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result EnvVarResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an environment variable.
func (s *EnvVarsService) Update(ctx context.Context, vaultID, key string, req UpdateEnvVarRequest, environment, gitBranch string) (*EnvVarResponse, error) {
	path := fmt.Sprintf("/v1/vaults/%s/env-vars/%s", vaultID, url.PathEscape(key))
	q := url.Values{}
	if environment != "" {
		q.Set("environment", environment)
	}
	if gitBranch != "" {
		q.Set("git_branch", gitBranch)
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result EnvVarResponse
	err := s.client.doJSON(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an environment variable.
func (s *EnvVarsService) Delete(ctx context.Context, vaultID, key, environment, gitBranch string) error {
	path := fmt.Sprintf("/v1/vaults/%s/env-vars/%s", vaultID, url.PathEscape(key))
	q := url.Values{}
	if environment != "" {
		q.Set("environment", environment)
	}
	if gitBranch != "" {
		q.Set("git_branch", gitBranch)
	}
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	return s.client.doJSON(ctx, "DELETE", path, nil, nil)
}

// Resolve returns a merged key-value map of environment variables for a given
// environment, with source attribution (shared, vault, or branch override).
func (s *EnvVarsService) Resolve(ctx context.Context, vaultID, environment, gitBranch string) (*ResolveEnvVarsResponse, error) {
	q := url.Values{}
	q.Set("environment", environment)
	if gitBranch != "" {
		q.Set("git_branch", gitBranch)
	}
	path := fmt.Sprintf("/v1/vaults/%s/env-vars/resolve?%s", vaultID, q.Encode())
	var result ResolveEnvVarsResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListEnvironments lists available environments for a vault.
func (s *EnvVarsService) ListEnvironments(ctx context.Context, vaultID string) (*EnvironmentListResponse, error) {
	var result EnvironmentListResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/vaults/%s/environments", vaultID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateEnvironment creates a custom environment for a vault.
func (s *EnvVarsService) CreateEnvironment(ctx context.Context, vaultID string, req CreateEnvironmentRequest) (*EnvironmentResponse, error) {
	var result EnvironmentResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/vaults/%s/environments", vaultID), req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteEnvironment deletes a custom environment from a vault.
func (s *EnvVarsService) DeleteEnvironment(ctx context.Context, vaultID, slug string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/vaults/%s/environments/%s", vaultID, url.PathEscape(slug)), nil, nil)
}
