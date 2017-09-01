package settings


import (
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

type Config struct {
	Users []struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
	} `json:"users"`
	Port       int    `json:"port"`
	Address    string `json:"address"`
	LogNegroni bool   `json:"log-negroni"`
}

func Init(){
	log.Println("Init settings.")
	ReadConfig()
}

func ReadConfig() { // Use json.Decode for reading streams of JSON data
	f, err := os.Open("config.json")
	if err != nil {
		panic(e)
	}
	f_reader := bufio.NewReader(f)

	var config Config
	if err := json.NewDecoder(f_reader).Decode(config); err != nil {
		log.Println(err)
	}
}
/*
JSON TEMPLATE FOR CONFIG (WIP)
{
  "users": [
     {"username": "user1", "password": "thapassword",
                  "avatar": "path/for/avatar"
    },
    {"username": "user2", "password": "tha123235534",
                  "avatar": "path/for/avatar"
    },
    {"username": "user3", "password": "£$DFSFFS",
                  "avatar": "path/for/avatar"
    }
  ],
  "port" : 4000,
  "address" : "",
  "log-negroni": true
}

*/