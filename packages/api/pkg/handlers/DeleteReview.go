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
type deleteReviewBody struct {
	RestaurantId *int32 `body:"restaurantId" validate:"required"`
}

// DeleteReview
//
//	@Summary	delete a review
//
//	@Tags		Reviews
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Param		requestBody	body		deleteReviewBody	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/review/delete [delete]

func (handler Handler) DeleteReview(context echo.Context) (err error) {
	body, err := util.ValidateInput(&context, &deleteReviewBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = handler.Queries.DeleteReview(
		context.Request().Context(),
		sqlcClient.DeleteReviewParams{
			RestaurantID: *body.RestaurantId,
			UserID:       "00000000-0000-0000-0000-000000000000",
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.NoContent(http.StatusOK)
}
