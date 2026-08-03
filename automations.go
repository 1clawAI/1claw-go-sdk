package oneclaw

import (
	"context"
	"fmt"
)

// AutomationsService handles automation operations.
type AutomationsService struct {
	client *Client
}

// Create creates a new automation. agent_id and workflow_spec are required on params.
func (s *AutomationsService) Create(ctx context.Context, params CreateAutomationParams) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "POST", "/v1/automations", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists automations for the current organization.
func (s *AutomationsService) List(ctx context.Context) (*AutomationList, error) {
	var result AutomationList
	err := s.client.doJSON(ctx, "GET", "/v1/automations", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific automation.
func (s *AutomationsService) Get(ctx context.Context, automationID string) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/automations/%s", automationID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an automation.
func (s *AutomationsService) Update(ctx context.Context, automationID string, params UpdateAutomationParams) (*Automation, error) {
	var result Automation
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/automations/%s", automationID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an automation.
func (s *AutomationsService) Delete(ctx context.Context, automationID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/automations/%s", automationID), nil, nil)
}

// ListRuns lists recent runs for an automation.
func (s *AutomationsService) ListRuns(ctx context.Context, automationID string) (*AutomationRunList, error) {
	var result AutomationRunList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/automations/%s/runs", automationID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Trigger manually triggers an automation.
func (s *AutomationsService) Trigger(ctx context.Context, automationID string, input map[string]interface{}) (*AutomationRun, error) {
	var result AutomationRun
	body := map[string]interface{}{}
	if input != nil {
		body["input"] = input
	}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/automations/%s/trigger", automationID), body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
