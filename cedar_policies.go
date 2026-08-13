package oneclaw

import (
	"context"
	"fmt"
)

// CedarPoliciesService provides Cedar policy operations.
type CedarPoliciesService struct {
	client *Client
}

// Create creates a new Cedar policy.
func (s *CedarPoliciesService) Create(ctx context.Context, req CreateCedarPolicyRequest) (*CedarPolicyResponse, error) {
	var resp CedarPolicyResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/cedar-policies", req, &resp)
	return &resp, err
}

// List returns all Cedar policies for the org.
func (s *CedarPoliciesService) List(ctx context.Context) (*CedarPolicyListResponse, error) {
	var resp CedarPolicyListResponse
	err := s.client.doJSON(ctx, "GET", "/v1/org/cedar-policies", nil, &resp)
	return &resp, err
}

// Get retrieves a single Cedar policy by ID.
func (s *CedarPoliciesService) Get(ctx context.Context, policyID string) (*CedarPolicyResponse, error) {
	var resp CedarPolicyResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/org/cedar-policies/%s", policyID), nil, &resp)
	return &resp, err
}

// Delete removes a Cedar policy by ID.
func (s *CedarPoliciesService) Delete(ctx context.Context, policyID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/org/cedar-policies/%s", policyID), nil, nil)
}

// Test evaluates a Cedar policy against a test request.
func (s *CedarPoliciesService) Test(ctx context.Context, req CedarPolicyTestRequest) (*CedarPolicyTestResponse, error) {
	var resp CedarPolicyTestResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/cedar-policies/test", req, &resp)
	return &resp, err
}
