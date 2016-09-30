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

type (
	// Filing struct contains the data of a company's filing
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
		Links       struct {
			Document string `json:"document_metadata"`
			Self     string `json:"self"`
		} `json:"links"`
		Pages       int  `json:"pages"`
		PaperFiled  bool `json:"paper_filed"`
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

	// FilingResponse struct for API responses of Filing objects
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

// DownloadURL func. Takes *Filing. Returns (string, error)
func (c *Company) DownloadURL(f *Filing) (string, error) {
	var fu string
	var body map[string]string

	fmt.Printf("Document url: %s\n", f.Links.Document)
	resp, err := c.API.makeGetRequest(f.Links.Document, true, "application/pdf")
	if err != nil {
		return fu, err
	}
	fmt.Printf("Header: %+v", resp.Header)
	json.NewDecoder(resp.Body).Decode(&body)
	fmt.Printf("%+v", body)
	return fu, err
}

// GetFilings func. Takes *Company. Returns (FilingResponse, error)
func (c *Company) GetFilings() (FilingResponse, error) {
	var res FilingResponse
	body, err := c.API.callAPI("company/"+c.CompanyNumber+"/filing-history", false)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
