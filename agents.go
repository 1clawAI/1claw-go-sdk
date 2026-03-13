package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// CreateAgent creates a new agent.
func (s *AgentsService) CreateAgent(ctx context.Context, req openapi.CreateAgentRequest) (*openapi.AgentCreatedResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.CreateAgent(authCtx).
		CreateAgentRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// GetAgent retrieves an agent by ID.
func (s *AgentsService) GetAgent(ctx context.Context, agentID string) (*openapi.AgentResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.GetAgent(authCtx, agentID).Execute()
	return resp, wrapAPIError(err)
}

// ListAgents lists agents.
func (s *AgentsService) ListAgents(ctx context.Context) (*openapi.AgentListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.ListAgents(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// UpdateAgent updates an agent.
func (s *AgentsService) UpdateAgent(ctx context.Context, agentID string, req openapi.UpdateAgentRequest) (*openapi.AgentResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AgentsAPI.UpdateAgent(authCtx, agentID).
		UpdateAgentRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// DeleteAgent deletes an agent.
func (s *AgentsService) DeleteAgent(ctx context.Context, agentID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.AgentsAPI.DeleteAgent(authCtx, agentID).Execute()
	return wrapAPIError(err)
}
