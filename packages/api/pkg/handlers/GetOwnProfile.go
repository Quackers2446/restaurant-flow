package handlers

import (
	"net/http"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

// GetOwnProfile
//
//	@Summary	get own profile based on Authorization header
//
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	200				{object}	sqlcClient.User
//	@Param		Authorization	header		string	true	"access token"
//	@Failure	400				{object}	echo.HTTPError
//	@Failure	500				{object}	echo.HTTPError
//	@Router		/users/own-profile [get]
func (handler Handler) GetOwnProfile(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := handler.Queries.GetFullUser(context.Request().Context(), claims.Subject)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, user)
}
