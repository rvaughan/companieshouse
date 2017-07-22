package enum

import (
	"gopkg.in/yaml.v2"
)

type ENUM map[string]map[string]string
func (y ENUM) Get(s, k string) string {
	v, ok := y[string(s)][string(k)]
	if !ok {
		return ""
	}
	return v
}

var Constants, FilingHistoryDescriptions, MortgageDescriptions ENUM

func init() {
	if err := yaml.Unmarshal([]byte(filingHistoryDescriptionsYAML), &FilingHistoryDescriptions); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal([]byte(constantsYAML), &Constants); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal([]byte(mortgageDescriptionsYAML), &MortgageDescriptions); err != nil {
		panic(err)
	}
}
