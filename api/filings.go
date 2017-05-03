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
)

type (
	ContentType struct {
		ContentLength int    `json:"content_length"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
	}

	DocumentMetaData struct {
		CompanyNumber       string `json:"company_number"`
		Barcode             string `json:"barcode"`
		SignificantDate     string `json:"significan_date"`
		SignificantDateType string `json:"significan_date_type"`
		Category            string `json:"category"`
		CreatedAt           string `json:"created_at"`
		Links               struct {
			Self     string `json:"self"`
			Document string `json:"document"`
		} `json:"links"`
		Pages     int `json:"pages"`
		Resources []struct {
			ApplicationPDF      ContentType `json:"application/pdf"`
			ApplicationXhtmlXml ContentType `json:"application/xhtml+xml"`
		} `json:"resources"`
		UpdatedAt string `json:"updated_at"`
	}

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

// GetFile gets a Filing's document's content and returns the document content as a slice of bytes or returns an error
func (c *Company) GetDocument(f *Filing) ([]byte, error) {
	url := f.Links.DocumentMetaData
	resp, err := c.api.getResponse(url, "GET", nil, ContentTypePDF)
	if err != nil {
		return nil, err
	}

	m := &DocumentMetaData{}
	json.NewDecoder(resp.Body).Decode(m)
	resp.Body.Close()

	fmt.Printf("%s %+v\n\n", resp.Status, m)
	//fmt.Println("ID;", m.ID)

	resp, err = c.api.getResponse(m.Links.Document, "GET", nil, ContentTypePDF)
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

// OpenFile download a Filing's document's content and returns an error
func (c *Company) DownloadDocument(f *Filing, p string) error {
	b, err := c.GetDocument(f)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// GetFilings gets the json data for a company's filing hisotry from the Companies House REST API
// and returns a new FilingResponse and an error
func (c *Company) GetFilings() (*FilingResponse, error) {
	filings := &FilingResponse{}
	resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/filing-history", nil, false, ContentTypeJSON)
	if err != nil {
		return filings, err
	}

	err = json.Unmarshal(resp, &filings)
	if err != nil {
		return filings, err
	}

	return filings, err
}
