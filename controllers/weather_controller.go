package controllers

import (
    "net/http"
    "tukerin-platform/services"

    "github.com/labstack/echo/v4"
)

type WeatherController struct {
    weatherService *services.WeatherService
}

func NewWeatherController(weatherService *services.WeatherService) *WeatherController {
    return &WeatherController{weatherService}
}

func (wc *WeatherController) GetWeather(ctx echo.Context) error {
    city := ctx.QueryParam("city")
    if city == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{
            "error": "City parameter is required",
        })
    }

    weather, err := wc.weatherService.GetCurrentWeather(city)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{
            "error": err.Error(),
        })
    }

    return ctx.JSON(http.StatusOK, weather)
}