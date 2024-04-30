package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	initDatabase()
	migrate()

	fetchTop100Domains()
	fetchDNSRecords()
	fetchInternetDBResults()

	e := echo.New()

	e.GET("/domains", getDomains)
	e.GET("/dnsrecords", getDNSRecords)
	e.GET("/internetdbresults", getInternetDBResults)

	e.Start(":8080")
}
