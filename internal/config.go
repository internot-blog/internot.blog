package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type ModelConfig struct {
	User          string
	Password      string
	TextEndpoint  string
	ImageEndpoint string
}

type PromptConfig struct {
	TextPromptSubject  []string `json:"text_prompt_subject"`
	TextPromptAction   []string `json:"text_prompt_action"`
	ImagePromptSubject []string `json:"image_prompt_subject"`
	ImagePromptAction  []string `json:"image_prompt_action"`
}

type Config struct {
	ModelConfig  ModelConfig
	PromptConfig PromptConfig
}

func ReadPromptFile(promptPath string) PromptConfig {
	contents, err := os.ReadFile(promptPath)
	if err != nil {
		panic(fmt.Sprintln("Failed to open config file: ", err))
	}

	var out PromptConfig
	err = json.Unmarshal(contents, &out)
	if err != nil {
		panic(fmt.Sprintln("Failed to parse config file: ", err))
	}

	return out
}

func GetModelConfig() ModelConfig {
	user, exists := os.LookupEnv("AUTH_USER")
	if !exists {
		panic("Error: Environment variable 'AUTH_USER' not set.")
	}

	pass, exists := os.LookupEnv("AUTH_PASS")
	if !exists {
		panic("Error: Environment variable 'AUTH_PASS' not set.")
	}

	txtUrl, exists := os.LookupEnv("OLLAMA_URL")
	if !exists {
		panic("Error: Environment variable 'OLLAMA_URL' not set.")
	}

	imgUrl, exists := os.LookupEnv("STABLE_URL")
	if !exists {
		panic("Error: Environment variable 'STABLE_URL' not set.")
	}

	return ModelConfig{
		User:          user,
		Password:      pass,
		TextEndpoint:  txtUrl,
		ImageEndpoint: imgUrl,
	}
}

func GetConfig(promptPath string) Config {
	return Config{
		ModelConfig:  GetModelConfig(),
		PromptConfig: ReadPromptFile(promptPath),
	}
}
