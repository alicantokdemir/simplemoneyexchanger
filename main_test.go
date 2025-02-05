package main

import (
	"log"
	"testing"

	"github.com/alicantokdemir/simplemoneyexchanger/exchange"
)

func TestMain(t *testing.T) {

	exchanger := exchange.New()

	amount1 := exchange.Money{
		Amount:   10000,
		Currency: exchange.CurrencyUsd,
	}

	amount2 := exchange.Money{
		Amount:   12700,
		Currency: exchange.CurrencyJpy,
	}

	amount3 := exchange.Money{
		Amount:   320000,
		Currency: exchange.CurrencyTry,
	}

	convertedMoney, err := exchanger.Convert(amount1, exchange.CurrencyJpy)
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	log.Printf("Convert %s to %s\n", amount1.Print(), convertedMoney.Print())

	convertedMoney2, err := exchanger.Convert(amount2, exchange.CurrencyUsd)
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}
	log.Printf("Convert %s to %s\n", amount2.Print(), convertedMoney2.Print())

	convertedMoney3, err := exchanger.Convert(amount1, exchange.CurrencyTry)
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}
	log.Printf("Convert %s to %s\n", amount1.Print(), convertedMoney3.Print())

	convertedMoney4, err := exchanger.Convert(amount3, exchange.CurrencyUsd)
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}
	log.Printf("Convert %s to %s\n", amount3.Print(), convertedMoney4.Print())
}
