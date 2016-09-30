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
	// Address struct contains the structure of addresses used by Companies House
	Address struct {
		Premises string `json:"premises"`
		Address1 string `json:"address_line_1"`
		Address2 string `json:"address_line_2"`
		Locality string `json:"locality"`
		Region   string `json:"region"`
		Postcode string `json:"postal_code"`
		Country  string `json:"country"`
		CareOf   string `json:"care_of"`
		PoBox    string `json:"po_box"`
	}

	// PreviousName struct contains data of previous names and the time of use
	PreviousName struct {
		Name  string `json:"name"`
		From  string `json:"effective_from"`
		Until string `json:"ceased_on"`
	}

	// RefDate struct consists of Day and Month
	RefDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
	}

	// AnnualAccounts struct contains the last and next filing info for the Annual Accounts
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

	// AnnualReturn struct contains the last and next filing dates for the Annual Return
	AnnualReturn struct {
		LastMadeUpTo string `json:"last_made_up_to"`
		NextDue      string `json:"next_due"`
		NextMadeUpTo string `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// Links struct contains links to additional data and the original object (Self)
	Links struct {
		Psc              string `json:"persons_with_significant_control"`
		PscStatements    string `json:"persons_with_significant_control_statements`
		Registers        string `json:"registers"`
		UkEstablishments string `json:"uk_establishments"`
		Self             string `json:"self"`
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
		CompanyNumber       string `json:"companyNumber"`
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

// GetCompany func. Takes companyNumber string. Returns (Company, error)
func (a *ChAPI) GetCompany(companyNumber string) (Company, error) {
	var c Company

	body, err := a.callAPI("company/"+companyNumber, false)
	if err != nil {
		return Company{}, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return Company{}, err
	}

	c.API = a

	return c, err
}
