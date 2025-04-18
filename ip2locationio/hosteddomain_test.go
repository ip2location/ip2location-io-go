package ip2locationio

import (
	"testing"
)

// TestInvalidKeyHostedDomain calls LookUp with an
// invalid API key to see if it returns a result.
func TestInvalidKeyHostedDomain(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	hd, err := OpenHostedDomain(config)

	if err != nil {
		t.Fatalf(`OpenHostedDomain(config) = %+v, %v, error`, hd, err)
	}

	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)

	if err == nil {
		t.Fatalf(`hd.LookUp(ip, page) = %+v, %v, error`, res, err)
	}
}

// TestHostedDomain calls LookUp without an API key
// to see if it returns a result.
func TestHostedDomain(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	hd, err := OpenHostedDomain(config)

	if err != nil {
		t.Fatalf(`OpenHostedDomain(config) = %+v, %v, error`, hd, err)
	}

	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)

	if err != nil {
		t.Fatalf(`hd.LookUp(ip, page) = %+v, %v, error`, res, err)
	}
}

// TestURLHostedDomain calls LookUp with a bad URL.
func TestURLHostedDomain(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	hd, err := OpenHostedDomain(config)

	hd.baseUrl = hd.baseUrl + ".my"
	if err != nil {
		t.Fatalf(`OpenHostedDomain(config) = %+v, %v, error`, hd, err)
	}

	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)

	if err == nil {
		t.Fatalf(`hd.LookUp(ip, page) = %+v, %v, error`, res, err)
	}
}

// TestNonJSONHostedDomain calls LookUp with a non-JSON result URL.
func TestNonJSONHostedDomain(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	hd, err := OpenHostedDomain(config)

	hd.baseUrl = "ip2whois.com"
	if err != nil {
		t.Fatalf(`OpenHostedDomain(config) = %+v, %v, error`, hd, err)
	}

	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)

	if err == nil {
		t.Fatalf(`hd.LookUp(ip, page) = %+v, %v, error`, res, err)
	}
}

// TestHTTPResponseHostedDomain calls LookUp with a 404 URL.
func TestHTTPResponseHostedDomain(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	hd, err := OpenHostedDomain(config)

	hd.baseUrl = "ip2whois.com/notfound"
	if err != nil {
		t.Fatalf(`OpenHostedDomain(config) = %+v, %v, error`, hd, err)
	}

	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)

	if err == nil {
		t.Fatalf(`hd.LookUp(ip, page) = %+v, error`, res)
	}
}
