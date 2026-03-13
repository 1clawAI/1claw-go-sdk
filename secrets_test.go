package oneclaw

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecretsService_GetSecret(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Path may be encoded (db%2Fcreds) or not (db/creds) depending on client
		if r.URL.Path != "/v1/vaults/v1/secrets/db%2Fcreds" && r.URL.Path != "/v1/vaults/v1/secrets/db/creds" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": "s1", "path": "db/creds", "value": "secret123", "type": "generic", "version": 1, "created_at": "2024-01-01T00:00:00Z",
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Secrets.GetSecret(context.Background(), "v1", "db/creds")
	if err != nil {
		t.Fatalf("GetSecret error = %v", err)
	}
	if resp == nil || resp.Value != "secret123" {
		t.Errorf("resp = %+v", resp)
	}
}

func TestSecretsService_PutSecret(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s", r.Method)
		}
		var req struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Value != "my-secret" {
			t.Errorf("value = %q", req.Value)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// PutSecret returns SecretMetadataResponse (no value field)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": "s2", "path": "key", "type": "api_key", "version": 1, "created_at": "2024-01-01T00:00:00Z",
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Secrets.PutSecret(context.Background(), "v1", "key", "my-secret", "api_key")
	if err != nil {
		t.Fatalf("PutSecret error = %v", err)
	}
	if resp == nil {
		t.Error("resp is nil")
	}
}
