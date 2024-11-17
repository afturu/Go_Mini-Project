package routes

import (
	"tukerin-platform/controllers"
	"github.com/labstack/echo/v4"
)

func AIRoutes(e *echo.Echo) {
	e.GET("/recommendations", controllers.GetRecommendations)
}