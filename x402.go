package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// PaymentRequirement represents a 402 payment requirement from the API.
type PaymentRequirement = openapi.PaymentRequirement

// X402Signer signs EIP-712 payment payloads for x402 protocol.
// Implement this interface to enable automatic payment when API returns 402.
type X402Signer interface {
	GetAddress(ctx context.Context) (string, error)
	SignPayment(ctx context.Context, accept *PaymentRequirement) (signedPayloadHex string, err error)
}

// GetPaymentRequirement is a placeholder for x402 payment flow.
// The actual 402 handling is typically done via error inspection when API returns 402.
func (s *X402Service) GetPaymentRequirement(ctx context.Context) (*openapi.PaymentRequirement, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	// Admin API has GetX402Config - for org context we'd use billing/payment endpoints
	_ = authCtx
	return nil, nil // Stub: full x402 flow when payment endpoints are available
}
