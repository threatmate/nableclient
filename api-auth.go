package nableclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/threatmate/restapiclient"
)

// AuthAuthenticateResponse is the response from the authenticate endpoint.
type AuthAuthenticateResponse struct {
	Tokens struct {
		Access  AuthAuthenticateResponseToken `json:"access"`
		Refresh AuthAuthenticateResponseToken `json:"refresh"`
	} `json:"tokens"`
	Refresh  string `json:"refresh"`
	Validate string `json:"validate"`
}

// AuthAuthenticateResponseToken is a token from the authenticate endpoint.
type AuthAuthenticateResponseToken struct {
	Token         string `json:"token"`
	Type          string `json:"type"`
	ExpirySeconds int    `json:"expirySeconds"`
}

func (c *Client) Authenticate(ctx context.Context, apiKey string) error {
	var output AuthAuthenticateResponse
	// Note: Use the raw client directly.
	err := c.client.Do(ctx, http.MethodPost, "/api/auth/authenticate", nil, &output,
		restapiclient.OptionHeader("Authorization", "Bearer "+apiKey),
	)
	if err != nil {
		return fmt.Errorf("authenticate: %w", err)
	}

	c.lock.Lock()
	c.accessToken = output.Tokens.Access.Token
	c.lock.Unlock()

	return nil
}
