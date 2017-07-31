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
	"github.com/BalkanTech/companieshouse/api/enum"
	"github.com/BalkanTech/kit/stringutils"
)

type InsolvencyCaseDateType string

func (f InsolvencyCaseDateType) String() string {
	return enum.Constants.Get("insolvency_case_date_type", string(f))
}

type InsolvencyCaseType string
func (f InsolvencyCaseType) String() string {
	return enum.Constants.Get("insolvency_case_type", string(f))
}

type (
	// Insolvency contains the data of an insolvency case
	Insolvency struct {
		Dates []struct {
			Date string `json:"date"`
			Type InsolvencyCaseDateType `json:"type"`
		} `json:"dates"`
		Notes         []string `json:"notes"`
		Number        string   `json:"number"`
		Practitioners []struct {
			Address   Address `json:"address"`
			Appointed string  `json:"appointed_on"`
			Ceased    string  `json:"ceased_to_act_on"`
			Name      string  `json:"name"`
			Role      stringutils.TitledString  `json:"role"`
		} `json:"practitioners"`
		Type InsolvencyCaseType `json:"type"`
		Status stringutils.TitledString `json:"status"`
	}

	// InsolvenciesResponse contains the server response of a data request to the companies house API
	InsolvencyHistoryResponse struct {
		Etag   string       `json:"etag"`
		Status []string     `json:"status"`
		Cases  []Insolvency `json:"cases"`
	}
)