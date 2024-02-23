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
