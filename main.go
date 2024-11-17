package main

import (
    "github.com/labstack/echo/v4"
    "tukerin-platform/config"
    "tukerin-platform/routes"
)

func main() {
    config.InitDB()

    config.LoadEnvironment()

    e := echo.New()

    routes.InitRoutes(e)
    routes.WeatherRoutes(e)
	routes.AIRoutes(e)
	routes.LocationRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}