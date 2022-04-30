package models

type Customer struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Order struct {
	Item   string `json:"item"`
	Amount int64  `json:"amount"`
	Time   int64  `json:"time"`
}
