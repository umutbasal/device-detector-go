package devicedetector

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
	Client ClientResult
	Device DeviceResult
	OS     OperatingSystemResult
	Bot    BotResult
}

type ClientResult struct {
	Name          string
	Version       string
	Engine        string
	EngineVersion string
	Type          string
	URL           string
}

type DeviceResult struct {
	Type  DeviceType
	Brand string
	Model string
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
	Name     string
	Version  string
	Platform string
}

type BotResult struct {
	Name     string
	Category string
	URL      string
	Producer struct {
		Name string
		URL  string
	}
}
