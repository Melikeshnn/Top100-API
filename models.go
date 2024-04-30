package main

import "gorm.io/gorm"

// Domain modeli
type Domain struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// DNSRecord modeli
type DNSRecord struct {
	gorm.Model
	DomainID uint
	Domain   Domain
	Type     string
	Value    string
}

// InternetDBResult modeli
type InternetDBResult struct {
	gorm.Model
	IP          string
	Info        string
	DNSRecordID uint
}
