package main

import (
	"fmt"

	"github.com/internot-blog/internot.blog.git/internal"
)

func main() {
	promptConfig := internal.ReadPromptFile("./proompt.json")
	_ = promptConfig

	fmt.Println("")
}
