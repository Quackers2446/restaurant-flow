package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type getRestaurantReviewsQuery struct {
	// Page start
	Start uint16 `query:"start" validate:"omitempty,gt=0"`
	// Max entries
	Limit uint8 `query:"limit" validate:"omitempty,lte=50,gt=0"`
	// Column to order by
	OrderBy string `query:"orderBy" validate:"omitempty,oneof=created_at rating"`
	// Ascending or descending
	Order string `query:"order" validate:"omitempty,oneof=desc asc"`

	RestaurantId uint16 `param:"restaurantId" json:"restaurantId" validate:"required"`
}

// GetRestaurantReviews
//
//	@Summary	get reviews for a restaurant
//
//	@Tags		Restaurants, Reviews
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		sqlcClient.Review
//	@Param		requestQuery	query		getRestaurantReviewsQuery	false	"request query"
//	@Param		requestParams	path		getRestaurantReviewsParams	true	"request params"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/{restaurantId}/reviews [get]
func (handler Handler) GetRestaurantReviews(context echo.Context) (err error) {
	query, err := util.ValidateInput(&context, &getRestaurantReviewsQuery{Start: 0, Limit: 20, OrderBy: "created_at"})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	reviews, err := handler.Queries.GetRestaurantReviews(
		context.Request().Context(),
		sqlcClient.GetRestaurantReviewsParams{
			RestaurantID: int32(query.RestaurantId),
			Offset:       int32(query.Start),
			Limit:        int32(query.Limit),
			OrderBy:      query.OrderBy,
			Order:        query.Order,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, reviews)
}
