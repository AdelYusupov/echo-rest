package main

import (
	"echo-rest/cmd/handlers"
	"echo-rest/cmd/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDb()
	e.GET("/", handlers.Home)
	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurements)
	e.Logger.Fatal(e.Start(":8080"))
}
