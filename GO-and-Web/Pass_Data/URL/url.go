/////////////////url.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Create a function to extract the query from the URL
///					-- Use the function FormValue
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", reqvalueextract)
	http.ListenAndServe(":8080", nil)

}

func reqvalueextract(varwrite http.ResponseWriter, varreq *http.Request) {
	//Extract the query
	varvalue := varreq.FormValue("q")
	fmt.Fprintln(varwrite, "Query:"+varvalue)

}
