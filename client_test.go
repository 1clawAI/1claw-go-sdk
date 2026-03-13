package oneclaw

import (
	"context"
	"net/http"
	"testing"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

func TestNew_DefaultConfig(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if client == nil {
		t.Fatal("New() returned nil client")
	}
	if client.api == nil {
		t.Error("client.api is nil")
	}
}

func TestNew_WithBaseURL(t *testing.T) {
	client, err := New(WithBaseURL("https://custom.example.com"))
	if err != nil {
		t.Fatalf("New(WithBaseURL) error = %v", err)
	}
	cfg := client.api.GetConfig()
	if len(cfg.Servers) == 0 {
		t.Fatal("no servers configured")
	}
	if cfg.Servers[0].URL != "https://custom.example.com" {
		t.Errorf("base URL = %q, want https://custom.example.com", cfg.Servers[0].URL)
	}
}

func TestNew_WithToken(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.test"
	client, err := New(WithToken(token))
	if err != nil {
		t.Fatalf("New(WithToken) error = %v", err)
	}
	ctx := context.Background()
	authCtx, err := client.authContext(ctx)
	if err != nil {
		t.Fatalf("authContext error = %v", err)
	}
	if authCtx == nil {
		t.Fatal("authContext returned nil")
	}
	// Token should be in context for the underlying client to use
	if authCtx.Value(openapi.ContextAccessToken) != token {
		t.Error("token not set in auth context")
	}
}

func TestNew_WithAPIKey(t *testing.T) {
	client, err := New(WithAPIKey("ocv_test_key_placeholder"))
	if err != nil {
		t.Fatalf("New(WithAPIKey) error = %v", err)
	}
	if client.apiKey != "ocv_test_key_placeholder" {
		t.Errorf("apiKey = %q, want ocv_test_key_placeholder", client.apiKey)
	}
}

func TestNew_WithHTTPClient(t *testing.T) {
	custom := &http.Client{Transport: http.DefaultTransport}
	client, err := New(WithHTTPClient(custom))
	if err != nil {
		t.Fatalf("New(WithHTTPClient) error = %v", err)
	}
	cfg := client.api.GetConfig()
	if cfg.HTTPClient == nil {
		t.Error("HTTPClient not set")
	}
	// Client wraps custom transport with retry/idempotency
	if cfg.HTTPClient.Transport == nil {
		t.Error("HTTPClient.Transport not set")
	}
}

func TestNew_WithUserAgent(t *testing.T) {
	ua := "my-app/1.0"
	client, err := New(WithUserAgent(ua))
	if err != nil {
		t.Fatalf("New(WithUserAgent) error = %v", err)
	}
	cfg := client.api.GetConfig()
	if cfg.UserAgent != ua {
		t.Errorf("UserAgent = %q, want %q", cfg.UserAgent, ua)
	}
}

func TestNew_WithDebug(t *testing.T) {
	client, err := New(WithDebug(true))
	if err != nil {
		t.Fatalf("New(WithDebug) error = %v", err)
	}
	cfg := client.api.GetConfig()
	if !cfg.Debug {
		t.Error("Debug not set")
	}
}
