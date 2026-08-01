package oneclaw

import (
	"context"
	"fmt"
)

// MemoryService handles agent memory operations.
type MemoryService struct {
	client *Client
}

// Store stores a memory entry with automatic embedding.
func (s *MemoryService) Store(ctx context.Context, agentID string, params StoreMemoryParams) (*MemoryEntry, error) {
	var result MemoryEntry
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/memory", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches memory entries by semantic similarity.
func (s *MemoryService) Search(ctx context.Context, agentID string, params SearchMemoryParams) (*MemoryEntryList, error) {
	var result MemoryEntryList
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/memory/search", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists memory entries for an agent.
func (s *MemoryService) List(ctx context.Context, agentID string) (*MemoryEntryList, error) {
	var result MemoryEntryList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/memory", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific memory entry.
func (s *MemoryService) Get(ctx context.Context, agentID, entryID string) (*MemoryEntry, error) {
	var result MemoryEntry
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/memory/%s", agentID, entryID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a memory entry.
func (s *MemoryService) Delete(ctx context.Context, agentID, entryID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/memory/%s", agentID, entryID), nil, nil)
}

// Clear removes all memory entries for an agent.
func (s *MemoryService) Clear(ctx context.Context, agentID string) error {
	return s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/memory/clear", agentID), nil, nil)
}
