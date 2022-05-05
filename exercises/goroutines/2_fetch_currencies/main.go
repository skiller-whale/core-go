package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func noError(e error) {
	if e != nil {
		panic(e)
	}
}

func oneHundredPoundsIn(currencySymbol string) {
	resp, err := http.Get("https://api.frankfurter.app/latest?amount=100&from=GBP&to=" + currencySymbol)
	noError(err)
	body, err := ioutil.ReadAll(resp.Body)
	noError(err)
	fmt.Println(string(body))
}

func main() {
	currencies := []string{"USD", "EUR", "ISK", "JPY", "NZD", "SEK", "CAD", "CHF", "HKD", "KRW", "PLN"}
	start := time.Now()
	for _, currency := range currencies {
		oneHundredPoundsIn(currency)
	}
	fmt.Println("That took", time.Since(start))
}
