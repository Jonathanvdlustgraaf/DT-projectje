package types

import "time"

type Order struct {
	No         int         `json:"no"`
	Date       time.Time   `json:"date"`
	DateString string      `json:"datestring"`
	Email      string      `json:"email"`
	Lines      []OrderLine `json:"lines"`
	Total      float32     `json:"total"`
}

type OrderLine struct {
	CustomPizza
	Pizza       string `json:"pizza"`
	Qty         int    `json:"qty"`
}

type CustomerWithOrders struct {
	Customer
	Orders []Order
}
