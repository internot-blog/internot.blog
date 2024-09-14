package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct for LLM conversation message
type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Struct for LLM request payload
type textPayload struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"`
}

// Struct for LLM response
type textResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

// TODO:
func GenImage(Config) (string, error) {
	return "", nil
}

// TODO:
func GenText(config Config, prompt string) (string, error) {
	payload := textPayload{
		Model: "llama3.1:8b",
		Messages: []message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream: false,
	}

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create HTTP request with basic authentication
	req, err := http.NewRequest("POST", config.ModelConfig.TextEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.SetBasicAuth(config.ModelConfig.User, config.ModelConfig.Password)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Ignoring error here for simplicity
		return "", fmt.Errorf("failed with status code: %d, error: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal JSON response
	var response textResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Extract LLM response text
	return response.Message.Content, nil
}
