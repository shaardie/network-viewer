package server

import (
	"github.com/labstack/echo/v4"
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
	e.GET("/api/v1/subnet", s.subnetListAPI())
	e.POST("/api/v1/subnet", s.subnetCreateAPI())
	e.DELETE("/api/v1/subnet/:id", s.subnetDeleteAPI())
	e.GET("/api/v1/subnet/:id", s.subnetGetAPI())
	e.PUT("/api/v1/subnet/:id", s.subnetReplaceAPI())
	e.GET("/api/v1/ip", s.ipListAPI())
	e.DELETE("/api/v1/ip/:id", s.ipDeleteAPI())
}
