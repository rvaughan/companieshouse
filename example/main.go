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
	Text string
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
	log.Println("Got a \"%s\" request in rootHandler", r.Method)
	if r.Method == "POST" {
		q := r.FormValue("SearchText")
		t := r.FormValue("SearchFor")
		if q == "" || t == "" {
			data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", "Enter a search text and search type"})
		}

		switch t {
		case "C":
			r, err := ch.SearchCompany(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["Companies"] = r
			}
		case "O":
			r, err := ch.SearchOfficer(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["Officers"] = r
			}
		case "D":
			r, err := ch.SearchDisqualifiedOfficer(q, 0, 0)
			if err != nil {
				data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", err.Error()})
			} else {
				data.Data["DiscOfficers"] = r
			}
		default:
			data.Alerts = append(data.Alerts, Alert{"Error", "alert-danger", "Invalid search type"})
		}
	}
	t, err := template.ParseFiles("templates/base.html", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
		return
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
		}

		filings, err := company.GetFilings()
		if err == nil {
			data.Data["Filings"] = filings
		}
	}

	t, err := template.ParseFiles("templates/base.html", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
		return
	}
	t.ExecuteTemplate(w, "base", data)
}

