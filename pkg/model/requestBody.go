package model

type RequestBody struct {
	Amount float64 `json:"amount"`
	From string `json:"from"`
	To string `json:"to"`
}
