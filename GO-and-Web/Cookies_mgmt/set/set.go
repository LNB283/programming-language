/////////////////set.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Print out the cookie informaiton
///					-- Use functions
///						--- SetCookie : https://golang.org/pkg/net/http/#SetCookie
///						--- Cookie : https://golang.org/pkg/net/http/#Request.Cookie
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setcookie)
	http.HandleFunc("/readcookie", readcookie)
	http.ListenAndServe(":8080", nil)
}

//Create the cookie
func setcookie(varwrite http.ResponseWriter, varreq *http.Request) {
	http.SetCookie(varwrite, &http.Cookie{
		Name:  "Cookie-Kobe",
		Value: "KB24",
		Path:  "/",
	})
	fmt.Fprintln(varwrite, "Cookie created, Follow the path to check cookies")
	fmt.Fprintln(varwrite, "Firefox : Tool --> Web Developer --> Storage Inspector")
}

//Read the cookie
func readcookie(varwrite http.ResponseWriter, varreq *http.Request) {

	varcookie, err := varreq.Cookie("Cookie-Kobe")
	if err != nil {
		http.Error(varwrite, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(varwrite, "Cookie Value:", varcookie)
}
