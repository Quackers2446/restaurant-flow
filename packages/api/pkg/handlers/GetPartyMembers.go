package handlers

import (
	"net/http"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type getPartyMembersQuery struct {
	PartyId uint16 `param:"partyId" json:"partyId" validate:"required"`
}

// GetPartyMembers
//
//	@Summary	get members for a party
//
//	@Tags		Restaurants, Reviews
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		sqlcClient.PartyMembers
//	@Param		requestQuery	query		getRestaurantReviewsQuery	false	"request query"
//	@Param		requestParams	path		getRestaurantReviewsParams	true	"request params"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/restaurants/{restaurantId}/reviews [get]
func (handler Handler) GetPartyMembers(context echo.Context) (err error) {
	query, err := util.ValidateInput(&context, &getPartyMembersQuery{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	reviews, err := handler.Queries.GetPartyMembers(
		context.Request().Context(),
		int32(query.PartyId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, reviews)
}
