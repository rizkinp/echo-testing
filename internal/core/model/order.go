package model

import "time"

type Order struct {
	OrderID     string
	ProductCode string
	ProductName string
	Quantity    int
	OrderDate   time.Time
	Price       float64
	CustomerID  int
}
