# IP2Location.io Go SDK — IP Geolocation, Domain WHOIS & Hosted Domain Lookup

**IP2Location.io Go SDK** is a Go package for retrieving **IPv4 and IPv6 geolocation, network, and IP intelligence data** using the [IP2Location.io API](https://www.ip2location.io/). It also supports **domain WHOIS lookup**, domain and Punycode utilities, and **hosted-domain lookup**. Integrate these services into Go applications without building the API requests manually.

**Quick links:** [Get an API key](https://www.ip2location.io/pricing) · [Developer documentation](https://ip2location-io-go.readthedocs.io/en/latest/index.html) · [API reference](https://ip2location-io-go.readthedocs.io/en/latest/code.html) · [GitHub repository](https://github.com/ip2location/ip2location-io-go)

## Features

- **IPv4 and IPv6 geolocation:** Look up country, region, district, city, latitude, longitude, and ZIP/postal code.
- **Network and ASN information:** Retrieve available ASN, AS name, ISP, associated domain, connection speed, and usage type.
- **Time zone and localization data:** Access time zone information and available daylight saving details; request translated location names on eligible Plus and Security plans.
- **Proxy and threat-related signals:** Access plan-dependent proxy, VPN, Tor, data center, crawler, and related indicators, plus available fraud scores.
- **Additional IP intelligence:** Retrieve eligible mobile carrier (MCC/MNC), country metadata, weather station, elevation, and advertising category fields.
- **Domain WHOIS lookup:** Get available registration dates, registrar information, nameservers, domain status, and contact details.
- **Internationalized domain and URL utilities:** Convert between Unicode domain names and Punycode, extract domain names from URLs, and get domain extensions.
- **Hosted domains:** Retrieve domains hosted on an IPv4 or IPv6 address, with paginated responses.

> **Plan availability:** Not all response fields are provided by every plan, and values may be unavailable for particular addresses or domains. See [API plans](https://www.ip2location.io/pricing) and the [Go SDK API reference](https://ip2location-io-go.readthedocs.io/en/latest/code.html).

## SDK functions

The [official Go SDK documentation](https://ip2location-io-go.readthedocs.io/en/latest/index.html) describes the following constructors and methods:

| Function / method | Description |
| --- | --- |
| `ip2locationio.OpenConfiguration(apiKey)` | Creates the SDK configuration using an API key. |
| `ip2locationio.OpenIPGeolocation(config)` | Initializes an IP geolocation client. |
| `geo.LookUp(ip, language)` | Looks up IP geolocation and available intelligence data. Use `""` when no translation is requested. |
| `ip2locationio.OpenDomainWhois(config)` | Initializes a domain WHOIS client. |
| `whois.LookUp(domain)` | Retrieves available WHOIS information for a domain. |
| `whois.GetPunycode(domain)` | Converts an internationalized domain name to Punycode. |
| `whois.GetNormalText(domain)` | Converts Punycode to Unicode text. |
| `whois.GetDomainName(url)` | Extracts a domain name from a URL. |
| `whois.GetDomainExtension(domain)` | Extracts a domain extension. |
| `ip2locationio.OpenHostedDomain(config)` | Initializes a hosted-domain lookup client. |
| `hosted.LookUp(ip, page)` | Retrieves hosted domains for an IP address and page number. |

The Go SDK uses exported struct fields such as `CountryName`, `CityName`, `Asn`, `Proxy.IsVpn`, and `TotalDomains`, rather than the JSON-style snake_case keys used by some other SDKs.

## API access and query limits

IP2Location.io offers keyless geolocation requests through its REST API and a Free plan with an API key:

| API service | Keyless (no API key) | Free plan (with API key) |
| --- | --- | --- |
| IP Geolocation API | **1,000 queries/day** | **50,000 queries/month** |
| Domain WHOIS API | — | **500 queries/month** |
| Hosted Domain API | — | **50 queries/month** |

- **Keyless access:** Limited keyless lookups can be made directly through the [IP Geolocation REST API](https://www.ip2location.io/ip2location-documentation). The allowance resets at **00:00 UTC**. The Go SDK documentation requires an API key; its examples do not describe a keyless SDK lookup mode.
- **Free plan:** Register for an API key to use the published monthly allowances above.
- **Attribution:** Keyless and Free usage require attribution; check the [pricing page](https://www.ip2location.io/pricing) for current terms.
- **Paid plans:** See the [pricing page](https://www.ip2location.io/pricing) for higher quotas, additional fields, and current pricing.

For details about authentication and rate limits, see the [IP Geolocation API documentation](https://www.ip2location.io/ip2location-documentation).

## Requirements

- **Go 1.18 or later**, based on the repository's [`go.mod`](https://github.com/ip2location/ip2location-io-go/blob/main/go.mod) directive.
- A Go module project, using the standard [Go modules](https://go.dev/ref/mod) tooling.
- An [IP2Location.io API key](https://www.ip2location.io/pricing) for SDK lookups.
- Network connectivity to the IP2Location.io services.

The Go module declares `golang.org/x/net` and `golang.org/x/text` dependencies; Go modules handles these when the SDK is installed.

## Installation

From your Go project directory, initialize a module if you do not already have one:

```bash
go mod init example.com/ip2location-demo
```

Install the official SDK using the command in the [Go quickstart](https://ip2location-io-go.readthedocs.io/en/latest/quickstart.html):

```bash
go get github.com/ip2location/ip2location-io-go
```

Keep the API key outside source control by setting an environment variable:

**Linux / macOS terminal:**

```bash
export IP2LOCATION_API_KEY='YOUR_API_KEY'
```

**Windows Command Prompt:**

```bat
set IP2LOCATION_API_KEY=YOUR_API_KEY
```

## Quickstart: IP geolocation lookup

Create `main.go` with the following example, adapted from the [official Go geolocation sample](https://ip2location-io-go.readthedocs.io/en/latest/quickstart.html). The API key is read from the environment rather than written directly into the code.

```go
package main

import (
    "fmt"
    "os"

    "github.com/ip2location/ip2location-io-go/ip2locationio"
)

func main() {
    apiKey := os.Getenv("IP2LOCATION_API_KEY")
    if apiKey == "" {
        fmt.Fprintln(os.Stderr, "Set the IP2LOCATION_API_KEY environment variable.")
        os.Exit(1)
    }

    config, err := ip2locationio.OpenConfiguration(apiKey)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    geo, err := ip2locationio.OpenIPGeolocation(config)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    result, err := geo.LookUp("8.8.8.8", "")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    fmt.Printf("IP: %v\n", result.IP)
    fmt.Printf("Country: %v\n", result.CountryName)
    fmt.Printf("City: %v\n", result.CityName)
    fmt.Printf("ASN: %v\n", result.Asn)
    fmt.Printf("ISP: %v\n", result.Isp)
}
```

Run the example from the project directory:

```bash
go run main.go
```

The SDK returns values through Go structs; available attributes depend on the IP address and subscription plan. See the [Go IPGeolocation response reference](https://ip2location-io-go.readthedocs.io/en/latest/code.html) for supported fields.

## More usage examples

The following snippets assume `config` has been created with `ip2locationio.OpenConfiguration(apiKey)` as in the quickstart, and that you are using the `ip2locationio` package within a Go function. The examples use the SDK's documented constructors and `LookUp()` methods; check returned errors before using results.

### Request translated location names

The second argument to `geo.LookUp(ip, language)` is a language code. Translated continent, country, region, and city names are available for **Plus and Security** plans. Use an empty string to omit the translation request.

```go
geo, err := ip2locationio.OpenIPGeolocation(config)
if err != nil {
    fmt.Println(err)
    return
}

result, err := geo.LookUp("8.8.8.8", "es")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("Country: %v\n", result.CountryName)
fmt.Printf("Country translation: %v\n", result.Country.Translation.Value)
```

### Inspect ASN and network information

The Go response exposes basic ASN and network attributes, together with extended AS fields when available.

```go
geo, err := ip2locationio.OpenIPGeolocation(config)
if err != nil {
    fmt.Println(err)
    return
}

result, err := geo.LookUp("8.8.8.8", "")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("ASN: %v\n", result.Asn)
fmt.Printf("AS name: %v\n", result.AS)
fmt.Printf("ISP: %v\n", result.Isp)
fmt.Printf("AS domain: %v\n", result.ASInfo.ASDomain)
```

### Inspect available proxy indicators

Proxy and threat indicators depend on the subscription plan and queried IP. Do not assume an absent or unpopulated field is proof that no risk exists.

```go
geo, err := ip2locationio.OpenIPGeolocation(config)
if err != nil {
    fmt.Println(err)
    return
}

result, err := geo.LookUp("8.8.8.8", "")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("Proxy indicator: %v\n", result.IsProxy)
fmt.Printf("VPN indicator: %v\n", result.Proxy.IsVpn)
fmt.Printf("Tor indicator: %v\n", result.Proxy.IsTor)
fmt.Printf("Fraud score: %v\n", result.FraudScore)
```

### Look up domain WHOIS information

The `DomainWhois` client looks up available registration details using the IP2WHOIS API. WHOIS contact information may be unavailable or redacted.

```go
whois, err := ip2locationio.OpenDomainWhois(config)
if err != nil {
    fmt.Println(err)
    return
}

result, err := whois.LookUp("locaproxy.com")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("Domain: %v\n", result.Domain)
fmt.Printf("Created: %v\n", result.CreateDate)
fmt.Printf("Expires: %v\n", result.ExpireDate)
fmt.Printf("Registrar: %v\n", result.Registrar.Name)
fmt.Printf("Nameservers: %v\n", result.Nameservers)
```

See the [Go DomainWhois API reference](https://ip2location-io-go.readthedocs.io/en/latest/code.html) for the documented response fields.

### Use domain and Punycode utilities

The SDK provides helpers for internationalized domain names and parsing domains from URLs. These utilities are demonstrated in the [official quickstart](https://ip2location-io-go.readthedocs.io/en/latest/quickstart.html).

```go
whois, err := ip2locationio.OpenDomainWhois(config)
if err != nil {
    fmt.Println(err)
    return
}

punycode, err := whois.GetPunycode("täst.de")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Punycode:", punycode)

normalText, err := whois.GetNormalText("xn--tst-qla.de")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Unicode:", normalText)

domainName, err := whois.GetDomainName("https://www.example.com/path")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Domain:", domainName)

extension, err := whois.GetDomainExtension("example.com")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Extension:", extension)
```

These helpers do not require a WHOIS lookup request, but they are accessed through the SDK's `DomainWhois` client.

### Find domains hosted on an IP address

The [official hosted-domain example](https://ip2location-io-go.readthedocs.io/en/latest/quickstart.html) passes a page number to `LookUp(ip, page)` and reads pagination metadata from the response.

```go
hosted, err := ip2locationio.OpenHostedDomain(config)
if err != nil {
    fmt.Println(err)
    return
}

result, err := hosted.LookUp("8.8.8.8", 1)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("IP: %v\n", result.IP)
fmt.Printf("Total domains: %v\n", result.TotalDomains)
fmt.Printf("Page: %v / %v\n", result.Page, result.TotalPages)
fmt.Printf("Domains: %v\n", result.Domains)
```

Pass `2`, `3`, etc. to request additional pages when the response indicates more pages are available.

## Common use cases

- **Website and application localization:** Suggest region-appropriate language or content based on approximate IP location.
- **Traffic analytics and enrichment:** Add geographic, ASN, ISP, and network attributes to server-side event records.
- **Security review:** Use available VPN, Tor, proxy, crawler, and fraud indicators as part of a broader risk assessment rather than as sole evidence.
- **Domain research:** Review registration dates, registrar details, and nameservers from WHOIS responses.
- **Network and hosting research:** Investigate AS information and domains hosted on an IP address.

## Documentation and support

- [Go SDK quickstart and sample codes](https://ip2location-io-go.readthedocs.io/en/latest/quickstart.html)
- [Go SDK API reference and response fields](https://ip2location-io-go.readthedocs.io/en/latest/code.html)
- [Official GitHub repository](https://github.com/ip2location/ip2location-io-go)
- [Go module manifest and requirements](https://github.com/ip2location/ip2location-io-go/blob/main/go.mod)
- [IP2Location.io API plans and query limits](https://www.ip2location.io/pricing)
- [IP Geolocation API authentication and rate limits](https://www.ip2location.io/ip2location-documentation)
- [Report a Go SDK issue or request a feature](https://github.com/ip2location/ip2location-io-go/issues)
- [API and account support](https://www.ip2location.io/contact-us)

## License

Licensed under the **MIT License**. See [LICENSE](./LICENSE).
