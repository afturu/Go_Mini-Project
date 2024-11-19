package controllers

import (
	"tukerin-platform/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRecommendations(c echo.Context) error {
	item := c.QueryParam("item")
	if item == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Item parameter is required"})
	}

	// Membuat instance dari AIService
	aiService, err := services.NewAIService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Menggunakan service AI untuk mendapatkan rekomendasi
	response, err := aiService.FetchRecommendations(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"recommendation": response,
	})
}