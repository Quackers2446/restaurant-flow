package main

import (
	_ "restaurant-flow/docs"
	"restaurant-flow/pkg/db"
	"restaurant-flow/pkg/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Restaurant Flow
// @description	Restaurant reviews for UW Students
// @BasePath		/
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	DB := db.Connect()

	defer db.CloseConnection(DB)

	h := handlers.New(DB)

	// Routes
	e.GET("/dummy-table", h.GetDummyTable)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
