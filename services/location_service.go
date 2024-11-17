package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type NearbyItem struct {
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Address  string  `json:"address"`
}

type LocationResponse struct {
	Results []struct {
		Name      string `json:"name"`
		Distance  int    `json:"distance"`
		Address   string `json:"vicinity"`
		Latitude  string `json:"geometry.location.lat"`
		Longitude string `json:"geometry.location.lng"`
	} `json:"results"`
}

func FetchNearbyItems(latitude, longitude string) ([]NearbyItem, error) {
	apiKey := os.Getenv("LOCATION_API_KEY")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=5000&type=store&key=%s", latitude, longitude, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch nearby items")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response from location API")
	}

	var locationResponse LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationResponse); err != nil {
		return nil, errors.New("failed to decode location data")
	}

	var items []NearbyItem
	for _, result := range locationResponse.Results {
		items = append(items, NearbyItem{
			Name:     result.Name,
			Distance: float64(result.Distance),
			Address:  result.Address,
		})
	}

	return items, nil
}