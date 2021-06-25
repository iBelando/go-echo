package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "1,2021-06-22,NACHO,BELANDO,23")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hola Mundo</h1>")
	})

	e.GET("/no-content", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8080")
}
