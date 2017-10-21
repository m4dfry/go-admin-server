package settings

import (
	"log"
	"encoding/json"
	"os"
	"bufio"
)

const CONFIG_DEFAULT_PATH = "config.json"

type User struct {
	RealName string `json:"realname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type Config struct {
	Users		map[string](User)
	Port       int    `json:"port"`
	Address    string `json:"address"`
	LogNegroni bool   `json:"log-negroni"`
}

func Init(cp *string) *Config{
	log.Println("Init settings.")
	a := ReadConfig(cp)
	return a
}

func ReadConfig(cp *string) (*Config) { // Use json.Decode for reading streams of JSON data
	if *cp == "" { *cp = CONFIG_DEFAULT_PATH }
	log.Println("File config:", *cp)
	f, err := os.Open(*cp)
	if err != nil {
		panic(err)
	}
	f_reader := bufio.NewReader(f)

	var config Config
	if err := json.NewDecoder(f_reader).Decode(&config); err != nil {
		log.Println(err)
	}

	return &config//, users
}