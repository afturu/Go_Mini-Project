package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func FetchWeather(city string) (map[string]interface{}, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch weather data")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response from weather API")
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, errors.New("failed to decode weather data")
	}

	return map[string]interface{}{
		"city":        city,
		"temperature": weather.Main.Temp,
		"description": weather.Weather[0].Description,
	}, nil
}