package types

type AccessoryInfo struct {
	ProductName         string   `json:"productName" yaml:"productName"`
	HardwareBoardType   int      `json:"hardwareBoardType" yaml:"hardwareBoardType"`
	FirmwareBuildNumber int      `json:"firmwareBuildNumber" yaml:"firmwareBuildNumber"`
	FirmwareVersion     string   `json:"firmwareVersion" yaml:"firmwareVersion"`
	SerialNumber        string   `json:"serialNumber" yaml:"serialNumber"`
	DisplayName         string   `json:"displayName" yaml:"displayName"`
	Features            []string `json:"features" yaml:"features"`
}
