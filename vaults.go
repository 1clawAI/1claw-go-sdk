package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// Create creates a new vault.
func (s *VaultsService) Create(ctx context.Context, name string, description string) (*Vault, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := openapi.CreateVaultRequest{Name: name}
	if description != "" {
		req.Description = &description
	}
	resp, _, err := s.client.api.VaultsAPI.CreateVault(authCtx).
		CreateVaultRequest(req).
		Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return vaultFromAPI(resp), nil
}

// List lists vaults for the current user/org.
func (s *VaultsService) List(ctx context.Context) (*VaultList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.VaultsAPI.ListVaults(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return vaultListFromAPI(resp), nil
}

// Get retrieves a vault by ID.
func (s *VaultsService) Get(ctx context.Context, vaultID string) (*Vault, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.VaultsAPI.GetVault(authCtx, vaultID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return vaultFromAPI(resp), nil
}

// Delete deletes a vault.
func (s *VaultsService) Delete(ctx context.Context, vaultID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.VaultsAPI.DeleteVault(authCtx, vaultID).Execute()
	return wrapAPIError(err)
}
