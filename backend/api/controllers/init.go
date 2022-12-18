package controllers

import (
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	e.GET("/users/:id", GetUser)
	e.POST("/users", CreateUser)
}
