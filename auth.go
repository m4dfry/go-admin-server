package main

import (
	"log"
	"net/http"
	//"encoding/json"
)

func isAuth(req *http.Request) bool {
	// Get a session. Get() always returns a session, even if empty.
	session, err := sessionStore.Get(req, "session-auth")
	if err != nil {
		log.Println("Failed to read session. Reset.", err.Error(), http.StatusInternalServerError)
		return false
	}
	return (session.Values["foo"] == "bar")
}

func LoginRoute(res http.ResponseWriter, req *http.Request) {
	res.Write(pageBuff["login"])
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	username := req.Form["username"][0]
	password := req.Form["password"][0]
	log.Println("U:", username, "-P:", password)
	//LOG: that
	//TemporaryLogInRoute(res, req) */
}
func LogoutHandler(res http.ResponseWriter, req *http.Request) {
}
func TemporaryLogInRoute(w http.ResponseWriter, r *http.Request) {

	// Get a session. Get() always returns a session, even if empty.
	session, err := sessionStore.Get(r, "session-auth")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
	}

	// Set some session values.
	session.Values["foo"] = "bar"
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	w.Write([]byte("Login Success"))
}
