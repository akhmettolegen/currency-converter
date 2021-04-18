package model

type RequestBody struct {
	Amount float64 `json:"amount"`
	Currency Currency `json:"currency"`
}

type Currency string
const(
	CurrencyUSD Currency = "USD"
	CurrencyEUR  = "EUR"
	CurrencyRUB = "RUB"
	CurrencyKZT = "KZT"
)