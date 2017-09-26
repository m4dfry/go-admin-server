package main

import (
	"net/http"
)

func MoneyPostHandler(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	//TO-DO
	res.Write([]byte("true" + req.Form["date"][0]))

}

func MoneyGetHandler(res http.ResponseWriter, req *http.Request) {

}

