package oneclaw

import (
	"context"
	"fmt"
)

// Create creates a new agent.
func (s *AgentsService) Create(ctx context.Context, req CreateAgentParams) (*AgentCreated, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := createAgentParamsToAPI(req)
	resp, _, err := s.client.api.AgentsAPI.CreateAgent(authCtx).
		CreateAgentRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return agentCreatedFromAPI(resp), nil
}

// Get retrieves an agent by ID.
func (s *AgentsService) Get(ctx context.Context, agentID string) (*Agent, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.GetAgent(authCtx, agentID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	a := agentFromAPI(resp)
	return &a, nil
}

// List lists agents.
func (s *AgentsService) List(ctx context.Context) (*AgentList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.ListAgents(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return agentListFromAPI(resp), nil
}

// Update updates an agent.
func (s *AgentsService) Update(ctx context.Context, agentID string, req UpdateAgentParams) (*Agent, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := updateAgentParamsToAPI(req)
	resp, _, err := s.client.api.AgentsAPI.UpdateAgent(authCtx, agentID).
		UpdateAgentRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	a := agentFromAPI(resp)
	return &a, nil
}

// Delete deletes an agent.
func (s *AgentsService) Delete(ctx context.Context, agentID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.AgentsAPI.DeleteAgent(authCtx, agentID).Execute()
	return wrapAPIError(err)
}

// --- Delegations ---

// CreateDelegationParams contains options for creating a delegation.
type CreateDelegationParams struct {
	DelegateID          string                 `json:"delegate_id"`
	AllowedTools        []string               `json:"allowed_tools,omitempty"`
	BlockedTools        []string               `json:"blocked_tools,omitempty"`
	MaxDailyDelegations *int                   `json:"max_daily_delegations,omitempty"`
	MaxDepth            *int                   `json:"max_depth,omitempty"`
	Guardrails          map[string]interface{} `json:"guardrails,omitempty"`
	DelegationMode      string                 `json:"delegation_mode,omitempty"`
	ExpiresAt           string                 `json:"expires_at,omitempty"`
}

// UpdateDelegationParams contains options for updating a delegation.
type UpdateDelegationParams struct {
	AllowedTools        []string               `json:"allowed_tools,omitempty"`
	BlockedTools        []string               `json:"blocked_tools,omitempty"`
	MaxDailyDelegations *int                   `json:"max_daily_delegations,omitempty"`
	MaxDepth            *int                   `json:"max_depth,omitempty"`
	Guardrails          map[string]interface{} `json:"guardrails,omitempty"`
	DelegationMode      string                 `json:"delegation_mode,omitempty"`
	IsActive            *bool                  `json:"is_active,omitempty"`
	ExpiresAt           string                 `json:"expires_at,omitempty"`
}

// DelegationResponse represents a delegation record.
type DelegationResponse struct {
	ID                  string                 `json:"id"`
	OrgID               string                 `json:"org_id"`
	DelegatorID         string                 `json:"delegator_id"`
	DelegateID          string                 `json:"delegate_id"`
	DelegatorName       *string                `json:"delegator_name,omitempty"`
	DelegateName        *string                `json:"delegate_name,omitempty"`
	AllowedTools        []string               `json:"allowed_tools"`
	BlockedTools        []string               `json:"blocked_tools"`
	MaxDailyDelegations *int                   `json:"max_daily_delegations,omitempty"`
	MaxDepth            int                    `json:"max_depth"`
	Guardrails          map[string]interface{} `json:"guardrails"`
	DelegationMode      string                 `json:"delegation_mode"`
	IsActive            bool                   `json:"is_active"`
	CreatedBy           string                 `json:"created_by"`
	ExpiresAt           *string                `json:"expires_at,omitempty"`
	CreatedAt           string                 `json:"created_at"`
	UpdatedAt           string                 `json:"updated_at"`
	DelegationsToday    *int                   `json:"delegations_today,omitempty"`
}

// DelegationListResponse wraps a list of delegations.
type DelegationListResponse struct {
	Delegations []DelegationResponse `json:"delegations"`
}

// CreateDelegation creates a delegation from this agent to another (human-only).
func (s *AgentsService) CreateDelegation(ctx context.Context, agentID string, params CreateDelegationParams) (*DelegationResponse, error) {
	var resp DelegationResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/delegations", agentID), params, &resp)
	return &resp, err
}

// ListDelegations lists all delegations for an agent.
func (s *AgentsService) ListDelegations(ctx context.Context, agentID string) (*DelegationListResponse, error) {
	var resp DelegationListResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/delegations", agentID), nil, &resp)
	return &resp, err
}

// GetDelegation retrieves a specific delegation by ID.
func (s *AgentsService) GetDelegation(ctx context.Context, agentID, delegationID string) (*DelegationResponse, error) {
	var resp DelegationResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/delegations/%s", agentID, delegationID), nil, &resp)
	return &resp, err
}

// UpdateDelegation updates a delegation (human-only).
func (s *AgentsService) UpdateDelegation(ctx context.Context, agentID, delegationID string, params UpdateDelegationParams) (*DelegationResponse, error) {
	var resp DelegationResponse
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/delegations/%s", agentID, delegationID), params, &resp)
	return &resp, err
}

// RevokeDelegation revokes (deletes) a delegation (human-only).
func (s *AgentsService) RevokeDelegation(ctx context.Context, agentID, delegationID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/delegations/%s", agentID, delegationID), nil, nil)
}

// GetEffectiveDelegations returns effective delegations for runtime tool discovery.
func (s *AgentsService) GetEffectiveDelegations(ctx context.Context, agentID string) (*DelegationListResponse, error) {
	var resp DelegationListResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/delegations/effective", agentID), nil, &resp)
	return &resp, err
}

// ImportSmartAccount imports an existing Safe smart account for an agent.
func (s *AgentsService) ImportSmartAccount(ctx context.Context, agentID string, req ImportSmartAccountRequest) (map[string]interface{}, error) {
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/smart-accounts/import", agentID), req, &resp)
	return resp, err
}
