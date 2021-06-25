package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/division", Dividir)

	e.Start(":8080")
}

func Dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	a := 2500 / f
	return c.String(http.StatusOK, strconv.Itoa(a))
}
