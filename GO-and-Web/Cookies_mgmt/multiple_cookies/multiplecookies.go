/////////////////multiplecookies.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Print out the cookie informaiton
///					-- Use functions
///						--- SetCookie : https://golang.org/pkg/net/http/#SetCookie
///						--- Cookie : https://golang.org/pkg/net/http/#Request.Cookie
///				- Create and print out multiple cookies values
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", setcookie)
	http.HandleFunc("/readcookie", readcookie)
	http.HandleFunc("/multiplecookies", multiplecookie)
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

	varcookieKB, err := varreq.Cookie("Cookie-Kobe")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(varwrite, "Cookie KB Value:", varcookieKB)
	}

	varcookieMJ, err := varreq.Cookie("Cookie-Jordan")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(varwrite, "Cookie MJ Value:", varcookieMJ)
	}

	varcookieRH, err := varreq.Cookie("Cookie-Hachimura")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(varwrite, "Cookie RH Value:", varcookieRH)
	}

	varcookieTP, err := varreq.Cookie("Cookie-Parker")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(varwrite, "Cookie RH Value:", varcookieTP)
	}
}

//Create multiple cookies
func multiplecookie(varwrite http.ResponseWriter, varreq *http.Request) {

	http.SetCookie(varwrite, &http.Cookie{
		Name:  "Cookie-Jordan",
		Value: "MJ23",
	})

	http.SetCookie(varwrite, &http.Cookie{
		Name:  "Cookie-Hachimura",
		Value: "RH8",
	})

	http.SetCookie(varwrite, &http.Cookie{
		Name:  "Cookie-Parker",
		Value: "TP9",
	})

	fmt.Fprintln(varwrite, "Cookie created, Follow the path to check cookies")
	fmt.Fprintln(varwrite, "Firefox : Tool --> Web Developer --> Storage Inspector")
}
