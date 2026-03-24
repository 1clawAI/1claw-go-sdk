package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// APIKeyToken exchanges a user API key for a JWT. Use when you need to obtain a token manually.
func (s *AuthService) APIKeyToken(ctx context.Context, apiKey string) (*Token, error) {
	resp, _, err := s.client.api.AuthenticationAPI.ApiKeyToken(ctx).
		UserApiKeyTokenRequest(openapi.UserApiKeyTokenRequest{ApiKey: apiKey}).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return tokenFromAPI(resp), nil
}

// AgentToken exchanges agent credentials for a JWT.
func (s *AuthService) AgentToken(ctx context.Context, agentID, apiKey string) (*Token, error) {
	req := openapi.NewAgentTokenRequest(apiKey)
	if agentID != "" {
		req.SetAgentId(agentID)
	}
	resp, _, err := s.client.api.AuthenticationAPI.AgentToken(ctx).
		AgentTokenRequest(*req).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return tokenFromAPI(resp), nil
}

// Login authenticates with email and password.
func (s *AuthService) Login(ctx context.Context, email, password string) (*LoginResult, error) {
	resp, _, err := s.client.api.AuthenticationAPI.Login(ctx).
		LoginRequest(openapi.LoginRequest{Email: email, Password: password}).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return loginResultFromAPI(resp), nil
}

// Me returns the current user profile. Requires an authenticated context.
func (s *AuthService) Me(ctx context.Context) (*UserProfile, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AuthenticationAPI.GetMe(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return userProfileFromAPI(resp), nil
}
