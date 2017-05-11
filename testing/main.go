package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Descriptor struct {
	Description map[string]string `yaml:"description"`
}

func main() {
	d := Descriptor{}
	//D := &d{d: make(map[string]string)}
	yamlFile, err := ioutil.ReadFile("../api/yaml/filing_history_descriptions.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println(d.Description["liquidation-change-of-membership-of-creditors-or-liquidation-committee"])
}


/*func main() {
	ch := companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))
	now := time.Now()
	c, err := ch.GetCompany("06043462")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Request took: %s", time.Since(now))
	log.Println(c.CompanyName)
	log.Println(c.Officers)
	log.Println(c.Filings)
	log.Println(c.Charges)
	log.Println(c.InsolvencyHistory)
} */
