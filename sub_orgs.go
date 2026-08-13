package oneclaw

import (
	"context"
	"fmt"
)

// SubOrgsService provides sub-organization operations.
type SubOrgsService struct {
	client *Client
}

// Create creates a new sub-organization.
func (s *SubOrgsService) Create(ctx context.Context, req CreateSubOrgRequest) (*SubOrgResponse, error) {
	var resp SubOrgResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/sub-orgs", req, &resp)
	return &resp, err
}

// List returns all sub-organizations for the org.
func (s *SubOrgsService) List(ctx context.Context) (*SubOrgListResponse, error) {
	var resp SubOrgListResponse
	err := s.client.doJSON(ctx, "GET", "/v1/org/sub-orgs", nil, &resp)
	return &resp, err
}

// Get retrieves a single sub-organization by ID.
func (s *SubOrgsService) Get(ctx context.Context, subOrgID string) (*SubOrgResponse, error) {
	var resp SubOrgResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/org/sub-orgs/%s", subOrgID), nil, &resp)
	return &resp, err
}

// Archive archives (soft-deletes) a sub-organization.
func (s *SubOrgsService) Archive(ctx context.Context, subOrgID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/org/sub-orgs/%s", subOrgID), nil, nil)
}

// GrantPermission grants a permission to a sub-organization.
func (s *SubOrgsService) GrantPermission(ctx context.Context, subOrgID string, req SubOrgPermissionRequest) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/org/sub-orgs/%s/permissions", subOrgID), req, nil)
}

// RevokePermission revokes a permission from a sub-organization.
func (s *SubOrgsService) RevokePermission(ctx context.Context, subOrgID string, permissionID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/org/sub-orgs/%s/permissions/%s", subOrgID, permissionID), nil, nil)
}

// AddUser adds a user to a sub-organization.
func (s *SubOrgsService) AddUser(ctx context.Context, subOrgID string, req SubOrgAddUserRequest) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/org/sub-orgs/%s/users", subOrgID), req, nil)
}

// GenerateWallets generates treasury wallets for a sub-organization.
func (s *SubOrgsService) GenerateWallets(ctx context.Context, subOrgID string, chains []string) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if len(chains) > 0 {
		body["chains"] = chains
	}
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/org/sub-orgs/%s/wallets/generate", subOrgID), body, &resp)
	return resp, err
}
