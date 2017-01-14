/*
Golang Companies House REST service API
Copyright (C) 2016-2017, Balkan C & T OOD

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
	"errors"
)

type (
	// Filing contains the data of a company's filing
	Filing struct {
		Annotations []struct {
			Annotation  string `json:"annotation"`
			Date        string `json:"date"`
			Description string `json:"description"`
		} `json:"annotations"`
		Associated []struct {
			Type        string `json:"type"`
			Date        string `json:"date"`
			Description string `json:"description"`
		} `json:"associated_filings"`
		Barcode     string `json:"barcode"`
		Category    string `json:"category"`
		Subcategory string `json:"subcategory"`
		Date        string `json:"date"`
		Description string `json:"description"`
		Links       Links  `json:"links"`
		Pages       int    `json:"pages"`
		PaperFiled  bool   `json:"paper_filed"`
		Resolutions []struct {
			Category    string `json:"category"`
			Description string `json:"description"`
			DocumentID  string `json:"document_id"`
			ReceiveDate string `json:"receive_date"`
			Subcategory string `json:"subcategory"`
			Type        string `json:"type"`
		} `json:"resolutions"`
		TransactionID string `json:"transaction_id"`
		Type          string `json:"type"`
	}

	// FilingResponse contains the server response of a data request to the companies house API
	FilingResponse struct {
		Etag         string   `json:"etag"`
		Kind         string   `json:"kind"`
		Start        int      `json:"start_index"`
		ItemsPerPage int      `json:"items_per_page"`
		TotalResults int      `json:"total_results"`
		Status       string   `json:"filing_history_status"`
		Filings      []Filing `json:"items"`
	}
)

// GetDownloadURL returns the download URL for a document related to a company's filing
// and returns a string with the url and an error
func (c *Company) GetDownloadURL(f *Filing) (string, error) {
	resp, err := c.API.getResponse(f.Links.Document+"/content", ContentTypePDF)
	if err != nil {
		return "", err
	}

	if len(resp.Header.Get("Location")) == 0 {
		return "", errors.New("Response's header has no Location")
	}

	return resp.Header.Get("Location"), err
}

// GetFilings gets the json data for a company's filing hisotry from the Companies House REST API
// and returns a new FilingResponse and an error
func (c *Company) GetFilings() (*FilingResponse, error) {
	res := &FilingResponse{}
	body, err := c.API.callAPI("/company/"+c.CompanyNumber+"/filing-history", false, ContentTypeJSON)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}