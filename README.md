[![Go Report Card](https://goreportcard.com/badge/github.com/ip2location/ip2location-io-go/ip2locationio)](https://goreportcard.com/report/github.com/ip2location/ip2location-io-go/ip2locationio)

IP2Location.io Go SDK
=====================
This Go package enables user to query for an enriched data set, such as country, region, district, city, latitude & longitude, ZIP code, time zone, ASN, ISP, domain, net speed, IDD code, area code, weather station data, MNC, MCC, mobile brand, elevation, usage type, address type, advertisement category and proxy data with an IP address. It supports both IPv4 and IPv6 address lookup.

In addition, this package provides WHOIS lookup api that helps users to obtain domain information, WHOIS record, by using a domain name. The WHOIS API returns a comprehensive WHOIS data such as creation date, updated date, expiration date, domain age, the contact information of the registrant, mailing address, phone number, email address, nameservers the domain is using and much more.

This package requires API key to function. You may sign up for a free API key at https://www.ip2location.io/pricing.


Installation
============
```

go get github.com/ip2location/ip2location-io-go/ip2locationio

```


Usage Example
============
### Lookup IP Address Geolocation Data
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
	fmt.Printf("TimeZoneInfo.Sunrise => %+v\n", res.TimeZoneInfo.Sunrise)
	fmt.Printf("TimeZoneInfo.Sunset => %+v\n", res.TimeZoneInfo.Sunset)
	
	fmt.Printf("Geotargeting.Metro => %+v\n", res.Geotargeting.Metro)
	
	fmt.Printf("Proxy.LastSeen => %+v\n", res.Proxy.LastSeen)
	fmt.Printf("Proxy.ProxyType => %+v\n", res.Proxy.ProxyType)
	fmt.Printf("Proxy.Threat => %+v\n", res.Proxy.Threat)
	fmt.Printf("Proxy.Provider => %+v\n", res.Proxy.Provider)
}
```

### Lookup Domain Information
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


Response Parameter
============
### IP Geolocation Lookup function
| Parameter | Type | Description |
|---|---|---|
|IP|string|IP address.|
|CountryCode|string|Two-character country code based on ISO 3166.|
|CountryName|string|Country name based on ISO 3166.|
|RegionName|string|Region or state name.|
|CityName|string|City name.|
|Latitude|double|City latitude. Defaults to capital city latitude if city is unknown.|
|Longitude|double|City longitude. Defaults to capital city longitude if city is unknown.|
|ZipCode|string|ZIP/Postal code.|
|TimeZone|string|UTC time zone (with DST supported).|
|Asn|string|Autonomous system number (ASN).|
|AS|string|Autonomous system (AS) name.|
|Isp|string|Internet Service Provider or company's name.|
|Domain|string|Internet domain name associated with IP address range.|
|NetSpeed|string|Internet connection type. DIAL = dial-up, DSL = broadband/cable/fiber/mobile, COMP = company/T1|
|IddCode|string|The IDD prefix to call the city from another country.|
|AreaCode|string|A varying length number assigned to geographic areas for calls between cities.|
|WeatherStationCode|string|The special code to identify the nearest weather observation station.|
|WeatherStationName|string|The name of the nearest weather observation station.|
|Mcc|string|Mobile Country Codes (MCC) as defined in ITU E.212 for use in identifying mobile stations in wireless telephone networks, particularly GSM and UMTS networks.|
|Mnc|string|Mobile Network Code (MNC) is used in combination with a Mobile Country Code (MCC) to uniquely identify a mobile phone operator or carrier.|
|MobileBrand|string|Commercial brand associated with the mobile carrier.|
|Elevation|integer|Average height of city above sea level in meters (m).|
|UsageType|string|Usage type classification of ISP or company.|
|AddressType|string|IP address types as defined in Internet Protocol version 4 (IPv4) and Internet Protocol version 6 (IPv6).|
|Continent.Name|string|Continent name.|
|Continent.Code|string|Two-character continent code.|
|Continent.Hemisphere|array|The hemisphere of where the country located. The data in array format with first item indicates (north/south) hemisphere and second item indicates (east/west) hemisphere information.|
|Continent.Translation|object|Translation data based on the given lang code.|
|District|string|District or county name.|
|Country.Name|string|Country name based on ISO 3166.|
|Country.Alpha3Code|string|Three-character country code based on ISO 3166.|
|Country.NumericCode|string|Three-character country numeric code based on ISO 3166.|
|Country.Demonym|string|Native of the country.|
|Country.Flag|string|URL of the country flag image.|
|Country.Capital|string|Capital of the country.|
|Country.TotalArea|integer|Total area in km2.|
|Country.Population|integer|Population of the country.|
|Country.Currency|object|Currency of the country.|
|Country.Language|object|Language of the country.|
|Country.Tld|string|Country-Code Top-Level Domain.|
|Country.Translation|object|Translation data based on the given lang code.|
|Region.Name|string|Region or state name.|
|Region.Code|string|ISO3166-2 code.|
|Region.Translation|object|Translation data based on the given lang code.|
|City.Name|string| City name.|
|City.Translation|object|Translation data based on the given lang code.|
|TimeZoneInfo.Olson|string|Time zone in Olson format.|
|TimeZoneInfo.CurrentTime|string|Current time in ISO 8601 format.|
|TimeZoneInfo.GmtOffset|integer|GMT offset value in seconds.|
|TimeZoneInfo.IsDst|boolean|Indicate if the time zone value is in DST.|
|TimeZoneInfo.Sunrise|string|Time of sunrise. (hh:mm format in local time, i.e, 07:47)|
|TimeZoneInfo.Sunset|string|Time of sunset. (hh:mm format in local time, i.e 19:50)|
|Geotargeting.Metro|string|Metro code based on zip/postal code.|
|AdsCategory|string|The domain category code based on IAB Tech Lab Content Taxonomy.|
|AdsCategoryName|string|The domain category based on IAB Tech Lab Content Taxonomy. These categories are comprised of Tier-1 and Tier-2 (if available) level categories widely used in services like advertising, Internet security and filtering appliances.|
|IsProxy|boolean|Whether is a proxy or not.|
|Proxy.LastSeen|integer|Proxy last seen in days.|
|Proxy.ProxyType|string|Type of proxy.|
|Proxy.Threat|string|Security threat reported.|
|Proxy.Provider|string|Name of VPN provider if available.|

```json
{
    "ip":"8.8.8.8",
    "country_code":"US",
    "country_name":"United States of America",
    "region_name":"California",
    "city_name":"Mountain View",
    "latitude":37.405992,
    "longitude":-122.078515,
    "zip_code":"94043",
    "time_zone":"-07:00",
    "asn":"15169",
    "as":"Google LLC",
    "isp":"Google LLC",
    "domain":"google.com",
    "net_speed":"T1",
    "idd_code":"1",
    "area_code":"650",
    "weather_station_code":"USCA0746",
    "weather_station_name":"Mountain View",
    "mcc":"-",
    "mnc":"-",
    "mobile_brand":"-",
    "elevation":32,
    "usage_type":"DCH",
    "address_type":"Anycast",
    "continent":{
        "name":"North America",
        "code":"NA",
        "hemisphere":["north","west"],
        "translation":{
            "lang":"ko",
            "value":"북아메리카"
        }
    },
    "district": "Santa Clara County",
    "country":{
        "name":"United States of America",
        "alpha3_code":"USA",
        "numeric_code":840,
        "demonym":"Americans",
        "flag":"https://cdn.ip2location.io/assets/img/flags/us.png",
        "capital":"Washington, D.C.",
        "total_area":9826675,
        "population":331002651,
        "currency":{
            "code":"USD",
            "name":"United States Dollar",
            "symbol":"$"
        },
        "language":{
            "code":"EN",
            "name":"English"
        },
        "tld":"us",
        "translation":{
            "lang":"ko",
            "value":"미국"
        }
    },
    "region":{
        "name":"California",
        "code":"US-CA",
        "translation":{
            "lang":"ko",
            "value":"캘리포니아주"
        }
    },
    "city":{
        "name":"Mountain View",
        "translation":{
            "lang":null,
            "value":null
        }
    },
    "time_zone_info":{
        "olson":"America/Los_Angeles",
        "current_time":"2022-04-18T23:41:57-07:00",
        "gmt_offset":-25200,
        "is_dst":true,
        "sunrise":"06:27",
        "sunset":"19:47"
    },
    "geotargeting":{
        "metro":"807"
    },
    "ads_category":"IAB19",
    "ads_category_name":"Technology & Computing",
    "is_proxy":false,
    "proxy":{
        "last_seen":18,
        "proxy_type":"DCH",
        "threat":"-",
        "provider":"-"
    }
}
```

### Domain WHOIS Lookup function
| Parameter | Type | Description |
|---|---|---|
|Domain|string|Domain name.|
|DomainID|string|Domain name ID.|
|Status|string|Domain name status.|
|CreateDate|string|Domain name creation date.|
|UpdateDate|string|Domain name updated date.|
|ExpireDate|string|Domain name expiration date.|
|DomainAge|integer|Domain name age in day(s).|
|WhoisServer|string|WHOIS server name.|
|Registrar.IanaID|string|Registrar IANA ID.|
|Registrar.Name|string|Registrar name.|
|Registrar.Url|string|Registrar URL.|
|Registrant.Name|string|Registrant name.|
|Registrant.Organization|string|Registrant organization.|
|Registrant.StreetAddress|string|Registrant street address.|
|Registrant.City|string|Registrant city.|
|Registrant.Region|string|Registrant region.|
|Registrant.ZipCode|string|Registrant ZIP Code.|
|Registrant.Country|string|Registrant country.|
|Registrant.Phone|string|Registrant phone number.|
|Registrant.Fax|string|Registrant fax number.|
|Registrant.Email|string|Registrant email address.|
|Admin.Name|string|Admin name.|
|Admin.Organization|string|Admin organization.|
|Admin.StreetAddress|string|Admin street address.|
|Admin.City|string|Admin city.|
|Admin.Region|string|Admin region.|
|Admin.ZipCode|string|Admin ZIP Code.|
|Admin.Country|string|Admin country.|
|Admin.Phone|string|Admin phone number.|
|Admin.Fax|string|Admin fax number.|
|Admin.Email|string|Admin email address.|
|Tech.Name|string|Tech name.|
|Tech.Organization|string|Tech organization.|
|Tech.StreetAddress|string|Tech street address.|
|Tech.City|string|Tech city.|
|Tech.Region|string|Tech region.|
|Tech.ZipCode|string|Tech ZIP Code.|
|Tech.Country|string|Tech country.|
|Tech.Phone|string|Tech phone number.|
|Tech.Fax|string|Tech fax number.|
|Tech.Email|string|Tech email address.|
|Billing.Name|string|Billing name.|
|Billing.Organization|string|Billing organization.|
|Billing.StreetAddress|string|Billing street address.|
|Billing.City|string|Billing city.|
|Billing.Region|string|Billing region.|
|Billing.ZipCode|string|Billing ZIP Code.|
|Billing.Country|string|Billing country.|
|Billing.Phone|string|Billing phone number.|
|Billing.Fax|string|Billing fax number.|
|Billing.Email|string|Billing email address.|
|Nameservers|array|Name servers|

```json
{
    "domain": "locaproxy.com",
    "domain_id": "1710914405_DOMAIN_COM-VRSN",
    "status": "clientTransferProhibited https://icann.org/epp#clientTransferProhibited",
    "create_date": "2012-04-03T02:34:32Z",
    "update_date": "2021-12-03T02:54:57Z",
    "expire_date": "2024-04-03T02:34:32Z",
    "domain_age": 3863,
    "whois_server": "whois.godaddy.com",
    "registrar": {
        "iana_id": "146",
        "name": "GoDaddy.com, LLC",
        "url": "https://www.godaddy.com"
    },
    "registrant": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "admin": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "tech": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "billing": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "nameservers": ["vera.ns.cloudflare.com", "walt.ns.cloudflare.com"]
}
```


LICENCE
=====================
See the LICENSE file.
