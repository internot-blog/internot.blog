package internal

import (
	"fmt"
	"math/rand"
)

func GenImagePrompt(cfg Config) string {
	subject := cfg.PromptConfig.ImagePromptSubject
	action := cfg.PromptConfig.ImagePromptAction

	prompt_subject := subject[rand.Intn(len(subject))]
	prompt_action := action[rand.Intn(len(action))]

	return fmt.Sprintf("Generate me an ad for %s where %s is %s", prompt_subject, prompt_subject, prompt_action)
}

func GenTextPrompt(cfg Config) string {
	subject := cfg.PromptConfig.TextPromptSubject
	action := cfg.PromptConfig.TextPromptAction

	prompt_subject := subject[rand.Intn(len(subject))]
	prompt_action := action[rand.Intn(len(action))]

	return fmt.Sprintf("Write me a blog post about %s and how they %s", prompt_subject, prompt_action)
}
