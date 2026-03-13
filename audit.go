package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// QueryAuditEvents queries audit events.
func (s *AuditService) QueryAuditEvents(ctx context.Context, resourceID string, limit int32) (*openapi.AuditEventsResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := s.client.api.AuditAPI.QueryAuditEvents(authCtx)
	if resourceID != "" {
		req = req.ResourceId(resourceID)
	}
	if limit > 0 {
		req = req.Limit(limit)
	}
	resp, _, err := req.Execute()
	return resp, wrapAPIError(err)
}
