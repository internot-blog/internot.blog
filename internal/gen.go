package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ********** Image Structs **********

// Struct to represent model inputs for image generation
type modelInputs struct {
	Prompt            string  `json:"prompt"`
	NumInferenceSteps int     `json:"num_inference_steps"`
	GuidanceScale     float64 `json:"guidance_scale"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
}

// Struct to represent call inputs for the image generation request
type callInputs struct {
	ModelID       string `json:"MODEL_ID"`
	Pipeline      string `json:"PIPELINE"`
	Scheduler     string `json:"SCHEDULER"`
	SafetyChecker bool   `json:"safety_checker"`
}

// Struct for the image generation payload
type imagePayload struct {
	ModelInputs modelInputs `json:"modelInputs"`
	CallInputs  callInputs  `json:"callInputs"`
}

// Struct to represent the response for image generation
type imageResponse struct {
	ImageBase64 string `json:"image_base64"`
}

// ********** Text Structs **********

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
func GenImage(config Config, prompt string, width int, height int, steps int, guidance float64) ([]byte, error) {
	// Prepare the payload
	payload := imagePayload{
		ModelInputs: modelInputs{
			Prompt:            prompt,
			NumInferenceSteps: steps,
			GuidanceScale:     guidance,
			Width:             width,
			Height:            height,
		},
		CallInputs: callInputs{
			ModelID:       "stabilityai/stable-diffusion-2",
			Pipeline:      "StableDiffusionPipeline",
			Scheduler:     "LMSDiscreteScheduler",
			SafetyChecker: true,
		},
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create HTTP request with basic authentication
	req, err := http.NewRequest("POST", config.ModelConfig.ImageEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.SetBasicAuth(config.ModelConfig.User, config.ModelConfig.Password)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed with status code: %d, error: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response
	var response imageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Decode the base64 image
	imageData, err := base64.StdEncoding.DecodeString(response.ImageBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return imageData, nil
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
