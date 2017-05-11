package yaml

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"path"
	"runtime"
	"errors"
)

type yamlDescriptions struct {
	Descriptions map[string]string `yaml:"description"`
}

type StrEnum map[string]string

var FilingHistoryDescriptions StrEnum

func getDescriptions(f string) (*yamlDescriptions, error) {
	d := &yamlDescriptions{}
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("No caller information")
	}
	file := path.Join(path.Dir(filename), f)
	yamlFile, err := ioutil.ReadFile(file)
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
	d, err := getDescriptions("./filing_history_descriptions.yml")
	if err != nil {
		log.Fatal(err)
	}
	FilingHistoryDescriptions = d.Descriptions
}
