package routes

import (
	"my-echo-app/internal/transport/http/handlers"
	"my-echo-app/pkg/middleware"

	"github.com/labstack/echo/v4"
)

// RegisterOrderRoutes menghubungkan route Order dengan handler

func RegisterOrderRoutes(api *echo.Group, handler *handlers.OrderHandler) {
	orderGroup := api.Group("/orders")
	orderGroup.GET("", handler.GetAllOrders)
	orderGroup.GET("/:order_id", handler.GetOrder, middleware.AuthMiddleware)
	orderGroup.POST("", handler.CreateOrder)
}
