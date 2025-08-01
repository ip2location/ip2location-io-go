# Quickstart

## Dependencies

This module requires API key to function. You may sign up for a free API key at <https://www.ip2location.io/pricing>.

## Installation

Install this package using the command below:

``` bash
go get github.com/ip2location/ip2location-io-go
```

## Sample Codes

### Lookup IP Address Geolocation Data

You can make a geolocation data lookup for an IP address as below:

``` go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	apikey := "YOUR_API_KEY"
	
	config, err := ip2locationio.OpenConfiguration(apikey)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	ipl, err := ip2locationio.OpenIPGeolocation(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	ip := "8.8.8.8"
	lang := "en" // set to blank string if not applicable
	res, err := ipl.LookUp(ip, lang) // language parameter only available with Plus and Security plans
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	fmt.Printf("IP => %+v\n", res.IP)
	fmt.Printf("CountryCode => %+v\n", res.CountryCode)
	fmt.Printf("CountryName => %+v\n", res.CountryName)
	fmt.Printf("RegionName => %+v\n", res.RegionName)
	fmt.Printf("CityName => %+v\n", res.CityName)
	fmt.Printf("Latitude => %+v\n", res.Latitude)
	fmt.Printf("Longitude => %+v\n", res.Longitude)
	fmt.Printf("ZipCode => %+v\n", res.ZipCode)
	fmt.Printf("TimeZone => %+v\n", res.TimeZone)
	fmt.Printf("Asn => %+v\n", res.Asn)
	fmt.Printf("AS => %+v\n", res.AS)
	fmt.Printf("Isp => %+v\n", res.Isp)
	fmt.Printf("Domain => %+v\n", res.Domain)
	fmt.Printf("NetSpeed => %+v\n", res.NetSpeed)
	fmt.Printf("IddCode => %+v\n", res.IddCode)
	fmt.Printf("AreaCode => %+v\n", res.AreaCode)
	fmt.Printf("WeatherStationCode => %+v\n", res.WeatherStationCode)
	fmt.Printf("WeatherStationName => %+v\n", res.WeatherStationName)
	fmt.Printf("Mcc => %+v\n", res.Mcc)
	fmt.Printf("Mnc => %+v\n", res.Mnc)
	fmt.Printf("MobileBrand => %+v\n", res.MobileBrand)
	fmt.Printf("Elevation => %+v\n", res.Elevation)
	fmt.Printf("UsageType => %+v\n", res.UsageType)
	fmt.Printf("AddressType => %+v\n", res.AddressType)
	fmt.Printf("District => %+v\n", res.District)
	fmt.Printf("AdsCategory => %+v\n", res.AdsCategory)
	fmt.Printf("AdsCategoryName => %+v\n", res.AdsCategoryName)
	fmt.Printf("IsProxy => %+v\n", res.IsProxy)
	fmt.Printf("FraudScore => %+v\n", res.FraudScore)
	
	fmt.Printf("Continent.Name => %+v\n", res.Continent.Name)
	fmt.Printf("Continent.Code => %+v\n", res.Continent.Code)
	fmt.Printf("Continent.Hemisphere => %+v\n", res.Continent.Hemisphere)
	fmt.Printf("Continent.Translation.Lang => %+v\n", res.Continent.Translation.Lang)
	fmt.Printf("Continent.Translation.Value => %+v\n", res.Continent.Translation.Value)
	
	fmt.Printf("Country.Name => %+v\n", res.Country.Name)
	fmt.Printf("Country.Alpha3Code => %+v\n", res.Country.Alpha3Code)
	fmt.Printf("Country.NumericCode => %+v\n", res.Country.NumericCode)
	fmt.Printf("Country.Demonym => %+v\n", res.Country.Demonym)
	fmt.Printf("Country.Flag => %+v\n", res.Country.Flag)
	fmt.Printf("Country.Capital => %+v\n", res.Country.Capital)
	fmt.Printf("Country.TotalArea => %+v\n", res.Country.TotalArea)
	fmt.Printf("Country.Population => %+v\n", res.Country.Population)
	fmt.Printf("Country.Currency.Code => %+v\n", res.Country.Currency.Code)
	fmt.Printf("Country.Currency.Name => %+v\n", res.Country.Currency.Name)
	fmt.Printf("Country.Currency.Symbol => %+v\n", res.Country.Currency.Symbol)
	fmt.Printf("Country.Language.Code => %+v\n", res.Country.Language.Code)
	fmt.Printf("Country.Language.Name => %+v\n", res.Country.Language.Name)
	fmt.Printf("Country.Tld => %+v\n", res.Country.Tld)
	fmt.Printf("Country.Translation.Lang => %+v\n", res.Country.Translation.Lang)
	fmt.Printf("Country.Translation.Value => %+v\n", res.Country.Translation.Value)
	
	fmt.Printf("Region.Name => %+v\n", res.Region.Name)
	fmt.Printf("Region.Code => %+v\n", res.Region.Code)
	fmt.Printf("Region.Translation.Lang => %+v\n", res.Region.Translation.Lang)
	fmt.Printf("Region.Translation.Value => %+v\n", res.Region.Translation.Value)
	
	fmt.Printf("City.Name => %+v\n", res.City.Name)
	fmt.Printf("City.Translation.Lang => %+v\n", res.City.Translation.Lang)
	fmt.Printf("City.Translation.Value => %+v\n", res.City.Translation.Value)
	
	fmt.Printf("TimeZoneInfo.Olson => %+v\n", res.TimeZoneInfo.Olson)
	fmt.Printf("TimeZoneInfo.CurrentTime => %+v\n", res.TimeZoneInfo.CurrentTime)
	fmt.Printf("TimeZoneInfo.GmtOffset => %+v\n", res.TimeZoneInfo.GmtOffset)
	fmt.Printf("TimeZoneInfo.IsDst => %+v\n", res.TimeZoneInfo.IsDst)
	fmt.Printf("TimeZoneInfo.Abbreviation => %+v\n", res.TimeZoneInfo.Abbreviation)
	fmt.Printf("TimeZoneInfo.DstStartDate => %+v\n", res.TimeZoneInfo.DstStartDate)
	fmt.Printf("TimeZoneInfo.DstEndDate => %+v\n", res.TimeZoneInfo.DstEndDate)
	fmt.Printf("TimeZoneInfo.Sunrise => %+v\n", res.TimeZoneInfo.Sunrise)
	fmt.Printf("TimeZoneInfo.Sunset => %+v\n", res.TimeZoneInfo.Sunset)
	
	fmt.Printf("Geotargeting.Metro => %+v\n", res.Geotargeting.Metro)
	
	fmt.Printf("Proxy.LastSeen => %+v\n", res.Proxy.LastSeen)
	fmt.Printf("Proxy.ProxyType => %+v\n", res.Proxy.ProxyType)
	fmt.Printf("Proxy.Threat => %+v\n", res.Proxy.Threat)
	fmt.Printf("Proxy.Provider => %+v\n", res.Proxy.Provider)
	fmt.Printf("Proxy.IsVpn => %+v\n", res.Proxy.IsVpn)
	fmt.Printf("Proxy.IsTor => %+v\n", res.Proxy.IsTor)
	fmt.Printf("Proxy.IsDataCenter => %+v\n", res.Proxy.IsDataCenter)
	fmt.Printf("Proxy.IsPublicProxy => %+v\n", res.Proxy.IsPublicProxy)
	fmt.Printf("Proxy.IsWebProxy => %+v\n", res.Proxy.IsWebProxy)
	fmt.Printf("Proxy.IsWebCrawler => %+v\n", res.Proxy.IsWebCrawler)
	fmt.Printf("Proxy.IsResidentialProxy => %+v\n", res.Proxy.IsResidentialProxy)
	fmt.Printf("Proxy.IsSpammer => %+v\n", res.Proxy.IsSpammer)
	fmt.Printf("Proxy.IsScanner => %+v\n", res.Proxy.IsScanner)
	fmt.Printf("Proxy.IsBotnet => %+v\n", res.Proxy.IsBotnet)
}
```

### Lookup Domain Information

You can lookup domain information as below:

```go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	apikey := "YOUR_API_KEY"
	
	config, err := ip2locationio.OpenConfiguration(apikey)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	whois, err := ip2locationio.OpenDomainWhois(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	domain := "locaproxy.com"
	res, err := whois.LookUp(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	fmt.Printf("Domain => %+v\n", res.Domain)
	fmt.Printf("DomainID => %+v\n", res.DomainID)
	fmt.Printf("Status => %+v\n", res.Status)
	fmt.Printf("CreateDate => %+v\n", res.CreateDate)
	fmt.Printf("ExpireDate => %+v\n", res.ExpireDate)
	fmt.Printf("DomainAge => %+v\n", res.DomainAge)
	fmt.Printf("WhoisServer => %+v\n", res.WhoisServer)
	
	fmt.Printf("Registrar.IanaID => %+v\n", res.Registrar.IanaID)
	fmt.Printf("Registrar.Name => %+v\n", res.Registrar.Name)
	fmt.Printf("Registrar.Url => %+v\n", res.Registrar.Url)
	
	fmt.Printf("Registrant.Name => %+v\n", res.Registrant.Name)
	fmt.Printf("Registrant.Organization => %+v\n", res.Registrant.Organization)
	fmt.Printf("Registrant.StreetAddress => %+v\n", res.Registrant.StreetAddress)
	fmt.Printf("Registrant.City => %+v\n", res.Registrant.City)
	fmt.Printf("Registrant.Region => %+v\n", res.Registrant.Region)
	fmt.Printf("Registrant.ZipCode => %+v\n", res.Registrant.ZipCode)
	fmt.Printf("Registrant.Country => %+v\n", res.Registrant.Country)
	fmt.Printf("Registrant.Phone => %+v\n", res.Registrant.Phone)
	fmt.Printf("Registrant.Fax => %+v\n", res.Registrant.Fax)
	fmt.Printf("Registrant.Email => %+v\n", res.Registrant.Email)
	
	fmt.Printf("Admin.Name => %+v\n", res.Admin.Name)
	fmt.Printf("Admin.Organization => %+v\n", res.Admin.Organization)
	fmt.Printf("Admin.StreetAddress => %+v\n", res.Admin.StreetAddress)
	fmt.Printf("Admin.City => %+v\n", res.Admin.City)
	fmt.Printf("Admin.Region => %+v\n", res.Admin.Region)
	fmt.Printf("Admin.ZipCode => %+v\n", res.Admin.ZipCode)
	fmt.Printf("Admin.Country => %+v\n", res.Admin.Country)
	fmt.Printf("Admin.Phone => %+v\n", res.Admin.Phone)
	fmt.Printf("Admin.Fax => %+v\n", res.Admin.Fax)
	fmt.Printf("Admin.Email => %+v\n", res.Admin.Email)
	
	fmt.Printf("Tech.Name => %+v\n", res.Tech.Name)
	fmt.Printf("Tech.Organization => %+v\n", res.Tech.Organization)
	fmt.Printf("Tech.StreetAddress => %+v\n", res.Tech.StreetAddress)
	fmt.Printf("Tech.City => %+v\n", res.Tech.City)
	fmt.Printf("Tech.Region => %+v\n", res.Tech.Region)
	fmt.Printf("Tech.ZipCode => %+v\n", res.Tech.ZipCode)
	fmt.Printf("Tech.Country => %+v\n", res.Tech.Country)
	fmt.Printf("Tech.Phone => %+v\n", res.Tech.Phone)
	fmt.Printf("Tech.Fax => %+v\n", res.Tech.Fax)
	fmt.Printf("Tech.Email => %+v\n", res.Tech.Email)
	
	fmt.Printf("Billing.Name => %+v\n", res.Billing.Name)
	fmt.Printf("Billing.Organization => %+v\n", res.Billing.Organization)
	fmt.Printf("Billing.StreetAddress => %+v\n", res.Billing.StreetAddress)
	fmt.Printf("Billing.City => %+v\n", res.Billing.City)
	fmt.Printf("Billing.Region => %+v\n", res.Billing.Region)
	fmt.Printf("Billing.ZipCode => %+v\n", res.Billing.ZipCode)
	fmt.Printf("Billing.Country => %+v\n", res.Billing.Country)
	fmt.Printf("Billing.Phone => %+v\n", res.Billing.Phone)
	fmt.Printf("Billing.Fax => %+v\n", res.Billing.Fax)
	fmt.Printf("Billing.Email => %+v\n", res.Billing.Email)
	
	fmt.Printf("Nameservers => %+v\n", res.Nameservers)
}
```

### Convert Normal Text to Punycode

You can convert an international domain name to Punycode as below:

```go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	config, err := ip2locationio.OpenConfiguration("")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	whois, err := ip2locationio.OpenDomainWhois(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	domain = "täst.de"
	res, err := whois.GetPunycode(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", domain, res);
}
```

### Convert Punycode to Normal Text

You can convert a Punycode to international domain name as below:

```go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	config, err := ip2locationio.OpenConfiguration("")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	whois, err := ip2locationio.OpenDomainWhois(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	domain = "xn--tst-qla.de"
	res, err := whois.GetNormalText(domain)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", domain, res);
}
```

### Get Domain Name

You can extract the domain name from an url as below:

```go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	config, err := ip2locationio.OpenConfiguration("")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	whois, err := ip2locationio.OpenDomainWhois(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	testURL := "https://www.example.com/exe"
	res, err := whois.GetDomainName(testURL)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", testURL, res);
}
```

### Get Domain Extension

You can extract the domain extension from a domain name or url as below:

```go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	config, err := ip2locationio.OpenConfiguration("")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	whois, err := ip2locationio.OpenDomainWhois(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	testStr := "example.com"
	res, err := whois.GetDomainExtension(testStr)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v => %+v\n", testStr, res);
}
```

### Lookup IP Address Hosted Domains Data

You can make a hosted domains data lookup for an IP address as below:

``` go
package main

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"fmt"
)

func main() {
	apikey := "YOUR_API_KEY"
	
	config, err := ip2locationio.OpenConfiguration(apikey)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	hd, err := ip2locationio.OpenHostedDomain(config)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	ip := "8.8.8.8"
	page := 1
	res, err := hd.LookUp(ip, page)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	fmt.Printf("IP => %+v\n", res.IP)
	fmt.Printf("TotalDomains => %+v\n", res.TotalDomains)
	fmt.Printf("Page => %+v\n", res.Page)
	fmt.Printf("PerPage => %+v\n", res.PerPage)
	fmt.Printf("TotalPages => %+v\n", res.TotalPages)
	fmt.Printf("Domains => %+v\n", res.Domains)
}
```
