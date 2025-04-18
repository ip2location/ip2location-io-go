package ip2locationio

// The Configuration struct stores the IP2Location.io API key.
type Configuration struct {
	apiKey        string
	source        string
	sourceVersion string
}

// OpenConfiguration initializes with the IP2Location.io API key
func OpenConfiguration(apikey string) (*Configuration, error) {
	var config = &Configuration{}
	config.apiKey = apikey
	config.source = "sdk-go-iplio"
	config.sourceVersion = "1.2.0"
	return config, nil
}
