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
	return (session.Values["logged"] == "true")
}

func LoginRoute(res http.ResponseWriter, req *http.Request) {
	res.Write(pageBuff["login"])
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if value, ok := configs.Users[req.Form["username"][0]]; (ok && value.Password == req.Form["password"][0]) {
		SetSecureSession(res, req, "logged","true")
		SetSecureSession(res, req, "avatar", value.Avatar)
		SetCookieRealname(res, value.RealName)
		res.Write([]byte("true"))
	} else {
		res.Write([]byte("false"))
	}
}
func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	SetSecureSession(res, req, "logged","false")
}

func SetSecureSession(w http.ResponseWriter, r *http.Request, name string,value string) {
	// Get a session. Get() always returns a session, even if empty.
	session, _ := sessionStore.Get(r, "session-auth")
	// Set some session values.
	session.Values[name] = value
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}

func SetCookieRealname(w http.ResponseWriter, value string) {
	cookie := http.Cookie{Name: "realname", Value: value, Path: "/"}
	http.SetCookie(w, &cookie)
}
