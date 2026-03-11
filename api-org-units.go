package ncentralclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type GetOrgUnitsIDUsersResponse GenericPage[OrgUnitUser]

type OrgUnitUser struct {
	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	APIOnlyUser        bool     `json:"apiOnlyUser"`
	Description        string   `json:"description"`
	IsEnabled          bool     `json:"isEnabled"`
	IsLDAP             bool     `json:"isLDAP"`
	IsLocked           bool     `json:"isLocked"`
	LoggedInUser       bool     `json:"loggedInUser"`
	ReadOnly           bool     `json:"readOnly"`
	SupportUser        bool     `json:"supportUser"`
	UserID             int      `json:"userId"`
	UserName           string   `json:"userName"`
	AccessGroupIDs     []int    `json:"accessGroupIds"`
	CurrentSSOProvider string   `json:"currentSsoProvider"`
	CustomerTree       []string `json:"customerTree"`
	FullName           string   `json:"fullName"`
	RoleIDs            []int    `json:"roleIds"`
	TwoFactorEnabled   bool     `json:"twoFactorEnabled"`
}

// See: https://developer.n-able.com/n-central/reference/listusers
func (c *Client) GetOrgUnitsIDUsers(ctx context.Context, orgUnitID string) ([]OrgUnitUser, error) {
	output := []OrgUnitUser{}

	pageURL := "/api/org-units/" + url.PathEscape(orgUnitID) + "/users"
	for {
		var response GetOrgUnitsIDUsersResponse
		err := c.Do(ctx, http.MethodGet, pageURL, nil, &response)
		if err != nil {
			return nil, fmt.Errorf("could not get page: %w", err)
		}
		output = append(output, response.Data...)
		if response.Links.NextPage == nil {
			break
		}
		pageURL = *response.Links.NextPage
	}
	return output, nil
}
