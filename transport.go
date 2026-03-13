package oneclaw

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"time"
)

// retryTransport wraps an http.RoundTripper and retries on 5xx and temporary errors.
type retryTransport struct {
	inner     http.RoundTripper
	maxRetries int
}

func (r *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for i := 0; i <= r.maxRetries; i++ {
		resp, err = r.inner.RoundTrip(req)
		if err != nil {
			if isRetryableErr(err) && i < r.maxRetries {
				time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
				continue
			}
			return resp, err
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 && i < r.maxRetries {
			resp.Body.Close()
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
			continue
		}
		return resp, nil
	}
	return resp, err
}

func isRetryableErr(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "connection refused") ||
		strings.Contains(s, "connection reset") ||
		strings.Contains(s, "timeout") ||
		strings.Contains(s, "temporary failure")
}

// idempotencyTransport adds Idempotency-Key to POST requests to transaction endpoints.
type idempotencyTransport struct {
	inner http.RoundTripper
}

func (i *idempotencyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodPost && strings.Contains(req.URL.Path, "/transactions") {
		if req.Header.Get("Idempotency-Key") == "" {
			key := generateIdempotencyKey()
			req = req.Clone(req.Context())
			req.Header.Set("Idempotency-Key", key)
		}
	}
	return i.inner.RoundTrip(req)
}

func generateIdempotencyKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
