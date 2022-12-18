package controllers

import (
	"net/http"
	"vmd/api/models"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	models.DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, user)
}
