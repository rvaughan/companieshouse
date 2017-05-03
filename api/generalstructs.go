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

type (
	// Address struct contains the details of addresses
	Address struct {
		Premises     string `json:"premises"`
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		Locality     string `json:"locality"`
		Region       string `json:"region"`
		PostalCode   string `json:"postal_code"`
		Country      string `json:"country"`
		CareOf       string `json:"care_of"`
		PoBox        string `json:"po_box"`
	}

	// Links struct contains links to additional data and the original object (Self)
	Links struct {
		Psc              string `json:"persons_with_significant_control"`
		PscStatements    string `json:"persons_with_significant_control_statements`
		Registers        string `json:"registers"`
		UkEstablishments string `json:"uk_establishments"`
		DocumentMetaData string `json:"document_metadata"`
		Document         string `json:"document"`
		Self             string `json:"self"`
	}
)
