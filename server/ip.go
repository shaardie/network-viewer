package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shaardie/network-viewer/components"
	"github.com/shaardie/network-viewer/database"
)

func (s server) ipList() ([]database.IP, error) {
	ips := []database.IP{}
	if err := s.db.Find(&ips).Error; err != nil {
		return nil, fmt.Errorf("unable to get subnets, %w", err)
	}
	return ips, nil
}

func (s server) ipListAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		ips, err := s.ipList()
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		return c.JSON(http.StatusOK, ips)
	}
}

func (s server) ipListPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ips, err := s.ipList()
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}
		return components.IPListPage(ips).Render(c.Request().Context(), c.Response().Writer)
	}
}

func (s server) ipDelete(id uint) error {
	return s.db.Delete(&database.IP{}, id).Error
}

func (s server) ipDeleteAPI() echo.HandlerFunc {
	type input struct {
		ID uint `param:"id"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		if err := s.ipDelete(i.ID); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		return nil
	}
}

func (s server) ipDeletePage() echo.HandlerFunc {
	type input struct {
		ID uint `param:"id"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		if err := s.ipDelete(i.ID); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		return c.Redirect(http.StatusSeeOther, "/ip")
	}
}
