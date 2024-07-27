package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"
	"time"

	"github.com/labstack/echo/v4"
)

// Note: pointers allow validator to tell the difference between 0 and empty
// https://stackoverflow.com/questions/66632787/how-to-allow-zero0-value
type createPartyBody struct {
	MaxMembers   *uint8  `body:"max_members" validate:"gte=2"`
	RestaurantId int32   `body:"restaurantId" validate:"required"`
	Description  *string `body:"description"`
	Time         *int64  `body:"time"`
}

// CreateParty
//
//	@Summary	create a party
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200				{object}	sqlcClient.Party
//	@Param		requestBody		body		createPartyBody	true	"request body"
//	@Param		Authorization	header		string			true	"access token"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	401				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/party/create [post]
func (handler Handler) CreateParty(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	body, err := util.ValidateInput(&context, &createPartyBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	lastInsertId, err := handler.Queries.CreateParty(
		context.Request().Context(),
		sqlcClient.CreatePartyParams{
			MaxMembers:   int32(*body.MaxMembers),
			Description:  body.Description,
			RestaurantID: body.RestaurantId,
			Time:         time.Unix(*body.Time, 0),
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	party, err := handler.Queries.JoinParty(
		context.Request().Context(),
		sqlcClient.JoinPartyParams{
			UserID:  claims.Subject,
			PartyID: int32(lastInsertId),
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, party)
}
