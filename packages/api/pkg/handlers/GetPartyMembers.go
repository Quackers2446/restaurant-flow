package handlers

import (
	"net/http"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type getPartyMembersParams struct {
	PartyId int32 `param:"partyId" json:"partyId" validate:"required"`
}

// GetPartyMembers
//
//	@Summary	get members for a party
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		sqlcClient.PartyMember
//	@Param		requestParams	path		getPartyMembersParams	true	"request params"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/party/members/{partyId} [get]
func (handler Handler) GetPartyMembers(context echo.Context) (err error) {
	params, err := util.ValidateInput(&context, &getPartyMembersParams{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	members, err := handler.Queries.GetPartyMembers(
		context.Request().Context(),
		[]int32{params.PartyId},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, members)
}
