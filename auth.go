package oneclaw

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// refreshThreshold is how long before expiry we refresh the token.
const refreshThreshold = 60 * time.Second

// ensureToken exchanges API key for JWT if needed, or refreshes when close to expiry.
func (c *Client) ensureToken(ctx context.Context) error {
	if c.token != "" && time.Now().Add(refreshThreshold).Before(c.tokenExpiry) {
		return nil
	}
	if c.refreshToken != "" {
		if err := c.refreshAccessToken(ctx); err == nil {
			return nil
		}
		// Refresh failed, clear and fall through to re-exchange
		c.refreshToken = ""
	}
	if c.apiKey == "" {
		return fmt.Errorf("no token or API key configured")
	}

	// Use agent token if agent ID was set (via WithAgentID)
	if c.agentID != "" {
		return c.exchangeAgentToken(ctx)
	}
	return c.exchangeAPIKeyToken(ctx)
}

func (c *Client) exchangeAPIKeyToken(ctx context.Context) error {
	req := openapi.UserApiKeyTokenRequest{ApiKey: c.apiKey}
	resp, httpResp, err := c.api.AuthenticationAPI.ApiKeyToken(ctx).
		UserApiKeyTokenRequest(req).
		Execute()
	if err != nil {
		return c.wrapAuthError(err, httpResp)
	}
	c.token = resp.AccessToken
	if resp.ExpiresIn != nil {
		c.tokenExpiry = time.Now().Add(time.Duration(*resp.ExpiresIn) * time.Second)
	}
	if resp.RefreshToken != nil {
		c.refreshToken = *resp.RefreshToken
	}
	return nil
}

func (c *Client) refreshAccessToken(ctx context.Context) error {
	req := openapi.RefreshTokenRequest{RefreshToken: c.refreshToken}
	resp, httpResp, err := c.api.AuthenticationAPI.RefreshToken(ctx).
		RefreshTokenRequest(req).
		Execute()
	if err != nil {
		return c.wrapAuthError(err, httpResp)
	}
	c.token = resp.AccessToken
	if resp.ExpiresIn != nil {
		c.tokenExpiry = time.Now().Add(time.Duration(*resp.ExpiresIn) * time.Second)
	}
	if resp.RefreshToken != nil {
		c.refreshToken = *resp.RefreshToken
	}
	return nil
}

func (c *Client) exchangeAgentToken(ctx context.Context) error {
	req := openapi.AgentTokenRequest{
		AgentId: c.agentID,
		ApiKey:  c.apiKey,
	}
	resp, httpResp, err := c.api.AuthenticationAPI.AgentToken(ctx).
		AgentTokenRequest(req).
		Execute()
	if err != nil {
		return c.wrapAuthError(err, httpResp)
	}
	c.token = resp.AccessToken
	if resp.ExpiresIn != nil {
		c.tokenExpiry = time.Now().Add(time.Duration(*resp.ExpiresIn) * time.Second)
	}
	return nil
}

func (c *Client) wrapAuthError(err error, httpResp *http.Response) error {
	var body []byte
	if gerr, ok := err.(*openapi.GenericOpenAPIError); ok {
		body = gerr.Body()
	}
	if httpResp != nil {
		if len(body) == 0 {
			body, _ = io.ReadAll(httpResp.Body)
			httpResp.Body.Close()
		}
		var msg, detail string
		var pd struct {
			Title  string `json:"title"`
			Detail string `json:"detail"`
		}
		if json.Unmarshal(body, &pd) == nil {
			msg, detail = pd.Title, pd.Detail
		}
		if msg == "" {
			msg = httpResp.Status
		}
		return &AuthError{
			StatusCode: httpResp.StatusCode,
			Message:    msg,
			Detail:     detail,
			Body:       body,
		}
	}
	return fmt.Errorf("%w: %v", ErrUnauthorized, err)
}
