package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "strings"
    "io"
)

type IPInfo struct {
    IP       string `json:"ip"`
    Hostname string `json:"hostname"`
    City     string `json:"city"`
    Region   string `json:"region"`
    Country  string `json:"country"`
    Loc      string `json:"loc"`
    Org      string `json:"org"`
    Postal   string `json:"postal"`
    Timezone string `json:"timezone"`
    ASN      struct {
        ASN    string `json:"asn"`
        Name   string `json:"name"`
        Domain string `json:"domain"`
        Route  string `json:"route"`
        Type   string `json:"type"`
    } `json:"asn"`
    Company struct {
        Name   string `json:"name"`
        Domain string `json:"domain"`
        Type   string `json:"type"`
    } `json:"company"`
    Privacy struct {
        VPN     bool   `json:"vpn"`
        Proxy   bool   `json:"proxy"`
        Tor     bool   `json:"tor"`
        Relay   bool   `json:"relay"`
        Hosting bool   `json:"hosting"`
        Service string `json:"service"`
    } `json:"privacy"`
    Abuse struct {
        Address string `json:"address"`
        Country string `json:"country"`
        Email   string `json:"email"`
        Name    string `json:"name"`
        Network string `json:"network"`
        Phone   string `json:"phone"`
    } `json:"abuse"`
    Domains struct {
        Page    int      `json:"page"`
        Total   int      `json:"total"`
        Domains []string `json:"domains"`
    } `json:"domains"`
}

func fetchIPInfo() (IPInfo, error) {
    resp, err := http.Get("https://ipinfo.io/widget")
    if err != nil {
        return IPInfo{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return IPInfo{}, err
    }

    var info IPInfo
    err = json.Unmarshal(body, &info)
    if err != nil {
        return IPInfo{}, err
    }

    return info, nil
}

func generateHTML(info IPInfo) string {
    html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>IP Information</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        .accordion-button { font-weight: bold; }
        .info-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 1rem; }
        .info-item { margin-bottom: 0; }
    </style>
</head>
<body class="bg-light">
    <div class="container my-5">
        <h1 class="text-center mb-4">IP Information for ` + info.IP + `</h1>

        <div class="accordion" id="ipInfoAccordion">
            <div class="accordion-item">
                <h2 class="accordion-header" id="generalInfo">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseGeneral">
                        General Information
                    </button>
                </h2>
                <div id="collapseGeneral" class="accordion-collapse collapse show" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>IP:</strong> ` + info.IP + `</p>
                            <p class="info-item"><strong>Hostname:</strong> ` + info.Hostname + `</p>
                            <p class="info-item"><strong>City:</strong> ` + info.City + `</p>
                            <p class="info-item"><strong>Region:</strong> ` + info.Region + `</p>
                            <p class="info-item"><strong>Country:</strong> ` + info.Country + `</p>
                            <p class="info-item"><strong>Location:</strong> ` + info.Loc + `</p>
                            <p class="info-item"><strong>Organization:</strong> ` + info.Org + `</p>
                            <p class="info-item"><strong>Postal:</strong> ` + info.Postal + `</p>
                            <p class="info-item"><strong>Timezone:</strong> ` + info.Timezone + `</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="asnInfo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseASN">
                        ASN Information
                    </button>
                </h2>
                <div id="collapseASN" class="accordion-collapse collapse" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>ASN:</strong> ` + info.ASN.ASN + `</p>
                            <p class="info-item"><strong>Name:</strong> ` + info.ASN.Name + `</p>
                            <p class="info-item"><strong>Domain:</strong> ` + info.ASN.Domain + `</p>
                            <p class="info-item"><strong>Route:</strong> ` + info.ASN.Route + `</p>
                            <p class="info-item"><strong>Type:</strong> ` + info.ASN.Type + `</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="companyInfo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseCompany">
                        Company Information
                    </button>
                </h2>
                <div id="collapseCompany" class="accordion-collapse collapse" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>Name:</strong> ` + info.Company.Name + `</p>
                            <p class="info-item"><strong>Domain:</strong> ` + info.Company.Domain + `</p>
                            <p class="info-item"><strong>Type:</strong> ` + info.Company.Type + `</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="privacyInfo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapsePrivacy">
                        Privacy Information
                    </button>
                </h2>
                <div id="collapsePrivacy" class="accordion-collapse collapse" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>VPN:</strong> ` + strconv.FormatBool(info.Privacy.VPN) + `</p>
                            <p class="info-item"><strong>Proxy:</strong> ` + strconv.FormatBool(info.Privacy.Proxy) + `</p>
                            <p class="info-item"><strong>Tor:</strong> ` + strconv.FormatBool(info.Privacy.Tor) + `</p>
                            <p class="info-item"><strong>Relay:</strong> ` + strconv.FormatBool(info.Privacy.Relay) + `</p>
                            <p class="info-item"><strong>Hosting:</strong> ` + strconv.FormatBool(info.Privacy.Hosting) + `</p>
                            <p class="info-item"><strong>Service:</strong> ` + info.Privacy.Service + `</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="abuseInfo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseAbuse">
                        Abuse Information
                    </button>
                </h2>
                <div id="collapseAbuse" class="accordion-collapse collapse" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>Address:</strong> ` + info.Abuse.Address + `</p>
                            <p class="info-item"><strong>Country:</strong> ` + info.Abuse.Country + `</p>
                            <p class="info-item"><strong>Email:</strong> ` + info.Abuse.Email + `</p>
                            <p class="info-item"><strong>Name:</strong> ` + info.Abuse.Name + `</p>
                            <p class="info-item"><strong>Network:</strong> ` + info.Abuse.Network + `</p>
                            <p class="info-item"><strong>Phone:</strong> ` + info.Abuse.Phone + `</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="domainsInfo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseDomains">
                        Domains Information
                    </button>
                </h2>
                <div id="collapseDomains" class="accordion-collapse collapse" data-bs-parent="#ipInfoAccordion">
                    <div class="accordion-body">
                        <div class="info-grid">
                            <p class="info-item"><strong>Page:</strong> ` + strconv.Itoa(info.Domains.Page) + `</p>
                            <p class="info-item"><strong>Total:</strong> ` + strconv.Itoa(info.Domains.Total) + `</p>
                            <p class="info-item"><strong>Domains:</strong> ` + strings.Join(info.Domains.Domains, ", ") + `</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
`
    return html
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		info, err := fetchIPInfo()
		if err != nil {
			http.Error(w, "Error fetching IP information", http.StatusInternalServerError)
			return
		}

		html := generateHTML(info)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}