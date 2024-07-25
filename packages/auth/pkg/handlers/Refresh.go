package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"restaurant-flow-auth/pkg/util"
	"time"

	"github.com/labstack/echo/v4"
)

type refreshResponse struct {
	AccessToken string `json:"accessToken"`
}

func (handler Handler) Refresh(context echo.Context) (err error) {
	cookie, err := context.Cookie("refresh_token")

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	refreshToken := cookie.Value

	fmt.Println(refreshToken)

	decodedRefreshToken, err := base64.URLEncoding.DecodeString(refreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	currentSession, err := handler.Queries.GetSession(context.Request().Context(), decodedRefreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if currentSession.ExpiresAt.Before(time.Now()) || !currentSession.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("refresh token expired"))
	}

	err = handler.Queries.UpdateSessionLastUsed(context.Request().Context(), currentSession.SessionID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	accessToken, err := util.GenerateJWT(string(*currentSession.UserIDText))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, refreshResponse{AccessToken: accessToken})
}
