package exchange

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
)

type Currency struct {
	Code           string
	Symbol         string
	SymbolPosition string
	Divisor        int
}

type Money struct {
	Amount   int64
	Currency Currency
}

func (m *Money) Print() string {
	currencyFmt := fmt.Sprintf("%.*f", int64(math.Log10(float64(m.Currency.Divisor))), float64(m.Amount)/float64(m.Currency.Divisor))

	if m.Currency.SymbolPosition == "before" {
		return m.Currency.Symbol + currencyFmt
	} else {
		return currencyFmt + m.Currency.Symbol
	}
}

type ExchangeRate map[string]map[string]Money

var exchangeRates ExchangeRate
var CurrencyUsd Currency
var CurrencyJpy Currency
var CurrencyTry Currency
var CurrencyEur Currency
var CurrencyGbp Currency

func init() {
	CurrencyUsd = Currency{
		Code:           "USD",
		Symbol:         "$",
		SymbolPosition: "before",
		Divisor:        10000,
	}

	CurrencyJpy = Currency{
		Code:           "JPY",
		Symbol:         "¥",
		SymbolPosition: "after",
		Divisor:        100,
	}

	CurrencyTry = Currency{
		Code:           "TRY",
		Symbol:         "₺",
		SymbolPosition: "before",
		Divisor:        10000,
	}

	CurrencyEur = Currency{
		Code:           "EUR",
		Symbol:         "€",
		SymbolPosition: "before",
		Divisor:        10000,
	}

	exchangeRates = map[string]map[string]Money{
		"USD": {
			"JPY": Money{
				Amount:   15300,
				Currency: CurrencyJpy,
			},
			"TRY": Money{
				Amount:   360000,
				Currency: CurrencyTry,
			},
			"EUR": Money{
				Amount:   9600,
				Currency: CurrencyEur,
			},
		},
		"JPY": {
			"USD": Money{
				Amount:   65,
				Currency: CurrencyUsd,
			},
			"TRY": Money{
				Amount:   2353,
				Currency: CurrencyTry,
			},
			"EUR": Money{
				Amount:   6,
				Currency: CurrencyEur,
			},
		},
		"TRY": {
			"USD": Money{
				Amount:   300,
				Currency: CurrencyUsd,
			},
			"JPY": Money{
				Amount:   425,
				Currency: CurrencyJpy,
			},
			"EUR": Money{
				Amount:   285,
				Currency: CurrencyEur,
			},
		},
		"EUR": {
			"USD": Money{
				Amount:   10417,
				Currency: CurrencyUsd,
			},
			"JPY": Money{
				Amount:   12500,
				Currency: CurrencyJpy,
			},
			"TRY": Money{
				Amount:   90000,
				Currency: CurrencyTry,
			},
		},
	}
}

type MoneyExchanger interface {
	GetExchangeRate(fromCurrencyCode string, toCurrencyCode string) (Money, error)
	Convert(fromMoney Money, toCurrency Currency) (Money, error)
}

func New() MoneyExchanger {
	return &SimpleMoneyExchanger{
		exchangeRates: exchangeRates,
	}
}

type SimpleMoneyExchanger struct {
	exchangeRates ExchangeRate
}

func (s *SimpleMoneyExchanger) LoadExchangeRates(rates ExchangeRate) {
	s.exchangeRates = rates
}

func (s *SimpleMoneyExchanger) GetExchangeRate(fromCurrencyCode string, toCurrencyCode string) (Money, error) {
	exchangeRate, ok := s.exchangeRates[fromCurrencyCode][toCurrencyCode]

	if !ok {
		return Money{}, errors.New("Exchange rate not found")
	}

	return exchangeRate, nil
}

func (s *SimpleMoneyExchanger) Convert(fromMoney Money, toCurrency Currency) (Money, error) {
	exchangeRate, err := s.GetExchangeRate(fromMoney.Currency.Code, toCurrency.Code)

	if err != nil {
		return Money{}, err
	}

	fromMoneyAmount := float64(fromMoney.Amount) / float64(fromMoney.Currency.Divisor)
	exchangeRateAmount := float64(exchangeRate.Amount) / float64(exchangeRate.Currency.Divisor)
	convertedAmount := (fromMoneyAmount * exchangeRateAmount) * float64(toCurrency.Divisor)

	return Money{
		Amount:   int64(convertedAmount),
		Currency: toCurrency,
	}, nil
}
