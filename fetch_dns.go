package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func fetchDNSRecords() {
	file, err := os.Open("top_100_domains.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		nsRecords, err := net.LookupNS(domain)
		if err != nil {
			fmt.Println("Failed to lookup NS records for", domain, ":", err)
			continue
		}

		for _, ns := range nsRecords {
			dnsRecord := DNSRecord{
				Domain: Domain{Name: domain},
				Type:   "NS",
				Value:  ns.Host,
			}
			db.Create(&dnsRecord)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("DNS records fetched and saved.")
}
