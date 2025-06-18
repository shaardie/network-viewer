package server

import (
	"github.com/labstack/echo/v4"
	"github.com/shaardie/network-viewer/components"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
}

func New(db *gorm.DB) server {
	return server{
		db: db,
	}
}

func (s server) SetupRoutes(e *echo.Echo) {
	e.GET("/subnet", s.subnetListPage())
	e.GET("/subnet/delete/:id", s.subnetDeletePage())
	e.GET("/subnet/create", s.subnetCreateFormPage())
	e.POST("/subnet/create", s.subnetCreatePage())

	e.GET("/ip", s.ipListPage())
	e.GET("/ip/delete/:id", s.ipDeletePage())

	e.GET("/api/v1/subnet", s.subnetListAPI())
	e.POST("/api/v1/subnet", s.subnetCreateAPI())
	e.DELETE("/api/v1/subnet/:id", s.subnetDeleteAPI())
	e.GET("/api/v1/subnet/:id", s.subnetGetAPI())
	e.GET("/api/v1/ip", s.ipListAPI())
	e.DELETE("/api/v1/ip/:id", s.ipDeleteAPI())
}

func (s server) rootPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return components.Home().Render(c.Request().Context(), c.Response().Writer)
	}
}
