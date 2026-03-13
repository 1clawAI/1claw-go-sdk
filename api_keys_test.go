package oneclaw

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIKeysService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/auth/api-keys" || r.Method != http.MethodGet {
			t.Errorf("path=%s method=%s", r.URL.Path, r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"keys": []map[string]interface{}{
				{"id": "k1", "name": "dev-key"},
			},
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL), WithToken("eyJ.test"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.APIKeys.List(context.Background())
	if err != nil {
		t.Fatalf("List error = %v", err)
	}
	if resp == nil || len(resp.Keys) != 1 {
		t.Errorf("resp = %+v", resp)
	}
}
