/*
Golang Companies House REST service API
Copyright (C) 2016, Balkan Technologies EOOD & Co. KD

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package companieshouse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChAPI is the body for the Companies House API and contains only the apikey
type ChAPI struct {
	apikey string
}

// SetAPIKey func. Takes apikey string.
func (a *ChAPI) SetAPIKey(apikey string) {
	a.apikey = apikey
}

func (a *ChAPI) getRequest(path string, fullURL bool) (*http.Request, error) {
	var err error
	var url string

	apikey := a.apikey
	apiurl := "api.companieshouse.gov.uk/"

	if fullURL {
		fmt.Println("Full url")
		url = path
	} else {
		fmt.Println("No full url")
		url = "https://" + apiurl + path
	}

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	req.SetBasicAuth(apikey, "")

	return req, err
}

func (a *ChAPI) executeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

func (a *ChAPI) makeGetRequest(path string, fullURL bool, contentType string) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = a.getRequest(path, fullURL)
	if err != nil {
		return resp, err
	}

	if contentType != "" {
		req.Header.Set("Acccept", contentType)
	}
	resp, err = a.executeRequest(req)

	return resp, err
}

func (a *ChAPI) callAPI(path string, fullURL bool) ([]byte, error) {
	resp, err := a.makeGetRequest(path, fullURL, "")
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
