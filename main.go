package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const BaseUrl = "https://api.fixer.io/"

type Result struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	resp, err := http.Get(BaseUrl + "latest")
	if err != nil {
		log.Panicf("Execute GET-request error: %v", err)
	}

	decoder := json.NewDecoder(resp.Body)
	result := Result{}
	if err := decoder.Decode(&result); err != nil {
		log.Panicf("Parse response of GET-request error: %v", err)
	}

	log.Printf("Base: %v", result.Base)
	log.Printf("Date: %v", result.Date)
	log.Println("Rates: [")
	for k, v := range result.Rates {
		log.Printf("\t%v: %v", k, v)
	}
	log.Println("]")
}
