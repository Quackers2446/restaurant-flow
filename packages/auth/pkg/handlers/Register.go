package handlers

import (
	"encoding/base64"
	"net/http"
	"restaurant-flow-auth/pkg/sqlcClient"
	"restaurant-flow-auth/pkg/util"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	passwordValidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const minEntropyBits = 50 // Between 50 and 70 is "reasonable"
const bcryptCost = 14

type registerBody struct {
	Email    string `body:"email" validate:"required,email"`
	Password string `body:"password" validate:"required"`
}

func (handler Handler) Register(context echo.Context) (err error) {
	body, err := util.ValidateInput(&context, &registerBody{})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = passwordValidator.Validate(body.Password, minEntropyBits)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcryptCost)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	newUserId := uuid.New().String()

	transaction, err := handler.DB.BeginTx(context.Request().Context(), nil)

	defer transaction.Rollback()

	_, err = transaction.Exec( /*sql*/ `
		insert into user set
			user_id=unhex(replace(?,'-','')),
			email=?;
	`, newUserId, body.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = transaction.Exec( /*sql*/ `
		insert into user_auth set
			user_id=unhex(replace(?,'-','')),
			password_hash=?;
	`, newUserId, hashedPassword)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = transaction.Commit()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Same as login:
	refreshToken, err := util.GenerateRandomBytes(64)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	exp := time.Now().Add(24 * time.Hour * refreshTokenExpDays)

	err = handler.Queries.CreateSession(context.Request().Context(), sqlcClient.CreateSessionParams{
		UserID:     newUserId,
		IpAddr:     context.Echo().IPExtractor(context.Request()),
		UserAgent:  context.Request().UserAgent(),
		ExpiresAt:  exp,
		RefreshKey: refreshToken,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	accessToken, err := util.GenerateJWT(newUserId)

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
