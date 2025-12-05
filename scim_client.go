package byo

import (
	"context"
	"encoding/json"
)

// ScimClient handles ScimClient operations
type ScimClient struct {
	client *PropelAuthClient
	// Sub-clients
	Management *ScimManagementClient
}

func newScimClient(parent *PropelAuthClient) *ScimClient {
	c := &ScimClient{
		client: parent,
	}
	c.Management = newScimManagementClient(parent)
	return c
}

// ScimRequest executes the ScimRequest command
func (c *ScimClient) ScimRequest(ctx context.Context, cmd ScimRequestCommand) (ScimResponse, error) {
	var rawResp json.RawMessage
	err := c.client.request(ctx, "ScimRequest", cmd, &rawResp, parseScimClientFacingError)
	if err != nil {
		return nil, err
	}
	return UnmarshalScimResponse(rawResp)
}

// LinkScimUser executes the LinkScimUser command
func (c *ScimClient) LinkScimUser(ctx context.Context, cmd LinkScimUserCommand) (*ScimResponseCompleted, error) {
	var resp ScimResponseCompleted
	err := c.client.request(ctx, "LinkScimUser", cmd, &resp, parseLinkScimUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CommitScimUserChange executes the CommitScimUserChange command
func (c *ScimClient) CommitScimUserChange(ctx context.Context, cmd CommitScimUserChangeCommand) (*ScimResponseCompleted, error) {
	var resp ScimResponseCompleted
	err := c.client.request(ctx, "CommitScimUserChange", cmd, &resp, parseCommitScimUserChangeError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetScimUser executes the GetScimUser command
func (c *ScimClient) GetScimUser(ctx context.Context, cmd GetScimUserCommand) (*GetScimUserResponse, error) {
	var resp GetScimUserResponse
	err := c.client.request(ctx, "GetScimUser", cmd, &resp, parseGetScimUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

