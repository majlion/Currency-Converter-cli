package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const baseURL = "https://api.exchangerate-api.com/v4/latest/USD"

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	// Fetching exchange rates
	rates, err := fetchExchangeRates()
	if err != nil {
		fmt.Println("Error fetching exchange rates:", err)
		return
	}

	// Getting the info from command-line arguments
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: currencyconverter <amount> <source_currency> <target_currency>")
		return
	}

	amount := args[0]
	sourceCurrency := strings.ToUpper(args[1])
	targetCurrency := strings.ToUpper(args[2])

	// Converting the currency
	convertedAmount := convertCurrency(amount, sourceCurrency, targetCurrency, rates)
	if convertedAmount == -1 {
		fmt.Println("Error converting currency. Invalid currency code.")
		return
	}

	// Printing the conversion result
	fmt.Printf("%s %s = %.2f %s\n", amount, sourceCurrency, convertedAmount, targetCurrency)
}

func fetchExchangeRates() (map[string]float64, error) {
	response, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRates ExchangeRates
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return nil, err
	}

	return exchangeRates.Rates, nil
}

func convertCurrency(amount, sourceCurrency, targetCurrency string, rates map[string]float64) float64 {
	if sourceCurrency == targetCurrency {
		return 0
	}

	sourceRate, ok := rates[sourceCurrency]
	if !ok {
		return -1
	}

	targetRate, ok := rates[targetCurrency]
	if !ok {
		return -1
	}

	amountFloat := parseFloat(amount)
	if amountFloat == -1 {
		return -1
	}

	return amountFloat * (targetRate / sourceRate)
}

func parseFloat(amount string) float64 {
	value, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return -1
	}
	return value
}
