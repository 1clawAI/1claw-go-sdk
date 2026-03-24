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

// --- Access/Policy adapters ---

func policyFromAPI(o *openapi.PolicyResponse) Policy {
	if o == nil {
		return Policy{}
	}
	p := Policy{
		ID:                o.Id,
		VaultID:           o.VaultId,
		SecretPathPattern: o.SecretPathPattern,
		PrincipalType:     o.PrincipalType,
		PrincipalID:       o.PrincipalId,
		Permissions:       o.Permissions,
		CreatedAt:         o.CreatedAt,
	}
	if o.Conditions != nil {
		p.Conditions = o.Conditions
	}
	if o.ExpiresAt != nil {
		p.ExpiresAt = o.ExpiresAt
	}
	if o.CreatedBy != nil {
		p.CreatedBy = *o.CreatedBy
	}
	if o.CreatedByType != nil {
		p.CreatedByType = *o.CreatedByType
	}
	return p
}

func policyListFromAPI(o *openapi.PolicyListResponse) *PolicyList {
	if o == nil {
		return nil
	}
	list := &PolicyList{}
	for i := range o.Policies {
		list.Policies = append(list.Policies, policyFromAPI(&o.Policies[i]))
	}
	return list
}

func createPolicyParamsToAPI(p CreatePolicyParams) openapi.CreatePolicyRequest {
	req := openapi.CreatePolicyRequest{
		SecretPathPattern: p.SecretPathPattern,
		PrincipalType:     p.PrincipalType,
		PrincipalId:       p.PrincipalID,
		Permissions:       p.Permissions,
	}
	if p.Conditions != nil {
		req.Conditions = p.Conditions
	}
	if p.ExpiresAt != nil {
		req.ExpiresAt = p.ExpiresAt
	}
	return req
}

func updatePolicyParamsToAPI(p UpdatePolicyParams) openapi.UpdatePolicyRequest {
	var req openapi.UpdatePolicyRequest
	req.Permissions = p.Permissions
	req.Conditions = p.Conditions
	req.ExpiresAt = p.ExpiresAt
	return req
}

// --- Sharing adapters ---

func shareFromAPI(o *openapi.ShareResponse) Share {
	if o == nil {
		return Share{}
	}
	s := Share{}
	if o.Id != nil {
		s.ID = *o.Id
	}
	if o.ShareUrl != nil {
		s.ShareURL = *o.ShareUrl
	}
	if o.RecipientType != nil {
		s.RecipientType = *o.RecipientType
	}
	if o.RecipientEmail != nil {
		s.RecipientEmail = *o.RecipientEmail
	}
	if o.ExpiresAt != nil {
		s.ExpiresAt = o.ExpiresAt
	}
	if o.MaxAccessCount != nil {
		v := *o.MaxAccessCount
		s.MaxAccessCount = &v
	}
	return s
}

func shareListItemFromAPI(o *openapi.ShareListItem) Share {
	if o == nil {
		return Share{}
	}
	s := Share{}
	if o.Id != nil {
		s.ID = *o.Id
	}
	if o.RecipientType != nil {
		s.RecipientType = *o.RecipientType
	}
	if o.RecipientEmail != nil {
		s.RecipientEmail = *o.RecipientEmail
	}
	if o.ExpiresAt != nil {
		s.ExpiresAt = o.ExpiresAt
	}
	if o.MaxAccessCount != nil {
		v := *o.MaxAccessCount
		s.MaxAccessCount = &v
	}
	return s
}

func shareListFromAPI(o *openapi.ShareListResponse) *ShareList {
	if o == nil {
		return nil
	}
	list := &ShareList{}
	for i := range o.Shares {
		list.Shares = append(list.Shares, shareListItemFromAPI(&o.Shares[i]))
	}
	return list
}

func createShareParamsToAPI(p CreateShareParams) openapi.CreateShareRequest {
	req := openapi.CreateShareRequest{
		RecipientType: p.RecipientType,
		ExpiresAt:     p.ExpiresAt,
	}
	if p.RecipientID != "" {
		req.RecipientId = &p.RecipientID
	}
	if p.Email != "" {
		req.Email = &p.Email
	}
	if len(p.Permissions) > 0 {
		req.Permissions = p.Permissions
	}
	if p.MaxAccessCount != nil {
		req.MaxAccessCount = p.MaxAccessCount
	}
	if p.Passphrase != "" {
		req.Passphrase = &p.Passphrase
	}
	if len(p.IPAllowlist) > 0 {
		req.IpAllowlist = p.IPAllowlist
	}
	return req
}

// --- Org adapters ---

func orgMemberFromAPI(o *openapi.OrgMemberResponse) OrgMember {
	if o == nil {
		return OrgMember{}
	}
	m := OrgMember{}
	if o.Id != nil {
		m.ID = *o.Id
	}
	if o.Email != nil {
		m.Email = *o.Email
	}
	if o.DisplayName != nil {
		m.DisplayName = *o.DisplayName
	}
	if o.Role != nil {
		m.Role = *o.Role
	}
	if o.AuthMethod != nil {
		m.AuthMethod = *o.AuthMethod
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt
	}
	return m
}

func orgMemberListFromAPI(o *openapi.OrgMemberListResponse) *OrgMemberList {
	if o == nil {
		return nil
	}
	list := &OrgMemberList{}
	for i := range o.Members {
		list.Members = append(list.Members, orgMemberFromAPI(&o.Members[i]))
	}
	return list
}

func inviteMemberParamsToAPI(p InviteMemberParams) openapi.InviteMemberRequest {
	req := openapi.InviteMemberRequest{Email: p.Email}
	if p.Role != "" {
		req.Role = &p.Role
	}
	return req
}

func updateMemberRoleParamsToAPI(p UpdateMemberRoleParams) openapi.UpdateMemberRoleRequest {
	return openapi.UpdateMemberRoleRequest{Role: p.Role}
}

func inviteMemberResultFromAPI(o *openapi.InviteMemberResponse) *InviteMemberResult {
	if o == nil {
		return nil
	}
	r := &InviteMemberResult{}
	if o.Message != nil {
		r.Message = *o.Message
	}
	if o.Email != nil {
		r.Email = *o.Email
	}
	return r
}

// --- Chain adapters ---

func chainFromAPI(o *openapi.ChainResponse) Chain {
	if o == nil {
		return Chain{}
	}
	c := Chain{}
	if o.Id != nil {
		c.ID = *o.Id
	}
	if o.Name != nil {
		c.Name = *o.Name
	}
	if o.DisplayName != nil {
		c.DisplayName = *o.DisplayName
	}
	if o.ChainId != nil {
		c.ChainID = o.ChainId
	}
	if o.RpcUrl != nil {
		c.RPCURL = *o.RpcUrl
	}
	if o.WsUrl != nil {
		c.WSURL = *o.WsUrl
	}
	if o.ExplorerUrl != nil {
		c.ExplorerURL = *o.ExplorerUrl
	}
	if o.NativeCurrency != nil {
		c.NativeCurrency = *o.NativeCurrency
	}
	if o.IsTestnet != nil {
		c.IsTestnet = *o.IsTestnet
	}
	if o.IsEnabled != nil {
		c.IsEnabled = *o.IsEnabled
	}
	if o.CreatedAt != nil {
		c.CreatedAt = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		c.UpdatedAt = o.UpdatedAt
	}
	return c
}

func chainListFromAPI(o *openapi.ChainListResponse) *ChainList {
	if o == nil {
		return nil
	}
	list := &ChainList{}
	for i := range o.Chains {
		list.Chains = append(list.Chains, chainFromAPI(&o.Chains[i]))
	}
	return list
}

// --- Billing adapters ---

func subscriptionFromAPI(o *openapi.SubscriptionResponse) *Subscription {
	if o == nil {
		return nil
	}
	s := &Subscription{}
	if o.Tier != nil {
		s.Tier = *o.Tier
	}
	if o.Interval != nil {
		s.Interval = *o.Interval
	}
	if o.PeriodEnd != nil {
		s.PeriodEnd = o.PeriodEnd
	}
	if o.Status != nil {
		s.Status = *o.Status
	}
	if o.CreditBalanceCents != nil {
		s.CreditBalanceCents = *o.CreditBalanceCents
	}
	if o.CreditBalanceUsd != nil {
		s.CreditBalanceUsd = *o.CreditBalanceUsd
	}
	if o.OverageMethod != nil {
		s.OverageMethod = *o.OverageMethod
	}
	if o.Usage != nil {
		s.Usage = o.Usage
	}
	return s
}

func creditBalanceFromAPI(o *openapi.CreditBalanceResponse) *CreditBalance {
	if o == nil {
		return nil
	}
	c := &CreditBalance{}
	if o.BalanceCents != nil {
		c.BalanceCents = *o.BalanceCents
	}
	if o.BalanceUsd != nil {
		c.BalanceUsd = *o.BalanceUsd
	}
	if o.ExpiringWithin90Days != nil {
		c.ExpiringWithin90Days = o.ExpiringWithin90Days
	}
	return c
}

// --- Audit adapters ---

func auditEventFromAPI(o *openapi.AuditEvent) AuditEvent {
	if o == nil {
		return AuditEvent{}
	}
	e := AuditEvent{}
	if o.Id != nil {
		e.ID = *o.Id
	}
	if o.Action != nil {
		e.Action = *o.Action
	}
	if o.ActorId != nil {
		e.ActorID = *o.ActorId
	}
	if o.ActorType != nil {
		e.ActorType = *o.ActorType
	}
	if o.ResourceType != nil {
		e.ResourceType = *o.ResourceType
	}
	if o.ResourceId != nil {
		e.ResourceID = *o.ResourceId
	}
	if o.OrgId != nil {
		e.OrgID = *o.OrgId
	}
	if o.Details != nil {
		e.Details = o.Details
	}
	if o.IpAddress != nil {
		e.IPAddress = *o.IpAddress
	}
	if o.CreatedAt != nil {
		e.CreatedAt = o.CreatedAt
	}
	return e
}

func auditEventsFromAPI(o *openapi.AuditEventsResponse) *AuditEvents {
	if o == nil {
		return nil
	}
	a := &AuditEvents{}
	for i := range o.Events {
		a.Events = append(a.Events, auditEventFromAPI(&o.Events[i]))
	}
	if o.Count != nil {
		a.Count = *o.Count
	}
	return a
}

// --- X402 adapters ---

func paymentRequirementFromAPI(o *openapi.PaymentRequirement) *PaymentRequirement {
	if o == nil {
		return nil
	}
	pr := &PaymentRequirement{}
	if o.X402Version != nil {
		v := *o.X402Version
		pr.X402Version = &v
	}
	if o.Description != nil {
		pr.Description = *o.Description
	}
	if len(o.Accepts) > 0 {
		pr.Accepts = make([]PaymentRequirementAccept, len(o.Accepts))
		for i := range o.Accepts {
			pr.Accepts[i] = PaymentRequirementAccept{
				Scheme:                  ptrStr(o.Accepts[i].Scheme),
				Network:                 ptrStr(o.Accepts[i].Network),
				PayTo:                   ptrStr(o.Accepts[i].PayTo),
				Price:                   ptrStr(o.Accepts[i].Price),
				RequiredDeadlineSeconds: o.Accepts[i].RequiredDeadlineSeconds,
			}
		}
	}
	return pr
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// --- Treasury adapters ---

func treasurySignerFromAPI(o *openapi.TreasurySignerResponse) TreasurySigner {
	if o == nil {
		return TreasurySigner{}
	}
	s := TreasurySigner{}
	if o.Id != nil {
		s.ID = *o.Id
	}
	if o.SignerType != nil {
		s.SignerType = *o.SignerType
	}
	if o.SignerId != nil {
		s.SignerID = *o.SignerId
	}
	if o.SignerAddress != nil {
		s.SignerAddress = *o.SignerAddress
	}
	if o.AddedAt != nil {
		s.AddedAt = o.AddedAt
	}
	return s
}

func treasuryFromAPI(o *openapi.TreasuryResponse) *Treasury {
	if o == nil {
		return nil
	}
	t := &Treasury{}
	if o.Id != nil {
		t.ID = *o.Id
	}
	if o.Name != nil {
		t.Name = *o.Name
	}
	if o.SafeAddress != nil {
		t.SafeAddress = *o.SafeAddress
	}
	if o.Chain != nil {
		t.Chain = *o.Chain
	}
	if o.ChainId != nil {
		v := *o.ChainId
		t.ChainID = &v
	}
	if o.Threshold != nil {
		v := *o.Threshold
		t.Threshold = &v
	}
	if o.CreatedBy != nil {
		t.CreatedBy = *o.CreatedBy
	}
	for i := range o.Signers {
		t.Signers = append(t.Signers, treasurySignerFromAPI(&o.Signers[i]))
	}
	if o.CreatedAt != nil {
		t.CreatedAt = o.CreatedAt
	}
	return t
}

func treasuryListFromAPI(o *openapi.ListTreasuries200Response) *TreasuryList {
	if o == nil {
		return nil
	}
	list := &TreasuryList{}
	for i := range o.Treasuries {
		if tr := treasuryFromAPI(&o.Treasuries[i]); tr != nil {
			list.Treasuries = append(list.Treasuries, *tr)
		}
	}
	return list
}

func treasuryAccessRequestFromAPI(o *openapi.AccessRequestResponse) TreasuryAccessRequest {
	if o == nil {
		return TreasuryAccessRequest{}
	}
	a := TreasuryAccessRequest{}
	if o.Id != nil {
		a.ID = *o.Id
	}
	if o.TreasuryId != nil {
		a.TreasuryID = *o.TreasuryId
	}
	if o.AgentId != nil {
		a.AgentID = *o.AgentId
	}
	if o.Status != nil {
		a.Status = *o.Status
	}
	if o.Reason != nil {
		a.Reason = *o.Reason
	}
	if o.RequestedAt != nil {
		a.RequestedAt = o.RequestedAt
	}
	if o.ResolvedBy != nil {
		a.ResolvedBy = *o.ResolvedBy
	}
	if o.ResolvedAt != nil {
		a.ResolvedAt = o.ResolvedAt
	}
	return a
}

func treasuryAccessRequestListFromAPI(o *openapi.ListTreasuryAccessRequests200Response) *TreasuryAccessRequestList {
	if o == nil {
		return nil
	}
	list := &TreasuryAccessRequestList{}
	for i := range o.Requests {
		list.Requests = append(list.Requests, treasuryAccessRequestFromAPI(&o.Requests[i]))
	}
	return list
}
