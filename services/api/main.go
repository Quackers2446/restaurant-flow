package main

import (
	"restaurant-flow/pkg/db"
	"restaurant-flow/pkg/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
