package byo

import (
	"context"
)

// ApiKeyClient handles ApiKeyClient operations
type ApiKeyClient struct {
	client *PropelAuthClient
}

func newApiKeyClient(parent *PropelAuthClient) *ApiKeyClient {
	c := &ApiKeyClient{
		client: parent,
	}
	return c
}

// CreateAPIKey executes the CreateApiKey command
func (c *ApiKeyClient) CreateAPIKey(ctx context.Context, cmd CreateApiKeyCommand) (*CreateApiKeyResponse, error) {
	var resp CreateApiKeyResponse
	err := c.client.request(ctx, "CreateApiKey", cmd, &resp, parseCreateApiKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateAPIKey executes the ValidateApiKey command
func (c *ApiKeyClient) ValidateAPIKey(ctx context.Context, cmd ValidateApiKeyCommand) (*ValidateApiKeyResponse, error) {
	var resp ValidateApiKeyResponse
	err := c.client.request(ctx, "ValidateApiKey", cmd, &resp, parseValidateApiKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeAPIKey executes the RevokeApiKey command
func (c *ApiKeyClient) RevokeAPIKey(ctx context.Context, cmd RevokeApiKeyCommand) (*RevokeApiKeyResponse, error) {
	var resp RevokeApiKeyResponse
	err := c.client.request(ctx, "RevokeApiKey", cmd, &resp, parseRevokeApiKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PatchAPIKey executes the PatchApiKey command
func (c *ApiKeyClient) PatchAPIKey(ctx context.Context, cmd PatchApiKeyCommand) (*PatchApiKeyResponse, error) {
	var resp PatchApiKeyResponse
	err := c.client.request(ctx, "PatchApiKey", cmd, &resp, parsePatchApiKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAPIKeys executes the GetApiKeys command
func (c *ApiKeyClient) GetAPIKeys(ctx context.Context, cmd GetApiKeysCommand) (*GetApiKeysResponse, error) {
	var resp GetApiKeysResponse
	err := c.client.request(ctx, "GetApiKeys", cmd, &resp, parseGetApiKeysError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

