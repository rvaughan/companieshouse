/*
Golang Companies House REST service API
Copyright (C) 2017, Balkan C & T OOD

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
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Content Types
const (
	ContentTypeJSON      = "application/json"
	ContentTypePDF       = "application/pdf"
	ContentTypeTextPlain = "text/plain; charset=utf-8"
)

const (
	defaultURL         = "https://api.companieshouse.gov.uk"
	defaultDocumentURL = "https://document-api.companieshouse.gov.uk"
)

// API implements access to the Companies House REST API
type API struct {
	apiKey              string
	apiURL              string
	overWriteDefaultURL bool
}

type QueryParams map[string]interface{}

// setAPIURL allows to override the apiURL value of API. Such overriding is meant for proper functioning of unit tests
// and should normally not be used in production
func (a *API) setAPIURL(u string) {
	a.apiURL = u
	a.overWriteDefaultURL = true
}

//SetAPIKey sets the apiKey value for the API struct and sets an inital value to the Companies House REST url
func (a *API) SetAPIKey(k string) {
	a.apiKey = k
	a.apiURL = defaultURL
	a.overWriteDefaultURL = false
}

func (a *API) constructURL(path string) string {
	return a.apiURL + path
}

func (a *API) prepareRequest(url string, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return req, err
	}
	req.SetBasicAuth(a.apiKey, "")

	return req, err
}

func (a *API) getResponse(url string, method string, params QueryParams, contentType string) (*http.Response, error) {
	if a.overWriteDefaultURL {
		url = strings.Replace(url, defaultURL, a.apiURL, -1)
		url = strings.Replace(url, defaultDocumentURL, a.apiURL, -1)
	}
	req, err := a.prepareRequest(url, method)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()

	if contentType != "" {
		req.Header.Set("Accept", contentType)
	}

	client := &http.Client{}
	return client.Do(req)
}

// CallAPI makes a call to the companies house API via the provided path
// Set fullURL to true if the path is a full URL
// Set contentType to the desired content type. If contentType is an empty string then ContentTypeJSON will be used.
// Returns the response's body as a slice of bytes and an error
func (a *API) CallAPI(path string, params QueryParams, fullURL bool, contentType string) ([]byte, error) {
	var url string

	if fullURL {
		url = path
	} else {
		url = a.constructURL(path)
	}

	if contentType == "" {
		contentType = ContentTypeJSON
	}

	resp, err := a.getResponse(url, "GET", params, contentType)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func CompaniesHouseAPI(key string) *API {
	api := &API{}
	api.SetAPIKey(key)
	return api
}

// Todo: Officer Disqualifications
// Todo: UK Establshment companies
// Todo: Persons with significant control
// Todo: Document downloads
