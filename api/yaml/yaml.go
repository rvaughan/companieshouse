package yaml

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type yamlDescriptions struct {
	Descriptions map[string]string `yaml:"description"`
}

type StrEnum map[string]string

var FilingHistoryDescriptions StrEnum

func getDescriptions(f string) (*yamlDescriptions, error) {
	d := &yamlDescriptions{}
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func init() {
	FilingHistoryDescriptions = make(StrEnum)
	d, err := getDescriptions("filing_history_descriptions.yml")
	if err != nil {
		log.Fatal(err)
	}
	FilingHistoryDescriptions = d.Descriptions
}
