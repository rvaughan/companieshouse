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
	// Charge contains the data of a company's charges
	Charge struct {
		Etag                 string `json:"etag"`
		AcquiredOn           ChDate `json:"acquired_on"`
		AssetsCeasedReleased string `json:"assets_ceased_released"`
		ChargeCode           string `json:"charge_code"`
		ChargeNumber         int    `json:"charge_number"`
		Classification       struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"classification"`
		CoveringInstrumentDate ChDate `json:"covering_instrument_date"`
		CreatedOn                ChDate `json:"created_on"`
		DeliveredOn              ChDate `json:"delivered_on"`
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
		ResolvedOn            ChDate `json:"resolved_on"`
		SatisfiedOn           ChDate `json:"satisfied_on"`
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
			DeliveredOn      ChDate `json:"delivered_on"`
			FilingType       string `json:"filing_type"`
			InsolvencyNumber int    `json:"insolvency_case_number"`
			Links            struct {
				Filing     string `json:"filing"`
				Insolvency `json:"insolvency_case"`
			} `json:"links"`
			TransactionID string `json:"transaction_id"`
		} `json:"transactions"`
	}

	// ChargesResponse contains the server response of a data request to the companies house API
	ChargesResponse struct {
		Etag          string   `json:"etag"`
		PartSatisfiedCount int      `json:"part_satisfied_count"`
		SatisfiedCount    int      `json:"satisfied_count"`
		TotalCount         int      `json:"total_count"`
		UnfileteredCount   int      `json:"unfiletered_count"`
		Items       []*Charge `json:"items"`
	}
)

func (c *Company) getCharges() (*ChargesResponse, error) {
	charges := &ChargesResponse{}

	resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/charges", nil, false, ContentTypeJSON)
	if err != nil {
		return charges, err
	}

	err = json.Unmarshal(resp, charges)
	if err != nil {
		return charges, err
	}

	return charges, err
}

// GetCharges gets the json data for a company's charges from the Companies House REST API
// and returns a new ChargesResponse and an error
func (a *API) GetCharges(c string) (<-chan *ChargesResponse, <-chan error) {
	r := make(chan *ChargesResponse, 1)
	e := make(chan error, 1)

	go func() {
		defer close(r)
		defer close(e)

		charges := &ChargesResponse{}

		resp, err := a.CallAPI("/company/" + c + "/charges", nil, false, ContentTypeJSON)
		if err != nil {
			r <- nil
			e <- err
			return
		}

		err = json.Unmarshal(resp, charges)
		if err != nil {
			r <- nil
			e <- err
			return
		}
		r <- charges
		e <- nil
	}()

	return r, e
}
