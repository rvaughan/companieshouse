package companieshouse

type (
	SearchAddress struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		CareOf       string `json:"care_of"`
		Country      string `json:"country"`
		Locality     string `json:"locality"`
		POBox        string `json:"po_box"`
		PostalCode   string `json:"postal_code"`
		Region       string `json:"region"`
	}

	SearchResult struct {
		Address               SearchAddress `json:"address"`
		AddressSnippet        string        `json:"address_snippet"`
		Description           string        `json:"description"`
		DescriptionIdentifier []string      `json:"description_identifier"`
		Kind                  string        `json:"kind"`
		Links                 struct {
			Self string `json:"self"`
		} `json:"links"`
		Matches struct {
			AddressSnippet []int `json:"address_snippet"`
			Snippet        []int `json:"snippet"`
			Title          []int `json:"title"`
		}
		Snippet string `json:"snippet"`
		Title   string `json:"title"`
	}

	CompanySearchResult struct {
		SearchResult
		CompanyNumber   string `json:"company_number"`
		CompanyStatus   CompanyStatus `json:"company_status"`
		CompanyType     CompanyType `json:"company_type"`
		DateOfCessation string `json:"date_of_cessation"`
		DataOfCreation  string `json:"date_of_creation"`
	}

	OfficerSearchResult struct {
		SearchResult
		AppointmentCount int `json:"appointment_count"`
		DateOfBirth      struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"date_of_birth"`
	}

	DisqualifiedOfficerSearchResult struct {
		SearchResult
		DateOfBirth           DateOfBirth `json:"date_of_birth"`
		Description           string      `json:"description"`
		DescriptionIdentifier []string    `json:"description_identifier"`
	}

	SearchResults struct {
		Etag         string `json:"etag"`
		ItemsPerPage int    `json:"ItemsPerPage"`
		Kind         string `json:"Kind"`
		StartIndex   int    `json:"start_index"`
		TotalResults int    `json:"total_results"`
	}

	GeneralSearchResults struct {
		SearchResults
		Items []SearchResult `json:"items"`
	}

	CompanySearchResults struct {
		SearchResults
		Items []CompanySearchResult `json:"items"`
	}

	OfficerSearchResults struct {
		SearchResults
		Items []OfficerSearchResult `json:"items"`
	}

	DisqualifiedOfficerSearchResults struct {
		SearchResults
		Items []DisqualifiedOfficerSearchResult `json:"items"`
	}
)