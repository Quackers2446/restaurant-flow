package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type dummyTable struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

// GetDummyTable godoc
//
//	@Tags		Dummy
//	@Summary	get dummy table
//	@Produce	json
//	@Success	200	{array}		dummyTable	"dummy table rows"
//	@Failure	500	{object}	echo.HTTPError
//	@Router		/dummy-table [get]
func (handler Handler) GetDummyTable(context echo.Context) (err error) {
	data := []dummyTable{}

	err = handler.DB.Select(&data, "SELECT * FROM `dummy_table`")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, data)
}
