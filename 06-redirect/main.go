package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	go func() {
		e.Pre(middleware.HTTPSRedirect())
	}()

	e.GET("/", index)
	e.GET("/divicion", anterior)
	e.GET("/division", nueva)
	e.Logger.Fatal(e.StartTLS(":443", "./certificates/cert.pem", "./certificates/key.pem"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "HOLA MUNDO!")
}

func anterior(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/division")
}

func nueva(c echo.Context) error {
	return c.String(http.StatusOK, "Este es el contenido super especial")
}
