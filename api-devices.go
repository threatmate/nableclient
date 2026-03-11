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
	StillLoggedIn            string   `json:"stillLoggedIn"`
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

type GetDevicesIDAssetsResponse GenericResult[DeviceAsset]

type DeviceAssetList[T any] struct {
	List []T `json:"list"`
}

type DeviceAssetExtraUSBDeviceItem struct {
	Index        int    `json:"_index"`
	Caption      string `json:"caption"`
	Manufacturer string `json:"manufacturer"`
	Status       string `json:"status"`
}

type DeviceAssetExtraOSFeatureItem struct {
	Index  int    `json:"_index"`
	PValue string `json:"pvalue"`
	PKey   string `json:"pkey"`
}

type DeviceAssetExtraMemoryItem struct {
	Index        int    `json:"_index"`
	SerialNumber string `json:"serialnumber"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	PartNumber   string `json:"partnumber"`
	Speed        string `json:"speed"`
	Manufacturer string `json:"manufacturer"`
	Capacity     string `json:"capacity"`
}

type DeviceAssetExtraOS struct {
	LicenseType    string  `json:"licensetype"`
	InstallDate    string  `json:"installdate"`
	SerialNumber   string  `json:"serialnumber"`
	Publisher      string  `json:"publisher"`
	CSDVersion     *string `json:"csdversion"`
	LastBootupTime string  `json:"lastbootuptime"`
	SupportedOS    string  `json:"supportedos"`
	LicenseKey     string  `json:"licensekey"`
}

type DeviceAssetExtraMediaAccessDeviceItem struct {
	Index     int    `json:"_index"`
	MediaType string `json:"mediatype"`
	UniqueID  string `json:"uniqueid"`
}

type DeviceAssetExtraFolderForShareItem struct {
	Index     int    `json:"_index"`
	Path      string `json:"path"`
	ShareName string `json:"sharename"`
}

type DeviceAssetExtraPrinterItem struct {
	Index         int    `json:"_index"`
	Path          string `json:"path"`
	Port          string `json:"port"`
	Name          string `json:"name"`
	SystemDefault string `json:"systemdefault"`
}

type DeviceAssetExtraMotherboard struct {
	Product      string `json:"product"`
	SerialNumber string `json:"serialnumber"`
	BIOSVersion  string `json:"biosversion"`
	Version      string `json:"version"`
	Manufacturer string `json:"manufacturer"`
}

type DeviceAssetExtraPhysicalDriveItem struct {
	Index        int    `json:"_index"`
	SerialNumber string `json:"serialnumber"`
	ModelNumber  string `json:"modelnumber"`
	Capacity     string `json:"capacity"`
}

type DeviceAssetExtraProcessor struct {
	MaxClockSpeed string  `json:"maxclockSpeed"`
	CPUID         string  `json:"cpuid"`
	Vendor        *string `json:"vendor"`
	Description   string  `json:"description"`
	Architecture  *string `json:"architecture"`
}

type DeviceAssetExtraVideoControllerItem struct {
	Index             int    `json:"_index"`
	Name              string `json:"name"`
	VideoControllerID string `json:"videocontrollerid"`
	Description       string `json:"description"`
	AdapterRAM        string `json:"adapterram"`
}

type DeviceAssetExtraSOCustomer struct {
	CustomerID   string `json:"customerid"`
	CustomerName string `json:"customername"`
}

type DeviceAssetExtraApplicationItem struct {
	Index            int     `json:"_index"`
	LicenseType      *string `json:"licensetype"`
	InstallationDate *string `json:"installationdate"`
	DisplayName      string  `json:"displayname"`
	Publisher        string  `json:"publisher"`
	Version          string  `json:"version"`
	LicenseKey       *string `json:"licensekey"`
}

type DeviceAssetExtraPortItem struct {
	Index       int    `json:"_index"`
	Port        string `json:"port"`
	ServiceName string `json:"servicename"`
}

type DeviceAssetExtraService struct {
	Index          int    `json:"_index"`
	StartupType    string `json:"startuptype"`
	Caption        string `json:"caption"`
	ServiceName    string `json:"servicename"`
	ExecutableName string `json:"executablename"`
	UserAccount    string `json:"useraccount"`
}

type DeviceAssetExtraComputerSystem struct {
	PopulatedMemorySlots string `json:"populatedmemory_slots"`
	TotalMemorySlots     string `json:"totalmemory_slots"`
	SystemType           string `json:"systemtype"`
	UUID                 string `json:"uuid"`
}

type DeviceAssetExtraLogicalDeviceItem struct {
	Index       int    `json:"_index"`
	MaxCapacity string `json:"maxcapacity"`
	VolumeName  string `json:"volumename"`
}

type DeviceAssetExtraDevice struct {
	TakeControlUUID               string `json:"takecontroluuid"`
	LastLoggedInUserStillLoggedIn string `json:"lastloggedinuser_stillloggedin"`
	LastLoggedInUserSessionType   string `json:"lastloggedinuser_sessiontype"`
	CustomerID                    string `json:"customerid"`
	WarrantyExpirationDate        string `json:"warrantyexpirationdate"`
	LastLoggedInUserDomain        string `json:"lastloggedinuser_domain"`
	CreatedOn                     string `json:"createdon"`
	LastLoggedInUser              string `json:"lastloggedinuser"`
	NCentralAssetTag              string `json:"ncentralassettag"`
}

type DeviceAssetExtraCustomer struct {
	CustomerID   string `json:"customerid"`
	CustomerName string `json:"customername"`
}

type DeviceAssetExtra struct {
	USBDevice         DeviceAssetList[DeviceAssetExtraUSBDeviceItem]         `json:"usbDevice"`
	OSFeatures        DeviceAssetList[DeviceAssetExtraOSFeatureItem]         `json:"osFeatures"`
	Memory            DeviceAssetList[DeviceAssetExtraMemoryItem]            `json:"memory"`
	OS                DeviceAssetExtraOS                                     `json:"os"`
	MediaAccessDevice DeviceAssetList[DeviceAssetExtraMediaAccessDeviceItem] `json:"mediaaccessdevice"`
	FolderForShare    DeviceAssetList[DeviceAssetExtraFolderForShareItem]    `json:"folderforshare"`
	Printer           DeviceAssetList[DeviceAssetExtraPrinterItem]           `json:"printer"`
	Motherboard       DeviceAssetExtraMotherboard                            `json:"motherboard"`
	PhysicalDrive     DeviceAssetList[DeviceAssetExtraPhysicalDriveItem]     `json:"physicaldrive"`
	Processor         DeviceAssetExtraProcessor                              `json:"processor"`
	VideoController   DeviceAssetList[DeviceAssetExtraVideoControllerItem]   `json:"videocontroller"`
	SOCustomer        DeviceAssetExtraSOCustomer                             `json:"socustomer"`
	Application       DeviceAssetList[DeviceAssetExtraApplicationItem]       `json:"application"`
	Port              DeviceAssetList[DeviceAssetExtraPortItem]              `json:"port"`
	Service           DeviceAssetExtraService                                `json:"service"`
	ComputerSystem    DeviceAssetExtraComputerSystem                         `json:"computersystem"`
	LogicalDevice     DeviceAssetList[DeviceAssetExtraLogicalDeviceItem]     `json:"logicaldevice"`
	Device            DeviceAssetExtraDevice                                 `json:"device"`
	Customer          DeviceAssetExtraCustomer                               `json:"customer"`
}

type DeviceAssetOS struct {
	ReportedOS     string `json:"reportedos"`
	OSArchitecture string `json:"osarchitecture"`
	Version        string `json:"version"`
}

type DeviceAssetApplicationItem struct {
	Index       int    `json:"_index"`
	DisplayName string `json:"displayname"`
}

type DeviceAssetComputerSystem struct {
	SerialNumber        string `json:"serialnumber"`
	NetBIOSName         string `json:"netbiosname"`
	Model               string `json:"model"`
	TotalPhysicalMemory string `json:"totalphysicalmemory"`
	Manufacturer        string `json:"manufacturer"`
}

type DeviceAssetNetworkAdapterItem struct {
	Index       int    `json:"_index"`
	IPAddress   string `json:"ipaddress"`
	DNSServer   string `json:"dnsserver"`
	Description string `json:"description"`
	DHCPServer  string `json:"dhcpserver"`
	MACAddress  string `json:"macaddress"`
	Gateway     string `json:"gateway"`
}

type DeviceAssetDevice struct {
	LongName    string `json:"longname"`
	Deleted     string `json:"deleted"`
	LastLogin   string `json:"lastlogin"`
	DeviceClass string `json:"deviceclass"`
	DeviceID    string `json:"deviceid"`
	URI         string `json:"uri"`
}

type DeviceAssetProcessor struct {
	Name          string `json:"name"`
	NumberOfCores string `json:"numberofcores"`
	NumberOfCPUs  string `json:"numberofcpus"`
}

type DeviceAsset struct {
	Extra          DeviceAssetExtra                               `json:"_extra"`
	OS             DeviceAssetOS                                  `json:"os"`
	Application    DeviceAssetList[DeviceAssetApplicationItem]    `json:"application"`
	ComputerSystem DeviceAssetComputerSystem                      `json:"computersystem"`
	NetworkAdapter DeviceAssetList[DeviceAssetNetworkAdapterItem] `json:"networkadapter"`
	Device         DeviceAssetDevice                              `json:"device"`
	Processor      DeviceAssetProcessor                           `json:"processor"`
}

// See: https://developer.n-able.com/n-central/reference/getassetinfo
func (c *Client) GetDevicesIDAssets(ctx context.Context, deviceID int) (DeviceAsset, error) {
	output := DeviceAsset{}

	pageURL := "/api/devices/" + fmt.Sprintf("%d", deviceID) + "/assets"
	{
		var response GetDevicesIDAssetsResponse
		err := c.Do(ctx, http.MethodGet, pageURL, nil, &response)
		if err != nil {
			return output, fmt.Errorf("could not get page: %w", err)
		}
		output = response.Data
	}
	return output, nil
}
