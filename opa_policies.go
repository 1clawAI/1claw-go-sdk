package oneclaw

import (
	"context"
	"fmt"
)

// OpaPoliciesService provides OPA policy operations.
type OpaPoliciesService struct {
	client *Client
}

// Create creates a new OPA policy.
func (s *OpaPoliciesService) Create(ctx context.Context, req CreateOpaPolicyRequest) (*OpaPolicyResponse, error) {
	var resp OpaPolicyResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/opa-policies", req, &resp)
	return &resp, err
}

// List returns all OPA policies for the org.
func (s *OpaPoliciesService) List(ctx context.Context) (*OpaPolicyListResponse, error) {
	var resp OpaPolicyListResponse
	err := s.client.doJSON(ctx, "GET", "/v1/org/opa-policies", nil, &resp)
	return &resp, err
}

// Get retrieves a single OPA policy by ID.
func (s *OpaPoliciesService) Get(ctx context.Context, policyID string) (*OpaPolicyResponse, error) {
	var resp OpaPolicyResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/org/opa-policies/%s", policyID), nil, &resp)
	return &resp, err
}

// Delete removes an OPA policy by ID.
func (s *OpaPoliciesService) Delete(ctx context.Context, policyID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/org/opa-policies/%s", policyID), nil, nil)
}

// Test evaluates an OPA policy against a test input.
func (s *OpaPoliciesService) Test(ctx context.Context, req OpaPolicyTestRequest) (*OpaPolicyTestResponse, error) {
	var resp OpaPolicyTestResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/opa-policies/test", req, &resp)
	return &resp, err
}
