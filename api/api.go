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

const(
	defaultURL         = "https://api.companieshouse.gov.uk"
	defaultDocumentURL = "https://document-api.companieshouse.gov.uk"
)

// ChAPI implements access to the Companies House REST API
type ChAPI struct {
	apiKey              string
	apiURL              string
	overWriteDefaultURL bool
}

// setAPIURL allows to override the apiURL value of ChAPI. Such overriding is meant for proper functioning of unit tests
// and should normally not be used in production
func (a *ChAPI) setAPIURL(u string) {
	a.apiURL = u
	a.overWriteDefaultURL = true
}

//SetAPIKey sets the apiKey value for the ChAPI struct and sets an inital value to the Companies House REST url
func (a *ChAPI) SetAPIKey(k string) {
	a.apiKey = k
	a.apiURL = defaultURL
	a.overWriteDefaultURL = false
}

func (a *ChAPI) constructURL(path string) string {
	return a.apiURL + path
}

func (a *ChAPI) prepareRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return req, err
	}
	req.SetBasicAuth(a.apiKey, "")

	return req, err
}

func (a *ChAPI) getResponse(url string, contentType string) (*http.Response, error) {
	if a.overWriteDefaultURL {
		url = strings.Replace(url, defaultURL, a.apiURL, -1)
		url = strings.Replace(url, defaultDocumentURL, a.apiURL, -1)
	}
	req, err := a.prepareRequest(url)
	if err != nil {
		return &http.Response{}, err
	}

	if contentType != "" {
		req.Header.Set("Accept", contentType)
	}
	client := &http.Client{}

	return client.Do(req)
}

// CallAPI will make a call to the companies house API to the provided path
// Set fullURL to true if the path is a full URL
// Set contentType to the desired content type. If contentType is an empty string then ContentTypeJSON will be used.
// Returns the response's body as a slice of bytes and an error
func (a *ChAPI) CallAPI(path string, fullURL bool, contentType string) ([]byte, error) {
	var url string

	if fullURL {
		url = path
	} else {
		url = a.constructURL(path)
	}

	if contentType == "" {
		contentType = ContentTypeJSON
	}

	resp, err := a.getResponse(url, contentType)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}