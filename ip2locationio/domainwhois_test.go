package ip2locationio

import (
	"testing"
)

// TestInvalidKeyWhois calls LookUp without an API key
// to see if it returns a result.
func TestInvalidKeyWhois(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "locaproxy.com"
	res, err := whois.LookUp(domain)

	if err == nil {
		t.Fatalf(`whois.LookUp(domain) = %+v, error`, res)
	}
}

// TestBadURLWhois calls LookUp with a bad URL.
func TestBadURLWhois(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)
	whois.baseUrl = whois.baseUrl + ".my"

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "locaproxy.com"
	res, err := whois.LookUp(domain)

	if err == nil {
		t.Fatalf(`whois.LookUp(domain) = %+v, error`, res)
	}
}

// TestNonJSONWhois calls LookUp with non-JSON result URL.
func TestNonJSONWhois(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)
	whois.baseUrl = "ip2location.io"

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "locaproxy.com"
	res, err := whois.LookUp(domain)

	if err == nil {
		t.Fatalf(`whois.LookUp(domain) = %+v, error`, res)
	}
}

// TestJSONWhois calls LookUp with JSON result URL.
func TestJSONWhois(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)
	whois.baseUrl = "ip2location.io/get-ip.json"

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "locaproxy.com"
	res, err := whois.LookUp(domain)

	if err != nil {
		t.Fatalf(`whois.LookUp(domain) = %+v, error`, res)
	}
}

// TestGetPunycode calls GetPunycode
// to see if it returns the correct Punycode.
func TestGetPunycode(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "täst.de"
	expected := "xn--tst-qla.de"
	res, err := whois.GetPunycode(domain)

	if err != nil || res != expected {
		t.Fatalf(`whois.GetPunycode(domain) = %+v, error`, res)
	}
}

// TestGetBadPunycode calls GetPunycode with space
// to see if it throws error.
func TestGetBadPunycode(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "fake this"
	_, err = whois.GetPunycode(domain)

	if err == nil {
		t.Fatalf(`whois.GetPunycode(domain) = %+v, error`, domain)
	}
}

// TestGetNormalText calls GetNormalText
// to see if it returns the correct UTF8 domain.
func TestGetNormalText(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "xn--tst-qla.de"
	expected := "täst.de"
	res, err := whois.GetNormalText(domain)

	if err != nil || res != expected {
		t.Fatalf(`whois.GetNormalText(domain) = %+v, error`, res)
	}
}

// TestGetBadNormalText calls GetNormalText with space
// to see if it throws error.
func TestGetBadNormalText(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	domain := "fake this"
	_, err = whois.GetNormalText(domain)

	if err == nil {
		t.Fatalf(`whois.GetNormalText(domain) = %+v, error`, domain)
	}
}

// TestGetDomainName calls GetDomainName
// to see if it returns the correct domain.
func TestGetDomainName(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	url := "https://www.example.com/exe"
	expected := "example.com"
	res, err := whois.GetDomainName(url)

	if err != nil || res != expected {
		t.Fatalf(`whois.GetDomainName(url) = %+v, error`, res)
	}
}

// TestGetBadDomainName calls GetDomainName with bad domain
// to see if it throws error.
func TestGetBadDomainName(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	url := "fake this"
	_, err = whois.GetDomainName(url)

	if err == nil {
		t.Fatalf(`whois.GetDomainName(url) = %+v, error`, url)
	}
}

// TestGetDomainExtension calls GetDomainExtension
// to see if it returns the correct domain extension.
func TestGetDomainExtension(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	str := "example.com"
	expected := ".com"
	res, err := whois.GetDomainExtension(str)

	if err != nil || res != expected {
		t.Fatalf(`whois.GetDomainExtension(str) = %+v, error`, res)
	}
}

// TestGetBadDomainExtension calls GetDomainExtension with bad domain
// to see if it throws error.
func TestGetBadDomainExtension(t *testing.T) {
	config, err := OpenConfiguration("")

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	whois, err := OpenDomainWhois(config)

	if err != nil {
		t.Fatalf(`OpenDomainWhois(config) = %+v, %v, error`, whois, err)
	}

	str := "example"
	res, err := whois.GetDomainExtension(str)

	if err == nil {
		t.Fatalf(`whois.GetDomainExtension(str) = %+v, error`, res)
	}
}
