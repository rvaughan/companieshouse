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
	"fmt"
	"errors"
	"github.com/BalkanTech/companieshouse/api/enum"
)

type CompanyType string

func (f CompanyType) String() string {
	return enum.Constants.Get("company_type", string(f))
}

type AccountType string

func (f AccountType) String() string {
	return enum.Constants.Get("account_type", string(f))
}

type ForeignAccountType string

func (f ForeignAccountType) String() string {
	return enum.Constants.Get("foreign_account_type", string(f))
}

type CompanyStatus string

func (f CompanyStatus) String() string {
	return enum.Constants.Get("company_status", string(f))
}

type CompanyStatusDetail string

func (f CompanyStatusDetail) String() string {
	return enum.Constants.Get("company_status_detail", string(f))
}

type TermsOfAccountPublication string

func (f TermsOfAccountPublication) String() string {
	return enum.Constants.Get("terms_of_account_publication", string(f))
}

type Jurisdiction string

func (f Jurisdiction) String() string {
	return enum.Constants.Get("jurisdiction", string(f))
}

type PartialDataAvailable string

func (f PartialDataAvailable) String() string {
	return enum.Constants.Get("partial_data_available", string(f))
}

type SICCode string

func (f SICCode) String() string {
	desc := enum.Constants.Get("sic_descriptions", string(f))
	if desc == "" {
		fmt.Sprintf("%s - Unknown", string(f), desc)
	}
	return fmt.Sprintf("%s - %s", string(f), desc)
}

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
			MadeUpTo      ChDate      `json:"made_up_to"`
			Type          AccountType `json:"type"`
			PeriodEndOn   ChDate      `json:"period_end_on"`
			PeriodStartOn ChDate      `json:"period_start_on"`
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
	ForeignCompanyDetails struct {
		AccountingRequirement struct {
			ForeignAccountType        ForeignAccountType        `json:"foreign_account_type"`
			TermsOfAccountPublication TermsOfAccountPublication `json:"terms_of_account_publication"`
		} `json:"accounting_requirement"`
		Accounts struct {
			From RefDate `json:"account_period_from"`
			To   RefDate `json:"account_period_to"`
			Term struct {
				Months string `json:"months"`
			} `json:"must_file_within"`
		} `json:"accounts`
		LegalForm                     string `json:"legal_form"`
		CompanyType                   string `json:"company_type"`
		BusinessActivity              string `json:"business_activity"`
		GovernedBy                    string `json:"governed_by"`
		IsACreditFinancialInstitution bool   `json:"is_a_credit_financial_institution"`
		RegistrationNumber            string `json:"registration_number"`
		OriginatingRegistry           struct {
			Country string `json:"country"`
			Name    string `json:"name"`
		} `json:"originating_registry"`
	}

	// Company struct contains basic company data
	Company struct {
		api                        *API                  `json:"-"`
		Accounts                   Accounts              `json:"accounts"`
		AnnualReturn               AnnualReturn          `json:"annual_return"`
		BranchCompanyDetails       Branch                `json:"branch_company_details"`
		CanFile                    bool                  `json:"can_file"`
		CompanyName                string                `json:"company_name"`
		CompanyNumber              string                `json:"company_number"`
		CompanyStatus              CompanyStatus         `json:"company_status"`
		CompanyStatusDetail        CompanyStatusDetail   `json:"company_status_detail"`
		ConfirmationStatement      AnnualReturn          `json:"confirmation_statement"`
		DateOfCessation            ChDate                `json:"date_of_cessation"`
		DateOfCreation             ChDate                `json:"date_of_creation"`
		Etag                       string                `json:"etag"`
		ForeignCompanyDetails      ForeignCompanyDetails `json:"foreign_company_details"`
		HasBeenLiquidated          bool                  `json:"has_been_liquidated"`
		HasCharges                 bool                  `json:"has_charges"`
		HasInsolvencyHistory       bool                  `json:"has_insolvency_history"`
		IsCommunityInterestCompany bool                  `json:"is_community_interest_company"`
		Jurisdiction               Jurisdiction          `json:"jurisdiction"`
		LastFullMembersListDate    ChDate                `json:"last_full_members_list_date"`
		Links                      struct {
			Charges                                 string `json:"charges"`
			FilingHistory                           string `json:"filing_history"`
			Insolvency                              string `json:"insolvency"`
			Officers                                string `json:"officers"`
			PersonsWithSignificantControl           string `json:"persons_with_significant_control"`
			PersonsWithSignificantControlStatements string `json:"persons_with_significant_control_statements"`
			Registers                               string `json:"registers"`
			Self                                    string `json:"self"`
		} `json:"links"`
		PartialDataAvailable                 PartialDataAvailable `json:"partial_data_available"`
		PreviousCompanyNames                 []PreviousName       `json:"previous_company_names"`
		RegisteredOfficeAddress              Address              `json:"registered_office_address"`
		RegisteredOfficeIsInDispute          bool                 `json:"registered_office_is_in_dispute"`
		SICCodes                             []SICCode            `json:"sic_codes"`
		Type                                 CompanyType          `json:"type"`
		UndeliverableRegisteredOfficeAddress bool                 `json:"undeliverable_registered_office_address"`

		Officers          OfficerResponse           `json:"-"`
		FilingHistory     FilingResponse            `json:"-"`
		Charges           ChargesResponse           `json:"-"`
		InsolvencyHistory InsolvencyHistoryResponse `json:"-"`
	}
)

func (c Company) HasTasks() bool {
	return c.AnnualReturn != (AnnualReturn{}) || c.ConfirmationStatement != (AnnualReturn{}) || c.Accounts != (Accounts{})
}

// GetFiling
func (c Company) GetFiling(tid string) (*Filing, error) {
	for _, f := range c.FilingHistory.Items {
		if f.TransactionID == tid {
			return &f, nil
		}
	}
	return nil, errors.New("Not found")
}