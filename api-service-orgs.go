package nableclient

import (
	"context"
	"fmt"
	"net/http"
)

type GetServiceOrgsResponse GenericPage[ServiceOrg]

type ServiceOrg struct {
	SOID              string  `json:"soId"`
	SOName            string  `json:"soName"`
	OrgUnitType       string  `json:"orgUnitType"`
	ParentID          string  `json:"parentId"`
	ExternalID        string  `json:"externalId"`
	ExternalID2       string  `json:"externalId2"`
	ContactFirstName  string  `json:"contactFirstName"`
	ContactLastName   string  `json:"contactLastName"`
	Phone             string  `json:"phone"`
	ContactTitle      string  `json:"contactTitle"`
	ContactEmail      string  `json:"contactEmail"`
	ContactPhone      string  `json:"contactPhone"`
	ContactPhoneExt   string  `json:"contactPhoneExt"`
	ContactDepartment string  `json:"contactDepartment"`
	Street1           string  `json:"street1"`
	Street2           string  `json:"street2"`
	City              string  `json:"city"`
	StateProv         string  `json:"stateProv"`
	Country           *string `json:"country"`
	PostalCode        string  `json:"postalCode"`
	IsSystem          bool    `json:"isSystem"`
	IsServiceOrg      bool    `json:"isServiceOrg"`
}

// See: https://developer.n-able.com/n-central/reference/listServiceOrgs
func (c *Client) GetServiceOrgs(ctx context.Context) ([]ServiceOrg, error) {
	output := []ServiceOrg{}

	pageURL := "/api/service-orgs"
	for {
		var response GetServiceOrgsResponse
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
