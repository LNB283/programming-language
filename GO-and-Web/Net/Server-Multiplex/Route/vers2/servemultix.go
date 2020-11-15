/////////////////servemultix.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Understand routing
///				- based on the URL , print out the appropriate result
///				- create 2 functions which print out on the browser 2 different results based on the URL
///				- Analyze the URL
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"io"
	"net/http"
)

type typehandler1 int
type typehandler2 int

func (varhand1 typehandler1) ServeHTTP(varresult1 http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varresult1, "Kobe #24")
}

func (varhand2 typehandler2) ServeHTTP(varresult2 http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varresult2, "Jordan #23")
}

func main() {
	var MJ23 typehandler2
	var KB24 typehandler1
	// localhost:8080/MJ : Works good. localhost:8080/MJ/test : Works good too. We give the possibility to analyze the entire URL even if the URL doesn't match exactly
	http.Handle("/MJ/", MJ23)
	// localhost:8080/KB : Works good. localhost:8080/KB/test : Doesn't work. we need exact match with the URL
	http.Handle("/KB", KB24)
	// If URL match : OK. If URL doesn't match, print out 404 error by using nil variable
	http.ListenAndServe(":8080", nil)
}
