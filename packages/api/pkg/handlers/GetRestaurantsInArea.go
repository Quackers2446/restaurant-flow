package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Really we don't need to be this precise:
// "How much precision do you really need? 6 decimal places gives you enough precision to
// distinguish two people kissing each other. 8 can tell your fingers apart. FLOAT distinguishes
// two items 1.7m (5.6ft) apart. All of those are ludicrously excessive for "map" applications!"
// https://stackoverflow.com/questions/12504208/what-mysql-data-type-should-be-used-for-latitude-longitude-with-8-decimal-places
type getRestaurantsInAreaQuery struct {
	Lat float64 `query:"lat" validate:"required" example:"43.472587" default:"43.472587"`
	Lng float64 `query:"lng" validate:"required" example:"-80.537681" default:"-80.537681"`
	// Radius in meters
	Radius uint32 `query:"radius" validate:"required,lte=10000" example:"200" default:"200"`
}

// GetRestaurantsInArea
//
//	@Summary	get all restaurants in an area defined by latitude, longitude, and radius in meters
//	@Tags		Restaurants
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		sqlcClient.GetRestaurantsInAreaRow
//	@Param		requestQuery	query		getRestaurantsInAreaQuery	false	"request query"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/in-area [get]
func (handler Handler) GetRestaurantsInArea(context echo.Context) (err error) {
	query, err := util.ValidateInput[getRestaurantsInAreaQuery](&context, nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	restaurants, err := handler.Queries.GetRestaurantsInArea(
		context.Request().Context(),
		sqlcClient.GetRestaurantsInAreaParams{
			Lat:    query.Lat,
			Lng:    query.Lng,
			Radius: float64(query.Radius),
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, restaurants)
}
