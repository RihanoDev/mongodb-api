package main

import (
	"mongodb-api/config"
	"mongodb-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
	// })

	//run database
	config.ConnectDB()

	//routes
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start((":6000")))
}
