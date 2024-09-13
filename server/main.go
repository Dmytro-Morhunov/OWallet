package main

import (
	"OWallet.com/app/controller"
	"OWallet.com/database"
	_ "OWallet.com/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title OWallet API
// @version 1.0
// @description This is OWallet API

// @host localhost:8080
func main() {
	// PostgresSQL connect
	database.ConnectToDB()

	// Echo instance
	e := echo.New()
	apiGroup := e.Group("/api")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User Routes
	controller.InitUserController(apiGroup)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
