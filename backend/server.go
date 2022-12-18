package main

import (
	"vmd/api/echo"
	"vmd/api/models"
)

func main() {
	models.ConnectDatabase()
	echo.Echo()
}
