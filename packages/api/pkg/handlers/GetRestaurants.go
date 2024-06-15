package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type getRestaurantsQuery struct {
	Start   uint16 `query:"start"`
	Limit   uint8  `query:"limit" validate:"lte=50"`
	OrderBy string `query:"orderBy" validate:"omitempty,oneof=name created_at updated_at avg_rating"`
	Order   string `query:"order" validate:"omitempty,oneof=desc asc"`
}

type getRestaurantsResult struct {
	sqlcClient.GetRestaurantsRow

	Tags []*sqlcClient.Tag `json:"tags"`
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
//	@Success	200		{array}		getRestaurantsResult
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
	index := make(map[int32]*getRestaurantsResult)

	var data []*getRestaurantsResult = util.Map(
		restaurants,
		func(restaurant *sqlcClient.GetRestaurantsRow) *getRestaurantsResult {
			resultItem := &getRestaurantsResult{GetRestaurantsRow: *restaurant, Tags: []*sqlcClient.Tag{}}
			index[restaurant.RestaurantID] = resultItem // SIDE EFFECT

			return resultItem
		},
	)

	tags, err := handler.Queries.GetTags(
		context.Request().Context(),
		util.Map(restaurants, func(restaurant *sqlcClient.GetRestaurantsRow) *int32 { return &restaurant.RestaurantID }),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, tag := range tags {
		index[*tag.RestaurantID].Tags = append(index[*tag.RestaurantID].Tags, &tag)
	}

	return context.JSON(http.StatusOK, data)
}
