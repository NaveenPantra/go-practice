package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Error: ", err)
		return
	}
	data := struct {
		Method  string
		QParams url.Values
		Url     *url.URL
	}{
		r.Method,
		r.Form,
		r.URL,
	}
	if err := tpl.ExecuteTemplate(w, "000_form.gohtml", data); err != nil {
		log.Println("Error", err)
		return
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("000_form.gohtml"))
}

func main() {
	var s server

	http.ListenAndServe(":8080", s)
}
