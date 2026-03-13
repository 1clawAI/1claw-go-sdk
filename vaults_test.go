package oneclaw

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVaultsService_ListVaults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/vaults" || r.Method != http.MethodGet {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer eyJ.test" {
			t.Error("missing or invalid Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"vaults": []map[string]interface{}{
				{"id": "v1", "name": "my-vault", "created_at": "2024-01-01T00:00:00Z"},
			},
		})
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithToken("eyJ.test"),
	)
	if err != nil {
		t.Fatalf("New error = %v", err)
	}

	resp, err := client.Vaults.ListVaults(context.Background())
	if err != nil {
		t.Fatalf("ListVaults error = %v", err)
	}
	if resp == nil || len(resp.Vaults) != 1 {
		t.Errorf("resp = %+v", resp)
	}
}

func TestVaultsService_CreateVault(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/vaults" || r.Method != http.MethodPost {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		var req struct {
			Name string `json:"name"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Name != "test-vault" {
			t.Errorf("name = %q", req.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": "v-new", "name": "test-vault", "created_at": "2024-01-01T00:00:00Z",
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Vaults.CreateVault(context.Background(), "test-vault", "")
	if err != nil {
		t.Fatalf("CreateVault error = %v", err)
	}
	if resp == nil || resp.Id != "v-new" {
		t.Errorf("resp = %+v", resp)
	}
}
