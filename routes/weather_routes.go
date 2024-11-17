package routes

import (
	"tukerin-platform/controllers"
	"github.com/labstack/echo/v4"
)

func WeatherRoutes(e *echo.Echo) {
	e.GET("/weather", controllers.GetWeather)
}