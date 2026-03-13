package oneclaw

import (
	"context"
)

// ListMembers lists organization members.
func (s *OrgService) ListMembers(ctx context.Context) (*OrgMemberList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.OrganizationAPI.ListOrgMembers(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return orgMemberListFromAPI(resp), nil
}

// InviteMember invites a member by email.
func (s *OrgService) InviteMember(ctx context.Context, req InviteMemberParams) (*InviteMemberResult, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := inviteMemberParamsToAPI(req)
	resp, _, err := s.client.api.OrganizationAPI.InviteMember(authCtx).
		InviteMemberRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return inviteMemberResultFromAPI(resp), nil
}

// UpdateMemberRole updates a member's role.
func (s *OrgService) UpdateMemberRole(ctx context.Context, userID string, req UpdateMemberRoleParams) (*OrgMember, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	apiReq := updateMemberRoleParamsToAPI(req)
	resp, _, err := s.client.api.OrganizationAPI.UpdateMemberRole(authCtx, userID).
		UpdateMemberRoleRequest(apiReq).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	m := orgMemberFromAPI(resp)
	return &m, nil
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
