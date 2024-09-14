package internal

import (
	"bytes"
	"fmt"
	"math/rand"
	"text/template"
)

func GenImagePrompt(cfg Config) string {
	subject := cfg.PromptConfig.ImagePromptSubject
	action := cfg.PromptConfig.ImagePromptAction

	promptSubject := subject[rand.Intn(len(subject))]
	promptAction := action[rand.Intn(len(action))]

	return fmt.Sprintf("Generate me an ad for %s where %s is %s", promptSubject, promptSubject, promptAction)
}

func GenTextPrompt(cfg Config) string {
	subject := cfg.PromptConfig.TextPromptSubject
	action := cfg.PromptConfig.TextPromptAction

	promptSubject := subject[rand.Intn(len(subject))]
	promptActioin := action[rand.Intn(len(action))]

	tmpl, err := template.ParseFiles("templates/prompt.txt")
	// tmpl, err := template.ParseFiles("../templates/prompt.txt")
	if err != nil {
		panic(fmt.Sprintln("Error reading templates/prompt.txt: ", err))
	}

	data := struct {
		PromptSubject string
		PromptAction  string
	}{
		PromptSubject: promptSubject,
		PromptAction:  promptActioin,
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		panic(fmt.Sprintln("Error executing template: ", err))
	}

	return result.String()
}
