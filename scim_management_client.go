package byo

import (
	"context"
)

// ScimManagementClient handles ScimManagementClient operations
type ScimManagementClient struct {
	client *PropelAuthClient
}

func newScimManagementClient(parent *PropelAuthClient) *ScimManagementClient {
	c := &ScimManagementClient{
		client: parent,
	}
	return c
}

// GetScimUsers executes the GetScimUsers command
func (c *ScimManagementClient) GetScimUsers(ctx context.Context, cmd GetScimUsersCommand) (*GetScimUsersResponse, error) {
	var resp GetScimUsersResponse
	err := c.client.request(ctx, "GetScimUsers", cmd, &resp, parseGetScimUsersError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateScimConnection executes the CreateScimConnection command
func (c *ScimManagementClient) CreateScimConnection(ctx context.Context, cmd CreateScimConnectionCommand) (*CreateScimConnectionResponse, error) {
	var resp CreateScimConnectionResponse
	err := c.client.request(ctx, "CreateScimConnection", cmd, &resp, parseCreateScimConnectionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchScimConnection executes the FetchScimConnection command
func (c *ScimManagementClient) FetchScimConnection(ctx context.Context, cmd FetchScimConnectionCommand) (*FetchScimConnectionResponse, error) {
	var resp FetchScimConnectionResponse
	err := c.client.request(ctx, "FetchScimConnection", cmd, &resp, parseFetchScimConnectionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PatchScimConnection executes the PatchScimConnection command
func (c *ScimManagementClient) PatchScimConnection(ctx context.Context, cmd PatchScimConnectionCommand) (*PatchScimConnectionResponse, error) {
	var resp PatchScimConnectionResponse
	err := c.client.request(ctx, "PatchScimConnection", cmd, &resp, parsePatchScimConnectionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetScimAPIKey executes the ResetScimApiKey command
func (c *ScimManagementClient) ResetScimAPIKey(ctx context.Context, cmd ResetScimApiKeyCommand) (*ResetScimApiKeyResponse, error) {
	var resp ResetScimApiKeyResponse
	err := c.client.request(ctx, "ResetScimApiKey", cmd, &resp, parseResetScimApiKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteScimConnection executes the DeleteScimConnection command
func (c *ScimManagementClient) DeleteScimConnection(ctx context.Context, cmd DeleteScimConnectionCommand) (*DeleteScimConnectionResponse, error) {
	var resp DeleteScimConnectionResponse
	err := c.client.request(ctx, "DeleteScimConnection", cmd, &resp, parseDeleteScimConnectionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

