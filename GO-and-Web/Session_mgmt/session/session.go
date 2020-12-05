/////////////////session.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create 2 html page
///					-- index : collect user info
///					-- info : print out the user info
///				- session.go
///					-- collect information
///					-- save the info by using a map of string and type string
///					-- save the session by using a map of string and type string
////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var vartemplate *template.Template

//userID , user
var dbUsers = map[string]user{}

//sessionID , userID
var dbSessions = map[string]string{}

func init() {
	vartemplate = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/info", info)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(varwrite http.ResponseWriter, varreq *http.Request) {

	// get cookie
	varcookie, err := varreq.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		varcookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(varwrite, varcookie)
	}

	// if the user exists already, get user
	var varuser user
	if varusername, ok := dbSessions[varcookie.Value]; ok {
		varuser = dbUsers[varusername]
	}

	// process form submission
	if varreq.Method == http.MethodPost {
		varusername := varreq.FormValue("username")
		varfirstname := varreq.FormValue("firstname")
		varlastname := varreq.FormValue("lastname")
		varuser = user{varusername, varfirstname, varlastname}
		dbSessions[varcookie.Value] = varusername
		dbUsers[varusername] = varuser
	}

	vartemplate.ExecuteTemplate(varwrite, "index.gohtml", varuser)
}

func info(varwrite http.ResponseWriter, varreq *http.Request) {

	// get cookie
	varcookie, err := varreq.Cookie("session")
	if err != nil {
		http.Redirect(varwrite, varreq, "/", http.StatusSeeOther)
		return
	}
	varusername, ok := dbSessions[varcookie.Value]
	if !ok {
		http.Redirect(varwrite, varreq, "/", http.StatusSeeOther)
		return
	}
	varuser := dbUsers[varusername]
	vartemplate.ExecuteTemplate(varwrite, "info.gohtml", varuser)
}
