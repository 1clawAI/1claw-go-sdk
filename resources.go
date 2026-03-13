package oneclaw

// AuthService provides authentication operations.
type AuthService struct {
	client *Client
}

// VaultsService provides vault operations.
type VaultsService struct {
	client *Client
}

// SecretsService provides secret operations.
type SecretsService struct {
	client *Client
}

// AgentsService provides agent operations.
type AgentsService struct {
	client *Client
}

// APIKeysService provides API key management operations.
type APIKeysService struct {
	client *Client
}

// SharingService provides secret sharing operations.
type SharingService struct {
	client *Client
}

// AccessService provides access policy operations.
type AccessService struct {
	client *Client
}

// OrgService provides organization operations.
type OrgService struct {
	client *Client
}

// ChainsService provides chain registry operations.
type ChainsService struct {
	client *Client
}

// BillingService provides billing operations.
type BillingService struct {
	client *Client
}

// AuditService provides audit log operations.
type AuditService struct {
	client *Client
}

// X402Service provides x402 payment operations.
type X402Service struct {
	client *Client
}
