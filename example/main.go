package main

import (
	"fmt"
	"github.com/BalkanTech/companieshouse/api"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"bytes"
)

type Alert struct {
	Title   string
	Class   string
	Message string
}

type Data struct {
	Alerts []Alert
	Data   map[string]interface{}
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
	r.HandleFunc("/document/{cid}/{fid}", showFilingHandler)

	http.ListenAndServe(":8000", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := newData()
	t, err := template.ParseFiles("templates/base.html", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w, "Error while parsing template: %s", err.Error())
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
	t, err := template.ParseFiles(
		"templates/base.html",
		"templates/company.html",
		"templates/tab_general.html",
		"templates/tab_officers.html",
		"templates/tab_filings.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error while parsing template: %s", err.Error())
		return
	}

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
		data.Data["Company"] = *company
	}

	t.ExecuteTemplate(w, "base", data)
}

func showFilingHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	cid, ok := v["cid"]
	if !ok {
		http.NotFound(w, r)
		return
	}

	fid, ok := v["fid"]
	if !ok {
		http.NotFound(w, r)
		return
	}

	c, err := ch.GetCompany(cid)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	f, err := c.GetFiling(fid)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	d, err := c.GetDocument(f)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	b := bytes.NewBuffer(d)

	// stream straight to client(browser)
	w.Header().Set("Content-type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s.%s", f.Barcode, "pdf"))

	if _, err := b.WriteTo(w); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
