package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//func getUser(req *http.Request) user {
func getUser(varwrite http.ResponseWriter, varreq *http.Request) user {
	//var u user
	// get cookie
	c, err := varreq.Cookie("session")
	if err != nil {
		//return u
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(varwrite, c)

	var u user
	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
