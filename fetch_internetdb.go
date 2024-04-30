package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InternetDBAPIResponse struct {
	Info string `json:"info"`
}

func fetchInternetDBResults() {
	var dnsRecords []DNSRecord
	db.Find(&dnsRecords)

	for _, record := range dnsRecords {
		if record.Type == "A" || record.Type == "AAAA" {
			resp, err := http.Get(fmt.Sprintf("https://internetdb.shodan.io/%s", record.Value))
			if err != nil {
				fmt.Println("Failed to query InternetDB API for IP address", record.Value, ":", err)
				continue
			}
			defer resp.Body.Close()

			var apiResponse InternetDBAPIResponse
			if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
				fmt.Println("Failed to decode API response:", err)
				continue
			}

			internetDBResult := InternetDBResult{
				IP:          record.Value,
				Info:        apiResponse.Info,
				DNSRecordID: record.ID,
			}
			db.Create(&internetDBResult)
		}
	}
	fmt.Println("Fetched InternetDB results and saved.")
}
