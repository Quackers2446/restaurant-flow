package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type updateReviewBody struct {
	Rating       *uint8  `body:"rating" validate:"required,gte=0,lte=10"`
	IsAnonymous  bool    `body:"isAnonymous" default:"false"`
	RestaurantId *int32  `body:"restaurantId" validate:"required"`
	Comments     *string `body:"comments"`
}

// UpdateReview
//
//	@Summary	update a review
//
//	@Tags		Reviews
//	@Accept		json
//	@Produce	json
//	@Success	200			{object}	sqlcClient.Review
//	@Param		requestBody	body		updateReviewBody	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/review/update [post]

func (handler Handler) UpdateReview(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	body, err := util.ValidateInput(&context, &updateReviewBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err2 := handler.Queries.UpdateReview(
		context.Request().Context(),
		sqlcClient.UpdateReviewParams{
			Rating:       *body.Rating,
			Comments:     body.Comments,
			RestaurantID: *body.RestaurantId,
			UserID:       claims.Subject,
			IsAnonymous:  util.ToPointer(util.BoolToInt[int8](body.IsAnonymous)),
		},
	)

	if err2 != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err2)
	}

	review, err := handler.Queries.GetUpdatedReview(
		context.Request().Context(),
		sqlcClient.GetUpdatedReviewParams{
			UserID:       claims.Subject,
			RestaurantID: *body.RestaurantId,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, review)
}
