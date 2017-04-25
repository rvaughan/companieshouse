package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
	"github.com/BalkanTech/companieshouse/api"
	"github.com/gorilla/mux"
	"log"
)

type Alert struct {
	Title string
	Class string
	Message string
}

type Data struct {
	Alerts []Alert
	Data map[string]interface{}
}

func newData() *Data {
	return &Data{Data: make(map[string]interface{})}
}

var ch = companieshouse.CompaniesHouseAPI(os.Getenv("CH_API_KEY"))

func main() {
	r := mux.NewRouter()
	r.Handle("/static", http.FileServer(http.Dir("static")))
	r.HandleFunc("/favicon.ico", http.NotFound)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/company/{id}", companyHandler)

	http.ListenAndServe(":8000", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := newData()
	t, err := template.ParseFiles("templates/base.html", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
		return
	}

	log.Printf("Got a \"%s\" request in rootHandler", r.Method)
	if r.Method == "POST" {
		q := r.FormValue("SearchText")
		f := r.FormValue("SearchFor")
		if q == "" || f == "" {
			data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", "Enter a search text and search type"})
			t.ExecuteTemplate(w, "base", data)
			return
		}

		switch f {
		case "C":
			r, err := ch.SearchCompany(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["Companies"] = *r
			}
		case "O":
			r, err := ch.SearchOfficer(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["Officers"] = *r
			}
		case "D":
			r, err := ch.SearchDisqualifiedOfficer(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["DisqOfficers"] = *r
			}
		default:
			data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", "Invalid search type"})
		}
	}
	t.ExecuteTemplate(w, "base", data)
}

func companyHandler(w http.ResponseWriter, r *http.Request) {
	data := newData()
	v := mux.Vars(r)
	id, ok := v["id"]
	if !ok {
		http.NotFound(w, r)
		return
	}
	company, err := ch.GetCompany(id)
	if err != nil {
		data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
	} else {
		data.Data["Company"] = company

		officers, err := company.GetOfficers()
		if err == nil {
			data.Data["Officers"] = officers
		} else {
			log.Println("Error when getting officers:", err)
		}


		filings, err := company.GetFilings()
		if err == nil {
			data.Data["Filings"] = filings
		} else {
			log.Println("Error when getting filings:", err)
		}

		charges, err := company.GetCharges()
		if err == nil {
			data.Data["Charges"] = charges
		} else {
			log.Println("Error when getting charges:", err)
		}


		insolvencyDetails, err := company.GetInsolvencyDetails()
		if err == nil {
			data.Data["InsolvencyDetails"] = insolvencyDetails
		} else {
			log.Println("Error while getting insolvency details:", err)
		}
	}

	t, err := template.ParseFiles("templates/base.html", "templates/company.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
		return
	}
	t.ExecuteTemplate(w, "base", data)
}

