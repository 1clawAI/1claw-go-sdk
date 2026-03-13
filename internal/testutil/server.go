package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// NewTestServer creates an httptest.Server with the given handler and registers
// server.Close to run at test cleanup.
func NewTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return server
}
