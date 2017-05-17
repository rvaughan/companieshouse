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
	"github.com/BalkanTech/companieshouse/api/yaml"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
	"html/template"
)

func ReplaceBetween(s *string, t, st, et string) {
	var p int = len(t)
	var useSt bool = true
	str := *s

	for {
		i := strings.Index(str, t)
		if i == -1 {
			*s = str
			break
		}

		if useSt {
			str = str[:i] + st + str[i + p:]
		} else {
			str = str[:i] + et + str[i + p:]
		}
		useSt = !useSt
	}
}

type FilingDescription string

func (fd FilingDescription) String() string {
	d, ok := yaml.FilingHistoryDescriptions[string(fd)]
	if !ok {
		return ""
	}
	return d
}

func (fd FilingDescription) AsHTML() template.HTML {
	d, ok := yaml.FilingHistoryDescriptions[string(fd)]
	if !ok {
		return ""
	}
	ReplaceBetween(&d, "**", "<strong>", "</strong>")
	return template.HTML(d)
}

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
			Date        ChDate `json:"date"`
			Description string `json:"description"`
		} `json:"annotations"`
		AssociatedFilings []struct {
			Type        string `json:"type"`
			Date        ChDate `json:"date"`
			Description string `json:"description"`
			DescriptionValues struct {
				Description FilingDescription `json:"description"`
			} `json:"description_values"`
		} `json:"associated_filings"`
		Barcode     string            `json:"barcode"`
		Category    string            `json:"category"`
		Subcategory string            `json:"subcategory"`
		Date        ChDate            `json:"date"`
		Description FilingDescription `json:"description"`
		DescriptionValues struct {
			Description FilingDescription `json:"description"`
		} `json:"description_values"`
		Links       Links             `json:"links"`
		Pages       int               `json:"pages"`
		PaperFiled  bool              `json:"paper_filed"`
		Resolutions []struct {
			Category    string `json:"category"`
			Description FilingDescription `json:"description"`
			DescriptionValues struct {
				Description FilingDescription `json:"description"`
			} `json:"description_values"`
			DocumentID  string `json:"document_id"`
			ReceiveDate ChDate `json:"receive_date"`
			Subcategory string `json:"subcategory"`
			Type        string `json:"type"`
		} `json:"resolutions"`
		TransactionID string `json:"transaction_id"`
		Type          string `json:"type"`
	}

	// FilingResponse contains the server response of a data request to the companies house API
	FilingResponse struct {
		Etag                string    `json:"etag"`
		FilingHistoryStatus string    `json:"filing_history_status"`
		Kind                string    `json:"kind"`
		StartIndex          int       `json:"start_index"`
		ItemsPerPage        int       `json:"items_per_page"`
		TotalCount          int       `json:"total_count"`
		Items               []*Filing `json:"items"`
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

func (a *API) GetFilings(c string) (<-chan *FilingResponse, <-chan error) {
	r := make(chan *FilingResponse, 1)
	e := make(chan error, 1)

	go func() {
		defer close(r)
		defer close(e)

		f := &FilingResponse{}

		resp, err := a.CallAPI("/company/"+c+"/filing-history", nil, false, ContentTypeJSON)
		if err != nil {
			r <- nil
			e <- err
			return
		}

		err = json.Unmarshal(resp, f)
		if err != nil {
			r <- nil
			e <- err
			return
		}

		r <- f
		e <- nil
	}()

	return r, e
}

// GetFilings gets the json data for a company's filing hisotry from the Companies House REST API
// and returns a new FilingResponse and an error
func (c *Company) getFilings() (*FilingResponse, error) {
	filings := &FilingResponse{}
	resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/filing-history", nil, false, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &filings)
	if err != nil {
		return nil, err
	}

	return filings, err
}

// GetFiling
func (c *Company) GetFiling(tid string) (*Filing, error) {
	for _, f := range c.Filings.Items {
		if f.TransactionID == tid {
			return f, nil
		}
	}
	return nil, errors.New("Not found")
}
