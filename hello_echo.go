package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("hello.html")
}

func main() {
	e := echo.New()
	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/bye", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bye, World!")
	})
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1323"))

}
