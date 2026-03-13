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
