package handlers

import (
	"errors"
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

	currentSession, err := handler.Queries.GetSession(context.Request().Context(), refreshToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if currentSession.ExpiresAt.Before(time.Now()) {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("refresh token expired"))

	}

	err = handler.Queries.UpdateSessionLastUsed(context.Request().Context(), currentSession.SessionID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sessions, err := handler.Queries.GetUserSessions(context.Request().Context(), string(currentSession.UserID))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusNoContent, sessions)
}
