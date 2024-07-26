package handlers

import (
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type getPartiesQuery struct {
	// Page start
	Start uint16 `query:"start" validate:"omitempty,gt=0"`
	// Max entries
	Limit uint8 `query:"limit" validate:"omitempty,lte=50,gt=0"`
	// Column to order by
	OrderBy string `query:"orderBy" validate:"omitempty,oneof=time"`
	// Ascending or descending
	Order string `query:"order" validate:"omitempty,oneof=desc asc"`
}

type getPartiesResult struct {
	*sqlcClient.Party

	PartyMembers []*sqlcClient.GetPartyMembersRow `json:"partyMembers"`
}

// GetParties
//
//	@Summary	get parties
//
//	@Tags		Party
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		getPartiesResult
//	@Param		requestParams	path		getPartiesQuery	true	"request params"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/party/all [get]
func (handler Handler) GetParties(context echo.Context) (err error) {
	query, err := util.ValidateInput(&context, &getPartiesQuery{Start: 0, Limit: 20, OrderBy: "time"})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	parties, err := handler.Queries.GetParties(
		context.Request().Context(),
		sqlcClient.GetPartiesParams{
			Offset:  int32(query.Start),
			Limit:   int32(query.Limit),
			OrderBy: query.OrderBy,
			Order:   query.Order,
		},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	partyIndex := make(map[int32]*getPartiesResult)

	var data []*getPartiesResult = util.Map(
		parties,
		func(party *sqlcClient.Party) *getPartiesResult {
			resultItem := &getPartiesResult{
				Party:        party,
				PartyMembers: []*sqlcClient.GetPartyMembersRow{},
			}
			partyIndex[party.PartyID] = resultItem

			return resultItem
		},
	)

	// Add members
	members, err := handler.Queries.GetPartyMembers(
		context.Request().Context(),
		util.Map(parties, func(party *sqlcClient.Party) int32 { return party.PartyID }),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, member := range members {
		partyIndex[member.PartyID].PartyMembers = append(partyIndex[member.PartyID].PartyMembers, member)
	}

	return context.JSON(http.StatusOK, data)
}
