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
)

type (
	// Insolvency contains the data of an insolvency case
	Insolvency struct {
		Dates []struct {
			Date string `json:"date"`
			Type string `json:"type"`
		} `json:"dates"`
		Notes         []string `json:"notes"`
		Number        string   `json:"number"`
		Practitioners []struct {
			Address   Address `json:"address"`
			Appointed string  `json:"appointed_on"`
			Ceased    string  `json:"ceased_to_act_on"`
			Name      string  `json:"name"`
			Role      string  `json:"role"`
		} `json:"practitioners"`
		Type string `json:"type"`
	}

	// InsolvenciesResponse contains the server response of a data request to the companies house API
	InsolvencyHistoryResponse struct {
		Etag   string       `json:"etag"`
		Status []string     `json:"status"`
		Cases  []*Insolvency `json:"cases"`
	}
)

// GetInsolvencyDetails gets the json data for a company's insolvency details from the Companies House REST API
// and returns a new InsolvenciesResponse and an error
func (c *Company) GetInsolvencyDetails() (*InsolvencyHistoryResponse, error) {
	insolvencies := &InsolvencyHistoryResponse{}
	resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/insolvency", nil, false, ContentTypeJSON)
	if err != nil {
		return insolvencies, err
	}

	err = json.Unmarshal(resp, insolvencies)
	if err != nil {
		return insolvencies, err
	}
	return insolvencies, err
}

func (c *Company) GetInsolvencyHistory() (<-chan *InsolvencyHistoryResponse, <-chan error) {
	r := make(chan *InsolvencyHistoryResponse, 1)
	e := make(chan error, 1)

	go func() {
		ih := &InsolvencyHistoryResponse{}
		resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/insolvency", nil, false, ContentTypeJSON)
		if err != nil {
			e <- err
		}

		err = json.Unmarshal(resp, &ih)
		if err != nil {
			e <- err
		}
		r <- ih
		e <- nil
		close(r)
		close(e)
	}()

	return r, e
}
