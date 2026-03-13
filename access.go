package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// CreatePolicy creates an access policy on a vault.
func (s *AccessService) CreatePolicy(ctx context.Context, vaultID string, req openapi.CreatePolicyRequest) (*openapi.PolicyResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.PoliciesAPI.CreatePolicy(authCtx, vaultID).
		CreatePolicyRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// ListPolicies lists policies on a vault.
func (s *AccessService) ListPolicies(ctx context.Context, vaultID string) (*openapi.PolicyListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.PoliciesAPI.ListPolicies(authCtx, vaultID).Execute()
	return resp, wrapAPIError(err)
}

// UpdatePolicy updates a policy.
func (s *AccessService) UpdatePolicy(ctx context.Context, vaultID, policyID string, req openapi.UpdatePolicyRequest) (*openapi.PolicyResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.PoliciesAPI.UpdatePolicy(authCtx, vaultID, policyID).
		UpdatePolicyRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// DeletePolicy deletes a policy.
func (s *AccessService) DeletePolicy(ctx context.Context, vaultID, policyID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.PoliciesAPI.DeletePolicy(authCtx, vaultID, policyID).Execute()
	return wrapAPIError(err)
}
