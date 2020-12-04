/////////////////id.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- ListeandServer on port 8080 - locaslhost
///				- Print out the session ID information on the terminal
///					-- Use functions
///						--- SetCookie : https://golang.org/pkg/net/http/#SetCookie
///						--- Cookie : https://golang.org/pkg/net/http/#Request.Cookie
///						--- package requires :  github.com/satori/go.uuid (perform the command : go get github.com/satori/go.uuid)
///
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	id := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", id)
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
