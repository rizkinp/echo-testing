package usecase

import (
	"my-echo-app/internal/core/dto"
)

type OrderUseCase interface {
	GetAllOrders() ([]dto.OrderResponse, error)
	GetOrder(id string) (dto.OrderResponse, error)
	CreateOrder(order dto.OrderRequest) (dto.OrderResponse, error)
}
