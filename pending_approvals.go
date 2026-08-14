package oneclaw

import (
	"context"
	"fmt"
)

// PendingApprovalsService manages consensus-based approval workflows.
type PendingApprovalsService struct {
	client *Client
}

// Submit creates a pending approval for a consensus-gated action.
func (s *PendingApprovalsService) Submit(ctx context.Context, req SubmitPendingApprovalRequest) (*SubmitPendingApprovalResponse, error) {
	var resp SubmitPendingApprovalResponse
	err := s.client.doJSON(ctx, "POST", "/v1/pending-approvals", req, &resp)
	return &resp, err
}

// List returns pending approvals for the org.
func (s *PendingApprovalsService) List(ctx context.Context, params ListPendingApprovalsParams) (*PendingApprovalListResponse, error) {
	path := "/v1/pending-approvals"
	q := ""
	if params.Status != "" {
		q = fmt.Sprintf("status=%s", params.Status)
	}
	if params.AgentID != "" {
		if q != "" {
			q += "&"
		}
		q += fmt.Sprintf("agent_id=%s", params.AgentID)
	}
	if params.Limit > 0 {
		if q != "" {
			q += "&"
		}
		q += fmt.Sprintf("limit=%d", params.Limit)
	}
	if params.Offset > 0 {
		if q != "" {
			q += "&"
		}
		q += fmt.Sprintf("offset=%d", params.Offset)
	}
	if q != "" {
		path += "?" + q
	}
	var resp PendingApprovalListResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &resp)
	return &resp, err
}

// Get retrieves a pending approval with signatures.
func (s *PendingApprovalsService) Get(ctx context.Context, id string) (*PendingApprovalResponse, error) {
	var resp PendingApprovalResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/pending-approvals/%s", id), nil, &resp)
	return &resp, err
}

// Approve approves or rejects a pending approval.
func (s *PendingApprovalsService) Approve(ctx context.Context, id string, req ApprovePendingApprovalRequest) (*PendingApprovalResponse, error) {
	var resp PendingApprovalResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/pending-approvals/%s/approve", id), req, &resp)
	return &resp, err
}

// Execute marks an approved action as executed (human-only).
func (s *PendingApprovalsService) Execute(ctx context.Context, id string) (*ExecutePendingApprovalResponse, error) {
	var resp ExecutePendingApprovalResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/pending-approvals/%s/execute", id), nil, &resp)
	return &resp, err
}

// Cancel cancels a pending approval.
func (s *PendingApprovalsService) Cancel(ctx context.Context, id string) (*PendingApprovalResponse, error) {
	var resp PendingApprovalResponse
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/pending-approvals/%s/cancel", id), nil, &resp)
	return &resp, err
}
