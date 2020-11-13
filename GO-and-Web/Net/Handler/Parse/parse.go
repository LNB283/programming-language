/////////////////parse.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Based on the information input on the index.gohtml
///				- Gather the information from the request
///				- Parse the index.gothml file
///				- Write the answer
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"log"
	"net/http"
)

var vartemplate *template.Template

type vartest int

func init() {
	vartemplate = template.Must(template.ParseFiles("index.gohtml"))
}

func (test vartest) ServeHTTP(wreq http.ResponseWriter, varreq *http.Request) {
	err := varreq.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	vartemplate.ExecuteTemplate(wreq, "index.gohtml", varreq.Form)

}

func main() {

	var test vartest
	http.ListenAndServe(":8080", test)
}
