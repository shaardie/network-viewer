package main

import (
	"github.com/shaardie/network-viewer/components"
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

	e.GET("/", func(c echo.Context) error {
		component := components.Home()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	// ip
	e.GET("/ip", func(c echo.Context) error {
		ips := []database.IP{}
		if err := db.Find(&ips).Error; err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		component := components.IPListPage(ips)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
