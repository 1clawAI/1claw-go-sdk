package oneclaw

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

func TestAuth_ApiKeyToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "eyJ.test.jwt",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "eyJ.agent.jwt",
			"token_type":   "Bearer",
		})
	}))
	defer server.Close()

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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"title":  "Unauthorized",
			"detail": "Invalid API key",
		})
	}))
	defer server.Close()

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
