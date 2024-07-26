package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"
	"strings"

	"github.com/labstack/echo/v4"
)

type getRestaurantParams struct {
	RestaurantId uint16 `param:"id" json:"id" validate:"required"`
}

type getRestaurantResult struct {
	*sqlcClient.GetRestaurantRow

	Tags         []*sqlcClient.Tag                        `json:"tags"`
	OpeningHours map[string]([]*sqlcClient.OpeningPeriod) `json:"openingHours"`
}

// GetRestaurant
//
//	@Summary	get a restaurant
//
//	@Tags		Restaurants
//	@Accept		json
//	@Produce	json
//	@Param		requestParams	path		getRestaurantParams	true	"request params"
//	@Success	200				{object}	getRestaurantsResult
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	401				{object}	echo.HTTPError
//	@Failure	404				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/{id} [get]
func (handler Handler) GetRestaurant(context echo.Context) (err error) {
	params, err := util.ValidateInput(&context, &getRestaurantParams{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	restaurant, err := handler.Queries.GetRestaurant(
		context.Request().Context(),
		int32(params.RestaurantId),
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var data = &getRestaurantResult{
		GetRestaurantRow: restaurant,
		Tags:             []*sqlcClient.Tag{},
		OpeningHours:     make(map[string]([]*sqlcClient.OpeningPeriod)),
	}

	// Add tags
	tags, err := handler.Queries.GetTags(
		context.Request().Context(),
		[]int32{data.RestaurantID},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	data.Tags = tags

	// Add opening hours
	openingHours, err := handler.Queries.GetOpeningHours(
		context.Request().Context(),
		[]int32{data.RestaurantID},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, hour := range openingHours {
		typeKey := strings.ToLower(string(hour.Type))
		openingHoursForRestaurant := data.OpeningHours

		val, ok := openingHoursForRestaurant[typeKey]

		if ok {
			openingHoursForRestaurant[typeKey] = append(val, &hour.OpeningPeriod)
		} else {
			openingHoursForRestaurant[typeKey] = []*sqlcClient.OpeningPeriod{&hour.OpeningPeriod}
		}

	}

	return context.JSON(http.StatusOK, data)
}
