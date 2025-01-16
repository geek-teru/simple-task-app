package router

import (
	"github.com/geek-teru/simple-task-app/handler"

	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, userHandler handler.UserHandler) {
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.GET("/healthcheck", handler.Healthcheck)

	e.POST("/signup", userHandler.SignUp)
	e.POST("/signin", userHandler.SignUp)

	u := e.Group("user/")
	u.GET("user/", userHandler.GetUserProfile)
	u.PUT("user/", userHandler.UpdateUserProfile)

}
