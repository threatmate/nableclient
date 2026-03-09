package nableclient

import (
	"context"
	"fmt"
	"net/http"
)

type GetCustomersResponse GenericPage[Customer]

type Customer struct {
	CustomerID        string  `json:"customerId"`
	CustomerName      string  `json:"customerName"`
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

// See: https://developer.n-able.com/n-central/reference/listcustomers_1
func (c *Client) GetCustomers(ctx context.Context) ([]Customer, error) {
	output := []Customer{}

	pageURL := "/api/customers"
	for {
		var response GetCustomersResponse
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
