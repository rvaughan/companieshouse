package main

import (
	"net/http"
	"fmt"
	"html/template"
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

func main() {
	http.Handle("/static", http.FileServer(http.Dir("static")))
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
		return
	}
	t.ExecuteTemplate(w, "content", nil)
}

