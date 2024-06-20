package handlers

// TODO: Implemented authentication. This file is using the placeholder user.

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type createReviewInput struct {
	Rating       *uint8  `body:"rating" validate:"required,gte=0,lte=10" default:"10"`
	IsAnonymous  bool    `body:"isAnonymous" default:"false"`
	RestaurantId *int32  `body:"restaurantId" validate:"required"`
	Comments     *string `body:"comments"`
}

// CreateReview
//
//	@Summary	create a review
//
//	@Accept		json
//	@Produce	json
//	@Success	200			{object}	sqlcClient.Review
//	@Param		requestBody	body		createReviewInput	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/review/create [post]
func (handler Handler) CreateReview(context echo.Context) (err error) {
	input := createReviewInput{}

	if err = context.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = context.Validate(&input); err != nil {
		return err
	}

	lastInsertId, err := handler.Queries.CreateReview(
		context.Request().Context(),
		sqlcClient.CreateReviewParams{
			Rating:       *input.Rating,
			Comments:     input.Comments,
			RestaurantID: *input.RestaurantId,
			UserID:       "00000000-0000-0000-0000-000000000000",
			IsAnonymous:  util.ToPointer(util.BoolToInt[int8](input.IsAnonymous)),
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
