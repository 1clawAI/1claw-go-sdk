package oneclaw

import (
	"context"
)

// Create creates a share link for a secret.
func (s *SharingService) Create(ctx context.Context, secretID string, req CreateShareParams) (*Share, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := createShareParamsToAPI(req)
	resp, _, err := s.client.api.SharingAPI.CreateShare(authCtx, secretID).
		CreateShareRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	sh := shareFromAPI(resp)
	return &sh, nil
}

// ListOutbound lists shares you have sent.
func (s *SharingService) ListOutbound(ctx context.Context) (*ShareList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SharingAPI.ListOutboundShares(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return shareListFromAPI(resp), nil
}

// ListInbound lists shares sent to you.
func (s *SharingService) ListInbound(ctx context.Context) (*ShareList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SharingAPI.ListInboundShares(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return shareListFromAPI(resp), nil
}

// Revoke revokes a share link.
func (s *SharingService) Revoke(ctx context.Context, shareID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.SharingAPI.RevokeShare(authCtx, shareID).Execute()
	return wrapAPIError(err)
}
