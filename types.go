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

// --- Access/Policy types ---

// Policy represents an access policy.
type Policy struct {
	ID                string
	VaultID           string
	SecretPathPattern  string
	PrincipalType      string
	PrincipalID       string
	Permissions       []string
	Conditions        map[string]interface{}
	ExpiresAt         *time.Time
	CreatedBy         string
	CreatedByType     string
	CreatedAt         time.Time
}

// PolicyList is the response from listing policies.
type PolicyList struct {
	Policies []Policy
}

// CreatePolicyParams are parameters for creating a policy.
type CreatePolicyParams struct {
	SecretPathPattern string
	PrincipalType     string
	PrincipalID       string
	Permissions       []string
	Conditions        map[string]interface{}
	ExpiresAt         *time.Time
}

// UpdatePolicyParams are parameters for updating a policy.
type UpdatePolicyParams struct {
	SecretPathPattern *string
	Permissions       []string
	Conditions        map[string]interface{}
	ExpiresAt         *time.Time
}

// --- Sharing types ---

// Share represents a share.
type Share struct {
	ID               string
	ShareURL         string
	RecipientType    string
	RecipientEmail   string
	ExpiresAt        *time.Time
	MaxAccessCount   *int32
}

// ShareList is the response from listing shares.
type ShareList struct {
	Shares []Share
}

// CreateShareParams are parameters for creating a share.
type CreateShareParams struct {
	RecipientType   string
	RecipientID     string
	Email           string
	Permissions     []string
	MaxAccessCount  *int32
	ExpiresAt       time.Time
	Passphrase      string
	IPAllowlist     []string
}

// --- Org types ---

// OrgMember represents an organization member.
type OrgMember struct {
	ID          string
	Email       string
	DisplayName string
	Role        string
	AuthMethod  string
	CreatedAt   *time.Time
}

// OrgMemberList is the response from listing org members.
type OrgMemberList struct {
	Members []OrgMember
}

// InviteMemberParams are parameters for inviting a member.
type InviteMemberParams struct {
	Email string
	Role  string
}

// InviteMemberResult is the response from inviting a member.
type InviteMemberResult struct {
	Message string
	Email   string
}

// UpdateMemberRoleParams are parameters for updating a member's role.
type UpdateMemberRoleParams struct {
	Role string
}

// --- Chain types ---

// Chain represents a blockchain.
type Chain struct {
	ID             string
	Name           string
	DisplayName    string
	ChainID        *int32
	RPCURL         string
	WSURL          string
	ExplorerURL    string
	NativeCurrency string
	IsTestnet      bool
	IsEnabled      bool
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// ChainList is the response from listing chains.
type ChainList struct {
	Chains []Chain
}

// --- Billing types ---

// Subscription is the billing subscription response.
type Subscription struct {
	Tier               string
	Interval           string
	PeriodEnd          *time.Time
	Status             string
	CreditBalanceCents int32
	CreditBalanceUsd   string
	OverageMethod      string
	Usage              interface{}
}

// CreditBalance is the credit balance response.
type CreditBalance struct {
	BalanceCents           int32
	BalanceUsd             string
	ExpiringWithin90Days   interface{}
}

// --- Audit types ---

// AuditEvent represents an audit log event.
type AuditEvent struct {
	ID           string
	Action       string
	ActorID      string
	ActorType    string
	ResourceType string
	ResourceID   string
	OrgID        string
	Details      map[string]interface{}
	IPAddress    string
	CreatedAt    *time.Time
}

// AuditEvents is the response from querying audit events.
type AuditEvents struct {
	Events []AuditEvent
	Count  int32
}

// --- X402 types ---

// PaymentRequirementAccept represents a payment option in a 402 response.
type PaymentRequirementAccept struct {
	Scheme                 string
	Network                string
	PayTo                  string
	Price                  string
	RequiredDeadlineSeconds *int32
}

// PaymentRequirement represents a 402 payment requirement from the API.
type PaymentRequirement struct {
	X402Version *int32
	Accepts     []PaymentRequirementAccept
	Description string
}

// --- Treasury types ---

// TreasurySigner is a signer on a treasury Safe.
type TreasurySigner struct {
	ID            string
	SignerType    string
	SignerID      string
	SignerAddress string
	AddedAt       *time.Time
}

// Treasury is a treasury (Safe multisig) record.
type Treasury struct {
	ID          string
	Name        string
	SafeAddress string
	Chain       string
	ChainID     *int32
	Threshold   *int32
	CreatedBy   string
	Signers     []TreasurySigner
	CreatedAt   *time.Time
}

// TreasuryList is the response from listing treasuries.
type TreasuryList struct {
	Treasuries []Treasury
}

// CreateTreasurySignerEntry is a signer to attach when creating a treasury.
type CreateTreasurySignerEntry struct {
	SignerType    string
	SignerID      string
	SignerAddress string
}

// CreateTreasuryParams are parameters for creating a treasury.
type CreateTreasuryParams struct {
	Name        string
	SafeAddress string
	Chain       *string
	ChainID     *int32
	Threshold   *int32
	Signers     []CreateTreasurySignerEntry
}

// UpdateTreasuryParams are parameters for patching a treasury (all fields optional).
type UpdateTreasuryParams struct {
	Name      *string
	Threshold *int32
}

// TreasuryAccessRequest is an agent access request for a treasury.
type TreasuryAccessRequest struct {
	ID          string
	TreasuryID  string
	AgentID     string
	Status      string
	Reason      string
	RequestedAt *time.Time
	ResolvedBy  string
	ResolvedAt  *time.Time
}

// TreasuryAccessRequestList is the response from listing access requests.
type TreasuryAccessRequestList struct {
	Requests []TreasuryAccessRequest
}

// AddTreasurySignerParams are parameters for adding a signer.
type AddTreasurySignerParams struct {
	SignerType    string
	SignerID      string
	SignerAddress string
}
