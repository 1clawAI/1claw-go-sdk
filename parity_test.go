package oneclaw

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Table-driven tests for resource parity with TS SDK.
func TestResourceParity_Chains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chains" || r.Method != http.MethodGet {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"chains": []map[string]interface{}{
				{"id": "base", "name": "Base", "chain_id": 8453},
			},
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Chains.ListChains(context.Background())
	if err != nil {
		t.Fatalf("ListChains error = %v", err)
	}
	if resp == nil || len(resp.Chains) != 1 {
		t.Errorf("resp = %+v", resp)
	}
}

func TestResourceParity_Sharing(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/shares/outbound" {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"shares": []map[string]interface{}{},
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Sharing.ListOutboundShares(context.Background())
	if err != nil {
		t.Fatalf("ListOutboundShares error = %v", err)
	}
	if resp == nil {
		t.Error("resp is nil")
	}
}
