package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler Handler) GetDummyTable(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}
