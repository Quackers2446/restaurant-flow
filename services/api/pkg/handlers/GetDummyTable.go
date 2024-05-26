package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DummyTable struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

func (handler Handler) GetDummyTable(context echo.Context) error {
	data := []DummyTable{}

	err := handler.DB.Select(&data, "SELECT * FROM `dummyTable`")

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, data)
}
