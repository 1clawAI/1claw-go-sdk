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
