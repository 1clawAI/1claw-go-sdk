package oneclaw

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
	"github.com/1clawAI/1claw-go-sdk/internal/testutil"
)

func TestAuth_ApiKeyToken_Success(t *testing.T) {
	server := testutil.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/auth/api-key-token" {
			t.Errorf("path = %q, want /v1/auth/api-key-token", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		var req struct {
			APIKey string `json:"api_key"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			return
		}
		if req.APIKey != "ocv_test_key" {
			t.Errorf("api_key = %q, want ocv_test_key", req.APIKey)
		}
		body := testutil.LoadJSON(t, "testdata/auth/api_key_token_success.json")
		testutil.WriteJSON(t, w, http.StatusOK, body)
	}))

	client, err := New(
		WithBaseURL(server.URL),
		WithAPIKey("ocv_test_key"),
	)
	if err != nil {
		t.Fatalf("New error = %v", err)
	}

	ctx := context.Background()
	authCtx, err := client.authContext(ctx)
	if err != nil {
		t.Fatalf("authContext error = %v", err)
	}
	if authCtx.Value(openapi.ContextAccessToken) != "eyJ.test.jwt" {
		t.Error("token not set after API key exchange")
	}
}

func TestAuth_AgentToken_Success(t *testing.T) {
	server := testutil.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/auth/agent-token" {
			t.Errorf("path = %q, want /v1/auth/agent-token", r.URL.Path)
		}
		var req struct {
			AgentID string `json:"agent_id"`
			APIKey  string `json:"api_key"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			return
		}
		if req.AgentID != "agent-123" || req.APIKey != "ocv_agent_key" {
			t.Errorf("agent_id=%q api_key=%q", req.AgentID, req.APIKey)
		}
		testutil.WriteJSON(t, w, http.StatusOK, map[string]interface{}{
			"access_token": "eyJ.agent.jwt",
			"token_type":   "Bearer",
		})
	}))

	client, err := New(
		WithBaseURL(server.URL),
		WithAPIKey("ocv_agent_key"),
		WithAgentID("agent-123"),
	)
	if err != nil {
		t.Fatalf("New error = %v", err)
	}

	ctx := context.Background()
	authCtx, err := client.authContext(ctx)
	if err != nil {
		t.Fatalf("authContext error = %v", err)
	}
	if authCtx.Value(openapi.ContextAccessToken) != "eyJ.agent.jwt" {
		t.Error("agent token not set")
	}
}

func TestAuth_ApiKeyToken_401(t *testing.T) {
	server := testutil.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.LoadJSON(t, "testdata/auth/api_key_token_401.json")
		testutil.WriteJSON(t, w, http.StatusUnauthorized, body)
	}))

	client, err := New(
		WithBaseURL(server.URL),
		WithAPIKey("invalid"),
	)
	if err != nil {
		t.Fatalf("New error = %v", err)
	}

	ctx := context.Background()
	_, err = client.authContext(ctx)
	if err == nil {
		t.Fatal("expected error for 401")
	}
	var authErr *AuthError
	if !errors.As(err, &authErr) {
		t.Errorf("error type = %T, want *AuthError", err)
	}
	if authErr.StatusCode != 401 {
		t.Errorf("StatusCode = %d, want 401", authErr.StatusCode)
	}
}
