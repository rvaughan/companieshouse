package main

import (
	"github.com/BalkanTech/companieshouse/api"
	"log"
	"os"
	"time"
)

func main() {
	ch := companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))
	now := time.Now()
	c, err := ch.GetCompany("06043462")
	if err != nil {
		log.Fatal(err)
	}

	/*f, err := c.GetFilings()
	if err != nil {
		log.Fatal(err)
	}

	d := f.Items[len(f.Items)-1]
	err = c.DownloadDocument(d, "test.pdf")
	if err != nil {
		log.Fatal(err)
	}*/

	log.Printf("Request took: %s", time.Since(now))
	log.Println(c.Officers)
	log.Println(c.Charges)
	log.Println(c.InsolvencyHistory)
}
