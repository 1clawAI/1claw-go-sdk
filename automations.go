package oneclaw

import (
	"context"
	"fmt"
)

// AutomationsService handles automation operations.
type AutomationsService struct {
	client *Client
}

// Create creates a new automation for an agent.
func (s *AutomationsService) Create(ctx context.Context, agentID string, params CreateAutomationParams) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/automations", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists automations for an agent.
func (s *AutomationsService) List(ctx context.Context, agentID string) (*AutomationList, error) {
	var result AutomationList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/automations", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific automation.
func (s *AutomationsService) Get(ctx context.Context, agentID, automationID string) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/automations/%s", agentID, automationID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an automation.
func (s *AutomationsService) Update(ctx context.Context, agentID, automationID string, params UpdateAutomationParams) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/automations/%s", agentID, automationID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an automation.
func (s *AutomationsService) Delete(ctx context.Context, agentID, automationID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/automations/%s", agentID, automationID), nil, nil)
}

// ListRuns lists recent runs for an automation.
func (s *AutomationsService) ListRuns(ctx context.Context, agentID, automationID string) (*AutomationRunList, error) {
	var result AutomationRunList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/automations/%s/runs", agentID, automationID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Trigger manually triggers an automation.
func (s *AutomationsService) Trigger(ctx context.Context, agentID, automationID string) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/automations/%s/trigger", agentID, automationID), nil, nil)
}
