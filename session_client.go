package byo

import (
	"context"
)

// SessionClient handles SessionClient operations
type SessionClient struct {
	client *PropelAuthClient
	// Sub-clients
	Device *DeviceClient
}

func newSessionClient(parent *PropelAuthClient) *SessionClient {
	c := &SessionClient{
		client: parent,
	}
	c.Device = newDeviceClient(parent)
	return c
}

// Create executes the CreateSession command
func (c *SessionClient) Create(ctx context.Context, cmd CreateSessionCommand) (*CreateSessionResponse, error) {
	var resp CreateSessionResponse
	err := c.client.request(ctx, "CreateSession", cmd, &resp, parseCreateSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Validate executes the ValidateSession command
func (c *SessionClient) Validate(ctx context.Context, cmd ValidateSessionCommand) (*ValidateSessionResponse, error) {
	var resp ValidateSessionResponse
	err := c.client.request(ctx, "ValidateSession", cmd, &resp, parseValidateSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateAndRefresh executes the ValidateAndRefreshSession command
func (c *SessionClient) ValidateAndRefresh(ctx context.Context, cmd ValidateAndRefreshSessionCommand) (*ValidateAndRefreshSessionResponse, error) {
	var resp ValidateAndRefreshSessionResponse
	err := c.client.request(ctx, "ValidateAndRefreshSession", cmd, &resp, parseValidateAndRefreshSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAllForUser executes the FetchAllSessionsForUser command
func (c *SessionClient) FetchAllForUser(ctx context.Context, cmd FetchAllSessionsForUserCommand) (*FetchAllSessionsForUserResponse, error) {
	var resp FetchAllSessionsForUserResponse
	err := c.client.request(ctx, "FetchAllSessionsForUser", cmd, &resp, parseFetchAllSessionsForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchAll executes the FetchAllSessions command
func (c *SessionClient) FetchAll(ctx context.Context, cmd FetchAllSessionsCommand) (*FetchAllSessionsResponse, error) {
	var resp FetchAllSessionsResponse
	err := c.client.request(ctx, "FetchAllSessions", cmd, &resp, parseFetchAllSessionsError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchByID executes the FetchSessionById command
func (c *SessionClient) FetchByID(ctx context.Context, cmd FetchSessionByIdCommand) (*SessionInfo, error) {
	var resp SessionInfo
	err := c.client.request(ctx, "FetchSessionById", cmd, &resp, parseFetchSessionByIdError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Update executes the UpdateSession command
func (c *SessionClient) Update(ctx context.Context, cmd UpdateSessionCommand) (*UpdateSessionResponse, error) {
	var resp UpdateSessionResponse
	err := c.client.request(ctx, "UpdateSession", cmd, &resp, parseUpdateSessionError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMany executes the UpdateSessions command
func (c *SessionClient) UpdateMany(ctx context.Context, cmd UpdateSessionsCommand) (*UpdateSessionsResponse, error) {
	var resp UpdateSessionsResponse
	err := c.client.request(ctx, "UpdateSessions", cmd, &resp, parseUpdateSessionsError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateByID executes the InvalidateSessionById command
func (c *SessionClient) InvalidateByID(ctx context.Context, cmd InvalidateSessionByIdCommand) (*InvalidateSessionByIdResponse, error) {
	var resp InvalidateSessionByIdResponse
	err := c.client.request(ctx, "InvalidateSessionById", cmd, &resp, parseInvalidateSessionByIdError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateByToken executes the InvalidateSessionByToken command
func (c *SessionClient) InvalidateByToken(ctx context.Context, cmd InvalidateSessionByTokenCommand) (*InvalidateSessionByTokenResponse, error) {
	var resp InvalidateSessionByTokenResponse
	err := c.client.request(ctx, "InvalidateSessionByToken", cmd, &resp, parseInvalidateSessionByTokenError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateAllForUser executes the InvalidateAllSessionsForUser command
func (c *SessionClient) InvalidateAllForUser(ctx context.Context, cmd InvalidateAllSessionsForUserCommand) (*InvalidateAllSessionsForUserResponse, error) {
	var resp InvalidateAllSessionsForUserResponse
	err := c.client.request(ctx, "InvalidateAllSessionsForUser", cmd, &resp, parseInvalidateAllSessionsForUserError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvalidateAllForUserExceptOne executes the InvalidateAllSessionsForUserExceptOne command
func (c *SessionClient) InvalidateAllForUserExceptOne(ctx context.Context, cmd InvalidateAllSessionsForUserExceptOneCommand) (*InvalidateAllSessionsForUserExceptOneResponse, error) {
	var resp InvalidateAllSessionsForUserExceptOneResponse
	err := c.client.request(ctx, "InvalidateAllSessionsForUserExceptOne", cmd, &resp, parseInvalidateAllSessionsForUserExceptOneError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateStatelessToken executes the CreateStatelessToken command
func (c *SessionClient) CreateStatelessToken(ctx context.Context, cmd CreateStatelessTokenCommand) (*CreateStatelessTokenResponse, error) {
	var resp CreateStatelessTokenResponse
	err := c.client.request(ctx, "CreateStatelessToken", cmd, &resp, parseCreateStatelessTokenError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJwks executes the GetJwks command
func (c *SessionClient) GetJwks(ctx context.Context, cmd GetJwksCommand) (*GetJwksResponse, error) {
	var resp GetJwksResponse
	err := c.client.request(ctx, "GetJwks", cmd, &resp, parseGetJwksError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RotateStatelessTokenKey executes the RotateStatelessTokenKey command
func (c *SessionClient) RotateStatelessTokenKey(ctx context.Context, cmd RotateStatelessTokenKeyCommand) (*RotateStatelessTokenKeyResponse, error) {
	var resp RotateStatelessTokenKeyResponse
	err := c.client.request(ctx, "RotateStatelessTokenKey", cmd, &resp, parseRotateStatelessTokenKeyError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

