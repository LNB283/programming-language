/////////////////form.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- create html templates
///					-- Header: contain information
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"net/http"
	"text/template"
)

var vartemplate *template.Template

func init() {
	vartemplate = template.Must(template.ParseGlob("Templates/*"))
}

type nbaplayer struct {
	FirstName string
	LastName  string
	Number    string
	Retired   bool
}

func main() {
	http.HandleFunc("/", reqexec)
	http.ListenAndServe(":8080", nil)
}

func reqexec(varwrite http.ResponseWriter, varreq *http.Request) {
	fname := varreq.FormValue("FName")
	lname := varreq.FormValue("LName")
	numb := varreq.FormValue(("Number"))
	infoRetired := varreq.FormValue("Retired") == "on"

	err := vartemplate.ExecuteTemplate(varwrite, "index.gohtml", nbaplayer{fname, lname, numb, infoRetired})
	if err != nil {
		http.Error(varwrite, err.Error(), 500)
		log.Fatalln(err)
	}
}
