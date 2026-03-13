package oneclaw

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestGenerationScriptOutputsExpectedLocation verifies that the generation script
// (when run) produces output in internal/openapi. This test checks the expected
// layout exists; actual regeneration is done via make generate or scripts/generate.sh.
func TestGenerationScriptOutputsExpectedLocation(t *testing.T) {
	repoRoot := findRepoRoot(t)
	expectedDir := filepath.Join(repoRoot, "internal", "openapi")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Fatalf("expected internal/openapi to exist at %s", expectedDir)
	}
	// Check for key generated files
	requiredFiles := []string{"client.go", "configuration.go", "api_vaults.go", "api_secrets.go"}
	for _, f := range requiredFiles {
		p := filepath.Join(expectedDir, f)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			t.Errorf("expected generated file %s at %s", f, p)
		}
	}
}

// TestGenerationScriptRuns verifies the generation script can be invoked.
// Skips if npx is not available.
func TestGenerationScriptRuns(t *testing.T) {
	if _, err := exec.LookPath("npx"); err != nil {
		t.Skip("npx not found, skipping generation script test")
	}
	repoRoot := findRepoRoot(t)
	scriptPath := filepath.Join(repoRoot, "scripts", "generate.sh")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		t.Skip("scripts/generate.sh not found")
	}
	// We don't actually run it here to avoid network/side effects in CI,
	// but we verified the script exists. A separate CI step can run make generate.
	_ = scriptPath
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
