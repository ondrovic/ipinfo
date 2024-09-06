package main

import (
	"fmt"
	"ipinfo-web/internal/utils"
	"net/http"
)

const (
	baseApiUrl = "http://ip-api.com/json"
	fields     = "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,offset,currency,isp,org,as,asname,reverse,mobile,proxy,hosting,query"
)

func main() {

    // gran the ip info
	rawJson, err := utils.FetchIPInfo(baseApiUrl, fields)

	if err != nil {
		fmt.Printf("error getting IP info: %v", err)
		return
	}

    // output in json format
    http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application.json")
		w.Write([]byte(rawJson))

	})

    // output in bootstrap html template populated by data
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		data, err := utils.ParseIPInfo(rawJson)
		if err != nil {
			http.Error(w, "Error parsing data", http.StatusInternalServerError)
			return
		}

		output, err := utils.GenerateHTMLOutput(data)

		if err != nil {
			http.Error(w, "Error populating template with data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(output))

	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
