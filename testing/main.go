package main

import (

)
import (
	"github.com/BalkanTech/companieshouse/api"
	"os"
	"fmt"
	"log"
)

func main() {
	ch := companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))
	r, err := ch.Search("51 Portland Road", 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *r)
}

