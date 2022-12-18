package echo

import (
	"vmd/api/controllers"
	"vmd/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Echo() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	controllers.InitRouting(e)
	e.Logger.Fatal(e.Start(config.GetEnv("PORT")))
}
