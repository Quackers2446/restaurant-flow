package main

import (
	"fmt"
	"net/http"
	_ "restaurant-flow/docs"
	"restaurant-flow/pkg/db"
	"restaurant-flow/pkg/handlers"
	"restaurant-flow/pkg/sqlcClient"
	"restaurant-flow/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
			"http://localhost:3334",
			"http://127.0.0.1:3334",
		},
		AllowCredentials: true,
	}))

	DB := db.Connect()

	defer db.CloseConnection(DB)

	// Setup sqlc
	queries := sqlcClient.New(DB)

	h := handlers.New(DB, queries)

	// Routes
	e.GET("/dummy-table", h.GetDummyTable)
	e.GET("/restaurants", h.GetRestaurants)
	e.GET("/restaurants/in-area", h.GetRestaurantsInArea)
	e.GET("/restaurants/search", h.GetRestaurantsSearch)
	e.GET("/restaurants/:id", h.GetRestaurant)

	e.POST("/review/create", h.CreateReview)
	e.POST("/review/update", h.UpdateReview)
	e.DELETE("/review/delete", h.DeleteReview)
	e.GET("/restaurants/:restaurantId/reviews", h.GetRestaurantReviews)
	e.POST("/internal/register", h.InternalRegister)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
