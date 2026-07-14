// Package oneclaw provides the 1Claw Go SDK for secret management.
package oneclaw

import "time"

// --- Vault types ---

// Vault represents a vault.
type Vault struct {
	ID              string
	Name            string
	Description     string
	CreatedBy       string
	CreatedByType   string
	CreatedAt       time.Time
	CMEKEnabled     bool
	CMEKFingerprint string
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
	ID                      string
	Name                    string
	Description             string
	AuthMethod              string
	Scopes                  []string
	IsActive                bool
	IntentsAPIEnabled       bool
	TxToAllowlist           []string
	TxMaxValueEth           string
	TxDailyLimitEth         string
	TxAllowedChains         []string
	TokenTtlSeconds         *int32
	VaultIDs                []string
	ClientCertFingerprint   string
	OIDCIssuer              string
	OIDCClientID            string
	SSHPublicKey            string
	ECDHPublicKey           string
	ShroudEnabled           bool
	TxTokenAllowlist        []string               `json:"tx_token_allowlist,omitempty"`
	TxKnownTokensOnly       bool                   `json:"tx_known_tokens_only,omitempty"`
	XRPLAllowedTxTypes      []string               `json:"xrpl_allowed_tx_types,omitempty"`
	PerChainGuardrails      map[string]interface{} `json:"per_chain_guardrails,omitempty"`
	TxSpentTodayByChain     map[string]string      `json:"tx_spent_today_by_chain,omitempty"`
	ExecutionIntentsEnabled bool                   `json:"execution_intents_enabled,omitempty"`
	ExecutionGuardrails     map[string]interface{} `json:"execution_guardrails,omitempty"`
	IntentsRequireTee       bool                   `json:"intents_require_tee,omitempty"`
	ExecutionRequireTee     bool                   `json:"execution_require_tee,omitempty"`
	TxMaxPerDay             *int                   `json:"tx_max_per_day,omitempty"`
	TxOverheadBudget        map[string]string      `json:"tx_overhead_budget,omitempty"`
	SolanaAtaAllowlist      []string               `json:"solana_ata_allowlist,omitempty"`
	TxCountToday            *int64                 `json:"tx_count_today,omitempty"`
	TxOverheadTodayByChain  map[string]string      `json:"tx_overhead_today_by_chain,omitempty"`
	CreatedAt               time.Time
	ExpiresAt               *time.Time
	ApiKeyExpiresAt         *time.Time
	LastActiveAt            *time.Time
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
	Name                  string
	Description           string
	AuthMethod            string
	Scopes                []string
	ExpiresAt             *time.Time
	ApiKeyExpiresAt       *time.Time
	IntentsAPIEnabled     bool
	TxToAllowlist         []string
	TxMaxValueEth         string
	TxDailyLimitEth       string
	TxAllowedChains       []string
	TokenTtlSeconds       *int32
	VaultIDs              []string
	ClientCertFingerprint string
	OIDCIssuer            string
	OIDCClientID          string
	ShroudEnabled         bool
	// ShroudConfig can be a struct that marshals to the API's expected shape.
	// Pass nil to use defaults.
	ShroudConfig            interface{}
	TxTokenAllowlist        []string               `json:"tx_token_allowlist,omitempty"`
	TxKnownTokensOnly       bool                   `json:"tx_known_tokens_only,omitempty"`
	XRPLAllowedTxTypes      []string               `json:"xrpl_allowed_tx_types,omitempty"`
	PerChainGuardrails      map[string]interface{} `json:"per_chain_guardrails,omitempty"`
	ExecutionIntentsEnabled bool                   `json:"execution_intents_enabled,omitempty"`
	ExecutionGuardrails     map[string]interface{} `json:"execution_guardrails,omitempty"`
	IntentsRequireTee       bool                   `json:"intents_require_tee,omitempty"`
	ExecutionRequireTee     bool                   `json:"execution_require_tee,omitempty"`
	TxMaxPerDay             *int                   `json:"tx_max_per_day,omitempty"`
	TxOverheadBudget        map[string]string      `json:"tx_overhead_budget,omitempty"`
	SolanaAtaAllowlist      []string               `json:"solana_ata_allowlist,omitempty"`
}

// UpdateAgentParams are parameters for updating an agent.
type UpdateAgentParams struct {
	Name                    *string
	Scopes                  []string
	IsActive                *bool
	ExpiresAt               *time.Time
	ApiKeyExpiresAt         *time.Time
	IntentsAPIEnabled       *bool
	TxToAllowlist           []string
	TxMaxValueEth           *string
	TxDailyLimitEth         *string
	TxAllowedChains         []string
	TokenTtlSeconds         *int32
	VaultIDs                []string
	ShroudEnabled           *bool
	ShroudConfig            interface{}
	TxTokenAllowlist        []string               `json:"tx_token_allowlist,omitempty"`
	TxKnownTokensOnly       *bool                  `json:"tx_known_tokens_only,omitempty"`
	XRPLAllowedTxTypes      []string               `json:"xrpl_allowed_tx_types,omitempty"`
	PerChainGuardrails      map[string]interface{} `json:"per_chain_guardrails,omitempty"`
	ExecutionIntentsEnabled *bool                  `json:"execution_intents_enabled,omitempty"`
	ExecutionGuardrails     map[string]interface{} `json:"execution_guardrails,omitempty"`
	IntentsRequireTee       *bool                  `json:"intents_require_tee,omitempty"`
	TxMaxPerDay             *int                   `json:"tx_max_per_day,omitempty"`
	TxOverheadBudget        map[string]string      `json:"tx_overhead_budget,omitempty"`
	SolanaAtaAllowlist      []string               `json:"solana_ata_allowlist,omitempty"`
	ExecutionRequireTee     *bool                  `json:"execution_require_tee,omitempty"`
}

// --- API Key types ---

// APIKey represents an API key (metadata only).
type APIKey struct {
	ID         string
	Name       string
	KeyPrefix  string
	Scopes     []string
	IsActive   bool
	CreatedAt  *time.Time
	ExpiresAt  *time.Time
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
	SecretPathPattern string
	PrincipalType     string
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
	ID             string
	ShareURL       string
	RecipientType  string
	RecipientEmail string
	ExpiresAt      *time.Time
	MaxAccessCount *int32
}

// ShareList is the response from listing shares.
type ShareList struct {
	Shares []Share
}

// CreateShareParams are parameters for creating a share.
type CreateShareParams struct {
	RecipientType  string
	RecipientID    string
	Email          string
	Permissions    []string
	MaxAccessCount *int32
	ExpiresAt      time.Time
	Passphrase     string
	IPAllowlist    []string
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
	BalanceCents         int32
	BalanceUsd           string
	ExpiringWithin90Days interface{}
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
	Scheme                  string
	Network                 string
	PayTo                   string
	Price                   string
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

// --- Platform types ---

// PlatformApp represents a platform app.
type PlatformApp struct {
	ID                string
	Name              string
	Slug              string
	Description       string
	LogoURL           string
	APIKeyPrefix      string
	OIDCJwksURL       string
	OIDCIssuer        string
	OIDCAudience      string
	RedirectURIs      []string
	WebhookURL        string
	IsActive          bool
	BillingModel      string
	AuthMode          string
	MaxConnectedUsers *int32
	ConnectedUsers    int32
	ApiKeyExpiresAt   *time.Time
	ApiKeyRotatedAt   *time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

// PlatformAppList is the response from listing platform apps.
type PlatformAppList struct {
	Apps []PlatformApp
}

// PlatformAppCreated is the response from creating a platform app.
type PlatformAppCreated struct {
	PlatformApp
	APIKey string
}

// CreatePlatformAppRequest are parameters for creating a platform app.
type CreatePlatformAppRequest struct {
	Name              string
	Slug              string
	Description       string
	OIDCJwksURL       string
	OIDCIssuer        string
	OIDCAudience      string
	RedirectURIs      []string
	BillingModel      string
	AuthMode          string
	MaxConnectedUsers *int32
	ApiKeyExpiresAt   *time.Time
}

// UpdatePlatformAppRequest are parameters for updating a platform app.
type UpdatePlatformAppRequest struct {
	Name              *string
	Description       *string
	LogoURL           *string
	OIDCJwksURL       *string
	OIDCIssuer        *string
	OIDCAudience      *string
	RedirectURIs      []string
	WebhookURL        *string
	BillingModel      *string
	AuthMode          *string
	MaxConnectedUsers *int32
	IsActive          *bool
}

// --- Signing Key types ---

// SigningKey represents a multi-chain signing key for an agent.
type SigningKey struct {
	ID         string
	AgentID    string
	Chain      string
	Curve      string
	PublicKey  string
	Address    string
	KeyVersion int32
	IsActive   bool
	CreatedAt  time.Time
	RotatedAt  *time.Time
}

// SigningKeyList is the response from listing signing keys.
type SigningKeyList struct {
	Keys []SigningKey
}

// CreateSigningKeyParams are parameters for provisioning a signing key.
type CreateSigningKeyParams struct {
	Chain string
}

// SigningKeyExport is the response from exporting a signing key's private key.
type SigningKeyExport struct {
	Chain      string  `json:"chain"`
	Curve      string  `json:"curve"`
	PublicKey  string  `json:"public_key"`
	Address    *string `json:"address,omitempty"`
	PrivateKey string  `json:"private_key"`
	KeyVersion int     `json:"key_version"`
	AgentID    string  `json:"agent_id"`
}

// SignIntentParams are parameters for the unified sign endpoint.
type SignIntentParams struct {
	IntentType        string                 `json:"intent_type"`
	Chain             string                 `json:"chain"`
	SigningKeyPath    string                 `json:"signing_key_path,omitempty"`
	Message           string                 `json:"message,omitempty"`
	TypedData         interface{}            `json:"typed_data,omitempty"`
	Hash              string                 `json:"hash,omitempty"`
	To                string                 `json:"to,omitempty"`
	Value             string                 `json:"value,omitempty"`
	TxType            *int32                 `json:"tx_type,omitempty"`
	Data              string                 `json:"data,omitempty"`
	Nonce             *int64                 `json:"nonce,omitempty"`
	GasLimit          *int64                 `json:"gas_limit,omitempty"`
	GasPrice          string                 `json:"gas_price,omitempty"`
	MaxFeePerGas      string                 `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFee    string                 `json:"max_priority_fee_per_gas,omitempty"`
	AccessList        interface{}            `json:"access_list,omitempty"`
	MaxFeePerBlobGas  string                 `json:"max_fee_per_blob_gas,omitempty"`
	BlobVersionedHash []string               `json:"blob_versioned_hashes,omitempty"`
	AuthorizationList interface{}            `json:"authorization_list,omitempty"`
	SignOnly          *bool                  `json:"sign_only,omitempty"`
	DestinationTag    *int64                 `json:"destination_tag,omitempty"`
	Memo              string                 `json:"memo,omitempty"`
	FeeRateSatPerVB   *int64                 `json:"fee_rate_sat_per_vbyte,omitempty"`
	FeeLimitSun       *int64                 `json:"fee_limit_sun,omitempty"`
	TokenMint         string                 `json:"token_mint,omitempty"`
	TokenDecimals     *int32                 `json:"token_decimals,omitempty"`
	TTL               *int64                 `json:"ttl,omitempty"`
	XrplTxJSON        map[string]interface{} `json:"xrpl_tx_json,omitempty"`
}

// SignIntentResult is the response from the unified sign endpoint.
type SignIntentResult struct {
	IntentType    string
	Chain         string
	From          string
	Signature     string
	SignedTx      string
	TxHash        string
	MessageHash   string
	TypedDataHash string
	TxType        *int32
}

// --- Known Token types ---

// KnownToken represents a known/verified token in the chain registry.
type KnownToken struct {
	ID              string  `json:"id"`
	Chain           string  `json:"chain"`
	Symbol          string  `json:"symbol"`
	Name            string  `json:"name"`
	ContractAddress string  `json:"contract_address"`
	Decimals        int     `json:"decimals"`
	IsTestnet       bool    `json:"is_testnet"`
	IsVerified      bool    `json:"is_verified"`
	LogoURL         *string `json:"logo_url,omitempty"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// KnownTokenListResponse is the response from listing known tokens.
type KnownTokenListResponse struct {
	Tokens []KnownToken `json:"tokens"`
}

// --- Payment Card types ---

// OrderCardRequest are parameters for ordering a prepaid or gift card.
type OrderCardRequest struct {
	Kind         string  `json:"kind"`
	AmountUSD    string  `json:"amount_usd"`
	LasoServerID *string `json:"laso_server_id,omitempty"`
	Country      *string `json:"country,omitempty"`
}

// CardResponse is a masked card view — never contains PAN/CVV.
type CardResponse struct {
	ID             string                 `json:"id"`
	AgentID        *string                `json:"agent_id,omitempty"`
	Issuer         string                 `json:"issuer"`
	Kind           string                 `json:"kind"`
	Brand          *string                `json:"brand,omitempty"`
	Last4          *string                `json:"last4,omitempty"`
	ExpMonth       *int                   `json:"exp_month,omitempty"`
	ExpYear        *int                   `json:"exp_year,omitempty"`
	Currency       string                 `json:"currency"`
	OrderAmountUSD *string                `json:"order_amount_usd,omitempty"`
	Balance        *string                `json:"balance,omitempty"`
	Status         string                 `json:"status"`
	StorageMode    string                 `json:"storage_mode"`
	RevealPolicy   map[string]interface{} `json:"reveal_policy"`
	VoidAfter      *string                `json:"void_after,omitempty"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
}

// CardListResponse is the response from listing cards.
type CardListResponse struct {
	Cards []CardResponse `json:"cards"`
}

// CardRevealResponse contains full card details (sensitive).
type CardRevealResponse struct {
	ID         string      `json:"id"`
	PAN        *string     `json:"pan,omitempty"`
	CVV        *string     `json:"cvv,omitempty"`
	ExpMonth   *int        `json:"exp_month,omitempty"`
	ExpYear    *int        `json:"exp_year,omitempty"`
	Brand      *string     `json:"brand,omitempty"`
	Redemption interface{} `json:"redemption,omitempty"`
	Disclaimer string      `json:"disclaimer"`
}

// UpdateCardRequest are parameters for updating a card's reveal policy.
type UpdateCardRequest struct {
	AgentReveal     *bool   `json:"agent_reveal,omitempty"`
	MaxReveals      *int    `json:"max_reveals,omitempty"`
	RevealExpiresAt *string `json:"reveal_expires_at,omitempty"`
	VoidAfter       *string `json:"void_after,omitempty"`
}

// ImportCardRequest are parameters for manually importing a card.
type ImportCardRequest struct {
	PAN      string  `json:"pan"`
	CVV      string  `json:"cvv"`
	ExpMonth int     `json:"exp_month"`
	ExpYear  int     `json:"exp_year"`
	Brand    *string `json:"brand,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Balance  *string `json:"balance,omitempty"`
	AgentID  *string `json:"agent_id,omitempty"`
}

// SearchGiftCardsRequest are parameters for searching gift-card brands.
type SearchGiftCardsRequest struct {
	Query   *string `json:"query,omitempty"`
	Country *string `json:"country,omitempty"`
}

// SearchGiftCardsResponse is the response from searching gift cards.
type SearchGiftCardsResponse struct {
	Results []map[string]interface{} `json:"results"`
}

// --- Binding types (Execution Intents) ---

// CredentialSource describes how a binding's credential is sourced:
// inline (copied into __agent-keys) or vault_ref (live pointer to an existing secret).
type CredentialSource struct {
	Type    string                 `json:"type"`
	Value   map[string]interface{} `json:"value,omitempty"`
	VaultID string                 `json:"vault_id,omitempty"`
	Path    string                 `json:"path,omitempty"`
}

// Binding represents a named credential handle for execution intents.
type Binding struct {
	ID                   string                 `json:"id"`
	AgentID              string                 `json:"agent_id"`
	BindingType          string                 `json:"binding_type"`
	Name                 string                 `json:"name"`
	Config               map[string]interface{} `json:"config"`
	Guardrails           map[string]interface{} `json:"guardrails"`
	IsActive             bool                   `json:"is_active"`
	CredentialSet        bool                   `json:"credential_set,omitempty"`
	CredentialSourceType *string                `json:"credential_source_type,omitempty"`
	CredentialVaultID    *string                `json:"credential_vault_id,omitempty"`
	CredentialPath       *string                `json:"credential_path,omitempty"`
	CreatedAt            string                 `json:"created_at"`
	UpdatedAt            string                 `json:"updated_at"`
}

// BindingList is the response from listing bindings.
type BindingList struct {
	Bindings []Binding `json:"bindings"`
}

// CreateBindingParams are parameters for creating a binding.
type CreateBindingParams struct {
	Name             string                 `json:"name"`
	BindingType      string                 `json:"binding_type"`
	Config           map[string]interface{} `json:"config,omitempty"`
	Guardrails       map[string]interface{} `json:"guardrails,omitempty"`
	Credential       map[string]interface{} `json:"credential,omitempty"`
	CredentialSource *CredentialSource      `json:"credential_source,omitempty"`
}

// UpdateBindingParams are parameters for updating a binding.
type UpdateBindingParams struct {
	Config           map[string]interface{} `json:"config,omitempty"`
	Guardrails       map[string]interface{} `json:"guardrails,omitempty"`
	IsActive         *bool                  `json:"is_active,omitempty"`
	Credential       map[string]interface{} `json:"credential,omitempty"`
	CredentialSource *CredentialSource      `json:"credential_source,omitempty"`
}

// RotateCredentialParams are parameters for rotating a binding credential.
type RotateCredentialParams struct {
	Credential map[string]interface{} `json:"credential"`
}

// TestBindingParams are optional parameters for testing binding connectivity.
type TestBindingParams struct {
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
}

// TestBindingResult is the response from testing a binding.
type TestBindingResult struct {
	Success   bool    `json:"success"`
	LatencyMs int64   `json:"latency_ms"`
	Error     *string `json:"error,omitempty"`
}

// ExecuteParams are parameters for executing an intent through a binding.
type ExecuteParams struct {
	Binding       string                 `json:"binding"`
	IntentType    string                 `json:"intent_type"`
	ExecutionMode string                 `json:"execution_mode,omitempty"`
	Params        map[string]interface{} `json:"params"`
}

// ExecuteResult is the response from executing an intent.
type ExecuteResult struct {
	ExecutionID       string                 `json:"execution_id"`
	Status            string                 `json:"status"`
	Result            map[string]interface{} `json:"result,omitempty"`
	Error             *string                `json:"error,omitempty"`
	DurationMs        *int64                 `json:"duration_ms,omitempty"`
	RedactionsApplied *int32                 `json:"redactions_applied,omitempty"`
}

// ExecutionEvent represents a single execution event for audit/billing.
type ExecutionEvent struct {
	ID                string                 `json:"id"`
	AgentID           string                 `json:"agent_id"`
	BindingID         string                 `json:"binding_id"`
	IntentType        string                 `json:"intent_type"`
	ExecutionMode     string                 `json:"execution_mode"`
	Status            string                 `json:"status"`
	RequestSummary    map[string]interface{} `json:"request_summary,omitempty"`
	ResultSummary     map[string]interface{} `json:"result_summary,omitempty"`
	ErrorMessage      *string                `json:"error_message,omitempty"`
	DurationMs        *int64                 `json:"duration_ms,omitempty"`
	CostCents         *int32                 `json:"cost_cents,omitempty"`
	RedactionsApplied *int32                 `json:"redactions_applied,omitempty"`
	CreatedAt         string                 `json:"created_at"`
}

// ExecutionEventList is the response from listing execution events.
type ExecutionEventList struct {
	Events []ExecutionEvent `json:"events"`
}
