package main

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/shaardie/network-viewer/components"
	"github.com/shaardie/network-viewer/database"
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

	e.GET("/", func(c echo.Context) error {
		component := components.Home()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/subnet", func(c echo.Context) error {
		subnets := []database.Subnet{}
		if err := db.Find(&subnets).Error; err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		component := components.SubnetListPage(subnets)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/subnet/create", func(c echo.Context) error {
		component := components.SubnetCreateOrUpdatePage(nil)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/subnet/delete/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "id missing")
		}

		if err := db.Delete(&database.Subnet{}, id).Error; err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}

		return c.Redirect(http.StatusSeeOther, "/subnet")
	})

	e.POST("/subnet/create", func(c echo.Context) error {
		_, net, err := net.ParseCIDR(c.FormValue("subnet"))
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}
		scannerEnabled := false
		if c.FormValue("enabled") == "on" {
			scannerEnabled = true
		}

		hours, err := strconv.Atoi(c.FormValue("hours"))
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}
		minutes, err := strconv.Atoi(c.FormValue("minutes"))
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}

		seconds, err := strconv.Atoi(c.FormValue("seconds"))
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}

		subnet := database.Subnet{
			Subnet:         database.IPNet{IPNet: net},
			ScannerEnabled: scannerEnabled,
			ScannerInterval: time.Duration(hours)*time.Hour +
				time.Duration(minutes)*time.Minute +
				time.Duration(seconds)*time.Second,
			Metadata: database.Metadata{
				Comment: c.FormValue("comment"),
			},
		}
		if err := db.Save(&subnet).Error; err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}

		return c.Redirect(http.StatusSeeOther, "/subnet")
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

	e.GET("/ip/delete/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "id missing")
		}

		if err := db.Delete(&database.IP{}, id).Error; err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}

		return c.Redirect(http.StatusSeeOther, "/ip")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
