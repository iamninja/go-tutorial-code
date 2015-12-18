package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Name		string `json:"SiteName"`
	URL 		string `json:"SiteURL"`
	Database	struct {
		Name 		string
		Host		string
		Port		int
		Username	string
		Password	string
	}
}

func main() {
	conf := Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil { panic(err) }

	err = json.Unmarshal(data, &conf)
	if err != nil { panic(err) }

	fmt.Printf("Site: %s (%s)\n", conf.Name, conf.URL)

	db := conf.Database
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}