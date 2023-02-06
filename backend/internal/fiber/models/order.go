package models

type Order struct {
	OrderID int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}
