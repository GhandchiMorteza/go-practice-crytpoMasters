package main

import (
	"fmt"
	"sync"

	"test.com/crypto/api"
)

func main() {
	currencies := [3]string {"BTC", "ETH", "BCH"}

	var wg sync.WaitGroup
	wg.Add(len(currencies))
	for _, v := range currencies {
		go func(v string) {
			defer wg.Done()
			getCurrencyData(v)
		}(v)
	}
	wg.Wait()
}

func getCurrencyData(currencyName string) {
	rate, err := api.GetRate(currencyName)
	
	if err != nil {
		fmt.Printf("error happens %v\n", err)
	} else {
		fmt.Printf("The rate for %s is %.2f\n", rate.Currency, rate.Price)
	}
}