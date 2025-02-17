package repository

import (
	"my-echo-app/internal/core/model"
)

type OrderRepository interface {
	GetAllOrders() ([]model.Order, error)
	GetOrder(id string) (model.Order, error)
	CreateOrder(order model.Order) (model.Order, error)
}
