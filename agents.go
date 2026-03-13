package oneclaw

import (
	"context"
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
