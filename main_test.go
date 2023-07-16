package main

import (
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	rates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"GBP": 0.73,
	}

	amount := "100"
	sourceCurrency := "USD"
	targetCurrency := "EUR"

	convertedAmount := convertCurrency(amount, sourceCurrency, targetCurrency, rates)
	expectedAmount := 85.0

	if convertedAmount != expectedAmount {
		t.Errorf("Currency conversion incorrect. Expected: %.2f, Got: %.2f", expectedAmount, convertedAmount)
	}
}

func TestParseFloat(t *testing.T) {
	amount := "100.50"
	expectedFloat := 100.50

	parsedFloat := parseFloat(amount)

	if parsedFloat != expectedFloat {
		t.Errorf("Float parsing incorrect. Expected: %.2f, Got: %.2f", expectedFloat, parsedFloat)
	}
}

func TestFetchExchangeRates(t *testing.T) {
	rates, err := fetchExchangeRates()

	if err != nil {
		t.Errorf("Error fetching exchange rates: %v", err)
	}

	if len(rates) == 0 {
		t.Error("Exchange rates not fetched successfully")
	}
}
