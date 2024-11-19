package routes

import (
    "os"

    "github.com/labstack/echo/v4"
    "tukerin-platform/controllers"
    "tukerin-platform/services"
)

func WeatherRoutes(e *echo.Echo) {
    apiKey := os.Getenv("WEATHER_API_KEY")
    if apiKey == "" {
        panic("WEATHER_API_KEY is not set in environment variables")
    }

    weatherService := services.NewWeatherService(apiKey)
    weatherController := controllers.NewWeatherController(weatherService)

    e.GET("/api/weather", weatherController.GetWeather) // Endpoint untuk mendapatkan cuaca
}