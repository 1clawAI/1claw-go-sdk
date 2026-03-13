// Package oneclaw provides the 1Claw Go SDK for secret management.
package oneclaw

import "time"

// --- Vault types ---

// Vault represents a vault.
type Vault struct {
	ID               string
	Name             string
	Description       string
	CreatedBy         string
	CreatedByType     string
	CreatedAt         time.Time
	CMEKEnabled       bool
	CMEKFingerprint   string
}

// VaultList is the response from listing vaults.
type VaultList struct {
	Vaults []Vault
}

// --- Secret types ---

// Secret represents a secret with its decrypted value.
type Secret struct {
	ID            string
	Path          string
	Type          string
	Value         string
	Version       int32
	Metadata      map[string]interface{}
	CreatedBy     string
	CreatedAt     time.Time
	ExpiresAt     *time.Time
	CMEKEncrypted bool
}

// SecretMetadata represents secret metadata without the value.
type SecretMetadata struct {
	ID        string
	Path      string
	Type      string
	Version   int32
	Metadata  map[string]interface{}
	CreatedAt time.Time
	ExpiresAt *time.Time
}

// SecretList is the response from listing secrets.
type SecretList struct {
	Secrets []SecretMetadata
}

// --- Agent types ---

// Agent represents an agent.
type Agent struct {
	ID                   string
	Name                 string
	Description          string
	AuthMethod           string
	Scopes               []string
	IsActive             bool
	IntentsAPIEnabled    bool
	TxToAllowlist        []string
	TxMaxValueEth        string
	TxDailyLimitEth      string
	TxAllowedChains      []string
	TokenTtlSeconds      *int32
	VaultIDs             []string
	ClientCertFingerprint string
	OIDCIssuer           string
	OIDCClientID         string
	SSHPublicKey         string
	ECDHPublicKey        string
	ShroudEnabled        bool
	CreatedAt            time.Time
	ExpiresAt            *time.Time
	LastActiveAt         *time.Time
}

// AgentCreated is the response from creating an agent.
type AgentCreated struct {
	Agent  Agent
	APIKey string
}

// AgentList is the response from listing agents.
type AgentList struct {
	Agents []Agent
}

// CreateAgentParams are parameters for creating an agent.
type CreateAgentParams struct {
	Name                 string
	Description          string
	AuthMethod           string
	Scopes               []string
	ExpiresAt            *time.Time
	IntentsAPIEnabled    bool
	TxToAllowlist        []string
	TxMaxValueEth        string
	TxDailyLimitEth      string
	TxAllowedChains      []string
	TokenTtlSeconds      *int32
	VaultIDs             []string
	ClientCertFingerprint string
	OIDCIssuer           string
	OIDCClientID         string
	ShroudEnabled        bool
	// ShroudConfig can be a struct that marshals to the API's expected shape.
	// Pass nil to use defaults.
	ShroudConfig interface{}
}

// UpdateAgentParams are parameters for updating an agent.
type UpdateAgentParams struct {
	Name              *string
	Scopes            []string
	IsActive          *bool
	ExpiresAt         *time.Time
	IntentsAPIEnabled *bool
	TxToAllowlist     []string
	TxMaxValueEth     *string
	TxDailyLimitEth   *string
	TxAllowedChains   []string
	TokenTtlSeconds   *int32
	VaultIDs          []string
	ShroudEnabled     *bool
	ShroudConfig      interface{}
}

// --- API Key types ---

// APIKey represents an API key (metadata only).
type APIKey struct {
	ID        string
	Name      string
	KeyPrefix string
	Scopes    []string
	IsActive  bool
	CreatedAt *time.Time
	ExpiresAt *time.Time
	LastUsedAt *time.Time
}

// APIKeyCreated is the response from creating an API key.
type APIKeyCreated struct {
	Key    APIKey
	APIKey string // Full key, shown once
}

// APIKeyList is the response from listing API keys.
type APIKeyList struct {
	Keys []APIKey
}

// --- Auth types ---

// Token represents an auth token response.
type Token struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    *int32
	RefreshToken string
}

// LoginResult is the response from email/password login.
type LoginResult struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    *int32
	RefreshToken string
	MFARequired  bool
	MFAToken     string
}

// UserProfile represents the current user's profile.
type UserProfile struct {
	ID              string
	Email           string
	DisplayName     string
	AuthMethod      string
	Role            string
	EmailVerified   bool
	MarketingEmails bool
	TotpEnabled     bool
	CreatedAt       *time.Time
}
