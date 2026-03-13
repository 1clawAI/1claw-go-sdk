package oneclaw

import (
	"context"
)

// Subscription returns subscription, usage, and credit summary.
func (s *BillingService) Subscription(ctx context.Context) (*Subscription, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.BillingAPI.BillingSubscription(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return subscriptionFromAPI(resp), nil
}

// CreditBalance returns credit balance.
func (s *BillingService) CreditBalance(ctx context.Context) (*CreditBalance, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.BillingAPI.BillingCreditBalance(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return creditBalanceFromAPI(resp), nil
}
