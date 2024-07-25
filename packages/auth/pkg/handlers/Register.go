package handlers

import (
	"net/http"
	"restaurant-flow-auth/pkg/util"

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

	return context.NoContent(http.StatusNoContent)
}
