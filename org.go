package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// ListOrgMembers lists organization members.
func (s *OrgService) ListOrgMembers(ctx context.Context) (*openapi.OrgMemberListResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.OrganizationAPI.ListOrgMembers(authCtx).Execute()
	return resp, wrapAPIError(err)
}

// InviteMember invites a member by email.
func (s *OrgService) InviteMember(ctx context.Context, req openapi.InviteMemberRequest) (*openapi.InviteMemberResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.OrganizationAPI.InviteMember(authCtx).
		InviteMemberRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// UpdateMemberRole updates a member's role.
func (s *OrgService) UpdateMemberRole(ctx context.Context, userID string, req openapi.UpdateMemberRoleRequest) (*openapi.OrgMemberResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.OrganizationAPI.UpdateMemberRole(authCtx, userID).
		UpdateMemberRoleRequest(req).
		Execute()
	return resp, wrapAPIError(err)
}

// RemoveMember removes a member from the organization.
func (s *OrgService) RemoveMember(ctx context.Context, userID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.OrganizationAPI.RemoveMember(authCtx, userID).Execute()
	return wrapAPIError(err)
}
