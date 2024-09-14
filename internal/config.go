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
	// TODO: add (optional) text prompt starter maybe?
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

func FetchConfig() ModelConfig {
	user := os.Getenv("AUTH_USER")
	if user == "" {
		panic("Error: Environment variable 'AUTH_USER' not set.")
	}

	pass := os.Getenv("AUTH_PASS")
	if pass == "" {
		panic("Error: Environment variable 'AUTH_PASS' not set.")
	}

	txtUrl := os.Getenv("OLLAMA_URL")
	if txtUrl == "" {
		panic("Error: Environment variable 'OLLAMA_URL' not set.")
	}

	imgUrl := os.Getenv("STABLE_URL")
	if imgUrl == "" {
		panic("Error: Environment variable 'STABLE_URL' not set.")
	}

	return ModelConfig{
		User:          user,
		Password:      pass,
		TextEndpoint:  imgUrl,
		ImageEndpoint: txtUrl,
	}
}
