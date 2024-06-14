package handlers

import (
	"net/http"
	"restaurant-flow/pkg/httputil"

	"github.com/labstack/echo/v4"
)

// GetRestaurants
//
//	@Summary	get all restaurants
//	@Produce	json
//	@Success	200	{array}		sqlcClient.GetRestaurantsRow
//	@Failure	500	{object}	httputil.HTTPError
//	@Router		/restaurants [get]
func (handler Handler) GetRestaurants(context echo.Context) error {
	data, err := handler.Queries.GetRestaurants(context.Request().Context())

	if err != nil {
		return httputil.NewError(context, http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, data)
}
