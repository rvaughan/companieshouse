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
	Insolvency struct {
		Dates []struct {
			Date string `json:"date"`
			Type string `json:"type"`
		} `json:"dates"`
		Notes         []string `json:"notes"`
		Number        int      `json:"number"`
		Practitioners []struct {
			Address   Address `json:"address"`
			Appointed string  `json:"appointed_on"`
			Ceased    string  `json:"ceased_to_act_on"`
			Name      string  `json:"name"`
			Role      string  `json:"role"`
		} `json:"practitioners"`
		Type string `json:"type"`
	}

	// FilingResponse struct for API responses of Filing objects
	InsolvenciesResponse struct {
		Etag   string       `json:"etag"`
		Status []string     `json:"status"`
		Cases  []Insolvency `json:"cases"`
	}
)

// GetInsolvencyDetails func. Takes *Company. Returns (FilingResponse, error)
func (c *Company) GetInsolvencyDetails() (InsolvenciesResponse, error) {
	var res InsolvenciesResponse
	body, err := c.API.callAPI("company/"+c.CompanyNumber+"/insolvency", false)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
