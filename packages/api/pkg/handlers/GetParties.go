package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetParties
//
//	@Summary	get parties
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		sqlcClient.Party
//	@Failure	400	{object}	echo.HTTPError
//	@Failure	500	{object}	echo.HTTPError
//	@Router		/party/all [get]
func (handler Handler) GetParties(context echo.Context) (err error) {
	parties, err := handler.Queries.GetParties(
		context.Request().Context(),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, parties)
}
