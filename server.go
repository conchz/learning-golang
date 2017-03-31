package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderWWWAuthenticate,
			echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
