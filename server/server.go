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
	e.GET("/subnet", s.subnetListPage())
	e.GET("/api/v1/subnet", s.subnetListAPI())
	e.GET("/subnet/delete/:id", s.subnetDeletePage())
	e.DELETE("/api/v1/subnet/:id", s.subnetDeleteAPI())
}
