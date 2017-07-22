package main

import (
	"github.com/BalkanTech/companieshouse/api"
	"log"
	"os"
	"time"
)

func main() {
	ch := companieshouse.NewAPI(os.Getenv("CH_API_KEY"))
	now := time.Now()
	c, err := ch.GetCompany("OC392483", companieshouse.WithOfficers())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Request took: %s", time.Since(now))
	log.Println(c.CompanyName)
	log.Println(c.SICCodes)

	log.Println(c.Officers)
	//log.Println(c.Filings)
	//log.Println(c.Charges)
	//log.Println(c.InsolvencyHistory)
}
