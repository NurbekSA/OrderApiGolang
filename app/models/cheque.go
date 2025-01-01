package models

type Cheque struct {
	Id      string `json:"Id"`
	orderId string `json:"orderId"`
	body    string `json:"body"`
}
