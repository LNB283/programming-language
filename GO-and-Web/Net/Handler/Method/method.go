package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type vartest int

func (testreq vartest) ServeHTTP(wreq http.ResponseWriter, varreq *http.Request) {
	err := varreq.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	vardata := struct {
		Method      string
		Submissions url.Values
	}{
		varreq.Method,
		varreq.Form,
	}
	vartemplate.ExecuteTemplate(wreq, "index.gohtml", vardata)
}

var vartemplate *template.Template

func init() {
	vartemplate = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d vartest
	http.ListenAndServe(":8080", d)
}
