package routes

import (
	"tukerin-platform/controllers"
	"github.com/labstack/echo/v4"
)

func LocationRoutes(e *echo.Echo) {
	e.GET("/locations/nearby", controllers.GetNearbyItems)
}