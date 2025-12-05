package byo

import (
	"context"
)

// SsoClient handles SsoClient operations
type SsoClient struct {
	client *PropelAuthClient
	// Sub-clients
	Management *SsoManagementClient
}

func newSsoClient(parent *PropelAuthClient) *SsoClient {
	c := &SsoClient{
		client: parent,
	}
	c.Management = newSsoManagementClient(parent)
	return c
}

// InitiateOIDCLogin executes the InitiateOidcLogin command
func (c *SsoClient) InitiateOIDCLogin(ctx context.Context, cmd InitiateOidcLoginCommand) (*InitiateOidcLoginResponse, error) {
	var resp InitiateOidcLoginResponse
	err := c.client.request(ctx, "InitiateOidcLogin", cmd, &resp, parseInitiateOidcLoginError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CompleteOIDCLogin executes the CompleteOidcLogin command
func (c *SsoClient) CompleteOIDCLogin(ctx context.Context, cmd CompleteOidcLoginCommand) (*CompleteOidcLoginResponse, error) {
	var resp CompleteOidcLoginResponse
	err := c.client.request(ctx, "CompleteOidcLogin", cmd, &resp, parseCompleteOidcLoginError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

