package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"

	"github.com/labstack/echo/v4"
)

type getRestaurantsQuery struct {
	Start   uint16 `query:"start"`
	Limit   uint8  `query:"limit" validate:"lte=50"`
	OrderBy string `query:"orderBy" validate:"omitempty,oneof=name created_at updated_at avg_rating"`
	Order   string `query:"order" validate:"omitempty,oneof=desc asc"`
}

// GetRestaurants
//
//	@Summary	get all restaurants
//
//	@Accept		json
//	@Produce	json
//	@Param		start	query		integer	false	"page start"				minimum(0)
//	@Param		limit	query		integer	false	"max entries"				maximum(50)
//	@Param		orderBy	query		string	false	"column to order by"		Enums(name, created_at, updated_at, avg_rating)
//	@Param		order	query		string	false	"ascending or descending"	Enums(asc, desc)
//	@Success	200		{array}		sqlcClient.GetRestaurantsRow
//	@Failure	400		{object}	echo.HTTPError
//	@Failure	500		{object}	echo.HTTPError
//	@Router		/restaurants [get]
func (handler Handler) GetRestaurants(context echo.Context) (err error) {
	query := getRestaurantsQuery{Start: 0, Limit: 20, OrderBy: "name"}

	if err = context.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = context.Validate(&query); err != nil {
		return err
	}

	data, err := handler.Queries.GetRestaurants(
		context.Request().Context(),
		sqlcClient.GetRestaurantsParams{
			Offset:  int32(query.Start),
			Limit:   int32(query.Limit),
			OrderBy: query.OrderBy,
			Order:   query.Order,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, data)
}
