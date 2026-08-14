package oneclaw

import "context"

// PolicyBackendService manages Cedar/OPA enforcement settings.
type PolicyBackendService struct {
	client *Client
}

// GetSettings returns org policy backend configuration.
func (s *PolicyBackendService) GetSettings(ctx context.Context) (*PolicyBackendSettings, error) {
	var resp PolicyBackendSettings
	err := s.client.doJSON(ctx, "GET", "/v1/org/settings/policy-backend", nil, &resp)
	return &resp, err
}

// UpdateSettings patches org policy backend configuration.
func (s *PolicyBackendService) UpdateSettings(ctx context.Context, req UpdatePolicyBackendSettingsRequest) (*PolicyBackendSettings, error) {
	var resp PolicyBackendSettings
	err := s.client.doJSON(ctx, "PATCH", "/v1/org/settings/policy-backend", req, &resp)
	return &resp, err
}

// GetShadowReport returns shadow mode divergence statistics.
func (s *PolicyBackendService) GetShadowReport(ctx context.Context) (*PolicyShadowReport, error) {
	var resp PolicyShadowReport
	err := s.client.doJSON(ctx, "GET", "/v1/org/policy-shadow-report", nil, &resp)
	return &resp, err
}
