package main

import (
	"github.com/shaardie/network-viewer/database"
	"github.com/shaardie/network-viewer/server"
	"github.com/shaardie/network-viewer/subnetscanner"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init database
	db, err := database.Init("./network-viewer.db")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	scanner := subnetscanner.New(db)
	scanner.Start()

	s := server.New(db)
	s.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
