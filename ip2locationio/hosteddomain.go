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

// The HostedDomainResult struct stores all of the available
// hosted domains info found in the IP2Location.io API.
type HostedDomainResult struct {
	IP           string   `json:"ip"`
	TotalDomains int      `json:"total_domains"`
	Page         int      `json:"page"`
	PerPage      int      `json:"per_page"`
	TotalPages   int      `json:"total_pages"`
	Domains      []string `json:"domains"`
}

// The HostedDomainError struct stores errors
// returned by the IP2Location.io API.
type HostedDomainError struct {
	Error struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"error"`
}

// The HostedDomain struct is the main object used to query the IP2Location.io API.
type HostedDomain struct {
	configuration *Configuration
	baseUrl       string
}

// OpenHostedDomain initializes with the Configuration object
func OpenHostedDomain(config *Configuration) (*HostedDomain, error) {
	var hd = &HostedDomain{}
	hd.configuration = config
	hd.baseUrl = "domains.ip2whois.com/domains"
	return hd, nil
}

// LookUp will return all hosted domains fields based on the queried IP address
func (a *HostedDomain) LookUp(ip string, page int) (HostedDomainResult, error) {
	var res HostedDomainResult
	var ex HostedDomainError

	myUrl := "https://" + a.baseUrl + "?key=" + url.QueryEscape(a.configuration.apiKey) + "&ip=" + url.QueryEscape(ip) + "&page=" + strconv.Itoa(page) + "&source=" + url.QueryEscape(a.configuration.source) + "&source_version=" + url.QueryEscape(a.configuration.sourceVersion)

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
