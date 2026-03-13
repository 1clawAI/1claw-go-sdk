package oneclaw

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

func TestIdempotencyTransport_AddsKeyToTransactionPost(t *testing.T) {
	var gotKey string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.Header.Get("Idempotency-Key")
		if r.URL.Path != "/v1/agents/a1/transactions" || r.Method != http.MethodPost {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	_, _, err = client.api.TransactionsAPI.SubmitTransaction(ctx, "a1").
		SubmitTransactionRequest(openapi.SubmitTransactionRequest{
			To:    "0x000000000000000000000000000000000000dEaD",
			Value: "0",
			Chain: "base",
		}).
		Execute()
	if err != nil {
		t.Fatalf("SubmitTransaction error = %v", err)
	}
	if gotKey == "" || len(gotKey) != 32 {
		t.Errorf("Idempotency-Key = %q (want 32-char hex)", gotKey)
	}
}

func TestRetryTransport_RetriesOn5xx(t *testing.T) {
	// Retry happens inside the transport; ListVaults will succeed after retries
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	_, err = client.Vaults.ListVaults(ctx)
	if err != nil {
		t.Fatalf("ListVaults error = %v", err)
	}
	if attempts != 2 {
		t.Errorf("attempts = %d, want 2", attempts)
	}
}
