package router

import (
	"github.com/geek-teru/simple-task-app/handler"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	// e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
	// 	AllowMethods: []string{http.MethodGet},
	// }))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	e.GET("/healthcheck", handler.Healthcheck)

	return e
}
