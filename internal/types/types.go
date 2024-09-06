package types

// IpInfo represents the information related to an IP address.
//
// This struct contains various fields that describe details about an IP address,
// such as its geographical location, timezone, ISP information, and more. Each
// field is tagged with JSON annotations to facilitate marshalling and unmarshalling
// of JSON data.
//
// Fields:
// - Status: A string representing the status of the IP info request (e.g., "success").
// - Message: A string containing any additional message or error details (optional).
// - Continent: The name of the continent where the IP address is located.
// - ContinentCode: The ISO code of the continent where the IP address is located.
// - Country: The name of the country where the IP address is located.
// - CountryCode: The ISO code of the country where the IP address is located.
// - Region: The name of the region or state where the IP address is located.
// - RegionName: The name of the region or state in a human-readable format.
// - City: The city where the IP address is located.
// - District: The district or neighborhood within the city (optional).
// - Zip: The ZIP or postal code of the IP address's location.
// - Lat: The latitude coordinate of the IP address's location.
// - Lon: The longitude coordinate of the IP address's location.
// - Timezone: The timezone of the IP address's location.
// - Offset: The timezone offset in minutes from UTC.
// - Currency: The currency code used in the IP address's location.
// - Isp: The name of the Internet Service Provider (ISP) for the IP address.
// - Org: The organization associated with the IP address.
// - As: The Autonomous System (AS) number associated with the IP address.
// - AsName: The name of the Autonomous System (AS) associated with the IP address.
// - Reverse: The reverse DNS hostname associated with the IP address.
// - Mobile: A boolean indicating whether the IP address is associated with a mobile device.
// - Proxy: A boolean indicating whether the IP address is associated with a proxy server.
// - Hosting: A boolean indicating whether the IP address is associated with a hosting provider.
// - Query: The IP address that was queried.
//
// Example usage:
//
//	ipInfo := IpInfo{
//	    Status:        "success",
//	    Continent:     "Europe",
//	    ContinentCode: "EU",
//	    Country:       "Germany",
//	    CountryCode:   "DE",
//	    Region:        "Berlin",
//	    City:          "Berlin",
//	    Lat:           52.52,
//	    Lon:           13.405,
//	    Timezone:      "Europe/Berlin",
//	    Offset:        120,
//	    Currency:      "EUR",
//	    Isp:           "Deutsche Telekom",
//	    Org:           "Deutsche Telekom AG",
//	    As:            "AS3320",
//	    AsName:        "Deutsche Telekom AG",
//	    Reverse:       "example.de",
//	    Mobile:        false,
//	    Proxy:         false,
//	    Hosting:       false,
//	    Query:         "8.8.8.8",
//	}
type IpInfo struct {
	Status        string  `json:"status"`
	Message       string  `json:"message,omitempty"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district,omitempty"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	AsName        string  `json:"asname"`
	Reverse       string  `json:"reverse"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
	Query         string  `json:"query"`
}
