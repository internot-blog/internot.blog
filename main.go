package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/internot-blog/internot.blog.git/internal"
)

func MakeText(cfg internal.Config, textFilename string, prompt string) {
	fmt.Println("Prompt:", prompt)

	response, err := internal.GenText(cfg, prompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	internal.SaveText(response, textFilename)
	fmt.Println(response)
}

func MakeImage(cfg internal.Config, imageFilename string, prompt string, seed int64) {
	width := 512
	height := 512
	steps := 10
	guidance := 7.5

	imageData, err := internal.GenImage(cfg, prompt, seed, width, height, steps, guidance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	internal.SaveImage(imageData, imageFilename)
	fmt.Println("Image saved.")
}

func MakeLogos(cfg internal.Config, prompt string) {
	err := os.MkdirAll("logos", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	for i := 0; i < 100; i++ {
		seed := rand.Int63()
		MakeImage(cfg, fmt.Sprintf("logos/%d.png", seed), prompt, seed)
	}
}

func main() {
	// 1. Load config
	cfg := internal.GetConfig("./proompt.json")

	// 2. Generate post ID

	postId := internal.GenUniqueId()
	postDir := "posts/" + postId

	err := os.MkdirAll(postDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// 3. Generate post contents

	prompt := internal.GenTextPrompt(cfg)
	MakeText(cfg, postDir+"/post.md", prompt)
	seed := rand.Int63()
	MakeImage(cfg, fmt.Sprintf("%s/image.%d.png", postDir, seed), prompt, seed)

	// Site logo generator

	// prompt := "dead internet theory blog website logo named 'internot.blog'"
	// MakeLogos(cfg, prompt)
}
