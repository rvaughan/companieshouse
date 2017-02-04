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
	"github.com/BalkanTech/companieshouse/api/testhandlers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestCompanies is a slice of an anonymous struct containing the default test data of companies with existing test data
// tests using this slice should ALWAYS pass
var TestCompanies = []struct {
	company string
}{
	{"12345678"},
	{"OC123456"},
	{"FC123456"},
}

// TestGetCompany is a unit test of the GetCompany function
// The test checks:
// 1) Tests which are supposed to pass, consisting of:
//	a) If a call can be made b) If the response contains the expected content
// 2) If the expected error is returned when calling the function with an invalid input data
// 3) If the expected error is returned when a call without proper authentication is made
func TestGetCompany(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testhandlers.GetCompanyHandler))
	defer server.Close()

	api := API{}
	api.SetAPIKey("12345")
	api.setAPIURL(server.URL)

	t.Log("Testing the GetCompany function")
	{
		t.Logf("\tUsing server \"%s\". Testing calls supposed to pass", server.URL)
		{
			for _, test := range TestCompanies {
				c := &Company{api: &api}
				data, _ := testhandlers.GetResponse(test.company, testhandlers.Companies)
				json.Unmarshal([]byte(data), c)

				t.Logf("\t\tWhen checking company\"%s (%s)\"", c.CompanyName, c.CompanyNumber)
				resp, err := api.GetCompany(test.company)
				{
					if err != nil {
						t.Fatalf("\t\t\tShould be able to make the call, but got error: %s. %v", err, testFailed)
					}
					t.Log("\t\t\tShould be able to make the call.", testOK)

					if reflect.DeepEqual(*resp, *c) != true {
						t.Fatalf("\t\t\tShould have proper content, received %+v, but got %+v. %v", resp, c, testFailed)
					}
					t.Log("\t\t\tShould have proper content. ", testOK)
				}
			}
		}

		t.Logf("\tUsing server \"%s\". Testing call with an non-existing company number", server.URL)
		{
			expected := "404 Not Found"
			_, err := api.GetCompany("87654321")
			if err.Error() != expected {
				t.Fatalf("\t\tExpected to fail with \"%s\", but got \"%s\" instead. %v", expected, err, testFailed)
			}
			t.Logf("\t\tExpected to fail with \"%s\". %v", expected, testOK)
		}

		t.Logf("\tUsing server \"%s\". Testing call with invalid authorization", server.URL)
		{
			api.SetAPIKey("invalid")
			api.setAPIURL(server.URL)

			expected := "401 Unauthorized"
			_, err := api.GetCompany("09999801")
			if err.Error() != expected {
				t.Fatalf("\t\tExpected to fail with \"%s\", but got \"%s\" instead. %v", expected, err, testFailed)
			}
			t.Logf("\t\tExpected to fail with \"%s\". %v", expected, testOK)
		}
	}
}
