package byo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// HTTPDoer interface allows for testing
type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

// PropelAuthClient is the main client for PropelAuth API
type PropelAuthClient struct {
	url            string
	integrationKey string
	httpClient     HTTPDoer

	// Sub-clients
	Passkeys *PasskeyClient
	Session *SessionClient
	Scim *ScimClient
	Impersonation *ImpersonationClient
	Sso *SsoClient
}

// NewClient creates a new PropelAuth client
func NewClient(url, integrationKey string, opts ...ClientOption) (*PropelAuthClient, error) {
	if url == "" {
		return nil, fmt.Errorf("url is required")
	}
	if integrationKey == "" {
		return nil, fmt.Errorf("integration key is required")
	}

	c := &PropelAuthClient{
		url:            strings.TrimSuffix(url, "/"),
		integrationKey: integrationKey,
		httpClient:     &http.Client{Timeout: 30 * time.Second},
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	// Initialize sub-clients
	c.Passkeys = newPasskeyClient(c)
	c.Session = newSessionClient(c)
	c.Scim = newScimClient(c)
	c.Impersonation = newImpersonationClient(c)
	c.Sso = newSsoClient(c)

	return c, nil
}

// ClientOption is a functional option for configuring the client
type ClientOption func(*PropelAuthClient)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient HTTPDoer) ClientOption {
	return func(c *PropelAuthClient) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *PropelAuthClient) {
		if client, ok := c.httpClient.(*http.Client); ok {
			client.Timeout = timeout
		}
	}
}

// Ping executes the Ping command
func (c *PropelAuthClient) Ping(ctx context.Context, cmd PingCommand) (*PingResponse, error) {
	var resp PingResponse
	err := c.request(ctx, "Ping", cmd, &resp, parsePingError)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ErrorParser is a function type for parsing error responses
type ErrorParser func([]byte) error

// request is the internal method for making HTTP requests
func (c *PropelAuthClient) request(ctx context.Context, commandName string, data interface{}, response interface{}, parseError ErrorParser) error {
	payload := map[string]interface{}{
		"command": commandName,
	}
	if data != nil {
		payload["data"] = data
	}

	body, err := json.Marshal(payload)
	if err != nil {
		errJSON := fmt.Sprintf(`{"type":"UnexpectedError","details":"failed to marshal request: %s"}`, err.Error())
		return parseError([]byte(errJSON))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+"/run-command", bytes.NewReader(body))
	if err != nil {
		errJSON := fmt.Sprintf(`{"type":"UnexpectedError","details":"failed to create request: %s"}`, err.Error())
		return parseError([]byte(errJSON))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.integrationKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Check for context errors
		if errors.Is(err, context.DeadlineExceeded) {
			errJSON := `{"type":"UnexpectedError","details":"request timeout"}`
			return parseError([]byte(errJSON))
		}
		if errors.Is(err, context.Canceled) {
			errJSON := `{"type":"UnexpectedError","details":"request canceled"}`
			return parseError([]byte(errJSON))
		}
		// Network error
		errJSON := fmt.Sprintf(`{"type":"UnexpectedError","details":"network error: %s"}`, err.Error())
		return parseError([]byte(errJSON))
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errJSON := fmt.Sprintf(`{"type":"UnexpectedError","details":"failed to read response: %s"}`, err.Error())
		return parseError([]byte(errJSON))
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.Unmarshal(bodyBytes, response); err != nil {
			errJSON := fmt.Sprintf(`{"type":"UnexpectedError","details":"failed to parse success response: %s"}`, err.Error())
			return parseError([]byte(errJSON))
		}
		return nil
	}

	// Parse error response using the provided error parser
	return parseError(bodyBytes)
}
