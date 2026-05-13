package oneclaw

import (
	"context"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// Create creates a treasury (Safe multisig).
func (s *TreasuryService) Create(ctx context.Context, params CreateTreasuryParams) (*Treasury, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := openapi.NewCreateTreasuryRequest(params.Name, params.SafeAddress)
	if params.Chain != nil {
		req.SetChain(*params.Chain)
	}
	if params.ChainID != nil {
		req.SetChainId(*params.ChainID)
	}
	if params.Threshold != nil {
		req.SetThreshold(*params.Threshold)
	}
	if len(params.Signers) > 0 {
		signers := make([]openapi.CreateTreasurySignerEntry, 0, len(params.Signers))
		for _, e := range params.Signers {
			signers = append(signers, *openapi.NewCreateTreasurySignerEntry(e.SignerType, e.SignerID, e.SignerAddress))
		}
		req.SetSigners(signers)
	}
	resp, _, err := s.client.api.TreasuryAPI.CreateTreasury(authCtx).CreateTreasuryRequest(*req).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return treasuryFromAPI(resp), nil
}

// List lists treasuries in the org.
func (s *TreasuryService) List(ctx context.Context) (*TreasuryList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.TreasuryAPI.ListTreasuries(authCtx).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return treasuryListFromAPI(resp), nil
}

// Get returns treasury details.
func (s *TreasuryService) Get(ctx context.Context, treasuryID string) (*Treasury, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.TreasuryAPI.GetTreasury(authCtx, treasuryID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return treasuryFromAPI(resp), nil
}

// Update patches treasury name and/or threshold.
func (s *TreasuryService) Update(ctx context.Context, treasuryID string, params UpdateTreasuryParams) (*Treasury, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	req := openapi.NewUpdateTreasuryRequest()
	if params.Name != nil {
		req.SetName(*params.Name)
	}
	if params.Threshold != nil {
		req.SetThreshold(*params.Threshold)
	}
	resp, _, err := s.client.api.TreasuryAPI.UpdateTreasury(authCtx, treasuryID).UpdateTreasuryRequest(*req).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return treasuryFromAPI(resp), nil
}

// Delete deletes a treasury and its signers.
func (s *TreasuryService) Delete(ctx context.Context, treasuryID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.TreasuryAPI.DeleteTreasury(authCtx, treasuryID).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	return nil
}

// AddSigner adds a signer to a treasury.
func (s *TreasuryService) AddSigner(ctx context.Context, treasuryID string, params AddTreasurySignerParams) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	req := openapi.NewAddSignerRequest(params.SignerType, params.SignerID, params.SignerAddress)
	_, err = s.client.api.TreasuryAPI.AddTreasurySigner(authCtx, treasuryID).AddSignerRequest(*req).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	return nil
}

// RemoveSigner removes a signer from a treasury.
func (s *TreasuryService) RemoveSigner(ctx context.Context, treasuryID, signerID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.TreasuryAPI.RemoveTreasurySigner(authCtx, treasuryID, signerID).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	return nil
}

// ListAccessRequests lists access requests for a treasury.
func (s *TreasuryService) ListAccessRequests(ctx context.Context, treasuryID string) (*TreasuryAccessRequestList, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.TreasuryAPI.ListTreasuryAccessRequests(authCtx, treasuryID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	return treasuryAccessRequestListFromAPI(resp), nil
}

// RequestAccess submits an access request (agent-only).
func (s *TreasuryService) RequestAccess(ctx context.Context, treasuryID string) (*TreasuryAccessRequest, error) {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, _, err := s.client.api.TreasuryAPI.RequestTreasuryAccess(authCtx, treasuryID).Execute()
	if err != nil {
		return nil, wrapAPIError(err)
	}
	ar := treasuryAccessRequestFromAPI(resp)
	return &ar, nil
}

// ApproveAccess approves a treasury access request.
func (s *TreasuryService) ApproveAccess(ctx context.Context, treasuryID, requestID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.TreasuryAPI.ApproveTreasuryAccess(authCtx, treasuryID, requestID).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	return nil
}

// DenyAccess denies a treasury access request.
func (s *TreasuryService) DenyAccess(ctx context.Context, treasuryID, requestID string) error {
	authCtx, err := s.client.authContext(ctx)
	if err != nil {
		return err
	}
	_, err = s.client.api.TreasuryAPI.DenyTreasuryAccess(authCtx, treasuryID, requestID).Execute()
	if err != nil {
		return wrapAPIError(err)
	}
	return nil
}

// ── Treasury Wallets (multi-chain, human-only, Pro+) ──────────────

// TreasuryWallet represents a generated treasury wallet for a specific chain.
type TreasuryWallet struct {
	ID           string `json:"id"`
	Chain        string `json:"chain"`
	Curve        string `json:"curve"`
	PublicKeyHex string `json:"public_key_hex"`
	Address      string `json:"address"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
}

// TreasuryWalletList is the response from listing treasury wallets.
type TreasuryWalletList struct {
	Wallets []TreasuryWallet `json:"wallets"`
}

// GenerateTreasuryWalletsParams configures which chains to generate wallets for.
type GenerateTreasuryWalletsParams struct {
	Chains []string `json:"chains,omitempty"`
}

// TreasuryWalletExport contains the exported private key for a treasury wallet.
type TreasuryWalletExport struct {
	Chain         string `json:"chain"`
	Address       string `json:"address"`
	PrivateKeyHex string `json:"private_key_hex"`
}

// GenerateWallets generates treasury wallets for the specified chains.
// Human-only, requires Pro+ subscription.
func (s *TreasuryService) GenerateWallets(ctx context.Context, params GenerateTreasuryWalletsParams) (*TreasuryWalletList, error) {
	var result TreasuryWalletList
	err := s.client.doJSON(ctx, "POST", "/v1/treasury/wallets/generate", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListWallets returns all active treasury wallets for the calling user.
func (s *TreasuryService) ListWallets(ctx context.Context) (*TreasuryWalletList, error) {
	var result TreasuryWalletList
	err := s.client.doJSON(ctx, "GET", "/v1/treasury/wallets", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWallet returns the active treasury wallet for a specific chain.
func (s *TreasuryService) GetWallet(ctx context.Context, chain string) (*TreasuryWallet, error) {
	var result TreasuryWallet
	err := s.client.doJSON(ctx, "GET", "/v1/treasury/wallets/"+chain, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportWallet exports the private key for a treasury wallet (audit-logged).
// Requires re-authentication via the account password.
func (s *TreasuryService) ExportWallet(ctx context.Context, chain, password string) (*TreasuryWalletExport, error) {
	var result TreasuryWalletExport
	err := s.client.doJSONWithHeaders(ctx, "POST", "/v1/treasury/wallets/"+chain+"/export", nil, &result, map[string]string{
		"X-Auth-Confirm": password,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RotateWallet rotates the key for a treasury wallet on the given chain.
func (s *TreasuryService) RotateWallet(ctx context.Context, chain string) (*TreasuryWallet, error) {
	var result TreasuryWallet
	err := s.client.doJSON(ctx, "POST", "/v1/treasury/wallets/"+chain+"/rotate", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeactivateWallet deactivates the treasury wallet for a specific chain.
func (s *TreasuryService) DeactivateWallet(ctx context.Context, chain string) error {
	return s.client.doJSON(ctx, "DELETE", "/v1/treasury/wallets/"+chain, nil, nil)
}
