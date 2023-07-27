package models

type Item struct {
	Id         string `json:"id,omitempty"`
	CustomerId string `json:"customer_id"`
	Title      string `json:"title"`
	Price      int64  `json:"price"`
}
