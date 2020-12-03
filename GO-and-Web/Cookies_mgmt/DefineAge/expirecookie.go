/////////////////mexpirecookie.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Create 3 actions
///					-- Set the cookie
///					-- Read the cookie and print out his value
///					-- Delete the cookie by changing is MaxAge value
///					-- Use functions
///						--- SetCookie : https://golang.org/pkg/net/http/#SetCookie
///						--- Cookie : https://golang.org/pkg/net/http/#Request.Cookie
///						--- MaxAge : https://golang.org/pkg/net/http/#Cookie
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/setcookie", setcookie)
	http.HandleFunc("/readcookie", readcookie)
	http.HandleFunc("/expirecookie", expirecookie)
	http.ListenAndServe(":8080", nil)

}

func index(varwrite http.ResponseWriter, varreq *http.Request) {
	fmt.Fprintln(varwrite, `<h1><a href="/setcookie">set a cookie</a></h1>`)
}

func setcookie(varwrite http.ResponseWriter, varreq *http.Request) {
	http.SetCookie(varwrite, &http.Cookie{
		Name:  "Cookie-session",
		Value: "KB24",
		Path:  "/",
	})
	fmt.Fprintln(varwrite, `<h1><a href="/readcookie">read</a></h1>`)
}

func readcookie(varwrite http.ResponseWriter, varreq *http.Request) {
	varcookie, err := varreq.Cookie("Cookie-session")
	if err != nil {
		http.Redirect(varwrite, varreq, "/setcookie", http.StatusSeeOther)
		return
	}
	//print out the cookie value
	fmt.Fprintf(varwrite, `<h1>Cookie:<br>%v</h1><h1><a href="/expirecookie">expire</a></h1>`, varcookie)
}

func expirecookie(varwrite http.ResponseWriter, varreq *http.Request) {
	varcookie, err := varreq.Cookie("Cookie-session")
	if err != nil {
		http.Redirect(varwrite, varreq, "/setcookie", http.StatusSeeOther)
		return
	}
	//Delete the cookie by changing the age
	varcookie.MaxAge = -1
	//Initialize cookie
	http.SetCookie(varwrite, varcookie)
	//Go back to main page
	http.Redirect(varwrite, varreq, "/", http.StatusSeeOther)
}
