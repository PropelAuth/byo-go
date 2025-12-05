package byo

import (
	"context"
)

// SsoManagementClient handles SsoManagementClient operations
type SsoManagementClient struct {
	client *PropelAuthClient
}

func newSsoManagementClient(parent *PropelAuthClient) *SsoManagementClient {
	c := &SsoManagementClient{
		client: parent,
	}
	return c
}

// CreateOIDCClient executes the CreateOidcClient command
func (c *SsoManagementClient) CreateOIDCClient(ctx context.Context, cmd CreateOidcClientCommand) (*CreateOidcClientResponse, error) {
	var resp CreateOidcClientResponse
	err := c.client.request(ctx, "CreateOidcClient", cmd, &resp, parseCreateOidcClientError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchOIDCClient executes the FetchOidcClient command
func (c *SsoManagementClient) FetchOIDCClient(ctx context.Context, cmd FetchOidcClientCommand) (*FetchOidcClientResponse, error) {
	var resp FetchOidcClientResponse
	err := c.client.request(ctx, "FetchOidcClient", cmd, &resp, parseFetchOidcClientError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PatchOIDCClient executes the PatchOidcClient command
func (c *SsoManagementClient) PatchOIDCClient(ctx context.Context, cmd PatchOidcClientCommand) (*PatchOidcClientResponse, error) {
	var resp PatchOidcClientResponse
	err := c.client.request(ctx, "PatchOidcClient", cmd, &resp, parsePatchOidcClientError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOIDCClient executes the DeleteOidcClient command
func (c *SsoManagementClient) DeleteOIDCClient(ctx context.Context, cmd DeleteOidcClientCommand) (*DeleteOidcClientResponse, error) {
	var resp DeleteOidcClientResponse
	err := c.client.request(ctx, "DeleteOidcClient", cmd, &resp, parseDeleteOidcClientError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

