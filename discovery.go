package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// DiscoveryService provides agent discovery and inter-agent communication.
type DiscoveryService struct {
	client *Client
}

// OrgDirectoryParams contains options for listing org agents.
type OrgDirectoryParams struct {
	Query    string
	Tags     string
	Page     int
	PageSize int
}

// OrgDirectoryAgent represents an agent in the org directory.
type OrgDirectoryAgent struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	PublicDescription       *string  `json:"public_description"`
	PublicTags              []string `json:"public_tags"`
	A2AURL                  *string  `json:"a2a_url"`
	MCPURL                  *string  `json:"mcp_url"`
	IntentsAPIEnabled       bool     `json:"intents_api_enabled"`
	ExecutionIntentsEnabled bool     `json:"execution_intents_enabled"`
	MemoryEnabled           bool     `json:"memory_enabled"`
	ShroudEnabled           bool     `json:"shroud_enabled"`
}

// OrgDirectoryResponse is the response from the org directory endpoint.
type OrgDirectoryResponse struct {
	Agents   []OrgDirectoryAgent `json:"agents"`
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

// DelegateTaskParams contains options for delegating a task to another agent.
type DelegateTaskParams struct {
	Message  string `json:"message"`
	Model    string `json:"model,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// OrgDirectory lists agents in the caller's organization for sub-agent discovery.
func (s *DiscoveryService) OrgDirectory(ctx context.Context, params *OrgDirectoryParams) (*OrgDirectoryResponse, error) {
	path := "/v1/agents/org-directory"
	if params != nil {
		q := url.Values{}
		if params.Query != "" {
			q.Set("q", params.Query)
		}
		if params.Tags != "" {
			q.Set("tags", params.Tags)
		}
		if params.Page > 0 {
			q.Set("page", fmt.Sprintf("%d", params.Page))
		}
		if params.PageSize > 0 {
			q.Set("page_size", fmt.Sprintf("%d", params.PageSize))
		}
		if qs := q.Encode(); qs != "" {
			path += "?" + qs
		}
	}
	var resp OrgDirectoryResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &resp)
	return &resp, err
}

// Search searches the public agent directory.
func (s *DiscoveryService) Search(ctx context.Context, query string, tags string) (*DirectorySearchResult, error) {
	path := "/v1/agents/directory"
	q := url.Values{}
	if query != "" {
		q.Set("q", query)
	}
	if tags != "" {
		q.Set("tags", tags)
	}
	if qs := q.Encode(); qs != "" {
		path += "?" + qs
	}
	var result DirectorySearchResult
	err := s.client.doJSON(ctx, "GET", path, nil, &result)
	return &result, err
}

// GetAgentCard retrieves an agent's public discovery card.
func (s *DiscoveryService) GetAgentCard(ctx context.Context, agentID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/card", agentID), nil, &result)
	return result, err
}

// UpdateDiscovery updates an agent's discovery settings (human-only).
func (s *DiscoveryService) UpdateDiscovery(ctx context.Context, agentID string, params UpdateListingParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/discovery", agentID), params, &result)
	return result, err
}

// DelegateTask sends a task to another agent via chat (inter-agent communication).
func (s *DiscoveryService) DelegateTask(ctx context.Context, agentID string, params *DelegateTaskParams) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"message": params.Message,
		"mode":    "llm",
	}
	if params.Model != "" {
		body["model"] = params.Model
	}
	if params.Provider != "" {
		body["provider"] = params.Provider
	}
	var resp map[string]interface{}
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/chat", agentID), body, &resp)
	return resp, err
}
