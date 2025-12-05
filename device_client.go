package byo

import (
	"context"
)

// DeviceClient handles DeviceClient operations
type DeviceClient struct {
	client *PropelAuthClient
}

func newDeviceClient(parent *PropelAuthClient) *DeviceClient {
	c := &DeviceClient{
		client: parent,
	}
	return c
}

// CreateChallenge executes the CreateDeviceChallenge command
func (c *DeviceClient) CreateChallenge(ctx context.Context, cmd CreateDeviceChallengeCommand) (*CreateDeviceChallengeResponse, error) {
	var resp CreateDeviceChallengeResponse
	err := c.client.request(ctx, "CreateDeviceChallenge", cmd, &resp, parseCreateDeviceChallengeError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Register executes the RegisterDevice command
func (c *DeviceClient) Register(ctx context.Context, cmd RegisterDeviceCommand) (*RegisterDeviceResponse, error) {
	var resp RegisterDeviceResponse
	err := c.client.request(ctx, "RegisterDevice", cmd, &resp, parseRegisterDeviceError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

