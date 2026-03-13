package oneclaw

import (
	"context"
)

// List lists enabled chains.
func (s *ChainsService) List(ctx context.Context) (*ChainList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.ChainsAPI.ListChains(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return chainListFromAPI(resp), nil
}

// Get gets a chain by name or ID.
func (s *ChainsService) Get(ctx context.Context, identifier string) (*Chain, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.ChainsAPI.GetChain(authCtx, identifier).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	c := chainFromAPI(resp)
	return &c, nil
}
