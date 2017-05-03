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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const (
	CompanyNumber        int = 2
	ContentTypeJSON          = "application/json"
	ContentTypePDF           = "application/pdf"
	ContentTypeTextPlain     = "text/plain; charset=utf-8"
)

func getFromPath(p string, e int) (string, error) {
	s := regexp.MustCompile("//*").Split(p, -1)
	if len(s)-1 < e {
		return "", errors.New("Not existing")
	}
	return s[e], nil
}

type TestResponse struct {
	Text string
}

func GetResponseHandler(w http.ResponseWriter, r *http.Request) {
	response := TestResponse{
		Text: "Hello, client",
	}

	if r.Header.Get("Accept") == ContentTypePDF {
		w.Header().Set("Content-type", ContentTypeTextPlain)
		w.Header().Set("Location", "http://www.example.com/application.pdf")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")

	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", ContentTypeJSON)
		json.NewEncoder(w).Encode(&response)
	}
}

func GetResponse(n string, d map[string]string) (string, bool) {
	c, ok := d[n]
	return c, ok
}

var NotFoundResponse string = `{
  "errors": [
    {
      "type": "ch:service",
      "error": "not-found"
    }
  ]
}`

var NotAuthenticatedResponse string = `{
  "errors": [
    {
      "type": "ch:service",
      "error": "not-authenticated"
    }
  ]
}`

func Authenticated(r *http.Request) bool {
	return r.Header.Get("Authorization") == "Basic MTIzNDU6"
}

func GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, NotAuthenticatedResponse)
		return
	}

	w.Header().Set("Content-type", ContentTypeJSON)

	n, err := getFromPath(r.URL.Path, CompanyNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	resp, ok := GetResponse(n, Companies)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
	return
}

func GetOfficersHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, NotAuthenticatedResponse)
		return
	}

	w.Header().Set("Content-type", ContentTypeJSON)

	n, err := getFromPath(r.URL.Path, CompanyNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	resp, ok := GetResponse(n, Officers)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
	return
}

func GetInsolvenciesHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, NotAuthenticatedResponse)
		return
	}

	w.Header().Set("Content-type", ContentTypeJSON)

	n, err := getFromPath(r.URL.Path, CompanyNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	resp, ok := GetResponse(n, Insolvencies)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
	return
}

func GetChargesHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, NotAuthenticatedResponse)
		return
	}

	w.Header().Set("Content-type", ContentTypeJSON)

	n, err := getFromPath(r.URL.Path, CompanyNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	resp, ok := GetResponse(n, Charges)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
	return
}

func GetFilingsHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, NotAuthenticatedResponse)
		return
	}

	w.Header().Set("Content-type", ContentTypeJSON)

	n, err := getFromPath(r.URL.Path, CompanyNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	resp, ok := GetResponse(n, Filings)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, NotFoundResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
	return
}

//Todo: testhandlers.go:30:2: exported const CompanyNumber should have comment (or a comment on this block) or be unexported
//Todo: testhandlers.go:44:6: exported type TestResponse should have comment or be unexported
//Todo: testhandlers.go:48:1: exported function GetResponseHandler should have comment or be unexported
//Todo: testhandlers.go:66:1: exported function GetResponse should have comment or be unexported
//Todo: testhandlers.go:71:5: exported var NotFoundResponse should have comment or be unexported
//Todo: testhandlers.go:71:22: should omit type string from declaration of var NotFoundResponse; it will be inferred from the right-hand side
//Todo: testhandlers.go:80:5: exported var NotAuthenticatedResponse should have comment or be unexported
//Todo: testhandlers.go:80:30: should omit type string from declaration of var NotAuthenticatedResponse; it will be inferred from the right-hand side
//Todo: testhandlers.go:89:1: exported function Authenticated should have comment or be unexported
//Todo: testhandlers.go:93:1: exported function GetCompanyHandler should have comment or be unexported
//Todo: testhandlers.go:121:1: exported function GetOfficersHandler should have comment or be unexported
//Todo: testhandlers.go:149:1: exported function GetInsolvenciesHandler should have comment or be unexported
//Todo: testhandlers.go:177:1: exported function GetChargesHandler should have comment or be unexported
//Todo: testhandlers.go:205:1: exported function GetFilingsHandler should have comment or be unexported
