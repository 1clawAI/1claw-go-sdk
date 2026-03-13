package oneclaw

import (
	"context"
)

// X402Signer signs EIP-712 payment payloads for x402 protocol.
// Implement this interface to enable automatic payment when API returns 402.
type X402Signer interface {
	GetAddress(ctx context.Context) (string, error)
	SignPayment(ctx context.Context, accept *PaymentRequirement) (signedPayloadHex string, err error)
}

// PaymentRequirement returns the payment requirement for a 402 response.
// Stub: full x402 flow when payment endpoints are available.
func (s *X402Service) PaymentRequirement(ctx context.Context) (*PaymentRequirement, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	_ = authCtx
	return nil, nil
}
