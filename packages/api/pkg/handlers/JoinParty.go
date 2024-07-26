package handlers

// TODO: Implemented authentication. This file is using the placeholder user.

import (
	"fmt"
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type joinPartyQuery struct {
	PartyId *int32 `param:"partyId" validate:"required"`
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
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/review/create [post]
func (handler Handler) JoinParty(context echo.Context) (err error) {
	body, err := util.ValidateInput(&context, &joinPartyQuery{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	party, err := handler.Queries.GetPartyDetails(
		context.Request().Context(),
		*body.PartyId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	partySize, err := handler.Queries.GetPartySize(
		context.Request().Context(),
		*body.PartyId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if party.MaxMembers < int32(partySize)+1 {
		fmt.Println(partySize)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	joined, err := handler.Queries.JoinParty(
		context.Request().Context(),
		sqlcClient.JoinPartyParams{
			PartyID: *body.PartyId,
			UserID:  "00000000-0000-0000-0000-000000000002",
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// return context.NoContent(http.StatusOK)
	return context.JSON(http.StatusOK, joined)
}
