package httputil

import "github.com/labstack/echo/v4"

// NewError example
func NewError(context echo.Context, statusCode int, err error) error {
	return context.JSON(statusCode, HTTPError{
		Code:    statusCode,
		Message: err.Error(),
	})
}

type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"example error"`
}
