package controllers

import (
	"tukerin-platform/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetWeather(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "City parameter is required"})
	}

	response, err := services.FetchWeather(city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}