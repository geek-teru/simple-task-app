package router

import (
	"os"

	"github.com/geek-teru/simple-task-app/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, userHandler handler.UserHandler) {
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.GET("/healthcheck", handler.Healthcheck)

	e.POST("/signup", userHandler.SignUp)
	e.POST("/signin", userHandler.SignIn)

	u := e.Group("/user")
	u.Use(JwtAuth())
	u.GET("", userHandler.GetUserProfile)
	u.PUT("", userHandler.UpdateUserProfile)

}

func JwtAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "header:Authorization:Bearer ",
	})
}
