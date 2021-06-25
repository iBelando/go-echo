package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/img/:name", func(c echo.Context) error {
		p := c.Param("name")
		switch {
		case p == "inline":
			return c.Inline("img/wallpaper.png", "gdg-wallpaper")
		case p == "att":
			return c.Attachment("img/wallpaper.png", "gdg-wallpaper")
		case p == "file":
			return c.File("img/wallpaper.png")
		}
		return c.HTML(http.StatusNotFound, "<h1>No Encontrado</h1>")
	})
	e.Start(":8080")
}
