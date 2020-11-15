/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Understand routing
///				- based on the URL , print out the appropriate result
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"io"
	"net/http"
)

type typehandler int

func (handler typehandler) ServeHTTP(varwrite http.ResponseWriter, varreq *http.Request) {

	switch varreq.URL.Path {
	case "/Kobe":
		io.WriteString(varwrite, "Kobe, number #24")
	case "/Michael":
		io.WriteString(varwrite, "Jordan, number #23")
	}
}

func main() {

	var varhd typehandler
	http.ListenAndServe(":8080", varhd)

}
