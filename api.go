package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getDomains(c echo.Context) error {
	var domains []Domain
	db.Find(&domains)
	return c.JSON(http.StatusOK, domains)
}

func getDNSRecords(c echo.Context) error {
	var dnsRecords []DNSRecord
	db.Find(&dnsRecords)
	return c.JSON(http.StatusOK, dnsRecords)
}

func getInternetDBResults(c echo.Context) error {
	var internetDBResults []InternetDBResult
	db.Find(&internetDBResults)
	return c.JSON(http.StatusOK, internetDBResults)
}
