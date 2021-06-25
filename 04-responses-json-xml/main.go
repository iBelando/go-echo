package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	ps := make([]Person, 0)

	p := Person{
		FirstName: "Ignacio",
		LastName:  "Belando",
		Age:       23,
	}

	ps = append(ps, p)
	ps = append(ps, p)
	ps = append(ps, p)
	ps = append(ps, p)

	e := echo.New()
	e.GET("/json", func(c echo.Context) error {

		return c.JSON(http.StatusOK, ps)
	})
	e.GET("/xml", func(c echo.Context) error {

		return c.XML(http.StatusOK, ps)
	})

	e.Start(":8080")
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
