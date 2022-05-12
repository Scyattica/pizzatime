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
	Data []data `json:"data"`
	Meta meta   `json:"meta"`
}

type data struct {
	Id                 int    `json:"id"`
	Date               string `json:"date"`
	Home_team          teams  `json:"home_team"`
	Home_team_score    int    `json:"home_team_score"`
	Period             int    `json:"period"`
	Postseason         bool   `json:"postseason"`
	Season             int    `json:"season"`
	Status             string `json:"status"`
	Time               string `json:"time"`
	Vistor_team        teams  `json:"visitor_team"`
	Visitor_team_score int    `json:"visitor_team_score"`
}

type meta struct {
	Total_pages  int
	Current_page int
	Next_page    int
	Per_page     int
	Total_count  int
}

type teams struct {
	Id           int `json:"id"`
	Abbreviation string
	City         string
	Conference   string
	Division     string
	Full_name    string
	Name         string
}

func pizzaurl() (pizzaurl string) {
	time := time.Now()
	return fmt.Sprintf("https://www.balldontlie.io/api/v1/games?team_ids[]=16&dates[]=%d-%d-%d", time.Year(), time.Month(), time.AddDate(0, 0, -1).Day())
}

func getresult() (result *loljson) {
	url := pizzaurl()
	//fmt.Printf(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	jr := &loljson{}
	json.Unmarshal(body, &jr)
	return jr
}

func main() {
	parseresults(getresult())
}

func parseresults(uj *loljson) (pizzatime bool) {
	if uj.Meta.Total_count != 1 {
		fmt.Printf("I don't think the heat played yesterday.")
		return false // heat didn't play.
	} else {
		if uj.Data[0].Home_team.Name == "Heat" {
			if uj.Data[0].Home_team_score > uj.Data[0].Visitor_team_score {
				fmt.Println("Heat won! probably should buy pizza.")
			}
		} else if uj.Data[0].Vistor_team.Name == "Heat" {
			if uj.Data[0].Visitor_team_score > uj.Data[0].Home_team_score {
				fmt.Println("Heat won! probably should buy pizza.")
			}
		} else {
			fmt.Println("Heat didn't play/something else happened. ")
		}
	}
	return false
}
