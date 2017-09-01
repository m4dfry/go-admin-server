package settings


import (
	"fmt"
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