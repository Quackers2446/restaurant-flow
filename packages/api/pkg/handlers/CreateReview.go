package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type createReviewBody struct {
	Rating       *uint8  `body:"rating" validate:"required,gte=0,lte=10"`
	IsAnonymous  bool    `body:"isAnonymous" default:"false"`
	RestaurantId *int32  `body:"restaurantId" validate:"required"`
	Comments     *string `body:"comments"`
}

// CreateReview
//
//	@Summary	create a review
//
//	@Tags		Reviews
//	@Accept		json
//	@Produce	json
//	@Success	200			{object}	sqlcClient.Review
//	@Param		requestBody	body		createReviewBody	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	401			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/review/create [post]
func (handler Handler) CreateReview(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	body, err := util.ValidateInput(&context, &createReviewBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	lastInsertId, err := handler.Queries.CreateReview(
		context.Request().Context(),
		sqlcClient.CreateReviewParams{
			Rating:       *body.Rating,
			Comments:     body.Comments,
			RestaurantID: *body.RestaurantId,
			UserID:       claims.Subject,
			IsAnonymous:  util.ToPointer(util.BoolToInt[int8](body.IsAnonymous)),
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	review, err := handler.Queries.GetReview(
		context.Request().Context(),
		int32(lastInsertId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, review)
}
