package models

type Customer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Order struct {
	Item   string `json:"item"`
	Amount int64  `json:"amount"`
	Time   string `json:"time"`
}
