package controllers

import (
	"tukerin-platform/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNearbyItems(c echo.Context) error {
	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")
	if latitude == "" || longitude == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Latitude and longitude parameters are required"})
	}

	response, err := services.FetchNearbyItems(latitude, longitude)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}