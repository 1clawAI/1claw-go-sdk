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

	resp, err := client.Chains.List(context.Background())
	if err != nil {
		t.Fatalf("ListChains error = %v", err)
	}
	if resp == nil || len(resp.Chains) != 1 {
		t.Errorf("resp = %+v", resp)
	}
}

func TestResourceParity_Treasury(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/treasury" || r.Method != http.MethodGet {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"treasuries": []map[string]interface{}{},
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Treasury.List(context.Background())
	if err != nil {
		t.Fatalf("Treasury.List error = %v", err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
}

func TestResourceParity_Bindings(t *testing.T) {
	agentID := "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

	t.Run("List", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expected := "/v1/agents/" + agentID + "/bindings"
			if r.URL.Path != expected || r.Method != http.MethodGet {
				t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"bindings": []map[string]interface{}{
					{"id": "b1", "agent_id": agentID, "binding_type": "http", "name": "my-api", "is_active": true},
				},
			})
		}))
		defer server.Close()

		client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := client.Bindings.List(context.Background(), agentID)
		if err != nil {
			t.Fatalf("Bindings.List error = %v", err)
		}
		if resp == nil || len(resp.Bindings) != 1 {
			t.Errorf("resp = %+v", resp)
		}
	})

	t.Run("Execute", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expected := "/v1/agents/" + agentID + "/execute"
			if r.URL.Path != expected || r.Method != http.MethodPost {
				t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"execution_id": "exec-1",
				"status":       "success",
			})
		}))
		defer server.Close()

		client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := client.Bindings.Execute(context.Background(), agentID, ExecuteParams{
			Binding:    "my-api",
			IntentType: "http",
			Params:     map[string]interface{}{"url": "https://example.com"},
		})
		if err != nil {
			t.Fatalf("Bindings.Execute error = %v", err)
		}
		if resp == nil || resp.ExecutionID != "exec-1" {
			t.Errorf("resp = %+v", resp)
		}
	})
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

	resp, err := client.Sharing.ListOutbound(context.Background())
	if err != nil {
		t.Fatalf("ListOutboundShares error = %v", err)
	}
	if resp == nil {
		t.Error("resp is nil")
	}
}
