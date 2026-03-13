package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// ListChains lists enabled chains.
func (s *ChainsService) ListChains(ctx context.Context) (*openapi.ChainListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.ChainsAPI.ListChains(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// GetChain gets a chain by name or ID.
func (s *ChainsService) GetChain(ctx context.Context, identifier string) (*openapi.ChainResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.ChainsAPI.GetChain(authCtx, identifier).Execute()
	return resp, wrapAPIError(err)
}
