package oneclaw

import (
	"context"
	"fmt"
)

// ContractAbisService manages the org contract ABI registry.
type ContractAbisService struct {
	client *Client
}

// Create registers a contract ABI.
func (s *ContractAbisService) Create(ctx context.Context, req CreateContractAbiRequest) (*ContractAbiResponse, error) {
	var resp ContractAbiResponse
	err := s.client.doJSON(ctx, "POST", "/v1/org/contract-abis", req, &resp)
	return &resp, err
}

// List returns org contract ABIs, optionally filtered by chain.
func (s *ContractAbisService) List(ctx context.Context, chain string) (*ContractAbiListResponse, error) {
	path := "/v1/org/contract-abis"
	if chain != "" {
		path = fmt.Sprintf("/v1/org/contract-abis?chain=%s", chain)
	}
	var resp ContractAbiListResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &resp)
	return &resp, err
}

// Get retrieves a contract ABI by ID.
func (s *ContractAbisService) Get(ctx context.Context, id string) (*ContractAbiResponse, error) {
	var resp ContractAbiResponse
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/org/contract-abis/%s", id), nil, &resp)
	return &resp, err
}

// Delete removes a contract ABI.
func (s *ContractAbisService) Delete(ctx context.Context, id string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/org/contract-abis/%s", id), nil, nil)
}
