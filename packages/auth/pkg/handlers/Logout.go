package handlers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (handler Handler) Logout(context echo.Context) (err error) {
	cookie, err := context.Cookie("refresh_token")

	if cookie == nil || err != nil {
		return context.NoContent(http.StatusNoContent)
	}

	refreshToken := cookie.Value

	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.HttpOnly = true

	context.SetCookie(cookie)

	decodedRefreshToken, err := base64.URLEncoding.DecodeString(refreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	currentSession, err := handler.Queries.GetSession(context.Request().Context(), decodedRefreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = handler.Queries.InvalidateSession(context.Request().Context(), currentSession.SessionID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.NoContent(http.StatusNoContent)
}
