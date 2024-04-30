package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func fetchTop100Domains() {
	resp, err := http.Get("https://builtwith.com/top-1m")
	if err != nil {
		fmt.Println("Failed to fetch top 1M domains:", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("top_100_domains.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(resp.Body)
	count := 0
	for scanner.Scan() {
		if count < 100 {
			domain := strings.TrimSpace(scanner.Text())
			file.WriteString(domain + "\n")
			count++
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response body:", err)
	}
	fmt.Println("Fetched top 100 domains and saved to top_100_domains.txt.")
}
