package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// ApiKeyToken exchanges a user API key for a JWT. Use when you need to obtain a token manually.
func (s *AuthService) ApiKeyToken(ctx context.Context, apiKey string) (*openapi.TokenResponse, error) {
	resp, _, err := s.client.api.AuthenticationAPI.ApiKeyToken(ctx).
		UserApiKeyTokenRequest(openapi.UserApiKeyTokenRequest{ApiKey: apiKey}).
		Execute()
	return resp, wrapAPIError(err)
}

// AgentToken exchanges agent credentials for a JWT.
func (s *AuthService) AgentToken(ctx context.Context, agentID, apiKey string) (*openapi.TokenResponse, error) {
	resp, _, err := s.client.api.AuthenticationAPI.AgentToken(ctx).
		AgentTokenRequest(openapi.AgentTokenRequest{AgentId: agentID, ApiKey: apiKey}).
		Execute()
	return resp, wrapAPIError(err)
}

// Login authenticates with email and password.
func (s *AuthService) Login(ctx context.Context, email, password string) (*openapi.LoginResponse, error) {
	resp, _, err := s.client.api.AuthenticationAPI.Login(ctx).
		LoginRequest(openapi.LoginRequest{Email: email, Password: password}).
		Execute()
	return resp, wrapAPIError(err)
}

// GetMe returns the current user profile. Requires an authenticated context.
func (s *AuthService) GetMe(ctx context.Context) (*openapi.UserProfileResponse, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.AuthenticationAPI.GetMe(authCtx).Execute()
	return resp, wrapAPIError(err)
}
