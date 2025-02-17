package converter

import (
	"my-echo-app/internal/core/dto"
	"my-echo-app/internal/core/model"
)

// ConvertOrderToDTO mengonversi model.Order ke dto.OrderResponse
func ConvertOrderToDTO(order model.Order) dto.OrderResponse {
	return dto.OrderResponse{
		OrderID:     order.OrderID,
		ProductCode: order.ProductCode,
		ProductName: order.ProductName,
		Quantity:    order.Quantity,
		OrderDate:   order.OrderDate,
		Price:       order.Price,
		CustomerID:  order.CustomerID,
	}
}

// ConvertOrdersToDTO mengonversi slice dari model.Order ke slice dari dto.OrderResponse
func ConvertOrdersToDTO(orders []model.Order) []dto.OrderResponse {
	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ConvertOrderToDTO(order))
	}
	return orderResponses
}

// ConvertOrderRequestToModel mengonversi dto.OrderRequest ke model.Order
func ConvertOrderRequestToModel(order dto.OrderRequest) model.Order {
	return model.Order{
		OrderID:     order.OrderID,
		ProductCode: order.ProductCode,
		ProductName: order.ProductName,
		Quantity:    order.Quantity,
		OrderDate:   order.OrderDate,
		Price:       order.Price,
		CustomerID:  order.CustomerID,
	}
}
