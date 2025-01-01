package models

type Order struct {
	Id              string   `json:"id"`
	UserId          string   `json:"userId"`
	BoxesId         []string `json:"boxesId"`
	Status          string   `json:"status"`
	BookingDateTime int64    `json:"bookingDateTime"`
	RentalPeriod    int      `json:"rentalPeriod"`
	TotalPrice      float64  `json:"totalPrice"`
}
