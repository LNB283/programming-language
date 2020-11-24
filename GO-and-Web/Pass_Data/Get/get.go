/////////////////get.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Create a function to retrieve value from GET method
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", reqvalueextract)
	http.ListenAndServe(":8080", nil)

}

func reqvalueextract(varwrite http.ResponseWriter, varreq *http.Request) {
	//Extract the query
	varvalue := varreq.FormValue("q")
	varwrite.Header().Set("Content-Type", "text/html; charset=utf-8")
	varpostform := `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`
	io.WriteString(varwrite, varpostform+varvalue)
}
