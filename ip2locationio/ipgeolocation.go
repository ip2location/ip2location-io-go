package ip2locationio

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// The IPGeolocationResult struct stores all of the available
// geolocation info found in the IP2Location.io API.
type IPGeolocationResult struct {
	IP                 string  `json:"ip"`
	CountryCode        string  `json:"country_code"`
	CountryName        string  `json:"country_name"`
	RegionName         string  `json:"region_name"`
	CityName           string  `json:"city_name"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	ZipCode            string  `json:"zip_code"`
	TimeZone           string  `json:"time_zone"`
	Asn                string  `json:"asn"`
	AS                 string  `json:"as"`
	Isp                string  `json:"isp"`
	Domain             string  `json:"domain"`
	NetSpeed           string  `json:"net_speed"`
	IddCode            string  `json:"idd_code"`
	AreaCode           string  `json:"area_code"`
	WeatherStationCode string  `json:"weather_station_code"`
	WeatherStationName string  `json:"weather_station_name"`
	Mcc                string  `json:"mcc"`
	Mnc                string  `json:"mnc"`
	MobileBrand        string  `json:"mobile_brand"`
	Elevation          int     `json:"elevation"`
	UsageType          string  `json:"usage_type"`
	AddressType        string  `json:"address_type"`
	Continent          struct {
		Name        string   `json:"name"`
		Code        string   `json:"code"`
		Hemisphere  []string `json:"hemisphere"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"continent"`
	District string `json:"district"`
	Country  struct {
		Name        string `json:"name"`
		Alpha3Code  string `json:"alpha3_code"`
		NumericCode int    `json:"numeric_code"`
		Demonym     string `json:"demonym"`
		Flag        string `json:"flag"`
		Capital     string `json:"capital"`
		TotalArea   int    `json:"total_area"`
		Population  int    `json:"population"`
		Currency    struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"currency"`
		Language struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"language"`
		Tld         string `json:"tld"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"country"`
	Region struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"region"`
	City struct {
		Name        string `json:"name"`
		Translation struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"translation"`
	} `json:"city"`
	TimeZoneInfo struct {
		Olson       string `json:"olson"`
		CurrentTime string `json:"current_time"`
		GmtOffset   int    `json:"gmt_offset"`
		IsDst       bool   `json:"is_dst"`
		Sunrise     string `json:"sunrise"`
		Sunset      string `json:"sunset"`
	} `json:"time_zone_info"`
	Geotargeting struct {
		Metro string `json:"metro"`
	} `json:"geotargeting"`
	AdsCategory     string `json:"ads_category"`
	AdsCategoryName string `json:"ads_category_name"`
	IsProxy         bool   `json:"is_proxy"`
	Proxy           struct {
		LastSeen           int    `json:"last_seen"`
		ProxyType          string `json:"proxy_type"`
		Threat             string `json:"threat"`
		Provider           string `json:"provider"`
		IsVpn              bool   `json:"is_vpn"`
		IsTor              bool   `json:"is_tor"`
		IsDataCenter       bool   `json:"is_data_center"`
		IsPublicProxy      bool   `json:"is_public_proxy"`
		IsWebProxy         bool   `json:"is_web_proxy"`
		IsWebCrawler       bool   `json:"is_web_crawler"`
		IsResidentialProxy bool   `json:"is_residential_proxy"`
		IsSpammer          bool   `json:"is_spammer"`
		IsScanner          bool   `json:"is_scanner"`
		IsBotnet           bool   `json:"is_botnet"`
	} `json:"proxy"`
}

// The IPGeolocationError struct stores errors
// returned by the IP2Location.io API.
type IPGeolocationError struct {
	Error struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"error"`
}

// The IPGeolocation struct is the main object used to query the IP2Location.io API.
type IPGeolocation struct {
	configuration *Configuration
	baseUrl       string
}

// OpenIPGeolocation initializes with the Configuration object
func OpenIPGeolocation(config *Configuration) (*IPGeolocation, error) {
	var ipl = &IPGeolocation{}
	ipl.configuration = config
	ipl.baseUrl = "api.ip2location.io"
	return ipl, nil
}

// LookUp will return all geolocation fields based on the queried IP address
func (a *IPGeolocation) LookUp(ip string, lang string) (IPGeolocationResult, error) {
	var res IPGeolocationResult
	var ex IPGeolocationError

	// myUrl := "https://" + a.baseUrl + "?key=" + url.QueryEscape(a.configuration.apiKey) + "&ip=" + url.QueryEscape(ip) + "&lang=" + url.QueryEscape(lang) + "&source=" + url.QueryEscape(a.configuration.source) + "&source_version=" + url.QueryEscape(a.configuration.sourceVersion)
	
	myUrl := "https://" + a.baseUrl + "?ip=" + url.QueryEscape(ip) + "&source=" + url.QueryEscape(a.configuration.source) + "&source_version=" + url.QueryEscape(a.configuration.sourceVersion)

	if a.configuration.apiKey != "" {
		myUrl = myUrl + "&key=" + url.QueryEscape(a.configuration.apiKey)
	}

	if lang != "" {
		myUrl = myUrl + "&lang=" + url.QueryEscape(lang)
	}

	resp, err := http.Get(myUrl)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return res, err
		}

		err = json.Unmarshal(bodyBytes, &res)

		if err != nil {
			return res, err
		}

		return res, nil
	} else if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return res, err
		}

		bodyStr := string(bodyBytes[:])
		if strings.Contains(bodyStr, "error_message") {
			err = json.Unmarshal(bodyBytes, &ex)

			if err != nil {
				return res, err
			}
			return res, errors.New("Error: " + ex.Error.ErrorMessage)
		}
	}

	return res, errors.New("Error HTTP " + strconv.Itoa(int(resp.StatusCode)))
}
