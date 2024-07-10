package handlers

// TODO: COMBINE THIS WITH GET RESTAURANTS

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"
	"strings"

	"github.com/labstack/echo/v4"
)

type getRestaurantsSearchQuery struct {
	Search string `query:"search" validate:"required"`
}

type getRestaurantsSearchResult struct {
	*sqlcClient.GetRestaurantsSearchRow

	Tags         []*sqlcClient.Tag                        `json:"tags"`
	OpeningHours map[string]([]*sqlcClient.OpeningPeriod) `json:"openingHours"`
}

// GetRestaurantsSearch
//
//	@Summary	get many restaurants with fulltext search
//
//	@Tags		Restaurants
//	@Accept		json
//	@Produce	json
//	@Param		requestQuery	query		getRestaurantsSearchQuery	true	"request query"
//	@Success	200				{array}		getRestaurantsSearchResult
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/search [get]
func (handler Handler) GetRestaurantsSearch(context echo.Context) (err error) {
	query, err := util.ValidateInput(&context, &getRestaurantsSearchQuery{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	restaurants, err := handler.Queries.GetRestaurantsSearch(
		context.Request().Context(),
		query.Search,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Due to limitations of SQLC (and basically every SQL library for Go that does not involve a
	// query builder) we have to perform one-to-many left "joins" ourselves, which is a giant pain
	// in the ass.
	// Normally we would use a query builder like  https://github.com/go-jet/jet but then we
	// wouldn't be writing SQL, which (I assume) is against the rules of the project.
	restaurantIndex := make(map[int32]*getRestaurantsSearchResult)
	googleRestaurantIndex := make(map[int32]*getRestaurantsSearchResult)

	var data []*getRestaurantsSearchResult = util.Map(
		restaurants,
		func(restaurant *sqlcClient.GetRestaurantsSearchRow) *getRestaurantsSearchResult {
			resultItem := &getRestaurantsSearchResult{
				GetRestaurantsSearchRow: restaurant,
				Tags:                    []*sqlcClient.Tag{},
				OpeningHours:            make(map[string]([]*sqlcClient.OpeningPeriod)),
			}
			restaurantIndex[restaurant.RestaurantID] = resultItem              // SIDE EFFECT
			googleRestaurantIndex[*restaurant.GoogleRestaurantID] = resultItem // SIDE EFFECT

			return resultItem
		},
	)

	// Add tags
	tags, err := handler.Queries.GetTags(
		context.Request().Context(),
		util.Map(restaurants, func(restaurant *sqlcClient.GetRestaurantsSearchRow) int32 { return restaurant.RestaurantID }),
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
		util.Map(restaurants, func(restaurant *sqlcClient.GetRestaurantsSearchRow) int32 {
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
