package handlers

import (
	"net/http"
	"restaurant-flow-auth/pkg/util"

	"github.com/labstack/echo/v4"
	passwordValidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const minEntropyBits = 60 // Between 50 and 70 is "reasonable"
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

	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcryptCost)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = handler.DB.Exec( /*sql*/ `
		start transaction;

		set new_user_id = unhex(uuid(),'-','')); -- Is this ok???

		insert into user set
			user_id=new_user_id,
			email=?;

		insert into user_auth
			user_id=new_user_id,
			password_hash=?;

		commit;
	`, body.Email, bytes)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.NoContent(http.StatusNoContent)
}
