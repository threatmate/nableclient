package nableclient

import (
	"context"
	"fmt"
	"net/http"
)

type GetDevicesResponse GenericPage[Device]

type Device struct {
	DeviceID                 int      `json:"deviceId"`
	URI                      string   `json:"uri"`
	RemoteControlURI         *string  `json:"remoteControlUri"`
	SourceURI                string   `json:"sourceUri"`
	LongName                 string   `json:"longName"`
	DeviceClass              string   `json:"deviceClass"`
	Description              string   `json:"description"`
	IsProbe                  bool     `json:"isProbe"`
	OSID                     string   `json:"osId"`
	SupportedOS              string   `json:"supportedOs"`
	DiscoveredName           string   `json:"discoveredName"`
	DeviceClassLabel         string   `json:"deviceClassLabel"`
	SupportedOSLabel         string   `json:"supportedOsLabel"`
	LastLoggedInUser         string   `json:"lastLoggedInUser"`
	StillLoggedIn            bool     `json:"stillLoggedIn"`
	LicenseMode              string   `json:"licenseMode"`
	OrgUnitID                int      `json:"orgUnitId"`
	SOID                     int      `json:"soId"`
	SOName                   string   `json:"soName"`
	CustomerID               int      `json:"customerId"`
	CustomerName             string   `json:"customerName"`
	SiteID                   *int     `json:"siteId"`
	SiteName                 *string  `json:"siteName"`
	ApplianceID              int      `json:"applianceId"`
	LastApplianceCheckinTime DateTime `json:"lastApplianceCheckinTime"`
}

// See: https://developer.n-able.com/n-central/reference/listdevices
func (c *Client) GetDevices(ctx context.Context) ([]Device, error) {
	output := []Device{}

	pageURL := "/api/devices"
	for {
		var response GetDevicesResponse
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
