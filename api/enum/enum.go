package enum

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"reflect"
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

func openFile(f string) ([]byte, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("No caller information")
	}
	file := path.Join(path.Dir(filename), f)
	return ioutil.ReadFile(file)
}

func loadFile(f string, d interface{}) error {
	if reflect.ValueOf(d).Kind() != reflect.Ptr {
		return errors.New("d should be a pointer")
	}

	yamlFile, err := openFile(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := loadFile("./filing_history_descriptions.yml", &FilingHistoryDescriptions)
	if err != nil {
		log.Fatal(err)
	}

	err = loadFile("./constants.yml", &Constants)
	if err != nil {
		log.Fatal(err)
	}

	err = loadFile("./mortgage_descriptions.yml", &MortgageDescriptions)
	if err != nil {
		log.Fatal(err)
	}
}
