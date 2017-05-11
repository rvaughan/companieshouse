package main

import (
	"github.com/BalkanTech/companieshouse/api"
	"time"
	"log"
	"os"
)

func main() {
	ch := companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))
	now := time.Now()
	c, err := ch.GetCompany("09999801")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Request took: %s", time.Since(now))
	log.Println(c.CompanyName)
	log.Println(c.Officers)
	log.Println(c.Filings)
	log.Println(c.Charges)
	log.Println(c.InsolvencyHistory)
}
