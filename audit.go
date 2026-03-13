package oneclaw

import (
	"context"
)

// Query queries audit events.
func (s *AuditService) Query(ctx context.Context, resourceID string, limit int32) (*AuditEvents, error) {
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
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return auditEventsFromAPI(resp), nil
}
