/////////////////startingfilesfolder2.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Serve the files on Public/pictures folder
///					-- func StripPrefix(prefix string, h Handler) Handler
///					-- func FileServer(root FileSystem) Handler
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"log"
	"net/http"
)

var vartemplate *template.Template

func init() {

	vartemplate = template.Must(vartemplate.ParseFiles("Kobe.gohtml"))
}

func main() {
	fileserv := http.FileServer(http.Dir("Public"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileserv))
	http.HandleFunc("/", kobe)
	http.ListenAndServe(":8080", nil)

}

func kobe(varwrite http.ResponseWriter, varreq *http.Request) {
	err := vartemplate.Execute(varwrite, nil)
	if err != nil {
		log.Fatalln("Template execution issue", err)
	}
}
