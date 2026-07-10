package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// Create creates a new binding for the given agent.
func (s *BindingsService) Create(ctx context.Context, agentID string, params CreateBindingParams) (*Binding, error) {
	var result Binding
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/bindings", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns all bindings for the given agent.
func (s *BindingsService) List(ctx context.Context, agentID string) (*BindingList, error) {
	var result BindingList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/bindings", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single binding by ID.
func (s *BindingsService) Get(ctx context.Context, agentID, bindingID string) (*Binding, error) {
	var result Binding
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/bindings/%s", agentID, url.PathEscape(bindingID)), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update patches a binding's config, guardrails, active status, or credential.
func (s *BindingsService) Update(ctx context.Context, agentID, bindingID string, params UpdateBindingParams) (*Binding, error) {
	var result Binding
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/bindings/%s", agentID, url.PathEscape(bindingID)), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a binding.
func (s *BindingsService) Delete(ctx context.Context, agentID, bindingID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/bindings/%s", agentID, url.PathEscape(bindingID)), nil, nil)
}

// Test verifies connectivity for a binding.
func (s *BindingsService) Test(ctx context.Context, agentID, bindingID string, params *TestBindingParams) (*TestBindingResult, error) {
	var body interface{}
	if params != nil {
		body = params
	} else {
		body = struct{}{}
	}
	var result TestBindingResult
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/bindings/%s/test", agentID, url.PathEscape(bindingID)), body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Execute runs an intent through a binding.
func (s *BindingsService) Execute(ctx context.Context, agentID string, params ExecuteParams) (*ExecuteResult, error) {
	var result ExecuteResult
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/execute", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListExecutions returns execution events for the given agent.
// Accepts optional limit and offset for pagination.
func (s *BindingsService) ListExecutions(ctx context.Context, agentID string, limit, offset *int) (*ExecutionEventList, error) {
	path := fmt.Sprintf("/v1/agents/%s/executions", agentID)
	qs := url.Values{}
	if limit != nil {
		qs.Set("limit", fmt.Sprintf("%d", *limit))
	}
	if offset != nil {
		qs.Set("offset", fmt.Sprintf("%d", *offset))
	}
	if q := qs.Encode(); q != "" {
		path += "?" + q
	}
	var result ExecutionEventList
	err := s.client.doJSON(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
