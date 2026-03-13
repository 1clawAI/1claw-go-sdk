package oneclaw

import (
	"context"
)

// Create creates an access policy on a vault.
func (s *AccessService) Create(ctx context.Context, vaultID string, req CreatePolicyParams) (*Policy, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := createPolicyParamsToAPI(req)
	resp, _, err := s.client.api.PoliciesAPI.CreatePolicy(authCtx, vaultID).
		CreatePolicyRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	p := policyFromAPI(resp)
	return &p, nil
}

// List lists policies on a vault.
func (s *AccessService) List(ctx context.Context, vaultID string) (*PolicyList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.PoliciesAPI.ListPolicies(authCtx, vaultID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return policyListFromAPI(resp), nil
}

// Update updates a policy.
func (s *AccessService) Update(ctx context.Context, vaultID, policyID string, req UpdatePolicyParams) (*Policy, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := updatePolicyParamsToAPI(req)
	resp, _, err := s.client.api.PoliciesAPI.UpdatePolicy(authCtx, vaultID, policyID).
		UpdatePolicyRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	p := policyFromAPI(resp)
	return &p, nil
}

// Delete deletes a policy.
func (s *AccessService) Delete(ctx context.Context, vaultID, policyID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.PoliciesAPI.DeletePolicy(authCtx, vaultID, policyID).Execute()
	return wrapAPIError(err)
}
