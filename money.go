package main

import (
	"net/http"
	"log"

	"gopkg.in/mgo.v2"
	"net"
	"time"
	"crypto/tls"

	"fmt"
	"strings"
	"math/big"
)

type Reg struct {
	Date int64
	Descr string
	Type string
	Amount float64
}

func MoneyPostHandler(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	sDate := req.Form["date"][0]
	descr := req.Form["description"][0]
	typereg := req.Form["type"][0]
	sAmount := strings.TrimSpace(req.Form["amount"][0])

	log.Println("MONEY POTS REG LOG \nDate:" + sDate + "\n"+
		"Description:" + descr + "\n" +
		"Type:" + typereg + "\n" +
		"Amount:" + req.Form["amount"][0] + "-\n" )

	// CONVERT DATE FROM STRING
	conv := "2006-01-02"
	tDate, err := time.Parse(conv , req.Form["date"][0])

	if err != nil {
		fmt.Println(err)
	}

	// CONVERT AMOUNT FROM STRIN
	fAmount := new(big.Float)
	_ , err = fmt.Sscan(sAmount, fAmount)
	if err != nil {
		log.Println("error scanning value:", err)
	}
	f64Amount, _ := fAmount.Float64()

	dialInfo := &mgo.DialInfo{
		Addrs:    configs.Money.Addrs,
		Database: configs.Money.Database ,
		Username: configs.Money.Username,
		Password: configs.Money.Password,
		Source: configs.Money.Source,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Timeout: time.Second * 10,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	//session, err := mgo.Dial("mongodb://:@?ssl=true&replicaSet=&authSource=admin")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("money").C("reg")
	err = c.Insert(&Reg{tDate.Unix(), descr, typereg,f64Amount})
	if err != nil {
		log.Fatal(err)
	}

	res.Write([]byte("true"))
}

func MoneyGetHandler(res http.ResponseWriter, req *http.Request) {

}

