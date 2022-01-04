package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(r.Form)
	if err := tpl.ExecuteTemplate(w, "000_form.gohtml", r.Form); err != nil {
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
