package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"strings"

	"github.com/labstack/echo/v4"
)

type getRestaurantInput struct {
	RestaurantId uint16 `param:"id" validate:"required"`
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
//	@Accept		json
//	@Produce	json
//	@Param		requestParams	path		getRestaurantInput	true	"request params"
//	@Success	200				{object}	getRestaurantsResult
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	404				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/{id} [get]
func (handler Handler) GetRestaurant(context echo.Context) (err error) {
	input := getRestaurantInput{}

	if err = context.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = context.Validate(&input); err != nil {
		return err
	}

	restaurant, err := handler.Queries.GetRestaurant(
		context.Request().Context(),
		int32(input.RestaurantId),
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
		[]*int32{&data.RestaurantID},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	data.Tags = tags

	// Add opening hours
	openingHours, err := handler.Queries.GetOpeningHours(
		context.Request().Context(),
		[]*int32{&data.RestaurantID},
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
