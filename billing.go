package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// GetSubscription returns subscription, usage, and credit summary.
func (s *BillingService) GetSubscription(ctx context.Context) (*openapi.SubscriptionResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.BillingAPI.BillingSubscription(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// GetCreditBalance returns credit balance.
func (s *BillingService) GetCreditBalance(ctx context.Context) (*openapi.CreditBalanceResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.BillingAPI.BillingCreditBalance(authCtx).Execute()
	return resp, wrapAPIError(err)
}
