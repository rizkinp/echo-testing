package dto

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

type OrderRequest struct {
	OrderID     string
	ProductCode string
	ProductName string
	Quantity    int
	OrderDate   time.Time
	Price       float64
	CustomerID  int
}

type OrderResponse struct {
	OrderID     string
	ProductCode string
	ProductName string
	Quantity    int
	OrderDate   time.Time
	Price       float64
	CustomerID  int
}
