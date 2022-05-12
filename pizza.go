package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type loljson struct {
	data data `json:"data"`
	meta meta   `json:"meta"`
}

type data struct {
	id                 string
	date               string
	home_team          teams
	home_team_score    int
	period             int
	postseason         bool
	season             int
	status             string
	time               string
	vistor_team        teams
	visitor_team_score int
}

type meta struct {
	total_pages  int
	current_page int
	next_page    int
	per_page     int
	total_count  int
}

type teams struct {
	id           int
	abbreviation string
	city         string
	conference   string
	division     string
	full_name    string
	name         string
}

func pizzaurl() (pizzaurl string) {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("https://www.balldontlie.io/api/v1/games?seasons[]=%d&team_ids[]=16&dates[]=%d-%d-%d", year-1, year, int(month), day)
}

func getresult() (result string) {
	url := pizzaurl()
	println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(body))
	jr := &loljson{}
	json.Unmarshal(body, &jr)
	fmt.Println(jr)
	return "lol"
}

func main() {
	println("Hi")
	getresult()
}
