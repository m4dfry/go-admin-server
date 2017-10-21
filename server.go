package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"flag"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"

	"github.com/m4dfry/go-admin-server/settings"
	"github.com/m4dfry/go-admin-server/plugins"
)

var sessionStore *sessions.CookieStore
var pageBuff map[string][]byte
var configs *settings.Config

func usage(){
	// DO NOTHING FOR NOW

}

func init() {
	var config_path = flag.String("c", "", "Custom configuration file path")

	flag.Usage = func() {
		usage()
		flag.PrintDefaults()
	}

	flag.Parse()

	token := make([]byte, 32)
	rand.Seed(time.Now().Unix())
	rand.Read(token)
	log.Println("Session Key:", token)
	sessionStore = sessions.NewCookieStore(token)

	configs = settings.Init(config_path)

	pageBuff = make(map[string][]byte)
	pageBuff["index"] = LoadFile("pages/index.html")
	pageBuff["login"] = LoadFile("pages/page_login.html")
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.PathPrefix("/plugins/").Handler(http.FileServer(http.Dir("static/")))
	router.PathPrefix("/img/").Handler(http.FileServer(http.Dir("static/")))

	user := router.PathPrefix("/user").Subrouter()
	user.Path("/avatar").Methods("GET").HandlerFunc(UserAvatar)

	//	To login/logout/signup:
	//	/auth/login
	//	/auth/logout
	auth := router.PathPrefix("/auth").Subrouter()
	auth.Path("/login").Methods("POST").HandlerFunc(LoginHandler)
	auth.Path("/logout").Methods("GET").HandlerFunc(LogoutHandler)

	plugin := router.PathPrefix("/plugin").Subrouter()
	plugin.Path("/test").Methods("POST").HandlerFunc(plugins.PluginHandler)

	router.PathPrefix("/ajax").Handler(negroni.New(
		negroni.HandlerFunc(AuthMiddleware),
		negroni.Wrap(http.FileServer(http.Dir("pages/"))),
	))

	router.HandleFunc("/", IndexRoute)
	router.NotFoundHandler = http.HandlerFunc(PageRootRoute)

	//n := negroni.Classic()
	n := negroni.New()
	if configs.LogNegroni {
		negroniLogger := negroni.NewLogger()
		negroniLogger.SetFormat(" {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} \n")
		n.Use(negroniLogger)
	}

	// Doesn't work as expected
	//n.Use(negroni.NewStatic(http.Dir("public")))
	//n.Use(negroni.HandlerFunc(AuthMiddleware))
	n.UseHandler(router)
	n.Run(configs.Address + ":" + strconv.Itoa(configs.Port))
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
func PageRootRoute(res http.ResponseWriter, req *http.Request) {
	target := "http://" + req.Host
	http.Redirect(res, req, target, http.StatusTemporaryRedirect)
}

func UserAvatar(res http.ResponseWriter, req *http.Request) {
	session, _ := sessionStore.Get(req, "session-auth")
	res.Write(LoadFile(session.Values["avatar"].(string)))
}

func LoadFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
