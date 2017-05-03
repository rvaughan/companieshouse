package main

import (
	"github.com/BalkanTech/companieshouse/api"
	"log"
	"os"
)

func main() {
	ch := companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))
	c, err := ch.GetCompany("09999801")
	if err != nil {
		log.Fatal(err)
	}

	f, err := c.GetFilings()
	if err != nil {
		log.Fatal(err)
	}

	d := f.Filings[len(f.Filings)-1]
	err = c.DownloadDocument(&d, "test.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
