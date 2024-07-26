package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (handler Handler) GetSessions(context echo.Context) (err error) {
	cookie, err := context.Cookie("refresh_token")

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	refreshToken := cookie.Value

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

	fmt.Printf("CURRENT SESSION %#v\n", currentSession)

	err = handler.Queries.UpdateSessionLastUsed(context.Request().Context(), currentSession.SessionID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sessions, err := handler.Queries.GetUserSessions(context.Request().Context(), string(*currentSession.UserIDText))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, sessions)
}
