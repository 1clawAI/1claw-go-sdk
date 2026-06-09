// Package oneclaw provides the 1Claw Go SDK for secret management.
package oneclaw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// Client is the 1Claw SDK client.
type Client struct {
	api          *openapi.APIClient
	token        string
	refreshToken string
	tokenExpiry  time.Time
	apiKey       string
	agentID      string
	tokenMu      sync.Mutex

	// Resource clients
	Auth        *AuthService
	Vaults      *VaultsService
	Secrets     *SecretsService
	Agents      *AgentsService
	APIKeys     *APIKeysService
	Sharing     *SharingService
	Access      *AccessService
	Org         *OrgService
	Chains      *ChainsService
	Billing     *BillingService
	Audit       *AuditService
	X402        *X402Service
	Treasury    *TreasuryService
	SigningKeys *SigningKeysService
	Platform   *PlatformService
}

// Option configures the Client.
type Option func(*config)

type config struct {
	baseURL    string
	token      string
	apiKey     string
	agentID    string
	httpClient *http.Client
	userAgent  string
	debug      bool
}

// WithBaseURL sets the API base URL.
func WithBaseURL(url string) Option {
	return func(c *config) {
		c.baseURL = url
	}
}

// WithToken sets a pre-obtained JWT. Takes precedence over API key.
func WithToken(token string) Option {
	return func(c *config) {
		c.token = token
	}
}

// WithAPIKey sets the API key. It will be exchanged for a JWT on first use.
func WithAPIKey(apiKey string) Option {
	return func(c *config) {
		c.apiKey = apiKey
	}
}

// WithAgentID sets the agent ID for agent token flow. Use with WithAPIKey.
func WithAgentID(agentID string) Option {
	return func(c *config) {
		c.agentID = agentID
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *config) {
		c.httpClient = client
	}
}

// WithUserAgent sets the User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *config) {
		c.userAgent = ua
	}
}

// WithDebug enables debug logging of requests/responses.
func WithDebug(debug bool) Option {
	return func(c *config) {
		c.debug = debug
	}
}

// New creates a new 1Claw SDK client.
func New(opts ...Option) (*Client, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	openapiCfg := openapi.NewConfiguration()
	baseTransport := http.DefaultTransport
	if cfg.httpClient != nil && cfg.httpClient.Transport != nil {
		baseTransport = cfg.httpClient.Transport
	}
	httpClient := &http.Client{
		Transport: &idempotencyTransport{
			inner: &retryTransport{
				inner:      baseTransport,
				maxRetries: 2,
			},
		},
	}
	if cfg.httpClient != nil {
		httpClient.Timeout = cfg.httpClient.Timeout
	}
	openapiCfg.HTTPClient = httpClient
	if cfg.userAgent != "" {
		openapiCfg.UserAgent = cfg.userAgent
	}
	openapiCfg.Debug = cfg.debug

	if cfg.baseURL != "" {
		openapiCfg.Servers = openapi.ServerConfigurations{
			{URL: cfg.baseURL, Description: "Custom"},
		}
	}

	api := openapi.NewAPIClient(openapiCfg)
	client := &Client{
		api:     api,
		token:   cfg.token,
		apiKey:  cfg.apiKey,
		agentID: cfg.agentID,
	}
	if cfg.token != "" {
		client.tokenExpiry = time.Now().Add(24 * time.Hour) // assume long-lived when passed directly
	}
	client.Auth = &AuthService{client: client}
	client.Vaults = &VaultsService{client: client}
	client.Secrets = &SecretsService{client: client}
	client.Agents = &AgentsService{client: client}
	client.APIKeys = &APIKeysService{client: client}
	client.Sharing = &SharingService{client: client}
	client.Access = &AccessService{client: client}
	client.Org = &OrgService{client: client}
	client.Chains = &ChainsService{client: client}
	client.Billing = &BillingService{client: client}
	client.Audit = &AuditService{client: client}
	client.X402 = &X402Service{client: client}
	client.Treasury = &TreasuryService{client: client}
	client.SigningKeys = &SigningKeysService{client: client}
	client.Platform = &PlatformService{client: client}
	return client, nil
}

// authContext returns a context with the bearer token injected for API calls.
// Exchanges API key for JWT if needed. Callers must use this context when invoking API methods.
func (c *Client) authContext(ctx context.Context) (context.Context, error) {
	if err := c.ensureToken(ctx); err != nil {
		return nil, err
	}
	if c.token != "" {
		return context.WithValue(ctx, openapi.ContextAccessToken, c.token), nil
	}
	return ctx, nil
}

// doJSON performs authenticated JSON HTTP requests for endpoints not yet in the
// generated openapi client. body may be nil for requests without a body.
func (c *Client) doJSON(ctx context.Context, method, path string, body any, result any) error {
	return c.doJSONWithHeaders(ctx, method, path, body, result, nil)
}

// doJSONWithHeaders is like doJSON but allows setting additional request headers.
func (c *Client) doJSONWithHeaders(ctx context.Context, method, path string, body any, result any, headers map[string]string) error {
	if err := c.ensureToken(ctx); err != nil {
		return err
	}

	baseURL := "https://api.1claw.xyz"
	if len(c.api.GetConfig().Servers) > 0 {
		baseURL = c.api.GetConfig().Servers[0].URL
	}

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.api.GetConfig().HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, 50*1024*1024))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(data))
	}
	if result != nil && len(data) > 0 {
		return json.Unmarshal(data, result)
	}
	return nil
}
