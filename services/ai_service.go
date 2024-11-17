package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Prompt    string `json:"prompt"`
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func FetchRecommendations(item string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/completions"

	prompt := fmt.Sprintf("Berikan rekomendasi barang untuk kategori: %s", item)
	requestBody := OpenAIRequest{
		Prompt:    prompt,
		Model:     "text-davinci-003",
		MaxTokens: 100,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", errors.New("failed to create request")
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to call OpenAI API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("OpenAI API call failed")
	}

	var openAIResp OpenAIResponse
	json.NewDecoder(resp.Body).Decode(&openAIResp)

	return openAIResp.Choices[0].Text, nil
}