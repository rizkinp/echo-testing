package routes

import (
	"my-echo-app/internal/transport/http/handlers"
	apiRoutes "my-echo-app/internal/transport/http/routes/api"

	"github.com/labstack/echo/v4"
)

// RoutesConfig
type RoutesConfig struct {
	OrderHandler *handlers.OrderHandler
}

// SetupRoutes /api
func SetupRoutes(e *echo.Echo, cfg *RoutesConfig) {
	api := e.Group("/api")

	apiRoutes.RegisterOrderRoutes(api, cfg.OrderHandler)
}
