package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shaardie/network-viewer/components"
	"github.com/shaardie/network-viewer/database"
)

func (s server) subnetList() ([]database.Subnet, error) {
	subnets := []database.Subnet{}
	if err := s.db.Find(&subnets).Error; err != nil {
		return nil, fmt.Errorf("unable to get subnets, %w", err)
	}
	return subnets, nil
}

func (s server) subnetListAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		subnets, err := s.subnetList()
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		return c.JSON(http.StatusOK, subnets)
	}
}

func (s server) subnetListPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		subnets, err := s.subnetList()
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		return components.SubnetListPage(subnets).Render(c.Request().Context(), c.Response().Writer)
	}
}

func (s server) subnetDelete(id uint) error {
	return s.db.Delete(&database.Subnet{}, id).Error
}

func (s server) subnetDeleteAPI() echo.HandlerFunc {
	type input struct {
		ID uint `param:"id"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		if err := s.subnetDelete(i.ID); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		return nil
	}
}

func (s server) subnetDeletePage() echo.HandlerFunc {
	type input struct {
		ID uint `param:"id"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		if err := s.subnetDelete(i.ID); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		return c.Redirect(http.StatusSeeOther, "/subnet")
	}
}

func (s server) subnetCreate(subnet *database.Subnet) error {
	return s.db.Create(subnet).Error
}

func (s server) subnetCreateAPI() echo.HandlerFunc {
	type input struct {
		Subnet          string        `json:"subnet"`
		ScannerEnabled  bool          `json:"scanner_enabled"`
		ScannerInterval time.Duration `json:"scanner_interval"`
		Comment         string        `json:"comment"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		_, ipNet, err := net.ParseCIDR(i.Subnet)
		if err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		subnet := database.Subnet{
			Metadata: database.Metadata{
				Comment: i.Comment,
			},
			ScannerInterval: i.ScannerInterval,
			ScannerEnabled:  i.ScannerEnabled,
			Subnet: database.IPNet{
				IPNet: ipNet,
			},
		}
		err = s.subnetCreate(&subnet)
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}

		return c.JSON(http.StatusCreated, subnet)
	}
}

func (s server) subnetCreateFormPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		component := components.SubnetCreateOrUpdatePage(nil)
		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}

func (s server) subnetCreatePage() echo.HandlerFunc {
	type input struct {
		Subnet         string `form:"subnet"`
		ScannerEnabled bool   `form:"scanner_enabled"`
		ScannerHours   uint   `form:"scanner_hours"`
		ScannerMinutes uint   `form:"scanner_minutes"`
		ScannerSeconds uint   `form:"scanner_seconds"`
		Comment        string `form:"comment"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}
		_, ipNet, err := net.ParseCIDR(i.Subnet)
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}
		subnet := database.Subnet{
			Metadata: database.Metadata{
				Comment: i.Comment,
			},
			ScannerInterval: time.Duration(i.ScannerHours)*time.Hour +
				time.Duration(i.ScannerMinutes)*time.Minute +
				time.Duration(i.ScannerSeconds)*time.Second,
			ScannerEnabled: i.ScannerEnabled,
			Subnet: database.IPNet{
				IPNet: ipNet,
			},
		}
		err = s.subnetCreate(&subnet)
		if err != nil {
			component := components.SubnetCreateOrUpdatePage(err)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}

		return c.Redirect(http.StatusSeeOther, "/subnet")
	}
}
