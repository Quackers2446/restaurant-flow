package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (handler Handler) ClearSessions(context echo.Context) (err error) {
	cookie, err := context.Cookie("refresh_token")

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	refreshToken := cookie.Value

	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.HttpOnly = true

	context.SetCookie(cookie)

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

	err = handler.Queries.InvalidateAllSessions(context.Request().Context(), string(currentSession.UserID))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.NoContent(http.StatusNoContent)
}
