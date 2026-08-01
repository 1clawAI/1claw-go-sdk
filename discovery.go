package oneclaw

import (
	"context"
	"fmt"
)

// DiscoveryService handles agent discovery operations.
type DiscoveryService struct {
	client *Client
}

// Publish publishes an agent to the discovery directory.
func (s *DiscoveryService) Publish(ctx context.Context, agentID string, params PublishParams) (*DirectoryListing, error) {
	var result DirectoryListing
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/discovery", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Unpublish removes an agent from the discovery directory.
func (s *DiscoveryService) Unpublish(ctx context.Context, agentID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/discovery", agentID), nil, nil)
}

// GetListing retrieves the directory listing for an agent.
func (s *DiscoveryService) GetListing(ctx context.Context, agentID string) (*DirectoryListing, error) {
	var result DirectoryListing
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/discovery", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateListing updates an agent's discovery listing.
func (s *DiscoveryService) UpdateListing(ctx context.Context, agentID string, params UpdateListingParams) (*DirectoryListing, error) {
	var result DirectoryListing
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/discovery", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches the agent discovery directory.
func (s *DiscoveryService) Search(ctx context.Context, query string) (*DirectorySearchResult, error) {
	var result DirectorySearchResult
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/directory?q=%s", query), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
