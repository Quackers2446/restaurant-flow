package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type leavePartyBody struct {
	PartyId int32 `body:"PartyId" validate:"required"`
}

// LeaveParty
//
//	@Summary	leave a party
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Param		requestBody		body		leavePartyBody	true	"request body"
//	@Param		Authorization	header		string			true	"access token"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	401				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/party/leave [delete]
func (handler Handler) LeaveParty(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	body, err := util.ValidateInput(&context, &leavePartyBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = handler.Queries.LeaveParty(
		context.Request().Context(),
		sqlcClient.LeavePartyParams{
			PartyID: body.PartyId,
			UserID:  claims.Subject,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.NoContent(http.StatusOK)
}
