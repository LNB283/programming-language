/////////////////ShowPicture.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Show a picture when the template is executed
///					-- Use http.ServeFile
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootpage)
	http.HandleFunc("/Kobe/", kobemindset)
	http.HandleFunc("/KobeMindset1.jpg", picture)
	http.ListenAndServe(":8080", nil)

}

func rootpage(varwrite http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varwrite, "Root Page")
}

func kobemindset(varwrite http.ResponseWriter, varreq *http.Request) {
	vartemplate, err := template.ParseFiles("kobe.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	vartemplate.ExecuteTemplate(varwrite, "kobe.gohtml", nil)
}

func picture(varwrite http.ResponseWriter, varreq *http.Request) {
	http.ServeFile(varwrite, varreq, "KobeMindset1.jpg")
}
