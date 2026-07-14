package oneclaw

import (
	"context"
	"fmt"
)

// Order orders a prepaid or gift card for the given agent via the x402
// payment flow. An Idempotency-Key header is auto-generated for each call.
func (s *CardsService) Order(ctx context.Context, agentID string, params OrderCardRequest) (*CardResponse, error) {
	var result CardResponse
	headers := map[string]string{
		"Idempotency-Key": generateIdempotencyKey(),
	}
	err := s.client.doJSONWithHeaders(ctx, "POST", fmt.Sprintf("/v1/agents/%s/cards/order", agentID), params, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns all cards visible to the caller (agents see only their own).
func (s *CardsService) List(ctx context.Context) (*CardListResponse, error) {
	var result CardListResponse
	err := s.client.doJSON(ctx, "GET", "/v1/cards", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a single card by ID (masked — last4 only).
func (s *CardsService) Get(ctx context.Context, cardID string) (*CardResponse, error) {
	var result CardResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/cards/%s", cardID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Reveal returns full card details (PAN, CVV). Humans must pass their account
// password (sent as X-Auth-Confirm). Agents may reveal only when a human has
// enabled a per-card reveal policy.
func (s *CardsService) Reveal(ctx context.Context, cardID, password string) (*CardRevealResponse, error) {
	var result CardRevealResponse
	headers := map[string]string{"X-Auth-Confirm": password}
	err := s.client.doJSONWithHeaders(ctx, "POST", fmt.Sprintf("/v1/cards/%s/reveal", cardID), nil, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update patches a card's reveal policy and/or void_after (human-only).
func (s *CardsService) Update(ctx context.Context, cardID string, params UpdateCardRequest) (*CardResponse, error) {
	var result CardResponse
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/cards/%s", cardID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Void locks a card at the 1Claw level (forward-looking only).
func (s *CardsService) Void(ctx context.Context, cardID string) (*CardResponse, error) {
	var result CardResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/cards/%s/void", cardID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Refresh proxies a Laso refresh for a reference-mode card's balance/status.
func (s *CardsService) Refresh(ctx context.Context, cardID string) (*CardResponse, error) {
	var result CardResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/cards/%s/refresh", cardID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Import manually imports an existing card (human-only, full storage mode).
func (s *CardsService) Import(ctx context.Context, params ImportCardRequest) (*CardResponse, error) {
	var result CardResponse
	err := s.client.doJSON(ctx, "POST", "/v1/cards/import", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchGiftCards searches available Laso gift-card brands/servers.
func (s *CardsService) SearchGiftCards(ctx context.Context, params SearchGiftCardsRequest) (*SearchGiftCardsResponse, error) {
	var result SearchGiftCardsResponse
	err := s.client.doJSON(ctx, "POST", "/v1/cards/gift-cards/search", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
