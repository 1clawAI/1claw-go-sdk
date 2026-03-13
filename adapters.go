package oneclaw

import (
	"encoding/json"

	"github.com/1clawAI/1claw-go-sdk/internal/openapi"
)

// --- Vault adapters ---

func vaultFromAPI(o *openapi.VaultResponse) *Vault {
	if o == nil {
		return nil
	}
	v := &Vault{
		ID:   o.Id,
		Name: o.Name,
		CreatedAt: o.CreatedAt,
	}
	if o.Description != nil {
		v.Description = *o.Description
	}
	if o.CreatedBy != nil {
		v.CreatedBy = *o.CreatedBy
	}
	if o.CreatedByType != nil {
		v.CreatedByType = *o.CreatedByType
	}
	if o.CmekEnabled != nil {
		v.CMEKEnabled = *o.CmekEnabled
	}
	if o.CmekFingerprint != nil {
		v.CMEKFingerprint = *o.CmekFingerprint
	}
	return v
}

func vaultListFromAPI(o *openapi.VaultListResponse) *VaultList {
	if o == nil {
		return nil
	}
	list := &VaultList{}
	for i := range o.Vaults {
		list.Vaults = append(list.Vaults, *vaultFromAPI(&o.Vaults[i]))
	}
	return list
}

// --- Secret adapters ---

func secretFromAPI(o *openapi.SecretResponse) *Secret {
	if o == nil {
		return nil
	}
	s := &Secret{
		ID:      o.Id,
		Path:    o.Path,
		Type:    o.Type,
		Value:   o.Value,
		Version: o.Version,
		CreatedAt: o.CreatedAt,
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.CreatedBy != nil {
		s.CreatedBy = *o.CreatedBy
	}
	if o.ExpiresAt != nil {
		s.ExpiresAt = o.ExpiresAt
	}
	if o.CmekEncrypted != nil {
		s.CMEKEncrypted = *o.CmekEncrypted
	}
	return s
}

func secretMetadataFromAPI(o *openapi.SecretMetadataResponse) SecretMetadata {
	if o == nil {
		return SecretMetadata{}
	}
	s := SecretMetadata{
		ID:        o.Id,
		Path:      o.Path,
		Type:      o.Type,
		Version:   o.Version,
		CreatedAt: o.CreatedAt,
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.ExpiresAt != nil {
		s.ExpiresAt = o.ExpiresAt
	}
	return s
}

func secretListFromAPI(o *openapi.SecretListResponse) *SecretList {
	if o == nil {
		return nil
	}
	list := &SecretList{}
	for i := range o.Secrets {
		list.Secrets = append(list.Secrets, secretMetadataFromAPI(&o.Secrets[i]))
	}
	return list
}

// --- Agent adapters ---

func agentFromAPI(o *openapi.AgentResponse) Agent {
	if o == nil {
		return Agent{}
	}
	a := Agent{
		ID:                o.Id,
		Name:              o.Name,
		AuthMethod:        o.AuthMethod,
		Scopes:            o.Scopes,
		IsActive:          o.IsActive,
		IntentsAPIEnabled: o.IntentsApiEnabled,
		ShroudEnabled:     o.ShroudEnabled,
		CreatedAt:         o.CreatedAt,
	}
	if o.Description != nil {
		a.Description = *o.Description
	}
	if o.TxToAllowlist != nil {
		a.TxToAllowlist = o.TxToAllowlist
	}
	if o.TxMaxValueEth != nil {
		a.TxMaxValueEth = *o.TxMaxValueEth
	}
	if o.TxDailyLimitEth != nil {
		a.TxDailyLimitEth = *o.TxDailyLimitEth
	}
	if o.TxAllowedChains != nil {
		a.TxAllowedChains = o.TxAllowedChains
	}
	if o.TokenTtlSeconds != nil {
		a.TokenTtlSeconds = o.TokenTtlSeconds
	}
	if o.VaultIds != nil {
		a.VaultIDs = o.VaultIds
	}
	if o.ClientCertFingerprint != nil {
		a.ClientCertFingerprint = *o.ClientCertFingerprint
	}
	if o.OidcIssuer != nil {
		a.OIDCIssuer = *o.OidcIssuer
	}
	if o.OidcClientId != nil {
		a.OIDCClientID = *o.OidcClientId
	}
	if o.SshPublicKey != nil {
		a.SSHPublicKey = *o.SshPublicKey
	}
	if o.EcdhPublicKey != nil {
		a.ECDHPublicKey = *o.EcdhPublicKey
	}
	if o.ExpiresAt != nil {
		a.ExpiresAt = o.ExpiresAt
	}
	if o.LastActiveAt != nil {
		a.LastActiveAt = o.LastActiveAt
	}
	return a
}

func agentCreatedFromAPI(o *openapi.AgentCreatedResponse) *AgentCreated {
	if o == nil {
		return nil
	}
	ac := &AgentCreated{
		Agent: agentFromAPI(&o.Agent),
	}
	if o.ApiKey != nil {
		ac.APIKey = *o.ApiKey
	}
	return ac
}

func agentListFromAPI(o *openapi.AgentListResponse) *AgentList {
	if o == nil {
		return nil
	}
	list := &AgentList{}
	for i := range o.Agents {
		list.Agents = append(list.Agents, agentFromAPI(&o.Agents[i]))
	}
	return list
}

func createAgentParamsToAPI(p CreateAgentParams) openapi.CreateAgentRequest {
	req := openapi.CreateAgentRequest{Name: p.Name}
	if p.Description != "" {
		req.Description = &p.Description
	}
	if p.AuthMethod != "" {
		req.AuthMethod = &p.AuthMethod
	}
	if len(p.Scopes) > 0 {
		req.Scopes = p.Scopes
	}
	if p.ExpiresAt != nil {
		req.ExpiresAt = p.ExpiresAt
	}
	req.IntentsApiEnabled = &p.IntentsAPIEnabled
	if len(p.TxToAllowlist) > 0 {
		req.TxToAllowlist = p.TxToAllowlist
	}
	if p.TxMaxValueEth != "" {
		req.TxMaxValueEth = &p.TxMaxValueEth
	}
	if p.TxDailyLimitEth != "" {
		req.TxDailyLimitEth = &p.TxDailyLimitEth
	}
	if len(p.TxAllowedChains) > 0 {
		req.TxAllowedChains = p.TxAllowedChains
	}
	if p.TokenTtlSeconds != nil {
		req.TokenTtlSeconds = p.TokenTtlSeconds
	}
	if len(p.VaultIDs) > 0 {
		req.VaultIds = p.VaultIDs
	}
	if p.ClientCertFingerprint != "" {
		req.ClientCertFingerprint = &p.ClientCertFingerprint
	}
	if p.OIDCIssuer != "" {
		req.OidcIssuer = &p.OIDCIssuer
	}
	if p.OIDCClientID != "" {
		req.OidcClientId = &p.OIDCClientID
	}
	shroudEnabled := p.ShroudEnabled
	req.ShroudEnabled = &shroudEnabled
	if p.ShroudConfig != nil {
		if data, err := json.Marshal(p.ShroudConfig); err == nil {
			var cfg openapi.ShroudConfig
			if json.Unmarshal(data, &cfg) == nil {
				req.ShroudConfig = &cfg
			}
		}
	}
	return req
}

func updateAgentParamsToAPI(p UpdateAgentParams) openapi.UpdateAgentRequest {
	var req openapi.UpdateAgentRequest
	req.Name = p.Name
	req.Scopes = p.Scopes
	req.IsActive = p.IsActive
	req.ExpiresAt = p.ExpiresAt
	req.IntentsApiEnabled = p.IntentsAPIEnabled
	req.TxToAllowlist = p.TxToAllowlist
	req.TxMaxValueEth = p.TxMaxValueEth
	req.TxDailyLimitEth = p.TxDailyLimitEth
	req.TxAllowedChains = p.TxAllowedChains
	req.TokenTtlSeconds = p.TokenTtlSeconds
	req.VaultIds = p.VaultIDs
	req.ShroudEnabled = p.ShroudEnabled
	if p.ShroudConfig != nil {
		if data, err := json.Marshal(p.ShroudConfig); err == nil {
			var cfg openapi.ShroudConfig
			if json.Unmarshal(data, &cfg) == nil {
				req.ShroudConfig = &cfg
			}
		}
	}
	return req
}

// --- API Key adapters ---

func apiKeyFromAPI(o *openapi.ApiKeyResponse) APIKey {
	if o == nil {
		return APIKey{}
	}
	a := APIKey{}
	if o.Id != nil {
		a.ID = *o.Id
	}
	if o.Name != nil {
		a.Name = *o.Name
	}
	if o.KeyPrefix != nil {
		a.KeyPrefix = *o.KeyPrefix
	}
	if o.Scopes != nil {
		a.Scopes = o.Scopes
	}
	if o.IsActive != nil {
		a.IsActive = *o.IsActive
	}
	if o.CreatedAt != nil {
		a.CreatedAt = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		a.ExpiresAt = o.ExpiresAt
	}
	if o.LastUsedAt != nil {
		a.LastUsedAt = o.LastUsedAt
	}
	return a
}

func apiKeyCreatedFromAPI(o *openapi.ApiKeyCreatedResponse) *APIKeyCreated {
	if o == nil {
		return nil
	}
	ac := &APIKeyCreated{}
	if o.Key != nil {
		ac.Key = apiKeyFromAPI(o.Key)
	}
	if o.ApiKey != nil {
		ac.APIKey = *o.ApiKey
	}
	return ac
}

func apiKeyListFromAPI(o *openapi.ApiKeyListResponse) *APIKeyList {
	if o == nil {
		return nil
	}
	list := &APIKeyList{}
	for i := range o.Keys {
		list.Keys = append(list.Keys, apiKeyFromAPI(&o.Keys[i]))
	}
	return list
}

// --- Auth adapters ---

func tokenFromAPI(o *openapi.TokenResponse) *Token {
	if o == nil {
		return nil
	}
	t := &Token{
		AccessToken: o.AccessToken,
		TokenType:  o.TokenType,
		ExpiresIn:  o.ExpiresIn,
	}
	if o.RefreshToken != nil {
		t.RefreshToken = *o.RefreshToken
	}
	return t
}

func loginResultFromAPI(o *openapi.LoginResponse) *LoginResult {
	if o == nil {
		return nil
	}
	r := &LoginResult{}
	if o.AccessToken != nil {
		r.AccessToken = *o.AccessToken
	}
	if o.TokenType != nil {
		r.TokenType = *o.TokenType
	}
	if o.ExpiresIn != nil {
		r.ExpiresIn = o.ExpiresIn
	}
	if o.RefreshToken != nil {
		r.RefreshToken = *o.RefreshToken
	}
	if o.MfaRequired != nil {
		r.MFARequired = *o.MfaRequired
	}
	if o.MfaToken != nil {
		r.MFAToken = *o.MfaToken
	}
	return r
}

func userProfileFromAPI(o *openapi.UserProfileResponse) *UserProfile {
	if o == nil {
		return nil
	}
	p := &UserProfile{}
	if o.Id != nil {
		p.ID = *o.Id
	}
	if o.Email != nil {
		p.Email = *o.Email
	}
	if o.DisplayName != nil {
		p.DisplayName = *o.DisplayName
	}
	if o.AuthMethod != nil {
		p.AuthMethod = *o.AuthMethod
	}
	if o.Role != nil {
		p.Role = *o.Role
	}
	if o.EmailVerified != nil {
		p.EmailVerified = *o.EmailVerified
	}
	if o.MarketingEmails != nil {
		p.MarketingEmails = *o.MarketingEmails
	}
	if o.TotpEnabled != nil {
		p.TotpEnabled = *o.TotpEnabled
	}
	if o.CreatedAt != nil {
		p.CreatedAt = o.CreatedAt
	}
	return p
}
