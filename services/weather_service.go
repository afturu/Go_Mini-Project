package services

import (
    "encoding/json"
    "errors"
    "net/http"
)

type WeatherService struct {
    APIKey string
}

func NewWeatherService(apiKey string) *WeatherService {
    return &WeatherService{APIKey: apiKey}
}

type WeatherResponse struct {
    Location struct {
        Name    string `json:"name"`
        Region  string `json:"region"`
        Country string `json:"country"`
    } `json:"location"`
    Current struct {
        TempC      float64 `json:"temp_c"`
        TempF      float64 `json:"temp_f"`
        Condition  struct {
            Text string `json:"text"`
            Icon string `json:"icon"`
        } `json:"condition"`
    } `json:"current"`
}

func (ws *WeatherService) GetCurrentWeather(city string) (*WeatherResponse, error) {
    url := "http://api.weatherapi.com/v1/current.json?key=" + ws.APIKey + "&q=" + city + "&aqi=no"
    
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to fetch weather data")
    }

    var weather WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
        return nil, err
    }

    return &weather, nil
}