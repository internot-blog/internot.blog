package main

import (
	"fmt"

	"github.com/internot-blog/internot.blog.git/internal"
)

func main() {
	cfg := internal.GetConfig("./proompt.json")

	llmPrompt := internal.GenTextPrompt(cfg)
	fmt.Println("Prompt:", llmPrompt)

	response, err := internal.GenText(cfg, llmPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)

	imagePrompt := llmPrompt
	width := 512
	height := 512
	steps := 150
	guidance := 7.5

	imageData, err := internal.GenImage(cfg, imagePrompt, width, height, steps, guidance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	internal.SaveImage(imageData, "generated_image.png")
}
