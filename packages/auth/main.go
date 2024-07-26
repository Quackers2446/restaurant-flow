package main

import (
	"fmt"
	"net/http"

	"restaurant-flow-auth/pkg/db"
	"restaurant-flow-auth/pkg/handlers"
	"restaurant-flow-auth/pkg/sqlcClient"
	"restaurant-flow-auth/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// For stdout coloring
const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorDarkGray = 90
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
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Resembling common log format
		Format: fmt.Sprintf(
			"%s - - [%s] %s %s %s ${latency_human} ${error}\n",
			util.Color("${remote_ip}", colorCyan),
			util.Color("${time_rfc3339}", colorDarkGray),
			util.Color(`"${method} ${path} ${protocol}"`, colorGreen),
			util.Color("${status}", colorCyan),
			util.Color("${bytes_out}", colorCyan),
		),
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		},
		AllowCredentials: true,
	}))

	DB := db.Connect()

	defer db.CloseConnection(DB)

	// Setup sqlc
	queries := sqlcClient.New(DB)

	h := handlers.New(DB, queries)

	// Routes
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/get-sessions", h.GetSessions)
	e.GET("/refresh", h.Refresh)
	e.POST("/clear-sessions", h.ClearSessions)
	e.POST("/logout", h.Logout)

	// Start server
	e.Logger.Fatal(e.Start(":3334"))
}
