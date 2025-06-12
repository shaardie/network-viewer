package server

import (
	"fmt"
	"net/http"

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

func (s server) subnetCreate(subnet *database.Subnet) error {
	return s.db.Create(subnet).Error
}

func (s server) subnetDelete(id uint) error {
	return s.db.Delete(&database.Subnet{}, id).Error
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
