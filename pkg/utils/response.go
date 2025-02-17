package utils

import (
	"github.com/labstack/echo/v4"
)

// JSONResponseStruct digunakan untuk memastikan urutan JSON tetap
type JSONResponseStruct struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponseStruct digunakan untuk memastikan error JSON tetap urutannya
type ErrorResponseStruct struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// JSONResponse untuk mengembalikan response JSON yang terstruktur dengan urutan yang benar
func JSONResponse(c echo.Context, status int, message string, data interface{}) error {
	response := JSONResponseStruct{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.JSON(status, response)
}

// ErrorResponse mengembalikan response JSON untuk error dengan urutan yang benar
func ErrorResponse(c echo.Context, status int, message string) error {
	response := ErrorResponseStruct{
		Status: status,
		Error:  message,
	}
	return c.JSON(status, response)
}
