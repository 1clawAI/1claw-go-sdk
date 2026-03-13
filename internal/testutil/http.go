package testutil

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

// WriteJSON writes status and body as JSON to w, setting Content-Type to application/json.
func WriteJSON(t *testing.T, w http.ResponseWriter, status int, body any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		t.Errorf("WriteJSON: %v", err)
	}
}

// LoadJSON reads a JSON file from path (relative to repo root) and unmarshals into a map.
// Use for testdata fixtures.
func LoadJSON(t *testing.T, path string) map[string]interface{} {
	t.Helper()
	root := findRepoRoot(t)
	fullPath := filepath.Join(root, path)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("LoadJSON %q: %v", path, err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("LoadJSON %q: %v", path, err)
	}
	return out
}

func findRepoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find repo root (go.mod)")
		}
		dir = parent
	}
}
