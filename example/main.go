package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func main() {
	http.Handle("/static", http.FileServer(http.Dir("static")))
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("base.html", "root.html")
	if err != nil {
		fmt.Fprintf(w,"Error while parsing template: %s", err.Error())
	}
	t.ExecuteTemplate(w, "content", nil)
}

