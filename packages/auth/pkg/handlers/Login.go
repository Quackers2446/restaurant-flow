package handlers

import (
	"encoding/base64"
	"net/http"
	"restaurant-flow-auth/pkg/sqlcClient"
	"restaurant-flow-auth/pkg/util"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Email    string `body:"email" validate:"required,email"`
	Password string `body:"password" validate:"required"`
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}

const refreshTokenExpDays = 90

func (handler Handler) Login(context echo.Context) (err error) {
	body, err := util.ValidateInput(&context, &loginBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	loginDetails, err := handler.Queries.GetLoginDetails(context.Request().Context(), body.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginDetails.PasswordHash), []byte(body.Password))

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	refreshToken, err := util.GenerateRandomBytes(64)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	exp := time.Now().Add(24 * time.Hour * refreshTokenExpDays)

	err = handler.Queries.CreateSession(context.Request().Context(), sqlcClient.CreateSessionParams{
		UserID:     *loginDetails.UserIDText,
		IpAddr:     context.Echo().IPExtractor(context.Request()),
		UserAgent:  context.Request().UserAgent(),
		ExpiresAt:  exp,
		RefreshKey: refreshToken,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	accessToken, err := util.GenerateJWT(*loginDetails.UserIDText)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = base64.URLEncoding.EncodeToString(refreshToken)
	cookie.Expires = exp
	cookie.HttpOnly = true

	context.SetCookie(cookie)

	return context.JSON(http.StatusOK, loginResponse{
		AccessToken: accessToken,
	})
}
