package byo

import (
	"context"
)

// PasskeyClient handles PasskeyClient operations
type PasskeyClient struct {
	client *PropelAuthClient
}

func newPasskeyClient(parent *PropelAuthClient) *PasskeyClient {
	c := &PasskeyClient{
		client: parent,
	}
	return c
}

// StartRegistration executes the StartPasskeyRegistration command
func (c *PasskeyClient) StartRegistration(ctx context.Context, cmd StartPasskeyRegistrationCommand) (*StartPasskeyRegistrationResponse, error) {
	var resp StartPasskeyRegistrationResponse
	err := c.client.request(ctx, "StartPasskeyRegistration", cmd, &resp, parseStartPasskeyRegistrationError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FinishRegistration executes the FinishPasskeyRegistration command
func (c *PasskeyClient) FinishRegistration(ctx context.Context, cmd FinishPasskeyRegistrationCommand) (*FinishPasskeyRegistrationResponse, error) {
	var resp FinishPasskeyRegistrationResponse
	err := c.client.request(ctx, "FinishPasskeyRegistration", cmd, &resp, parseFinishPasskeyRegistrationError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartAuthentication executes the StartPasskeyAuthentication command
func (c *PasskeyClient) StartAuthentication(ctx context.Context, cmd StartPasskeyAuthenticationCommand) (*StartPasskeyAuthenticationResponse, error) {
	var resp StartPasskeyAuthenticationResponse
	err := c.client.request(ctx, "StartPasskeyAuthentication", cmd, &resp, parseStartPasskeyAuthenticationError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FinishAuthentication executes the FinishPasskeyAuthentication command
func (c *PasskeyClient) FinishAuthentication(ctx context.Context, cmd FinishPasskeyAuthenticationCommand) (*FinishPasskeyAuthenticationResponse, error) {
	var resp FinishPasskeyAuthenticationResponse
	err := c.client.request(ctx, "FinishPasskeyAuthentication", cmd, &resp, parseFinishPasskeyAuthenticationError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAllPasskeysForUser executes the FetchAllPasskeysForUser command
func (c *PasskeyClient) FetchAllPasskeysForUser(ctx context.Context, cmd FetchAllPasskeysForUserCommand) (*FetchAllPasskeysForUserResponse, error) {
	var resp FetchAllPasskeysForUserResponse
	err := c.client.request(ctx, "FetchAllPasskeysForUser", cmd, &resp, parseFetchAllPasskeysForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeregisterPasskey executes the DeregisterPasskey command
func (c *PasskeyClient) DeregisterPasskey(ctx context.Context, cmd DeregisterPasskeyCommand) (*DeregisterPasskeyResponse, error) {
	var resp DeregisterPasskeyResponse
	err := c.client.request(ctx, "DeregisterPasskey", cmd, &resp, parseDeregisterPasskeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeregisterAllPasskeysForUser executes the DeregisterAllPasskeysForUser command
func (c *PasskeyClient) DeregisterAllPasskeysForUser(ctx context.Context, cmd DeregisterAllPasskeysForUserCommand) (*DeregisterAllPasskeysForUserResponse, error) {
	var resp DeregisterAllPasskeysForUserResponse
	err := c.client.request(ctx, "DeregisterAllPasskeysForUser", cmd, &resp, parseDeregisterAllPasskeysForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

