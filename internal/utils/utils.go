package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"ipinfo-web/internal/templates"
	"ipinfo-web/internal/types"
	"net/http"
	"strings"
)

// FetchIPInfo performs an HTTP GET request to the specified URI with the given fields
// as query parameters and retrieves the response body.
//
// Parameters:
//   - uri: The base URL to which the HTTP GET request is made.
//   - fields: A query parameter specifying which fields to include in the response.
//
// Returns:
//   - []byte: The body of the HTTP response as a byte slice.
//   - error: An error if the HTTP request fails or if reading the response body fails.
//
// Example:
//
//	uri := "https://example.com/api/ipinfo"
//	fields := "ip,city,country"
//	data, err := FetchIPInfo(uri, fields)
//	if err != nil {
//	    log.Fatalf("Failed to fetch IP info: %v", err)
//	}
//	fmt.Println(string(data))
//
// Notes:
//   - Ensure that the fields parameter is properly URL-encoded if it contains special characters.
//   - Consider adding timeouts or context management for production use to handle potential delays in network requests.
func FetchIPInfo(uri, fields string) ([]byte, error) {
	fullUrl := fmt.Sprintf("%s?fields=%s", uri, fields)

	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// ParseIPInfo parses raw JSON data into an IpInfo structure.
//
// This function takes a byte slice containing raw JSON data and attempts to
// unmarshal it into an IpInfo object. If the JSON data is successfully parsed,
// it returns the IpInfo object and a nil error. If an error occurs during the
// unmarshalling process, it returns the default IpInfo value and the encountered
// error.
//
// Parameters:
// - rawJson: A byte slice containing the raw JSON data to be parsed.
//
// Returns:
//   - IpInfo: The parsed IpInfo structure derived from the JSON data.
//   - error: An error object that will be non-nil if the unmarshalling fails,
//     otherwise it will be nil.
//
// Example usage:
//
//	rawJson := []byte(`{"ip": "192.168.1.1", "hostname": "example.local"}`)
//	ipInfo, err := ParseIPInfo(rawJson)
//	if err != nil {
//	    log.Fatalf("Failed to parse IP info: %v", err)
//	}
//	fmt.Println(ipInfo)
func ParseIPInfo(rawJson []byte) (types.IpInfo, error) {
	var ipInfos types.IpInfo
	err := json.Unmarshal(rawJson, &ipInfos)
	return ipInfos, err
}

// GenerateHTMLOutput generates an HTML output string from an IpInfo structure.
//
// This function converts the provided IpInfo data into a JSON representation and
// then uses an embedded HTML template to generate an HTML output string. The JSON
// data is safely injected into the template using a custom function. If any errors
// occur during the marshaling of data, reading the template file, or parsing the
// template, the function returns an empty string and an appropriate error.
//
// Parameters:
// - data: An IpInfo structure containing the data to be included in the HTML output.
//
// Returns:
//   - string: The generated HTML output as a string.
//   - error: An error object that will be non-nil if any error occurs during the
//     processing, otherwise it will be nil.
//
// Example usage:
//
//	ipInfo := types.IpInfo{IP: "192.168.1.1", Hostname: "example.local"}
//	htmlOutput, err := GenerateHTMLOutput(ipInfo)
//	if err != nil {
//	    log.Fatalf("Failed to generate HTML output: %v", err)
//	}
//	fmt.Println(htmlOutput)
func GenerateHTMLOutput(data types.IpInfo) (string, error) {
	// Convert the snippets to JSON
	ipInfoData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling snippets to JSON: %v", err)
	}

	tmplContent, err := templates.HtmlTemplateFS.ReadFile("html/template.html")
	if err != nil {
		return "", fmt.Errorf("error reading embedded HTML template file: %v", err)
	}

	// Create a template with a custom function to safely inject JSON
	tmpl, err := template.New("output").Funcs(template.FuncMap{
		"safeJSON": func(v interface{}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		},
	}).Parse(string(tmplContent))
	if err != nil {
		return "", fmt.Errorf("error parsing HTML template: %v", err)
	}

	// Prepare the data for the template
	templateData := struct {
		IpInfoJson template.JS
	}{
		IpInfoJson: template.JS(ipInfoData),
	}

	var output strings.Builder
	err = tmpl.Execute(&output, templateData)
	if err != nil {
		return "", fmt.Errorf("error executing HTML template: %v", err)
	}

	return output.String(), nil
}
