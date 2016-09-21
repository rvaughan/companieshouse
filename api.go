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
	// ChAPI is the body for the Companies House API and contains only the apikey
	ChAPI struct {
		apikey string
	}

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

	// DateOfBirth struct consists of Day(int), Month (int) and Year (int)
	DateOfBirth struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}

	// Identification struct
	Identification struct {
		IDType             string `json:"identification_type"`
		Authority          string `json:"legal_authority"`
		LegalForm          string `json:"legal_form"`
		PlaceRegistered    string `json:"place_registered"`
		RegistrationNumber string `json:"registration_number"`
	}

	// Officer struct contains Officer data
	Officer struct {
		Address            Address     `json:"address"`
		AppointedOn        string      `json:"appointed_on"`
		CountryOfResidence string      `json:"country_of_residence"`
		Dob                DateOfBirth `json:"date_of_birth"`
		FomerNames         []struct {
			Forenames string `json:"forenames"`
			Surname   string `json:"surname"`
		} `json:"former_names"`
		Identification Identification `json:"identification"`
		Links          struct {
			Officer struct {
				Appointments string `json:"appointments"`
			} `json:"officer"`
		} `json:"links"`
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
		Occupation  string `json:"occupation"`
		Role        string `json:"officer_role"`
		ResignedOn  string `json:"resigned_on"`
	}

	// OfficerResponse struct for API responses of Officer objects
	OfficerResponse struct {
		Etag                   string    `json:"etag"`
		Kind                   string    `json:"kind"`
		Start                  int       `json:"start_index"`
		ItemsPerPage           int       `json:"items_per_page"`
		TotalResults           int       `json:"total_results"`
		ActiveAppointments     int       `json:"active_count"`
		ResignatedAppointments int       `json:"resigned_count"`
		Officers               []Officer `json:"items"`
		Links                  struct {
			self string `json:"self"`
		} `json:"Links"`
	}

	// Filing struct contains the data of a company's filing
	Filing struct {
		Annotations []struct {
			Annotation  string `json:"annotation"`
			Date        string `json:"date"`
			Description string `json:"description"`
		} `json:"annotations"`
		Associated []struct {
			Type        string `json:"type"`
			Date        string `json:"date"`
			Description string `json:"description"`
		} `json:"associated_filings"`
		Barcode     string `json:"barcode"`
		Category    string `json:"category"`
		Subcategory string `json:"subcategory"`
		Date        string `json:"date"`
		Description string `json:"description"`
		Links       struct {
			Document string `json:"document_metadata"`
			Self     string `json:"self"`
		} `json:"links"`
		Pages       int  `json:"pages"`
		PaperFiled  bool `json:"paper_filed"`
		Resolutions []struct {
			Category    string `json:"category"`
			Description string `json:"description"`
			DocumentID  string `json:"document_id"`
			ReceiveDate string `json:"receive_date"`
			Subcategory string `json:"subcategory"`
			Type        string `json:"type"`
		} `json:"resolutions"`
		TransactionID string `json:"transaction_id"`
		Type          string `json:"type"`
	}

	// FilingResponse struct for API responses of Filing objects
	FilingResponse struct {
		Etag         string   `json:"etag"`
		Kind         string   `json:"kind"`
		Start        int      `json:"start_index"`
		ItemsPerPage int      `json:"items_per_page"`
		TotalResults int      `json:"total_results"`
		Status       string   `json:"filing_history_status"`
		Filings      []Filing `json:"items"`
	}
)

// SetAPIKey func. Takes apikey string.
func (a *ChAPI) SetAPIKey(apikey string) {
	a.apikey = apikey
}

// DownloadURL func. Takes *Filing. Returns (string, error)
func (a *ChAPI) DownloadURL(f *Filing) (string, error) {
	var fu string
	var body map[string]string

	fmt.Printf("Document url: %s\n", f.Links.Document)
	resp, err := a.makeGetRequest(f.Links.Document, true, "application/pdf")
	if err != nil {
		return fu, err
	}
	fmt.Printf("Header: %+v", resp.Header)
	json.NewDecoder(resp.Body).Decode(&body)
	fmt.Printf("%+v", body)
	return fu, err
}

// GetFilings func. Takes *Company. Returns (FilingResponse, error)
func (a *ChAPI) GetFilings(c *Company) (FilingResponse, error) {
	var res FilingResponse
	body, err := a.callAPI("company/"+c.CompanyNumber+"/filing-history", false)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// GetOfficers func. Takes *Company. Returns (OfficerResponse, error)
func (a *ChAPI) GetOfficers(c *Company) (OfficerResponse, error) {
	var res OfficerResponse
	body, err := a.callAPI("company/"+c.CompanyNumber+"/officers", false)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

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

	return c, err
}

func (a *ChAPI) getRequest(path string, fullURL bool) (*http.Request, error) {
	var err error
	var url string

	apikey := a.apikey
	apiurl := "api.companieshouse.gov.uk/"

	if fullURL {
		fmt.Println("Full url")
		url = path
	} else {
		fmt.Println("No full url")
		url = "https://" + apiurl + path
	}

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	req.SetBasicAuth(apikey, "")

	return req, err
}

func (a *ChAPI) executeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

func (a *ChAPI) makeGetRequest(path string, fullURL bool, contentType string) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = a.getRequest(path, fullURL)
	if err != nil {
		return resp, err
	}

	if contentType != "" {
		req.Header.Set("Acccept", contentType)
	}
	resp, err = executeRequest(req)

	return resp, err
}

func (a *ChAPI) callAPI(path string, fullURL bool) ([]byte, error) {
	resp, err := a.makeGetRequest(path, fullURL, "")
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
