package oneclaw

import (
	"context"
	"fmt"
)

// WebhooksService handles webhook registration and management.
type WebhooksService struct {
	client *Client
}

// Webhook represents a registered event webhook.
type Webhook struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Events    []string `json:"events"`
	IsActive  bool     `json:"is_active"`
	Secret    bool     `json:"secret_configured"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// WebhookList is the response from listing webhooks.
type WebhookList struct {
	Webhooks []Webhook `json:"webhooks"`
}

// CreateWebhookParams are parameters for registering a webhook.
type CreateWebhookParams struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
	Secret string   `json:"secret,omitempty"`
}

// UpdateWebhookParams are parameters for updating a webhook.
type UpdateWebhookParams struct {
	URL      *string  `json:"url,omitempty"`
	Events   []string `json:"events,omitempty"`
	IsActive *bool    `json:"is_active,omitempty"`
}

// Create registers a new webhook.
func (s *WebhooksService) Create(ctx context.Context, params CreateWebhookParams) (*Webhook, error) {
	var result Webhook
	err := s.client.doJSON(ctx, "POST", "/v1/webhooks", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists webhooks for the organization.
func (s *WebhooksService) List(ctx context.Context) (*WebhookList, error) {
	var result WebhookList
	err := s.client.doJSON(ctx, "GET", "/v1/webhooks", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a webhook by ID.
func (s *WebhooksService) Get(ctx context.Context, webhookID string) (*Webhook, error) {
	var result Webhook
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/webhooks/%s", webhookID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a webhook.
func (s *WebhooksService) Update(ctx context.Context, webhookID string, params UpdateWebhookParams) (*Webhook, error) {
	var result Webhook
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/webhooks/%s", webhookID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a webhook.
func (s *WebhooksService) Delete(ctx context.Context, webhookID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/webhooks/%s", webhookID), nil, nil)
}
