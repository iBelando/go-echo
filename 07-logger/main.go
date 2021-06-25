package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	myLog, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir o crear el archivo de logs: %v", err)
	}
	defer myLog.Close()

	logConfig := middleware.LoggerConfig{
		Output: myLog,
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}

	e.Use(middleware.LoggerWithConfig(logConfig))

	log.Println(e.Start(":7878"))

}
