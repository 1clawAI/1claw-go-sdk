// Package openapi provides the generated 1Claw API client.
// This file contains tests that verify the generation pipeline and module layout.

package openapi

import (
	"testing"
)

// TestGeneratedClientCompiles verifies the SDK module compiles and can instantiate
// the internal generated client.
func TestGeneratedClientCompiles(t *testing.T) {
	cfg := NewConfiguration()
	if cfg == nil {
		t.Fatal("NewConfiguration returned nil")
	}
	client := NewAPIClient(cfg)
	if client == nil {
		t.Fatal("NewAPIClient returned nil")
	}
	// Ensure resource services are wired
	if client.VaultsAPI == nil {
		t.Error("VaultsAPI is nil")
	}
	if client.SecretsAPI == nil {
		t.Error("SecretsAPI is nil")
	}
	if client.AuthenticationAPI == nil {
		t.Error("AuthenticationAPI is nil")
	}
}
