package handlers

import (
	"errors"
	"net/http"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/labstack/echo/v4"
)

type registerBody struct {
	UserId   string `body:"userId" validate:"required,uuid"`
	Email    string `body:"email" validate:"required,email"`
	Username string `body:"username" validate:"required"`
	Name     string `body:"name" validate:"required"`
}

// InternalRegister
//
//	@Summary	Register a user, only to be called from the auth service
//
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	204
//	@Param		requestBody	body		registerBody	true	"request body"
//	@Failure	400			{object}	echo.HTTPError
//	@Failure	500			{object}	echo.HTTPError
//	@Router		/internal/register [post]
func (handler Handler) InternalRegister(context echo.Context) (err error) {
	_, claims, err := util.ValidateTokenHeader(&context.Request().Header)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if claims.Subject != "uwEats" {
		return echo.NewHTTPError(http.StatusForbidden, errors.New("not allowed"))
	}

	body, err := util.ValidateInput(&context, &registerBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = handler.Queries.CreateUser(context.Request().Context(), sqlcClient.CreateUserParams{
		UserID:   body.UserId,
		Username: body.Username,
		Email:    body.Email,
		Name:     body.Name,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return context.NoContent(http.StatusNoContent)
}
