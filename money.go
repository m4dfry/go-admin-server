package main

import "net/http"

func MoneyPostHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//req.Form["username"][0]

	ok := true
	if ok {

		res.Write([]byte("true"))
	} else {
		res.Write([]byte("false"))
	}
}

func MoneyGetHandler(res http.ResponseWriter, req *http.Request) {

}

