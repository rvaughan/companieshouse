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
	"github.com/rvaughan/companieshouse/api/enum"
)

type IdentificationType string

func (f IdentificationType) String() string {
	return enum.Constants.Get("identification_type", string(f))
}

type OfficerRole string

func (f OfficerRole) String() string {
	return enum.Constants.Get("officer_role", string(f))
}

type (
	// DateOfBirth struct consists of Day(int), Month (int) and Year (int)
	OfficerDateOfBirth struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}

	// Identification struct
	Identification struct {
		IdentificationType IdentificationType `json:"identification_type"`
		Authority          string             `json:"legal_authority"`
		LegalForm          string             `json:"legal_form"`
		PlaceRegistered    string             `json:"place_registered"`
		RegistrationNumber string             `json:"registration_number"`
	}

	// Officer struct contains the data of a company's officers
	Officer struct {
		Address            Address            `json:"address"`
		AppointedOn        ChDate             `json:"appointed_on"`
		CountryOfResidence string             `json:"country_of_residence"`
		DateOfBirth        OfficerDateOfBirth `json:"date_of_birth"`
		FormerNames        []struct {
			Forenames string `json:"forenames"`
			Surname   string `json:"surname"`
		} `json:"former_names"`
		Identification Identification `json:"identification"`
		Links          struct {
			Officer struct {
				Appointments string `json:"appointments"`
			} `json:"officer"`
		} `json:"links"`
		Name        string      `json:"name"`
		Nationality string      `json:"nationality"`
		Occupation  string      `json:"occupation"`
		OfficerRole OfficerRole `json:"officer_role"`
		ResignedOn  ChDate      `json:"resigned_on"`
	}

	// OfficerResponse contains the server response of a data request to the companies house API
	OfficerResponse struct {
		Etag          string    `json:"etag"`
		Kind          string    `json:"kind"`
		Start         int       `json:"start_index"`
		ItemsPerPage  int       `json:"items_per_page"`
		TotalResults  int       `json:"total_results"`
		ActiveCount   int       `json:"active_count"`
		InactiveCount int       `json:"inactive_count"`
		ResignedCount int       `json:"resigned_count"`
		Items         []Officer `json:"items"`
		Links         struct {
			self string `json:"self"`
		} `json:"Links"`
	}
)
