package main

import (
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
	auth.Path("/login").Methods("POST").HandlerFunc(LoginHandler)
	auth.Path("/logout").Methods("GET").HandlerFunc(LogoutHandler)

	router.PathPrefix("/ajax").Handler(negroni.New(
		negroni.HandlerFunc(AuthMiddleware),
		negroni.Wrap(http.FileServer(http.Dir("pages/"))),
	))

	router.HandleFunc("/", IndexRoute)
	router.NotFoundHandler = http.HandlerFunc(Page404Route)
	//n := negroni.Classic()
	negroniLogger := negroni.NewLogger()
	negroniLogger.SetFormat(" {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} \n")
	n := negroni.New(negroniLogger)

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

func LoadPage(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
