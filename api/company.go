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
)

type (
	// PreviousName struct contains data of a company's previous names and the time of use
	PreviousName struct {
		Name  string `json:"name"`
		From  string `json:"effective_from"`
		Until string `json:"ceased_on"`
	}

	// RefDate struct consists of Day and Month
	RefDate struct {
		Day   string `json:"day"`
		Month string `json:"month"`
	}

	// AnnualAccounts struct contains a company's last and next filing info for the Annual Accounts
	AnnualAccounts struct {
		RefDate      RefDate `json:"accounting_reference_date"`
		LastAccounts struct {
			MadeUpTo string `json:"made_up_to"`
			Type     string `json:"type"`
		} `json:"last_accounts"`
		NextDue      string `json:"next_due"`
		NextMadeUpTo string `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// AnnualReturn struct contains a company's the last and next filing dates for the Annual Return
	AnnualReturn struct {
		LastMadeUpTo string `json:"last_made_up_to"`
		NextDue      string `json:"next_due"`
		NextMadeUpTo string `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// Branch struct contains data of a Branch
	Branch struct {
		Activity            string `json:"business_activity"`
		ParentCompanyCame   string `json:"Parent_company_name"`
		ParentCompanyNumber string `json:Parent_companyNumber`
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
		API                 *ChAPI
		Etag                string `json:"etag"`
		CompanyNumber       string `json:"company_number"`
		CompanyName         string `json:"company_name"`
		CanFile             bool   `json:"can_file"`
		CompanyType         string `json:"type"`
		CompanyStatus       string `json:"company_status"`
		CompanyStatusDetail string `json:"company_status_detail"`
		DateOfCessation     string `json:"date_of_cessation"`
		DateOfCreation      string `json:"date_of_creation"`

		Charges           bool   `json:"has_charges"`
		InsolvencyHistory bool   `json:"has_insolvency_history"`
		Cic               bool   `json:"is_community_interest_company"`
		Jurisdiction      string `json:"jurisdiction"`
		LastFmList        string `json:"last_full_members_list_date"`
		Liquidated        bool   `json:"has_been_liquidated"`
		RoUndeliverable   bool   `json:"undeliverable_registered_office_address"`
		RoDispute         bool   `json:"registered_office_is_in_dispute"`

		RegisteredOffice      Address        `json:"registered_office_address"`
		Return                AnnualReturn   `json:"annual_return"`
		ConfirmationStatement AnnualReturn   `json:"confirmation_statement"`
		Accounts              AnnualAccounts `json:"accounts"`
		SicCodes              []string       `json:"sic_codes"`
		PreviousNames         []PreviousName `json:"previous_company_names"`
		Links                 Links          `json:"links"`
		Branch                Branch         `json:"branch_company_details"`
		ForeignCompany        ForeignCompany `json:"foreign_company_details"`
	}
)

// GetCompany gets the json data for a company from the Companies House REST API
// and returns a new Company and an error
func (a *ChAPI) GetCompany(companyNumber string) (*Company, error) {
	c := &Company{}

	body, err := a.callAPI("/company/"+companyNumber, false, ContentTypeJSON)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}

	c.API = a

	return c, err
}
