package handlers

import (
	"my-echo-app/internal/core/dto"
	"my-echo-app/internal/interfaces/usecase"
	"my-echo-app/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Struct
type OrderHandler struct {
	orderUseCase usecase.OrderUseCase
}

// Cnnstructor
func NewOrderHandler(orderUseCase usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

// @Summary Get all orders
// @Description Get all orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.JSONResponseStruct
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(c echo.Context) error {
	orders, err := h.orderUseCase.GetAllOrders()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders")
	}
	return utils.JSONResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// @Summary Get order by ID
// @Description Get order by ID
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} utils.JSONResponseStruct
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c echo.Context) error {
	orderID := c.Param("id")
	order, err := h.orderUseCase.GetOrder(orderID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Order not found")
	}
	return utils.JSONResponse(c, http.StatusOK, "Order retrieved successfully", order)
}

// @Summary Create order
// @Description Create order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param order body dto.OrderRequest true "Order details"
// @Success 200 {object} utils.JSONResponseStruct
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var order dto.OrderRequest
	if err := c.Bind(&order); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}
	newOrder, err := h.orderUseCase.CreateOrder(order)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create order")
	}
	return utils.JSONResponse(c, http.StatusCreated, "Order created successfully", newOrder)
}
