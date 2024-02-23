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
