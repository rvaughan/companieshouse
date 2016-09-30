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
	Charge struct {
		ID                   string `json:"id"`
		Etag                 string `json:"etag"`
		AcquiredOn           string `json:"acquired_on"`
		AssetsCeasedReleased string `json:"assets_ceased_released"`
		ChargeCode           string `json:"charge_code"`
		ChargeNumber         int    `json:"charge_number"`
		Classification       []struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"classification"`
		CoveringInstrumentDate string `json:"covering_instrument_date"`
		Created                string `json:"created_on"`
		Delivered              string `json:"delivered_on"`
		Cases                  []struct {
			Number int `json:"case_number"`
			Links  struct {
				Case string `json:"case"`
			} `json:"links"`
			TransactionID int `json:"transaction_id"`
		} `json:"insolvency_cases"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		MoreThanFourEntitled bool `json:"more_than_four_persons_entitled"`
		Particulars          struct {
			ChargorBareTrustee      bool   `json:"chargor_acting_as_bare_trustee"`
			FixedCharge             bool   `json:"contains_fixed_charge"`
			FloatingCharge          bool   `json:"floating_charge_covers_all"`
			FloatingChargeCoversAll bool   `json:"chargor_acting_as_bare_trustee"`
			NegativePledge          bool   `json:"contains_negative_pledge"`
			Description             string `json:"description"`
			Type                    string `json:"type"`
		} `json:"particulars"`
		PersonsEntitled []struct {
			Name string `json:"name"`
		} `json:"persons_entitled"`
		Resolved            string `json:"resolved_on"`
		Satisfied           string `json:"satisfied_on"`
		ScottishAlterations struct {
			AlterationsToOrder        bool `json:"has_alterations_to_order"`
			AlterationsToProhibitions bool `json:"has_alterations_to_prohibitions"`
			RestrictingProvisions     bool `json:"has_restricting_provisions"`
		} `json:"scottish_alterations"`
		SecuredDetails struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"secured_details"`
		Status       string `json:"status"`
		Transactions []struct {
			DeliveredOn      string `json:"delivered_on"`
			FilingType       string `json:"filing_type"`
			InsolvencyNumber int    `json:"insolvency_case_number"`
			Links            struct {
				Filing         string `json:"filing"`
				InsolvencyCase `json:"insolvency_case"`
			} `json:"links"`
			TransactionID string `json:"transaction_id"`
		} `json:"transactions"`
	}

	// FilingResponse struct for API responses of Filing objects
	ChargesResponse struct {
		Etag          string   `json:"etag"`
		PartSatisfied int      `json:"part_satisfied_count"`
		Satisfied     int      `json:"satisfied_count"`
		Total         int      `json:"total_count"`
		Unfiletered   int      `json:"unfiletered_count"`
		Charges       []Charge `json:"items"`
	}
)

// GetCharges func. Takes *Company. Returns (FilingResponse, error)
func (c *Company) GetCharges() (ChargesResponse, error) {
	var res ChargesResponse
	body, err := c.API.callAPI("company/"+c.CompanyNumber+"/charges", false)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
