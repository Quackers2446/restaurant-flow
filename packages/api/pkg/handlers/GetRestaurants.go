package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"
	"strings"

	"github.com/labstack/echo/v4"
)

type getRestaurantsQuery struct {
	// Page start
	Start uint16 `query:"start" validate:"omitempty,gt=0"`
	// Max entries
	Limit uint8 `query:"limit" validate:"omitempty,lte=50,gt=0"`
	// Column to order by
	OrderBy string `query:"orderBy" validate:"omitempty,oneof=name created_at updated_at avg_rating"`
	// Ascending or descending
	Order string `query:"order" validate:"omitempty,oneof=desc asc"`
}

type getRestaurantsResult struct {
	*sqlcClient.GetRestaurantsRow

	Tags         []*sqlcClient.Tag                        `json:"tags"`
	OpeningHours map[string]([]*sqlcClient.OpeningPeriod) `json:"openingHours"`
}

// GetRestaurants
//
//	@Summary	get many restaurants
//
//	@Tags		Restaurants
//	@Accept		json
//	@Produce	json
//	@Param		requestQuery	query		getRestaurantsQuery	false	"request query"
//	@Success	200				{array}		getRestaurantsResult
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants [get]
func (handler Handler) GetRestaurants(context echo.Context) (err error) {
	query, err := util.ValidateInput(&context, &getRestaurantsQuery{Start: 0, Limit: 20, OrderBy: "name"})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	restaurants, err := handler.Queries.GetRestaurants(
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

	// Due to limitations of SQLC (and basically every SQL library for Go that does not involve a
	// query builder) we have to perform one-to-many left "joins" ourselves, which is a giant pain
	// in the ass.
	// Normally we would use a query builder like  https://github.com/go-jet/jet but then we
	// wouldn't be writing SQL, which (I assume) is against the rules of the project.
	restaurantIndex := make(map[int32]*getRestaurantsResult)
	googleRestaurantIndex := make(map[int32]*getRestaurantsResult)

	var data []*getRestaurantsResult = util.Map(
		restaurants,
		func(restaurant *sqlcClient.GetRestaurantsRow) *getRestaurantsResult {
			resultItem := &getRestaurantsResult{
				GetRestaurantsRow: restaurant,
				Tags:              []*sqlcClient.Tag{},
				OpeningHours:      make(map[string]([]*sqlcClient.OpeningPeriod)),
			}
			restaurantIndex[restaurant.RestaurantID] = resultItem              // SIDE EFFECT
			googleRestaurantIndex[*restaurant.GoogleRestaurantID] = resultItem // SIDE EFFECT

			return resultItem
		},
	)

	// Add tags
	tags, err := handler.Queries.GetTags(
		context.Request().Context(),
		util.Map(restaurants, func(restaurant *sqlcClient.GetRestaurantsRow) int32 { return restaurant.RestaurantID }),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, tag := range tags {
		restaurantIndex[tag.RestaurantID].Tags = append(restaurantIndex[tag.RestaurantID].Tags, tag)
	}

	// Add opening hours
	openingHours, err := handler.Queries.GetOpeningHours(
		context.Request().Context(),
		util.Map(restaurants, func(restaurant *sqlcClient.GetRestaurantsRow) int32 {
			return restaurant.GoogleRestaurant.GoogleRestaurantID
		}),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, hour := range openingHours {
		typeKey := strings.ToLower(string(hour.Type))
		openingHoursForRestaurant := googleRestaurantIndex[hour.GoogleRestaurantID].OpeningHours

		val, ok := openingHoursForRestaurant[typeKey]

		if ok {
			openingHoursForRestaurant[typeKey] = append(val, &hour.OpeningPeriod)
		} else {
			openingHoursForRestaurant[typeKey] = []*sqlcClient.OpeningPeriod{&hour.OpeningPeriod}
		}

	}

	return context.JSON(http.StatusOK, data)
}
