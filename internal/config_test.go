package internal_test

import (
	"fmt"
	"testing"

	"github.com/internot-blog/internot.blog.git/internal"
)

func TestReadPromptFile(t *testing.T) {
	promptConfig := internal.ReadPromptFile("../proompt.json")
	fmt.Println(len(promptConfig.TextPromptAction), len(promptConfig.ImagePromptAction), len(promptConfig.TextPromptSubject), len(promptConfig.ImagePromptSubject))
}
