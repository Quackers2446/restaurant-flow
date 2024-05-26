package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DummyTable struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func (handler Handler) GetDummyTable(context echo.Context) error {
	rows, err := handler.DB.Query("SELECT * FROM `dummyTable`")
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	var data []DummyTable

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var row DummyTable

		if err := rows.Scan(&row.Id, &row.Description); err != nil {
			return err
		}

		data = append(data, row)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, data)
}
