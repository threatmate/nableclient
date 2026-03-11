package simulator

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/threatmate/ncentralclient"
)

type Universe struct {
	APIUsers    []*APIUser
	ServiceOrgs []*ncentralclient.ServiceOrg
	OrgUnits    []*OrgUnit
	Customers   []*ncentralclient.Customer
	Devices     []*Device
}

type OrgUnit struct {
	OrgUnitID string
	Users     []*ncentralclient.OrgUnitUser
}

type Device struct {
	Device *ncentralclient.Device
	Asset  *ncentralclient.DeviceAsset
}

type APIUser struct {
	Username      string
	APIKey        string
	lock          sync.Mutex
	accessTokens  []APIUserToken
	refreshTokens []APIUserToken
}

type APIUserToken struct {
	Token     string
	Type      string
	ExpiresAt time.Time
}

func (u *APIUser) AuthenticateAPIKey(apiKey string) (output ncentralclient.PostAuthAuthenticateResponse, err error) {
	if apiKey != u.APIKey {
		return output, fmt.Errorf("invalid API key")
	}

	u.lock.Lock()
	defer u.lock.Unlock()

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

	output.Tokens.Access = ncentralclient.PostAuthAuthenticateResponseToken{
		Token:         accessToken.Token,
		Type:          accessToken.Type,
		ExpirySeconds: int(time.Until(accessToken.ExpiresAt).Seconds()),
	}
	output.Tokens.Refresh = ncentralclient.PostAuthAuthenticateResponseToken{
		Token:         refreshToken.Token,
		Type:          refreshToken.Type,
		ExpirySeconds: int(time.Until(refreshToken.ExpiresAt).Seconds()),
	}
	return output, nil
}
