package internal_test

import (
	"fmt"
	"testing"

	"github.com/internot-blog/internot.blog.git/internal"
)

var promptFile string = "../proompt.json"

func TestReadPromptFile(t *testing.T) {
	promptConfig := internal.ReadPromptFile(promptFile)
	fmt.Println(len(promptConfig.TextPromptAction), len(promptConfig.ImagePromptAction), len(promptConfig.TextPromptSubject), len(promptConfig.ImagePromptSubject))
}

func TestGetModelConfig(t *testing.T) {
	internal.GetModelConfig()
}

func TestGetConfig(t *testing.T) {
	internal.GetConfig(promptFile)
}
