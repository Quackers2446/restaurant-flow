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
type joinPartyParams struct {
	PartyId *int32 `param:"partyId" validate:"required"`
}

// Join party
//
//	@Summary	join a party
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200			{object}	sqlcClient.Review
//	@Param		requestBody	body		joinPartyParams	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	401			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/party/join/{partyId} [post]
func (handler Handler) JoinParty(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	params, err := util.ValidateInput(&context, &joinPartyParams{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	party, err := handler.Queries.GetPartyDetails(
		context.Request().Context(),
		*params.PartyId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	partySize, err := handler.Queries.GetPartySize(
		context.Request().Context(),
		*params.PartyId,
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
			PartyID: *params.PartyId,
			UserID:  claims.Subject,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// return context.NoContent(http.StatusOK)
	return context.JSON(http.StatusOK, joined)
}
