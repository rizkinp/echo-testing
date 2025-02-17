package app

import (
	"log"
	"my-echo-app/internal/core/converter"
	"my-echo-app/internal/core/dto"
	"my-echo-app/internal/interfaces/repository"
	"my-echo-app/internal/interfaces/usecase"
)

// NewOrderService membuat instance dari OrderService
func NewOrderService(repo repository.OrderRepository) usecase.OrderUseCase {
	return &orderService{repo: repo}
}

type orderService struct {
	repo repository.OrderRepository
}

// CreateOrder implements usecase.OrderUseCase.
func (s *orderService) CreateOrder(order dto.OrderRequest) (dto.OrderResponse, error) {
	convertedOrder := converter.ConvertOrderRequestToModel(order)
	newOrder, err := s.repo.CreateOrder(convertedOrder)
	if err != nil {
		log.Printf("❌ Error Creating Order: %v", err) // Tambahkan log ini
		return dto.OrderResponse{}, err
	}
	return converter.ConvertOrderToDTO(newOrder), nil
}

// GetAllOrders mengambil semua order dan mengonversi ke DTO menggunakan `order_converter.go`
func (s *orderService) GetAllOrders() ([]dto.OrderResponse, error) {
	orders, err := s.repo.GetAllOrders()
	if err != nil {
		log.Printf("❌ Error Getting All Orders: %v", err)
		return nil, err
	}
	return converter.ConvertOrdersToDTO(orders), nil
}

// GetOrder mengambil satu order berdasarkan ID dan mengonversi ke DTO menggunakan `order_converter.go`
func (s *orderService) GetOrder(id string) (dto.OrderResponse, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		log.Printf("❌ Error Getting Order: %v", err)
		return dto.OrderResponse{}, err
	}
	return converter.ConvertOrderToDTO(order), nil
}
