package oneclaw

import (
	"context"
	"fmt"
	"net/url"
)

// PortfolioService provides portfolio aggregation operations.
type PortfolioService struct {
	client *Client
}

// PortfolioParams configures the portfolio query.
type PortfolioParams struct {
	Chains        string
	IncludeTokens bool
}

// Get returns the aggregated portfolio view across all wallets and signing keys.
func (s *PortfolioService) Get(ctx context.Context, params *PortfolioParams) (*PortfolioResponse, error) {
	path := "/v1/portfolio"
	if params != nil {
		q := url.Values{}
		if params.Chains != "" {
			q.Set("chains", params.Chains)
		}
		if params.IncludeTokens {
			q.Set("include_tokens", "true")
		}
		if qs := q.Encode(); qs != "" {
			path = fmt.Sprintf("%s?%s", path, qs)
		}
	}
	var resp PortfolioResponse
	err := s.client.doJSON(ctx, "GET", path, nil, &resp)
	return &resp, err
}
