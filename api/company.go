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
	// PreviousName struct contains data of a company's previous names and the time of use
	PreviousName struct {
		Name          string `json:"name"`
		EffectiveFrom ChDate `json:"effective_from"`
		CeasedOn      ChDate `json:"ceased_on"`
	}

	// RefDate struct consists of Day and Month
	RefDate struct {
		Day   string `json:"day"`
		Month string `json:"month"`
	}

	// Accounts struct contains a company's last and next filing info for the Annual Accounts
	Accounts struct {
		AccountingReferenceDate RefDate `json:"accounting_reference_date"`
		LastAccounts            struct {
			MadeUpTo      ChDate `json:"made_up_to"`
			Type          string `json:"type"`
			PeriodEndOn   ChDate `json:"period_end_on"`
			PeriodStartOn ChDate `json:"period_start_on"`
		} `json:"last_accounts"`
		NextAccounts struct {
			DueOn         ChDate `json:"due_on"`
			Overdue       bool   `json:"overdue"`
			PeriodEndOn   ChDate `json:"period_end_on"`
			PeriodStartOn ChDate `json:"period_start_on"`
		} `json:"next_accounts"`
		NextDue      ChDate `json:"next_due"`
		NextMadeUpTo ChDate `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// AnnualReturn struct contains a company's the last and next filing dates for the Annual Return
	AnnualReturn struct {
		LastMadeUpTo ChDate `json:"last_made_up_to"`
		NextDue      ChDate `json:"next_due"`
		NextMadeUpTo ChDate `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// Branch struct contains data of a Branch
	Branch struct {
		BusinessActivity    string `json:"business_activity"`
		ParentCompanyName   string `json:"parent_company_name"`
		ParentCompanyNumber string `json:"parent_company_number"`
	}

	// ForeignCompany struct contains data of Foreign Companies
	ForeignCompany struct {
		AccountingRequirement struct {
			AccountType string `json:"foreign_account_type"`
			Terms       string `json:"terms_of_account_publication"`
		}
		Accounts struct {
			From RefDate `json:"account_period_from"`
			To   RefDate `json:"account_period_to"`
			Term struct {
				Months string `json:"months"`
			} `json:"must_file_within"`
		} `json:"accounts`
		LegalForm           string `json:"legal_form"`
		BusinessActivity    string `json:"business_activity"`
		GovernedBy          string `json:"governed_by"`
		FinanceInstitution  bool   `json:"is_a_credit_finance_institution"`
		RegistrationNumber  string `json:"registration_number"`
		OriginatingRegistry struct {
			Country string `json:"country"`
			Name    string `json:"name"`
		} `json:"originating_registry"`
	}

	// Company struct contains basic company data
	Company struct {
		api                 *API   `json:"-"`
		Etag                string `json:"etag"`
		CompanyNumber       string `json:"company_number"`
		CompanyName         string `json:"company_name"`
		CanFile             bool   `json:"can_file"`
		Type                string `json:"type"`
		CompanyStatus       string `json:"company_status"`
		CompanyStatusDetail string `json:"company_status_detail"`
		DateOfCessation     ChDate `json:"date_of_cessation"`
		DateOfCreation      ChDate `json:"date_of_creation"`

		HasCharges                           bool   `json:"has_charges"`
		HasInsolvencyHistory                 bool   `json:"has_insolvency_history"`
		IsCommunityInterestCompany           bool   `json:"is_community_interest_company"`
		Jurisdiction                         string `json:"jurisdiction"`
		LastFullMemberListDate               ChDate `json:"last_full_members_list_date"`
		Liquidated                           bool   `json:"has_been_liquidated"`
		UndeliverableRegisteredOfficeAddress bool   `json:"undeliverable_registered_office_address"`
		RegisteredOfficeIsInDispute          bool   `json:"registered_office_is_in_dispute"`

		RegisteredOfficeAddress Address        `json:"registered_office_address"`
		AnnualReturn            AnnualReturn   `json:"annual_return"`
		ConfirmationStatement   AnnualReturn   `json:"confirmation_statement"`
		Accounts                Accounts       `json:"accounts"`
		SICCodes                []string       `json:"sic_codes"`
		PreviousCompanyNames    []PreviousName `json:"previous_company_names"`
		Links                   Links          `json:"links"`
		BranchCompanyDetails    Branch         `json:"branch_company_details"`
		ForeignCompanyDetails   ForeignCompany `json:"foreign_company_details"`
	}
)

func (c Company) HasTasks() bool {
	return c.AnnualReturn != (AnnualReturn{}) || c.ConfirmationStatement != (AnnualReturn{}) || c.Accounts != (Accounts{})
}

// GetCompany gets the json data for a company from the Companies House REST API
// and returns a new Company and an error
func (a *API) GetCompany(companyNumber string) (*Company, error) {
	c := &Company{}

	resp, err := a.CallAPI("/company/"+companyNumber, nil, false, ContentTypeJSON)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(resp, &c)
	if err != nil {
		return nil, err
	}

	c.api = a

	return c, err
}

// Todo: SIC code description
