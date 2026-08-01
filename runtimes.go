package oneclaw

import (
	"context"
	"fmt"
)

// RuntimesService handles runtime operations.
type RuntimesService struct {
	client *Client
}

// Create creates a new runtime deployment.
func (s *RuntimesService) Create(ctx context.Context, params CreateRuntimeParams) (*Runtime, error) {
	var result Runtime
	err := s.client.doJSON(ctx, "POST", "/v1/runtimes", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists runtimes.
func (s *RuntimesService) List(ctx context.Context) (*RuntimeList, error) {
	var result RuntimeList
	err := s.client.doJSON(ctx, "GET", "/v1/runtimes", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a runtime by ID.
func (s *RuntimesService) Get(ctx context.Context, runtimeID string) (*Runtime, error) {
	var result Runtime
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/runtimes/%s", runtimeID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a runtime.
func (s *RuntimesService) Update(ctx context.Context, runtimeID string, params UpdateRuntimeParams) (*Runtime, error) {
	var result Runtime
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/runtimes/%s", runtimeID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a runtime.
func (s *RuntimesService) Delete(ctx context.Context, runtimeID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/runtimes/%s", runtimeID), nil, nil)
}

// Restart restarts a runtime.
func (s *RuntimesService) Restart(ctx context.Context, runtimeID string) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/runtimes/%s/restart", runtimeID), nil, nil)
}
