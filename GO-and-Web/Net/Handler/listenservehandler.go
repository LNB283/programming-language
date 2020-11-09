/////////////////listenservehandler.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- listen communication on localhost
///				- Print out on the browser a message
///				- Use Firefox, more convenient than Chrome because chrome generates too many requests
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net/http"
)

type testhandler int

//reference : https://golang.org/src/net/http/server.go#L2041
func (varthandler testhandler) ServeHTTP(varwww http.ResponseWriter, varrequest *http.Request) {
	fmt.Fprintln(varwww, "This is a test")
}

func main() {
	var th testhandler
	// reference : https://golang.org/src/net/http/server.go#L3118
	http.ListenAndServe(":8080", th)
}
