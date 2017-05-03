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

package testhandlers

import (
	"testing"
)

const testOK = "\u2713"
const testFailed = "\u2717"

func TestGetFromPath(t *testing.T) {
	var tests = []struct {
		path         string
		element      int
		expected     string
		expect_error bool
	}{
		{"/company/09999801", CompanyNumber, "09999801", false},
		{"/company/09999801/officers", CompanyNumber, "09999801", false},
		{"/company", CompanyNumber, "", true},
	}

	t.Log("Testing the getFromPath function")
	{
		for _, test := range tests {
			t.Logf("\tWhen checking url \"%s\"", test.path)
			resp, err := getFromPath(test.path, test.element)
			{
				if err != nil && !test.expect_error {
					t.Fatalf("\t\t\tShould return error: %t. %v", test.expect_error, testFailed)
				}
				t.Logf("\t\tShould return error: %t. %v", test.expect_error, testOK)

				if test.expected != resp {
					t.Fatalf("\t\t\tShould receive \"%s\", but got %s. %v", test.expected, resp, testFailed)
				}
				t.Logf("\t\tShould receive \"%s\". %v", test.expected, testOK)
			}
		}
	}
}

//Todo: testhandlers_test.go:33:3: don't use underscores in Go names; struct field expect_error should be expectError
