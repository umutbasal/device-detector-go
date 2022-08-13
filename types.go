package devicedetector

// This types are based to js library to parse result

type DeviceDetectorOptions struct {
	SkipBotDetection  bool
	VersionTruncation VersionTruncation
}

type VersionTruncation int

const (
	VC VersionTruncation = iota
	VC1
	VC2
	VC3
)

type DeviceDetectorResult struct {
	Client ClientResult          `json:"client"`
	Device DeviceResult          `json:"device"`
	OS     OperatingSystemResult `json:"os"`
	Bot    BotResult             `json:"bot"`
}

type ClientResult struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Engine        string `json:"engine"`
	EngineVersion string `json:"engineVersion"`
	Type          string `json:"type"`
	URL           string `json:"url"`
}

type DeviceResult struct {
	Type  DeviceType `json:"type"`
	Brand string     `json:"brand"`
	Model string     `json:"model"`
}

type DeviceType string

const (
	DeviceTypeDesktop             DeviceType = "desktop"
	DeviceTypeSmartphone          DeviceType = "smartphone"
	DeviceTypeTablet              DeviceType = "tablet"
	DeviceTypeTelevision          DeviceType = "television"
	DeviceTypeSmartDisplay        DeviceType = "smart display"
	DeviceTypeCamera              DeviceType = "camera"
	DeviceTypeCar                 DeviceType = "car"
	DeviceTypeConsole             DeviceType = "console"
	DeviceTypePortableMediaPlayer DeviceType = "portable media player"
	DeviceTypePhablet             DeviceType = "phablet"
	DeviceTypeWearable            DeviceType = "wearable"
	DeviceTypeSmartSpeaker        DeviceType = "smart speaker"
	DeviceTypeFeaturePhone        DeviceType = "feature phone"
	DeviceTypePeripheral          DeviceType = "peripheral"
)

type OperatingSystemResult struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Platform string `json:"platform"`
}

type BotResult struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	URL      string `json:"url"`
	Producer struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"producer"`
}
