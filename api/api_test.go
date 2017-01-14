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
	"github.com/BalkanTech/companieshouse/testhandlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const testOK = "\u2713"
const testFailed = "\u2717"

var api ChAPI

func init() {
	api.SetAPIKey("12345")
}

// TestURlCreation is a unit test for the urlCreation function
// The test will check if urls are properly formed based upon the apiURL value of the ChAPI struct
// The function consists of two tests, one without overriding the default apiURL (set by the SetAPIKey function)
// and the other test with overriding the apiURL to a server from the httptest package to ensure proper behaviour of unit tests as well
func TestURLCreation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testhandlers.GetResponseHandler))
	defer server.Close()

	var tests = []struct {
		path     string
		expected string
	}{
		{"", "https://api.companieshouse.gov.uk"},
		{"/test", "https://api.companieshouse.gov.uk/test"},
	}

	t.Log("Testing the API's constructUrl function without overriding the apiUrl")
	{
		for _, test := range tests {
			t.Logf("\tChecking input \"%s\".", test.path)
			{
				r := api.constructURL(test.path)
				if r != test.expected {
					t.Fatalf("\t\tExpected \"%s\", got \"%s\". %v", test.expected, r, testFailed)
				}
				t.Logf("\t\tExpected \"%s\", got \"%s\". %v", test.expected, r, testOK)
			}

		}
	}

	api.setAPIURL(server.URL)
	var tests2 = []struct {
		path     string
		expected string
	}{
		{"", server.URL},
		{"/test", server.URL + "/test"},
	}

	t.Log("Testing the API's constructUrl function with overriding the apiUrl")
	{
		for _, test := range tests2 {
			t.Logf("\tChecking input \"%s\".", test.path)
			{
				r := api.constructURL(test.path)
				if r != test.expected {
					t.Fatalf("\t\tExpected \"%s\", got \"%s\". %v", test.expected, r, testFailed)
				}
				t.Logf("\t\tExpected \"%s\", got \"%s\". %v", test.expected, r, testOK)
			}

		}
	}
}

// TestPrepareRequest is a unit test for the prepareRequest function
// The test checks: 1) The ability to make a call 2) If the proper method is set in the response
// 3) If the URL Host is properly set in the response 4) If the reqsponse's header contains proper Authorization data
func TestPrepareRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testhandlers.GetResponseHandler))
	defer server.Close()

	method := "GET"
	url, _ := url.Parse(server.URL)
	header := "Basic MTIzNDU6"

	t.Log("Testing the API prepareRequest function")
	{
		resp, err := api.prepareRequest(server.URL)
		if err != nil {
			t.Fatal("\tShould be able to make the call.", testFailed, err)
		}
		t.Log("\tShould be able to make the call.", testOK)

		if resp.Method != method {
			t.Fatalf("\tShould receive a \"%s\" method, got \"%s\" instead. %v", method, resp.Method, testFailed)
		}
		t.Logf("\tShould receive a \"%s\" method. %v", method, testOK)

		if resp.URL.Host != url.Host {
			t.Fatalf("\tShould receive a \"%s\" URL.Host, got \"%s\" instead. %v", url.Host, resp.URL.Host, testFailed)
		}
		t.Logf("\tShould receive a \"%s\" URL.Host. %v", url.Host, testOK)

		t.Log("\tChecking the header")
		{
			if resp.Header.Get("Authorization") != header {
				t.Fatalf("\t\tShould receive a \"%s\" Authorization, got \"%s\" instead. %v", header, resp.Header["Authorization"][0], testFailed)
			}
			t.Logf("\t\tShould receive a \"%s\" Authorization. %v", header, testOK)
		}
	}
}

// TestPrepareRequest is a unit test for the getResponse function
// The test checks: 1) If the call can be made 2) The response's status 3) The ability to decode the response
// 4) If the decoded response contains the expected content
func TestGetResponse(t *testing.T) {
	statusCode := http.StatusOK

	server := httptest.NewServer(http.HandlerFunc(testhandlers.GetResponseHandler))
	defer server.Close()

	api.setAPIURL(server.URL)

	t.Log("Testing getResponse function")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\" without a content type", server.URL, statusCode)
		{
			resp, err := api.getResponse(server.URL, "")
			{
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", testFailed, err)
				}
				t.Log("\t\tShould be able to make the Get call.", testOK)

				defer resp.Body.Close()

				if resp.StatusCode != statusCode {
					t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, testFailed, resp.StatusCode)
				}
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, testOK)

				j := testhandlers.TestResponse{}

				if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
					t.Fatal("\tShould decode the response.", testFailed)
				}
				t.Log("\tShould decode the response.", testOK)

				if j.Text == "Hello, client" {
					t.Log("\tShould have a Text.", testOK)
				} else {
					t.Error("\tShould have a Text.", testFailed, j.Text)
				}
			}
		}

		contentType := ContentTypePDF
		expectedLocation := "http://www.example.com/application.pdf"

		t.Logf("\tWhen checking \"%s\" for status code \"%d\" with a content type %s", server.URL, statusCode, contentType)
		{
			resp, err := api.getResponse(server.URL, contentType)
			{
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", testFailed, err)
				}
				t.Log("\t\tShould be able to make the Get call.", testOK)

				defer resp.Body.Close()

				if resp.StatusCode != statusCode {
					t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, testFailed, resp.StatusCode)
				}
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, testOK)

				if resp.Header.Get("Content-Type") != ContentTypeTextPlain {
					t.Fatalf("\t\tShould receive a \"%s\" content type, received %s instead. %v", ContentTypeTextPlain, resp.Header.Get("Content-Type"), testFailed)
				}
				t.Logf("\t\tShould receive a \"%s\" content type. %v", ContentTypeTextPlain, testOK)

				if resp.Header.Get("Location") != "http://www.example.com/application.pdf" {
					t.Fatalf("\t\tShould receive a \"%s\" location, received %s instead. %v", expectedLocation, resp.Header.Get("Location"), testFailed)
				}
				t.Logf("\t\tShould receive a \"%s\" location. %v", expectedLocation, testOK)
			}
		}

		url := "https://document-api.companieshouse.gov.uk/document/n1EjP_MALLs8xZp5hs86iHcYDli0TE-n6t4HUDeZuq8/content"
		t.Logf("\tWhen checking \"%s\" for status code \"%d\" with a content type %s to see if overriding works", url, statusCode, contentType)
		{
			resp, err := api.getResponse(url, contentType)
			{
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", testFailed, err)
				}
				t.Log("\t\tShould be able to make the Get call.", testOK)

				defer resp.Body.Close()

				if resp.StatusCode != statusCode {
					t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, testFailed, resp.StatusCode)
				}
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, testOK)

				if resp.Header.Get("Content-Type") != ContentTypeTextPlain {
					t.Fatalf("\t\tShould receive a \"%s\" content type, received %s instead. %v", ContentTypeTextPlain, resp.Header.Get("Content-Type"), testFailed)
				}
				t.Logf("\t\tShould receive a \"%s\" content type. %v", ContentTypeTextPlain, testOK)

				if resp.Header.Get("Location") != "http://www.example.com/application.pdf" {
					t.Fatalf("\t\tShould receive a \"%s\" location, received %s instead. %v", expectedLocation, resp.Header.Get("Location"), testFailed)
				}
				t.Logf("\t\tShould receive a \"%s\" location. %v", expectedLocation, testOK)
			}
		}
	}
}

// TestCallAPI is a unit test for the callAPI function
// The tests checks: 1) If a call can be made 2) If the response can be unmarshaled
// 3) If the unmarshaled response has the expected content
func TestCallAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testhandlers.GetResponseHandler))
	defer server.Close()

	api.setAPIURL(server.URL)

	t.Log("Testing the callAPI function")
	{
		t.Logf("\tWhen checking \"%s\"", server.URL)
		{
			resp, err := api.callAPI(server.URL+"/test", true, ContentTypeJSON)
			{
				if err != nil {
					t.Fatal("\t\tShould be able to make the call.", testFailed, err)
				}
				t.Log("\t\tShould be able to make the call.", testOK)

				j := &testhandlers.TestResponse{}
				if err := json.Unmarshal(resp, &j); err != nil {
					t.Fatal("\tShould unmarshal the response.", testFailed)
				}
				t.Log("\tShould unmarshal the response.", testOK)

				if j.Text == "Hello, client" {
					t.Log("\tShould have a Text.", testOK)
				} else {
					t.Error("\tShould have a Text.", testFailed, j.Text)
				}
			}
		}
	}
}
