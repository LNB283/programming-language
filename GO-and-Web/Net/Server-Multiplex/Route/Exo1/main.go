/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Understand routing
///				- based on the URL , print out the appropriate result
///				- create 3 functions which print out on the browser 3 different results based on the URL
///				- Analyze the URL
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"io"
	"net/http"
)

type typehandler1 int
type typehandler2 int
type typehandler3 int

func (varhand1 typehandler1) ServeHTTP(varresult1 http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varresult1, "Kobe #24")
}

func (varhand2 typehandler2) ServeHTTP(varresult2 http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varresult2, "Jordan #23")
}

func (varhand3 typehandler3) ServeHTTP(varresult3 http.ResponseWriter, varreq *http.Request) {
	io.WriteString(varresult3, "root")
}

func main() {

	var KB24 typehandler1
	var MJ23 typehandler2
	var root typehandler3
	http.Handle("/KB24", KB24)
	http.Handle("/MJ23", MJ23)
	http.Handle("/", root)
	http.ListenAndServe(":8080", nil)
}
