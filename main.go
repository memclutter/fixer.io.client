package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

const BaseUrl = "https://api.fixer.io/"

type Result struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

var (
	period  string
	base    string
	symbols string
)

func init() {
	flag.StringVar(&period, "period", "latest", "Get historical rates for any day since 1999. Default latest")
	flag.StringVar(&base, "base", "EUR", "Rates are quoted against the Euro by default. ")
	flag.StringVar(&symbols, "symbols", "", "Request specific exchange rates by setting the symbols parameter.")
	flag.Parse()
}

func main() {
	req, err := http.NewRequest("GET", BaseUrl+period, nil)
	if err != nil {
		log.Panicf("Create GET-request error: %v", err)
	}

	q := req.URL.Query()
	q.Add("base", base)
	q.Add("symbols", symbols)

	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
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
