package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shaardie/network-viewer/database"
	"gorm.io/gorm"
)

type Subnet struct {
	ID              uint          `json:"id"`
	Subnet          string        `json:"subnet"`
	ScannerEnabled  bool          `json:"scanner_enabled"`
	ScannerInterval time.Duration `json:"scanner_interval"`
	LastScan        time.Time     `json:"last_scan"`
	Comment         string        `json:"comment"`
}

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
		os := make([]Subnet, 0, len(subnets))
		for _, sn := range subnets {
			os = append(os, Subnet{
				ID:              sn.ID,
				Subnet:          sn.Subnet.String(),
				ScannerEnabled:  sn.ScannerEnabled,
				ScannerInterval: sn.ScannerInterval,
				LastScan:        sn.LastScan,
				Comment:         sn.Comment,
			})
		}
		return c.JSON(http.StatusOK, os)
	}
}

func (s server) subnetGetAPI() echo.HandlerFunc {
	type input struct {
		ID uint `param:"id"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		sn := database.Subnet{
			Model: gorm.Model{
				ID: i.ID,
			},
		}
		if err := s.db.First(&sn).Error; err != nil {
			return echo.ErrBadRequest.SetInternal(err)
		}
		o := Subnet{
			ID:              sn.ID,
			Subnet:          sn.Subnet.String(),
			ScannerEnabled:  sn.ScannerEnabled,
			ScannerInterval: sn.ScannerInterval,
			LastScan:        sn.LastScan,
			Comment:         sn.Comment,
		}
		return c.JSON(http.StatusOK, o)
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
			ScannerInterval: i.ScannerInterval,
			ScannerEnabled:  i.ScannerEnabled,
			Subnet: database.IPNet{
				IPNet: ipNet,
			},
			Comment: i.Comment,
		}
		err = s.subnetCreate(&subnet)
		if err != nil {
			return echo.ErrInternalServerError.SetInternal(err)
		}

		return c.JSON(http.StatusCreated, subnet)
	}
}

func (s server) subnetReplaceAPI() echo.HandlerFunc {
	type input struct {
		ID              uint          `param:"id"`
		Subnet          string        `json:"subnet"`
		ScannerEnabled  bool          `json:"scanner_enabled"`
		ScannerInterval time.Duration `json:"scanner_interval"`
		Comment         string        `json:"comment"`
	}
	return func(c echo.Context) error {
		var i input
		if err := c.Bind(&i); err != nil {
			return echo.ErrBadRequest.SetInternal(fmt.Errorf("unable to bind, %w", err))
		}
		_, ipNet, err := net.ParseCIDR(i.Subnet)
		if err != nil {
			return echo.ErrBadRequest.SetInternal(fmt.Errorf("unable to parse cidr, %w", err))
		}
		subnet := database.Subnet{
			Model: gorm.Model{
				ID: i.ID,
			},
			ScannerInterval: i.ScannerInterval,
			ScannerEnabled:  i.ScannerEnabled,
			Subnet: database.IPNet{
				IPNet: ipNet,
			},
			Comment: i.Comment,
		}
		if err := s.db.Save(subnet).Error; err != nil {
			return echo.ErrInternalServerError.SetInternal(fmt.Errorf("unable to save subnet, %w", err))
		}

		return c.JSON(http.StatusOK, subnet)
	}
}
