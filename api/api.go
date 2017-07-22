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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Content Types
const (
	ContentTypeJSON      = "application/json"
	ContentTypePDF       = "application/pdf"
	ContentTypeTextPlain = "text/plain; charset=utf-8"
	defaultURL           = "https://api.companieshouse.gov.uk"
)

// API implements access to the Companies House REST API
type API struct {
	apiKey string
	url    string
}

type QueryParams map[string]interface{}

func (a *API) setURL(url string) {
	a.url = url
}

func (a *API) constructURL(path string) string {
	p := strings.ToLower(path)
	if strings.HasPrefix(p, "http") || strings.HasPrefix(p, "https") {
		return path
	}
	return a.url + path
}

func (a *API) prepareRequest(url string, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, a.constructURL(url), nil)
	if err != nil {
		return nil, fmt.Errorf("Prepare request: %q", err)
	}
	req.SetBasicAuth(a.apiKey, "")

	return req, nil
}

func (a *API) getResponse(url string, method string, params QueryParams, contentType string) (*http.Response, error) {
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
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Get response do: %q", err)
	}

	return r, nil
}

// CallAPI makes a call to the companies house API via the provided path
// Set fullURL to true if the path is a full URL
// Set contentType to the desired content type. If contentType is an empty string then ContentTypeJSON will be used.
// Returns the response's body as a slice of bytes and an error
func (a *API) CallAPI(path string, params QueryParams, contentType string) ([]byte, error) {
	if contentType == "" {
		contentType = ContentTypeJSON
	}

	resp, err := a.getResponse(path, "GET", params, contentType)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("CallAPI to %q returns status: %q", path, resp.Status)
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("CallAPI read body: %q", err)
	}
	return r, nil
}

type companyOption func(*Company)

func WithOfficers() companyOption {
	return func(c *Company) {
		if c.CompanyNumber != "" {
			o, err := c.api.GetOfficers(c.CompanyNumber)
			if err != nil {
				return
			}
			c.Officers = *o
		}
	}
}

func WithFilings() companyOption {
	return func(c *Company) {
		if c.CompanyNumber != "" {
			f, err := c.api.GetFilingHistory(c.CompanyNumber)
			if err != nil {
				return
			}
			c.FilingHistory = *f
		}
	}
}

func WithCharges() companyOption {
	return func(c *Company) {
		if c.CompanyNumber != "" && c.HasCharges {
			ch, err := c.api.GetCharges(c.CompanyNumber)
			if err != nil {
				return
			}
			c.Charges = *ch
		}
	}
}

func WithInsolvencyHistory() companyOption {
	return func(c *Company) {
		if c.CompanyNumber != "" && c.HasInsolvencyHistory {
			i, err := c.api.GetInsolvencyHistory(c.CompanyNumber)
			if err != nil {
				return
			}
			c.InsolvencyHistory = *i
		}
	}
}

// GetCompany gets the json data for a company from the Companies House REST API
// and returns a new Company and an error
func (a *API) GetCompany(companyNumber string, options ...companyOption) (*Company, error) {
	c := Company{api: a}
	// Fetch details
	resp, err := a.CallAPI("/company/"+companyNumber, nil, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &c)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option(&c)
	}

	return &c, nil
}

// GetCharges gets the json data for a company's charges from the Companies House REST API
// and returns a new ChargesResponse and an error
func (a *API) GetCharges(c string) (*ChargesResponse, error) {
	ch := ChargesResponse{}

	resp, err := a.CallAPI("/company/"+c+"/charges", nil, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &ch)
	if err != nil {
		return nil, err
	}

	return &ch, nil
}

// GetInsolvencyDetails gets the json data for a company's insolvency details from the Companies House REST API
// and returns a new InsolvenciesResponse and an error
func (a *API) GetInsolvencyHistory(c string) (*InsolvencyHistoryResponse, error) {
	i := InsolvencyHistoryResponse{}
	resp, err := a.CallAPI("/company/"+c+"/insolvency", nil, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// GetOfficers gets the json data for a company's officers from the Companies House REST API
// and returns a new OfficersResponse and an error
func (a *API) GetOfficers(c string) (*OfficerResponse, error) {
	o := OfficerResponse{}
	resp, err := a.CallAPI("/company/"+c+"/officers", nil, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &o)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

// GetFilings gets the json data for a company's filing hisotry from the Companies House REST API
// and returns a new FilingResponse and an error
func (a *API) GetFilingHistory(c string) (*FilingResponse, error) {
	f := FilingResponse{}
	resp, err := a.CallAPI("/company/"+c+"/filing-history", nil, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &f)
	if err != nil {
		return nil, err
	}

	return &f, err
}

// GetFile gets a Filing's document's content and returns the document content as a slice of bytes or returns an error
func (a *API) GetDocument(f *Filing) ([]byte, error) {
	url := f.Links.DocumentMetaData
	resp, err := a.getResponse(url, "GET", nil, ContentTypePDF)
	if err != nil {
		return nil, err
	}

	m := &DocumentMetaData{}
	json.NewDecoder(resp.Body).Decode(m)
	resp.Body.Close()

	resp, err = a.getResponse(m.Links.Document, "GET", nil, ContentTypePDF)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// DownloadDocument download a Filing's document's content and returns an error
func (a *API) DownloadDocument(f *Filing, p string) error {
	b, err := a.GetDocument(f)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) SearchCompany(q string, i int, s int) (*CompanySearchResults, error) {
	r := &CompanySearchResults{}

	var p QueryParams = QueryParams{"q": q}
	if i > 0 {
		p["items_per_page"] = i
	}
	if s > 0 {
		p["start_index"] = s
	}

	resp, err := a.CallAPI("/search/companies", p, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *API) SearchOfficer(q string, i int, s int) (*OfficerSearchResults, error) {
	r := &OfficerSearchResults{}

	var p QueryParams = QueryParams{"q": q}
	if i > 0 {
		p["items_per_page"] = i
	}
	if s > 0 {
		p["start_index"] = s
	}

	resp, err := a.CallAPI("/search/officers", p, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *API) SearchDisqualifiedOfficer(q string, i int, s int) (*DisqualifiedOfficerSearchResults, error) {
	r := &DisqualifiedOfficerSearchResults{}

	var p QueryParams = QueryParams{"q": q}
	if i > 0 {
		p["items_per_page"] = i
	}
	if s > 0 {
		p["start_index"] = s
	}

	resp, err := a.CallAPI("/search/disqualified-officers", p, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *API) Search(q string, i int, s int) (*GeneralSearchResults, error) {
	r := &GeneralSearchResults{}

	var p QueryParams = QueryParams{"q": q}
	if i > 0 {
		p["items_per_page"] = i
	}
	if s > 0 {
		p["start_index"] = s
	}

	resp, err := a.CallAPI("/search", p, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewAPI(key string) *API {
	return &API{apiKey: key, url: defaultURL}
}

// Todo: Officer Disqualifications
// Todo: UK Establshment companies
// Todo: People with significant control
