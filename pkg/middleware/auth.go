package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware mengecek apakah request memiliki token autentikasi
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		return next(c)
	}
}
