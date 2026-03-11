package nablesimulator

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/threatmate/nableclient"
)

type Universe struct {
	APIUsers    []*APIUser
	ServiceOrgs []*nableclient.ServiceOrg
	Users       []*nableclient.OrgUnitUser
	Customers   []*nableclient.Customer
	Devices     []*nableclient.Device
}

type APIUser struct {
	Username      string
	APIKey        string
	accessTokens  []APIUserToken
	refreshTokens []APIUserToken
}

type APIUserToken struct {
	Token     string
	Type      string
	ExpiresAt time.Time
}

func (u *APIUser) AuthenticateAPIKey(apiKey string) (output nableclient.PostAuthAuthenticateResponse, err error) {
	if apiKey != u.APIKey {
		return output, fmt.Errorf("invalid API key")
	}

	accessToken := APIUserToken{
		Token:     uuid.New().String(),
		Type:      "Bearer",
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	u.accessTokens = append(u.accessTokens, accessToken)

	refreshToken := APIUserToken{
		Token:     uuid.New().String(),
		Type:      "Bearer",
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	u.refreshTokens = append(u.refreshTokens, refreshToken)

	output.Tokens.Access = nableclient.PostAuthAuthenticateResponseToken{
		Token:         accessToken.Token,
		Type:          accessToken.Type,
		ExpirySeconds: int(time.Until(accessToken.ExpiresAt).Seconds()),
	}
	output.Tokens.Refresh = nableclient.PostAuthAuthenticateResponseToken{
		Token:         refreshToken.Token,
		Type:          refreshToken.Type,
		ExpirySeconds: int(time.Until(refreshToken.ExpiresAt).Seconds()),
	}
	return output, nil
}
