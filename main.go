package main

import (
    "github.com/labstack/echo/v4"
    "tukerin-platform/config"
    "tukerin-platform/routes"
)

func main() {
    config.InitDB()

    e := echo.New()

    routes.InitRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}