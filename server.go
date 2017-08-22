package main

import (
	//"strings"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
)

var sessionStore *sessions.CookieStore
var pageBuff map[string][]byte

func init() {
	token := make([]byte, 32)
	rand.Seed(time.Now().Unix())
	rand.Read(token)
	log.Println("Session Key:", token)
	sessionStore = sessions.NewCookieStore(token)

	pageBuff = make(map[string][]byte)
	pageBuff["index"] = LoadPage("pages/index.html")
	pageBuff["login"] = LoadPage("pages/page_login.html")
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.PathPrefix("/plugins/").Handler(http.FileServer(http.Dir("static/")))
	router.PathPrefix("/img/").Handler(http.FileServer(http.Dir("static/")))

	//	To login/logout/signup:
	//	/auth/login
	//	/auth/logout
	auth := router.PathPrefix("/auth").Subrouter()
	auth.Path("/login").HandlerFunc(LoginHandler)
	auth.Path("/logout").HandlerFunc(LogoutHandler)

	router.PathPrefix("/ajax").Handler(negroni.New(
		negroni.HandlerFunc(AuthMiddleware),
		negroni.Wrap(http.FileServer(http.Dir("pages/"))),
	))

	router.HandleFunc("/", IndexRoute)
	router.NotFoundHandler = http.HandlerFunc(Page404Route)
	//n := negroni.Classic()
	n := negroni.New(negroni.NewLogger())

	// Doesn't work as expected
	//n.Use(negroni.NewStatic(http.Dir("public")))
	//n.Use(negroni.HandlerFunc(AuthMiddleware))
	n.UseHandler(router)
	n.Run(":4000")
	//log.Println("Starting server on :4000")
	//http.ListenAndServe(":4000", router)
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Check authorize system")

	if isAuth(r) {
		next(w, r)
	} else {
		log.Println("Failed to authorize to the system : requested ", r.URL.Path)
		LoginRoute(w, r)
		return
	}
}

func IndexRoute(res http.ResponseWriter, req *http.Request) {
	if isAuth(req) {
		res.Write(pageBuff["index"])
	} else {
		LoginRoute(res, req)
	}
}
func Page404Route(res http.ResponseWriter, req *http.Request) {
	target := "http://" + req.Host
	http.Redirect(res, req, target, http.StatusTemporaryRedirect)
}
func LoginRoute(res http.ResponseWriter, req *http.Request) {
	res.Write(pageBuff["login"])
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	TemporaryLogInRoute(res, req)
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

func LoadPage(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
func isAuth(req *http.Request) bool {
	// Get a session. Get() always returns a session, even if empty.
	session, err := sessionStore.Get(req, "session-auth")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to read session ", err.Error(), http.StatusInternalServerError)
		return false
	}
	return (session.Values["foo"] == "bar")
}
