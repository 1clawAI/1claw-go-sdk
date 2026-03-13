package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// CreateShare creates a share link for a secret.
func (s *SharingService) CreateShare(ctx context.Context, secretID string, req openapi.CreateShareRequest) (*openapi.ShareResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SharingAPI.CreateShare(authCtx, secretID).
		CreateShareRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// ListOutboundShares lists shares you have sent.
func (s *SharingService) ListOutboundShares(ctx context.Context) (*openapi.ShareListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SharingAPI.ListOutboundShares(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// ListInboundShares lists shares sent to you.
func (s *SharingService) ListInboundShares(ctx context.Context) (*openapi.ShareListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.SharingAPI.ListInboundShares(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// RevokeShare revokes a share link.
func (s *SharingService) RevokeShare(ctx context.Context, shareID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.SharingAPI.RevokeShare(authCtx, shareID).Execute()
	return wrapAPIError(err)
}
