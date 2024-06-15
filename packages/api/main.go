package main

import (
	"net/http"
	_ "restaurant-flow/docs"
	"restaurant-flow/pkg/db"
	"restaurant-flow/pkg/handlers"
	"restaurant-flow/pkg/sqlcClient"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// https://echo.labstack.com/docs/request#validate-data
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

// @title			Restaurant Flow
// @description	Restaurant reviews for UW Students
// @BasePath		/
func main() {
	// Echo instance
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Resembling common log format
		Format: `${remote_ip} - - [${time_rfc3339}] "${method} ${path} ${protocol}" ${status} ${bytes_out} ${latency_human} ${error}` + "\n",
	}))
	e.Use(middleware.Recover())

	DB := db.Connect()

	defer db.CloseConnection(DB)

	// Setup sqlc
	queries := sqlcClient.New(DB)

	h := handlers.New(DB, queries)

	// Routes
	e.GET("/dummy-table", h.GetDummyTable)
	e.GET("/restaurants", h.GetRestaurants)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
