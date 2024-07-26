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
type createTagBody struct {
	Name		 *string  `body:"name" validate:"required"`
	RestaurantId *int32  `body:"restaurant_id" validate:"required"`
}

// CreateTag
//
//	@Summary	create a tag
//
//	@Tags		Tag
//	@Accept		json
//	@Produce	json
//	@Success	200			{object}	sqlcClient.Tag
//	@Param		requestBody	body		createTagBody	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/tag/create [post]
func (handler Handler) CreateTag(context echo.Context) (err error) {
	body, err := util.ValidateInput(&context, &createTagBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	lastInsertId, err := handler.Queries.CreateTag(
		context.Request().Context(),
		sqlcClient.CreateTagParams{
			RestaurantID: *body.RestaurantId,
			Name:		  body.Name,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	review, err := handler.Queries.GetTag(
		context.Request().Context(),
		int32(lastInsertId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, review)
}
