package byo

import (
	"context"
)

// ImpersonationClient handles ImpersonationClient operations
type ImpersonationClient struct {
	client *PropelAuthClient
}

func newImpersonationClient(parent *PropelAuthClient) *ImpersonationClient {
	c := &ImpersonationClient{
		client: parent,
	}
	return c
}

// Create executes the CreateImpersonationSession command
func (c *ImpersonationClient) Create(ctx context.Context, cmd CreateImpersonationSessionCommand) (*CreateImpersonationSessionResponse, error) {
	var resp CreateImpersonationSessionResponse
	err := c.client.request(ctx, "CreateImpersonationSession", cmd, &resp, parseCreateImpersonationSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Validate executes the ValidateImpersonationSession command
func (c *ImpersonationClient) Validate(ctx context.Context, cmd ValidateImpersonationSessionCommand) (*ValidateImpersonationSessionResponse, error) {
	var resp ValidateImpersonationSessionResponse
	err := c.client.request(ctx, "ValidateImpersonationSession", cmd, &resp, parseValidateImpersonationSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchByID executes the FetchImpersonationSessionById command
func (c *ImpersonationClient) FetchByID(ctx context.Context, cmd FetchImpersonationSessionByIdCommand) (*ImpersonationSessionInfo, error) {
	var resp ImpersonationSessionInfo
	err := c.client.request(ctx, "FetchImpersonationSessionById", cmd, &resp, parseFetchImpersonationSessionByIdError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAllForEmployee executes the FetchAllImpersonationSessionsForEmployee command
func (c *ImpersonationClient) FetchAllForEmployee(ctx context.Context, cmd FetchAllImpersonationSessionsForEmployeeCommand) (*FetchAllImpersonationSessionsForEmployeeResponse, error) {
	var resp FetchAllImpersonationSessionsForEmployeeResponse
	err := c.client.request(ctx, "FetchAllImpersonationSessionsForEmployee", cmd, &resp, parseFetchAllImpersonationSessionsForEmployeeError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAllForUser executes the FetchAllImpersonationSessionsForUser command
func (c *ImpersonationClient) FetchAllForUser(ctx context.Context, cmd FetchAllImpersonationSessionsForUserCommand) (*FetchAllImpersonationSessionsForUserResponse, error) {
	var resp FetchAllImpersonationSessionsForUserResponse
	err := c.client.request(ctx, "FetchAllImpersonationSessionsForUser", cmd, &resp, parseFetchAllImpersonationSessionsForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAllActive executes the FetchAllActiveImpersonationSessions command
func (c *ImpersonationClient) FetchAllActive(ctx context.Context, cmd FetchAllActiveImpersonationSessionsCommand) (*FetchAllActiveImpersonationSessionsResponse, error) {
	var resp FetchAllActiveImpersonationSessionsResponse
	err := c.client.request(ctx, "FetchAllActiveImpersonationSessions", cmd, &resp, parseFetchAllActiveImpersonationSessionsError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateByID executes the InvalidateImpersonationSessionById command
func (c *ImpersonationClient) InvalidateByID(ctx context.Context, cmd InvalidateImpersonationSessionByIdCommand) (*InvalidateImpersonationSessionByIdResponse, error) {
	var resp InvalidateImpersonationSessionByIdResponse
	err := c.client.request(ctx, "InvalidateImpersonationSessionById", cmd, &resp, parseInvalidateImpersonationSessionByIdError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateByToken executes the InvalidateImpersonationSessionByToken command
func (c *ImpersonationClient) InvalidateByToken(ctx context.Context, cmd InvalidateImpersonationSessionByTokenCommand) (*InvalidateImpersonationSessionByTokenResponse, error) {
	var resp InvalidateImpersonationSessionByTokenResponse
	err := c.client.request(ctx, "InvalidateImpersonationSessionByToken", cmd, &resp, parseInvalidateImpersonationSessionByTokenError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateAllForEmployee executes the InvalidateAllImpersonationSessionsForEmployee command
func (c *ImpersonationClient) InvalidateAllForEmployee(ctx context.Context, cmd InvalidateAllImpersonationSessionsForEmployeeCommand) (*InvalidateAllImpersonationSessionsForEmployeeResponse, error) {
	var resp InvalidateAllImpersonationSessionsForEmployeeResponse
	err := c.client.request(ctx, "InvalidateAllImpersonationSessionsForEmployee", cmd, &resp, parseInvalidateAllImpersonationSessionsForEmployeeError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateAllForUser executes the InvalidateAllImpersonationSessionsForUser command
func (c *ImpersonationClient) InvalidateAllForUser(ctx context.Context, cmd InvalidateAllImpersonationSessionsForUserCommand) (*InvalidateAllImpersonationSessionsForUserResponse, error) {
	var resp InvalidateAllImpersonationSessionsForUserResponse
	err := c.client.request(ctx, "InvalidateAllImpersonationSessionsForUser", cmd, &resp, parseInvalidateAllImpersonationSessionsForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

