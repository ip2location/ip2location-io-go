package ip2locationio

import (
	"testing"
)

// TestInvalidKeyIPGeolocation calls LookUp with an
// invalid API key to see if it returns a result.
func TestInvalidKeyIPGeolocation(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := ""
	res, err := ipl.LookUp(ip, lang)

	if err == nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, %v, error`, res, err)
	}
}

// TestLang calls LookUp without an API key
// so it should not allow the Lang parameter.
func TestLang(t *testing.T) {
	dummy := "DUMMY"
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := "en"
	res, err := ipl.LookUp(ip, lang)

	if err == nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, %v, error`, res, err)
	}
}

// TestIPGeolocation calls LookUp without an API key
// to see if it returns a result.
func TestIPGeolocation(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := ""
	res, err := ipl.LookUp(ip, lang)

	if err != nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, %v, error`, res, err)
	}
}

// TestURLIPGeolocation calls LookUp with a bad URL.
func TestURLIPGeolocation(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	ipl.baseUrl = ipl.baseUrl + ".my"
	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := ""
	res, err := ipl.LookUp(ip, lang)

	if err == nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, %v, error`, res, err)
	}
}

// TestNonJSONIPGeolocation calls LookUp with a non-JSON result URL.
func TestNonJSONIPGeolocation(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	ipl.baseUrl = "ip2location.io"
	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := ""
	res, err := ipl.LookUp(ip, lang)

	if err == nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, %v, error`, res, err)
	}
}

// TestHTTPResponseIPGeolocation calls LookUp with a 404 URL.
func TestHTTPResponseIPGeolocation(t *testing.T) {
	dummy := "" // blank key
	config, err := OpenConfiguration(dummy)

	if err != nil {
		t.Fatalf(`OpenConfiguration(dummy) = %+v, %v, error`, config, err)
	}

	ipl, err := OpenIPGeolocation(config)

	ipl.baseUrl = "ip2location.io/notfound"
	if err != nil {
		t.Fatalf(`OpenIPGeolocation(config) = %+v, %v, error`, ipl, err)
	}

	ip := "8.8.8.8"
	lang := ""
	res, err := ipl.LookUp(ip, lang)

	if err == nil {
		t.Fatalf(`ipl.LookUp(ip, lang) = %+v, error`, res)
	}
}
