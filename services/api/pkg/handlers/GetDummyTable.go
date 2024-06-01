package handlers

import (
	"net/http"
	"restaurant-flow/pkg/httputil"

	"github.com/labstack/echo/v4"
)

type dummyTable struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

// GetDummyTable godoc
//
//	@Summary	get dummy table
//	@Produce	json
//	@Success	200	{array}		dummyTable	"dummy table rows"
//	@Failure	500	{object}	httputil.HTTPError
//	@Router		/dummy-table [get]
func (handler Handler) GetDummyTable(context echo.Context) error {
	data := []dummyTable{}

	err := handler.DB.Select(&data, "SELECT * FROM `dummyTable`")

	if err != nil {
		return httputil.NewError(context, http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, data)
}
