# How to Use:

- Initialize a Currency

```
Currency{
    Code:           "USD",
    Symbol:         "$",
    SymbolPosition: "before",
    Divisor:        10000,
}
```

- Define exchange rates

```
exchangeRates = map[string]map[string]Money{
    "USD": {
        "TRY": Money{
            Amount:   360000,
            Currency: CurrencyTry,
        },
        "EUR": Money{
            Amount:   9600,
            Currency: CurrencyEur,
        },
    }
}
```

- Convert

```
amount := exchange.Money{
    Amount:   10000,
    Currency: exchange.CurrencyUsd,
}

convertedMoney, err := exchanger.Convert(amount1, exchange.CurrencyTry)
```

- Print Money in human readable format

```
amount := exchange.Money{
    Amount:   10000,
    Currency: exchange.CurrencyUsd,
}
amount.Print() // $1.0000

amount2 := exchange.Money{
    Amount:   15300,
    Currency: exchange.CurrencyJpy,
}
amount2.Print() // 153.00¥
```

## Test
```go test```

### Idea behind:
https://martinfowler.com/eaaCatalog/money.html