package main

func migrate() {
	db.AutoMigrate(&domain.Domain{})
	db.AutoMigrate(&dnsrecord.DNSRecord{})
	db.AutoMigrate(&internetdbresult.InternetDBResult{})
}
