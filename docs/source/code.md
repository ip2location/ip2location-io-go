# IP2Location.io Go API

## Configuration Class

```{py:function} OpenConfiguration(apikey)
Configuration registry. Set the IP2Location.io API key.

:param string apikey: (Required) The IP2Location.io API key.
:return: Returns the configuration in object.
:rtype: object
```

## IPGeolocation Class

```{py:function} OpenIPGeolocation(config)
Initialize the IPGeolocation class.

:param object config: (Required) The IP2Location.io Configuration object returned by Configuration class.
```

```{py:function} LookUp(ip, language)
Retrieve geolocation information for an IP address.

:param string ip: (Required) The IP address (IPv4 or IPv6).
:param string language: (Optional) Translation information(ISO639-1). The translation is only applicable for continent, country, region and city name.
:::{note}
Note: This parameter is only available for Plus and Security plan only.
:::
:return: Returns the geolocation information in JSON. Refer below table for the fields avaliable in the JSON
:rtype: json

**RETURN FIELDS**
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
|time_zone_info.Abbreviation|string|The time zone abbreviation of the Olson time zone, for example EST and EEST.|
|time_zone_info.DstStartDate|string|The date (UTC) of Daylight Saving Time (DST) begins.|
|time_zone_info.DstEndDate|string|The date (UTC) of Daylight Saving Time (DST) ends.|
|TimeZoneInfo.Sunrise|string|Time of sunrise. (hh:mm format in local time, i.e, 07:47)|
|TimeZoneInfo.Sunset|string|Time of sunset. (hh:mm format in local time, i.e 19:50)|
|Geotargeting.Metro|string|Metro code based on zip/postal code.|
|AdsCategory|string|The domain category code based on IAB Tech Lab Content Taxonomy.|
|AdsCategoryName|string|The domain category based on IAB Tech Lab Content Taxonomy. These categories are comprised of Tier-1 and Tier-2 (if available) level categories widely used in services like advertising, Internet security and filtering appliances.|
|IsProxy|boolean|Whether is a proxy or not.|
|FraudScore|integer|Potential risk score (0 - 99) associated with IP address. A higher IP2Proxy Fraud Score indicates a greater likelihood of fraudulent activity and a lower reputation.|
|Proxy.LastSeen|integer|Proxy last seen in days.|
|Proxy.ProxyType|string|Type of proxy.|
|Proxy.Threat|string|Security threat reported.|
|Proxy.Provider|string|Name of VPN provider if available.|
|Proxy.IsVpn|boolean|Anonymizing VPN services.|
|Proxy.IsTor|boolean|Tor Exit Nodes.|
|Proxy.IsDataCenter|boolean|Hosting Provider, Data Center or Content Delivery Network.|
|Proxy.IsPublicProxy|boolean|Public Proxies.|
|Proxy.IsWebProxy|boolean|Web Proxies.|
|Proxy.IsWebCrawler|boolean|Search Engine Robots.|
|Proxy.IsResidentialProxy|boolean|Residential proxies.|
|Proxy.IsSpammer|boolean|Email and forum spammers.|
|Proxy.IsScanner|boolean|Network security scanners.|
|Proxy.IsBotnet|boolean|Malware infected devices.|
```

## DomainWhois Class

**_Please note that this class is no longer being maintained and has been migrated to a new repository. We recommend that you use the [IP2WHOIS Go SDK](https://github.com/ip2whois/ip2whois-go) going forward._**

```{py:function} OpenDomainWhois(config)
Initialize the DomainWhois class.

:param object config: (Required) The IP2Location.io Configuration object returned by Configuration class.
```

```{py:function} LookUp(domain)
Lookup domain WHOIS information.

:param string domain: (Required) The domain name.
:return: IP2WHOIS Domain WHOIS result in JSON. Refer below table for the fields avaliable in the JSON.
:rtype: json

**RETURN FIELDS**
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
```

```{py:function} GetPunycode(domain)
Get Punycode for domain name.

:param string domain: (Required) Domain name.
:return: Returns punycode in text.
:rtype: string
```

```{py:function} GetNormalText(domain)
Get Normal Text.

:param string domain: (Required) The punycode domain name.
:return: Returns normal domain name in text.
:rtype: string
```

```{py:function} GetDomainName(url)
Get domain name from a URL.

:param string url: (Required) The URL that you want to extract domain name from.
:return: Returns the extracted domain name in text.
:rtype: string
```

```{py:function} GetDomainExtension(str)
Get domain extension from a URL or domain.

:param string str: (Required) The URL or domain.
:return: Returns the domain extension in text.
:rtype: string
```

## HostedDomain Class

```{py:function} OpenHostedDomain(config)
Initialize the HostedDomain class.

:param object config: (Required) The IP2Location.io Configuration object returned by Configuration class.
```

```{py:function} LookUp(ip, page)
Retrieve hosted domains information for an IP address.

:param string ip: (Required) The IP address (IPv4 or IPv6).
:param string page: (Required) Page of the result.
:return: Returns the hosted domains information in JSON. Refer below table for the fields avaliable in the JSON
:rtype: json

**RETURN FIELDS**
| Parameter | Type | Description |
|---|---|---|
|IP|string|IP address.|
|TotalDomains|integer|Total number of hosted domains found.|
|Page|integer|Current lookup page.|
|PerPage|integer|Number of domains displayed in the page.|
|TotalPages|integer|Total pages of the hosted domains.|
|Domains|array|Hosted domains of the lookup IP Address.|
```
